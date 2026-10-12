package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 bool }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 bool; F1 int8; F2 int32 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 uint16; F1 int64; F2 bool }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 uint; F1 bool }
type GenS11 struct { F0 uint32; F1 bool }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint }
func genFn0(p0 uint, p1 int64) int64 {
    v5 := GenS1{F0: 63, F1: 2, F2: (32767 < 65535)}
    p0 -= (p0 | p0)
    var v6 bool = true
    p1 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 42
    var v3 int32 = 1
    var v4 int16 = -299
    v4 -= 1000
    fmt.Println(v3)
    v3 = (^(v3))
    v2 += int(v4)
    v7 := GenS2{F0: 1, F1: 32767, F2: 155}
    v8 := []int{127, -353, (^(v2))}
    _ = len(v8)
    v2 += v2
    if ((!false) && true) {
        v9 := GenS1{F0: 127, F1: 65535, F2: false}
        _ = len(v1)
        v2 = (v2 ^ 335)
        _ = v9
    } else {
        if ((v3 * v3) > v3) {
            var v10 float32 = (-2.25 + float32(float64(float32(0.0))))
            if false {
                for fi11 := 0; fi11 < 10; fi11++ {
                    fmt.Println(v3, v10, v3)
                    fmt.Println(v4, v2)
                    var v12 uint16 = 42
                    _ = fi11
                    _ = v12
                }
                v13 := GenS3{F0: true, F1: 0, F2: -542}
                v10 -= (float32(0.1) / v10)
                _ = v13
            } else {
                v4 = v4
            }
            if ("hello" >= "hello") {
                var v14 uint8 = 153
                _ = v14
            } else {
                if true {
                    switch "foo bar" {
                    case "a":
                        var v15 uint16 = 337
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 6; fi16++ {
                            v3 = 1
                            _ = fi16
                        }
                    }
                } else {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v18 := GenS4{F0: genFn0(255, genFn0(4294967295, genFn0(18446744073709551615, -328))), F1: v3}
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v8 = append(v8, (v2 ^ v2))
                            if true {
                                if (!(42 == 42)) {
                                    v2 += 9223372036854775807
                                    switch (^(fi19)) {
                                    case 32768:
                                        switch "a" {
                                        case "foo bar":
                                            v2 = fi17
                                        case "\n\t":
                                            var v20 uint32 = 0
                                            _ = v20
                                        }
                                    case 42:
                                        if false {
                                            v10 = (float32((1e10 + -2.25)) + v10)
                                        } else {
                                            _ = len(v8)
                                            v2 = (v2 - 316)
                                            v21 := GenS4{F0: 768, F1: v3}
                                            _ = v21
                                        }
                                        var v22 int64 = -670
                                        v3 = (v3 ^ (-(v3)))
                                        _ = v22
                                    case -1:
                                        var v23 float32 = v10
                                        _ = v23
                                    }
                                    var v24 uint16 = 1
                                    _ = v24
                                }
                            }
                            var v25 float64 = -2.25
                            _ = fi19
                            _ = v25
                        }
                        v10 = v10
                        _ = fi17
                        _ = v18
                    }
                    v26 := GenS5{F0: 127, F1: -1}
                    var v27 uint32 = 1
                    _ = v26
                    _ = v27
                }
                v28 := GenS6{F0: 799}
                _ = v28
            }
            _ = v10
        } else {
            switch (v2 | (v2 * 635)) {
            case -9223372036854775808:
                v1["\n\t"] = (1 < 65535)
            case 4294967295:
                v3 = 2147483647
            case 42:
                v3 -= v3
            default:
                fmt.Println(v4)
            }
            var v29 uint8 = 255
            _ = v29
        }
        var v30 uint32 = 2
        _ = v30
    }
    if false {
        fmt.Println(v2, v2)
        var v31 float32 = ((3.14159 * 0.1) / float32(1e10))
        v32 := v1[("\n\t" + "a")]
        _ = v31
        _ = v32
    } else {
        for fi33 := 0; fi33 < 18; fi33++ {
            v2 -= -9223372036854775808
            _ = fi33
        }
        switch "a" {
        case "":
            v4 -= 661
        default:
            for fi34 := 0; fi34 < 1; fi34++ {
                if (false || (127 >= (fi34 / ((660) | 1)))) {
                    var v35 int64 = 42
                    _ = v35
                } else {
                    var v36 int = v2
                    if (!true) {
                        for fi37 := 0; fi37 < 10; fi37++ {
                            if (32767 <= 0) {
                                if false {
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(fi34, v2)
                                    case "a":
                                        v38 := GenS3{F0: ((v2 % ((2147483647) | 1)) < -9223372036854775808), F1: 1, F2: (v3 | (v3 - -1))}
                                        _ = v38
                                    case "foo bar":
                                        var v39 int = (fi37 / (((^(fi37))) | 1))
                                        _ = v39
                                    default:
                                        v36 += v2
                                    }
                                    if (1 > 4294967295) {
                                        var v40 int8 = -125
                                        if (!(!false)) {
                                            fmt.Println(v2, v2, v40)
                                            var v41 bool = (true || (2 <= 0))
                                            switch ("" + "a") {
                                            case "foo bar":
                                                fmt.Println(v4, fi37)
                                            case "\n\t":
                                                fmt.Println(fi34, fi37, v2)
                                            default:
                                                v42 := GenS7{F0: (v3 * -724), F1: 22}
                                                _ = v42
                                            }
                                            _ = v41
                                        }
                                        delete(v1, ("a" + ("hello" + "hello")))
                                        _ = v40
                                    } else {
                                        v4 = 1
                                        if (true || (707 <= (v3 / (((v3 * (v3 + v3))) | 1)))) {
                                            if false {
                                                v2 -= v2
                                                var v43 int64 = 9223372036854775807
                                                _ = v43
                                            } else {
                                                var v44 bool = false
                                                var v45 uint32 = 1
                                                v46 := GenS8{F0: 423, F1: 42, F2: (-9223372036854775808 < 0)}
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            }
                                        }
                                        v2 = 310
                                        v2 = (fi37 ^ -573)
                                        var v47 int32 = (v3 >> ((446) & 63))
                                        v48 := GenS9{F0: 9223372036854775807, F1: 1}
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                if ((float32(0.0) * 3.14159) >= 3.14159) {
                                    fmt.Println(v36)
                                    var v49 uint16 = 32767
                                    if ((v2 + v2) == v36) {
                                        var v50 float64 = float64(3.14159)
                                        _ = v50
                                    } else {
                                        if (!(true || true)) {
                                            switch 4294967295 {
                                            case 4294967295:
                                                if false {
                                                    _ = len(v1)
                                                    var v51 int32 = (v3 - -596)
                                                    _ = v51
                                                }
                                            case 0:
                                                if true {
                                                    if false {
                                                        fmt.Println(fi34, fi34)
                                                        v4 += v4
                                                    } else {
                                                        v52 := GenS10{F0: 635, F1: (0.1 == -2.25)}
                                                        _ = v52
                                                    }
                                                    if false {
                                                        var v53 int64 = 9223372036854775807
                                                        v4 -= (v4 * v4)
                                                        v4 = v4
                                                        v54 := v1[("a" + "\n\t")]
                                                        fmt.Println(v4)
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        v1["foo bar"] = (42 < genFn0(27, 1))
                                                        fmt.Println(fi37, v3, v49)
                                                        var v55 uint16 = (v49 | v49)
                                                        _ = v55
                                                    }
                                                    var v56 int32 = v3
                                                    _ = v56
                                                }
                                            case 2147483647:
                                                v57 := v1[("\n\t" + "hello")]
                                                _ = v57
                                            default:
                                                v58 := v1[((("a" + "foo bar") + ("" + "foo bar")) + (("\n\t" + "foo bar") + "\n\t"))]
                                                _ = v58
                                            }
                                            v1[""] = true
                                            v59 := GenS8{F0: v49, F1: 1, F2: (v4 <= v4)}
                                            _ = v59
                                        }
                                        var v60 uint16 = v49
                                        _ = v60
                                    }
                                    switch (v2 * 0) {
                                    case -413:
                                        var v61 float32 = 0.1
                                        _ = v61
                                    default:
                                        var v62 int64 = -1
                                        _ = v62
                                    }
                                    v4 = v4
                                    _ = v49
                                } else {
                                    fmt.Println(v3)
                                }
                                v2 = v36
                                switch "a" {
                                case "foo bar":
                                    var v63 uint = 457
                                    _ = v63
                                default:
                                    var v64 uint64 = 0
                                    _ = v64
                                }
                            }
                            _ = fi37
                        }
                        var v65 int64 = 9223372036854775807
                        _ = v65
                    } else {
                        if (((false || ((!false) && ((v4 >> ((65535) & 63)) != (v4 % ((32767) | 1))))) || false) && true) {
                            var v66 int32 = (v3 ^ (-(v3)))
                            if (4294967295 <= 455) {
                                v36 -= -325
                                var v67 int8 = -128
                                _ = v67
                            }
                            if (-1 == 42) {
                                _ = len(v8)
                                v66 += (v3 - 1)
                                if (-25 != -905) {
                                    var v68 uint = 255
                                    _ = v68
                                }
                            } else {
                                switch "" {
                                case "foo bar":
                                    switch "" {
                                    case "foo bar":
                                        var v69 bool = false
                                        _ = v69
                                    case "hello":
                                        if (2 != 2) {
                                            if (!false) {
                                                v36 = v36
                                                fmt.Println(fi34, fi34)
                                            } else {
                                                if (!(!false)) {
                                                    for fi70 := 0; fi70 < 2; fi70++ {
                                                        v66 = v66
                                                        _ = fi70
                                                    }
                                                }
                                                for fi71 := 0; fi71 < 16; fi71++ {
                                                    if (!(-1 < (-(v4)))) {
                                                        v36 = v36
                                                        var v72 int64 = 0
                                                        _ = v72
                                                    } else {
                                                        v8 = append(v8, (fi71 & 938))
                                                        fmt.Println(v3)
                                                    }
                                                    v73 := GenS11{F0: 0, F1: (562 > 4294967295)}
                                                    _ = fi71
                                                    _ = v73
                                                }
                                                switch (v36 / ((v2) | 1)) {
                                                case 0:
                                                    var v74 int64 = -9223372036854775808
                                                    _ = v74
                                                case 31:
                                                    if len(v8) > 0 {
                                                        v8[0] = fi34
                                                    }
                                                default:
                                                    v4 -= (v4 & v4)
                                                }
                                            }
                                        } else {
                                            switch (fi34 + (v36 / ((v2) | 1))) {
                                            case -2147483648:
                                                v36 = (-(v2))
                                            case 658:
                                                var v75 int64 = -9223372036854775808
                                                _ = v75
                                            case 128:
                                                for fi76 := 0; fi76 < 10; fi76++ {
                                                    v77 := GenS12{F0: 1}
                                                    v2 = v2
                                                    var v78 int16 = (-(v4))
                                                    _ = fi76
                                                    _ = v77
                                                    _ = v78
                                                }
                                            default:
                                                fmt.Println(v3, v3)
                                            }
                                            fmt.Println(v36, v66)
                                            v3 += v3
                                        }
                                    default:
                                        var v79 float64 = 1.5
                                        _ = v79
                                    }
                                }
                                switch v2 {
                                case 9223372036854775807:
                                    v36 = 4294967295
                                case -9223372036854775808:
                                    fmt.Println(fi34, v2)
                                case 32767:
                                    v80 := GenS8{F0: 146, F1: 1, F2: ((2 < 0) || false)}
                                    _ = v80
                                }
                                v4 = v4
                            }
                            _ = v66
                        } else {
                            v81 := GenS2{F0: 2, F1: 32767, F2: 648}
                            var v82 uint16 = 65535
                            v83 := GenS10{F0: 42, F1: (731 == 1)}
                            _ = v81
                            _ = v82
                            _ = v83
                        }
                    }
                    _ = v36
                }
                _ = fi34
            }
        }
        v2 = (v2 + (v2 ^ 127))
    }
    v2 -= v2
    var v84 float64 = ((((1e10 - float64(((1.5 * 0.0) - (0.0 + float32(3.14159))))) * 0.0) / ((1.5 + (0.1 * 1.5)) / 1.0)) + 3.14159)
    _ = len(v1)
    for fi85 := 0; fi85 < 18; fi85++ {
        v4 -= v4
        if len(v8) > 0 {
            v8[0] = 42
        }
        if ((!false) || ((!(0 >= 42)) || (!true))) {
            if (255 != 0) {
                var v86 uint64 = 2
                v87 := GenS11{F0: 265, F1: ((v2 * 0) >= (v2 * (v2 >> ((409) & 63))))}
                v86 -= v86
                _ = v86
                _ = v87
            } else {
                v84 -= v84
                if true {
                    switch "a" {
                    case "a":
                        for fi88 := 0; fi88 < 10; fi88++ {
                            var v89 int = v2
                            if true {
                                v2 += 9223372036854775807
                            }
                            fmt.Println(fi85, fi88)
                            if (18446744073709551615 != 2) {
                                v90 := GenS9{F0: 0, F1: (^(v3))}
                                v8 = append(v8, fi85)
                                _ = v90
                            } else {
                                v84 = float64(-2.25)
                                fmt.Println(v4, v4, v4)
                                fmt.Println(v84, v2)
                            }
                            _ = fi88
                            _ = v89
                        }
                    case "hello":
                        if (false || ((9223372036854775807 == 0) || false)) {
                            v91 := v1[("\n\t" + "\n\t")]
                            _ = v91
                        } else {
                            if (1 < 797) {
                                v92 := GenS12{F0: 248}
                                _ = v92
                            }
                        }
                    default:
                        v2 -= v2
                    }
                    v93 := v1[("" + ("hello" + (("hello" + "hello") + ("\n\t" + ""))))]
                    v93 = (!v93)
                    _ = v93
                } else {
                    var v94 int = (fi85 / ((2147483647) | 1))
                    fmt.Println(v3, fi85, v94)
                    v2 = fi85
                    _ = v94
                }
            }
        } else {
            v84 -= 1e10
            for fi95 := 0; fi95 < 14; fi95++ {
                var v96 uint16 = 0
                v3 = (-(v3))
                _ = fi95
                _ = v96
            }
        }
        _ = fi85
    }
    for fi97 := 0; fi97 < 18; fi97++ {
        fmt.Println(fi97, v4, v84)
        var v98 uint8 = 1
        _ = fi97
        _ = v98
    }
    for fi99 := 0; fi99 < 19; fi99++ {
        if false {
            v3 -= v3
            _ = len(v1)
        } else {
            if len(v8) > 0 {
                v8[0] = (fi99 << ((255) & 63))
            }
            switch (v2 % ((v2) | 1)) {
            case 32767:
                switch 9223372036854775807 {
                case 127:
                    for fi100 := 0; fi100 < 4; fi100++ {
                        fmt.Println(v2)
                        if (false || true) {
                            var v101 float64 = ((v84 * -2.25) - (v84 / ((-2.25 / -2.25) + float64((0.1 * (1.5 + 0.1))))))
                            _ = v101
                        }
                        if len(v8) > 0 {
                            v8[0] = (v2 & (fi99 * fi100))
                        }
                        _ = fi100
                    }
                }
            case -73:
                if true {
                    delete(v1, "foo bar")
                    fmt.Println(v4, fi99, fi99)
                    if false {
                        for fi102 := 0; fi102 < 10; fi102++ {
                            v84 = v84
                            var v103 uint8 = 0
                            _ = fi102
                            _ = v103
                        }
                    }
                }
            case -2147483648:
                if true {
                    if true {
                        if false {
                            var v104 uint = 18446744073709551615
                            _ = v104
                        } else {
                            v3 = 1
                            fmt.Println(v4)
                        }
                    } else {
                        if ("foo bar" > "foo bar") {
                            v1[("foo bar" + ("foo bar" + "\n\t"))] = (1 <= 167)
                            if (0 > 2) {
                                var v105 string = "\n\t"
                                _ = v105
                            }
                            v4 += v4
                        } else {
                            delete(v1, ("\n\t" + "hello"))
                            if (1e10 > 1.5) {
                                for fi106 := 0; fi106 < 5; fi106++ {
                                    switch "hello" {
                                    case "a":
                                        v84 += v84
                                    case "":
                                        switch "" {
                                        case "foo bar":
                                            switch 2147483647 {
                                            case -2147483648:
                                                fmt.Println(fi106)
                                            default:
                                                v2 -= fi106
                                            }
                                        case "hello":
                                            v84 += (v84 * v84)
                                        case "":
                                            if (18446744073709551615 != 65535) {
                                                v3 = (v3 | 0)
                                                _ = len(v8)
                                                var v107 int16 = -1
                                                v1["\n\t"] = false
                                                fmt.Println(v2)
                                                _ = v107
                                            } else {
                                                var v108 int16 = (v4 - (v4 * v4))
                                                v1[(("\n\t" + ("foo bar" + "a")) + ("" + ("foo bar" + "foo bar")))] = true
                                                fmt.Println(fi99, v3)
                                                fmt.Println(v3, v84, v84)
                                                v2 = (v2 + int(v108))
                                                if len(v8) > 0 {
                                                    v8[0] = (^(fi99))
                                                }
                                                var v109 uint32 = 2
                                                _ = v108
                                                _ = v109
                                            }
                                        }
                                    case "hello":
                                        v4 += v4
                                    default:
                                        switch "foo bar" {
                                        case "":
                                            v2 = v2
                                        default:
                                            var v110 uint = 869
                                            _ = v110
                                        }
                                    }
                                    var v111 int = fi99
                                    fmt.Println(v2, fi99, v3)
                                    _ = fi106
                                    _ = v111
                                }
                                delete(v1, (("a" + ((("a" + "hello") + "\n\t") + "a")) + (("" + ("\n\t" + "\n\t")) + "hello")))
                                v112 := GenS1{F0: 1, F1: 412, F2: false}
                                _ = v112
                            } else {
                                delete(v1, ("\n\t" + "foo bar"))
                            }
                        }
                    }
                    fmt.Println(fi99, v84)
                    v113 := GenS13{F0: 2}
                    _ = v113
                } else {
                    var v114 string = (("\n\t" + "\n\t") + ("foo bar" + "hello"))
                    v4 += v4
                    v4 = 32767
                    _ = v114
                }
            }
            switch (v2 >> ((520) & 63)) {
            case -1:
                var v115 int16 = (v4 ^ (v4 * v4))
                _ = v115
            default:
                switch "a" {
                case "":
                    v8 = append(v8, -9223372036854775808)
                case "foo bar":
                    v116 := GenS5{F0: 255, F1: (v4 ^ (-(v4)))}
                    _ = v116
                case "a":
                    if true {
                        v4 += v4
                        fmt.Println(v2)
                        var v117 string = ""
                        _ = v117
                    }
                }
            }
        }
        _ = fi99
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v84
}
