package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 int; F1 uint64; F2 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 string; F2 string }
type GenS9 struct { F0 bool; F1 uint }
type GenS10 struct { F0 string; F1 int64; F2 string }
type GenS11 struct { F0 bool; F1 uint }
type GenS12 struct { F0 int; F1 int8; F2 string }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint8; F1 uint64 }
type GenS15 struct { F0 bool; F1 uint8; F2 int8 }
type GenS16 struct { F0 uint16; F1 uint }

func main() {
    v1 := []int{127, 42, 0}
    var v2 float64 = ((1.5 + (0.0 / -2.25)) * (float64(float32(0.1)) / float64(3.14159)))
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    v5 := []uint32{2, 65535, 4294967295, 65535}
    v6 := []int{-2147483648, 32768}
    var v7 uint64 = 623
    delete(v4, 1)
    v7 = (v7 >> ((uint(v7)) & 63))
    if (!true) {
        fmt.Println(v7)
        if ((!false) || false) {
            var v8 int = 2147483647
            delete(v3, "foo bar")
            if len(v5) > 0 {
                v5[0] = 326
            }
            _ = v8
        }
        if true {
            if (778 < 1) {
                v7 = v7
                v9 := GenS1{F0: 42}
                v2 = 3.14159
                _ = v9
            } else {
                switch "\n\t" {
                case "":
                    fmt.Println(v7, v7)
                case "foo bar":
                    fmt.Println(v7, v2)
                default:
                    v2 += (-2.25 + 1.5)
                }
                for fi10 := 0; fi10 < 19; fi10++ {
                    switch fi10 {
                    case 4294967295:
                        if len(v5) > 0 {
                            v5[0] = 454
                        }
                    case -384:
                        v11 := GenS1{F0: (fi10 & fi10)}
                        _ = v11
                    default:
                        for fi12 := 0; fi12 < 18; fi12++ {
                            var v13 int32 = 1
                            _ = fi12
                            _ = v13
                        }
                    }
                    switch (^(fi10)) {
                    case 32768:
                        for fi14 := 0; fi14 < 2; fi14++ {
                            v2 = 0.1
                            _ = fi14
                        }
                    default:
                        v7 = (v7 & 45)
                    }
                    var v15 uint64 = (^(v7))
                    _ = fi10
                    _ = v15
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        v18 := GenS2{F0: "hello", F1: (fi16 << ((42) & 63))}
                        fmt.Println(v7)
                        _ = fi17
                        _ = v18
                    }
                    v7 -= v7
                    if (((v2 + v2) / (v2 - (v2 / ((v2 - v2) + 1.5)))) != (v2 + v2)) {
                        v19 := GenS1{F0: 42}
                        fmt.Println(fi16)
                        var v20 bool = false
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int32 = 1
                        for fi22 := 0; fi22 < 9; fi22++ {
                            if true {
                                if (247 >= 234) {
                                    v7 -= (v7 | v7)
                                } else {
                                    if (!(42 != 42)) {
                                        v21 += 1
                                        v23 := GenS3{F0: ""}
                                        _ = v23
                                    } else {
                                        if (127 != -128) {
                                            v24 := GenS2{F0: "", F1: -9223372036854775808}
                                            fmt.Println(v2, v7, v7)
                                            delete(v4, 147)
                                            _ = v24
                                        } else {
                                            v25 := v4[2147483648]
                                            fmt.Println(fi16)
                                            var v26 bool = false
                                            _ = v25
                                            _ = v26
                                        }
                                        v27 := GenS4{F0: v7, F1: -128, F2: 2}
                                        switch (fi22 + (-(fi22))) {
                                        case 9223372036854775807:
                                            v21 -= v21
                                        case 168:
                                            v28 := GenS2{F0: "hello", F1: (fi16 | 1)}
                                            _ = v28
                                        case -8:
                                            v29 := GenS4{F0: 0, F1: -65, F2: 211}
                                            _ = v29
                                        }
                                        _ = v27
                                    }
                                    switch ("foo bar" + "hello") {
                                    case "":
                                        v2 += (float64(0.1) * -2.25)
                                    case "foo bar":
                                        var v30 int16 = 1
                                        _ = v30
                                    }
                                    v7 += (^(v7))
                                    if true {
                                        switch fi22 {
                                        case -1:
                                            v2 -= ((v2 / (v2 * (float64(float32(1e10)) * (v2 - (float64(((1.5 * 0.1) + 1e10)) / v2))))) + (float64((-2.25 * ((1e10 + 1e10) * ((-2.25 * 0.1) - 0.1)))) + 0.0))
                                        default:
                                            if false {
                                                v21 = v21
                                                if true {
                                                    v21 = (v21 * (v21 & v21))
                                                    if false {
                                                        v7 += (v7 >> ((0) & 63))
                                                        v2 = v2
                                                        switch (fi16 >> ((2) & 63)) {
                                                        case 128:
                                                            var v31 uint32 = 42
                                                            _ = v31
                                                        default:
                                                            if (!true) {
                                                                var v32 float32 = 0.1
                                                                v32 += v32
                                                                if (false && (!("" <= "a"))) {
                                                                    var v33 uint32 = 4294967295
                                                                    if false {
                                                                        var v34 int = fi16
                                                                        var v35 float64 = v2
                                                                        var v36 bool = true
                                                                        _ = v34
                                                                        _ = v35
                                                                        _ = v36
                                                                    } else {
                                                                        switch fi16 {
                                                                        case 32767:
                                                                            if (fi16 <= (fi22 % (((^(fi16))) | 1))) {
                                                                                switch (fi16 & -9223372036854775808) {
                                                                                case -9223372036854775808:
                                                                                    v37 := GenS3{F0: "\n\t"}
                                                                                    _ = v37
                                                                                case 9223372036854775807:
                                                                                    var v38 uint = 4294967295
                                                                                    _ = v38
                                                                                }
                                                                            } else {
                                                                                v33 = v33
                                                                                if false {
                                                                                    var v39 bool = false
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = fi16
                                                                                    }
                                                                                    _ = v39
                                                                                } else {
                                                                                    fmt.Println(v21, fi22)
                                                                                    switch ("hello" + "") {
                                                                                    case "a":
                                                                                        v40 := GenS3{F0: "hello"}
                                                                                        _ = v40
                                                                                    case "\n\t":
                                                                                        var v41 uint16 = 65535
                                                                                        _ = v41
                                                                                    case "hello":
                                                                                        if false {
                                                                                            if (!false) {
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = (v33 >> ((42) & 63))
                                                                                                }
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = (^(fi16))
                                                                                                }
                                                                                                fmt.Println(v7)
                                                                                            }
                                                                                        } else {
                                                                                            if false {
                                                                                                if true {
                                                                                                    if (0 < 0) {
                                                                                                        fmt.Println(v32)
                                                                                                        if (9223372036854775807 >= -1) {
                                                                                                            v42 := GenS1{F0: fi16}
                                                                                                            var v43 uint16 = 2
                                                                                                            _ = v42
                                                                                                            _ = v43
                                                                                                        }
                                                                                                        if false {
                                                                                                            fmt.Println(v21, fi22)
                                                                                                            switch fi16 {
                                                                                                            case 1:
                                                                                                                if true {
                                                                                                                    if false {
                                                                                                                        if len(v5) > 0 {
                                                                                                                            v5[0] = v33
                                                                                                                        }
                                                                                                                        fmt.Println(v33, v32)
                                                                                                                        var v44 float64 = 0.1
                                                                                                                        _ = v44
                                                                                                                    } else {
                                                                                                                        switch int(v21) {
                                                                                                                        case 128:
                                                                                                                            var v45 int = fi16
                                                                                                                            _ = v45
                                                                                                                        case 42:
                                                                                                                            v1 = append(v1, fi16)
                                                                                                                        }
                                                                                                                        var v46 int8 = 1
                                                                                                                        _ = v46
                                                                                                                    }
                                                                                                                    if false {
                                                                                                                        v2 = ((0.0 / 1.0) + (((v2 / v2) - v2) * (float64(v32) + -2.25)))
                                                                                                                        v3[("" + "")] = ((!(!false)) && (152 > 928))
                                                                                                                        var v47 uint16 = 505
                                                                                                                        _ = v47
                                                                                                                    }
                                                                                                                    var v48 float32 = v32
                                                                                                                    _ = v48
                                                                                                                }
                                                                                                            case -445:
                                                                                                                v21 = (v21 >> ((65535) & 63))
                                                                                                            case 2147483647:
                                                                                                                v32 = v32
                                                                                                            default:
                                                                                                                if true {
                                                                                                                    fmt.Println(v2)
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v49 uint16 = 255
                                                                                                            _ = v49
                                                                                                        }
                                                                                                    }
                                                                                                    v3[(("" + "a") + "foo bar")] = false
                                                                                                    v50 := GenS2{F0: "hello", F1: (fi16 << ((18446744073709551615) & 63))}
                                                                                                    var v51 int = 9223372036854775807
                                                                                                    _ = v50
                                                                                                    _ = v51
                                                                                                }
                                                                                                var v52 uint16 = 42
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    fmt.Println(v7, v2)
                                                                                                case "foo bar":
                                                                                                    var v53 int8 = -68
                                                                                                    _ = v53
                                                                                                default:
                                                                                                    v33 += (^(v33))
                                                                                                }
                                                                                                _ = v52
                                                                                            }
                                                                                            if true {
                                                                                                v32 += (((3.14159 / ((v32 - -2.25) - (float32(-2.25) - float32((v2 * -2.25))))) * v32) * 1.5)
                                                                                                fmt.Println(v7, fi16, v21)
                                                                                                v32 -= 0.0
                                                                                            }
                                                                                            switch "a" {
                                                                                            case "foo bar":
                                                                                                v54 := GenS5{F0: v33, F1: v21}
                                                                                                _ = v54
                                                                                            }
                                                                                            v32 = (v32 * (0.1 - v32))
                                                                                            if (128 >= int(v33)) {
                                                                                                fmt.Println(v32)
                                                                                                switch (((("hello" + ("\n\t" + ("a" + ("hello" + "hello")))) + "a") + ("a" + "a")) + ("hello" + "\n\t")) {
                                                                                                case "hello":
                                                                                                    v55 := GenS4{F0: 1, F1: 42, F2: 0}
                                                                                                    _ = v55
                                                                                                }
                                                                                                var v56 int32 = v21
                                                                                                fmt.Println(v7, v56)
                                                                                                var v57 int = 32768
                                                                                                _ = v56
                                                                                                _ = v57
                                                                                            } else {
                                                                                                v21 += (v21 | v21)
                                                                                                var v58 float64 = float64(v32)
                                                                                                _ = len(v1)
                                                                                                fmt.Println(v7, v7, v32)
                                                                                                fmt.Println(v7, fi16)
                                                                                                var v59 int32 = (v21 ^ v21)
                                                                                                if false {
                                                                                                    var v60 int32 = v21
                                                                                                    if false {
                                                                                                        _ = len(v5)
                                                                                                        v7 -= (v7 & (v7 >> ((42) & 63)))
                                                                                                    }
                                                                                                    if false {
                                                                                                        v61 := GenS1{F0: fi22}
                                                                                                        v62 := GenS2{F0: "\n\t", F1: fi22}
                                                                                                        _ = v61
                                                                                                        _ = v62
                                                                                                    }
                                                                                                    _ = v60
                                                                                                } else {
                                                                                                    v59 -= v59
                                                                                                }
                                                                                                if (((((((v2 + (-2.25 - (float64(v32) + v58))) - (float64(float32((0.1 / float64(0.1)))) - ((1.5 - v58) / v58))) * 1.5) - v58) + v58) + v2) < v2) {
                                                                                                    v33 = (v33 - v33)
                                                                                                    if true {
                                                                                                        var v63 float64 = v58
                                                                                                        v33 += v33
                                                                                                        fmt.Println(fi16, v21)
                                                                                                        _ = v63
                                                                                                    } else {
                                                                                                        v64 := v3[("hello" + ("hello" + (("hello" + "a") + ("" + ("\n\t" + ("foo bar" + ("foo bar" + "foo bar")))))))]
                                                                                                        var v65 uint = 824
                                                                                                        _ = v64
                                                                                                        _ = v65
                                                                                                    }
                                                                                                    if (!(!false)) {
                                                                                                        var v66 bool = (855 >= 592)
                                                                                                        _ = v66
                                                                                                    } else {
                                                                                                        v33 = 1
                                                                                                        var v67 int16 = -1
                                                                                                        _ = v67
                                                                                                    }
                                                                                                } else {
                                                                                                    switch "foo bar" {
                                                                                                    case "":
                                                                                                        fmt.Println(v7, v7)
                                                                                                    case "a":
                                                                                                        if ((fi22 & (fi22 << ((42) & 63))) == int(v7)) {
                                                                                                            _ = len(v1)
                                                                                                            _ = len(v4)
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        var v68 int16 = -335
                                                                                                        _ = v68
                                                                                                    }
                                                                                                    var v69 int16 = 354
                                                                                                    _ = v69
                                                                                                }
                                                                                                v58 = v58
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    var v70 int32 = (v59 >> ((65535) & 63))
                                                                                                    _ = v70
                                                                                                case "foo bar":
                                                                                                    v58 -= 1.5
                                                                                                case "hello":
                                                                                                    var v71 uint = 0
                                                                                                    v72 := GenS2{F0: "hello", F1: 9223372036854775807}
                                                                                                    _ = v71
                                                                                                    _ = v72
                                                                                                default:
                                                                                                    v32 += v32
                                                                                                }
                                                                                                var v73 string = "foo bar"
                                                                                                _ = v58
                                                                                                _ = v59
                                                                                                _ = v73
                                                                                            }
                                                                                        }
                                                                                        v74 := GenS4{F0: v7, F1: 42, F2: 251}
                                                                                        v75 := v3[("a" + "\n\t")]
                                                                                        _ = v74
                                                                                        _ = v75
                                                                                    default:
                                                                                        v21 += v21
                                                                                    }
                                                                                    v76 := GenS6{F0: fi16, F1: (v7 | v7), F2: 176}
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                        case -456:
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                v77 := GenS7{F0: false}
                                                                                _ = v77
                                                                            default:
                                                                                if ((-125 >= 127) && (0 < 1)) {
                                                                                    fmt.Println(v32, v7)
                                                                                    delete(v4, (v33 >> ((856) & 63)))
                                                                                    switch 305 {
                                                                                    case 565:
                                                                                        if (255 <= 128) {
                                                                                            v6 = append(v6, (fi16 << ((826) & 63)))
                                                                                        } else {
                                                                                            v7 -= 18446744073709551615
                                                                                            if (817 <= 65535) {
                                                                                                if false {
                                                                                                    if (!((-128 >= 127) || (0.1 <= (v2 * float64(1e10))))) {
                                                                                                        fmt.Println(fi22, fi16)
                                                                                                        var v78 uint16 = 65535
                                                                                                        var v79 uint16 = (^(v78))
                                                                                                        _ = v78
                                                                                                        _ = v79
                                                                                                    } else {
                                                                                                        switch (("foo bar" + "\n\t") + "\n\t") {
                                                                                                        case "foo bar":
                                                                                                            var v80 int8 = 1
                                                                                                            var v81 uint32 = v33
                                                                                                            v2 = v2
                                                                                                            if (1 == 65535) {
                                                                                                                if true {
                                                                                                                    v7 = v7
                                                                                                                    v33 += v81
                                                                                                                    _ = len(v1)
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v82 uint16 = 742
                                                                                                                if true {
                                                                                                                    v83 := GenS4{F0: 720, F1: (v80 - 42), F2: 127}
                                                                                                                    _ = v83
                                                                                                                }
                                                                                                                switch (-(fi16)) {
                                                                                                                case 42:
                                                                                                                    var v84 uint8 = 1
                                                                                                                    _ = v84
                                                                                                                case 868:
                                                                                                                    fmt.Println(fi16)
                                                                                                                }
                                                                                                                _ = v82
                                                                                                            }
                                                                                                            _ = v80
                                                                                                            _ = v81
                                                                                                        case "\n\t":
                                                                                                            var v85 uint64 = (^(v7))
                                                                                                            _ = v85
                                                                                                        default:
                                                                                                            if (v21 != (v21 / ((v21) | 1))) {
                                                                                                                v86 := GenS7{F0: true}
                                                                                                                v32 = v32
                                                                                                                _ = v86
                                                                                                            } else {
                                                                                                                v2 = 0.0
                                                                                                                v87 := v3[("a" + "hello")]
                                                                                                                _ = v87
                                                                                                            }
                                                                                                        }
                                                                                                        _ = len(v1)
                                                                                                    }
                                                                                                    if true {
                                                                                                        var v88 uint64 = 4294967295
                                                                                                        v33 += v33
                                                                                                        _ = v88
                                                                                                    }
                                                                                                } else {
                                                                                                    switch "a" {
                                                                                                    case "\n\t":
                                                                                                        if true {
                                                                                                            v33 = (v33 % (((v33 >> ((2) & 63))) | 1))
                                                                                                            switch "hello" {
                                                                                                            case "a":
                                                                                                                var v89 uint16 = 160
                                                                                                                if true {
                                                                                                                    fmt.Println(v21, v33, v2)
                                                                                                                    var v90 uint16 = v89
                                                                                                                    var v91 int32 = v21
                                                                                                                    var v92 float32 = 0.0
                                                                                                                    _ = v90
                                                                                                                    _ = v91
                                                                                                                    _ = v92
                                                                                                                } else {
                                                                                                                    v32 += (float32(0.0) + v32)
                                                                                                                    v93 := GenS6{F0: 658, F1: v7, F2: 127}
                                                                                                                    _ = v93
                                                                                                                }
                                                                                                                v7 = (v7 + v7)
                                                                                                                _ = v89
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = len(v6)
                                                                                                }
                                                                                            }
                                                                                            var v94 bool = (570 > 0)
                                                                                            _ = v94
                                                                                        }
                                                                                    case -851:
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                        fmt.Println(v7, v21)
                                                                                        v7 = v7
                                                                                    case -64:
                                                                                        if (("a" != "a") && false) {
                                                                                            v95 := GenS1{F0: (fi16 * fi22)}
                                                                                            _ = v95
                                                                                        }
                                                                                    }
                                                                                    v96 := GenS5{F0: 42, F1: v21}
                                                                                    v6 = append(v6, fi22)
                                                                                    _ = v96
                                                                                } else {
                                                                                    v97 := v4[(v33 - v33)]
                                                                                    _ = v97
                                                                                }
                                                                            }
                                                                        case 715:
                                                                            _ = len(v4)
                                                                        }
                                                                    }
                                                                    if true {
                                                                        switch 4294967295 {
                                                                        case 4294967295:
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        default:
                                                                            var v98 uint64 = (^(v7))
                                                                            _ = v98
                                                                        }
                                                                        switch (^(fi16)) {
                                                                        case -2147483648:
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        case 42:
                                                                            fmt.Println(v33, v33)
                                                                        default:
                                                                            if false {
                                                                                if ((!(128 == 2)) || (!false)) {
                                                                                    if (!true) {
                                                                                        var v99 string = "a"
                                                                                        v100 := GenS8{F0: (fi22 >> ((2) & 63)), F1: v99, F2: "hello"}
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                        switch "a" {
                                                                                        case "\n\t":
                                                                                            v101 := GenS2{F0: "foo bar", F1: fi16}
                                                                                            v32 += v32
                                                                                            v102 := GenS9{F0: false, F1: 4294967295}
                                                                                            if (false && (!(2147483648 != (v33 & uint32(v7))))) {
                                                                                                v4[2147483648] = "hello"
                                                                                                var v103 uint16 = 255
                                                                                                var v104 string = v99
                                                                                                _ = v103
                                                                                                _ = v104
                                                                                            } else {
                                                                                                var v105 int64 = -644
                                                                                                _ = v105
                                                                                            }
                                                                                            var v106 bool = ((-(fi16)) > (fi22 >> ((1) & 63)))
                                                                                            _ = v101
                                                                                            _ = v102
                                                                                            _ = v106
                                                                                        default:
                                                                                            switch (fi16 | 1) {
                                                                                            case 402:
                                                                                                switch (-(fi16)) {
                                                                                                case 0:
                                                                                                    if true {
                                                                                                        if (125 >= 719) {
                                                                                                            v32 += (v32 / v32)
                                                                                                        } else {
                                                                                                            v7 += (v7 << ((460) & 63))
                                                                                                        }
                                                                                                        v99 = v99
                                                                                                        v107 := GenS9{F0: false, F1: 65535}
                                                                                                        var v108 float32 = (-2.25 / v32)
                                                                                                        _ = v107
                                                                                                        _ = v108
                                                                                                    }
                                                                                                case 126:
                                                                                                    var v109 int32 = (v21 / ((157) | 1))
                                                                                                    _ = v109
                                                                                                case 606:
                                                                                                    fmt.Println(v33, v99, v2)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v100
                                                                                        _ = v99
                                                                                    } else {
                                                                                        v3[("foo bar" + "hello")] = ((false && false) || ((fi22 & (fi22 / (((fi16 & 0)) | 1))) == (fi22 | (fi22 * fi16))))
                                                                                        v110 := GenS10{F0: "hello", F1: -1, F2: ("a" + (("" + "foo bar") + "\n\t"))}
                                                                                        if (((v2 - ((((3.14159 / ((3.14159 + (((0.1 - v2) + (v2 * 0.0)) - ((v2 - 3.14159) / (v2 / (v2 / v2))))) * (1e10 / float64(0.1)))) + ((3.14159 / float64(0.1)) - (v2 * v2))) + ((v2 - 0.1) - (v2 + v2))) + v2)) >= 0.1) && ((v2 >= v2) || (30 == 42))) {
                                                                                            if len(v6) > 0 {
                                                                                                v6[0] = (^(fi22))
                                                                                            }
                                                                                            fmt.Println(v2, fi16)
                                                                                            if true {
                                                                                                var v111 uint = 4294967295
                                                                                                if (!((v7 << ((4294967295) & 63)) <= uint64(v33))) {
                                                                                                    v21 = (v21 * -2147483648)
                                                                                                    var v112 uint16 = 793
                                                                                                    v111 += (^(v111))
                                                                                                    _ = v112
                                                                                                } else {
                                                                                                    var v113 int = -1
                                                                                                    fmt.Println(fi16, v21)
                                                                                                    var v114 string = "\n\t"
                                                                                                    _ = v113
                                                                                                    _ = v114
                                                                                                }
                                                                                                var v115 int64 = int64(v111)
                                                                                                _ = v111
                                                                                                _ = v115
                                                                                            }
                                                                                        } else {
                                                                                            v2 += (0.0 + float64(0.1))
                                                                                            var v116 uint8 = 45
                                                                                            v21 = (v21 * v21)
                                                                                            _ = v116
                                                                                        }
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            if ((128 >= 96) && true) {
                                                                                                v32 -= float32(-2.25)
                                                                                                _ = len(v6)
                                                                                                if false {
                                                                                                    v117 := GenS5{F0: v33, F1: -2147483648}
                                                                                                    fmt.Println(v32, fi22)
                                                                                                    _ = v117
                                                                                                }
                                                                                            }
                                                                                        case "a":
                                                                                            var v118 int = (fi22 / ((-637) | 1))
                                                                                            _ = v118
                                                                                        case "\n\t":
                                                                                            v119 := v3["foo bar"]
                                                                                            _ = v119
                                                                                        default:
                                                                                            if (true && (0 <= (v33 * (^(v33))))) {
                                                                                                if false {
                                                                                                    var v120 int8 = -128
                                                                                                    _ = v120
                                                                                                } else {
                                                                                                    v21 += -1
                                                                                                    var v121 uint64 = 4294967295
                                                                                                    _ = v121
                                                                                                }
                                                                                                switch "foo bar" {
                                                                                                case "\n\t":
                                                                                                    if true {
                                                                                                        v5 = append(v5, v33)
                                                                                                        if true {
                                                                                                            v122 := GenS10{F0: "", F1: 42, F2: ("a" + "")}
                                                                                                            var v123 uint = uint(v33)
                                                                                                            _ = v122
                                                                                                            _ = v123
                                                                                                        }
                                                                                                        v2 -= 1.5
                                                                                                    }
                                                                                                case "hello":
                                                                                                    if true {
                                                                                                        delete(v4, 914)
                                                                                                        v21 += 0
                                                                                                    } else {
                                                                                                        var v124 int64 = 508
                                                                                                        var v125 bool = ((v7 | v7) < v7)
                                                                                                        _ = v124
                                                                                                        _ = v125
                                                                                                    }
                                                                                                }
                                                                                                var v126 float32 = 3.14159
                                                                                                if (-22 > -1) {
                                                                                                    var v127 float64 = v2
                                                                                                    _ = v127
                                                                                                }
                                                                                                _ = v126
                                                                                            }
                                                                                        }
                                                                                        v128 := GenS2{F0: "", F1: fi22}
                                                                                        _ = v110
                                                                                        _ = v128
                                                                                    }
                                                                                    v129 := v4[uint32(v7)]
                                                                                    fmt.Println(v2)
                                                                                    _ = v129
                                                                                }
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = (^(v33))
                                                                                }
                                                                            }
                                                                        }
                                                                        var v130 int32 = 42
                                                                        _ = v130
                                                                    }
                                                                    _ = v33
                                                                } else {
                                                                    delete(v4, 380)
                                                                    var v131 uint64 = (v7 / (((v7 << ((18446744073709551615) & 63))) | 1))
                                                                    _ = v131
                                                                }
                                                                _ = v32
                                                            } else {
                                                                if (true && (false && true)) {
                                                                    v2 = (v2 / 1.0)
                                                                }
                                                                if true {
                                                                    v21 = (v21 - v21)
                                                                } else {
                                                                    v132 := GenS4{F0: (v7 >> ((262) & 63)), F1: 100, F2: 42}
                                                                    var v133 uint8 = 2
                                                                    v133 -= v133
                                                                    var v134 uint64 = (^(v7))
                                                                    _ = v132
                                                                    _ = v133
                                                                    _ = v134
                                                                }
                                                                v21 = v21
                                                            }
                                                        }
                                                    }
                                                    v2 = v2
                                                } else {
                                                    if true {
                                                        if false {
                                                            if (("foo bar" + "\n\t") == ("\n\t" + "hello")) {
                                                                v6 = append(v6, 32767)
                                                                var v135 int8 = 100
                                                                fmt.Println(v135, fi22, v21)
                                                                _ = v135
                                                            }
                                                            v136 := GenS11{F0: true, F1: 18446744073709551615}
                                                            if true {
                                                                var v137 uint64 = 775
                                                                v138 := GenS6{F0: fi22, F1: v137, F2: 0}
                                                                _ = v137
                                                                _ = v138
                                                            }
                                                            _ = v136
                                                        }
                                                        if (1 >= 1) {
                                                            v2 -= (-2.25 * (3.14159 + 1.5))
                                                        } else {
                                                            v21 += -911
                                                            var v139 uint32 = 2147483648
                                                            var v140 uint32 = v139
                                                            _ = v139
                                                            _ = v140
                                                        }
                                                    }
                                                }
                                            } else {
                                                var v141 int64 = -9223372036854775808
                                                fmt.Println(fi22, fi16, v7)
                                                if false {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v142 string = ""
                                                    var v143 int64 = (v141 ^ 1)
                                                    _ = v142
                                                    _ = v143
                                                }
                                                _ = v141
                                            }
                                        }
                                    }
                                }
                                v7 = (v7 & v7)
                            } else {
                                switch "a" {
                                case "\n\t":
                                    v7 += 0
                                case "a":
                                    var v144 int16 = 32767
                                    _ = v144
                                default:
                                    var v145 float32 = float32(v2)
                                    _ = v145
                                }
                                if (0 > 0) {
                                    _ = len(v1)
                                    v7 += (v7 >> ((65535) & 63))
                                } else {
                                    var v146 int8 = 127
                                    if ((v146 - (v146 + -5)) < int8(v7)) {
                                        v147 := GenS8{F0: (fi22 % ((fi16) | 1)), F1: "a", F2: "a"}
                                        var v148 string = ""
                                        _ = v147
                                        _ = v148
                                    } else {
                                        _ = len(v3)
                                        v149 := GenS12{F0: fi22, F1: 0, F2: "hello"}
                                        _ = v149
                                    }
                                    v5 = append(v5, 859)
                                    _ = v146
                                }
                                v5 = append(v5, 4294967295)
                            }
                            _ = fi22
                        }
                        _ = v21
                    }
                    _ = fi16
                }
            }
            var v150 uint = 2
            _ = v150
        } else {
            for fi151 := 0; fi151 < 6; fi151++ {
                v152 := GenS13{F0: 0}
                switch -2147483648 {
                case 291:
                    fmt.Println(fi151, v7)
                case 9223372036854775807:
                    if false {
                        for fi153 := 0; fi153 < 10; fi153++ {
                            v2 -= 1.5
                            var v154 float32 = float32(v2)
                            v154 = float32(v2)
                            v155 := GenS14{F0: 0, F1: 407}
                            _ = fi153
                            _ = v154
                            _ = v155
                        }
                        if (!true) {
                            var v156 int16 = 32767
                            _ = v156
                        }
                        var v157 float64 = (((((float64((float32((v2 * (v2 / 1.0))) - 1.5)) / 1.0) * v2) - 1.5) * 0.1) - float64(1e10))
                        _ = v157
                    }
                default:
                    switch (^(fi151)) {
                    case 2147483647:
                        var v158 int8 = 127
                        _ = v158
                    case 9223372036854775807:
                        if len(v1) > 0 {
                            v1[0] = fi151
                        }
                    default:
                        v7 = 1
                    }
                }
                delete(v3, ((("a" + (("\n\t" + ("hello" + ("" + "a"))) + "a")) + "") + ""))
                _ = fi151
                _ = v152
            }
        }
    }
    var v159 uint16 = 711
    v4[65535] = ("\n\t" + "foo bar")
    var v160 uint8 = 211
    var v161 bool = false
    if true {
        if true {
            switch "a" {
            case "a":
                v162 := GenS6{F0: -9223372036854775808, F1: (v7 & (^(v7))), F2: v160}
                _ = v162
            case "":
                v7 += (^(v7))
            case "foo bar":
                v161 = (189 < 219)
            default:
                v163 := GenS15{F0: (0 != 0), F1: (v160 & 128), F2: 127}
                _ = v163
            }
            var v164 float32 = 1e10
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v164
        }
    } else {
        var v165 int32 = -2147483648
        v2 = (0.0 / 1.5)
        if (-466 < -824) {
            if (490 == (^(v159))) {
                var v166 int16 = -235
                _ = v166
            } else {
                var v167 int8 = 100
                _ = v167
            }
            v168 := GenS14{F0: v160, F1: v7}
            _ = v168
        } else {
            for fi169 := 0; fi169 < 17; fi169++ {
                switch "a" {
                case "a":
                    for fi170 := 0; fi170 < 3; fi170++ {
                        v171 := GenS16{F0: (v159 ^ 255), F1: 256}
                        v161 = v161
                        var v172 float64 = float64((1e10 + (0.0 - (1e10 - 0.0))))
                        _ = fi170
                        _ = v171
                        _ = v172
                    }
                default:
                    v161 = v161
                }
                var v173 uint8 = v160
                v174 := GenS13{F0: v7}
                _ = fi169
                _ = v173
                _ = v174
            }
        }
        _ = v165
    }
    var v175 uint8 = (v160 - v160)
    v160 -= 1
    var v176 float64 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v159:", v159)
    fmt.Println("v160:", v160)
    fmt.Println("v161:", v161)
    fmt.Println("v175:", v175)
    fmt.Println("v176:", v176)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v159
    _ = v160
    _ = v161
    _ = v175
    _ = v176
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
