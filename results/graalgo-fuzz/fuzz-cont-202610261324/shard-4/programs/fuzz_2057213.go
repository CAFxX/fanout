package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint }
type GenS5 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int16; F1 int64; F2 uint8 }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 uint8; F1 int8 }
func genFn0(p0 int16) string {
    fmt.Println(p0)
    if true {
        var v5 bool = (-979 != -2147483648)
        if (3.14159 == 1e10) {
            var v6 string = "foo bar"
            _ = v6
        } else {
            v7 := GenS1{F0: 0, F1: 32767, F2: 0}
            var v8 float64 = 1.5
            _ = v7
            _ = v8
        }
        _ = v5
    } else {
        switch 42 {
        case 127:
            var v9 int32 = 42
            _ = v9
        case 0:
            var v10 uint = 0
            _ = v10
        }
        if ("" > ((("foo bar" + "hello") + "a") + "\n\t")) {
            p0 -= 1
            var v11 bool = false
            for fi12 := 0; fi12 < 5; fi12++ {
                if v11 {
                    v13 := GenS2{F0: 1}
                    for fi14 := 0; fi14 < 18; fi14++ {
                        var v15 int16 = 1000
                        v16 := GenS3{F0: 840, F1: 42}
                        fmt.Println(p0, fi14, v11)
                        v17 := GenS4{F0: 4294967295, F1: 762, F2: 65535}
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = v13
                } else {
                    v18 := []int{}
                    if (!v11) {
                        var v19 bool = v11
                        p0 -= p0
                        _ = v19
                    } else {
                        var v20 uint8 = 231
                        _ = len(v18)
                        var v21 string = "foo bar"
                        _ = v20
                        _ = v21
                    }
                    if (65535 < 42) {
                        p0 = -1
                        for fi22 := 0; fi22 < 16; fi22++ {
                            switch ("foo bar" + "\n\t") {
                            case "a":
                                if true {
                                    v11 = v11
                                    v23 := GenS4{F0: 361, F1: -1, F2: 0}
                                    _ = v23
                                } else {
                                    p0 += (p0 / (((p0 | (p0 | 1000))) | 1))
                                    fmt.Println(fi22, v11)
                                    fmt.Println(v11, v11)
                                }
                            case "":
                                var v24 bool = (0 == 4294967295)
                                _ = v24
                            case "foo bar":
                                v11 = ((201 != 42) || v11)
                            }
                            var v25 int16 = 1
                            fmt.Println(fi22, v11, p0)
                            _ = fi22
                            _ = v25
                        }
                    }
                    _ = v18
                }
                if v11 {
                    v26 := GenS3{F0: 0, F1: 0}
                    var v27 uint8 = 54
                    if (0.1 > float32(0.0)) {
                        fmt.Println(v11, v27, fi12)
                    } else {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            v11 = true
                            var v29 uint = 1
                            switch fi12 {
                            case 512:
                                var v30 int16 = (-(p0))
                                _ = v30
                            case -9223372036854775808:
                                var v31 int8 = -128
                                _ = v31
                            case 4294967295:
                                var v32 int64 = 1
                                _ = v32
                            default:
                                if v11 {
                                    v33 := []int{(fi12 ^ fi12)}
                                    v34 := GenS5{F0: 1, F1: 65535, F2: 207}
                                    _ = len(v33)
                                    _ = v33
                                    _ = v34
                                }
                            }
                            _ = fi28
                            _ = v29
                        }
                        fmt.Println(p0, v27)
                        var v35 uint8 = (v27 ^ v27)
                        _ = v35
                    }
                    _ = v26
                    _ = v27
                } else {
                    p0 += (^(p0))
                }
                v36 := GenS1{F0: 568, F1: 9223372036854775807, F2: 42}
                _ = fi12
                _ = v36
            }
            _ = v11
        }
        p0 = (p0 % (((p0 / ((42) | 1))) | 1))
    }
    v37 := []int{128, 1, -2147483648, 128}
    _ = p0
    _ = v37
    return "foo bar"
}

func main() {
    var v1 uint64 = 432
    v2 := []string{}
    var v3 int32 = -1
    v4 := make(map[int]int)
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 uint16 = 0
        var v40 float64 = (0.0 / (-2.25 / 1.5))
        _ = fi38
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        v3 -= (v3 + (-(v3)))
        if (1 < 2) {
            for fi42 := 0; fi42 < 2; fi42++ {
                if (false || false) {
                    var v43 int32 = (v3 - (^(v3)))
                    _ = v43
                }
                v4[(fi42 ^ (-(fi41)))] = (fi42 * -960)
                _ = fi42
            }
            var v44 uint8 = 128
            if false {
                v45 := GenS4{F0: (v1 << ((4294967295) & 63)), F1: -9223372036854775808, F2: 18446744073709551615}
                var v46 int8 = 100
                for fi47 := 0; fi47 < 19; fi47++ {
                    var v48 int8 = (v46 + -107)
                    var v49 int8 = v48
                    switch (("hello" + "a") + "foo bar") {
                    case "foo bar":
                        if (!(0 != 1)) {
                            v49 -= v46
                            v50 := GenS1{F0: 65535, F1: fi41, F2: 2147483648}
                            switch (fi47 + fi47) {
                            case -1:
                                v51 := GenS1{F0: 42, F1: fi41, F2: 844}
                                _ = v51
                            case 473:
                                if true {
                                    v44 = uint8(v1)
                                    v48 -= int8(fi41)
                                    fmt.Println(v46, v46, fi47)
                                } else {
                                    fmt.Println(fi41)
                                    v3 += (^(v3))
                                    _ = len(v4)
                                }
                            case 0:
                                var v52 int = int(v44)
                                _ = v52
                            }
                            _ = v50
                        }
                    case "hello":
                        v53 := GenS5{F0: (v3 & (v3 & 1)), F1: 967, F2: 1}
                        _ = v53
                    case "a":
                        fmt.Println(v49, v48)
                    }
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                _ = v45
                _ = v46
            }
            _ = v44
        }
        switch "\n\t" {
        case "":
            if true {
                _ = len(v2)
            } else {
                v1 = (v1 & (v1 - v1))
            }
        case "hello":
            v54 := GenS6{F0: 1000, F1: 21, F2: 0}
            _ = v54
        default:
            for fi55 := 0; fi55 < 12; fi55++ {
                switch ("\n\t" + "a") {
                case "foo bar":
                    var v56 uint32 = 4294967295
                    _ = v56
                case "a":
                    var v57 uint16 = 2
                    _ = v57
                case "":
                    if (((66 >= 42) || (0.0 >= 0.1)) || (!(!(false && (((2 > 18446744073709551615) && false) || true))))) {
                        var v58 int16 = 272
                        _ = v58
                    }
                    delete(v4, fi55)
                default:
                    v59 := v4[fi41]
                    _ = v59
                }
                v3 -= 0
                var v60 uint16 = 1
                _ = fi55
                _ = v60
            }
        }
        _ = fi41
    }
    var v61 int8 = -95
    switch 0 {
    case 0:
        v3 -= 1
    case -9223372036854775808:
        var v62 string = ""
        _ = v62
    case 42:
        if false {
            fmt.Println(v1, v61)
            for fi63 := 0; fi63 < 2; fi63++ {
                var v64 uint32 = 2
                v65 := GenS5{F0: (v3 ^ (v3 % (((v3 - (v3 & (v3 + (v3 | v3))))) | 1))), F1: 1, F2: v64}
                for fi66 := 0; fi66 < 9; fi66++ {
                    if false {
                        var v67 int = 65
                        v1 = (v1 << ((410) & 63))
                        switch "" {
                        case "foo bar":
                            v64 = v64
                        case "a":
                            var v68 uint = 255
                            _ = v68
                        }
                        _ = v67
                    }
                    _ = fi66
                }
                _ = fi63
                _ = v64
                _ = v65
            }
        } else {
            var v69 int8 = v61
            for fi70 := 0; fi70 < 19; fi70++ {
                var v71 int8 = 24
                if (v1 == (v1 - 491)) {
                    v71 = -110
                    var v72 uint16 = 72
                    _ = v72
                }
                switch fi70 {
                case 587:
                    v73 := GenS5{F0: v3, F1: 42, F2: 0}
                    _ = v73
                case 32768:
                    if ((^(v1)) <= (v1 | (v1 * v1))) {
                        fmt.Println(fi70, v71, v1)
                        for fi74 := 0; fi74 < 14; fi74++ {
                            fmt.Println(v69, v71)
                            var v75 bool = false
                            fmt.Println(v1)
                            v76 := GenS5{F0: (-(v3)), F1: 1, F2: 2}
                            _ = fi74
                            _ = v75
                            _ = v76
                        }
                    }
                case 128:
                    var v77 uint16 = 42
                    _ = v77
                default:
                    switch "" {
                    case "":
                        var v78 uint8 = 0
                        _ = v78
                    case "a":
                        v79 := GenS3{F0: 1, F1: 303}
                        _ = v79
                    case "foo bar":
                        if false {
                            switch -49 {
                            case -603:
                                var v80 string = ("\n\t" + "")
                                _ = v80
                            case 2147483647:
                                switch (fi70 ^ (fi70 ^ (fi70 * fi70))) {
                                case -9223372036854775808:
                                    delete(v4, 684)
                                case 128:
                                    v2 = append(v2, ("a" + "foo bar"))
                                }
                            }
                        } else {
                            if false {
                                if (!true) {
                                    if (false && false) {
                                        v3 += 42
                                    }
                                    fmt.Println(v61, fi70)
                                } else {
                                    for fi81 := 0; fi81 < 16; fi81++ {
                                        var v82 uint32 = 1
                                        var v83 uint64 = (v1 / (((v1 & v1)) | 1))
                                        var v84 bool = false
                                        _ = fi81
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        for fi85 := 0; fi85 < 6; fi85++ {
                                            v1 = (v1 & v1)
                                            if true {
                                                v1 = v1
                                                if ((false && true) && true) {
                                                    v4[fi85] = -1
                                                    if (255 >= uint16(v1)) {
                                                        v86 := GenS1{F0: 2, F1: fi85, F2: 0}
                                                        if len(v2) > 0 {
                                                            v2[0] = ""
                                                        }
                                                        _ = v86
                                                    }
                                                } else {
                                                    v87 := GenS5{F0: v3, F1: 0, F2: 2}
                                                    switch "a" {
                                                    case "a":
                                                        delete(v4, (fi85 << ((520) & 63)))
                                                    case "\n\t":
                                                        var v88 float64 = (1e10 * 1.5)
                                                        _ = v88
                                                    case "foo bar":
                                                        fmt.Println(v71)
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = ""
                                                    }
                                                    _ = v87
                                                }
                                                var v89 uint32 = 238
                                                _ = v89
                                            } else {
                                                v3 = -48
                                                v90 := v4[(-(fi70))]
                                                if false {
                                                    var v91 int64 = 9223372036854775807
                                                    if (!false) {
                                                        delete(v4, 32768)
                                                    } else {
                                                        if false {
                                                            var v92 int64 = v91
                                                            _ = v92
                                                        } else {
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v93 float64 = 3.14159
                                                                _ = v93
                                                            case "foo bar":
                                                                v3 = (v3 + v3)
                                                            }
                                                            var v94 int32 = (-(v3))
                                                            if false {
                                                                var v95 float64 = 1e10
                                                                var v96 float32 = (float32((1.5 * -2.25)) / -2.25)
                                                                _ = v95
                                                                _ = v96
                                                            }
                                                            _ = v94
                                                        }
                                                        _ = len(v2)
                                                        v69 -= (^(v69))
                                                        fmt.Println(v91, fi70, v3)
                                                    }
                                                    _ = v91
                                                } else {
                                                    var v97 int16 = -848
                                                    _ = v97
                                                }
                                                _ = v90
                                            }
                                            _ = fi85
                                        }
                                    case "foo bar":
                                        v98 := GenS4{F0: v1, F1: 9223372036854775807, F2: 0}
                                        _ = v98
                                    }
                                }
                            }
                            if (1.5 > (3.14159 * (0.1 - (3.14159 - -2.25)))) {
                                v61 = (v71 << ((18446744073709551615) & 63))
                                var v99 int32 = (v3 ^ v3)
                                _ = v99
                            } else {
                                for fi100 := 0; fi100 < 1; fi100++ {
                                    if (!(!true)) {
                                        var v101 uint16 = 255
                                        _ = v101
                                    }
                                    _ = fi100
                                }
                                v3 -= v3
                                for fi102 := 0; fi102 < 5; fi102++ {
                                    v103 := GenS7{F0: 18446744073709551615, F1: -201}
                                    _ = fi102
                                    _ = v103
                                }
                            }
                        }
                    }
                }
                _ = fi70
                _ = v71
            }
            _ = v69
        }
    }
    var v104 uint = 4294967295
    if false {
        for fi105 := 0; fi105 < 16; fi105++ {
            if (true && (!(1 == 255))) {
                v61 += v61
            } else {
                var v106 uint64 = (v1 ^ v1)
                var v107 uint32 = 2147483648
                _ = v106
                _ = v107
            }
            v61 = v61
            v108 := GenS8{F0: 2, F1: 1}
            _ = fi105
            _ = v108
        }
        v109 := GenS9{F0: 976, F1: true}
        v61 = v61
        _ = v109
    }
    fmt.Println(v1, v1, v61)
    var v110 uint8 = 127
    v111 := GenS10{F0: 202, F1: (v61 * -53)}
    switch 9223372036854775807 {
    case 127:
        v112 := v4[9223372036854775807]
        _ = v112
    default:
        var v113 int = -566
        _ = v113
    }
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v110:", v110)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v104
    _ = v110
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v61
}
