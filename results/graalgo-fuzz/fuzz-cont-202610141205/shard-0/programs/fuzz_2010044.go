package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint; F1 int; F2 uint16 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 int32; F1 uint64 }
func genFn0(p0 int32, p1 uint, p2 string) uint {
    if false {
        var v5 bool = (!false)
        p2 = p2
        _ = v5
    } else {
        var v6 bool = true
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{}
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint16 = 521
            var v11 float64 = 3.14159
            var v12 bool = false
            v13 := GenS1{F0: 2, F1: 0, F2: 127}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
            _ = v13
        }
        var v14 float32 = (1.5 - 3.14159)
        _ = fi7
        _ = v14
        _ = v8
    }
    var v15 uint = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return p1
}

func main() {
    var v1 float32 = -2.25
    var v2 int8 = 1
    var v3 float32 = (-2.25 - (v1 + 3.14159))
    v4 := make(map[int]bool)
    v16 := GenS1{F0: 127, F1: (v2 | (v2 * 42)), F2: 42}
    if true {
        if false {
            switch "a" {
            case "hello":
                v1 += 0.0
            default:
                fmt.Println(v1, v1)
            }
            if (1.5 >= (-2.25 + 3.14159)) {
                v3 = v1
                for fi17 := 0; fi17 < 9; fi17++ {
                    for fi18 := 0; fi18 < 15; fi18++ {
                        fmt.Println(fi18, v3)
                        if (1 != 11) {
                            v3 += -2.25
                            v19 := GenS1{F0: 42, F1: v2, F2: 183}
                            _ = v19
                        } else {
                            if (!(!((!((fi17 >> ((4294967295) & 63)) < (fi18 ^ 2147483647))) || (255 == 32767)))) {
                                v20 := v4[int(v2)]
                                if ((true && v20) && false) {
                                    if (!false) {
                                        v20 = (v2 > (v2 - v2))
                                        var v21 int = -1
                                        _ = v21
                                    } else {
                                        v2 = (v2 << ((289) & 63))
                                        var v22 uint8 = 127
                                        _ = v22
                                    }
                                }
                                v3 -= ((float32(3.14159) - 1e10) + v3)
                                _ = v20
                            }
                            v23 := v4[fi17]
                            var v24 uint = 2
                            _ = v23
                            _ = v24
                        }
                        _ = fi18
                    }
                    _ = fi17
                }
            } else {
                if (2 <= 2147483648) {
                    v3 += (((float32(3.14159) - v3) / (((1e10 + v3) - (0.0 - 1.5)) - v3)) + float32(1.5))
                    v3 = v1
                } else {
                    v3 -= (0.1 * v1)
                    switch ((("a" + "hello") + (("foo bar" + "\n\t") + ("" + ("a" + ("" + "hello"))))) + "foo bar") {
                    case "a":
                        v4[775] = (65535 < 255)
                    }
                }
                var v25 float64 = ((0.1 - 1e10) / (0.0 - float64((((v3 * 1.5) - (v1 - 0.0)) - 3.14159))))
                v3 = v1
                _ = v25
            }
            v26 := []int{0, -430, -2147483648, 1}
            _ = v26
        } else {
            v1 = v3
            v27 := GenS1{F0: 128, F1: (^(v2)), F2: 68}
            _ = v27
        }
        var v28 bool = false
        if v28 {
            fmt.Println(v3, v2)
        }
        _ = v28
    } else {
        for fi29 := 0; fi29 < 2; fi29++ {
            delete(v4, 0)
            v2 = v2
            var v30 string = ("\n\t" + "\n\t")
            _ = fi29
            _ = v30
        }
        switch (("a" + ((("" + ("foo bar" + "hello")) + ("foo bar" + ((((("\n\t" + "a") + "\n\t") + "") + "hello") + "\n\t"))) + ("\n\t" + "a"))) + "foo bar") {
        case "a":
            v3 = ((v3 / (v3 + v3)) + v1)
        case "hello":
            for fi31 := 0; fi31 < 2; fi31++ {
                v3 = (v3 / 1e10)
                v32 := GenS2{F0: (v2 << ((1) & 63))}
                _ = fi31
                _ = v32
            }
        default:
            v33 := GenS2{F0: (v2 >> ((genFn0(-523, 18446744073709551615, ("a" + ("a" + "foo bar")))) & 63))}
            _ = v33
        }
        v34 := GenS2{F0: 42}
        _ = v34
    }
    var v35 float64 = 0.0
    var v36 int = -1
    var v37 int64 = int64(v36)
    if false {
        var v38 uint32 = 2
        _ = v38
    } else {
        for fi39 := 0; fi39 < 12; fi39++ {
            for fi40 := 0; fi40 < 4; fi40++ {
                switch v36 {
                case 2147483647:
                    v41 := GenS3{F0: -1, F1: -32768}
                    _ = v41
                }
                _ = fi40
            }
            if true {
                v3 -= float32(v35)
            } else {
                fmt.Println(v36, fi39)
                if true {
                    v36 += 2147483647
                    if false {
                        for fi42 := 0; fi42 < 7; fi42++ {
                            var v43 int = -212
                            _ = fi42
                            _ = v43
                        }
                        if false {
                            if true {
                                if true {
                                    var v44 uint32 = 2147483648
                                    var v45 uint16 = 1
                                    if true {
                                        for fi46 := 0; fi46 < 17; fi46++ {
                                            if false {
                                                var v47 uint64 = 0
                                                v48 := GenS1{F0: 244, F1: (v2 ^ int8(fi39)), F2: 255}
                                                fmt.Println(v47, fi46)
                                                var v49 int16 = -924
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            }
                                            var v50 uint64 = 122
                                            if (1 != 32767) {
                                                switch "" {
                                                case "foo bar":
                                                    var v51 uint8 = 1
                                                    _ = v51
                                                default:
                                                    delete(v4, -1)
                                                }
                                            }
                                            v52 := []int{fi39}
                                            _ = fi46
                                            _ = v50
                                            _ = v52
                                        }
                                        var v53 int = (^(v36))
                                        _ = v53
                                    } else {
                                        v35 -= (v35 - (v35 * v35))
                                        var v54 float64 = v35
                                        v2 += int8(v37)
                                        _ = v54
                                    }
                                    _ = v44
                                    _ = v45
                                }
                            }
                        } else {
                            v35 = 3.14159
                            var v55 int32 = -499
                            v56 := GenS2{F0: int8(v37)}
                            _ = v55
                            _ = v56
                        }
                    }
                    if true {
                        v57 := GenS2{F0: (v2 * v2)}
                        switch -458 {
                        case 127:
                            var v58 float32 = v1
                            _ = v58
                        }
                        _ = v57
                    }
                } else {
                    var v59 string = "foo bar"
                    v2 += (v2 ^ -122)
                    _ = v59
                }
            }
            var v60 uint16 = 312
            _ = fi39
            _ = v60
        }
    }
    var v61 uint64 = 4294967295
    v62 := GenS1{F0: 1, F1: (v2 % ((65) | 1)), F2: 128}
    var v63 string = "a"
    v64 := []int{(v36 << ((255) & 63)), 2147483647, v36, 227}
    if (v37 <= -912) {
        fmt.Println(v63, v36)
        if true {
            v2 += int8(v37)
        } else {
            var v65 string = v63
            _ = v65
        }
    }
    v1 = (v1 - (v3 * v1))
    if false {
        for fi66 := 0; fi66 < 20; fi66++ {
            v3 = v1
            var v67 float32 = v3
            _ = fi66
            _ = v67
        }
    } else {
        v61 = 18446744073709551615
        var v68 uint8 = 0
        _ = v68
    }
    var v69 uint16 = 604
    v36 += -53
    if (true && true) {
        v69 = (^(v69))
        switch v63 {
        case "hello":
            var v70 uint32 = 107
            _ = v70
        case "foo bar":
            if ((431 == 1) || (!false)) {
                switch ((v63 + ((((v63 + v63) + (v63 + (v63 + v63))) + (v63 + (v63 + ""))) + "foo bar")) + v63) {
                case "":
                    v71 := GenS3{F0: -1, F1: 74}
                    _ = v71
                case "foo bar":
                    v72 := GenS2{F0: v2}
                    _ = v72
                default:
                    if true {
                        for fi73 := 0; fi73 < 1; fi73++ {
                            for fi74 := 0; fi74 < 5; fi74++ {
                                v69 -= (v69 & v69)
                                fmt.Println(v3)
                                var v75 uint8 = uint8(v69)
                                _ = fi74
                                _ = v75
                            }
                            if (v37 != v37) {
                                v64 = append(v64, fi73)
                                switch ("" + v63) {
                                case "foo bar":
                                    v69 = 65535
                                case "a":
                                    if (0 < (v69 - 42)) {
                                        var v76 uint = 0
                                        _ = v76
                                    } else {
                                        v36 += 1
                                        if (4294967295 > 621) {
                                            if (!(!((v63 + v63) == v63))) {
                                                for fi77 := 0; fi77 < 6; fi77++ {
                                                    var v78 int64 = (v37 * (v37 - (v37 << ((genFn0(-974, 0, "\n\t")) & 63))))
                                                    v79 := GenS4{F0: 42, F1: 127, F2: (v69 + v69)}
                                                    v4[(fi77 + fi77)] = true
                                                    _ = fi77
                                                    _ = v78
                                                    _ = v79
                                                }
                                                v36 = -2147483648
                                            } else {
                                                _ = len(v64)
                                            }
                                            if true {
                                                if (false || (!((!false) && ("a" != (v63 + v63))))) {
                                                    if (float64((0.0 - 0.0)) >= v35) {
                                                        var v80 int8 = v2
                                                        var v81 uint32 = 36
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                } else {
                                                    fmt.Println(v63, v61, v69)
                                                    if (65535 >= 4294967295) {
                                                        var v82 bool = (v69 < 278)
                                                        _ = v82
                                                    } else {
                                                        for fi83 := 0; fi83 < 9; fi83++ {
                                                            if false {
                                                                var v84 uint16 = (v69 % (((v69 / ((v69) | 1))) | 1))
                                                                switch (v63 + v63) {
                                                                case "foo bar":
                                                                    fmt.Println(v35, v3)
                                                                    switch "a" {
                                                                    case "":
                                                                        var v85 string = v63
                                                                        _ = v85
                                                                    case "foo bar":
                                                                        v84 += v84
                                                                    case "a":
                                                                        var v86 uint16 = uint16(v61)
                                                                        _ = v86
                                                                    }
                                                                    fmt.Println(v61)
                                                                    var v87 uint = 822
                                                                    var v88 string = "hello"
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                                var v89 int = (fi73 - (^(v36)))
                                                                switch "" {
                                                                case "hello":
                                                                    var v90 int32 = 837
                                                                    _ = v90
                                                                case "\n\t":
                                                                    v91 := GenS4{F0: 0, F1: 2147483647, F2: (v84 ^ 790)}
                                                                    _ = v91
                                                                case "":
                                                                    v36 = fi83
                                                                default:
                                                                    v92 := GenS5{F0: v89, F1: (v2 * (^(v2)))}
                                                                    if (v61 <= v61) {
                                                                        v93 := GenS3{F0: 775, F1: -189}
                                                                        var v94 float32 = 1e10
                                                                        _ = v93
                                                                        _ = v94
                                                                    }
                                                                    _ = v92
                                                                }
                                                                switch (v36 >> ((uint(v61)) & 63)) {
                                                                case 1:
                                                                    if true {
                                                                        var v95 int = -9223372036854775808
                                                                        var v96 uint64 = v61
                                                                        switch ("hello" + "hello") {
                                                                        case "a":
                                                                            _ = len(v4)
                                                                            var v97 float64 = 1.5
                                                                            _ = v97
                                                                        case "":
                                                                            v98 := v4[v95]
                                                                            v4[0] = (!(!(((((v63 + v63) + (((v63 + (v63 + (v63 + "\n\t"))) + v63) + (v63 + v63))) + (v63 + v63)) + v63) > v63)))
                                                                            _ = v98
                                                                        case "hello":
                                                                            v35 = v35
                                                                            fmt.Println(v36)
                                                                            var v99 int16 = 32767
                                                                            _ = v99
                                                                        default:
                                                                            if false {
                                                                                if (v61 != 4294967295) {
                                                                                    v100 := GenS2{F0: v2}
                                                                                    var v101 int64 = v37
                                                                                    _ = v100
                                                                                    _ = v101
                                                                                } else {
                                                                                    switch (v95 | 128) {
                                                                                    case 128:
                                                                                        if ((!(-740 < 1)) || ((v2 / ((v2) | 1)) == -10)) {
                                                                                            var v102 int16 = 1
                                                                                            v103 := GenS3{F0: 249, F1: v102}
                                                                                            v63 = v63
                                                                                            var v104 int16 = (^(v102))
                                                                                            _ = v102
                                                                                            _ = v103
                                                                                            _ = v104
                                                                                        } else {
                                                                                            v96 = (v96 / ((v61) | 1))
                                                                                            var v105 uint16 = (v84 ^ 656)
                                                                                            _ = v105
                                                                                        }
                                                                                    case 1:
                                                                                        if len(v64) >= 2 {
                                                                                            v64 = v64[1:len(v64)]
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v2 = v2
                                                                            }
                                                                        }
                                                                        v63 = v63
                                                                        fmt.Println(v2, v95)
                                                                        if false {
                                                                            fmt.Println(v1)
                                                                        } else {
                                                                            if (false && (false && (341 <= v61))) {
                                                                                var v106 int64 = -519
                                                                                v4[128] = (!true)
                                                                                _ = v106
                                                                            } else {
                                                                                switch ((((((((v63 + v63) + ("a" + v63)) + (v63 + v63)) + ("a" + v63)) + ((v63 + ("hello" + (v63 + "a"))) + (v63 + v63))) + v63) + (v63 + (v63 + v63))) + (v63 + "")) {
                                                                                case "\n\t":
                                                                                    v1 = 0.1
                                                                                }
                                                                                v107 := GenS6{F0: 0, F1: 752, F2: 42}
                                                                                _ = v107
                                                                            }
                                                                            var v108 uint32 = 0
                                                                            fmt.Println(v1, fi83, v1)
                                                                            _ = v108
                                                                        }
                                                                        fmt.Println(v3, v84, v61)
                                                                        _ = v95
                                                                        _ = v96
                                                                    } else {
                                                                        v37 -= 1
                                                                        v89 -= (fi73 << ((255) & 63))
                                                                    }
                                                                case -170:
                                                                    switch v63 {
                                                                    case "\n\t":
                                                                        v2 = 13
                                                                    case "hello":
                                                                        var v109 float32 = (v1 + v1)
                                                                        _ = v109
                                                                    case "":
                                                                        var v110 float32 = v3
                                                                        _ = v110
                                                                    default:
                                                                        var v111 int64 = v37
                                                                        _ = v111
                                                                    }
                                                                    v35 = ((v35 * v35) - v35)
                                                                case -1:
                                                                    if false {
                                                                        v63 = "hello"
                                                                        fmt.Println(fi83, v2)
                                                                    } else {
                                                                        v36 -= (v89 * fi73)
                                                                    }
                                                                }
                                                                _ = v84
                                                                _ = v89
                                                            }
                                                            _ = fi83
                                                        }
                                                        var v112 uint16 = 42
                                                        v63 = v63
                                                        _ = v112
                                                    }
                                                    v1 -= v3
                                                }
                                            }
                                        }
                                        var v113 int = 32768
                                        _ = v113
                                    }
                                case "\n\t":
                                    v63 = ((v63 + v63) + v63)
                                default:
                                    _ = len(v4)
                                }
                            }
                            _ = fi73
                        }
                        v114 := GenS4{F0: genFn0(-381, uint(v61), v63), F1: (^(v36)), F2: 65535}
                        for fi115 := 0; fi115 < 14; fi115++ {
                            v37 -= -9223372036854775808
                            v116 := GenS1{F0: 42, F1: (v2 + int8(v36)), F2: 0}
                            for fi117 := 0; fi117 < 14; fi117++ {
                                switch int(v61) {
                                case 816:
                                    fmt.Println(v69, v3, v36)
                                case 32767:
                                    _ = len(v4)
                                case 589:
                                    v118 := GenS7{F0: v61, F1: -32768}
                                    _ = v118
                                }
                                v119 := GenS8{F0: 42, F1: v61}
                                var v120 bool = (v35 <= (1.5 / (((v35 * 1e10) + (1.5 - float64(v3))) / float64((v1 / v3)))))
                                _ = fi117
                                _ = v119
                                _ = v120
                            }
                            _ = fi115
                            _ = v116
                        }
                        _ = v114
                    } else {
                        if len(v64) >= 2 {
                            v64 = v64[1:len(v64)]
                        }
                        fmt.Println(v69, v63, v2)
                    }
                }
                v121 := v4[v36]
                _ = v121
            } else {
                var v122 int8 = (v2 + int8(v37))
                _ = v122
            }
        default:
            if false {
                fmt.Println(v36, v37)
                v63 = "hello"
            } else {
                var v123 string = v63
                var v124 uint8 = 128
                _ = v123
                _ = v124
            }
        }
    } else {
        v61 = (v61 << ((401) & 63))
        v125 := GenS8{F0: 42, F1: v61}
        var v126 float32 = (v3 * float32(v35))
        _ = v125
        _ = v126
    }
    fmt.Println(v61, v36)
    v127 := GenS4{F0: 18446744073709551615, F1: (-(v36)), F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v64:", len(v64), v64)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v127
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v69
}
