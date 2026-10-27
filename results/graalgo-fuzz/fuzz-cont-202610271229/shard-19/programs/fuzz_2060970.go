package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 int8; F2 uint64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint8; F1 uint32; F2 int8 }

func main() {
    v1 := []uint32{2147483647, 1, 2147483648, 0}
    v2 := make(map[uint32]int)
    var v3 bool = false
    v4 := v2[2147483647]
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int8 = 42
        v7 := GenS1{F0: -341}
        delete(v2, 598)
        _ = fi5
        _ = v6
        _ = v7
    }
    if (1 != 42) {
        for fi8 := 0; fi8 < 13; fi8++ {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint8 = 2
                v3 = (0 > 1)
                _ = fi9
                _ = v10
            }
            fmt.Println(v4, v3, v4)
            _ = fi8
        }
        switch -510 {
        case 1:
            if (0 < 127) {
                v3 = v3
                for fi11 := 0; fi11 < 14; fi11++ {
                    v12 := GenS1{F0: 258}
                    _ = fi11
                    _ = v12
                }
            }
        default:
            var v13 int64 = -1
            _ = v13
        }
        fmt.Println(v4, v4)
    }
    var v14 int8 = 52
    switch (("hello" + ((("\n\t" + (("" + "a") + "\n\t")) + "hello") + "hello")) + "foo bar") {
    case "a":
        switch "a" {
        case "hello":
            var v15 uint = 763
            _ = v15
        case "a":
            v16 := GenS2{F0: 18446744073709551615, F1: 65535}
            _ = v16
        default:
            var v17 uint32 = 2
            _ = v17
        }
    case "":
        if false {
            if v3 {
                var v18 float32 = 3.14159
                fmt.Println(v3)
                _ = v18
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                v20 := GenS3{F0: 486}
                if false {
                    var v21 int = (fi19 >> ((65535) & 63))
                    var v22 float32 = ((1.5 / float32(1e10)) + 1.5)
                    _ = v21
                    _ = v22
                } else {
                    if len(v1) > 0 {
                        v1[0] = 760
                    }
                    if (2 <= 111) {
                        for fi23 := 0; fi23 < 10; fi23++ {
                            var v24 bool = (!(!(0 < -1)))
                            v4 -= -582
                            _ = fi23
                            _ = v24
                        }
                    }
                }
                _ = fi19
                _ = v20
            }
            switch "a" {
            case "\n\t":
                var v25 float64 = float64((1e10 + -2.25))
                _ = v25
            case "":
                switch int(v14) {
                case 2147483647:
                    var v26 float64 = (1.5 - -2.25)
                    _ = v26
                case -9223372036854775808:
                    if v3 {
                        v27 := GenS3{F0: 376}
                        v4 += v4
                        var v28 uint32 = 4294967295
                        _ = v27
                        _ = v28
                    }
                case 42:
                    switch "" {
                    case "a":
                        switch "\n\t" {
                        case "foo bar":
                            if (42 == 2) {
                                v3 = v3
                                var v29 int = v4
                                for fi30 := 0; fi30 < 19; fi30++ {
                                    fmt.Println(v4, v29)
                                    for fi31 := 0; fi31 < 13; fi31++ {
                                        switch 32767 {
                                        case -816:
                                            switch 0 {
                                            case 42:
                                                fmt.Println(v14, v4)
                                            }
                                        case 1:
                                            fmt.Println(v29)
                                        case 386:
                                            delete(v2, 327)
                                        default:
                                            var v32 int64 = 9223372036854775807
                                            _ = v32
                                        }
                                        v33 := GenS4{F0: 0, F1: 18446744073709551615}
                                        switch fi30 {
                                        case -856:
                                            if ((1.5 - 3.14159) >= 3.14159) {
                                                if (1 == 1) {
                                                    if len(v1) > 0 {
                                                        v1[0] = 346
                                                    }
                                                    switch (v29 << ((509) & 63)) {
                                                    case 32768:
                                                        var v34 bool = false
                                                        _ = v34
                                                    case 0:
                                                        switch v29 {
                                                        case 32768:
                                                            var v35 uint8 = 42
                                                            _ = v35
                                                        case 9223372036854775807:
                                                            fmt.Println(v29)
                                                        case -9223372036854775808:
                                                            if ((1 <= 63) && v3) {
                                                                var v36 string = (("hello" + (("foo bar" + "hello") + ("a" + ("foo bar" + "")))) + ("" + (("hello" + ("\n\t" + ("foo bar" + ("\n\t" + "\n\t")))) + "hello")))
                                                                _ = v36
                                                            }
                                                        }
                                                    case -111:
                                                        switch (("" + "hello") + "a") {
                                                        case "":
                                                            var v37 float32 = -2.25
                                                            _ = v37
                                                        case "hello":
                                                            if (((-2.25 - -2.25) / 1.0) >= 1.5) {
                                                                var v38 uint32 = 1
                                                                _ = v38
                                                            }
                                                        case "foo bar":
                                                            v3 = (32767 <= 9223372036854775807)
                                                        }
                                                    }
                                                    if (v3 && v3) {
                                                        v39 := GenS5{F0: 0}
                                                        switch 0 {
                                                        case 888:
                                                            v40 := v2[2147483648]
                                                            _ = v40
                                                        case 4294967295:
                                                            v4 += (v4 % ((fi31) | 1))
                                                        default:
                                                            switch 1 {
                                                            case 42:
                                                                v29 = 32767
                                                            case -9223372036854775808:
                                                                var v41 uint64 = 2
                                                                _ = v41
                                                            }
                                                        }
                                                        fmt.Println(v3, fi31, fi31)
                                                        _ = v39
                                                    } else {
                                                        var v42 uint64 = 0
                                                        _ = v42
                                                    }
                                                    if v3 {
                                                        v43 := GenS2{F0: 0, F1: 576}
                                                        _ = v43
                                                    }
                                                    v4 = 9223372036854775807
                                                    var v44 uint16 = 416
                                                    _ = v44
                                                }
                                            }
                                        }
                                        v2[2147483648] = fi30
                                        if false {
                                            v1 = append(v1, 2)
                                            _ = len(v2)
                                            v45 := GenS4{F0: 0, F1: 4294967295}
                                            _ = v45
                                        }
                                        if (-654 <= 1) {
                                            fmt.Println(v4, v14)
                                            var v46 bool = v3
                                            _ = v46
                                        } else {
                                            var v47 string = "\n\t"
                                            v48 := GenS5{F0: 1}
                                            if len(v1) > 0 {
                                                v1[0] = 2147483647
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                        var v49 string = ("" + ("" + ("\n\t" + "foo bar")))
                                        _ = fi31
                                        _ = v33
                                        _ = v49
                                    }
                                    if true {
                                        var v50 uint64 = 1
                                        var v51 int64 = 0
                                        _ = v50
                                        _ = v51
                                    } else {
                                        v14 = v14
                                        if false {
                                            fmt.Println(v3, v4)
                                            var v52 uint32 = 861
                                            v53 := v2[v52]
                                            _ = v52
                                            _ = v53
                                        } else {
                                            for fi54 := 0; fi54 < 6; fi54++ {
                                                if ((((3.14159 + -2.25) / ((1e10 - 0.1) + 1.5)) != 0.1) || true) {
                                                    switch (("" + ("" + "a")) + "hello") {
                                                    case "\n\t":
                                                        var v55 bool = false
                                                        _ = v55
                                                    case "a":
                                                        var v56 int64 = 9223372036854775807
                                                        _ = v56
                                                    }
                                                    if v3 {
                                                        var v57 int = (-(fi54))
                                                        var v58 uint = 2
                                                        if v3 {
                                                            fmt.Println(v57)
                                                        } else {
                                                            fmt.Println(v14, v4)
                                                            var v59 int64 = 1
                                                            var v60 int = (v4 ^ (v29 + 4294967295))
                                                            if v3 {
                                                                fmt.Println(v14, v60)
                                                            }
                                                            var v61 uint8 = 152
                                                            _ = v59
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            fmt.Println(v14, v29, v29)
                                                        case "a":
                                                            if true {
                                                                v4 = -2147483648
                                                                v62 := GenS6{F0: 635, F1: 248}
                                                                _ = v62
                                                            } else {
                                                                var v63 uint8 = 1
                                                                _ = v63
                                                            }
                                                        case "":
                                                            var v64 float32 = 3.14159
                                                            _ = v64
                                                        default:
                                                            v3 = (!(755 > 18446744073709551615))
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v4 -= (v29 - v4)
                                                    }
                                                    v65 := GenS7{F0: 1}
                                                    switch v29 {
                                                    case -2147483648:
                                                        if ((1 > 80) && (v14 >= v14)) {
                                                            var v66 string = "foo bar"
                                                            _ = v66
                                                        }
                                                    default:
                                                        if (1.5 < 1e10) {
                                                            if true {
                                                                delete(v2, 406)
                                                                var v67 uint16 = 42
                                                                var v68 uint64 = 735
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                        }
                                                    }
                                                    v69 := v2[4294967295]
                                                    if len(v1) > 0 {
                                                        v1[0] = 2
                                                    }
                                                    v3 = v3
                                                    v69 += (v29 % ((v69) | 1))
                                                    _ = v65
                                                    _ = v69
                                                }
                                                _ = fi54
                                            }
                                        }
                                    }
                                    _ = fi30
                                }
                                _ = v29
                            } else {
                                var v70 float32 = 0.1
                                _ = v70
                            }
                        case "hello":
                            var v71 float32 = (3.14159 * (float32(1e10) * (0.0 * 1.5)))
                            _ = v71
                        case "\n\t":
                            v72 := GenS6{F0: 523, F1: 29}
                            _ = v72
                        }
                    case "hello":
                        var v73 int = 259
                        _ = v73
                    case "\n\t":
                        switch v4 {
                        case 0:
                            v3 = false
                        case 679:
                            if v3 {
                                var v74 float64 = (0.0 * 3.14159)
                                if len(v1) > 0 {
                                    v1[0] = 1
                                }
                                v75 := GenS4{F0: 17, F1: 18446744073709551615}
                                _ = v74
                                _ = v75
                            } else {
                                if v3 {
                                    for fi76 := 0; fi76 < 14; fi76++ {
                                        fmt.Println(v3)
                                        _ = fi76
                                    }
                                }
                                if (!false) {
                                    fmt.Println(v3)
                                    if true {
                                        v77 := GenS5{F0: -32768}
                                        fmt.Println(v14, v14)
                                        v14 -= v14
                                        _ = v77
                                    }
                                    for fi78 := 0; fi78 < 13; fi78++ {
                                        _ = len(v2)
                                        _ = fi78
                                    }
                                }
                            }
                        case 42:
                            if ((0.0 * ((3.14159 + 3.14159) + 3.14159)) > (0.0 * (1e10 * 3.14159))) {
                                var v79 float64 = 3.14159
                                var v80 int16 = 1000
                                _ = v79
                                _ = v80
                            } else {
                                v4 -= (v4 << ((0) & 63))
                                switch "\n\t" {
                                case "foo bar":
                                    v81 := GenS6{F0: 65535, F1: 227}
                                    _ = v81
                                case "a":
                                    var v82 uint16 = 430
                                    _ = v82
                                }
                            }
                        default:
                            v4 -= 244
                        }
                    }
                }
            }
        } else {
            for fi83 := 0; fi83 < 20; fi83++ {
                v14 -= 88
                v14 += 85
                _ = fi83
            }
            v14 = (^(v14))
            if true {
                var v84 uint16 = 42
                v3 = (v3 && (!v3))
                if (110 == 202) {
                    v14 -= (v14 * (v14 ^ (v14 << ((18446744073709551615) & 63))))
                    fmt.Println(v14, v3)
                } else {
                    v85 := GenS8{F0: 42, F1: v14, F2: 42}
                    if (false || true) {
                        for fi86 := 0; fi86 < 20; fi86++ {
                            v87 := GenS9{F0: -87}
                            _ = fi86
                            _ = v87
                        }
                        fmt.Println(v3)
                    } else {
                        _ = len(v1)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    for fi88 := 0; fi88 < 10; fi88++ {
                        var v89 float32 = (float32((-2.25 / 1e10)) - 1e10)
                        _ = fi88
                        _ = v89
                    }
                    _ = v85
                }
                _ = v84
            }
        }
    case "\n\t":
        v90 := GenS10{F0: 255, F1: 2147483648, F2: (-(v14))}
        _ = v90
    default:
        var v91 uint64 = 0
        _ = v91
    }
    v14 += v14
    delete(v2, 2147483647)
    for fi92 := 0; fi92 < 2; fi92++ {
        fmt.Println(fi92, fi92)
        _ = fi92
    }
    var v93 int32 = -599
    v94 := GenS4{F0: 42, F1: 326}
    if (42 > 18446744073709551615) {
        if (!((!v3) || v3)) {
            switch (v4 / (((v4 >> ((42) & 63))) | 1)) {
            case 42:
                var v95 int32 = v93
                _ = v95
            default:
                switch (v4 | int(v14)) {
                case 9223372036854775807:
                    v3 = true
                default:
                    var v96 int64 = 9223372036854775807
                    _ = v96
                }
            }
            if false {
                v4 -= v4
                v97 := GenS1{F0: 1000}
                _ = v97
            } else {
                fmt.Println(v3, v14, v93)
                v1 = append(v1, 65535)
            }
            switch (v4 * (v4 + (^(v4)))) {
            case -2147483648:
                var v98 int64 = -9223372036854775808
                _ = v98
            case 2147483647:
                _ = len(v2)
            default:
                v99 := v2[2]
                _ = v99
            }
        } else {
            v4 -= -512
            var v100 int32 = v93
            _ = v100
        }
    }
    var v101 uint64 = 68
    fmt.Println(v14, v93, v14)
    switch ("a" + (("\n\t" + "foo bar") + "foo bar")) {
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = 0
        }
    case "hello":
        v1 = append(v1, 65535)
    case "":
        var v102 string = "hello"
        _ = v102
    default:
        var v103 int8 = v14
        _ = v103
    }
    v104 := GenS2{F0: (v101 << ((65535) & 63)), F1: 255}
    v93 -= (v93 % (((v93 * (v93 >> ((737) & 63)))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v101
    _ = v104
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v93
    _ = v94
}
