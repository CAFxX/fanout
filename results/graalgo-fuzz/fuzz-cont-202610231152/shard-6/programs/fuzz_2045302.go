package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 bool; F2 int16 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64 }
func genFn0() uint {
    for fi5 := 0; fi5 < 17; fi5++ {
        if (733 >= -676) {
            var v6 uint16 = 823
            _ = v6
        }
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: 1, F1: ("\n\t" + "\n\t"), F2: 27}
            var v9 int16 = 394
            fmt.Println(fi5, v9)
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS1{F0: -766, F1: (("" + "foo bar") + "a"), F2: -128}
            v12 := GenS1{F0: -32768, F1: "", F2: 42}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi5
    }
    var v13 int32 = 1
    v13 = 638
    _ = v13
    return 0
}
func genFn1(p0 uint, p1 uint32) int16 {
    if (false || (320 == -648)) {
        for fi14 := 0; fi14 < 14; fi14++ {
            switch (("hello" + "hello") + ("foo bar" + "")) {
            case "":
                for fi15 := 0; fi15 < 13; fi15++ {
                    p1 += (p1 ^ p1)
                    _ = fi15
                }
            case "hello":
                v16 := []int{(fi14 % (((fi14 >> ((1) & 63))) | 1)), fi14, 128, fi14}
                _ = v16
            case "foo bar":
                v17 := []int{fi14, fi14}
                _ = v17
            default:
                for fi18 := 0; fi18 < 19; fi18++ {
                    fmt.Println(p0, p0)
                    fmt.Println(p0)
                    _ = fi18
                }
            }
            if (-32768 > 1) {
                switch "a" {
                case "":
                    fmt.Println(p1)
                case "hello":
                    if ("foo bar" <= "foo bar") {
                        var v19 int8 = -1
                        _ = v19
                    }
                default:
                    var v20 uint64 = uint64(p0)
                    _ = v20
                }
                switch "" {
                case "hello":
                    v21 := []int{fi14}
                    _ = v21
                case "foo bar":
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v23 := []int{(fi14 * 42), (fi14 * 0), (fi14 >> ((genFn0()) & 63))}
                        p1 = 4294967295
                        _ = fi22
                        _ = v23
                    }
                default:
                    fmt.Println(p1, fi14, fi14)
                }
                switch "" {
                case "hello":
                    var v24 uint8 = 128
                    _ = v24
                case "foo bar":
                    switch fi14 {
                    case 128:
                        p1 += p1
                    case -1:
                        v25 := []int{int(p1), fi14}
                        _ = v25
                    case 9223372036854775807:
                        v26 := []int{9223372036854775807, fi14}
                        _ = v26
                    }
                case "a":
                    for fi27 := 0; fi27 < 12; fi27++ {
                        fmt.Println(fi14, p1, p1)
                        var v28 uint32 = p1
                        var v29 int = (fi27 * fi27)
                        switch 42 {
                        case 9223372036854775807:
                            p1 += (p1 / ((p1) | 1))
                        case -556:
                            v30 := []int{}
                            _ = v30
                        default:
                            var v31 uint8 = 1
                            _ = v31
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
            }
            _ = fi14
        }
        v32 := []int{-889, 127}
        _ = v32
    } else {
        p0 -= p0
    }
    fmt.Println(p0)
    var v33 float64 = 0.0
    var v34 float32 = 0.0
    _ = p0
    _ = p1
    _ = v33
    _ = v34
    return -762
}

func main() {
    var v1 uint64 = 869
    v2 := make(map[string]bool)
    var v3 bool = true
    var v4 uint8 = 102
    for fi35 := 0; fi35 < 9; fi35++ {
        for fi36 := 0; fi36 < 10; fi36++ {
            if (true && (-739 != -9223372036854775808)) {
                switch "" {
                case "foo bar":
                    switch "hello" {
                    case "":
                        delete(v2, "\n\t")
                    case "a":
                        if (42 <= genFn1(2, 968)) {
                            v37 := GenS2{F0: 0, F1: -128, F2: 127}
                            v38 := GenS2{F0: 255, F1: -25, F2: 1}
                            _ = v37
                            _ = v38
                        }
                    case "foo bar":
                        var v39 int8 = 100
                        _ = v39
                    }
                    switch -967 {
                    case 2147483647:
                        var v40 float32 = 1.5
                        _ = v40
                    case -2147483648:
                        v41 := GenS2{F0: genFn0(), F1: -128, F2: 100}
                        _ = v41
                    case 719:
                        if ((!false) || v3) {
                            var v42 int = 128
                            _ = v42
                        } else {
                            v43 := GenS3{F0: 72}
                            v44 := GenS4{F0: (v4 ^ v4)}
                            _ = v43
                            _ = v44
                        }
                    }
                default:
                    var v45 uint32 = 2147483648
                    _ = v45
                }
            }
            v1 -= v1
            fmt.Println(fi36, fi36)
            _ = fi36
        }
        _ = fi35
    }
    for fi46 := 0; fi46 < 13; fi46++ {
        var v47 int8 = 0
        _ = fi46
        _ = v47
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v4)
    case "hello":
        v2[("hello" + "\n\t")] = ((755 <= 922) && false)
    default:
        if ((!v3) || ((1.5 > 1e10) || true)) {
            if true {
                switch 4294967295 {
                case 32767:
                    if (1 == 0) {
                        if (160 > (v4 ^ 2)) {
                            if false {
                                var v48 bool = (-9223372036854775808 >= -1)
                                _ = v48
                            } else {
                                for fi49 := 0; fi49 < 8; fi49++ {
                                    for fi50 := 0; fi50 < 3; fi50++ {
                                        v1 = 926
                                        v51 := GenS2{F0: 392, F1: -128, F2: 8}
                                        _ = len(v2)
                                        v52 := GenS1{F0: 1, F1: "hello", F2: -1}
                                        _ = fi50
                                        _ = v51
                                        _ = v52
                                    }
                                    _ = fi49
                                }
                                v3 = ("hello" > "foo bar")
                                if true {
                                    if v3 {
                                        var v53 uint16 = 35
                                        var v54 uint32 = 862
                                        var v55 float64 = (float64(1e10) + 1.5)
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                    if true {
                                        v1 = 0
                                        switch "hello" {
                                        case "hello":
                                            v2["a"] = (!(("foo bar" + "a") == ("a" + ("hello" + "a"))))
                                        case "":
                                            if ((!v3) || false) {
                                                fmt.Println(v3, v1, v1)
                                            } else {
                                                v1 += 42
                                                fmt.Println(v1, v3)
                                                v56 := v2[("" + "\n\t")]
                                                _ = v56
                                            }
                                        }
                                        v57 := GenS5{F0: ((0.0 / 3.14159) > -2.25), F1: (!(!(v3 && (127 > (^(v4)))))), F2: -1}
                                        _ = v57
                                    } else {
                                        var v58 uint8 = v4
                                        _ = v58
                                    }
                                    for fi59 := 0; fi59 < 9; fi59++ {
                                        var v60 uint8 = v4
                                        _ = fi59
                                        _ = v60
                                    }
                                } else {
                                    if false {
                                        var v61 int32 = -2147483648
                                        _ = v61
                                    }
                                }
                            }
                            var v62 uint8 = 42
                            switch ("hello" + (("a" + "hello") + ((((("a" + (("\n\t" + "\n\t") + "hello")) + "\n\t") + "") + ((("" + ("" + "a")) + "") + "\n\t")) + "foo bar"))) {
                            case "a":
                                if v3 {
                                    for fi63 := 0; fi63 < 10; fi63++ {
                                        for fi64 := 0; fi64 < 18; fi64++ {
                                            if (false && v3) {
                                                v2[(("\n\t" + ("" + ("foo bar" + ""))) + "hello")] = (-2147483648 == 1)
                                                if (1000 < 1000) {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if true {
                                                            v65 := []int{}
                                                            _ = v65
                                                        } else {
                                                            v66 := v2["foo bar"]
                                                            v1 = (^(v1))
                                                            v67 := GenS6{F0: 42, F1: v62}
                                                            if true {
                                                                v68 := GenS4{F0: v62}
                                                                _ = v68
                                                            } else {
                                                                fmt.Println(fi64)
                                                            }
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                    }
                                                    switch "hello" {
                                                    case "hello":
                                                        var v69 float32 = (-2.25 - 1e10)
                                                        switch "a" {
                                                        case "a":
                                                            var v70 int8 = 42
                                                            _ = v70
                                                        }
                                                        _ = v69
                                                    case "\n\t":
                                                        var v71 uint16 = 32767
                                                        _ = v71
                                                    case "a":
                                                        if v3 {
                                                            var v72 float64 = 1.5
                                                            if false {
                                                                v62 = v4
                                                                var v73 uint16 = 2
                                                                fmt.Println(v73, fi63, v62)
                                                                _ = v73
                                                            }
                                                            fmt.Println(v4, v3, v72)
                                                            _ = v72
                                                        } else {
                                                            v3 = ((0.0 - float64(0.0)) < 3.14159)
                                                            v3 = v3
                                                            v1 = (^(v1))
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v1, fi63, v62)
                                                    v3 = (v3 || true)
                                                }
                                            } else {
                                                v74 := GenS7{F0: 42, F1: 952}
                                                fmt.Println(v4, v4, v1)
                                                _ = v74
                                            }
                                            if v3 {
                                                v75 := GenS6{F0: 42, F1: (v4 - v62)}
                                                v62 = (v62 & v4)
                                                _ = v75
                                            }
                                            _ = fi64
                                        }
                                        if v3 {
                                            for fi76 := 0; fi76 < 16; fi76++ {
                                                v77 := GenS4{F0: v62}
                                                fmt.Println(v1)
                                                var v78 bool = v3
                                                _ = fi76
                                                _ = v77
                                                _ = v78
                                            }
                                            var v79 uint = 422
                                            _ = v79
                                        } else {
                                            for fi80 := 0; fi80 < 6; fi80++ {
                                                v62 = v62
                                                if (genFn1(0, 4294967295) <= 1000) {
                                                    fmt.Println(fi63, v3, fi80)
                                                    if v3 {
                                                        v81 := GenS8{F0: "foo bar"}
                                                        var v82 float64 = 0.1
                                                        switch 208 {
                                                        case 2147483647:
                                                            v62 -= 179
                                                        default:
                                                            v4 -= v4
                                                        }
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                } else {
                                                    if (false || ((fi63 / ((fi63) | 1)) < 2147483647)) {
                                                        v83 := GenS9{F0: 989}
                                                        v3 = v3
                                                        if v3 {
                                                            var v84 int32 = -908
                                                            v85 := []int{fi63, (fi63 << ((77) & 63)), fi80}
                                                            if v3 {
                                                                var v86 float64 = (((3.14159 + ((1.5 + 1e10) + (0.1 + 3.14159))) * 0.0) + 1e10)
                                                                v62 = (^(v62))
                                                                var v87 int = 42
                                                                _ = v86
                                                                _ = v87
                                                            } else {
                                                                var v88 int8 = 1
                                                                fmt.Println(v88)
                                                                var v89 int16 = 0
                                                                var v90 float32 = (3.14159 - 0.1)
                                                                _ = v88
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                            if false {
                                                                var v91 uint8 = v4
                                                                fmt.Println(v3, v4, v3)
                                                                _ = v91
                                                            }
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            v4 = 128
                                                            if v3 {
                                                                if (!(2 > 2147483648)) {
                                                                    v1 = (v1 << ((genFn0()) & 63))
                                                                }
                                                                fmt.Println(fi80, v62)
                                                                var v92 float64 = 1e10
                                                                _ = v92
                                                            }
                                                            fmt.Println(fi80)
                                                            switch (fi80 >> ((genFn0()) & 63)) {
                                                            case 128:
                                                                fmt.Println(v62, v3, v62)
                                                            case 2147483647:
                                                                v1 -= (v1 << ((255) & 63))
                                                                var v93 float32 = 1e10
                                                                _ = v93
                                                            default:
                                                                if ((float64(3.14159) < -2.25) || true) {
                                                                    v62 = 128
                                                                    v94 := []int{fi80, fi63, (fi80 & 592), -2147483648}
                                                                    if ((!(1 < 42)) && v3) {
                                                                        var v95 int = (fi63 ^ fi63)
                                                                        _ = v95
                                                                    }
                                                                    _ = v94
                                                                }
                                                            }
                                                            v96 := []int{fi63, (fi63 % (((-(fi80))) | 1)), fi80}
                                                            _ = v96
                                                        }
                                                        _ = v83
                                                    }
                                                    v97 := GenS2{F0: genFn0(), F1: -65, F2: -1}
                                                    _ = v97
                                                }
                                                v4 -= v4
                                                _ = fi80
                                            }
                                            fmt.Println(fi63, v1)
                                        }
                                        v98 := GenS5{F0: (false && v3), F1: (v3 || ((100 >= v62) || (906 == 65535))), F2: int16(fi63)}
                                        _ = fi63
                                        _ = v98
                                    }
                                    var v99 uint8 = v4
                                    _ = v99
                                } else {
                                    var v100 int8 = 0
                                    _ = v100
                                }
                            case "\n\t":
                                var v101 uint = 1
                                _ = v101
                            }
                            _ = v62
                        } else {
                            for fi102 := 0; fi102 < 5; fi102++ {
                                v103 := GenS2{F0: genFn0(), F1: -128, F2: 100}
                                v3 = v3
                                var v104 int = -256
                                _ = fi102
                                _ = v103
                                _ = v104
                            }
                        }
                        for fi105 := 0; fi105 < 10; fi105++ {
                            for fi106 := 0; fi106 < 6; fi106++ {
                                v4 = (v4 ^ v4)
                                if (0 > 0) {
                                    v107 := []int{fi106, (^(fi106))}
                                    fmt.Println(v1, v3, fi105)
                                    _ = v107
                                } else {
                                    var v108 float32 = 1.5
                                    _ = v108
                                }
                                _ = fi106
                            }
                            _ = fi105
                        }
                    } else {
                        var v109 bool = (42 == -312)
                        _ = v109
                    }
                default:
                    v110 := GenS6{F0: 9223372036854775807, F1: 8}
                    _ = v110
                }
            }
            var v111 float64 = 0.1
            _ = v111
        }
    }
    if (42 < 1) {
        v112 := v2["hello"]
        if false {
            var v113 float64 = 0.0
            _ = v113
        }
        _ = v112
    }
    v1 -= (^(v1))
    var v114 uint = 4294967295
    fmt.Println(v4)
    v1 += v1
    v4 = (v4 + v4)
    switch ("a" + "foo bar") {
    case "\n\t":
        for fi115 := 0; fi115 < 7; fi115++ {
            v2[""] = true
            if ((!v3) && (-2.25 <= (1e10 / float32(3.14159)))) {
                v116 := GenS8{F0: ("hello" + "a")}
                v114 += 424
                _ = v116
            }
            _ = fi115
        }
    case "hello":
        var v117 float64 = (((0.1 - 3.14159) - 1e10) + -2.25)
        _ = v117
    case "":
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v114
    _ = v2
    _ = v3
    _ = v4
}
