package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int; F1 int64; F2 int16 }
type GenS12 struct { F0 uint16; F1 int; F2 uint16 }
func genFn0(p0 uint) uint8 {
    p0 -= p0
    v7 := GenS1{F0: -747}
    v8 := []int{-9223372036854775808, 4294967295, 128, 9223372036854775807}
    switch "foo bar" {
    case "a":
        var v9 uint = (p0 / ((p0) | 1))
        _ = v9
    default:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint16) uint16 {
    switch "" {
    case "foo bar":
        v10 := GenS1{F0: -2147483648}
        _ = v10
    default:
        if false {
            var v11 int = 4294967295
            v11 -= (v11 ^ v11)
            _ = v11
        } else {
            if false {
                fmt.Println(p0)
                v12 := GenS2{F0: 87, F1: 910, F2: p0}
                var v13 uint64 = 2
                _ = v12
                _ = v13
            }
            v14 := []int{42, 862, -2147483648}
            var v15 string = "hello"
            _ = v14
            _ = v15
        }
    }
    switch 42 {
    case 784:
        var v16 uint32 = 974
        _ = v16
    case 0:
        var v17 int = 144
        _ = v17
    default:
        var v18 float64 = float64(0.0)
        _ = v18
    }
    p0 -= 32767
    v19 := []int{-607, 55, 128}
    _ = p0
    _ = v19
    return (p0 | (p0 << ((18446744073709551615) & 63)))
}

func main() {
    var v1 bool = (42 <= 218)
    v2 := []uint32{42, 509, 4294967295}
    var v3 float64 = 0.0
    v4 := make(map[int]string)
    var v5 uint32 = 2147483648
    var v6 float64 = (v3 + -2.25)
    fmt.Println(v1)
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS3{F0: 0, F1: 9223372036854775807, F2: 1}
        v22 := GenS3{F0: -9223372036854775808, F1: 42, F2: 42}
        var v23 int8 = 42
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    fmt.Println(v3, v1, v3)
    for fi24 := 0; fi24 < 11; fi24++ {
        fmt.Println(v6, v3, v6)
        fmt.Println(v1, fi24)
        _ = fi24
    }
    v25 := GenS2{F0: 127, F1: -32768, F2: 255}
    v26 := GenS1{F0: -665}
    if ((42 >= 42) && v1) {
        v3 = v6
        v1 = (-1 <= 127)
    } else {
        for fi27 := 0; fi27 < 1; fi27++ {
            if (float32(v3) <= 1.5) {
                v1 = (42 >= 0)
                _ = len(v4)
                v5 += v5
            } else {
                v1 = v1
                switch (("\n\t" + "hello") + "") {
                case "\n\t":
                    for fi28 := 0; fi28 < 5; fi28++ {
                        var v29 int64 = 42
                        v30 := GenS2{F0: 1, F1: 42, F2: 2}
                        _ = len(v4)
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case "a":
                    v31 := GenS1{F0: 2147483647}
                    _ = v31
                }
            }
            if (42 <= -2147483648) {
                var v32 bool = v1
                if (false || (v1 || false)) {
                    var v33 uint16 = genFn1(65535)
                    for fi34 := 0; fi34 < 13; fi34++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi34
                    }
                    _ = v33
                }
                _ = v32
            } else {
                for fi35 := 0; fi35 < 7; fi35++ {
                    var v36 int32 = 1
                    if v1 {
                        if v1 {
                            v4[1] = "a"
                        } else {
                            var v37 int32 = 42
                            switch ("" + "hello") {
                            case "":
                                v38 := GenS4{F0: int32(fi35)}
                                _ = v38
                            case "foo bar":
                                if (!v1) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v39 int8 = -77
                                        _ = v39
                                    case "a":
                                        _ = len(v4)
                                    case "foo bar":
                                        if (v1 && (!true)) {
                                            v6 = (-2.25 + v3)
                                            _ = len(v4)
                                        } else {
                                            switch fi27 {
                                            case 2147483647:
                                                if v1 {
                                                    v36 = v36
                                                } else {
                                                    v4[(fi27 * fi27)] = "hello"
                                                    if v1 {
                                                        var v40 string = ""
                                                        _ = v40
                                                    }
                                                }
                                            case 546:
                                                if false {
                                                    if v1 {
                                                        if v1 {
                                                            if false {
                                                                var v41 int32 = 2147483647
                                                                v42 := GenS4{F0: v41}
                                                                var v43 int64 = 42
                                                                _ = v41
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                v2 = append(v2, 0)
                                                            }
                                                            v3 += 1.5
                                                            fmt.Println(v1)
                                                        } else {
                                                            switch -188 {
                                                            case 9223372036854775807:
                                                                var v44 int = fi35
                                                                _ = v44
                                                            case -896:
                                                                if true {
                                                                    v36 -= v36
                                                                    if (!(0 < 1)) {
                                                                        v4[fi35] = ("a" + (("hello" + "\n\t") + "hello"))
                                                                        fmt.Println(v36, fi35)
                                                                    }
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v1 = v1
                                                                    case "a":
                                                                        v45 := v4[(fi35 * (fi27 - 4294967295))]
                                                                        _ = v45
                                                                    case "":
                                                                        if ((1 == -9223372036854775808) || (!false)) {
                                                                            _ = len(v4)
                                                                            var v46 bool = (32767 == 0)
                                                                            fmt.Println(v3)
                                                                            if (21 > genFn0(255)) {
                                                                                if v1 {
                                                                                    v5 -= v5
                                                                                    if true {
                                                                                        fmt.Println(v3)
                                                                                    }
                                                                                }
                                                                                v46 = false
                                                                            }
                                                                            _ = v46
                                                                        } else {
                                                                            fmt.Println(v36)
                                                                            v5 += 375
                                                                            v47 := GenS3{F0: 42, F1: -9223372036854775808, F2: genFn0(793)}
                                                                            _ = v47
                                                                        }
                                                                    default:
                                                                        switch "hello" {
                                                                        case "":
                                                                            v48 := GenS1{F0: -934}
                                                                            _ = v48
                                                                        }
                                                                    }
                                                                    v49 := GenS4{F0: (v37 + (v36 * (v36 * (-(v36)))))}
                                                                    _ = v49
                                                                }
                                                            case 1:
                                                                var v50 int8 = 1
                                                                _ = v50
                                                            default:
                                                                var v51 float32 = ((3.14159 * 0.0) - ((-2.25 + 1.5) + 3.14159))
                                                                _ = v51
                                                            }
                                                        }
                                                        fmt.Println(v3, fi27)
                                                        var v52 int64 = -9223372036854775808
                                                        _ = v52
                                                    }
                                                    v53 := GenS5{F0: -1, F1: 2}
                                                    v54 := GenS6{F0: 637}
                                                    var v55 int64 = 1
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    v56 := GenS2{F0: genFn0(42), F1: 42, F2: 1}
                                                    var v57 int8 = 19
                                                    var v58 int64 = -9223372036854775808
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                }
                                            default:
                                                v1 = v1
                                            }
                                        }
                                    }
                                    if (2 > 715) {
                                        v36 += (v37 << ((uint(v5)) & 63))
                                        var v59 float64 = (((((v3 - 3.14159) - float64(((((0.1 / 1e10) * (3.14159 * 1.5)) + float32(v3)) + (0.1 - (((((3.14159 + (-2.25 * (((((1e10 + (1.5 - 3.14159)) / 3.14159) / ((float32(v6) * 3.14159) - (((((0.1 + 1.5) / float32(3.14159)) / -2.25) * 0.0) / ((-2.25 + 3.14159) + 1.5)))) + 3.14159) + 3.14159))) * 0.1) + (-2.25 - 3.14159)) + -2.25) - -2.25))))) + (1e10 / v6)) * ((1.5 * 1e10) / v3)) + ((float64(((3.14159 / ((0.1 - 1.5) / 1.5)) + float32(v3))) + v6) / (0.1 - 0.0)))
                                        _ = v59
                                    } else {
                                        v60 := GenS7{F0: genFn1(65535), F1: v36}
                                        var v61 int16 = 1
                                        var v62 int = 0
                                        switch (fi35 ^ fi35) {
                                        case 42:
                                            if v1 {
                                                fmt.Println(v36)
                                                v36 = (v36 + (v36 / ((v37) | 1)))
                                                v3 += (1.5 / ((float64(3.14159) * (((float64((-2.25 + (((1.5 * 0.0) - -2.25) / -2.25))) - 0.1) - (v3 + v3)) - -2.25)) / 3.14159))
                                            }
                                            var v63 uint64 = 4294967295
                                            v63 = (v63 ^ (v63 ^ 595))
                                            _ = v63
                                        case 1:
                                            fmt.Println(v6, v37)
                                        default:
                                            var v64 uint32 = (v5 & v5)
                                            _ = v64
                                        }
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    }
                                } else {
                                    v65 := GenS8{F0: (v36 + v36)}
                                    _ = v65
                                }
                            }
                            var v66 float32 = 1e10
                            _ = v37
                            _ = v66
                        }
                        switch "hello" {
                        case "a":
                            switch "foo bar" {
                            case "foo bar":
                                switch fi35 {
                                case -1:
                                    var v67 uint16 = genFn1(255)
                                    _ = v67
                                case 4294967295:
                                    v5 += v5
                                }
                            case "hello":
                                if false {
                                    var v68 uint = 1
                                    switch 128 {
                                    case 0:
                                        v6 = v3
                                    default:
                                        v4[32768] = (((("" + "foo bar") + ("hello" + ("a" + (("hello" + "a") + ("hello" + "\n\t"))))) + ((("hello" + "") + "\n\t") + ("foo bar" + "a"))) + ((((("\n\t" + "\n\t") + "") + ("foo bar" + ("hello" + (("" + "hello") + ("\n\t" + "a"))))) + ("\n\t" + "\n\t")) + ("\n\t" + ("" + "a"))))
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v68
                                } else {
                                    var v69 int = (fi35 / ((fi27) | 1))
                                    var v70 uint = 816
                                    _ = v69
                                    _ = v70
                                }
                            default:
                                _ = len(v4)
                            }
                        case "foo bar":
                            v71 := GenS5{F0: 0, F1: 2}
                            _ = v71
                        }
                        var v72 uint = 18446744073709551615
                        _ = v72
                    }
                    if false {
                        if v1 {
                            var v73 uint = 320
                            if false {
                                v1 = v1
                                delete(v4, fi35)
                            }
                            fmt.Println(v5, v1, v1)
                            _ = v73
                        }
                        v4[(^(fi27))] = (("\n\t" + "\n\t") + ("hello" + ("hello" + ((("hello" + "") + "foo bar") + ("" + "\n\t")))))
                        fmt.Println(v3, v3, v3)
                    } else {
                        v74 := GenS9{F0: (v5 - (v5 >> ((2) & 63))), F1: (v1 || v1), F2: v5}
                        _ = v74
                    }
                    _ = fi35
                    _ = v36
                }
                v75 := GenS7{F0: 299, F1: 42}
                _ = v75
            }
            for fi76 := 0; fi76 < 15; fi76++ {
                fmt.Println(fi27, fi76, v3)
                _ = fi76
            }
            _ = fi27
        }
        v5 = 2147483647
    }
    if true {
        for fi77 := 0; fi77 < 1; fi77++ {
            for fi78 := 0; fi78 < 17; fi78++ {
                var v79 uint8 = 127
                var v80 int = fi77
                if (0 != 1) {
                    fmt.Println(fi78, v80, fi77)
                    v2 = append(v2, v5)
                    v79 -= (v79 - (v79 / (((v79 % (((v79 >> ((261) & 63))) | 1))) | 1)))
                } else {
                    v79 += (v79 * (v79 % ((v79) | 1)))
                    v3 += v3
                }
                _ = fi78
                _ = v79
                _ = v80
            }
            var v81 string = ("a" + "hello")
            v4[32768] = ((("foo bar" + "") + ((v81 + (v81 + v81)) + ((v81 + v81) + ("\n\t" + v81)))) + v81)
            _ = fi77
            _ = v81
        }
        if false {
            switch 4294967295 {
            case 0:
                v6 += float64(-2.25)
            case -9223372036854775808:
                v5 = 2147483647
            default:
                var v82 int = -9223372036854775808
                _ = v82
            }
            switch "hello" {
            case "hello":
                switch -377 {
                case 860:
                    v6 = v6
                case 4294967295:
                    v6 = -2.25
                }
            case "a":
                var v83 float64 = float64(float32(0.0))
                _ = v83
            case "\n\t":
                fmt.Println(v6, v3, v3)
            }
        }
    }
    delete(v4, -822)
    fmt.Println(v6)
    var v84 int8 = 100
    if (-9223372036854775808 != 535) {
        if false {
            v4[-2147483648] = "a"
            var v85 uint64 = 2
            fmt.Println(v3)
            _ = v85
        } else {
            v6 = (v6 * 1.5)
        }
        switch "\n\t" {
        case "a":
            for fi86 := 0; fi86 < 15; fi86++ {
                v87 := GenS10{F0: ""}
                for fi88 := 0; fi88 < 17; fi88++ {
                    fmt.Println(v5)
                    v2 = append(v2, v5)
                    _ = fi88
                }
                for fi89 := 0; fi89 < 9; fi89++ {
                    var v90 int64 = -9223372036854775808
                    var v91 int = fi89
                    v92 := v4[v91]
                    _ = fi89
                    _ = v90
                    _ = v91
                    _ = v92
                }
                _ = fi86
                _ = v87
            }
        case "foo bar":
            v1 = (!v1)
        case "":
            v6 = v6
        default:
            switch ((("foo bar" + ("foo bar" + ("hello" + ("" + "hello")))) + "") + ("hello" + (("" + "foo bar") + "foo bar"))) {
            case "hello":
                var v93 int32 = -473
                _ = v93
            case "foo bar":
                if v1 {
                    fmt.Println(v1, v6)
                    var v94 uint16 = 918
                    v5 += 4294967295
                    _ = v94
                } else {
                    var v95 int = -255
                    v96 := GenS4{F0: -184}
                    switch (^(v95)) {
                    case -531:
                        if (!((-701 != 42) || ((("" + ("a" + "\n\t")) > "") || v1))) {
                            var v97 uint64 = 1
                            _ = v97
                        } else {
                            for fi98 := 0; fi98 < 5; fi98++ {
                                fmt.Println(v3, v3, v5)
                                for fi99 := 0; fi99 < 18; fi99++ {
                                    v6 = float64(((0.0 - float32(((v6 - (v3 + (-2.25 - (1.5 * (v3 - (v3 / v6)))))) / (0.0 + (v6 / v3))))) * -2.25))
                                    var v100 uint8 = 255
                                    _ = fi99
                                    _ = v100
                                }
                                _ = fi98
                            }
                        }
                    case 128:
                        switch v95 {
                        case -1:
                            switch v95 {
                            case 652:
                                v5 += 42
                            default:
                                var v101 int8 = (v84 - v84)
                                _ = v101
                            }
                        case 0:
                            for fi102 := 0; fi102 < 8; fi102++ {
                                if (42 != 255) {
                                    v1 = v1
                                    var v103 int8 = v84
                                    if (127 < 15) {
                                        v3 = v3
                                    }
                                    _ = v103
                                }
                                v104 := GenS11{F0: v95, F1: 9223372036854775807, F2: 249}
                                _ = fi102
                                _ = v104
                            }
                        case 895:
                            v1 = ((v1 && true) && v1)
                        default:
                            for fi105 := 0; fi105 < 12; fi105++ {
                                for fi106 := 0; fi106 < 14; fi106++ {
                                    v107 := GenS6{F0: v5}
                                    v3 = v3
                                    _ = fi106
                                    _ = v107
                                }
                                v108 := GenS1{F0: 0}
                                _ = fi105
                                _ = v108
                            }
                        }
                    case 127:
                        v109 := GenS9{F0: 2, F1: (-9223372036854775808 < 9223372036854775807), F2: v5}
                        _ = v109
                    default:
                        if false {
                            if (("a" + "foo bar") == "") {
                                v6 += ((1.5 * 1e10) / 3.14159)
                                v110 := GenS12{F0: 32767, F1: -980, F2: 398}
                                _ = v110
                            }
                            for fi111 := 0; fi111 < 16; fi111++ {
                                var v112 int64 = -600
                                if (1 <= v112) {
                                    v113 := GenS4{F0: 42}
                                    var v114 string = "\n\t"
                                    _ = v113
                                    _ = v114
                                }
                                _ = fi111
                                _ = v112
                            }
                            for fi115 := 0; fi115 < 4; fi115++ {
                                v3 += -2.25
                                for fi116 := 0; fi116 < 5; fi116++ {
                                    v117 := GenS3{F0: -672, F1: 9223372036854775807, F2: 0}
                                    _ = fi116
                                    _ = v117
                                }
                                var v118 int = 4294967295
                                _ = fi115
                                _ = v118
                            }
                        } else {
                            v3 = 1e10
                        }
                    }
                    _ = v95
                    _ = v96
                }
            default:
                var v119 uint = 255
                _ = v119
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v84
}
