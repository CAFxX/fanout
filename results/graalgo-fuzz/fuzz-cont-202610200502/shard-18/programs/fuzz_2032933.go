package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int16 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 int; F1 int64; F2 int16 }
type GenS10 struct { F0 uint64; F1 int32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 bool; F1 int8 }
type GenS13 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS14 struct { F0 int; F1 int32; F2 int32 }
func genFn0(p0 uint16, p1 uint) uint64 {
    fmt.Println(p0, p1)
    v5 := GenS1{F0: 1, F1: 32767, F2: -452}
    var v6 bool = true
    switch 2147483647 {
    case 1:
        var v7 float32 = 0.1
        _ = v7
    case 0:
        p1 -= p1
    default:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 uint32 = 719
    for fi8 := 0; fi8 < 3; fi8++ {
        switch v1 {
        case "hello":
            if (850 != -1) {
                fmt.Println(fi8)
            }
        default:
            v4 = 33
        }
        var v9 int64 = 0
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            if true {
                if true {
                    v12 := GenS1{F0: 42, F1: 32767, F2: 32767}
                    switch "hello" {
                    case "foo bar":
                        v13 := GenS2{F0: -32, F1: 18446744073709551615}
                        _ = v13
                    case "a":
                        var v14 float32 = (1.5 * (-2.25 / 1.5))
                        _ = v14
                    default:
                        var v15 int = 247
                        _ = v15
                    }
                    v4 += v4
                    _ = v12
                }
            }
            if true {
                v16 := []int{-85, fi10}
                _ = v16
            }
            switch (("a" + v1) + ((v1 + (("\n\t" + (v1 + v1)) + v1)) + "hello")) {
            case "foo bar":
                var v17 uint8 = 127
                _ = v17
            case "a":
                switch (v1 + v1) {
                case "a":
                    if true {
                        v18 := []int{(-(fi11)), (-(fi11))}
                        _ = v18
                    }
                case "\n\t":
                    var v19 int8 = 42
                    _ = v19
                }
                var v20 bool = false
                var v21 uint16 = 1
                _ = v20
                _ = v21
            default:
                v22 := v3[(fi11 / ((fi11) | 1))]
                _ = v22
            }
            _ = len(v2)
            v23 := []int{fi10}
            _ = fi11
            _ = v23
        }
        if (true && true) {
            for fi24 := 0; fi24 < 4; fi24++ {
                v1 = "foo bar"
                _ = fi24
            }
            v25 := []int{fi10, fi10, 0}
            switch v1 {
            case "a":
                v1 = (v1 + (v1 + v1))
            case "hello":
                if (((false && true) || false) || (!false)) {
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                    switch v1 {
                    case "foo bar":
                        if len(v25) > 0 {
                            v25[0] = (fi10 + fi10)
                        }
                    case "hello":
                        if true {
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            v26 := GenS1{F0: 523, F1: 1, F2: -306}
                            _ = v26
                        }
                    case "\n\t":
                        delete(v3, (fi10 ^ -2147483648))
                    default:
                        switch 4294967295 {
                        case 128:
                            var v27 int16 = 833
                            _ = v27
                        case -2147483648:
                            if false {
                                if (!false) {
                                    v1 = v1
                                }
                                v4 = 488
                            } else {
                                switch v1 {
                                case "hello":
                                    var v28 uint32 = 1
                                    _ = v28
                                case "a":
                                    var v29 uint16 = 0
                                    _ = v29
                                default:
                                    v1 = v1
                                }
                                if len(v25) > 0 {
                                    v25[0] = 9223372036854775807
                                }
                                v3[128] = fi10
                            }
                        }
                    }
                    v30 := GenS3{F0: -608, F1: 41}
                    _ = v30
                }
            default:
                v31 := GenS4{F0: 1}
                _ = v31
            }
            _ = v25
        } else {
            var v32 uint = 0
            for fi33 := 0; fi33 < 19; fi33++ {
                var v34 string = v1
                switch fi10 {
                case -9223372036854775808:
                    switch fi33 {
                    case 128:
                        fmt.Println(fi10, v1, v34)
                    }
                default:
                    if (!(!false)) {
                        var v35 string = v34
                        var v36 string = "foo bar"
                        v37 := GenS5{F0: 42, F1: -53, F2: (^(v32))}
                        _ = v35
                        _ = v36
                        _ = v37
                    } else {
                        v38 := GenS6{F0: 100, F1: (v32 + v32)}
                        v39 := []int{fi33, (fi10 * fi33)}
                        _ = v38
                        _ = v39
                    }
                }
                if (2147483647 >= 42) {
                    var v40 int64 = -1
                    v41 := GenS6{F0: -95, F1: (v32 * v32)}
                    var v42 uint32 = (^(v4))
                    _ = v40
                    _ = v41
                    _ = v42
                } else {
                    switch 11 {
                    case 1:
                        if (!(uint64(v4) == 0)) {
                            if true {
                                if (1 == 1000) {
                                    if (false && (true || false)) {
                                        var v43 int8 = 1
                                        _ = v43
                                    } else {
                                        v44 := []int{int(v32)}
                                        fmt.Println(v34, v4)
                                        var v45 int8 = -1
                                        _ = v44
                                        _ = v45
                                    }
                                    var v46 uint = (v32 & v32)
                                    v46 = (v32 << (((v32 << ((v46) & 63))) & 63))
                                    _ = v46
                                }
                            }
                        }
                    case -1:
                        if (v4 < 1) {
                            switch ("foo bar" + v34) {
                            case "\n\t":
                                fmt.Println(fi33, v34, v4)
                            case "a":
                                var v47 int64 = -1
                                _ = v47
                            }
                            if false {
                                if (((v34 + "a") + (v34 + v1)) != ("foo bar" + ("a" + v1))) {
                                    v48 := []int{}
                                    _ = v48
                                } else {
                                    v49 := GenS7{F0: 1, F1: "foo bar"}
                                    v32 = v32
                                    var v50 int8 = -112
                                    _ = v49
                                    _ = v50
                                }
                            } else {
                                v1 = v1
                            }
                        }
                    default:
                        var v51 uint8 = 33
                        _ = v51
                    }
                    var v52 uint64 = 2
                    switch v34 {
                    case "\n\t":
                        v53 := GenS8{F0: 315, F1: 1}
                        _ = v53
                    case "hello":
                        fmt.Println(v32, v52)
                    case "":
                        if (!false) {
                            v54 := GenS9{F0: -830, F1: -626, F2: 0}
                            _ = v54
                        }
                    default:
                        fmt.Println(v34, fi10, v1)
                        var v55 string = "hello"
                        _ = v55
                    }
                    _ = v52
                }
                _ = fi33
                _ = v34
            }
            fmt.Println(v32, v4, v4)
            _ = v32
        }
        v4 -= 0
        _ = fi10
    }
    for fi56 := 0; fi56 < 15; fi56++ {
        for fi57 := 0; fi57 < 13; fi57++ {
            switch fi56 {
            case 238:
                if (v1 > (v1 + (v1 + v1))) {
                    v1 = "foo bar"
                    var v58 int = fi57
                    _ = v58
                }
            case 2147483647:
                v59 := GenS7{F0: 2147483648, F1: v1}
                _ = v59
            case 32767:
                if true {
                    if (((("" + (("foo bar" + "foo bar") + (v1 + (((v1 + v1) + v1) + ("" + v1))))) + (v1 + "")) + v1) > v1) {
                        v1 = (v1 + v1)
                        v4 -= (v4 - 73)
                    } else {
                        switch (v1 + v1) {
                        case "foo bar":
                            v60 := GenS10{F0: 4294967295, F1: 0}
                            _ = v60
                        case "":
                            var v61 int = 413
                            _ = v61
                        case "\n\t":
                            v62 := GenS1{F0: genFn0(165, 1), F1: 0, F2: 42}
                            _ = v62
                        default:
                            v4 = 2147483648
                            var v63 int16 = -32768
                            _ = v63
                        }
                        var v64 string = v1
                        v2[v64] = fi57
                        _ = v64
                    }
                } else {
                    fmt.Println(v4)
                    var v65 uint64 = 880
                    _ = v65
                }
            }
            _ = fi57
        }
        switch (v1 + v1) {
        case "hello":
            v66 := GenS3{F0: 1, F1: 175}
            _ = v66
        }
        var v67 uint64 = 2
        _ = fi56
        _ = v67
    }
    var v68 float64 = 1e10
    for fi69 := 0; fi69 < 5; fi69++ {
        for fi70 := 0; fi70 < 13; fi70++ {
            if true {
                var v71 int = (fi69 % ((fi69) | 1))
                switch (fi70 + 851) {
                case 288:
                    if (!(((v4 < v4) && ((((v4 - 492) > (v4 * 67)) && (v1 == "a")) || (-2147483648 != -816))) && true)) {
                        switch v1 {
                        case "a":
                            v72 := GenS11{F0: 1}
                            _ = v72
                        default:
                            if true {
                                if (-111 <= -128) {
                                    var v73 uint = 582
                                    if (37 <= 42) {
                                        var v74 uint32 = 310
                                        var v75 int = fi70
                                        _ = len(v3)
                                        _ = v74
                                        _ = v75
                                    }
                                    if true {
                                        switch (fi69 | fi70) {
                                        case 2147483647:
                                            var v76 bool = false
                                            _ = v76
                                        case 240:
                                            if true {
                                                var v77 uint64 = 2
                                                _ = v77
                                            } else {
                                                switch "foo bar" {
                                                case "hello":
                                                    if false {
                                                        v71 -= 912
                                                        fmt.Println(v4)
                                                    } else {
                                                        var v78 int16 = 451
                                                        var v79 uint32 = 2147483648
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                case "a":
                                                    _ = len(v3)
                                                }
                                                v73 -= 65535
                                                fmt.Println(v71, v68, v4)
                                            }
                                            var v80 string = (((("hello" + (((("a" + v1) + (v1 + (v1 + (v1 + "\n\t")))) + "") + v1)) + ((("\n\t" + ("" + ((v1 + v1) + v1))) + "\n\t") + v1)) + ((("\n\t" + v1) + (v1 + (((v1 + v1) + v1) + (((v1 + (v1 + v1)) + "") + (v1 + v1))))) + v1)) + (v1 + "a"))
                                            _ = v80
                                        case 879:
                                            switch (v1 + "\n\t") {
                                            case "":
                                                v68 -= v68
                                            case "\n\t":
                                                delete(v3, fi70)
                                            default:
                                                v73 -= v73
                                            }
                                        default:
                                            fmt.Println(v4, fi70, v1)
                                            switch v1 {
                                            case "":
                                                var v81 string = v1
                                                _ = v81
                                            case "foo bar":
                                                var v82 string = (v1 + v1)
                                                _ = v82
                                            default:
                                                v73 = v73
                                                var v83 float32 = 0.1
                                                var v84 int8 = -128
                                                _ = v83
                                                _ = v84
                                            }
                                            var v85 int = fi69
                                            _ = v85
                                        }
                                        if true {
                                            if (!false) {
                                                var v86 int64 = 9223372036854775807
                                                var v87 uint16 = 65535
                                                _ = v86
                                                _ = v87
                                            } else {
                                                var v88 string = (v1 + "\n\t")
                                                v89 := v3[-228]
                                                if ((v1 <= ("hello" + "")) || (!false)) {
                                                    var v90 uint8 = 128
                                                    v88 = v1
                                                    _ = v90
                                                }
                                                v89 = fi69
                                                _ = v88
                                                _ = v89
                                            }
                                            v1 = "\n\t"
                                        }
                                    }
                                    _ = v73
                                }
                                v1 = v1
                                delete(v2, v1)
                                if (!true) {
                                    if true {
                                        v68 += v68
                                        if (false && false) {
                                            var v91 uint64 = 2
                                            fmt.Println(fi69)
                                            v68 += (v68 + 1e10)
                                            var v92 uint8 = 0
                                            if true {
                                                v4 = (^(v4))
                                                var v93 uint32 = (^(v4))
                                                v91 -= v91
                                                _ = v93
                                            } else {
                                                fmt.Println(fi69)
                                                v1 = v1
                                                var v94 float64 = v68
                                                _ = v94
                                            }
                                            switch v1 {
                                            case "\n\t":
                                                v4 = 1
                                            case "":
                                                var v95 uint32 = v4
                                                _ = v95
                                            case "foo bar":
                                                v96 := []int{(v71 ^ (fi70 + (fi70 | (fi69 + 2147483647))))}
                                                _ = v96
                                            default:
                                                v92 += 42
                                            }
                                            _ = v91
                                            _ = v92
                                        }
                                        var v97 uint64 = 0
                                        _ = v97
                                    }
                                    var v98 int64 = -215
                                    fmt.Println(v71)
                                    _ = v98
                                } else {
                                    var v99 uint32 = 2147483647
                                    _ = v99
                                }
                            }
                        }
                        var v100 int16 = -32768
                        _ = v100
                    }
                case -519:
                    var v101 float32 = 3.14159
                    _ = v101
                }
                _ = v71
            } else {
                v102 := GenS6{F0: -26, F1: 255}
                v1 = "a"
                v4 = (v4 >> ((65535) & 63))
                var v103 int16 = -581
                _ = v102
                _ = v103
            }
            _ = len(v2)
            _ = fi70
        }
        var v104 int16 = 42
        _ = fi69
        _ = v104
    }
    if (v1 < v1) {
        var v105 bool = true
        switch (((v1 + (v1 + v1)) + v1) + v1) {
        case "\n\t":
            if (42 < 32768) {
                var v106 int16 = 1000
                _ = v106
            }
        }
        if true {
            v107 := GenS12{F0: true, F1: -128}
            _ = v107
        } else {
            v4 = 65535
            for fi108 := 0; fi108 < 10; fi108++ {
                var v109 int8 = 1
                fmt.Println(v109)
                _ = fi108
                _ = v109
            }
        }
        _ = v105
    }
    v110 := GenS6{F0: 1, F1: 16}
    if (!(42 > -2147483648)) {
        for fi111 := 0; fi111 < 6; fi111++ {
            for fi112 := 0; fi112 < 18; fi112++ {
                switch 296 {
                case 94:
                    fmt.Println(v1, v68)
                case 127:
                    v68 -= 1e10
                }
                fmt.Println(fi111, v68, v4)
                if true {
                    var v113 uint8 = 128
                    if true {
                        if true {
                            var v114 float32 = (1e10 / 1.0)
                            fmt.Println(fi111, fi111)
                            _ = v114
                        } else {
                            fmt.Println(v1, fi112)
                            v1 = (((v1 + "") + v1) + "")
                            var v115 bool = (2 > 0)
                            var v116 int = -806
                            v117 := GenS8{F0: 0, F1: 1000}
                            _ = v115
                            _ = v116
                            _ = v117
                        }
                    } else {
                        var v118 uint = 4294967295
                        if (42 == 2) {
                            v119 := GenS2{F0: 56, F1: (^(v118))}
                            fmt.Println(v1)
                            _ = v119
                        } else {
                            var v120 float64 = v68
                            _ = v120
                        }
                        _ = v118
                    }
                    var v121 int8 = -128
                    _ = v113
                    _ = v121
                } else {
                    _ = len(v2)
                    v68 += ((v68 - float64(((float32(float64(((1.5 * float32(v68)) - (1e10 + 1.5)))) + 0.0) - -2.25))) + ((float64(float32((v68 + v68))) - v68) / 1.0))
                }
                fmt.Println(fi111, v4)
                v122 := GenS13{F0: 1, F1: v4, F2: 216}
                _ = fi112
                _ = v122
            }
            var v123 string = (("" + (((((v1 + "a") + (v1 + "foo bar")) + v1) + v1) + v1)) + "foo bar")
            _ = fi111
            _ = v123
        }
    }
    v124 := []int{128, 32768, -1, -2147483648}
    for fi125 := 0; fi125 < 3; fi125++ {
        v4 = (^(v4))
        var v126 bool = (true || (42 >= 42))
        _ = fi125
        _ = v126
    }
    v127 := GenS14{F0: 127, F1: 42, F2: 0}
    var v128 int8 = -1
    fmt.Println(v128, v4)
    v68 += 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v124:", len(v124), v124)
    fmt.Println("v128:", v128)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v110
    _ = v124
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v68
}
