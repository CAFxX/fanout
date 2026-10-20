package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int64; F1 bool; F2 string }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32; F1 bool; F2 uint32 }
type GenS10 struct { F0 bool; F1 int32; F2 string }
type GenS11 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS12 struct { F0 int32; F1 int; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    var v3 bool = (false && (!true))
    var v4 string = "foo bar"
    var v5 uint32 = 2147483647
    v6 := make(map[uint32]string)
    if v3 {
        switch 42 {
        case 1:
            _ = len(v1)
        case 0:
            if ((196 != 624) || v3) {
                v4 = v4
                fmt.Println(v4, v4)
                v7 := []int{-2147483648, 4294967295, 42, -776}
                _ = v7
            }
        case 4294967295:
            fmt.Println(v4, v4)
        default:
            for fi8 := 0; fi8 < 8; fi8++ {
                if false {
                    v9 := []int{fi8, (fi8 | (fi8 ^ fi8)), 32768, 1}
                    for fi10 := 0; fi10 < 11; fi10++ {
                        var v11 float32 = 0.1
                        fmt.Println(v3, v3, v5)
                        var v12 float32 = -2.25
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = v9
                } else {
                    v13 := GenS1{F0: 2147483647, F1: 195}
                    delete(v6, v5)
                    v1[(v5 + v5)] = fi8
                    _ = v13
                }
                fmt.Println(v5, v3)
                _ = fi8
            }
        }
    }
    v14 := GenS1{F0: 1, F1: 0}
    var v15 int = 4294967295
    v5 = 294
    if (v4 <= "a") {
        var v16 uint32 = 1
        _ = v16
    } else {
        var v17 int32 = 42
        for fi18 := 0; fi18 < 2; fi18++ {
            if false {
                if v3 {
                    switch fi18 {
                    case 128:
                        v19 := GenS1{F0: (v17 | v17), F1: 160}
                        _ = v19
                    case -1:
                        for fi20 := 0; fi20 < 2; fi20++ {
                            v21 := GenS1{F0: -51, F1: 222}
                            _ = fi20
                            _ = v21
                        }
                    default:
                        var v22 uint32 = (^(v5))
                        _ = v22
                    }
                } else {
                    for fi23 := 0; fi23 < 3; fi23++ {
                        switch (v15 % ((685) | 1)) {
                        case 127:
                            switch "hello" {
                            case "a":
                                var v24 string = "a"
                                _ = v24
                            case "hello":
                                v5 = v5
                            case "\n\t":
                                if true {
                                    switch v4 {
                                    case "hello":
                                        switch v4 {
                                        case "\n\t":
                                            v4 = "\n\t"
                                        default:
                                            v3 = v3
                                        }
                                    case "\n\t":
                                        v2["hello"] = fi23
                                    default:
                                        fmt.Println(v17)
                                    }
                                    if v3 {
                                        if v3 {
                                            v6[(v5 >> ((788) & 63))] = v4
                                        }
                                        if v3 {
                                            var v25 uint8 = 128
                                            _ = v25
                                        }
                                        var v26 float32 = (((0.0 - -2.25) * ((0.1 * (((float32(3.14159) * (((-2.25 * 3.14159) + (float32((1.5 - float64(1e10))) - (1e10 + -2.25))) / 1e10)) * 0.1) * 0.1)) - 1e10)) / ((((float32(0.0) / 3.14159) - 1e10) + (float32(3.14159) / 1.0)) + (float32(0.1) + 0.1)))
                                        _ = v26
                                    } else {
                                        v27 := []int{(fi18 >> ((65535) & 63)), v15}
                                        if v3 {
                                            v15 = 128
                                            fmt.Println(v5, v3, fi23)
                                        }
                                        var v28 float32 = 3.14159
                                        _ = v27
                                        _ = v28
                                    }
                                } else {
                                    if (42 != 118) {
                                        v29 := GenS2{F0: 379, F1: false, F2: "a"}
                                        v15 += (v15 * (fi18 / ((-9223372036854775808) | 1)))
                                        _ = v29
                                    } else {
                                        v17 -= 2147483647
                                        var v30 uint = 255
                                        _ = v30
                                    }
                                    if v3 {
                                        v15 = (^(v15))
                                        fmt.Println(v15)
                                        var v31 uint64 = 2
                                        fmt.Println(v31, v3)
                                        v5 = v5
                                        _ = v31
                                    }
                                    v32 := []int{fi23, (fi18 ^ 9223372036854775807), (fi18 + 9223372036854775807)}
                                    _ = v32
                                }
                            default:
                                v3 = (fi23 >= fi23)
                            }
                            var v33 uint32 = (^(v5))
                            _ = v33
                        case 128:
                            var v34 uint64 = 929
                            _ = v34
                        default:
                            v35 := GenS3{F0: v4, F1: v3}
                            _ = v35
                        }
                        fmt.Println(fi18, v3, v15)
                        _ = fi23
                    }
                }
                if (-742 > 32767) {
                    v36 := []int{v15}
                    _ = v36
                }
            }
            var v37 float64 = 0.1
            _ = fi18
            _ = v37
        }
        _ = v17
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        fmt.Println(v5, v3)
        v39 := []int{-406, fi38}
        _ = fi38
        _ = v39
    }
    if v3 {
        if v3 {
            fmt.Println(v3, v15)
            fmt.Println(v4, v4)
        } else {
            for fi40 := 0; fi40 < 20; fi40++ {
                var v41 uint64 = 0
                v3 = (42 > -1)
                v42 := []int{fi40, (v15 + 128), 2147483647}
                _ = fi40
                _ = v41
                _ = v42
            }
            v43 := GenS4{F0: -32768, F1: 349}
            switch ((v4 + (v4 + (v4 + v4))) + ((((("hello" + v4) + "hello") + ("hello" + v4)) + v4) + v4)) {
            case "\n\t":
                if (0.0 < 0.1) {
                    v3 = (!true)
                    for fi44 := 0; fi44 < 3; fi44++ {
                        fmt.Println(v3, v15)
                        fmt.Println(v3)
                        if (-1 < -1) {
                            var v45 bool = true
                            if true {
                                var v46 int8 = -128
                                v46 += (v46 ^ (v46 - 54))
                                _ = v46
                            }
                            v3 = (v45 || v3)
                            _ = v45
                        } else {
                            v5 -= 1
                            switch ((v4 + "") + (v4 + v4)) {
                            case "foo bar":
                                v47 := GenS5{F0: 385, F1: (fi44 / ((v15) | 1))}
                                _ = v47
                            default:
                                var v48 uint8 = 2
                                _ = v48
                            }
                        }
                        _ = fi44
                    }
                } else {
                    for fi49 := 0; fi49 < 3; fi49++ {
                        v50 := []int{fi49, v15}
                        for fi51 := 0; fi51 < 14; fi51++ {
                            var v52 float64 = -2.25
                            fmt.Println(v4)
                            _ = fi51
                            _ = v52
                        }
                        if (!v3) {
                            var v53 int = fi49
                            switch v4 {
                            case "\n\t":
                                v54 := GenS6{F0: 42}
                                _ = v54
                            }
                            _ = v53
                        }
                        _ = fi49
                        _ = v50
                    }
                    fmt.Println(v3, v5, v4)
                    for fi55 := 0; fi55 < 15; fi55++ {
                        v5 -= 2147483648
                        _ = fi55
                    }
                }
            case "foo bar":
                for fi56 := 0; fi56 < 6; fi56++ {
                    if v3 {
                        if (float32(-2.25) >= 1e10) {
                            v3 = ((v5 << ((255) & 63)) <= v5)
                            v5 -= v5
                        } else {
                            var v57 int16 = 42
                            for fi58 := 0; fi58 < 12; fi58++ {
                                if v3 {
                                    var v59 float64 = 1e10
                                    fmt.Println(v3)
                                    var v60 int8 = -94
                                    _ = v59
                                    _ = v60
                                }
                                v15 = 691
                                delete(v2, v4)
                                _ = fi58
                            }
                            for fi61 := 0; fi61 < 14; fi61++ {
                                if v3 {
                                    fmt.Println(v57, v5)
                                    var v62 uint16 = 2
                                    _ = v62
                                }
                                v63 := GenS7{F0: true}
                                _ = fi61
                                _ = v63
                            }
                            _ = v57
                        }
                    } else {
                        v64 := []int{(-(fi56)), -9223372036854775808, v15}
                        _ = v64
                    }
                    _ = fi56
                }
            }
            _ = v43
        }
        var v65 uint = 2
        _ = v65
    } else {
        for fi66 := 0; fi66 < 13; fi66++ {
            for fi67 := 0; fi67 < 10; fi67++ {
                if true {
                    var v68 int = -2147483648
                    _ = v68
                }
                v15 = v15
                _ = fi67
            }
            switch "foo bar" {
            case "":
                fmt.Println(v5, fi66)
            }
            var v69 uint = 4294967295
            _ = fi66
            _ = v69
        }
        _ = len(v2)
    }
    for fi70 := 0; fi70 < 19; fi70++ {
        if true {
            var v71 uint16 = 42
            v72 := v1[42]
            if (733 == 2) {
                if ((((v4 + "") > "\n\t") || ((-2.25 < 1e10) || (-1 > v15))) || v3) {
                    switch 42 {
                    case 42:
                        var v73 int8 = 1
                        _ = v73
                    case -37:
                        var v74 uint32 = v5
                        _ = v74
                    case 32768:
                        if v3 {
                            var v75 int64 = 42
                            v76 := GenS8{F0: fi70}
                            switch v4 {
                            case "\n\t":
                                switch 69 {
                                case -409:
                                    if (627 > 65535) {
                                        for fi77 := 0; fi77 < 6; fi77++ {
                                            v71 = 269
                                            v78 := GenS3{F0: (v4 + v4), F1: (!("\n\t" > (v4 + "a")))}
                                            if false {
                                                v4 = v4
                                            }
                                            if false {
                                                if true {
                                                    v79 := v6[(v5 >> ((65535) & 63))]
                                                    _ = v79
                                                } else {
                                                    v80 := GenS9{F0: -177, F1: (((float64(0.1) + -2.25) < -2.25) && false), F2: (^(v5))}
                                                    v15 -= (fi77 << ((65535) & 63))
                                                    _ = v80
                                                }
                                                var v81 string = v4
                                                _ = v81
                                            }
                                            if (!(255 == 2)) {
                                                switch (v15 << ((uint(v5)) & 63)) {
                                                case -1:
                                                    var v82 float32 = (0.0 * (-2.25 / 1.0))
                                                    v83 := GenS10{F0: true, F1: 0, F2: ((v4 + "hello") + v4)}
                                                    _ = v82
                                                    _ = v83
                                                case 0:
                                                    v15 = (v72 / (((fi70 | (v15 + v72))) | 1))
                                                case 1:
                                                    v84 := GenS11{F0: 1, F1: 127, F2: 170}
                                                    _ = v84
                                                default:
                                                    if ((-1 > 120) || v3) {
                                                        var v85 int32 = -836
                                                        _ = v85
                                                    } else {
                                                        if (v3 && ((v71 & v71) >= v71)) {
                                                            v86 := GenS2{F0: (v75 >> ((65535) & 63)), F1: (-1 <= 153), F2: (v4 + v4)}
                                                            var v87 int32 = 554
                                                            var v88 bool = (!(!((376 >= (v87 << ((2) & 63))) || ((255 > 159) || v3))))
                                                            var v89 float64 = 1e10
                                                            var v90 uint16 = 1
                                                            var v91 bool = (!v88)
                                                            v92 := GenS9{F0: (v87 * (v87 << ((138) & 63))), F1: v91, F2: (^(v5))}
                                                            _ = v86
                                                            _ = v87
                                                            _ = v88
                                                            _ = v89
                                                            _ = v90
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                        v3 = (0 >= 0)
                                                    }
                                                }
                                            }
                                            v75 = (v75 / ((9223372036854775807) | 1))
                                            switch v4 {
                                            case "hello":
                                                switch "hello" {
                                                case "\n\t":
                                                    var v93 string = (v4 + ("\n\t" + v4))
                                                    _ = v93
                                                }
                                            case "a":
                                                var v94 int32 = -567
                                                _ = v94
                                            case "\n\t":
                                                var v95 int8 = -128
                                                _ = v95
                                            }
                                            _ = fi77
                                            _ = v78
                                        }
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            var v96 int = v72
                                            _ = v96
                                        case "foo bar":
                                            switch "a" {
                                            case "foo bar":
                                                var v97 bool = (4294967295 >= 1)
                                                _ = v97
                                            case "hello":
                                                if (370 != 65535) {
                                                    v71 -= 190
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                        }
                                    }
                                }
                            default:
                                var v98 bool = false
                                _ = v98
                            }
                            _ = v75
                            _ = v76
                        }
                    default:
                        fmt.Println(fi70, v3)
                    }
                    v72 = 42
                } else {
                    if v3 {
                        fmt.Println(v4, v4, v3)
                    } else {
                        var v99 float64 = 0.1
                        var v100 float32 = (0.0 - -2.25)
                        _ = len(v2)
                        _ = v100
                        _ = v99
                    }
                }
                fmt.Println(v71)
                v101 := GenS2{F0: -1, F1: (!(-2.25 != (-2.25 + 1e10))), F2: (("hello" + v4) + "hello")}
                _ = v101
            }
            _ = v71
            _ = v72
        } else {
            delete(v1, v5)
            if (!false) {
                v15 -= (^(fi70))
                for fi102 := 0; fi102 < 2; fi102++ {
                    v103 := []int{fi102, 128}
                    var v104 string = ((v4 + (v4 + v4)) + (v4 + (((((v4 + v4) + v4) + ("a" + "\n\t")) + (v4 + "hello")) + v4)))
                    _ = fi102
                    _ = v103
                    _ = v104
                }
                var v105 uint32 = (v5 ^ v5)
                _ = v105
            }
            v4 = (("hello" + "a") + v4)
        }
        if (1 != 761) {
            var v106 uint64 = 18446744073709551615
            fmt.Println(fi70, fi70, v3)
            _ = v106
        }
        _ = fi70
    }
    var v107 uint32 = 576
    v108 := []int{-9223372036854775808, v15}
    var v109 int64 = 1
    delete(v6, v107)
    v110 := GenS12{F0: 2147483647, F1: v15, F2: 164}
    var v111 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v108:", len(v108), v108)
    fmt.Println("v109:", v109)
    fmt.Println("v111:", v111)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v107
    _ = v108
    _ = v109
    _ = v110
    _ = v111
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
