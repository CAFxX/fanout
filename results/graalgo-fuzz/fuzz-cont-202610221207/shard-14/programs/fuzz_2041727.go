package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint; F1 uint64; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 string; F2 bool }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 uint8; F1 int; F2 uint32 }
type GenS8 struct { F0 bool; F1 uint16; F2 uint32 }
func genFn0(p0 uint32) uint8 {
    v5 := GenS1{F0: -847, F1: 9223372036854775807, F2: 255}
    fmt.Println(p0)
    v6 := GenS1{F0: -2147483648, F1: -1, F2: 1}
    _ = p0
    _ = v5
    _ = v6
    return 0
}
func genFn1() uint8 {
    if false {
        switch ("hello" + ("foo bar" + "foo bar")) {
        case "foo bar":
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := []int{}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi7
                _ = v8
            }
        }
    } else {
        var v9 uint = 0
        v10 := GenS1{F0: -1, F1: -9223372036854775808, F2: 42}
        _ = v10
        _ = v9
    }
    v11 := []int{128, 32767}
    _ = v11
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 float32 = -2.25
    var v12 bool = (2 != 1)
    if (65535 < 0) {
        v4 = float32((((0.0 / float64(-2.25)) - 1.5) + ((-2.25 * -2.25) * ((3.14159 - 0.1) + 0.1))))
        if ((-9223372036854775808 == -9223372036854775808) || (313 < 18446744073709551615)) {
            var v13 int32 = 532
            var v14 string = "\n\t"
            _ = v13
            _ = v14
        } else {
            switch ("\n\t" + "hello") {
            case "\n\t":
                v15 := []int{}
                _ = v15
            case "a":
                switch "a" {
                case "hello":
                    v16 := GenS1{F0: -1, F1: -1, F2: 1}
                    _ = v16
                case "a":
                    fmt.Println(v12, v4)
                case "\n\t":
                    v4 += v4
                }
            case "hello":
                if v12 {
                    var v17 int32 = 0
                    for fi18 := 0; fi18 < 17; fi18++ {
                        v4 = (1.5 * v4)
                        switch "foo bar" {
                        case "\n\t":
                            var v19 uint = 1
                            _ = v19
                        case "hello":
                            var v20 int16 = 42
                            _ = v20
                        case "":
                            v17 = (v17 << ((42) & 63))
                        }
                        if v12 {
                            v12 = (530 != 2147483647)
                            if false {
                                var v21 uint = 65535
                                var v22 int = fi18
                                _ = v21
                                _ = v22
                            }
                        }
                        _ = fi18
                    }
                    _ = v17
                } else {
                    if (((801 <= 381) && ((-9223372036854775808 <= 355) || (42 >= -32768))) && ((!(2 >= 2)) || false)) {
                        switch "hello" {
                        case "a":
                            fmt.Println(v4, v4)
                        case "hello":
                            var v23 uint = 255
                            _ = v23
                        }
                    } else {
                        v24 := []int{0}
                        var v25 float64 = float64((v4 / (v4 * -2.25)))
                        if v12 {
                            var v26 uint = 198
                            var v27 uint8 = genFn0(65535)
                            if (255 == 1) {
                                v28 := GenS2{F0: -2147483648, F1: v12}
                                _ = v28
                            }
                            _ = v26
                            _ = v27
                        } else {
                            for fi29 := 0; fi29 < 3; fi29++ {
                                if (32768 <= fi29) {
                                    for fi30 := 0; fi30 < 16; fi30++ {
                                        v31 := GenS2{F0: -55, F1: v12}
                                        switch "a" {
                                        case "\n\t":
                                            if false {
                                                if v12 {
                                                    v3[42] = 814
                                                    var v32 uint8 = 159
                                                    v25 -= (((v25 - (3.14159 / v25)) + (((float64(((float32(v25) - (-2.25 * v4)) / 0.1)) + (v25 * v25)) / v25) - ((1.5 * float64(3.14159)) * v25))) / v25)
                                                    _ = v32
                                                }
                                                var v33 uint64 = 1
                                                v34 := GenS2{F0: 990, F1: (!(-1 >= int32(fi29)))}
                                                var v35 uint = 1
                                                _ = len(v24)
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            } else {
                                                if (974 > 42) {
                                                    if false {
                                                        fmt.Println(v12, fi29)
                                                    }
                                                } else {
                                                    v36 := v2[135]
                                                    var v37 int = (-(v36))
                                                    _ = v36
                                                    _ = v37
                                                }
                                                if len(v24) >= 2 {
                                                    v24 = v24[1:len(v24)]
                                                }
                                                v4 = float32(v25)
                                            }
                                        case "foo bar":
                                            v25 -= (float64(((3.14159 / (v4 / v4)) / 1e10)) * v25)
                                        case "hello":
                                            _ = len(v24)
                                        }
                                        _ = fi30
                                        _ = v31
                                    }
                                    fmt.Println(v4, v4, v12)
                                } else {
                                    for fi38 := 0; fi38 < 3; fi38++ {
                                        switch 2147483647 {
                                        case 42:
                                            var v39 int = (fi38 & 503)
                                            _ = v39
                                        }
                                        var v40 int16 = -32768
                                        var v41 float32 = v4
                                        _ = fi38
                                        _ = v40
                                        _ = v41
                                    }
                                    if false {
                                        for fi42 := 0; fi42 < 5; fi42++ {
                                            if v12 {
                                                v43 := GenS2{F0: 467, F1: false}
                                                fmt.Println(v25)
                                                _ = v43
                                            }
                                            var v44 int32 = 1
                                            var v45 float64 = v25
                                            _ = fi42
                                            _ = v44
                                            _ = v45
                                        }
                                        v46 := GenS2{F0: 770, F1: (42 == -9223372036854775808)}
                                        for fi47 := 0; fi47 < 11; fi47++ {
                                            _ = len(v24)
                                            _ = fi47
                                        }
                                        _ = v46
                                    }
                                    for fi48 := 0; fi48 < 3; fi48++ {
                                        fmt.Println(fi48, v25)
                                        if false {
                                            fmt.Println(fi48, v25)
                                            v12 = v12
                                            v4 -= v4
                                        } else {
                                            fmt.Println(v12, fi48, v25)
                                            fmt.Println(v4)
                                        }
                                        if (1 < (fi48 + fi29)) {
                                            v4 += -2.25
                                        } else {
                                            if true {
                                                v4 -= v4
                                                if v12 {
                                                    var v49 int16 = 0
                                                    var v50 uint64 = 1
                                                    _ = v49
                                                    _ = v50
                                                }
                                                var v51 float64 = v25
                                                _ = v51
                                            }
                                            v25 -= 3.14159
                                        }
                                        _ = len(v24)
                                        if (v12 || true) {
                                            v52 := GenS3{F0: 805, F1: 4294967295, F2: true}
                                            if v12 {
                                                if true {
                                                    switch 128 {
                                                    case 389:
                                                        v25 += ((0.0 - v25) + 0.0)
                                                    default:
                                                        v25 = v25
                                                    }
                                                    fmt.Println(fi29, v25)
                                                } else {
                                                    var v53 uint8 = 62
                                                    var v54 int8 = 1
                                                    _ = v53
                                                    _ = v54
                                                }
                                            }
                                            _ = v52
                                        } else {
                                            fmt.Println(v4)
                                            _ = len(v24)
                                        }
                                        _ = fi48
                                    }
                                }
                                var v55 int = 128
                                var v56 int16 = 0
                                _ = fi29
                                _ = v55
                                _ = v56
                            }
                            var v57 uint16 = 42
                            switch ("a" + "hello") {
                            case "foo bar":
                                v25 = -2.25
                            case "a":
                                var v58 float64 = v25
                                _ = v58
                            }
                            _ = v57
                        }
                        _ = v24
                        _ = v25
                    }
                }
            }
            switch -1 {
            case 241:
                v59 := GenS3{F0: 609, F1: 528, F2: v12}
                _ = v59
            case 128:
                fmt.Println(v12)
            }
            for fi60 := 0; fi60 < 3; fi60++ {
                if (!false) {
                    v12 = (v12 && (float32(float64(((((((v4 * (v4 - v4)) / (((v4 / (v4 * v4)) - (0.1 - (((v4 - (v4 - ((v4 + (0.1 - v4)) * (((v4 / (1.5 - -2.25)) * v4) / -2.25)))) + (3.14159 + v4)) + -2.25))) - float32(3.14159))) * v4) / v4) / (v4 + v4)) - v4))) != 0.1))
                    if v12 {
                        v4 = float32((3.14159 / 1e10))
                        if ((v12 || (!(float32(0.1) == (float32(0.0) - ((float32((0.0 * -2.25)) - ((3.14159 - v4) - v4)) + float32((float64((v4 + (3.14159 / 1.0))) * (3.14159 - 3.14159)))))))) || (42 <= 9223372036854775807)) {
                            if (!false) {
                                var v61 int8 = 17
                                for fi62 := 0; fi62 < 5; fi62++ {
                                    v63 := []int{(fi62 ^ (fi62 + -912))}
                                    v12 = v12
                                    var v64 int8 = v61
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                }
                                _ = v61
                            } else {
                                v12 = ((127 > 127) && v12)
                            }
                            v65 := GenS4{F0: "a"}
                            v66 := GenS1{F0: -1, F1: -679, F2: 22}
                            _ = v65
                            _ = v66
                        }
                    } else {
                        fmt.Println(v12)
                        v4 = v4
                        v4 += (float32((1e10 - 1.5)) - 3.14159)
                    }
                    for fi67 := 0; fi67 < 13; fi67++ {
                        v4 = 3.14159
                        var v68 int32 = 42
                        switch ("a" + ("a" + "a")) {
                        case "hello":
                            var v69 float32 = v4
                            _ = v69
                        }
                        v70 := []int{fi60, 2147483647, fi67}
                        fmt.Println(v12, v12, fi67)
                        _ = fi67
                        _ = v68
                        _ = v70
                    }
                }
                fmt.Println(fi60, fi60, v12)
                _ = fi60
            }
        }
    } else {
        v71 := GenS1{F0: 641, F1: -645, F2: 128}
        var v72 uint8 = 2
        _ = v71
        _ = v72
    }
    var v73 bool = v12
    for fi74 := 0; fi74 < 4; fi74++ {
        var v75 float32 = v4
        v76 := GenS5{F0: fi74, F1: "foo bar", F2: true}
        var v77 uint16 = 0
        _ = fi74
        _ = v75
        _ = v76
        _ = v77
    }
    var v78 int8 = 1
    v79 := []int{127, 868, 4294967295}
    if len(v79) >= 2 {
        v79 = v79[1:len(v79)]
    }
    _ = len(v3)
    v12 = false
    for fi80 := 0; fi80 < 18; fi80++ {
        if ((true && ((true || v73) && v73)) && (100 < 1)) {
            var v81 int64 = 1
            if (!(!v73)) {
                if v73 {
                    switch "" {
                    case "":
                        if len(v79) >= 2 {
                            v79 = v79[1:len(v79)]
                        }
                    default:
                        switch fi80 {
                        case 32768:
                            if len(v79) > 0 {
                                v79[0] = fi80
                            }
                        case 32767:
                            _ = len(v79)
                        case 1:
                            if ((1 != 142) && v73) {
                                switch fi80 {
                                case 0:
                                    if v12 {
                                        switch fi80 {
                                        case -2147483648:
                                            switch ("" + ("" + "\n\t")) {
                                            case "":
                                                v12 = (v81 >= (v81 ^ v81))
                                            }
                                        case 9223372036854775807:
                                            _ = len(v3)
                                        }
                                    }
                                default:
                                    if v12 {
                                        var v82 uint16 = 42
                                        v81 += v81
                                        _ = v82
                                    }
                                }
                                var v83 uint16 = 65535
                                _ = v83
                            }
                        default:
                            if v12 {
                                if ((842 > 630) || false) {
                                    for fi84 := 0; fi84 < 3; fi84++ {
                                        v12 = true
                                        v12 = ((-586 != 475) && true)
                                        switch (^(fi80)) {
                                        case 518:
                                            if (-2.25 >= (0.0 - 0.0)) {
                                                if v12 {
                                                    v79 = append(v79, fi84)
                                                } else {
                                                    v81 += -830
                                                    var v85 int16 = -32768
                                                    switch "" {
                                                    case "hello":
                                                        if v73 {
                                                            if ("\n\t" <= "\n\t") {
                                                                var v86 uint32 = 1
                                                                _ = v86
                                                            }
                                                            var v87 uint32 = 4294967295
                                                            var v88 int = 42
                                                            _ = v87
                                                            _ = v88
                                                        } else {
                                                            switch 127 {
                                                            case 4294967295:
                                                                if (v4 == 0.1) {
                                                                    var v89 float64 = 0.0
                                                                    _ = v89
                                                                }
                                                            }
                                                            v85 -= 826
                                                        }
                                                    case "a":
                                                        v85 = (v85 >> ((1) & 63))
                                                    default:
                                                        v90 := GenS3{F0: 972, F1: 357, F2: v12}
                                                        _ = v90
                                                    }
                                                    _ = v85
                                                }
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v4 = v4
                                                case "":
                                                    fmt.Println(fi80, fi80)
                                                default:
                                                    v78 = -1
                                                }
                                            }
                                        case 1:
                                            v1[(^(fi84))] = 2147483647
                                        }
                                        _ = fi84
                                    }
                                    if true {
                                        v12 = (float64(((((v4 - 1.5) + v4) - ((v4 + float32(1.5)) - -2.25)) * (float32(0.1) / v4))) <= 0.0)
                                        v81 = (v81 - 0)
                                        v4 = v4
                                    } else {
                                        v91 := GenS2{F0: int32(v78), F1: v12}
                                        _ = v91
                                    }
                                } else {
                                    var v92 float64 = (3.14159 + (0.0 - 0.0))
                                    if (!false) {
                                        fmt.Println(v12, v12)
                                    }
                                    v4 -= v4
                                    _ = v92
                                }
                                if (v73 || (v73 || (2147483647 >= 0))) {
                                    v4 = v4
                                    v93 := v2[fi80]
                                    var v94 uint64 = 4294967295
                                    _ = v93
                                    _ = v94
                                } else {
                                    var v95 uint16 = 342
                                    var v96 string = ("" + ("" + "a"))
                                    var v97 uint32 = 2
                                    _ = v95
                                    _ = v96
                                    _ = v97
                                }
                            }
                        }
                    }
                    _ = len(v1)
                } else {
                    v98 := GenS6{F0: v81, F1: 147}
                    _ = v98
                }
            } else {
                v4 += (-2.25 / 0.1)
                if ((!v12) || v73) {
                    v2[fi80] = (fi80 >> ((153) & 63))
                    switch fi80 {
                    case 9223372036854775807:
                        for fi99 := 0; fi99 < 4; fi99++ {
                            if (false || (v73 || true)) {
                                fmt.Println(v12, v4, v73)
                                switch "a" {
                                case "a":
                                    v100 := GenS2{F0: 0, F1: v73}
                                    var v101 int = -1
                                    fmt.Println(fi99, v81)
                                    _ = v100
                                    _ = v101
                                case "\n\t":
                                    _ = len(v2)
                                case "hello":
                                    v4 -= 1e10
                                }
                                var v102 int32 = -1
                                _ = len(v2)
                                var v103 bool = (v73 || true)
                                _ = v102
                                _ = v103
                            }
                            fmt.Println(v4, v78)
                            var v104 uint16 = 42
                            _ = fi99
                            _ = v104
                        }
                    default:
                        v105 := GenS6{F0: v81, F1: genFn1()}
                        _ = v105
                    }
                    var v106 uint64 = 141
                    _ = v106
                } else {
                    switch fi80 {
                    case 2147483647:
                        v81 = (v81 * (v81 & 1))
                    case 1:
                        var v107 uint64 = 2
                        _ = v107
                    }
                    var v108 int32 = 0
                    var v109 uint32 = 2147483647
                    _ = v108
                    _ = v109
                }
                fmt.Println(v4, v4, v78)
            }
            _ = v81
        } else {
            v110 := GenS7{F0: 42, F1: 47, F2: 65535}
            _ = v110
        }
        v111 := GenS1{F0: int32(v78), F1: 1, F2: genFn1()}
        _ = fi80
        _ = v111
    }
    v112 := GenS8{F0: v12, F1: 929, F2: 2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v73:", v73)
    fmt.Println("v78:", v78)
    fmt.Println("slicelen v79:", len(v79), v79)
    _ = v1
    _ = v112
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v73
    _ = v78
    _ = v79
}
