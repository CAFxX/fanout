package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int32; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint }
type GenS5 struct { F0 int16; F1 int; F2 uint }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 int8; F2 bool }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32; F1 uint64 }
type GenS12 struct { F0 bool; F1 int; F2 int }
type GenS13 struct { F0 int8; F1 uint64 }
func genFn0(p0 int16, p1 string, p2 uint32) uint {
    p1 = ""
    switch 0 {
    case 32767:
        var v7 float64 = 1.5
        _ = v7
    }
    switch 4294967295 {
    case 614:
        fmt.Println(p1, p0, p0)
    case 32768:
        if (!(!(-1 == 9223372036854775807))) {
            p1 = ((p1 + p1) + (p1 + p1))
        }
    }
    var v8 int8 = 1
    if (2147483647 == -2147483648) {
        switch 623 {
        case 32768:
            for fi9 := 0; fi9 < 16; fi9++ {
                p0 = -32768
                if true {
                    var v10 int16 = 42
                    for fi11 := 0; fi11 < 10; fi11++ {
                        var v12 uint64 = 4294967295
                        v13 := []int{}
                        if (-1 < int64(v12)) {
                            fmt.Println(v10)
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                } else {
                    var v14 float32 = (float32(1.5) / 1.0)
                    var v15 bool = (true || true)
                    _ = v14
                    _ = v15
                }
                _ = fi9
            }
        default:
            for fi16 := 0; fi16 < 1; fi16++ {
                for fi17 := 0; fi17 < 11; fi17++ {
                    if (!((3.14159 * 3.14159) < 3.14159)) {
                        fmt.Println(v8, p2)
                    } else {
                        switch "hello" {
                        case "foo bar":
                            v18 := GenS1{F0: int16(fi17)}
                            _ = v18
                        case "hello":
                            v8 -= 1
                        }
                    }
                    var v19 uint64 = 0
                    p1 = (p1 + "a")
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(p1)
                    default:
                        p2 = p2
                    }
                    _ = fi17
                    _ = v19
                }
                _ = fi16
            }
        }
        v20 := []int{}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 17; fi21++ {
            v8 += (-(v8))
            switch (p1 + ((p1 + (((p1 + p1) + (p1 + p1)) + p1)) + "foo bar")) {
            case "hello":
                var v22 uint16 = 491
                _ = v22
            case "foo bar":
                if false {
                    var v23 uint16 = 1
                    _ = v23
                }
            case "":
                switch p1 {
                case "a":
                    var v24 uint16 = 42
                    _ = v24
                case "\n\t":
                    fmt.Println(p2)
                case "foo bar":
                    p2 = p2
                default:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        if (("hello" >= "foo bar") || (true && (1 <= 255))) {
                            var v26 int = -1
                            v8 += 37
                            var v27 int32 = 1
                            _ = v26
                            _ = v27
                        }
                        var v28 uint64 = 557
                        v29 := []int{0, (fi21 % ((fi21) | 1)), fi25}
                        switch fi25 {
                        case 812:
                            p2 -= p2
                        case 0:
                            if (-2.25 < 3.14159) {
                                var v30 string = (p1 + "\n\t")
                                _ = v30
                            } else {
                                p2 -= p2
                                v31 := GenS1{F0: int16(v28)}
                                var v32 float32 = 3.14159
                                _ = v31
                                _ = v32
                            }
                        default:
                            fmt.Println(fi21)
                        }
                        switch p1 {
                        case "hello":
                            var v33 uint8 = 255
                            _ = v33
                        case "a":
                            p0 = p0
                        case "foo bar":
                            var v34 uint16 = 1
                            _ = v34
                        }
                        _ = fi25
                        _ = v28
                        _ = v29
                    }
                }
            default:
                v8 -= (v8 + 0)
            }
            _ = fi21
        }
        for fi35 := 0; fi35 < 10; fi35++ {
            p0 = p0
            if false {
                var v36 uint = 2
                _ = v36
            } else {
                fmt.Println(fi35)
            }
            _ = fi35
        }
        v37 := []int{}
        _ = v37
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 255
}
func genFn1(p0 int64) uint64 {
    if (65535 <= 65535) {
        var v38 string = "foo bar"
        var v39 uint32 = 332
        _ = v38
        _ = v39
    } else {
        fmt.Println(p0)
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        for fi41 := 0; fi41 < 11; fi41++ {
            v42 := GenS2{F0: 29, F1: 1, F2: 78}
            v43 := GenS3{F0: 1, F1: 65535}
            _ = fi41
            _ = v42
            _ = v43
        }
        var v44 int64 = -713
        fmt.Println(v44)
        _ = fi40
        _ = v44
    }
    _ = p0
    return 585
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 42
    v4 := make(map[string]int)
    var v5 uint32 = (v3 & 65535)
    var v6 uint64 = 180
    for fi45 := 0; fi45 < 6; fi45++ {
        var v46 int64 = 724
        if (uint(v1) == genFn0(-139, ("\n\t" + ""), v5)) {
            if (!(1000 != -32768)) {
                v47 := GenS2{F0: 318, F1: -2147483648, F2: 127}
                switch (fi45 + (fi45 ^ (fi45 | fi45))) {
                case -1:
                    if (fi45 > (fi45 >> ((760) & 63))) {
                        for fi48 := 0; fi48 < 9; fi48++ {
                            var v49 uint8 = 255
                            v49 = v49
                            if true {
                                v50 := GenS4{F0: (v5 >> ((255) & 63)), F1: -10, F2: 469}
                                _ = v50
                            }
                            _ = fi48
                            _ = v49
                        }
                        v51 := GenS5{F0: 1000, F1: (fi45 + fi45), F2: 65535}
                        _ = v51
                    } else {
                        for fi52 := 0; fi52 < 5; fi52++ {
                            if (false && true) {
                                switch (^(fi52)) {
                                case 0:
                                    var v53 uint8 = 225
                                    _ = v53
                                case 127:
                                    if (1e10 <= 0.0) {
                                        var v54 float32 = float32(0.0)
                                        var v55 uint8 = 196
                                        _ = v54
                                        _ = v55
                                    }
                                }
                                fmt.Println(v46)
                                v56 := GenS6{F0: 18446744073709551615, F1: (v1 - (v1 << ((uint(v1)) & 63)))}
                                _ = v56
                            }
                            _ = fi52
                        }
                    }
                default:
                    var v57 int32 = 0
                    _ = v57
                }
                _ = v47
            }
            var v58 uint8 = 127
            if true {
                fmt.Println(v6, v58, v46)
                for fi59 := 0; fi59 < 19; fi59++ {
                    fmt.Println(v58)
                    delete(v4, "a")
                    _ = fi59
                }
                v60 := GenS3{F0: 42, F1: (v5 / (((v5 / ((v5) | 1))) | 1))}
                _ = v60
            } else {
                v61 := GenS7{F0: 0}
                _ = v61
            }
            _ = v58
        } else {
            v62 := GenS6{F0: v6, F1: 2}
            _ = v62
        }
        fmt.Println(v5, v6)
        _ = fi45
        _ = v46
    }
    var v63 uint64 = 0
    v64 := GenS8{F0: "hello", F1: 42, F2: true}
    fmt.Println(v6, v3)
    var v65 uint16 = 64
    fmt.Println(v3)
    switch ("a" + "hello") {
    case "":
        fmt.Println(v6, v63, v65)
    case "\n\t":
        v66 := GenS1{F0: 0}
        _ = v66
    case "foo bar":
        v63 = (^(v63))
    }
    v67 := v4["a"]
    if (0 != 1) {
        if false {
            fmt.Println(v3, v3, v5)
            v2[v5] = (v5 << ((276) & 63))
            var v68 float64 = (0.1 / -2.25)
            _ = v68
        }
        v69 := GenS5{F0: 1, F1: v67, F2: 255}
        _ = v69
    }
    var v70 float32 = ((1e10 + ((1.5 / -2.25) - (float32(-2.25) * ((((3.14159 + -2.25) * float32(3.14159)) / float32(1.5)) * (3.14159 / 0.1))))) + 1.5)
    v71 := []int{v67, (v67 % (((v67 << ((genFn0(528, "hello", (v3 >> ((401) & 63)))) & 63))) | 1)), (v67 / ((v67) | 1)), 127}
    if len(v71) > 0 {
        v71[0] = (v67 >> ((4294967295) & 63))
    }
    switch 42 {
    case 42:
        v72 := GenS9{F0: -515, F1: -2147483648}
        _ = v72
    case -2147483648:
        for fi73 := 0; fi73 < 17; fi73++ {
            for fi74 := 0; fi74 < 18; fi74++ {
                switch ("foo bar" + ("" + "hello")) {
                case "a":
                    var v75 float64 = float64((v70 + ((v70 / float32((3.14159 + float64(v70)))) / v70)))
                    _ = v75
                default:
                    var v76 int16 = -32768
                    _ = v76
                }
                v70 = (1.5 - ((1.5 + float32(((-2.25 - 0.1) / 1.5))) * v70))
                _ = fi74
            }
            for fi77 := 0; fi77 < 7; fi77++ {
                if false {
                    _ = len(v2)
                }
                _ = fi77
            }
            var v78 uint8 = 201
            _ = fi73
            _ = v78
        }
    case -63:
        fmt.Println(v67)
    default:
        if false {
            switch "\n\t" {
            case "hello":
                delete(v2, (v5 + 2147483647))
            case "a":
                _ = len(v71)
            }
            var v79 float32 = 0.0
            fmt.Println(v70, v5)
            _ = v79
        } else {
            v80 := GenS1{F0: 1}
            if (v1 != 65535) {
                switch ("foo bar" + ("foo bar" + "")) {
                case "":
                    for fi81 := 0; fi81 < 19; fi81++ {
                        var v82 float64 = (1.5 * 0.1)
                        v5 = (^(v5))
                        if true {
                            v83 := GenS10{F0: 4294967295}
                            _ = v83
                        } else {
                            switch -1 {
                            case 4294967295:
                                v70 = (v70 * v70)
                            }
                        }
                        _ = fi81
                        _ = v82
                    }
                case "foo bar":
                    if (-1 > 127) {
                        for fi84 := 0; fi84 < 16; fi84++ {
                            for fi85 := 0; fi85 < 11; fi85++ {
                                v65 -= v65
                                _ = fi85
                            }
                            if (715 > (v65 * (v65 << ((genFn0(1000, ("foo bar" + ("\n\t" + "a")), v5)) & 63)))) {
                                for fi86 := 0; fi86 < 14; fi86++ {
                                    var v87 uint8 = 150
                                    _ = fi86
                                    _ = v87
                                }
                            }
                            for fi88 := 0; fi88 < 14; fi88++ {
                                if true {
                                    var v89 uint16 = v65
                                    _ = v89
                                }
                                switch "" {
                                case "hello":
                                    v65 = 1
                                }
                                _ = fi88
                            }
                            _ = fi84
                        }
                        v65 = v65
                        delete(v4, (("" + (("" + "foo bar") + ("a" + ("" + ("foo bar" + "a"))))) + ""))
                    }
                }
                var v90 int64 = 42
                if (149 <= 42) {
                    if false {
                        var v91 uint8 = 100
                        v71 = append(v71, 0)
                        _ = v91
                    } else {
                        if false {
                            var v92 uint16 = (v65 ^ (^(v65)))
                            _ = v92
                        }
                        fmt.Println(v63)
                        var v93 int8 = -1
                        _ = v93
                    }
                    var v94 uint = 65535
                    v2[(v5 ^ 1)] = 4294967295
                    _ = v94
                } else {
                    for fi95 := 0; fi95 < 13; fi95++ {
                        fmt.Println(v70, v70, v5)
                        _ = fi95
                    }
                }
                _ = v90
            } else {
                if (163 > 128) {
                    v65 = v65
                    for fi96 := 0; fi96 < 1; fi96++ {
                        switch "hello" {
                        case "\n\t":
                            _ = len(v71)
                        case "hello":
                            var v97 int64 = -980
                            _ = v97
                        case "a":
                            var v98 int32 = 866
                            _ = v98
                        default:
                            var v99 uint16 = (v1 >> ((4294967295) & 63))
                            _ = v99
                        }
                        var v100 uint = 1
                        _ = fi96
                        _ = v100
                    }
                }
                fmt.Println(v67, v5)
                for fi101 := 0; fi101 < 4; fi101++ {
                    fmt.Println(v63)
                    _ = fi101
                }
            }
            switch "hello" {
            case "":
                var v102 int = -77
                _ = v102
            case "a":
                v67 = (v67 * v67)
            case "\n\t":
                _ = len(v71)
            default:
                if len(v71) >= 2 {
                    v71 = v71[1:len(v71)]
                }
            }
            _ = v80
        }
    }
    for fi103 := 0; fi103 < 2; fi103++ {
        var v104 bool = true
        var v105 int32 = 1
        if (42 != 111) {
            for fi106 := 0; fi106 < 18; fi106++ {
                v107 := GenS11{F0: (^(v3)), F1: 4294967295}
                _ = fi106
                _ = v107
            }
            if true {
                v65 += (v1 * v1)
                v6 -= 1
            }
        } else {
            v108 := GenS12{F0: true, F1: v67, F2: 42}
            v3 = (v3 ^ v5)
            _ = v108
        }
        _ = fi103
        _ = v104
        _ = v105
    }
    switch (v67 % (((-(v67))) | 1)) {
    case 0:
        if (v63 == v6) {
            var v109 bool = (1e10 == 3.14159)
            switch v67 {
            case 42:
                if (v109 || true) {
                    var v110 uint64 = 1
                    var v111 float32 = v70
                    if len(v71) > 0 {
                        v71[0] = v67
                    }
                    _ = v110
                    _ = v111
                } else {
                    v70 -= 1e10
                    if v109 {
                        for fi112 := 0; fi112 < 11; fi112++ {
                            v113 := v4["hello"]
                            var v114 uint32 = (v3 & 0)
                            _ = fi112
                            _ = v113
                            _ = v114
                        }
                        v6 -= 4294967295
                        v109 = (!false)
                    } else {
                        for fi115 := 0; fi115 < 6; fi115++ {
                            if (256 != -1) {
                                fmt.Println(v6, fi115, v3)
                                var v116 bool = true
                                _ = v116
                            } else {
                                v117 := GenS13{F0: -1, F1: (v63 / ((4294967295) | 1))}
                                if len(v71) >= 2 {
                                    v71 = v71[1:len(v71)]
                                }
                                _ = v117
                            }
                            v109 = true
                            for fi118 := 0; fi118 < 7; fi118++ {
                                v70 -= (float32(float64(1.5)) + v70)
                                v119 := GenS8{F0: (("" + "") + ""), F1: -48, F2: (!false)}
                                _ = fi118
                                _ = v119
                            }
                            _ = fi115
                        }
                        fmt.Println(v3, v1, v65)
                        var v120 string = ("" + "foo bar")
                        _ = v120
                    }
                    if len(v71) >= 2 {
                        v71 = v71[1:len(v71)]
                    }
                }
            }
            _ = v109
        } else {
            _ = len(v71)
        }
    }
    v121 := GenS10{F0: 134}
    fmt.Println(v5, v5, v65)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    fmt.Println("v70:", v70)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v67
    _ = v70
    _ = v71
}
