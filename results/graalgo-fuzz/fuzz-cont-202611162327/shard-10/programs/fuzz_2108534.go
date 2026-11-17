package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS7 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS8 struct { F0 bool; F1 string; F2 uint16 }
type GenS9 struct { F0 string; F1 uint32; F2 int16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int8; F1 int32 }

func main() {
    var v1 float64 = 0.1
    v2 := []string{"hello", ("" + "\n\t")}
    v3 := []int{9223372036854775807, -1, 9223372036854775807, 32768}
    v4 := make(map[int]string)
    v5 := []string{"foo bar", ("hello" + ""), "", ("" + "foo bar")}
    v6 := make(map[string]int)
    v1 = v1
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(fi8)
            v10 := GenS1{F0: 775, F1: 255}
            _ = fi9
            _ = v10
        }
        v11 := GenS2{F0: ((^(v7)) >= v7), F1: 1}
        _ = fi8
        _ = v11
    }
    fmt.Println(v1)
    v1 = float64(1e10)
    switch -1 {
    case 4294967295:
        var v12 uint16 = 1
        _ = v12
    case -9223372036854775808:
        if (!false) {
            delete(v4, 127)
            v1 = 0.0
            if (((!true) || ("a" < "a")) && ((442 == (^(v7))) && false)) {
                for fi13 := 0; fi13 < 15; fi13++ {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        switch (("" + ((("a" + "foo bar") + "hello") + "")) + "foo bar") {
                        case "a":
                            v15 := GenS2{F0: true, F1: 42}
                            _ = v15
                        case "foo bar":
                            delete(v4, -340)
                        case "\n\t":
                            v1 -= v1
                        default:
                            v1 -= (v1 / v1)
                            var v16 float32 = float32((v1 - 0.1))
                            fmt.Println(v16, v7, fi14)
                            _ = v16
                        }
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = fi14
                    }
                    var v17 uint8 = 10
                    _ = fi13
                    _ = v17
                }
                fmt.Println(v1, v7)
                var v18 uint32 = 2147483647
                _ = v18
            } else {
                if (false || (v1 < v1)) {
                    v7 = v7
                    var v19 bool = (!false)
                    _ = v19
                } else {
                    if true {
                        v20 := GenS3{F0: (v7 >> ((791) & 63)), F1: 2, F2: 21}
                        _ = v20
                    }
                    switch "\n\t" {
                    case "hello":
                        v1 -= (float64(((1.5 - 0.1) * float32(v1))) + float64(((((-2.25 - 1.5) * (-2.25 / -2.25)) * 0.0) * 0.0)))
                    case "\n\t":
                        var v21 uint16 = 42
                        _ = v21
                    case "a":
                        for fi22 := 0; fi22 < 10; fi22++ {
                            for fi23 := 0; fi23 < 20; fi23++ {
                                v24 := GenS1{F0: 2, F1: 42}
                                if (false && ((v1 != (v1 + (v1 + v1))) || false)) {
                                    switch ((("hello" + "foo bar") + "") + "a") {
                                    case "\n\t":
                                        if false {
                                            var v25 bool = (!true)
                                            switch (fi23 * 127) {
                                            case -9223372036854775808:
                                                fmt.Println(v1, v1)
                                                var v26 int32 = 2147483647
                                                var v27 uint32 = uint32(v7)
                                                _ = v26
                                                _ = v27
                                            default:
                                                if v25 {
                                                    v28 := GenS3{F0: v7, F1: 65535, F2: -1}
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch -548 {
                                                    case 2147483647:
                                                        fmt.Println(fi23, v7)
                                                    case -834:
                                                        v25 = (-1 < -892)
                                                    case -9223372036854775808:
                                                        if (0 >= 127) {
                                                            var v29 int64 = 615
                                                            if v25 {
                                                                switch (^(fi22)) {
                                                                case -500:
                                                                    if v25 {
                                                                        var v30 bool = (65535 < 875)
                                                                        _ = v30
                                                                    } else {
                                                                        var v31 int32 = 0
                                                                        v32 := GenS2{F0: (false && (1 >= 4294967295)), F1: 42}
                                                                        _ = v31
                                                                        _ = v32
                                                                    }
                                                                case 9223372036854775807:
                                                                    v33 := GenS4{F0: v25}
                                                                    _ = v33
                                                                case 2147483647:
                                                                    var v34 uint8 = 2
                                                                    _ = v34
                                                                default:
                                                                    if (float64(0.0) != -2.25) {
                                                                        fmt.Println(v25, v7, fi22)
                                                                        fmt.Println(fi23, v25, fi23)
                                                                        switch ("foo bar" + "\n\t") {
                                                                        case "":
                                                                            v1 = float64(0.1)
                                                                        case "foo bar":
                                                                            v3 = append(v3, 0)
                                                                        }
                                                                    } else {
                                                                        if ((("\n\t" + "a") + "") >= (("\n\t" + ("a" + "a")) + "hello")) {
                                                                            v35 := GenS5{F0: v29}
                                                                            delete(v6, "a")
                                                                            v25 = (!v25)
                                                                            _ = v35
                                                                        }
                                                                        v36 := GenS2{F0: true, F1: 18446744073709551615}
                                                                        _ = v36
                                                                    }
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = (fi23 / (((fi22 >> ((1) & 63))) | 1))
                                                                }
                                                                if true {
                                                                    fmt.Println(v25, v29, fi22)
                                                                    var v37 uint64 = (^(v7))
                                                                    v38 := GenS6{F0: 42, F1: -2147483648, F2: 201}
                                                                    v39 := GenS3{F0: (^(v7)), F1: 2147483647, F2: (-(v29))}
                                                                    _ = v37
                                                                    _ = v38
                                                                    _ = v39
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = ("\n\t" + "hello")
                                                                    }
                                                                    var v40 int32 = 0
                                                                    switch ("foo bar" + "hello") {
                                                                    case "foo bar":
                                                                        if (1000 > -32768) {
                                                                            fmt.Println(v29)
                                                                        } else {
                                                                            var v41 uint64 = (v7 % (((^(v7))) | 1))
                                                                            v1 -= v1
                                                                            _ = v41
                                                                        }
                                                                    case "\n\t":
                                                                        var v42 uint32 = 42
                                                                        if v25 {
                                                                            v43 := GenS4{F0: (!v25)}
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            if (3.14159 != (1.5 / 1.5)) {
                                                                                v44 := GenS7{F0: 235, F1: ("" >= (((("\n\t" + ((("" + "") + ("foo bar" + "a")) + "a")) + "foo bar") + "foo bar") + ("a" + "hello"))), F2: 0}
                                                                                var v45 uint32 = (v42 - 532)
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    _ = len(v5)
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        v7 = v7
                                                                                        var v46 float32 = 1e10
                                                                                        _ = v46
                                                                                    }
                                                                                case "":
                                                                                    v45 -= v42
                                                                                default:
                                                                                    v5 = append(v5, "hello")
                                                                                }
                                                                                _ = v44
                                                                                _ = v45
                                                                            }
                                                                            v40 = -2147483648
                                                                            _ = v43
                                                                        }
                                                                        _ = v42
                                                                    default:
                                                                        var v47 string = ""
                                                                        _ = v47
                                                                    }
                                                                    _ = v40
                                                                }
                                                            } else {
                                                                fmt.Println(v25)
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            if true {
                                                                v48 := GenS5{F0: (v29 % (((^(v29))) | 1))}
                                                                v1 += v1
                                                                _ = v48
                                                            }
                                                            fmt.Println(v1, v7, v29)
                                                            fmt.Println(fi23)
                                                            v1 += v1
                                                            _ = v29
                                                        }
                                                    default:
                                                        if (float64(-2.25) > (v1 * 1e10)) {
                                                            if v25 {
                                                                if true {
                                                                    fmt.Println(fi23, fi22)
                                                                    var v49 int8 = -63
                                                                    _ = v49
                                                                } else {
                                                                    var v50 uint8 = 255
                                                                    switch ("" + "hello") {
                                                                    case "\n\t":
                                                                        v1 -= ((0.0 + 0.1) - v1)
                                                                    case "a":
                                                                        switch "a" {
                                                                        case "":
                                                                            v1 += -2.25
                                                                        }
                                                                    case "foo bar":
                                                                        v50 = (^(v50))
                                                                    default:
                                                                        v1 = v1
                                                                        v51 := GenS1{F0: 255, F1: 104}
                                                                        v25 = (0.1 == float32(((float64(-2.25) * 0.1) - v1)))
                                                                        fmt.Println(v50, v1)
                                                                        _ = v51
                                                                    }
                                                                    v25 = ((0.1 != -2.25) && (v25 && (2147483647 < 606)))
                                                                    _ = v50
                                                                }
                                                                var v52 uint32 = 65535
                                                                var v53 uint8 = uint8(v7)
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                        } else {
                                                            v54 := GenS2{F0: false, F1: 1}
                                                            var v55 float32 = float32(float64(-2.25))
                                                            v56 := GenS5{F0: -1}
                                                            _ = v54
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    }
                                                    _ = v28
                                                } else {
                                                    _ = len(v2)
                                                    if v25 {
                                                        var v57 float64 = (1e10 - -2.25)
                                                        if ((42 > -42) || ((v7 < v7) || (1 != int64(v7)))) {
                                                            var v58 float32 = float32((v1 - v57))
                                                            v1 = v1
                                                            if v25 {
                                                                delete(v4, (fi23 ^ (fi22 / (((fi23 + (fi23 | fi22))) | 1))))
                                                                var v59 int8 = -128
                                                                switch (fi23 ^ 476) {
                                                                case 2147483647:
                                                                    if (255 >= uint16(v7)) {
                                                                        v60 := GenS8{F0: v25, F1: (("a" + ("a" + "hello")) + "a"), F2: 0}
                                                                        _ = v60
                                                                    } else {
                                                                        var v61 int = fi22
                                                                        _ = v61
                                                                    }
                                                                    v25 = (v58 != v58)
                                                                    switch (fi23 >> ((18446744073709551615) & 63)) {
                                                                    case 32768:
                                                                        var v62 uint32 = 1
                                                                        _ = v62
                                                                    case 127:
                                                                        var v63 float32 = v58
                                                                        _ = v63
                                                                    }
                                                                case 9223372036854775807:
                                                                    switch (fi22 + fi23) {
                                                                    case 775:
                                                                        delete(v4, (fi22 & fi23))
                                                                    case -903:
                                                                        v4[fi22] = "\n\t"
                                                                    case 42:
                                                                        fmt.Println(fi23, v7, fi22)
                                                                    default:
                                                                        if true {
                                                                            switch (^(fi22)) {
                                                                            case 32768:
                                                                                v7 = (v7 + 2)
                                                                            case -1:
                                                                                v5 = append(v5, "")
                                                                            case 9223372036854775807:
                                                                                if v25 {
                                                                                    v64 := GenS4{F0: ((fi23 * fi22) > fi22)}
                                                                                    _ = v64
                                                                                } else {
                                                                                    fmt.Println(v7, fi23, v57)
                                                                                    v7 -= (^(v7))
                                                                                    if (v59 <= 1) {
                                                                                        v65 := GenS6{F0: 128, F1: 0, F2: 0}
                                                                                        var v66 float64 = (float64(float32(((v1 / 0.1) / 1.0))) * ((float64((v58 * (((v58 / v58) + -2.25) - v58))) - v1) - 1e10))
                                                                                        var v67 uint8 = 128
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    } else {
                                                                                        fmt.Println(fi22)
                                                                                        if v25 {
                                                                                            v68 := v4[32767]
                                                                                            _ = v68
                                                                                        } else {
                                                                                            fmt.Println(v58, v58, v58)
                                                                                            v69 := GenS9{F0: ("foo bar" + "hello"), F1: 208, F2: 42}
                                                                                            switch (fi23 << ((288) & 63)) {
                                                                                            case 32767:
                                                                                                var v70 uint64 = v7
                                                                                                _ = v70
                                                                                            case 42:
                                                                                                _ = len(v3)
                                                                                            case 440:
                                                                                                v71 := GenS10{F0: 0}
                                                                                                _ = v71
                                                                                            default:
                                                                                                var v72 uint16 = 1
                                                                                                _ = v72
                                                                                            }
                                                                                            _ = v69
                                                                                        }
                                                                                        var v73 float64 = v57
                                                                                        v25 = v25
                                                                                        _ = v73
                                                                                    }
                                                                                    if ((1e10 + v58) != 0.1) {
                                                                                        v74 := GenS7{F0: 42, F1: false, F2: 0}
                                                                                        _ = v74
                                                                                    } else {
                                                                                        var v75 uint32 = 947
                                                                                        _ = v75
                                                                                    }
                                                                                    switch fi22 {
                                                                                    case 127:
                                                                                        if true {
                                                                                            v5 = append(v5, "foo bar")
                                                                                            if (373 == 2) {
                                                                                                if false {
                                                                                                    switch "hello" {
                                                                                                    case "":
                                                                                                        v4[-227] = ""
                                                                                                    case "a":
                                                                                                        if false {
                                                                                                            var v76 int = (fi23 >> ((0) & 63))
                                                                                                            _ = v76
                                                                                                        } else {
                                                                                                            fmt.Println(fi23, v57, v59)
                                                                                                            switch "hello" {
                                                                                                            case "a":
                                                                                                                v57 = -2.25
                                                                                                            default:
                                                                                                                var v77 int = (fi22 ^ (fi22 * 32767))
                                                                                                                _ = v77
                                                                                                            }
                                                                                                            var v78 uint = 18446744073709551615
                                                                                                            _ = v78
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        if (v59 < (^(v59))) {
                                                                                                            if v25 {
                                                                                                                fmt.Println(v58, fi23, v58)
                                                                                                            } else {
                                                                                                                v1 = v1
                                                                                                                v57 = ((1e10 + -2.25) * (v1 * (v1 / 1e10)))
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch -2147483648 {
                                                                                                            case 0:
                                                                                                                if ((v59 - 127) < v59) {
                                                                                                                    fmt.Println(v1)
                                                                                                                    var v79 float32 = v58
                                                                                                                    _ = v79
                                                                                                                }
                                                                                                            case -584:
                                                                                                                v3 = append(v3, fi22)
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        var v80 uint8 = 42
                                                                                                        _ = v80
                                                                                                    }
                                                                                                } else {
                                                                                                    v7 = 4294967295
                                                                                                    var v81 uint16 = 255
                                                                                                    _ = v81
                                                                                                }
                                                                                                v82 := GenS4{F0: false}
                                                                                                v83 := GenS11{F0: 1}
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                            } else {
                                                                                                if v25 {
                                                                                                    v58 = v58
                                                                                                    v84 := GenS12{F0: (fi22 >> ((1) & 63))}
                                                                                                    _ = v84
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v85 := GenS13{F0: 18446744073709551615, F1: 179, F2: (v7 ^ v7)}
                                                                                            _ = v85
                                                                                        }
                                                                                    case 32768:
                                                                                        var v86 int32 = -716
                                                                                        _ = v86
                                                                                    case 128:
                                                                                        if (!false) {
                                                                                            v1 -= 1.5
                                                                                            v3 = append(v3, fi22)
                                                                                            v25 = (128 != 2)
                                                                                            if ((623 == 4294967295) || (128 >= 1)) {
                                                                                                v25 = (!(!true))
                                                                                                if v25 {
                                                                                                    fmt.Println(v25)
                                                                                                    v59 -= 1
                                                                                                    v7 -= (^(v7))
                                                                                                } else {
                                                                                                    var v87 float64 = v1
                                                                                                    _ = v87
                                                                                                }
                                                                                            } else {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                if v25 {
                                                                                                    var v88 uint = 42
                                                                                                    var v89 bool = (v25 || (((1 < 255) && false) && (225 > 877)))
                                                                                                    _ = v88
                                                                                                    _ = v89
                                                                                                } else {
                                                                                                    fmt.Println(fi22, v57, fi23)
                                                                                                    if true {
                                                                                                        v90 := GenS10{F0: 1000}
                                                                                                        v91 := GenS13{F0: 4294967295, F1: 368, F2: (v7 * (v7 | v7))}
                                                                                                        _ = v90
                                                                                                        _ = v91
                                                                                                    }
                                                                                                    v59 += (v59 >> ((1) & 63))
                                                                                                }
                                                                                                if false {
                                                                                                    if (126 == -905) {
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            v92 := GenS2{F0: ((v25 && v25) && v25), F1: 943}
                                                                                                            _ = v92
                                                                                                        }
                                                                                                    }
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    var v93 bool = (32767 <= 1000)
                                                                                                    _ = v93
                                                                                                } else {
                                                                                                    var v94 uint32 = 58
                                                                                                    _ = v94
                                                                                                }
                                                                                            }
                                                                                            var v95 int32 = 312
                                                                                            _ = v95
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if v25 {
                                                                                    v6["a"] = fi23
                                                                                    switch fi23 {
                                                                                    case 202:
                                                                                        if (831 > 0) {
                                                                                            var v96 uint8 = 2
                                                                                            var v97 float64 = (float64((((float32((3.14159 - (v57 + (float64(v58) / v1)))) + 1e10) + ((((float32(float64(0.0)) + (v58 / (v58 * ((((v58 * v58) * (v58 * float32(-2.25))) * v58) / v58)))) + v58) * v58) - (float32(v57) - v58))) / float32((v1 - 1.5)))) - ((v1 * ((v1 / (float64(v58) - v57)) - float64(v58))) * (v57 - 0.1)))
                                                                                            v98 := GenS5{F0: -939}
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                        }
                                                                                    case 42:
                                                                                        if true {
                                                                                            var v99 float64 = v57
                                                                                            _ = v99
                                                                                        } else {
                                                                                            v1 = 1.5
                                                                                        }
                                                                                    case 2147483647:
                                                                                        var v100 string = "hello"
                                                                                        _ = v100
                                                                                    }
                                                                                } else {
                                                                                    if false {
                                                                                        v101 := GenS14{F0: 18446744073709551615}
                                                                                        _ = v101
                                                                                    }
                                                                                    v58 = (v58 - v58)
                                                                                    var v102 bool = (!(171 >= 246))
                                                                                    var v103 string = "hello"
                                                                                    v104 := GenS4{F0: v102}
                                                                                    fmt.Println(fi23)
                                                                                    _ = v102
                                                                                    _ = v103
                                                                                    _ = v104
                                                                                }
                                                                            }
                                                                            v59 += (v59 ^ 100)
                                                                            v1 = (0.0 * ((0.1 - ((-2.25 + float64((v58 * 0.0))) / 1e10)) + v1))
                                                                        } else {
                                                                            if (false && (1e10 == ((float32(v1) - 0.0) * v58))) {
                                                                                v2 = append(v2, ("hello" + ""))
                                                                                if v25 {
                                                                                    v57 += float64(1.5)
                                                                                    switch fi22 {
                                                                                    case 9223372036854775807:
                                                                                        switch ("\n\t" + "") {
                                                                                        case "":
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v1, fi22)
                                                                                            case "a":
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    var v105 int64 = -1
                                                                                                    _ = v105
                                                                                                default:
                                                                                                    fmt.Println(v1)
                                                                                                }
                                                                                            case "hello":
                                                                                                var v106 float64 = (v1 - v57)
                                                                                                _ = v106
                                                                                            default:
                                                                                                var v107 string = "foo bar"
                                                                                                _ = v107
                                                                                            }
                                                                                        case "\n\t":
                                                                                            var v108 uint32 = 2
                                                                                            _ = v108
                                                                                        }
                                                                                    }
                                                                                    var v109 bool = false
                                                                                    _ = v109
                                                                                }
                                                                                switch (fi22 % ((1) | 1)) {
                                                                                case 1:
                                                                                    v59 += -128
                                                                                }
                                                                            }
                                                                            delete(v6, "hello")
                                                                            if ((-128 < v59) && (!false)) {
                                                                                v1 += 3.14159
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    var v110 bool = v25
                                                                    _ = v110
                                                                }
                                                                _ = v59
                                                            }
                                                            _ = v58
                                                        }
                                                        v111 := GenS1{F0: 756, F1: 1}
                                                        _ = v111
                                                        _ = v57
                                                    } else {
                                                        fmt.Println(v25, v7)
                                                    }
                                                    v7 = 4294967295
                                                }
                                            }
                                            _ = v25
                                        }
                                        var v112 uint32 = 4294967295
                                        fmt.Println(fi22, v112)
                                        _ = v112
                                    case "":
                                        v113 := GenS13{F0: (v7 % ((v7) | 1)), F1: 0, F2: v7}
                                        _ = v113
                                    case "a":
                                        var v114 string = ""
                                        fmt.Println(fi22, fi22, fi23)
                                        _ = v114
                                    }
                                    switch fi23 {
                                    case 32767:
                                        fmt.Println(fi22, v7)
                                    case -839:
                                        var v115 int64 = int64(v7)
                                        _ = v115
                                    case 4294967295:
                                        fmt.Println(v1)
                                        v116 := GenS14{F0: 819}
                                        v117 := GenS14{F0: 2}
                                        _ = v116
                                        _ = v117
                                    default:
                                        v1 += (v1 * float64(0.0))
                                    }
                                }
                                var v118 int = (fi22 & (fi23 ^ (fi23 << ((0) & 63))))
                                _ = fi23
                                _ = v118
                                _ = v24
                            }
                            if (42 < 65535) {
                                v7 -= v7
                            } else {
                                switch "foo bar" {
                                case "a":
                                    if (774 > 576) {
                                        var v119 uint16 = 42
                                        _ = v119
                                    } else {
                                        for fi120 := 0; fi120 < 14; fi120++ {
                                            fmt.Println(fi22, v7, fi22)
                                            if (42 == 5) {
                                                v7 = (v7 / (((v7 / ((v7) | 1))) | 1))
                                                v1 = (-2.25 * v1)
                                            } else {
                                                v7 = 1
                                                v1 -= v1
                                            }
                                            var v121 float64 = v1
                                            _ = fi120
                                            _ = v121
                                        }
                                        fmt.Println(v1, v7)
                                        for fi122 := 0; fi122 < 13; fi122++ {
                                            v123 := GenS5{F0: 1}
                                            fmt.Println(fi22, fi22, fi22)
                                            _ = fi122
                                            _ = v123
                                        }
                                    }
                                default:
                                    var v124 bool = (-2.25 == (0.0 - 0.0))
                                    _ = v124
                                }
                                _ = len(v4)
                            }
                            _ = fi22
                        }
                    }
                }
            }
        }
    case 1:
        var v125 uint32 = 0
        _ = v125
    }
    v1 += v1
    fmt.Println(v7)
    if false {
        fmt.Println(v7, v7)
    } else {
        var v126 int32 = 0
        _ = v126
    }
    if false {
        var v127 int8 = -1
        for fi128 := 0; fi128 < 13; fi128++ {
            var v129 int16 = 1
            v127 += -57
            if (!(!(!true))) {
                var v130 string = "foo bar"
                _ = v130
            } else {
                v127 = (-(v127))
            }
            _ = fi128
            _ = v129
        }
        v1 -= float64((float32(0.0) + 0.1))
        _ = v127
    } else {
        var v131 uint32 = 4294967295
        _ = v131
    }
    v132 := GenS15{F0: 1, F1: 620}
    if false {
        var v133 int8 = -56
        fmt.Println(v1, v7, v7)
        v133 = v133
        _ = v133
    } else {
        v1 = (((((float64((0.1 * 0.1)) * float64(3.14159)) - v1) * v1) - float64(float32(3.14159))) * 1e10)
        var v134 int64 = 9223372036854775807
        v135 := v4[-9223372036854775808]
        _ = v134
        _ = v135
    }
    v7 += v7
    fmt.Println(v7, v1)
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v132
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
