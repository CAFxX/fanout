package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint32; F1 int; F2 uint32 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint32; F1 string; F2 int32 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 bool; F1 int; F2 uint64 }
type GenS8 struct { F0 uint64; F1 int; F2 uint64 }
type GenS9 struct { F0 int8; F1 uint32 }
type GenS10 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS11 struct { F0 int8 }
func genFn0(p0 int32, p1 bool) string {
    p0 += p0
    v5 := []int{128, 261, 127, 4294967295}
    for fi6 := 0; fi6 < 1; fi6++ {
        p0 -= p0
        _ = fi6
    }
    var v7 uint64 = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return "hello"
}

func main() {
    v1 := []uint32{4294967295, 2, 633}
    var v2 int = 741
    var v3 int32 = 155
    var v4 uint32 = 447
    v3 = 0
    if true {
        _ = len(v1)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v2 -= v2
    if true {
        v4 = (v4 + 2147483647)
    } else {
        if false {
            var v8 int32 = (v3 % ((int32(v2)) | 1))
            v3 = 12
            _ = v8
        }
        v4 = v4
        switch "foo bar" {
        case "":
            if (false || true) {
                switch ((("" + "\n\t") + "hello") + "a") {
                case "hello":
                    if (!(true && false)) {
                        switch (("a" + ("a" + "\n\t")) + "\n\t") {
                        case "hello":
                            for fi9 := 0; fi9 < 16; fi9++ {
                                v1 = append(v1, v4)
                                var v10 float32 = ((0.0 / 0.1) * (-2.25 / 1.5))
                                _ = fi9
                                _ = v10
                            }
                        case "":
                            var v11 int16 = -532
                            _ = v11
                        case "foo bar":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        switch (v2 / ((127) | 1)) {
                        case 32768:
                            switch v2 {
                            case 128:
                                if ("a" != (("hello" + "hello") + "\n\t")) {
                                    fmt.Println(v2)
                                } else {
                                    v12 := GenS1{F0: 570, F1: (v3 / ((-2147483648) | 1))}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v1 = append(v1, v4)
                                    _ = v12
                                }
                            default:
                                var v13 uint32 = (^(v4))
                                _ = v13
                            }
                        case 32767:
                            if false {
                                var v14 int8 = 47
                                if (false || true) {
                                    for fi15 := 0; fi15 < 16; fi15++ {
                                        v16 := GenS2{F0: (^(v4)), F1: fi15, F2: v4}
                                        for fi17 := 0; fi17 < 13; fi17++ {
                                            switch -1 {
                                            case -573:
                                                if false {
                                                    fmt.Println(v4)
                                                    var v18 uint64 = 390
                                                    v4 = v4
                                                    if false {
                                                        v4 -= (v4 | 2)
                                                    } else {
                                                        v19 := GenS3{F0: (^(v14)), F1: 1000}
                                                        fmt.Println(v4)
                                                        _ = v19
                                                    }
                                                    _ = v18
                                                } else {
                                                    var v20 string = ("foo bar" + ((((("a" + "foo bar") + "hello") + "a") + "a") + ""))
                                                    v1 = append(v1, (v4 >> ((1) & 63)))
                                                    _ = v20
                                                }
                                            case 0:
                                                v4 = (v4 | 2147483648)
                                            case 128:
                                                var v21 float32 = ((0.1 + 3.14159) - (1.5 * (1.5 + 0.1)))
                                                _ = v21
                                            default:
                                                var v22 uint = 4294967295
                                                _ = v22
                                            }
                                            _ = fi17
                                        }
                                        _ = fi15
                                        _ = v16
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v23 := GenS4{F0: (v4 ^ (v4 / ((v4) | 1))), F1: "", F2: (v3 << ((0) & 63))}
                                    _ = v23
                                } else {
                                    v3 = 2147483647
                                    v14 = (v14 ^ v14)
                                }
                                var v24 uint = 42
                                _ = v14
                                _ = v24
                            }
                        default:
                            v3 += -2147483648
                        }
                        var v25 uint = 0
                        _ = v25
                    } else {
                        switch -140 {
                        case -888:
                            v26 := GenS5{F0: 1, F1: 15}
                            _ = v26
                        case -9223372036854775808:
                            switch v2 {
                            case -1:
                                var v27 int64 = 722
                                _ = v27
                            case -2147483648:
                                var v28 uint64 = 2
                                _ = v28
                            case 32767:
                                for fi29 := 0; fi29 < 2; fi29++ {
                                    if len(v1) > 0 {
                                        v1[0] = 65535
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi29
                                }
                            }
                        }
                        v30 := GenS1{F0: 2, F1: v3}
                        _ = v30
                    }
                }
            }
        case "a":
            if false {
                var v31 uint32 = (v4 | (v4 >> ((557) & 63)))
                v1 = append(v1, v31)
                var v32 uint = 360
                _ = v31
                _ = v32
            }
        case "foo bar":
            v33 := GenS4{F0: v4, F1: "", F2: v3}
            _ = v33
        default:
            _ = len(v1)
        }
    }
    fmt.Println(v4)
    v3 = 172
    switch (v2 - (v2 & v2)) {
    case 128:
        v2 += v2
    case 2147483647:
        if true {
            fmt.Println(v2, v2)
        }
    }
    _ = len(v1)
    if (467 <= 297) {
        var v34 float64 = (1.5 + 0.1)
        var v35 uint64 = 0
        if true {
            var v36 int16 = 32767
            var v37 int = v2
            v4 -= (v4 - (v4 + (^(v4))))
            _ = v36
            _ = v37
        }
        _ = v34
        _ = v35
    }
    if len(v1) > 0 {
        v1[0] = v4
    }
    if false {
        var v38 uint32 = (v4 ^ v4)
        var v39 uint64 = 18446744073709551615
        _ = v38
        _ = v39
    } else {
        if (false && false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 -= (v2 + -2147483648)
            if false {
                if (!(v3 >= (v3 << ((18446744073709551615) & 63)))) {
                    var v40 uint8 = 107
                    for fi41 := 0; fi41 < 17; fi41++ {
                        _ = len(v1)
                        _ = fi41
                    }
                    _ = v40
                } else {
                    for fi42 := 0; fi42 < 12; fi42++ {
                        switch "\n\t" {
                        case "foo bar":
                            v4 -= 369
                        case "\n\t":
                            if (18446744073709551615 <= 879) {
                                var v43 int64 = -130
                                _ = v43
                            } else {
                                var v44 int = fi42
                                _ = v44
                            }
                        case "a":
                            var v45 string = "\n\t"
                            _ = v45
                        }
                        _ = len(v1)
                        switch "hello" {
                        case "a":
                            v2 -= int(v3)
                        default:
                            v2 = 612
                        }
                        _ = fi42
                    }
                    if (!false) {
                        v46 := GenS3{F0: 0, F1: 1}
                        v47 := GenS6{F0: true, F1: 42}
                        if (false && (false || false)) {
                            for fi48 := 0; fi48 < 18; fi48++ {
                                if len(v1) > 0 {
                                    v1[0] = (v4 % ((v4) | 1))
                                }
                                if false {
                                    for fi49 := 0; fi49 < 18; fi49++ {
                                        if (!true) {
                                            if (("" <= "a") || true) {
                                                if true {
                                                    v4 += 588
                                                    switch 32768 {
                                                    case 9223372036854775807:
                                                        var v50 string = "hello"
                                                        _ = v50
                                                    case 1:
                                                        var v51 uint64 = 4294967295
                                                        var v52 uint64 = 4294967295
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    v3 = v3
                                                }
                                            } else {
                                                v53 := GenS3{F0: -36, F1: 1}
                                                var v54 int32 = (v3 | v3)
                                                _ = v53
                                                _ = v54
                                            }
                                        } else {
                                            if (!false) {
                                                v3 = v3
                                                fmt.Println(fi48)
                                            } else {
                                                switch -2147483648 {
                                                case 128:
                                                    v1 = append(v1, v4)
                                                case 4294967295:
                                                    v55 := GenS7{F0: true, F1: fi49, F2: 2}
                                                    var v56 float64 = 3.14159
                                                    switch "a" {
                                                    case "a":
                                                        v1 = append(v1, v4)
                                                        v57 := GenS4{F0: v4, F1: ("a" + "\n\t"), F2: v3}
                                                        var v58 int8 = 0
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    _ = v55
                                                    _ = v56
                                                default:
                                                    v2 = fi48
                                                }
                                                _ = len(v1)
                                                switch ("a" + "hello") {
                                                case "":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "hello":
                                                    var v59 int8 = 42
                                                    _ = v59
                                                default:
                                                    fmt.Println(v2, fi48, fi49)
                                                }
                                            }
                                        }
                                        var v60 uint32 = v4
                                        _ = fi49
                                        _ = v60
                                    }
                                    var v61 string = "\n\t"
                                    _ = v61
                                } else {
                                    if true {
                                        var v62 int16 = int16(v4)
                                        _ = v62
                                    }
                                    v3 = v3
                                    if (69 >= 1) {
                                        for fi63 := 0; fi63 < 7; fi63++ {
                                            fmt.Println(fi63, v3)
                                            _ = fi63
                                        }
                                    }
                                }
                                _ = fi48
                            }
                            for fi64 := 0; fi64 < 8; fi64++ {
                                v65 := GenS6{F0: true, F1: 18446744073709551615}
                                if true {
                                    var v66 int16 = 42
                                    if true {
                                        if false {
                                            var v67 int16 = v66
                                            v4 = v4
                                            if false {
                                                if (false && false) {
                                                    var v68 uint64 = 2
                                                    _ = v68
                                                }
                                                var v69 int16 = v66
                                                _ = v69
                                            }
                                            _ = v67
                                        }
                                        var v70 uint8 = 147
                                        var v71 int64 = int64(v66)
                                        _ = v70
                                        _ = v71
                                    } else {
                                        if true {
                                            v72 := GenS8{F0: 4294967295, F1: (v2 << ((4294967295) & 63)), F2: 174}
                                            _ = v72
                                        } else {
                                            v2 = -2147483648
                                            if ((0.1 - 3.14159) == float32((-2.25 / float64(float32(((0.0 / 1.0) + 0.1)))))) {
                                                var v73 float64 = 1.5
                                                var v74 uint32 = v4
                                                _ = v73
                                                _ = v74
                                            } else {
                                                for fi75 := 0; fi75 < 20; fi75++ {
                                                    v66 += v66
                                                    v2 = 0
                                                    if (1000 < (v66 % ((1) | 1))) {
                                                        v76 := GenS1{F0: 0, F1: (v3 * v3)}
                                                        fmt.Println(fi75, v2, fi64)
                                                        if false {
                                                            switch (("\n\t" + "\n\t") + (("\n\t" + "\n\t") + "\n\t")) {
                                                            case "":
                                                                _ = len(v1)
                                                            case "\n\t":
                                                                if len(v1) > 0 {
                                                                    v1[0] = v4
                                                                }
                                                            default:
                                                                var v77 float64 = (0.0 * (1e10 + float64(1e10)))
                                                                _ = v77
                                                            }
                                                        }
                                                        _ = v76
                                                    }
                                                    var v78 int64 = 9223372036854775807
                                                    var v79 uint = 601
                                                    var v80 int = -9223372036854775808
                                                    _ = fi75
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                                var v81 uint64 = 224
                                                _ = v81
                                            }
                                        }
                                        v4 = 65535
                                    }
                                    _ = v66
                                } else {
                                    fmt.Println(v4, v2)
                                }
                                v82 := GenS2{F0: v4, F1: fi64, F2: (v4 | (^(v4)))}
                                _ = fi64
                                _ = v65
                                _ = v82
                            }
                            v4 += (v4 & v4)
                        } else {
                            var v83 float64 = float64((-2.25 + 0.1))
                            _ = v83
                        }
                        _ = v46
                        _ = v47
                    } else {
                        switch "\n\t" {
                        case "hello":
                            var v84 uint8 = 104
                            _ = v84
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "foo bar":
                            fmt.Println(v3)
                        }
                        if true {
                            v85 := GenS9{F0: -128, F1: (v4 >> ((18446744073709551615) & 63))}
                            _ = v85
                        } else {
                            for fi86 := 0; fi86 < 15; fi86++ {
                                if true {
                                    v87 := GenS3{F0: -49, F1: 32767}
                                    switch "a" {
                                    case "\n\t":
                                        var v88 int32 = v3
                                        _ = v88
                                    case "foo bar":
                                        var v89 uint = 18446744073709551615
                                        _ = v89
                                    default:
                                        switch (fi86 & (v2 & 0)) {
                                        case -555:
                                            v4 = (v4 % (((v4 % (((v4 % (((v4 | v4)) | 1))) | 1))) | 1))
                                        case 4294967295:
                                            for fi90 := 0; fi90 < 17; fi90++ {
                                                switch ("" + "foo bar") {
                                                case "":
                                                    fmt.Println(fi90)
                                                case "\n\t":
                                                    if len(v1) > 0 {
                                                        v1[0] = (v4 + (v4 << ((4294967295) & 63)))
                                                    }
                                                case "a":
                                                    v2 += -2147483648
                                                default:
                                                    var v91 int16 = -1
                                                    _ = v91
                                                }
                                                _ = fi90
                                            }
                                        default:
                                            if (-762 >= fi86) {
                                                if (217 == 247) {
                                                    if false {
                                                        _ = len(v1)
                                                        if false {
                                                            var v92 int64 = 741
                                                            _ = v92
                                                        } else {
                                                            var v93 float64 = 1.5
                                                            _ = len(v1)
                                                            _ = v93
                                                        }
                                                    }
                                                }
                                            } else {
                                                v2 = fi86
                                                var v94 int = v2
                                                _ = v94
                                            }
                                        }
                                    }
                                    _ = v87
                                } else {
                                    if (!(!("hello" > "\n\t"))) {
                                        var v95 uint8 = 255
                                        if true {
                                            if len(v1) > 0 {
                                                v1[0] = (v4 ^ (v4 + v4))
                                            }
                                        } else {
                                            fmt.Println(fi86)
                                            var v96 uint32 = 751
                                            _ = v96
                                        }
                                        _ = v95
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    for fi97 := 0; fi97 < 1; fi97++ {
                                        var v98 int64 = 1
                                        v98 += (-(v98))
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = fi97
                                        _ = v98
                                    }
                                }
                                for fi99 := 0; fi99 < 3; fi99++ {
                                    v100 := GenS10{F0: 32767, F1: 910, F2: 411}
                                    var v101 int32 = 202
                                    v3 += v3
                                    _ = fi99
                                    _ = v100
                                    _ = v101
                                }
                                _ = fi86
                            }
                        }
                        if (!true) {
                            v102 := GenS11{F0: -1}
                            fmt.Println(v3, v3, v4)
                            if (32767 <= 1000) {
                                for fi103 := 0; fi103 < 5; fi103++ {
                                    v2 = (fi103 / ((v2) | 1))
                                    _ = fi103
                                }
                                v2 -= v2
                            }
                            _ = v102
                        }
                    }
                    if false {
                        v4 -= v4
                        v1 = append(v1, v4)
                        switch ("a" + "hello") {
                        case "a":
                            switch (v2 ^ (v2 - (-(v2)))) {
                            case -498:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                if len(v1) > 0 {
                                    v1[0] = (^(v4))
                                }
                            }
                        case "foo bar":
                            var v104 uint64 = 2
                            _ = v104
                        case "\n\t":
                            var v105 float64 = (0.0 / 1e10)
                            _ = v105
                        default:
                            v3 += v3
                        }
                    } else {
                        fmt.Println(v4, v3)
                    }
                }
                v2 = 2147483647
                var v106 int32 = 1
                _ = v106
            } else {
                var v107 bool = (42 == 127)
                _ = v107
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
