package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint16; F1 bool }
type GenS11 struct { F0 int; F1 bool; F2 int }
func genFn0(p0 uint32, p1 uint32, p2 uint16) int8 {
    var v5 uint16 = p2
    switch 32768 {
    case -2147483648:
        var v6 string = (("hello" + "") + (("hello" + (((("" + ("hello" + "\n\t")) + (((("foo bar" + ("" + "")) + "hello") + "a") + "hello")) + "hello") + "hello")) + (("foo bar" + "foo bar") + "hello")))
        _ = v6
    case 253:
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{(fi7 * 0), int(p0)}
            v5 = (v5 & p2)
            _ = fi7
            _ = v8
        }
    case 389:
        v9 := GenS1{F0: -35}
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        if false {
            var v11 uint64 = 464
            var v12 uint8 = 255
            _ = v11
            _ = v12
        } else {
            v13 := GenS1{F0: -9223372036854775808}
            if false {
                var v14 float32 = -2.25
                _ = v14
            }
            v15 := GenS1{F0: int64(p2)}
            _ = v13
            _ = v15
        }
        _ = fi10
    }
    v16 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v5
    return -118
}
func genFn1(p0 uint, p1 uint8, p2 uint64) uint32 {
    v17 := GenS2{F0: -9223372036854775808}
    for fi18 := 0; fi18 < 19; fi18++ {
        switch (fi18 * fi18) {
        case 127:
            if (!false) {
                switch "" {
                case "a":
                    var v19 int16 = 0
                    _ = v19
                case "hello":
                    var v20 uint32 = 2147483648
                    _ = v20
                case "":
                    p0 = p0
                default:
                    if false {
                        p2 -= (^(p2))
                        var v21 uint32 = 65535
                        _ = v21
                    } else {
                        p0 += 0
                        fmt.Println(p2, fi18, p1)
                        if (p2 != p2) {
                            var v22 uint32 = 2147483647
                            if (uint8(p0) >= (p1 - (p1 - p1))) {
                                var v23 uint64 = p2
                                var v24 int32 = 1
                                _ = v23
                                _ = v24
                            }
                            for fi25 := 0; fi25 < 13; fi25++ {
                                p1 = (p1 - p1)
                                _ = fi25
                            }
                            _ = v22
                        } else {
                            p2 = (p2 + p2)
                            fmt.Println(fi18)
                            p2 = (p2 + p2)
                        }
                    }
                }
                if false {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        if false {
                            v27 := GenS2{F0: 0}
                            var v28 string = "foo bar"
                            _ = v27
                            _ = v28
                        } else {
                            if true {
                                p0 -= p0
                            } else {
                                v29 := []int{(fi26 % (((fi26 % ((0) | 1))) | 1)), -762, 1}
                                if true {
                                    if len(v29) > 0 {
                                        v29[0] = fi26
                                    }
                                } else {
                                    v29 = append(v29, 550)
                                    var v30 int32 = int32(fi18)
                                    p0 -= p0
                                    _ = v30
                                }
                                _ = v29
                            }
                        }
                        v31 := []int{(fi18 << ((p0) & 63)), -362, (fi18 - -9223372036854775808), 0}
                        var v32 int32 = 42
                        var v33 uint32 = 2147483647
                        _ = fi26
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    var v34 uint64 = (p2 / (((p2 + 2)) | 1))
                    _ = v34
                }
            }
        case 2147483647:
            var v35 float64 = 1.5
            _ = v35
        default:
            v36 := GenS1{F0: 9223372036854775807}
            _ = v36
        }
        var v37 int = 318
        _ = fi18
        _ = v37
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 31
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float32 = (3.14159 / 3.14159)
    v3 := make(map[uint32]int)
    var v4 float32 = (v2 - v2)
    if false {
        var v38 uint = 255
        if (206 >= 32767) {
            var v39 string = "foo bar"
            if (true && (("foo bar" + v39) < v39)) {
                v4 = (v4 + (v4 + v2))
                fmt.Println(v4)
            } else {
                var v40 int8 = -128
                v38 = v38
                fmt.Println(v4)
                _ = v40
            }
            _ = v39
        } else {
            _ = len(v3)
            v41 := GenS3{F0: "hello", F1: 42}
            var v42 string = ("a" + "")
            _ = v41
            _ = v42
        }
        _ = v38
    } else {
        var v43 int16 = -509
        v44 := []int{128, 4294967295, 42}
        v43 = 0
        _ = v43
        _ = v44
    }
    var v45 bool = false
    if v45 {
        switch (("\n\t" + "hello") + "\n\t") {
        case "a":
            switch 2147483647 {
            case 295:
                for fi46 := 0; fi46 < 12; fi46++ {
                    var v47 int32 = -160
                    switch ("" + ("" + "hello")) {
                    case "foo bar":
                        var v48 string = "foo bar"
                        _ = v48
                    case "":
                        v49 := v3[(v1 % ((2) | 1))]
                        _ = v49
                    default:
                        fmt.Println(v47)
                    }
                    _ = fi46
                    _ = v47
                }
            default:
                fmt.Println(v2, v1)
            }
        case "":
            var v50 int64 = -518
            _ = v50
        }
    } else {
        for fi51 := 0; fi51 < 18; fi51++ {
            var v52 uint64 = 472
            if (595 != -32768) {
                _ = len(v3)
                v53 := GenS4{F0: 42}
                var v54 float32 = (-2.25 / (v4 / v2))
                _ = v53
                _ = v54
            } else {
                v55 := GenS2{F0: 42}
                for fi56 := 0; fi56 < 13; fi56++ {
                    v57 := []int{int(v1), 32767}
                    v1 = v1
                    _ = fi56
                    _ = v57
                }
                if v45 {
                    v52 = v52
                    var v58 string = (("foo bar" + "foo bar") + "a")
                    var v59 int = 42
                    _ = v58
                    _ = v59
                } else {
                    if (28 == -128) {
                        var v60 float32 = (v4 - v4)
                        _ = v60
                    } else {
                        v1 = genFn1(255, 42, v52)
                    }
                    for fi61 := 0; fi61 < 16; fi61++ {
                        _ = len(v3)
                        fmt.Println(v52, fi51)
                        if (440 > 0) {
                            if (v1 >= v1) {
                                if (false || v45) {
                                    fmt.Println(v2, v4)
                                    var v62 uint8 = 0
                                    _ = v62
                                }
                            } else {
                                var v63 int64 = 1
                                v64 := GenS5{F0: genFn0((v1 << ((42) & 63)), v1, 255), F1: v52}
                                _ = v63
                                _ = v64
                            }
                            v65 := []int{}
                            fmt.Println(fi51, v45)
                            _ = v65
                        }
                        var v66 int64 = 42
                        v52 = v52
                        _ = fi61
                        _ = v66
                    }
                    v67 := []int{fi51}
                    _ = v67
                }
                _ = v55
            }
            _ = fi51
            _ = v52
        }
        if false {
            var v68 uint32 = (^(v1))
            var v69 uint64 = 18446744073709551615
            _ = v68
            _ = v69
        }
    }
    switch "hello" {
    case "":
        v70 := []int{}
        _ = v70
    default:
        v71 := []int{790, -2147483648}
        _ = v71
    }
    for fi72 := 0; fi72 < 14; fi72++ {
        for fi73 := 0; fi73 < 18; fi73++ {
            switch (fi73 ^ -476) {
            case -1:
                if ("hello" != "") {
                    if (!(false || v45)) {
                        v1 = (v1 ^ 239)
                        switch fi73 {
                        case 0:
                            v4 += (((3.14159 * -2.25) + 0.0) - v4)
                        default:
                            switch ((("hello" + "") + "") + ("\n\t" + ("" + ""))) {
                            case "hello":
                                var v74 float64 = ((1e10 / -2.25) - 0.1)
                                _ = v74
                            }
                        }
                        var v75 int8 = -1
                        _ = v75
                    }
                    var v76 uint64 = 0
                    var v77 float64 = 3.14159
                    _ = v76
                    _ = v77
                }
            case 42:
                v2 += v2
            case 128:
                v2 = (v2 - v2)
            }
            _ = fi73
        }
        v78 := v3[v1]
        v45 = (3.14159 >= (((1e10 / ((0.0 / (1.5 / 1e10)) + 3.14159)) / 3.14159) - 0.0))
        _ = fi72
        _ = v78
    }
    switch "foo bar" {
    case "\n\t":
        if v45 {
            v4 += float32(-2.25)
            for fi79 := 0; fi79 < 4; fi79++ {
                fmt.Println(v4, v2)
                v1 -= v1
                var v80 int8 = -61
                _ = fi79
                _ = v80
            }
            v2 += 1e10
        }
    case "":
        for fi81 := 0; fi81 < 12; fi81++ {
            var v82 float64 = (0.1 / 1.5)
            if (false || ((!v45) || v45)) {
                var v83 int32 = 0
                v84 := GenS3{F0: "", F1: 0}
                v1 = v1
                _ = v83
                _ = v84
            } else {
                switch "" {
                case "foo bar":
                    if (!v45) {
                        switch "\n\t" {
                        case "":
                            fmt.Println(v2, v45, v1)
                        case "\n\t":
                            v85 := []int{(fi81 + 32768), fi81, fi81, (fi81 * (fi81 - (fi81 % ((fi81) | 1))))}
                            _ = v85
                        default:
                            for fi86 := 0; fi86 < 13; fi86++ {
                                var v87 float64 = v82
                                var v88 int64 = 42
                                _ = fi86
                                _ = v87
                                _ = v88
                            }
                        }
                    }
                case "hello":
                    var v89 string = "a"
                    _ = v89
                default:
                    delete(v3, 1)
                }
                var v90 float32 = v2
                _ = v90
            }
            _ = fi81
            _ = v82
        }
    case "a":
        switch "foo bar" {
        case "a":
            fmt.Println(v1, v45)
        case "hello":
            v91 := []int{32767}
            _ = v91
        }
    }
    if (0 < 2) {
        if (2147483647 < -2147483648) {
            v2 += (v2 + v4)
            var v92 string = ""
            switch "foo bar" {
            case "hello":
                fmt.Println(v92, v92, v2)
            }
            _ = v92
        } else {
            switch -2147483648 {
            case 42:
                v3[42] = 1
            }
        }
        if v45 {
            v1 = (v1 & (v1 >> ((65535) & 63)))
            v93 := GenS6{F0: v45}
            _ = v93
        }
    }
    if true {
        switch "hello" {
        case "foo bar":
            var v94 int32 = 0
            _ = v94
        case "hello":
            for fi95 := 0; fi95 < 16; fi95++ {
                v96 := GenS7{F0: 981, F1: 763}
                _ = fi95
                _ = v96
            }
        case "\n\t":
            var v97 string = "\n\t"
            _ = v97
        default:
            v1 = (v1 << ((255) & 63))
        }
        if false {
            switch (("hello" + ("a" + "a")) + (("" + "foo bar") + ("hello" + "a"))) {
            case "\n\t":
                v98 := GenS4{F0: 9223372036854775807}
                _ = v98
            case "":
                switch ("a" + "hello") {
                case "":
                    if false {
                        v45 = (("hello" + "") >= "hello")
                        v1 -= (v1 << ((4294967295) & 63))
                        var v99 uint32 = v1
                        _ = v99
                    } else {
                        var v100 int = 9223372036854775807
                        v101 := GenS8{F0: 32767, F1: (v1 | (v1 | (v1 + 0))), F2: 765}
                        for fi102 := 0; fi102 < 9; fi102++ {
                            if v45 {
                                for fi103 := 0; fi103 < 8; fi103++ {
                                    v45 = v45
                                    v45 = false
                                    _ = fi103
                                }
                                v104 := GenS9{F0: 32767}
                                _ = v104
                            }
                            v2 += (v2 + (v2 / -2.25))
                            switch "hello" {
                            case "":
                                v100 = (-(v100))
                            case "\n\t":
                                var v105 uint = 578
                                _ = v105
                            default:
                                fmt.Println(v45, v2, v2)
                            }
                            _ = fi102
                        }
                        _ = v100
                        _ = v101
                    }
                }
            case "foo bar":
                var v106 int8 = 100
                _ = v106
            default:
                switch "hello" {
                case "":
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        v1 = v1
                    case "\n\t":
                        fmt.Println(v4, v1)
                    case "hello":
                        v107 := GenS10{F0: 1, F1: (((1e10 - (0.0 + float32(0.0))) / float32(3.14159)) <= v4)}
                        _ = v107
                    default:
                        v108 := GenS10{F0: 42, F1: true}
                        _ = v108
                    }
                case "a":
                    var v109 int32 = 42
                    _ = v109
                case "foo bar":
                    for fi110 := 0; fi110 < 18; fi110++ {
                        v4 = v2
                        for fi111 := 0; fi111 < 19; fi111++ {
                            v112 := []int{(^(fi110))}
                            var v113 int = -1
                            v114 := GenS1{F0: 9223372036854775807}
                            _ = fi111
                            _ = v112
                            _ = v113
                            _ = v114
                        }
                        fmt.Println(v2, v1, v45)
                        _ = fi110
                    }
                default:
                    v115 := v3[0]
                    _ = v115
                }
            }
            v45 = (0 > 1)
        } else {
            v3[v1] = 127
        }
    }
    var v116 float32 = ((v4 - (v2 - (v4 * ((((-2.25 - 1e10) + float32((((0.0 / (((1e10 * ((1.5 - float64(v4)) / 0.1)) * (1e10 / (((((-2.25 + float64(float32((1e10 * 0.1)))) - 1e10) + ((((0.1 - float64(0.1)) / -2.25) + 0.1) - 3.14159)) - 1e10) + 1.5))) * -2.25)) + (float64(((((0.1 - ((float32(1e10) - (1.5 - (v4 / float32(float64(1e10))))) * ((float32(((3.14159 + (0.1 / (3.14159 - 0.1))) - 0.0)) * v4) + v4))) * float32(1e10)) / v4) * v4)) + 1e10)) - 1e10))) / ((float32(((-2.25 * (1e10 * 0.1)) / float64(3.14159))) - v2) / float32(3.14159))) + ((0.0 - v4) / 1.5))))) / v2)
    if v45 {
        var v117 int32 = 226
        _ = v117
    }
    if false {
        if true {
            v116 += v116
        }
        if v45 {
            var v118 int = 1
            _ = v118
        }
    }
    v119 := GenS11{F0: -9223372036854775808, F1: v45, F2: 1}
    v120 := v3[(^(v1))]
    var v121 uint8 = 106
    v121 -= (v121 * (v121 + v121))
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v116
    _ = v119
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v45
}
