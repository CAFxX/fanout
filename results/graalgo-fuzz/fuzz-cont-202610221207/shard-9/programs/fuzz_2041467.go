package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int16; F1 int8; F2 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32; F1 int32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string; F1 int32; F2 bool }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int16; F1 int8 }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 int8; F1 uint64; F2 int8 }
type GenS15 struct { F0 uint8; F1 string }
type GenS16 struct { F0 uint64; F1 int8; F2 uint }
type GenS17 struct { F0 int16; F1 int64; F2 int32 }
func genFn0() int32 {
    if (42 == 1) {
        for fi7 := 0; fi7 < 13; fi7++ {
            switch (fi7 * fi7) {
            case 0:
                var v8 int64 = 42
                _ = v8
            case -334:
                var v9 float32 = 1.5
                _ = v9
            default:
                var v10 int64 = 917
                _ = v10
            }
            _ = fi7
        }
        var v11 float32 = 1e10
        _ = v11
    } else {
        if true {
            var v12 uint8 = 42
            _ = v12
        }
    }
    v13 := GenS1{F0: 32767, F1: 0, F2: 0}
    _ = v13
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "hello"
    var v3 int = -668
    v4 := []uint32{}
    var v5 bool = false
    v6 := make(map[string]string)
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS2{F0: 2}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v16 int32 = -943
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int32 = 0
    switch 32768 {
    case 264:
        var v18 int = v3
        _ = v18
    case 0:
        var v19 float64 = -2.25
        _ = v19
    case -652:
        v2 = v2
    default:
        var v20 int32 = (^(v17))
        _ = v20
    }
    switch v2 {
    case "\n\t":
        for fi21 := 0; fi21 < 20; fi21++ {
            v2 = ""
            _ = fi21
        }
    case "a":
        if v5 {
            v2 = v2
            var v22 string = v2
            v23 := GenS3{F0: (2 == 42), F1: 127}
            _ = v22
            _ = v23
        }
    }
    if (v5 || (!(3.14159 > 0.0))) {
        if (127 >= -1) {
            if v5 {
                fmt.Println(v17)
                var v24 uint32 = 2
                _ = v24
            } else {
                for fi25 := 0; fi25 < 11; fi25++ {
                    switch fi25 {
                    case 267:
                        switch (v2 + ((v2 + (("" + ("hello" + "a")) + ("hello" + (v2 + "a")))) + "foo bar")) {
                        case "foo bar":
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                for fi26 := 0; fi26 < 17; fi26++ {
                                    if (49 <= 100) {
                                        switch v2 {
                                        case "":
                                            v3 = v3
                                        case "\n\t":
                                            v27 := GenS4{F0: v2, F1: "\n\t"}
                                            _ = v27
                                        case "a":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        var v28 int8 = 42
                                        v2 = v2
                                        _ = v28
                                    }
                                    fmt.Println(v3, v17)
                                    _ = fi26
                                }
                            case 976:
                                v3 = v3
                            case 513:
                                v5 = v5
                            default:
                                fmt.Println(v3, v3, v2)
                            }
                        case "a":
                            v3 -= fi25
                        }
                    case -586:
                        switch (v2 + ((v2 + v2) + (v2 + ((v2 + ((((((v2 + ("a" + ("foo bar" + ("" + "")))) + v2) + (v2 + ("" + "a"))) + "foo bar") + "hello") + ("foo bar" + "a"))) + "")))) {
                        case "hello":
                            if v5 {
                                v29 := GenS2{F0: 16}
                                _ = v29
                            }
                        case "\n\t":
                            v17 += (-(v17))
                        case "foo bar":
                            v17 -= genFn0()
                        }
                    case 2147483647:
                        v30 := GenS5{F0: 172, F1: -1, F2: 16}
                        _ = v30
                    default:
                        for fi31 := 0; fi31 < 4; fi31++ {
                            if (4294967295 < 4294967295) {
                                if len(v4) > 0 {
                                    v4[0] = 327
                                }
                                fmt.Println(fi31, v3)
                                if ((0 > 2) && (v5 || (!true))) {
                                    if false {
                                        if (65535 != 725) {
                                            var v32 uint32 = 2147483647
                                            if false {
                                                if (!((-405 <= 1) || (995 == 0))) {
                                                    fmt.Println(fi25)
                                                    v32 -= (^(v32))
                                                    var v33 int16 = 5
                                                    _ = v33
                                                } else {
                                                    if v5 {
                                                        v34 := GenS6{F0: 69}
                                                        _ = v34
                                                    } else {
                                                        var v35 int32 = (v17 ^ (v17 - v17))
                                                        v35 = v35
                                                        v36 := GenS7{F0: -32768}
                                                        fmt.Println(v3, v5)
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                    var v37 uint16 = 32767
                                                    v32 = (v32 << ((0) & 63))
                                                    _ = v37
                                                }
                                                v5 = ((float64((0.1 / 3.14159)) - (float64(-2.25) / 1.0)) < ((0.0 + 1.5) * 1e10))
                                                v5 = v5
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            if (18446744073709551615 >= 4294967295) {
                                                fmt.Println(v3, fi31, v2)
                                                v32 -= 1
                                            } else {
                                                if v5 {
                                                    v3 = fi25
                                                } else {
                                                    v3 -= fi25
                                                    var v38 int8 = 100
                                                    var v39 uint64 = 38
                                                    _ = v38
                                                    _ = v39
                                                }
                                                fmt.Println(v3)
                                                var v40 string = ("a" + v2)
                                                _ = v40
                                            }
                                            var v41 uint = 585
                                            _ = v32
                                            _ = v41
                                        }
                                        var v42 float32 = float32((1e10 - (3.14159 * float64(((0.0 * float32(-2.25)) + 1.5)))))
                                        var v43 bool = true
                                        v17 = v17
                                        _ = v42
                                        _ = v43
                                    }
                                    v44 := GenS8{F0: 42, F1: v17}
                                    v45 := v1["a"]
                                    _ = v44
                                    _ = v45
                                }
                            }
                            v5 = v5
                            var v46 float32 = float32(1e10)
                            if v5 {
                                var v47 int8 = -126
                                v48 := GenS9{F0: v5}
                                _ = v47
                                _ = v48
                            }
                            if v5 {
                                var v49 uint16 = 826
                                fmt.Println(v5, v2, v5)
                                var v50 bool = (1 > 2)
                                _ = v49
                                _ = v50
                            } else {
                                v5 = true
                                var v51 int32 = (v17 >> ((593) & 63))
                                switch 448 {
                                case 128:
                                    delete(v6, v2)
                                default:
                                    var v52 int8 = -74
                                    _ = v52
                                }
                                _ = v51
                            }
                            _ = fi31
                            _ = v46
                        }
                    }
                    _ = fi25
                }
                if (((!(255 == 32767)) || (((!(v17 >= 1)) || true) && (!(255 == 321)))) || (!((v2 + "foo bar") != "hello"))) {
                    switch v3 {
                    case 42:
                        v17 += 1
                    case -1:
                        v53 := GenS10{F0: v2, F1: -1, F2: v5}
                        _ = v53
                    default:
                        var v54 uint = 18446744073709551615
                        _ = v54
                    }
                    if (-1 >= 100) {
                        var v55 int32 = (v17 - v17)
                        var v56 float64 = (-2.25 + 1.5)
                        _ = v55
                        _ = v56
                    } else {
                        for fi57 := 0; fi57 < 10; fi57++ {
                            v3 += 4294967295
                            var v58 int64 = 1
                            _ = fi57
                            _ = v58
                        }
                        switch v2 {
                        case "hello":
                            for fi59 := 0; fi59 < 7; fi59++ {
                                if len(v4) > 0 {
                                    v4[0] = 4294967295
                                }
                                _ = fi59
                            }
                        case "":
                            v60 := GenS5{F0: 680, F1: 0, F2: 817}
                            _ = v60
                        default:
                            switch "a" {
                            case "a":
                                var v61 string = ""
                                _ = v61
                            default:
                                if false {
                                    if false {
                                        fmt.Println(v3, v17, v17)
                                        for fi62 := 0; fi62 < 18; fi62++ {
                                            var v63 float32 = float32((3.14159 - 1e10))
                                            v3 -= 0
                                            _ = fi62
                                            _ = v63
                                        }
                                    }
                                    fmt.Println(v2, v3)
                                    if v5 {
                                        v5 = v5
                                    }
                                } else {
                                    v64 := GenS6{F0: 107}
                                    _ = v64
                                }
                            }
                        }
                        var v65 uint32 = 0
                        _ = v65
                    }
                } else {
                    if v5 {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v3, v17)
                        v2 = ((v2 + v2) + "a")
                    }
                }
                var v66 uint16 = 255
                _ = v66
            }
            var v67 uint = 65535
            _ = v67
        }
    } else {
        v5 = (65535 > 144)
        v2 = v2
    }
    v2 = ((v2 + v2) + "foo bar")
    switch (v3 ^ v3) {
    case 9223372036854775807:
        if (!v5) {
            for fi68 := 0; fi68 < 19; fi68++ {
                var v69 int32 = (v17 << ((1) & 63))
                if (-1 != -710) {
                    _ = len(v4)
                } else {
                    v5 = (-1 != 32767)
                    if (1 == 4294967295) {
                        if ((100 == 42) || false) {
                            v17 -= -2147483648
                            v5 = (((0.0 == (3.14159 * (1e10 - (1.5 * 1e10)))) && (false && v5)) || (!(!(!v5))))
                            for fi70 := 0; fi70 < 17; fi70++ {
                                v1[v2] = true
                                switch v2 {
                                case "hello":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "":
                                    if len(v4) > 0 {
                                        v4[0] = 316
                                    }
                                    fmt.Println(v17)
                                    var v71 float32 = 1e10
                                    _ = v71
                                case "a":
                                    fmt.Println(fi68, v17, v2)
                                }
                                v72 := GenS11{F0: (v17 / ((v17) | 1))}
                                v2 = "hello"
                                _ = fi70
                                _ = v72
                            }
                        }
                        var v73 float64 = (1e10 / 1.0)
                        _ = v73
                    } else {
                        for fi74 := 0; fi74 < 10; fi74++ {
                            var v75 uint = 744
                            fmt.Println(v75, v69)
                            _ = fi74
                            _ = v75
                        }
                    }
                }
                _ = fi68
                _ = v69
            }
        }
    }
    switch (v3 << ((2) & 63)) {
    case 230:
        var v76 float32 = (0.0 - 1e10)
        _ = v76
    default:
        var v77 uint8 = 42
        _ = v77
    }
    v4 = append(v4, 343)
    v78 := GenS10{F0: v2, F1: v17, F2: (false && (0.1 != ((0.1 * ((3.14159 * 1.5) / 1.0)) - float64((1.5 + 0.1)))))}
    switch v2 {
    case "foo bar":
        var v79 int32 = v17
        _ = v79
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "a":
        switch (v3 & 128) {
        case 157:
            switch v2 {
            case "foo bar":
                if false {
                    var v80 int16 = -710
                    fmt.Println(v5, v2)
                    switch v2 {
                    case "\n\t":
                        for fi81 := 0; fi81 < 17; fi81++ {
                            if ((42 <= 643) && (!(!(65535 < 310)))) {
                                v82 := GenS2{F0: 510}
                                _ = v82
                            }
                            v2 = "foo bar"
                            v83 := GenS10{F0: v2, F1: (^(v17)), F2: v5}
                            _ = fi81
                            _ = v83
                        }
                    case "foo bar":
                        var v84 int16 = v80
                        _ = v84
                    case "a":
                        _ = len(v4)
                    default:
                        delete(v1, v2)
                    }
                    _ = v80
                } else {
                    v1[((("\n\t" + "") + ("a" + (v2 + ""))) + (v2 + v2))] = v5
                    switch (((v2 + ("foo bar" + "\n\t")) + (v2 + (("hello" + ((v2 + "foo bar") + "hello")) + v2))) + "\n\t") {
                    case "\n\t":
                        fmt.Println(v3)
                    }
                    v3 -= v3
                }
            }
        case 705:
            switch -398 {
            case 2147483647:
                fmt.Println(v2, v2, v2)
            case -9223372036854775808:
                if false {
                    v85 := GenS4{F0: v2, F1: ""}
                    var v86 string = v2
                    _ = v85
                    _ = v86
                }
            case -727:
                if v5 {
                    v2 = ("hello" + (v2 + (v2 + "hello")))
                    v87 := v6[v2]
                    _ = v87
                }
            }
        }
    }
    var v88 float64 = (0.1 / 3.14159)
    var v89 float32 = 0.1
    fmt.Println(v2, v3, v5)
    if (1 > 31) {
        switch (v3 / (((v3 - -1)) | 1)) {
        case 0:
            v88 = (v88 / v88)
        case 32768:
            _ = len(v1)
        }
        v4 = append(v4, 600)
        var v90 uint = 496
        _ = v90
    } else {
        v91 := GenS12{F0: 1, F1: -64}
        var v92 uint16 = 613
        v93 := GenS13{F0: v5, F1: 0}
        _ = v91
        _ = v92
        _ = v93
    }
    for fi94 := 0; fi94 < 16; fi94++ {
        for fi95 := 0; fi95 < 17; fi95++ {
            v5 = true
            var v96 uint = 107
            _ = fi95
            _ = v96
        }
        fmt.Println(v5)
        var v97 float64 = 0.0
        _ = fi94
        _ = v97
    }
    for fi98 := 0; fi98 < 13; fi98++ {
        v2 = (v2 + v2)
        if v5 {
            fmt.Println(v88, v3, fi98)
            switch -2147483648 {
            case 1:
                for fi99 := 0; fi99 < 10; fi99++ {
                    switch (v2 + v2) {
                    case "hello":
                        v89 = (float32(-2.25) - (v89 * float32(v88)))
                    default:
                        var v100 int = (^(fi99))
                        _ = v100
                    }
                    var v101 int32 = 2147483647
                    if (1 == 42) {
                        _ = len(v6)
                    } else {
                        fmt.Println(fi99)
                    }
                    _ = fi99
                    _ = v101
                }
            case 9223372036854775807:
                v17 -= (v17 - (v17 / ((573) | 1)))
            }
            var v102 uint32 = 4294967295
            _ = v102
        }
        _ = fi98
    }
    switch -9223372036854775808 {
    case 399:
        var v103 bool = v5
        _ = v103
    case 32767:
        if v5 {
            v88 += (((v88 + v88) * v88) / v88)
            if v5 {
                switch ((((v2 + ("a" + (v2 + v2))) + v2) + (v2 + v2)) + "") {
                case "foo bar":
                    var v104 float32 = v89
                    _ = v104
                case "\n\t":
                    var v105 string = "\n\t"
                    _ = v105
                case "":
                    var v106 uint16 = 42
                    _ = v106
                }
                switch ((((v2 + "\n\t") + v2) + (((("\n\t" + v2) + (v2 + (v2 + (("a" + v2) + v2)))) + "") + "hello")) + "\n\t") {
                case "":
                    if (false || (1 < 115)) {
                        for fi107 := 0; fi107 < 10; fi107++ {
                            v108 := GenS9{F0: v5}
                            if false {
                                v109 := GenS14{F0: 1, F1: 42, F2: 127}
                                delete(v1, v2)
                                for fi110 := 0; fi110 < 4; fi110++ {
                                    if v5 {
                                        v89 += ((0.1 - (float32(v88) + (v89 / v89))) * (v89 - 0.1))
                                        v111 := GenS9{F0: v5}
                                        switch -2147483648 {
                                        case 9223372036854775807:
                                            v88 -= 0.0
                                        default:
                                            v112 := GenS15{F0: 255, F1: v2}
                                            _ = v112
                                        }
                                        _ = v111
                                    } else {
                                        fmt.Println(v5)
                                    }
                                    _ = fi110
                                }
                                _ = v109
                            } else {
                                for fi113 := 0; fi113 < 1; fi113++ {
                                    v2 = "a"
                                    v114 := GenS16{F0: 2, F1: -111, F2: 257}
                                    _ = fi113
                                    _ = v114
                                }
                                _ = len(v4)
                            }
                            _ = fi107
                            _ = v108
                        }
                    } else {
                        var v115 int32 = genFn0()
                        _ = v115
                    }
                case "\n\t":
                    for fi116 := 0; fi116 < 6; fi116++ {
                        fmt.Println(v3)
                        v89 += float32(((1.5 - v88) / 1.0))
                        _ = fi116
                    }
                case "foo bar":
                    v89 = 3.14159
                }
            }
            for fi117 := 0; fi117 < 16; fi117++ {
                if (1.5 > v89) {
                    var v118 uint64 = 761
                    _ = v118
                } else {
                    var v119 string = (v2 + "")
                    for fi120 := 0; fi120 < 19; fi120++ {
                        fmt.Println(v89, v17, fi120)
                        _ = fi120
                    }
                    _ = v119
                }
                fmt.Println(v2)
                v121 := GenS17{F0: 0, F1: 1, F2: v17}
                _ = fi117
                _ = v121
            }
        }
    default:
        var v122 bool = (v3 == (-(v3)))
        _ = v122
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v88
    _ = v89
}
