package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 int32; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 2
    var v3 uint8 = 241
    v4 := []int{-2147483648, 32767}
    var v5 string = ""
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(v2, v3, v3)
        v7 := GenS1{F0: 42, F1: 9223372036854775807, F2: 42}
        _ = fi6
        _ = v7
    }
    v5 = v5
    for fi8 := 0; fi8 < 6; fi8++ {
        v2 = (v2 % ((v2) | 1))
        v9 := GenS1{F0: 1, F1: (fi8 - fi8), F2: 373}
        var v10 float64 = -2.25
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int16 = 1
    var v12 int8 = -72
    var v13 int64 = 42
    for fi14 := 0; fi14 < 9; fi14++ {
        v4 = append(v4, (fi14 * int(v2)))
        v13 -= 0
        _ = fi14
    }
    var v15 bool = (true && (2 < v2))
    var v16 uint16 = v2
    var v17 int16 = 0
    v18 := GenS2{F0: (v12 + (v12 | (v12 - -1))), F1: true}
    fmt.Println(v17, v11)
    if (((true && true) && v15) && (187 != 255)) {
        v19 := GenS1{F0: 0, F1: 921, F2: 1}
        _ = v19
    } else {
        v3 -= (v3 + v3)
        var v20 int32 = -2147483648
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        _ = len(v4)
        v1[fi21] = 2147483647
        v17 = 0
        _ = fi21
    }
    var v22 float64 = 3.14159
    if v15 {
        if ((v17 ^ v17) >= (-(v11))) {
            for fi23 := 0; fi23 < 17; fi23++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi24 := 0; fi24 < 1; fi24++ {
                    v25 := v1[(fi24 + fi23)]
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
            switch 42 {
            case 32768:
                var v26 int16 = v11
                _ = v26
            case 91:
                switch 1 {
                case 1:
                    var v27 int16 = 0
                    _ = v27
                case -2147483648:
                    var v28 int32 = -2147483648
                    _ = v28
                }
            default:
                var v29 int = 127
                _ = v29
            }
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 int8 = -128
                _ = fi30
                _ = v31
            }
        } else {
            if (920 != 42) {
                for fi32 := 0; fi32 < 12; fi32++ {
                    if v15 {
                        v12 += (v12 + 0)
                        if (((!(1.5 < v22)) && true) || v15) {
                            for fi33 := 0; fi33 < 12; fi33++ {
                                v17 = -80
                                switch v5 {
                                case "hello":
                                    var v34 bool = v15
                                    _ = v34
                                case "":
                                    switch fi32 {
                                    case 176:
                                        var v35 float32 = 3.14159
                                        _ = v35
                                    case -1:
                                        var v36 string = v5
                                        _ = v36
                                    default:
                                        var v37 bool = v15
                                        if v15 {
                                            fmt.Println(v15)
                                            v17 = (-(v17))
                                            var v38 int16 = (-(v17))
                                            _ = v38
                                        } else {
                                            v22 = v22
                                        }
                                        var v39 uint64 = uint64(v3)
                                        _ = v37
                                        _ = v39
                                    }
                                case "a":
                                    v2 += uint16(v3)
                                default:
                                    v16 += v2
                                }
                                _ = fi33
                            }
                            for fi40 := 0; fi40 < 18; fi40++ {
                                if (v16 < v2) {
                                    fmt.Println(v2)
                                    if true {
                                        var v41 uint32 = 2
                                        _ = v41
                                    } else {
                                        v15 = v15
                                        v42 := GenS1{F0: 285, F1: (fi40 | fi40), F2: 4294967295}
                                        _ = v42
                                    }
                                    var v43 int = (fi40 % (((^(fi40))) | 1))
                                    _ = v43
                                }
                                if len(v4) > 0 {
                                    v4[0] = (fi32 % (((fi40 * (^(fi40)))) | 1))
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = fi40
                            }
                            if ((v11 % (((v17 * v17)) | 1)) < 42) {
                                if ((((false && ((!((-(fi32)) < fi32)) && v15)) && (v2 < 65535)) && v15) && true) {
                                    v15 = (v13 == 42)
                                } else {
                                    var v44 uint = 42
                                    _ = v44
                                }
                            } else {
                                v45 := GenS3{F0: ("\n\t" + v5)}
                                fmt.Println(fi32, v12, v17)
                                _ = v45
                            }
                        } else {
                            if true {
                                if v15 {
                                    fmt.Println(v22, v13)
                                    fmt.Println(v5)
                                    for fi46 := 0; fi46 < 11; fi46++ {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        fmt.Println(v17, fi46)
                                        _ = fi46
                                    }
                                } else {
                                    v11 = (-(v17))
                                    v47 := GenS4{F0: v3}
                                    _ = v47
                                }
                                for fi48 := 0; fi48 < 3; fi48++ {
                                    var v49 bool = (42 != 4294967295)
                                    _ = fi48
                                    _ = v49
                                }
                                for fi50 := 0; fi50 < 1; fi50++ {
                                    fmt.Println(v22)
                                    _ = fi50
                                }
                            } else {
                                switch (fi32 ^ 2147483647) {
                                case -9223372036854775808:
                                    for fi51 := 0; fi51 < 7; fi51++ {
                                        v52 := GenS5{F0: 42, F1: 0, F2: (v2 % ((v16) | 1))}
                                        fmt.Println(v16, v13)
                                        v17 -= 0
                                        switch "" {
                                        case "foo bar":
                                            _ = len(v4)
                                        case "hello":
                                            var v53 float32 = float32(v22)
                                            _ = v53
                                        }
                                        _ = fi51
                                        _ = v52
                                    }
                                default:
                                    var v54 int32 = -490
                                    _ = v54
                                }
                            }
                            for fi55 := 0; fi55 < 10; fi55++ {
                                var v56 int16 = (v17 / ((-207) | 1))
                                v16 = v2
                                v56 = (v11 ^ (v11 << ((255) & 63)))
                                _ = fi55
                                _ = v56
                            }
                        }
                        for fi57 := 0; fi57 < 8; fi57++ {
                            var v58 uint64 = 584
                            switch (fi57 - (fi57 >> ((311) & 63))) {
                            case 14:
                                fmt.Println(v5, v13, fi57)
                            case 386:
                                delete(v1, (-(fi57)))
                            default:
                                if (527 == (v58 | v58)) {
                                    v22 = 0.1
                                    if true {
                                        var v59 int64 = int64(v2)
                                        switch "a" {
                                        case "foo bar":
                                            v60 := GenS2{F0: v12, F1: v15}
                                            _ = v60
                                        case "":
                                            fmt.Println(v58)
                                        case "a":
                                            var v61 int16 = (-(v11))
                                            _ = v61
                                        }
                                        var v62 uint = 0
                                        _ = v59
                                        _ = v62
                                    }
                                }
                            }
                            _ = fi57
                            _ = v58
                        }
                    }
                    switch 451 {
                    case 9223372036854775807:
                        v11 -= 329
                    }
                    _ = fi32
                }
            }
            for fi63 := 0; fi63 < 2; fi63++ {
                var v64 bool = (!true)
                if (true || ((v3 * v3) >= v3)) {
                    v65 := GenS5{F0: -942, F1: 1, F2: v2}
                    for fi66 := 0; fi66 < 16; fi66++ {
                        delete(v1, fi66)
                        switch (fi66 >> ((523) & 63)) {
                        case 128:
                            switch (fi63 >> ((18446744073709551615) & 63)) {
                            case 947:
                                if (false && (-2147483648 != (fi63 / (((fi66 ^ 42)) | 1)))) {
                                    v67 := GenS6{F0: 1}
                                    var v68 uint64 = 133
                                    _ = v67
                                    _ = v68
                                } else {
                                    v64 = true
                                    switch v5 {
                                    case "":
                                        fmt.Println(v12)
                                    case "\n\t":
                                        var v69 bool = ((!true) && true)
                                        _ = v69
                                    case "foo bar":
                                        v2 = 0
                                    default:
                                        var v70 uint8 = v3
                                        _ = v70
                                    }
                                    var v71 bool = (!v64)
                                    _ = v71
                                }
                            default:
                                var v72 float32 = 0.1
                                _ = v72
                            }
                        default:
                            v2 += v16
                        }
                        var v73 int16 = (v17 * (v11 * 42))
                        var v74 int32 = 0
                        v75 := GenS2{F0: (v12 ^ (v12 * (v12 - (v12 - -62)))), F1: v64}
                        v76 := GenS7{F0: false, F1: 0}
                        _ = len(v4)
                        var v77 uint8 = (v3 << ((18446744073709551615) & 63))
                        _ = fi66
                        _ = v73
                        _ = v74
                        _ = v75
                        _ = v76
                        _ = v77
                    }
                    _ = v65
                } else {
                    fmt.Println(v15, v3)
                    if v64 {
                        var v78 int8 = v12
                        var v79 int32 = 61
                        v22 += ((float64(0.1) - v22) + (0.0 / (-2.25 + (3.14159 - v22))))
                        _ = v78
                        _ = v79
                    } else {
                        if (v22 >= (v22 + ((0.1 / (float64(3.14159) * float64(0.1))) + 1e10))) {
                            v80 := GenS1{F0: 793, F1: -526, F2: 18446744073709551615}
                            var v81 uint = 1
                            v3 -= uint8(v16)
                            _ = v80
                            _ = v81
                        }
                        var v82 int32 = 158
                        _ = v82
                    }
                    v2 -= 1
                }
                var v83 int16 = (v11 % ((v11) | 1))
                _ = fi63
                _ = v64
                _ = v83
            }
            switch ((v5 + v5) + v5) {
            case "a":
                v17 += -1
            default:
                switch 42 {
                case 1:
                    var v84 uint64 = 42
                    _ = v84
                case 42:
                    for fi85 := 0; fi85 < 5; fi85++ {
                        fmt.Println(v15, v13, v15)
                        var v86 uint8 = v3
                        v87 := GenS8{F0: "\n\t"}
                        _ = fi85
                        _ = v86
                        _ = v87
                    }
                case 0:
                    switch (v5 + "\n\t") {
                    case "a":
                        var v88 bool = v15
                        _ = v88
                    case "\n\t":
                        switch "" {
                        case "a":
                            v5 = ("\n\t" + v5)
                        case "foo bar":
                            for fi89 := 0; fi89 < 1; fi89++ {
                                _ = len(v4)
                                for fi90 := 0; fi90 < 13; fi90++ {
                                    var v91 int64 = -9223372036854775808
                                    switch (fi90 | (fi90 ^ fi90)) {
                                    case 9223372036854775807:
                                        v92 := v1[42]
                                        _ = v92
                                    case -645:
                                        v22 = 1e10
                                    case 2147483647:
                                        delete(v1, fi89)
                                    default:
                                        fmt.Println(v91, v91)
                                    }
                                    fmt.Println(v3)
                                    _ = fi90
                                    _ = v91
                                }
                                v15 = false
                                _ = fi89
                            }
                        case "hello":
                            var v93 uint = 1
                            _ = v93
                        default:
                            switch (v5 + ((v5 + "a") + v5)) {
                            case "a":
                                if (((1.5 - float32(0.0)) >= 0.0) && v15) {
                                    v94 := GenS9{F0: -1, F1: -1}
                                    for fi95 := 0; fi95 < 4; fi95++ {
                                        v2 -= v2
                                        for fi96 := 0; fi96 < 3; fi96++ {
                                            if v15 {
                                                v15 = false
                                                if v15 {
                                                    v97 := GenS2{F0: 91, F1: false}
                                                    var v98 int64 = (v13 - 0)
                                                    v3 = v3
                                                    _ = v97
                                                    _ = v98
                                                }
                                            }
                                            var v99 string = ((((v5 + v5) + v5) + v5) + ((v5 + v5) + v5))
                                            v100 := GenS10{F0: 1}
                                            _ = fi96
                                            _ = v100
                                            _ = v99
                                        }
                                        _ = fi95
                                    }
                                    _ = v94
                                }
                            case "foo bar":
                                switch 0 {
                                case 9223372036854775807:
                                    var v101 int32 = 0
                                    _ = v101
                                }
                            case "":
                                var v102 string = "a"
                                _ = v102
                            }
                        }
                    case "":
                        v17 += -1
                    default:
                        for fi103 := 0; fi103 < 19; fi103++ {
                            v16 += v16
                            _ = fi103
                        }
                    }
                default:
                    var v104 int16 = v17
                    _ = v104
                }
            }
        }
        v1[274] = 4294967295
        if (v22 < v22) {
            fmt.Println(v13, v11, v3)
            v17 -= 32767
        } else {
            if v15 {
                v15 = v15
                v13 = (v13 - -1)
                v4 = append(v4, -37)
            } else {
                for fi105 := 0; fi105 < 8; fi105++ {
                    v106 := GenS11{F0: 1}
                    v4 = append(v4, (fi105 | (fi105 % ((fi105) | 1))))
                    _ = fi105
                    _ = v106
                }
                fmt.Println(v11, v2)
                var v107 uint32 = 0
                _ = v107
            }
        }
    } else {
        if ((v12 ^ (^(v12))) < (v12 / ((89) | 1))) {
            _ = len(v4)
        } else {
            v11 += (v17 + 42)
            var v108 int64 = 42
            _ = v108
        }
        v109 := GenS12{F0: (^(v12)), F1: 4294967295}
        _ = v109
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
