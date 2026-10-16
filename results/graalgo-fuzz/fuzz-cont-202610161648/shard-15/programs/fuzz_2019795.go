package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint16; F2 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 uint8; F1 bool; F2 int16 }
type GenS9 struct { F0 int64; F1 int32; F2 uint32 }
func genFn0(p0 uint8, p1 uint8, p2 uint) int8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        switch "a" {
        case "hello":
            var v8 string = ("foo bar" + ("" + "hello"))
            _ = v8
        default:
            var v9 bool = false
            _ = v9
        }
        if true {
            var v10 int8 = 42
            v11 := []int{fi7, (fi7 * (fi7 << ((p2) & 63))), fi7}
            if false {
                if false {
                    var v12 uint32 = 508
                    for fi13 := 0; fi13 < 6; fi13++ {
                        var v14 uint64 = 42
                        var v15 uint = (^(p2))
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = v12
                }
                for fi16 := 0; fi16 < 6; fi16++ {
                    if (0.1 < 1.5) {
                        fmt.Println(p2)
                        fmt.Println(fi16)
                        var v17 int32 = 42
                        _ = v17
                    } else {
                        v18 := GenS1{F0: (fi7 / (((fi7 + 207)) | 1)), F1: -796}
                        v19 := GenS2{F0: (p1 >> (((^(p2))) & 63)), F1: (((("a" + (("a" + "") + "foo bar")) + "") + (("foo bar" + ("a" + "hello")) + "\n\t")) + (("a" + "a") + ("\n\t" + "a")))}
                        _ = v18
                        _ = v19
                    }
                    var v20 string = ("hello" + "\n\t")
                    p0 = 2
                    _ = fi16
                    _ = v20
                }
                for fi21 := 0; fi21 < 14; fi21++ {
                    v10 = (v10 & v10)
                    v22 := GenS2{F0: 127, F1: "foo bar"}
                    var v23 int8 = v10
                    switch "\n\t" {
                    case "\n\t":
                        switch "a" {
                        case "a":
                            var v24 uint16 = 42
                            _ = v24
                        }
                    case "a":
                        v25 := GenS3{F0: 0}
                        _ = v25
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            } else {
                for fi26 := 0; fi26 < 16; fi26++ {
                    p2 += 1
                    switch -2147483648 {
                    case -2147483648:
                        p2 = 255
                    case 1:
                        _ = len(v11)
                    }
                    var v27 int64 = -9223372036854775808
                    fmt.Println(p0, p0)
                    _ = fi26
                    _ = v27
                }
                if false {
                    if (((1.5 + (3.14159 * 3.14159)) >= (0.1 - 1e10)) && (0.0 <= 0.0)) {
                        fmt.Println(p2, p0)
                        v28 := GenS1{F0: fi7, F1: -32768}
                        for fi29 := 0; fi29 < 6; fi29++ {
                            var v30 uint32 = 260
                            _ = fi29
                            _ = v30
                        }
                        _ = v28
                    } else {
                        switch (fi7 * fi7) {
                        case 9223372036854775807:
                            fmt.Println(p1, fi7, fi7)
                        case 32768:
                            for fi31 := 0; fi31 < 4; fi31++ {
                                v10 += v10
                                v32 := GenS4{F0: p2, F1: 42, F2: false}
                                _ = fi31
                                _ = v32
                            }
                        default:
                            fmt.Println(p0, p0)
                        }
                    }
                }
            }
            _ = v10
            _ = v11
        }
        _ = fi7
    }
    v33 := GenS1{F0: -131, F1: 32767}
    var v34 uint64 = 275
    v35 := GenS4{F0: p2, F1: 65535, F2: ((!(0 > 901)) && (-20 >= 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v34
    _ = v35
    return -1
}
func genFn1(p0 int8, p1 int32) bool {
    if (p0 >= genFn0(123, 129, 0)) {
        if true {
            var v36 float32 = 3.14159
            _ = v36
        }
    } else {
        p0 -= genFn0(127, 2, 24)
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        p0 -= (p0 << ((1) & 63))
        if (!false) {
            var v38 uint64 = 990
            _ = v38
        }
        var v39 uint8 = 153
        _ = fi37
        _ = v39
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        var v41 int8 = p0
        _ = fi40
        _ = v41
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 uint16 = 1
    v2 := []uint32{}
    var v3 int64 = 9223372036854775807
    var v4 uint16 = 32767
    var v5 string = ("a" + ("foo bar" + ""))
    v6 := []string{("\n\t" + v5), (v5 + ("foo bar" + (("hello" + v5) + (v5 + v5))))}
    _ = len(v2)
    v42 := GenS1{F0: -399, F1: 32767}
    if false {
        v5 = ((v5 + (v5 + ((v5 + (("" + ((v5 + "a") + ("foo bar" + (v5 + v5)))) + ((((v5 + "foo bar") + ("foo bar" + (v5 + v5))) + (("a" + v5) + v5)) + "\n\t"))) + ("" + v5)))) + v5)
        if true {
            switch 9223372036854775807 {
            case -1:
                switch 42 {
                case 1:
                    var v43 float64 = 3.14159
                    _ = v43
                case -1:
                    if (42 < v4) {
                        if true {
                            switch v5 {
                            case "a":
                                var v44 int8 = 1
                                _ = v44
                            case "foo bar":
                                var v45 int64 = (^(v3))
                                _ = v45
                            default:
                                v4 += (v1 & v1)
                            }
                            var v46 float32 = float32((1e10 - (((((0.1 - 1.5) + (0.1 / 1.0)) + 3.14159) - 3.14159) / -2.25)))
                            _ = v46
                        }
                        v47 := GenS5{F0: 723}
                        _ = v47
                    } else {
                        if (1 < 255) {
                            for fi48 := 0; fi48 < 13; fi48++ {
                                v4 = v1
                                var v49 float64 = (0.1 / float64((1.5 * -2.25)))
                                _ = fi48
                                _ = v49
                            }
                            var v50 int32 = 42
                            _ = v50
                        } else {
                            var v51 uint64 = 252
                            _ = v51
                        }
                        var v52 uint64 = 0
                        for fi53 := 0; fi53 < 10; fi53++ {
                            var v54 float32 = 0.0
                            for fi55 := 0; fi55 < 2; fi55++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(v3, v5)
                                v1 = v1
                                _ = fi55
                            }
                            v56 := GenS2{F0: 1, F1: v5}
                            _ = fi53
                            _ = v54
                            _ = v56
                        }
                        _ = v52
                    }
                default:
                    fmt.Println(v5, v1)
                }
            case 32767:
                for fi57 := 0; fi57 < 18; fi57++ {
                    v58 := GenS6{F0: 1, F1: 341, F2: (v1 * v1)}
                    _ = fi57
                    _ = v58
                }
            default:
                var v59 int64 = (v3 - 9223372036854775807)
                _ = v59
            }
        } else {
            _ = len(v6)
            var v60 string = ""
            var v61 string = "\n\t"
            _ = v60
            _ = v61
        }
        for fi62 := 0; fi62 < 19; fi62++ {
            switch fi62 {
            case 2147483647:
                switch fi62 {
                case -1:
                    if true {
                        _ = len(v2)
                    }
                }
            case 127:
                v3 = v3
            case 884:
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
            }
            for fi63 := 0; fi63 < 17; fi63++ {
                v2 = append(v2, 166)
                fmt.Println(v1)
                var v64 uint = 304
                _ = fi63
                _ = v64
            }
            _ = fi62
        }
    } else {
        var v65 string = v5
        if true {
            switch -9223372036854775808 {
            case 323:
                v3 -= v3
            default:
                if true {
                    v1 += 2
                    switch "hello" {
                    case "a":
                        for fi66 := 0; fi66 < 4; fi66++ {
                            var v67 int32 = -741
                            _ = fi66
                            _ = v67
                        }
                    case "":
                        v4 -= v1
                    }
                } else {
                    v6 = append(v6, "hello")
                    for fi68 := 0; fi68 < 12; fi68++ {
                        if (v3 >= v3) {
                            v3 = -1
                            v69 := GenS2{F0: 127, F1: "hello"}
                            if false {
                                if true {
                                    for fi70 := 0; fi70 < 9; fi70++ {
                                        var v71 int = (fi70 - fi68)
                                        fmt.Println(fi70, v5, v65)
                                        var v72 float64 = (1.5 + 0.0)
                                        _ = fi70
                                        _ = v71
                                        _ = v72
                                    }
                                    v4 -= (v4 + v1)
                                }
                                v3 = (v3 & 1)
                                for fi73 := 0; fi73 < 12; fi73++ {
                                    var v74 int32 = 2147483647
                                    fmt.Println(fi68, fi73)
                                    _ = fi73
                                    _ = v74
                                }
                            }
                            _ = v69
                        }
                        v1 += (v4 / ((v4) | 1))
                        for fi75 := 0; fi75 < 2; fi75++ {
                            if (!(((fi68 + 9223372036854775807) < 1) && true)) {
                                fmt.Println(fi75, v65, fi75)
                                v76 := GenS7{F0: 18446744073709551615, F1: 18446744073709551615}
                                var v77 int8 = genFn0(0, 128, 935)
                                _ = v76
                                _ = v77
                            } else {
                                v78 := GenS1{F0: fi75, F1: -313}
                                v65 = v5
                                _ = v78
                            }
                            if false {
                                v1 -= (v4 ^ v1)
                                v79 := GenS8{F0: 128, F1: false, F2: 42}
                                fmt.Println(v5, v1)
                                var v80 uint32 = 2
                                var v81 bool = ((!true) || (!true))
                                v2 = append(v2, v80)
                                _ = v79
                                _ = v80
                                _ = v81
                            }
                            v82 := GenS7{F0: 18446744073709551615, F1: 4294967295}
                            _ = fi75
                            _ = v82
                        }
                        _ = fi68
                    }
                }
            }
            if len(v6) > 0 {
                v6[0] = ("hello" + "a")
            }
            var v83 uint = 18446744073709551615
            _ = v83
        }
        _ = v65
    }
    for fi84 := 0; fi84 < 12; fi84++ {
        switch -1 {
        case 127:
            var v85 float32 = (3.14159 + 0.1)
            _ = v85
        case 0:
            switch (v5 + ("foo bar" + v5)) {
            case "\n\t":
                for fi86 := 0; fi86 < 9; fi86++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v87 := GenS7{F0: 18446744073709551615, F1: 0}
                    var v88 int = fi84
                    v3 -= v3
                    var v89 float32 = 0.1
                    var v90 uint16 = v1
                    _ = fi86
                    _ = v87
                    _ = v88
                    _ = v89
                    _ = v90
                }
            case "hello":
                var v91 float32 = (float32(-2.25) - 0.0)
                _ = v91
            case "a":
                v92 := GenS9{F0: v3, F1: -1, F2: 42}
                _ = v92
            }
        }
        _ = fi84
    }
    if false {
        if (232 >= 169) {
            _ = len(v6)
            var v93 int32 = 2147483647
            _ = v93
        } else {
            if (v4 > (v4 & (v4 * (v1 + 978)))) {
                v3 -= -482
                switch "" {
                case "foo bar":
                    switch v5 {
                    case "":
                        var v94 int32 = int32(v1)
                        _ = v94
                    case "\n\t":
                        for fi95 := 0; fi95 < 13; fi95++ {
                            for fi96 := 0; fi96 < 6; fi96++ {
                                var v97 int = fi95
                                var v98 int16 = -657
                                _ = fi96
                                _ = v97
                                _ = v98
                            }
                            for fi99 := 0; fi99 < 8; fi99++ {
                                switch (v5 + "hello") {
                                case "\n\t":
                                    v3 -= (v3 | (v3 - 9223372036854775807))
                                default:
                                    if false {
                                        v4 += (v4 - v4)
                                        v2 = append(v2, 1)
                                        v5 = "a"
                                    }
                                }
                                _ = fi99
                            }
                            _ = fi95
                        }
                    default:
                        fmt.Println(v5, v4, v3)
                    }
                case "\n\t":
                    v3 += 0
                case "":
                    for fi100 := 0; fi100 < 8; fi100++ {
                        for fi101 := 0; fi101 < 20; fi101++ {
                            v3 -= (v3 ^ v3)
                            _ = fi101
                        }
                        v102 := GenS7{F0: 18446744073709551615, F1: 42}
                        v4 = v1
                        _ = fi100
                        _ = v102
                    }
                }
            }
            v6 = append(v6, v5)
            for fi103 := 0; fi103 < 14; fi103++ {
                v104 := GenS4{F0: 1, F1: v4, F2: false}
                var v105 int8 = -123
                v3 = v3
                _ = fi103
                _ = v104
                _ = v105
            }
        }
    }
    var v106 int16 = 621
    var v107 float32 = 0.0
    if (!false) {
        switch 255 {
        case 0:
            var v108 int = 42
            _ = v108
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
