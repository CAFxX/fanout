package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int8; F2 int }
type GenS4 struct { F0 string; F1 int16; F2 uint }
type GenS5 struct { F0 uint16; F1 bool; F2 int8 }
type GenS6 struct { F0 bool; F1 bool; F2 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 string; F2 uint16 }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int64 }

func main() {
    var v1 uint8 = 254
    v2 := []string{"hello"}
    var v3 int32 = 42
    var v4 float32 = -2.25
    var v5 int16 = 32767
    var v6 bool = true
    if v6 {
        v7 := GenS1{F0: 255}
        for fi8 := 0; fi8 < 14; fi8++ {
            v5 = v5
            if v6 {
                if v6 {
                    if v6 {
                        var v9 float32 = (v4 - 0.0)
                        _ = v9
                    }
                    if (233 > 856) {
                        v10 := GenS2{F0: 65535}
                        _ = v10
                    }
                    v2 = append(v2, ("foo bar" + "\n\t"))
                } else {
                    fmt.Println(v3)
                }
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
            }
            _ = fi8
        }
        _ = v7
    } else {
        v3 -= v3
        var v11 uint32 = uint32(v1)
        v12 := GenS1{F0: 65535}
        _ = v11
        _ = v12
    }
    if v6 {
        v4 -= -2.25
        for fi13 := 0; fi13 < 10; fi13++ {
            switch fi13 {
            case 1:
                for fi14 := 0; fi14 < 14; fi14++ {
                    v5 += (v5 ^ (v5 & (v5 * v5)))
                    _ = fi14
                }
            case -809:
                fmt.Println(v1, v3)
            case 32768:
                if v6 {
                    var v15 int = fi13
                    _ = v15
                }
            default:
                v2 = append(v2, "hello")
            }
            _ = fi13
        }
        if v6 {
            switch 42 {
            case -619:
                for fi16 := 0; fi16 < 18; fi16++ {
                    switch "foo bar" {
                    case "":
                        for fi17 := 0; fi17 < 4; fi17++ {
                            v5 = v5
                            var v18 uint8 = v1
                            _ = fi17
                            _ = v18
                        }
                    }
                    var v19 uint8 = v1
                    _ = fi16
                    _ = v19
                }
            case -682:
                var v20 bool = v6
                _ = v20
            case 9223372036854775807:
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            }
            v2 = append(v2, "")
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS2{F0: 65535}
                v23 := GenS1{F0: 32767}
                var v24 int64 = 98
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 12; fi25++ {
                v26 := GenS1{F0: 2}
                _ = fi25
                _ = v26
            }
            if (v6 && false) {
                v27 := GenS3{F0: v1, F1: int8(v3), F2: 0}
                v6 = v6
                _ = v27
            } else {
                switch ("" + (("a" + "\n\t") + ((("" + "\n\t") + "") + "foo bar"))) {
                case "a":
                    v2 = append(v2, ((("hello" + ("" + "hello")) + "hello") + (("\n\t" + ("a" + "hello")) + "")))
                case "hello":
                    v5 = -1
                case "":
                    _ = len(v2)
                default:
                    var v28 string = ""
                    _ = v28
                }
            }
        }
    }
    if v6 {
        fmt.Println(v1, v3)
        var v29 string = "foo bar"
        _ = v29
    }
    fmt.Println(v4, v5)
    switch -915 {
    case 32767:
        if ((^(v1)) == 0) {
            for fi30 := 0; fi30 < 14; fi30++ {
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                _ = fi30
            }
            v1 -= 188
        }
    default:
        var v31 uint32 = 1
        _ = v31
    }
    var v32 uint16 = 2
    fmt.Println(v4)
    switch "hello" {
    case "a":
        var v33 int = 429
        _ = v33
    case "":
        if (!v6) {
            for fi34 := 0; fi34 < 17; fi34++ {
                switch (("foo bar" + ("\n\t" + "\n\t")) + ("" + (("foo bar" + "\n\t") + "\n\t"))) {
                case "":
                    fmt.Println(fi34, v1, v6)
                case "a":
                    v6 = v6
                }
                _ = fi34
            }
            if true {
                if ((v6 && (307 < v32)) && false) {
                    for fi35 := 0; fi35 < 10; fi35++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi35
                    }
                    switch -809 {
                    case -1:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case 0:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    fmt.Println(v3, v6, v32)
                } else {
                    if v6 {
                        switch "foo bar" {
                        case "a":
                            v4 = float32(3.14159)
                        case "\n\t":
                            v36 := GenS4{F0: (("foo bar" + "") + ("a" + ("hello" + "\n\t"))), F1: 42, F2: 18446744073709551615}
                            _ = v36
                        default:
                            var v37 uint64 = 18446744073709551615
                            _ = v37
                        }
                        switch (("foo bar" + "hello") + "foo bar") {
                        case "foo bar":
                            fmt.Println(v3, v1, v3)
                        default:
                            v6 = (v6 || v6)
                        }
                    } else {
                        v3 += int32(v32)
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v38 uint = 255
                    _ = v38
                }
                if v6 {
                    if ((v1 > (v1 | v1)) && v6) {
                        v39 := GenS3{F0: (v1 - (v1 >> ((18446744073709551615) & 63))), F1: -105, F2: int(v5)}
                        _ = v39
                    }
                    var v40 int = -1
                    _ = v40
                } else {
                    v4 += (((0.0 * -2.25) / float32(3.14159)) - v4)
                }
                var v41 int64 = -673
                _ = v41
            }
            v2 = append(v2, ("foo bar" + ("a" + "foo bar")))
        } else {
            v6 = (true || (((("foo bar" + ("a" + (("\n\t" + ("hello" + ("a" + "foo bar"))) + "foo bar"))) + "a") + "\n\t") >= "foo bar"))
            switch "a" {
            case "foo bar":
                if (v6 || ((-1 >= -1) && v6)) {
                    for fi42 := 0; fi42 < 16; fi42++ {
                        switch fi42 {
                        case 32768:
                            v3 += (v3 / ((345) | 1))
                        }
                        if v6 {
                            if v6 {
                                switch (fi42 / ((fi42) | 1)) {
                                case 32767:
                                    for fi43 := 0; fi43 < 16; fi43++ {
                                        v6 = (v6 || v6)
                                        if v6 {
                                            v2 = append(v2, ("\n\t" + (("hello" + "hello") + ((("hello" + ((("\n\t" + ((("hello" + "hello") + "") + "hello")) + "foo bar") + "\n\t")) + "hello") + (("a" + ("hello" + "")) + ("hello" + (("\n\t" + "hello") + "a")))))))
                                            v3 += v3
                                        }
                                        v32 = v32
                                        v3 -= (v3 * (^(v3)))
                                        var v44 int8 = -37
                                        _ = fi43
                                        _ = v44
                                    }
                                default:
                                    for fi45 := 0; fi45 < 13; fi45++ {
                                        if v6 {
                                            v5 = (v5 | 1000)
                                            if ((v3 * v3) >= v3) {
                                                var v46 uint = 42
                                                fmt.Println(fi42)
                                                _ = v46
                                            } else {
                                                v47 := GenS5{F0: 1, F1: v6, F2: int8(fi45)}
                                                v2 = append(v2, "\n\t")
                                                _ = v47
                                            }
                                            v4 -= float32(-2.25)
                                        }
                                        v3 -= v3
                                        _ = fi45
                                    }
                                }
                                switch 127 {
                                case -1:
                                    var v48 bool = (22 <= 9223372036854775807)
                                    _ = v48
                                default:
                                    v6 = v6
                                }
                            } else {
                                if true {
                                    v49 := GenS5{F0: uint16(v1), F1: v6, F2: 0}
                                    _ = v49
                                } else {
                                    v2 = append(v2, "a")
                                    var v50 uint = 65535
                                    _ = v50
                                }
                            }
                        }
                        if (((fi42 / ((fi42) | 1)) == 9223372036854775807) || true) {
                            fmt.Println(v1)
                        }
                        _ = fi42
                    }
                    v51 := GenS6{F0: (-583 >= -833), F1: (2 != (v1 | (v1 % ((v1) | 1)))), F2: 1}
                    switch -1 {
                    case -299:
                        v3 = (v3 & (v3 >> ((65535) & 63)))
                    case 32767:
                        fmt.Println(v4, v3)
                    case 4294967295:
                        v5 += v5
                    }
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 20; fi52++ {
                        switch fi52 {
                        case -497:
                            v6 = v6
                        case 9223372036854775807:
                            for fi53 := 0; fi53 < 9; fi53++ {
                                var v54 int16 = (v5 >> ((747) & 63))
                                _ = fi53
                                _ = v54
                            }
                        case -598:
                            if (!((-36 >= 65) && false)) {
                                switch ((("\n\t" + "") + "\n\t") + "foo bar") {
                                case "":
                                    var v55 float64 = ((float64(0.0) - (((-2.25 - float64(v4)) + (float64(v4) + 1e10)) / 3.14159)) - 0.1)
                                    _ = v55
                                case "\n\t":
                                    var v56 int = 127
                                    _ = v56
                                case "a":
                                    var v57 int = int(v3)
                                    _ = v57
                                default:
                                    if (2 < 2147483648) {
                                        var v58 int8 = -10
                                        v32 = (v32 & (v32 | (v32 / (((v32 >> ((2) & 63))) | 1))))
                                        _ = v58
                                    } else {
                                        var v59 uint32 = 2
                                        switch "hello" {
                                        case "a":
                                            if len(v2) > 0 {
                                                v2[0] = ("\n\t" + "hello")
                                            }
                                        case "hello":
                                            var v60 int = (-(fi52))
                                            _ = v60
                                        case "foo bar":
                                            if (!((fi52 & (fi52 ^ 581)) <= fi52)) {
                                                v4 = ((v4 + float32(3.14159)) + (v4 + 3.14159))
                                                var v61 bool = v6
                                                _ = v61
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            _ = len(v2)
                                        case "foo bar":
                                            v59 -= v59
                                        }
                                        _ = v59
                                    }
                                }
                                switch "a" {
                                case "\n\t":
                                    v3 -= v3
                                case "foo bar":
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if v6 {
                                            if v6 {
                                                if v6 {
                                                    var v62 uint = 255
                                                    for fi63 := 0; fi63 < 17; fi63++ {
                                                        var v64 bool = (!(v32 != v32))
                                                        var v65 int32 = 2147483647
                                                        _ = fi63
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = "a"
                                                    }
                                                    _ = v62
                                                } else {
                                                    for fi66 := 0; fi66 < 10; fi66++ {
                                                        v6 = true
                                                        if v6 {
                                                            var v67 bool = (int16(v32) < -754)
                                                            v2 = append(v2, ("" + (("foo bar" + (("" + "hello") + "")) + "a")))
                                                            _ = v67
                                                        } else {
                                                            v2 = append(v2, "a")
                                                            var v68 string = ""
                                                            _ = v68
                                                        }
                                                        _ = fi66
                                                    }
                                                    if false {
                                                        var v69 int32 = (-(v3))
                                                        _ = v69
                                                    }
                                                }
                                                for fi70 := 0; fi70 < 17; fi70++ {
                                                    if (!v6) {
                                                        var v71 float64 = 0.1
                                                        v1 -= 255
                                                        if true {
                                                            if (!(!(!true))) {
                                                                fmt.Println(fi70)
                                                                var v72 float64 = (0.0 - ((3.14159 * v71) * (((v71 * v71) / 1.5) + v71)))
                                                                _ = v72
                                                            } else {
                                                                fmt.Println(v5, v71)
                                                                v73 := GenS7{F0: -9223372036854775808}
                                                                _ = v73
                                                            }
                                                        }
                                                        _ = v71
                                                    } else {
                                                        var v74 string = ""
                                                        _ = v74
                                                    }
                                                    if v6 {
                                                        if ((v6 && true) && false) {
                                                            var v75 int16 = (v5 | int16(v32))
                                                            _ = v75
                                                        } else {
                                                            if false {
                                                                v3 -= (v3 / ((v3) | 1))
                                                                fmt.Println(v4, v4, v5)
                                                            } else {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            switch fi52 {
                                                            case -571:
                                                                v4 = v4
                                                            default:
                                                                if (1.5 >= 3.14159) {
                                                                    if v6 {
                                                                        if (false && (18446744073709551615 > 282)) {
                                                                            var v76 int64 = -756
                                                                            _ = v76
                                                                        }
                                                                        var v77 float32 = v4
                                                                        var v78 float64 = 3.14159
                                                                        _ = v77
                                                                        _ = v78
                                                                    }
                                                                    switch "a" {
                                                                    case "":
                                                                        var v79 float64 = float64(v4)
                                                                        _ = v79
                                                                    default:
                                                                        v6 = true
                                                                    }
                                                                    if ((!(-128 > -91)) || true) {
                                                                        if v6 {
                                                                            var v80 float32 = float32(1e10)
                                                                            fmt.Println(fi70)
                                                                            _ = v80
                                                                        }
                                                                        v81 := GenS8{F0: -54, F1: "\n\t", F2: (v32 - 693)}
                                                                        var v82 uint64 = 2
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                } else {
                                                                    var v83 int = (fi52 >> ((2) & 63))
                                                                    v5 -= v5
                                                                    _ = v83
                                                                }
                                                            }
                                                        }
                                                        if (!true) {
                                                            v2 = append(v2, (((((("foo bar" + "") + "\n\t") + "a") + (("" + ("hello" + "a")) + ("a" + "foo bar"))) + "hello") + (("foo bar" + "hello") + (("foo bar" + ("hello" + "a")) + "foo bar"))))
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            case "hello":
                                                                switch 0 {
                                                                case -558:
                                                                    if (v6 && (v5 == (v5 - -169))) {
                                                                        var v84 uint = 4294967295
                                                                        if len(v2) > 0 {
                                                                            v2[0] = ("a" + ("a" + ("a" + "")))
                                                                        }
                                                                        v85 := GenS4{F0: ("foo bar" + ("" + "hello")), F1: 236, F2: (v84 ^ (v84 << (((v84 + v84)) & 63)))}
                                                                        var v86 uint = 936
                                                                        var v87 float32 = v4
                                                                        _ = v84
                                                                        _ = v85
                                                                        _ = v86
                                                                        _ = v87
                                                                    }
                                                                case 1:
                                                                    v32 += (v32 + v32)
                                                                }
                                                            }
                                                        } else {
                                                            if ((0 < 2) && (447 > v3)) {
                                                                v5 -= -87
                                                                v32 = v32
                                                                v4 = v4
                                                            }
                                                        }
                                                        v88 := GenS7{F0: 9223372036854775807}
                                                        _ = v88
                                                    }
                                                    _ = fi70
                                                }
                                            } else {
                                                v4 = ((1.5 / (v4 + v4)) - ((v4 + -2.25) / 0.1))
                                                if (1e10 >= v4) {
                                                    if len(v2) > 0 {
                                                        v2[0] = "\n\t"
                                                    }
                                                    for fi89 := 0; fi89 < 2; fi89++ {
                                                        var v90 int = fi89
                                                        if len(v2) > 0 {
                                                            v2[0] = (("foo bar" + "foo bar") + "a")
                                                        }
                                                        v91 := GenS9{F0: 1, F1: (-(v5))}
                                                        _ = fi89
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                } else {
                                                    var v92 uint8 = v1
                                                    for fi93 := 0; fi93 < 5; fi93++ {
                                                        var v94 uint16 = (v32 - v32)
                                                        var v95 uint8 = v92
                                                        switch fi93 {
                                                        case -839:
                                                            var v96 uint = 42
                                                            _ = v96
                                                        case 460:
                                                            var v97 int32 = v3
                                                            _ = v97
                                                        case 42:
                                                            var v98 uint = 42
                                                            _ = v98
                                                        default:
                                                            fmt.Println(v92)
                                                        }
                                                        _ = fi93
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    _ = v92
                                                }
                                            }
                                        } else {
                                            v3 = (v3 * (v3 >> ((18446744073709551615) & 63)))
                                            var v99 bool = ((4294967295 < 255) && ("foo bar" < "a"))
                                            if (1.5 == -2.25) {
                                                fmt.Println(fi52, v99)
                                                v1 = (v1 ^ v1)
                                            }
                                            _ = v99
                                        }
                                        v100 := GenS7{F0: 33}
                                        _ = v100
                                    }
                                case "hello":
                                    if ((false && v6) || v6) {
                                        if ("a" > "a") {
                                            if (0 < -1) {
                                                if v6 {
                                                    var v101 uint16 = (v32 ^ (v32 - v32))
                                                    _ = v101
                                                }
                                            } else {
                                                v3 = (v3 / ((0) | 1))
                                                v102 := GenS10{F0: 1}
                                                _ = v102
                                            }
                                        }
                                        v4 = v4
                                    }
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                var v103 uint64 = 1
                                _ = v103
                            } else {
                                var v104 uint32 = 2
                                fmt.Println(v104, v6, v3)
                                switch "a" {
                                case "a":
                                    for fi105 := 0; fi105 < 7; fi105++ {
                                        var v106 uint16 = (v32 ^ (v32 & 0))
                                        _ = fi105
                                        _ = v106
                                    }
                                case "":
                                    v104 = (v104 - (v104 - (v104 | v104)))
                                default:
                                    if v6 {
                                        v32 = (v32 - (v32 - (v32 >> ((0) & 63))))
                                    } else {
                                        for fi107 := 0; fi107 < 9; fi107++ {
                                            var v108 int = (fi107 & fi52)
                                            v109 := GenS11{F0: -1}
                                            v5 = (v5 >> ((611) & 63))
                                            _ = fi107
                                            _ = v108
                                            _ = v109
                                        }
                                    }
                                }
                                _ = v104
                            }
                        default:
                            v1 = 127
                        }
                        _ = len(v2)
                        for fi110 := 0; fi110 < 3; fi110++ {
                            var v111 int = 32768
                            _ = fi110
                            _ = v111
                        }
                        _ = fi52
                    }
                    v1 -= 42
                }
            case "hello":
                var v112 int = 954
                _ = v112
            case "":
                v1 = (^(v1))
            default:
                v113 := GenS11{F0: 9223372036854775807}
                _ = v113
            }
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v114 := GenS11{F0: 29}
    for fi115 := 0; fi115 < 19; fi115++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi115
    }
    for fi116 := 0; fi116 < 15; fi116++ {
        if v6 {
            v1 += v1
        }
        var v117 int = 32767
        if len(v2) > 0 {
            v2[0] = ((("a" + "hello") + "") + "a")
        }
        _ = fi116
        _ = v117
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
