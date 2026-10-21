package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 uint; F1 uint; F2 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8; F1 int32; F2 int32 }
type GenS11 struct { F0 int32 }
func genFn0(p0 uint, p1 uint32, p2 int32) string {
    p0 -= p0
    var v7 uint16 = 789
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("a" + ("foo bar" + "foo bar"))
}
func genFn1(p0 uint, p1 int16) uint {
    var v8 int8 = 42
    if false {
        if (!false) {
            if (!(false || (false || (787 <= -9223372036854775808)))) {
                v8 = (v8 + (v8 | v8))
            } else {
                p1 -= 672
                v9 := GenS1{F0: 32767}
                var v10 float32 = 0.0
                _ = v10
                _ = v9
            }
            p1 -= p1
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return 4294967295
}

func main() {
    v1 := []uint32{2147483648, 2147483648, 65535, 764}
    var v2 bool = true
    v3 := make(map[string]string)
    var v4 float64 = -2.25
    var v5 uint32 = 985
    var v6 float64 = 3.14159
    if len(v1) > 0 {
        v1[0] = v5
    }
    var v11 int16 = 538
    var v12 int32 = 379
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            v6 = v4
            _ = fi14
        }
        _ = fi13
    }
    switch ("" + ("hello" + "foo bar")) {
    case "foo bar":
        switch 4294967295 {
        case 42:
            if true {
                if (false || ((18446744073709551615 >= genFn1(1, -468)) && v2)) {
                    switch "" {
                    case "\n\t":
                        if v2 {
                            v15 := GenS1{F0: 451}
                            if v2 {
                                for fi16 := 0; fi16 < 12; fi16++ {
                                    if (v4 <= -2.25) {
                                        if v2 {
                                            var v17 int64 = 1
                                            for fi18 := 0; fi18 < 2; fi18++ {
                                                fmt.Println(v6)
                                                switch (fi18 * int(v11)) {
                                                case 4294967295:
                                                    v11 += v11
                                                default:
                                                    v2 = (127 > 42)
                                                }
                                                _ = fi18
                                            }
                                            v17 = 1
                                            _ = v17
                                        }
                                        if (v11 > v11) {
                                            var v19 uint32 = v5
                                            _ = v19
                                        }
                                    } else {
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v6)
                                        case "a":
                                            for fi20 := 0; fi20 < 1; fi20++ {
                                                fmt.Println(v11, v6, v2)
                                                v21 := GenS1{F0: 0}
                                                if (0.0 < 0.1) {
                                                    fmt.Println(fi16)
                                                    if v2 {
                                                        switch (-(fi20)) {
                                                        case -28:
                                                            _ = len(v1)
                                                        }
                                                        switch fi20 {
                                                        case 0:
                                                            v22 := GenS1{F0: 1}
                                                            v23 := GenS2{F0: 1, F1: "a"}
                                                            _ = v22
                                                            _ = v23
                                                        }
                                                        if (!(false || (!(!false)))) {
                                                            v2 = v2
                                                            var v24 bool = (9223372036854775807 == -1)
                                                            _ = v24
                                                        } else {
                                                            var v25 int64 = 909
                                                            var v26 int32 = (v12 + (v12 - v12))
                                                            _ = v25
                                                            _ = v26
                                                        }
                                                        switch "hello" {
                                                        case "":
                                                            v27 := GenS3{F0: -1}
                                                            _ = v27
                                                        case "a":
                                                            if false {
                                                                var v28 int64 = 227
                                                                fmt.Println(fi20, v4, v6)
                                                                var v29 float32 = (((1.5 - 1e10) / 1.0) + 0.1)
                                                                _ = v28
                                                                _ = v29
                                                            } else {
                                                                v2 = false
                                                                if v2 {
                                                                    var v30 uint8 = 127
                                                                    if ((v30 * v30) >= v30) {
                                                                        v31 := GenS4{F0: 42, F1: v30}
                                                                        v32 := GenS2{F0: 94, F1: (("hello" + "hello") + "")}
                                                                        _ = v31
                                                                        _ = v32
                                                                    }
                                                                    var v33 uint16 = 546
                                                                    _ = v30
                                                                    _ = v33
                                                                } else {
                                                                    v4 = v4
                                                                    delete(v3, "a")
                                                                    if (!false) {
                                                                        fmt.Println(v5, v11, v4)
                                                                        if (!(v2 || (!false))) {
                                                                            v11 -= (v11 + v11)
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v34 float64 = ((v4 + ((float64((-2.25 - 0.0)) / v6) * (v4 + (0.1 / 1.0)))) * v6)
                                                                            _ = v34
                                                                        } else {
                                                                            v12 = v12
                                                                        }
                                                                        var v35 uint = 18446744073709551615
                                                                        _ = v35
                                                                    }
                                                                    v36 := GenS3{F0: -128}
                                                                    if true {
                                                                        v2 = v2
                                                                    } else {
                                                                        v11 += v11
                                                                        var v37 int64 = -9223372036854775808
                                                                        v2 = (61 < 1)
                                                                        _ = v37
                                                                    }
                                                                    _ = v36
                                                                }
                                                            }
                                                        case "hello":
                                                            _ = len(v1)
                                                        default:
                                                            var v38 uint16 = 32767
                                                            _ = v38
                                                        }
                                                        var v39 uint = 65535
                                                        _ = v39
                                                    }
                                                } else {
                                                    v5 -= (^(v5))
                                                }
                                                _ = fi20
                                                _ = v21
                                            }
                                        case "":
                                            var v40 int32 = v12
                                            _ = v40
                                        default:
                                            if v2 {
                                                for fi41 := 0; fi41 < 20; fi41++ {
                                                    var v42 int8 = 34
                                                    _ = fi41
                                                    _ = v42
                                                }
                                                v11 += (-(v11))
                                            } else {
                                                var v43 string = "a"
                                                _ = v43
                                            }
                                        }
                                        switch 892 {
                                        case 32768:
                                            switch (fi16 >> ((461) & 63)) {
                                            case -2147483648:
                                                v5 = v5
                                            case 927:
                                                for fi44 := 0; fi44 < 2; fi44++ {
                                                    var v45 uint32 = (v5 + 4294967295)
                                                    _ = fi44
                                                    _ = v45
                                                }
                                            case 599:
                                                _ = len(v1)
                                            }
                                        }
                                        v6 += v6
                                    }
                                    delete(v3, ("" + ("\n\t" + "foo bar")))
                                    v11 -= v11
                                    _ = fi16
                                }
                                var v46 int64 = -9223372036854775808
                                _ = v46
                            }
                            if ((255 <= 239) || true) {
                                if v2 {
                                    v47 := GenS5{F0: 1, F1: -1}
                                    v5 = (v5 >> ((255) & 63))
                                    _ = v47
                                } else {
                                    v48 := GenS6{F0: 884, F1: uint(v5), F2: int8(v5)}
                                    _ = v48
                                }
                                for fi49 := 0; fi49 < 19; fi49++ {
                                    for fi50 := 0; fi50 < 13; fi50++ {
                                        var v51 bool = (1 > 18446744073709551615)
                                        switch "hello" {
                                        case "":
                                            var v52 uint64 = 18446744073709551615
                                            _ = v52
                                        case "\n\t":
                                            switch "hello" {
                                            case "hello":
                                                v53 := GenS7{F0: 609}
                                                _ = v53
                                            }
                                        case "foo bar":
                                            if (!false) {
                                                var v54 uint8 = 42
                                                _ = v54
                                            }
                                        default:
                                            if false {
                                                v55 := GenS4{F0: -1, F1: 128}
                                                if v51 {
                                                    if v2 {
                                                        if v51 {
                                                            v56 := GenS8{F0: (^(fi49)), F1: 128}
                                                            v57 := GenS8{F0: fi49, F1: 128}
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        v1 = append(v1, (^(v5)))
                                                    }
                                                    v2 = (true && (0.0 < 1.5))
                                                    var v58 uint16 = 419
                                                    var v59 int64 = 299
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = v55
                                            }
                                        }
                                        _ = len(v1)
                                        _ = fi50
                                        _ = v51
                                    }
                                    if (273 >= uint64(v5)) {
                                        var v60 uint16 = 32767
                                        fmt.Println(v6, v2)
                                        for fi61 := 0; fi61 < 20; fi61++ {
                                            var v62 uint = 65535
                                            v5 -= (v5 >> (((v62 | (^(v62)))) & 63))
                                            v63 := v3[""]
                                            _ = fi61
                                            _ = v62
                                            _ = v63
                                        }
                                        _ = v60
                                    } else {
                                        for fi64 := 0; fi64 < 3; fi64++ {
                                            var v65 int64 = 862
                                            v65 += 417
                                            if v2 {
                                                var v66 int8 = -1
                                                _ = v66
                                            }
                                            _ = fi64
                                            _ = v65
                                        }
                                        switch (fi49 / ((fi49) | 1)) {
                                        case 32767:
                                            for fi67 := 0; fi67 < 11; fi67++ {
                                                v4 -= v6
                                                _ = fi67
                                            }
                                        case 778:
                                            var v68 int = fi49
                                            _ = v68
                                        case 9223372036854775807:
                                            for fi69 := 0; fi69 < 1; fi69++ {
                                                if true {
                                                    if true {
                                                        var v70 uint32 = (^(v5))
                                                        _ = v70
                                                    } else {
                                                        fmt.Println(v12)
                                                        switch "hello" {
                                                        case "hello":
                                                            var v71 uint32 = v5
                                                            _ = v71
                                                        case "\n\t":
                                                            var v72 uint = 4294967295
                                                            _ = v72
                                                        }
                                                        if ((v5 > v5) && (!true)) {
                                                            if true {
                                                                fmt.Println(v6, fi69, v4)
                                                                if true {
                                                                    v5 += (v5 ^ v5)
                                                                    fmt.Println(v5, v11)
                                                                } else {
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        var v73 int16 = v11
                                                                        _ = v73
                                                                    }
                                                                    v4 -= 1.5
                                                                }
                                                                if v2 {
                                                                    v74 := GenS1{F0: 1}
                                                                    var v75 string = ("\n\t" + "")
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                                if (127 >= 42) {
                                                                    fmt.Println(fi49, v11)
                                                                    delete(v3, "")
                                                                } else {
                                                                    if (3.14159 != (3.14159 * 1.5)) {
                                                                        v6 = (((0.0 + v6) - v6) - (((0.0 / 3.14159) - (1.5 / 1.5)) + float64(float32(float64(float32(v4))))))
                                                                    }
                                                                    var v76 int = (fi49 << ((2) & 63))
                                                                    switch fi69 {
                                                                    case -9223372036854775808:
                                                                        if (!((v4 >= (0.1 - (v4 + ((0.0 + 1e10) / 1.0)))) || v2)) {
                                                                            delete(v3, "a")
                                                                            var v77 int = -9223372036854775808
                                                                            v4 += (1.5 * (v6 + ((float64(0.0) - (v4 - v4)) * v4)))
                                                                            _ = v77
                                                                        }
                                                                        if (true || (((1e10 / v6) - 0.1) == (3.14159 * 1e10))) {
                                                                            if v2 {
                                                                                fmt.Println(v5, v5)
                                                                            }
                                                                            var v78 uint16 = 32767
                                                                            _ = v78
                                                                        }
                                                                    }
                                                                    _ = v76
                                                                }
                                                                _ = len(v1)
                                                            } else {
                                                                v6 -= (0.1 - (0.0 * (float64(float32(((v4 + 1.5) * (float64(1e10) + 1e10)))) - (((1e10 * 3.14159) - v6) + v6))))
                                                                var v79 int16 = (v11 % ((v11) | 1))
                                                                _ = v79
                                                            }
                                                            var v80 uint16 = 910
                                                            var v81 bool = ((((v6 * (v4 + (1.5 - float64(3.14159)))) + v6) - 0.0) >= v4)
                                                            _ = v80
                                                            _ = v81
                                                        } else {
                                                            fmt.Println(v4, v6, v4)
                                                            switch (("\n\t" + "hello") + "hello") {
                                                            case "\n\t":
                                                                var v82 uint8 = 255
                                                                _ = v82
                                                            case "hello":
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v83 := v3["hello"]
                                                                    _ = v83
                                                                default:
                                                                    v2 = (v11 > (v11 * v11))
                                                                }
                                                                v3[("hello" + ("" + "foo bar"))] = (("\n\t" + (("foo bar" + (("" + ("" + ("" + ("foo bar" + "\n\t")))) + (("foo bar" + ("\n\t" + "a")) + "hello"))) + ("" + ("foo bar" + "a")))) + ("hello" + ("foo bar" + ("foo bar" + ""))))
                                                            }
                                                            v12 = 858
                                                        }
                                                    }
                                                    if v2 {
                                                        v11 = int16(v12)
                                                    }
                                                    var v84 uint64 = 4294967295
                                                    _ = v84
                                                } else {
                                                    v85 := GenS9{F0: genFn1(0, (v11 & (v11 >> ((817) & 63))))}
                                                    _ = len(v1)
                                                    _ = v85
                                                }
                                                v4 = float64(((1.5 + 0.0) + 0.0))
                                                _ = fi69
                                            }
                                        default:
                                            var v86 uint = 2
                                            _ = v86
                                        }
                                        fmt.Println(v4, v12, v5)
                                    }
                                    v11 = (v11 % ((int16(v5)) | 1))
                                    _ = fi49
                                }
                                for fi87 := 0; fi87 < 7; fi87++ {
                                    _ = len(v1)
                                    var v88 uint8 = 21
                                    v89 := GenS10{F0: 100, F1: (v12 - (v12 ^ -758)), F2: (v12 - v12)}
                                    _ = fi87
                                    _ = v88
                                    _ = v89
                                }
                            } else {
                                fmt.Println(v2)
                            }
                            _ = v15
                        } else {
                            switch 4294967295 {
                            case -9223372036854775808:
                                for fi90 := 0; fi90 < 20; fi90++ {
                                    switch (("hello" + "") + "a") {
                                    case "":
                                        v2 = (42 < 2)
                                    }
                                    _ = fi90
                                }
                            case -1:
                                v91 := v3[(("foo bar" + "\n\t") + "\n\t")]
                                _ = v91
                            case 9223372036854775807:
                                var v92 bool = (true && false)
                                _ = v92
                            default:
                                fmt.Println(v4, v4, v11)
                            }
                            for fi93 := 0; fi93 < 10; fi93++ {
                                var v94 bool = v2
                                v95 := v3[("" + ("\n\t" + ("a" + "hello")))]
                                v96 := GenS7{F0: (v12 ^ (-(v12)))}
                                _ = fi93
                                _ = v94
                                _ = v95
                                _ = v96
                            }
                        }
                    case "a":
                        if v2 {
                            var v97 uint32 = (^(v5))
                            var v98 int64 = 42
                            _ = v97
                            _ = v98
                        }
                    case "":
                        _ = len(v3)
                    default:
                        if v2 {
                            fmt.Println(v4)
                            fmt.Println(v11)
                        }
                    }
                } else {
                    var v99 int64 = -328
                    if true {
                        delete(v3, ("a" + ("foo bar" + "hello")))
                        v3["a"] = ("hello" + "foo bar")
                    }
                    v100 := GenS6{F0: 4294967295, F1: 1, F2: 0}
                    _ = v100
                    _ = v99
                }
            }
        case -93:
            fmt.Println(v4)
        default:
            var v101 float64 = v6
            _ = v101
        }
    case "\n\t":
        fmt.Println(v5, v6)
    case "a":
        if ((0.1 * v6) > 3.14159) {
            var v102 int64 = 9223372036854775807
            v5 -= v5
            for fi103 := 0; fi103 < 9; fi103++ {
                delete(v3, "hello")
                _ = fi103
            }
            _ = v102
        } else {
            v11 = v11
            v104 := GenS2{F0: 2, F1: ("a" + ((((("a" + "a") + "hello") + "a") + "foo bar") + "\n\t"))}
            _ = v104
        }
    }
    var v105 float32 = 1.5
    for fi106 := 0; fi106 < 15; fi106++ {
        v5 = v5
        _ = fi106
    }
    if v2 {
        if len(v1) > 0 {
            v1[0] = 0
        }
        if (!v2) {
            if ((true && v2) && ((v11 > -676) && false)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v107 uint16 = 32767
                _ = v107
            } else {
                switch (((("" + ("hello" + "\n\t")) + "foo bar") + "\n\t") + "\n\t") {
                case "foo bar":
                    v5 = (^(v5))
                }
                for fi108 := 0; fi108 < 14; fi108++ {
                    v109 := GenS11{F0: (v12 ^ (^(v12)))}
                    _ = fi108
                    _ = v109
                }
            }
            var v110 uint = 4294967295
            _ = v110
        } else {
            v2 = v2
        }
    }
    _ = len(v1)
    v12 += -503
    var v111 uint16 = 174
    switch 511 {
    case -180:
        var v112 uint8 = 232
        _ = v112
    case 0:
        v12 = -1
    case 457:
        v4 = v6
    }
    var v113 int32 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v105:", v105)
    fmt.Println("v11:", v11)
    fmt.Println("v111:", v111)
    fmt.Println("v113:", v113)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v105
    _ = v11
    _ = v111
    _ = v113
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
