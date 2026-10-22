package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 bool }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int16; F1 string; F2 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint; F1 uint32 }
type GenS11 struct { F0 string; F1 int32; F2 string }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint; F1 uint64 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int64; F1 bool; F2 uint8 }
type GenS16 struct { F0 uint16; F1 uint32; F2 string }
func genFn0(p0 int, p1 string) int {
    v5 := []int{127, (p0 | (p0 + p0))}
    switch p1 {
    case "a":
        p1 = ("" + p1)
    case "":
        for fi6 := 0; fi6 < 9; fi6++ {
            if len(v5) > 0 {
                v5[0] = p0
            }
            var v7 uint8 = 127
            v7 += 42
            _ = fi6
            _ = v7
        }
    case "hello":
        fmt.Println(p1)
    }
    if ((1e10 - 3.14159) != 1.5) {
        if true {
            var v8 uint8 = 2
            _ = v8
        } else {
            fmt.Println(p0, p1)
            if (false && false) {
                p1 = "a"
                p0 = 42
            } else {
                if len(v5) > 0 {
                    v5[0] = p0
                }
                var v9 uint64 = 127
                _ = v9
            }
        }
        switch ((p1 + p1) + p1) {
        case "foo bar":
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: 2147483647}
                if (351 >= -269) {
                    var v12 int8 = -128
                    var v13 int8 = (v12 / (((v12 | -128)) | 1))
                    switch p1 {
                    case "a":
                        v5 = append(v5, fi10)
                    case "":
                        var v14 uint16 = 32767
                        _ = v14
                    case "hello":
                        var v15 uint16 = 42
                        _ = v15
                    default:
                        var v16 uint16 = 0
                        _ = v16
                    }
                    _ = v12
                    _ = v13
                } else {
                    if true {
                        for fi17 := 0; fi17 < 10; fi17++ {
                            p1 = "foo bar"
                            _ = fi17
                        }
                        var v18 float32 = 0.1
                        for fi19 := 0; fi19 < 7; fi19++ {
                            var v20 uint = 18446744073709551615
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    } else {
                        if ((false && true) || (!(fi10 >= (p0 | p0)))) {
                            _ = len(v5)
                            if len(v5) > 0 {
                                v5[0] = p0
                            }
                            v21 := GenS2{F0: 110, F1: 2, F2: true}
                            _ = v21
                        }
                        p0 = -9223372036854775808
                        var v22 uint32 = 243
                        _ = v22
                    }
                    switch p0 {
                    case 32768:
                        var v23 int32 = -373
                        _ = v23
                    case 1:
                        p0 += fi10
                    default:
                        var v24 float64 = 3.14159
                        _ = v24
                    }
                    v25 := GenS3{F0: 4294967295, F1: p0}
                    _ = v25
                }
                if true {
                    var v26 float64 = (float64(0.0) / 1.0)
                    _ = len(v5)
                    fmt.Println(p0, v26)
                    _ = v26
                } else {
                    fmt.Println(p1, p0)
                    switch (p1 + "\n\t") {
                    case "":
                        fmt.Println(p1, p1)
                    case "a":
                        v27 := GenS4{F0: 756, F1: 1}
                        _ = v27
                    case "foo bar":
                        p1 = p1
                    }
                }
                _ = fi10
                _ = v11
            }
        default:
            v5 = append(v5, (p0 & p0))
        }
    }
    if false {
        switch ((p1 + ("foo bar" + p1)) + p1) {
        case "foo bar":
            var v28 uint = 2
            _ = v28
        default:
            for fi29 := 0; fi29 < 9; fi29++ {
                switch p0 {
                case -684:
                    if len(v5) > 0 {
                        v5[0] = (fi29 >> ((1) & 63))
                    }
                case 0:
                    p1 = ((((p1 + p1) + p1) + ("hello" + p1)) + p1)
                case 42:
                    switch "\n\t" {
                    case "foo bar":
                        for fi30 := 0; fi30 < 20; fi30++ {
                            var v31 bool = (-1 <= -2147483648)
                            switch (fi29 ^ (^(fi30))) {
                            case 288:
                                var v32 int8 = -88
                                _ = v32
                            case -9223372036854775808:
                                if true {
                                    v5 = append(v5, (fi29 + (-(fi29))))
                                } else {
                                    if ((!(1 != 902)) || (294 < 42)) {
                                        _ = len(v5)
                                        switch -348 {
                                        case 1:
                                            switch 32767 {
                                            case -290:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            default:
                                                var v33 string = ("a" + "hello")
                                                _ = v33
                                            }
                                        case 128:
                                            fmt.Println(fi29, p1, p0)
                                        case -2147483648:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            var v34 bool = v31
                                            _ = v34
                                        }
                                        p0 = p0
                                    }
                                    switch fi30 {
                                    case 32767:
                                        v35 := GenS4{F0: 9223372036854775807, F1: 451}
                                        _ = v35
                                    case 127:
                                        _ = len(v5)
                                    case 32768:
                                        fmt.Println(p0, p0, fi30)
                                    default:
                                        switch (fi30 >> ((578) & 63)) {
                                        case 1:
                                            switch (fi30 & 32768) {
                                            case -9223372036854775808:
                                                fmt.Println(fi30, v31, p0)
                                            case 2147483647:
                                                var v36 bool = ((float32(((0.1 / (1.5 * (-2.25 + 0.0))) / 1e10)) / -2.25) >= 0.1)
                                                _ = v36
                                            }
                                            p0 = (fi30 + fi29)
                                        default:
                                            v37 := GenS5{F0: 1000, F1: p1, F2: 775}
                                            fmt.Println(fi29, fi29, v31)
                                            if len(v5) > 0 {
                                                v5[0] = 128
                                            }
                                            _ = v37
                                        }
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = -2147483648
                                    }
                                }
                            case -590:
                                var v38 int32 = -1
                                _ = v38
                            default:
                                p0 += (p0 | (p0 | -1))
                            }
                            var v39 uint64 = 1
                            _ = fi30
                            _ = v31
                            _ = v39
                        }
                    case "a":
                        var v40 float32 = (1.5 * (3.14159 - 0.0))
                        _ = v40
                    case "":
                        v41 := GenS3{F0: 42, F1: p0}
                        _ = v41
                    }
                default:
                    for fi42 := 0; fi42 < 13; fi42++ {
                        p1 = p1
                        v43 := GenS6{F0: 4}
                        _ = fi42
                        _ = v43
                    }
                }
                if false {
                    switch "a" {
                    case "\n\t":
                        p1 = "foo bar"
                    case "a":
                        for fi44 := 0; fi44 < 20; fi44++ {
                            if true {
                                p0 += p0
                                if (!true) {
                                    if false {
                                        p0 = fi29
                                        var v45 int8 = -1
                                        var v46 string = (p1 + p1)
                                        _ = v45
                                        _ = v46
                                    } else {
                                        var v47 int32 = 42
                                        p0 += (fi29 + (p0 >> ((18446744073709551615) & 63)))
                                        switch -2147483648 {
                                        case -2147483648:
                                            switch p1 {
                                            case "hello":
                                                v5 = append(v5, 613)
                                            }
                                        case 32767:
                                            p0 -= fi44
                                        case 2147483647:
                                            var v48 bool = true
                                            _ = v48
                                        default:
                                            var v49 uint32 = 1
                                            _ = v49
                                        }
                                        _ = v47
                                    }
                                    v50 := GenS7{F0: -1}
                                    if (float64(float32(0.0)) == -2.25) {
                                        v51 := GenS7{F0: 1}
                                        fmt.Println(fi44)
                                        _ = v51
                                    } else {
                                        p0 -= fi44
                                    }
                                    v52 := GenS8{F0: 4294967295, F1: fi29}
                                    switch (p0 << ((569) & 63)) {
                                    case 32767:
                                        var v53 uint8 = 0
                                        _ = v53
                                    case 42:
                                        var v54 string = p1
                                        _ = v54
                                    }
                                    fmt.Println(p1, fi29)
                                    v55 := GenS9{F0: 255}
                                    _ = v50
                                    _ = v52
                                    _ = v55
                                }
                                p0 = fi44
                            }
                            var v56 uint64 = 87
                            var v57 int32 = -2147483648
                            _ = fi44
                            _ = v56
                            _ = v57
                        }
                    }
                }
                var v58 bool = false
                _ = fi29
                _ = v58
            }
        }
        if false {
            var v59 uint64 = 1
            v60 := GenS10{F0: 842, F1: 65535}
            for fi61 := 0; fi61 < 6; fi61++ {
                v62 := GenS11{F0: p1, F1: -2147483648, F2: p1}
                _ = fi61
                _ = v62
            }
            _ = v59
            _ = v60
        } else {
            if false {
                p1 = p1
                fmt.Println(p0, p0)
            }
            var v63 float32 = 1.5
            _ = v63
        }
    } else {
        p0 = p0
        fmt.Println(p0, p0)
        switch (p0 | p0) {
        case 9223372036854775807:
            if len(v5) > 0 {
                v5[0] = 454
            }
        case 313:
            var v64 string = p1
            _ = v64
        case 2147483647:
            var v65 int16 = 1000
            _ = v65
        default:
            v66 := GenS12{F0: (!false)}
            _ = v66
        }
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v5
    return (p0 / ((p0) | 1))
}

func main() {
    var v1 float32 = 3.14159
    v2 := []string{}
    v3 := make(map[string]int)
    var v4 int32 = -2147483648
    switch "a" {
    case "\n\t":
        v1 = v1
    default:
        v2 = append(v2, "a")
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    switch 32767 {
    case -2147483648:
        if (!(447 < 555)) {
            if true {
                if false {
                    for fi67 := 0; fi67 < 2; fi67++ {
                        if (42 == 2) {
                            v1 = v1
                            delete(v3, "foo bar")
                        } else {
                            var v68 float64 = 1e10
                            _ = v68
                        }
                        _ = fi67
                    }
                } else {
                    var v69 float64 = 1e10
                    _ = v69
                }
                if (!(1 <= 65535)) {
                    fmt.Println(v4, v4)
                } else {
                    v4 = v4
                    var v70 int16 = -32768
                    _ = v70
                }
            } else {
                if true {
                    switch ("" + "foo bar") {
                    case "hello":
                        v4 -= v4
                    case "":
                        var v71 uint32 = 2
                        _ = v71
                    case "a":
                        v4 += -687
                    default:
                        var v72 int64 = -9223372036854775808
                        _ = v72
                    }
                    for fi73 := 0; fi73 < 18; fi73++ {
                        v74 := GenS10{F0: 65535, F1: 0}
                        v75 := GenS13{F0: 397, F1: 668}
                        _ = fi73
                        _ = v74
                        _ = v75
                    }
                } else {
                    for fi76 := 0; fi76 < 15; fi76++ {
                        for fi77 := 0; fi77 < 20; fi77++ {
                            fmt.Println(v4)
                            v4 -= v4
                            var v78 float32 = v1
                            _ = fi77
                            _ = v78
                        }
                        fmt.Println(v4, v4, v1)
                        _ = fi76
                    }
                    var v79 uint16 = 498
                    delete(v3, "")
                    _ = v79
                }
            }
            var v80 uint64 = 731
            for fi81 := 0; fi81 < 9; fi81++ {
                fmt.Println(v1, v80, v80)
                for fi82 := 0; fi82 < 15; fi82++ {
                    v83 := GenS11{F0: "a", F1: (v4 ^ (v4 | 1)), F2: "a"}
                    var v84 int32 = (v4 - v4)
                    var v85 int32 = (^(v84))
                    _ = fi82
                    _ = v83
                    _ = v84
                    _ = v85
                }
                _ = fi81
            }
            _ = v80
        } else {
            fmt.Println(v1, v1)
        }
    default:
        var v86 float32 = v1
        _ = v86
    }
    if false {
        if (2 <= 351) {
            var v87 float32 = (float32(1e10) + float32((-2.25 * 1.5)))
            var v88 uint64 = 42
            _ = v87
            _ = v88
        } else {
            v89 := GenS7{F0: 42}
            _ = v89
        }
    }
    var v90 float32 = ((0.1 * ((0.1 + 1e10) * 1e10)) + float32((0.1 - 1.5)))
    for fi91 := 0; fi91 < 14; fi91++ {
        switch "hello" {
        case "foo bar":
            if (1 < 2) {
                v4 = v4
            } else {
                var v92 int64 = -109
                v1 += (v90 * float32(0.1))
                _ = v92
            }
        case "":
            v3[("\n\t" + "")] = genFn0(2147483647, "a")
        case "\n\t":
            for fi93 := 0; fi93 < 2; fi93++ {
                var v94 uint64 = 605
                if (true && false) {
                    if (0 < 42) {
                        v90 = v90
                        var v95 uint64 = v94
                        _ = v95
                    }
                    switch -666 {
                    case 479:
                        if false {
                            var v96 int8 = 127
                            delete(v3, "")
                            _ = v96
                        } else {
                            var v97 int32 = v4
                            _ = v97
                        }
                    default:
                        if true {
                            if (!false) {
                                fmt.Println(v90, v90)
                            } else {
                                v2 = append(v2, (("" + "a") + ""))
                            }
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v1, v90)
                            var v98 bool = true
                            _ = v98
                        }
                        var v99 string = ""
                        v94 = (v94 + v94)
                        v1 = (0.1 / float32(3.14159))
                        _ = v99
                    }
                    v4 = 0
                } else {
                    _ = len(v3)
                    var v100 uint = 255
                    _ = v100
                }
                if ((892 == 65535) || true) {
                    switch "hello" {
                    case "":
                        v94 = v94
                    case "hello":
                        switch "hello" {
                        case "":
                            v2 = append(v2, "\n\t")
                        }
                    }
                    if (true || false) {
                        v1 = 1e10
                        v94 = 18446744073709551615
                        var v101 uint64 = (v94 / ((v94) | 1))
                        _ = v101
                    }
                    if (!false) {
                        var v102 int64 = 1
                        v102 = v102
                        _ = v102
                    }
                } else {
                    v94 -= (v94 % ((v94) | 1))
                    if false {
                        if true {
                            var v103 int8 = 42
                            v1 = v90
                            fmt.Println(v4, v4)
                            _ = v103
                        } else {
                            v4 = 1
                        }
                        if (false && ((!true) || false)) {
                            fmt.Println(v90, v1)
                            _ = len(v2)
                            v94 = v94
                        }
                        v94 = v94
                    } else {
                        v104 := GenS5{F0: -841, F1: ((("hello" + "a") + ("a" + "\n\t")) + "\n\t"), F2: -109}
                        _ = v104
                    }
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                }
                _ = fi93
                _ = v94
            }
        default:
            v4 = (v4 & 2147483647)
        }
        switch "" {
        case "hello":
            v90 = (v90 + v90)
        case "\n\t":
            switch -2147483648 {
            case 9223372036854775807:
                v2 = append(v2, "a")
            case 289:
                if (((("a" + "") + ("foo bar" + "\n\t")) + "foo bar") <= ("a" + "")) {
                    if (!false) {
                        if true {
                            v105 := GenS14{F0: 18446744073709551615}
                            var v106 uint32 = 2147483648
                            _ = v105
                            _ = v106
                        }
                    } else {
                        var v107 uint8 = 130
                        for fi108 := 0; fi108 < 7; fi108++ {
                            v90 -= ((0.0 / (((3.14159 - (v1 * v90)) / v1) * float32((-2.25 / 3.14159)))) + (-2.25 * -2.25))
                            _ = fi108
                        }
                        _ = v107
                    }
                }
            default:
                var v109 string = ""
                _ = v109
            }
        case "":
            v90 = float32(0.1)
        default:
            v110 := GenS3{F0: 4294967295, F1: (fi91 - fi91)}
            _ = v110
        }
        _ = fi91
    }
    var v111 uint32 = 797
    var v112 float64 = ((3.14159 - 1.5) - float64(float32(1.5)))
    v111 = 692
    fmt.Println(v112, v112)
    delete(v3, "foo bar")
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1)
    _ = len(v3)
    v113 := GenS15{F0: 1, F1: true, F2: 2}
    if len(v2) > 0 {
        v2[0] = ("hello" + ("hello" + "\n\t"))
    }
    v114 := GenS16{F0: 0, F1: v111, F2: ("hello" + ("" + "a"))}
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v111
    _ = v112
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v90
}
