package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 bool; F1 uint16; F2 int8 }
type GenS3 struct { F0 int8; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint; F2 int }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS9 struct { F0 uint; F1 int8; F2 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS13 struct { F0 int16; F1 uint16 }
type GenS14 struct { F0 int16 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 bool = true
    var v3 uint = v1
    fmt.Println(v3, v2)
    fmt.Println(v3)
    if (!v2) {
        if false {
            var v4 uint16 = 933
            _ = v4
        }
        v5 := GenS1{F0: v3, F1: 1}
        _ = v5
    }
    switch 714 {
    case 128:
        v2 = ("foo bar" < "")
    case -351:
        var v6 float32 = 3.14159
        _ = v6
    case 1:
        switch "foo bar" {
        case "hello":
            if (-1 <= -1) {
                v7 := []int{9223372036854775807, 412, 127, 651}
                v3 -= (v3 / ((18446744073709551615) | 1))
                _ = v7
            }
        case "a":
            v2 = v2
        case "\n\t":
            v8 := GenS1{F0: (v3 | (^(v1))), F1: 32767}
            _ = v8
        default:
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := GenS2{F0: (1 <= 656), F1: 0, F2: 127}
                for fi11 := 0; fi11 < 12; fi11++ {
                    if ((3.14159 / (-2.25 + -2.25)) >= 1.5) {
                        v12 := GenS3{F0: 1, F1: 42, F2: 2147483647}
                        v13 := []int{-869, (fi11 * fi9), 42, 9223372036854775807}
                        if false {
                            switch (fi11 % (((fi11 | (fi9 - fi9))) | 1)) {
                            case 9223372036854775807:
                                if v2 {
                                    if len(v13) > 0 {
                                        v13[0] = fi9
                                    }
                                    var v14 float32 = 3.14159
                                    _ = v14
                                } else {
                                    v15 := GenS4{F0: 1}
                                    _ = v15
                                }
                            case 127:
                                v16 := GenS1{F0: (v1 - (v1 + v1)), F1: (fi9 * -567)}
                                _ = v16
                            case 32768:
                                v1 += (v1 << ((18446744073709551615) & 63))
                            default:
                                v17 := GenS1{F0: 1, F1: fi11}
                                _ = v17
                            }
                            var v18 uint8 = 127
                            v19 := GenS5{F0: -610, F1: 2, F2: (-(fi9))}
                            _ = v18
                            _ = v19
                        } else {
                            fmt.Println(fi11, v1)
                            v13 = append(v13, (fi11 + (fi9 + fi9)))
                        }
                        switch 0 {
                        case 42:
                            if len(v13) > 0 {
                                v13[0] = (fi11 + (fi9 / ((361) | 1)))
                            }
                            switch "hello" {
                            case "a":
                                var v20 uint32 = 4294967295
                                _ = v20
                            default:
                                v21 := GenS6{F0: 128, F1: -1}
                                _ = v21
                            }
                        case -2147483648:
                            fmt.Println(fi11)
                        case -390:
                            switch "hello" {
                            case "":
                                v2 = ((float64((1.5 * (1e10 - ((((0.0 + 0.1) / 1.0) * ((0.0 + 1e10) / (3.14159 * 1.5))) + -2.25)))) <= (0.1 + (float64((float32(0.1) * 0.0)) * 0.0))) || (0 >= 255))
                            case "\n\t":
                                fmt.Println(v2, v3, fi9)
                            case "foo bar":
                                fmt.Println(v3)
                            }
                        }
                        v1 = (^(v3))
                        _ = v12
                        _ = v13
                    } else {
                        fmt.Println(v2, v2, fi9)
                        fmt.Println(v1)
                        if true {
                            if (!v2) {
                                var v22 float64 = 1e10
                                fmt.Println(v22, fi9)
                                _ = v22
                            }
                            v1 -= v3
                            v2 = v2
                        }
                    }
                    var v23 int64 = 102
                    if (255 > 1) {
                        v23 = (^(v23))
                        var v24 uint64 = 388
                        _ = v24
                    }
                    _ = fi11
                    _ = v23
                }
                if true {
                    v25 := GenS1{F0: (^(v1)), F1: 786}
                    _ = v25
                } else {
                    switch "a" {
                    case "":
                        var v26 float64 = float64(((-2.25 + float32((((1.5 * float64(0.0)) + 1e10) - 1e10))) - (3.14159 * 0.0)))
                        _ = v26
                    default:
                        var v27 uint = v3
                        _ = v27
                    }
                    var v28 int32 = -465
                    _ = v28
                }
                _ = fi9
                _ = v10
            }
        }
    default:
        var v29 float32 = 0.1
        _ = v29
    }
    v30 := []int{-2147483648}
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 int = fi31
        switch (fi31 << ((855) & 63)) {
        case -953:
            if ((float32(-2.25) + (0.1 / 0.1)) < (3.14159 - 3.14159)) {
                if true {
                    switch (-(v32)) {
                    case 42:
                        var v33 uint64 = 2
                        _ = v33
                    case 525:
                        var v34 float64 = -2.25
                        _ = v34
                    case 32768:
                        v35 := GenS7{F0: v3}
                        _ = v35
                    }
                }
                var v36 uint8 = 244
                v36 = 255
                _ = v36
            }
        default:
            for fi37 := 0; fi37 < 4; fi37++ {
                v30 = append(v30, 32768)
                v1 = v3
                _ = fi37
            }
        }
        if (20 < 99) {
            if (4294967295 == 73) {
                v38 := GenS5{F0: 194, F1: (v3 - v3), F2: fi31}
                _ = v38
            }
            v2 = ((-2.25 >= ((0.0 * 1.5) - 3.14159)) && v2)
        } else {
            v30 = append(v30, -288)
        }
        _ = fi31
        _ = v32
    }
    switch 565 {
    case 32767:
        fmt.Println(v3, v2, v2)
    default:
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
    }
    v39 := GenS8{F0: 0, F1: 255, F2: 42}
    v2 = v2
    switch -505 {
    case 4294967295:
        var v40 uint64 = 2
        _ = v40
    case -9223372036854775808:
        var v41 int64 = -1
        _ = v41
    default:
        switch -1 {
        case 9223372036854775807:
            var v42 int = 4294967295
            _ = v42
        case 1:
            v43 := GenS9{F0: (v3 & 18446744073709551615), F1: 42, F2: 0}
            _ = v43
        default:
            for fi44 := 0; fi44 < 19; fi44++ {
                fmt.Println(v2)
                _ = fi44
            }
        }
    }
    var v45 int8 = 0
    var v46 int32 = 0
    var v47 int8 = 3
    v2 = ((v3 | 4294967295) == v3)
    var v48 uint = v1
    if v2 {
        var v49 uint8 = 177
        v46 -= v46
        _ = v49
    }
    if (v2 && (((-9223372036854775808 < -9223372036854775808) && v2) && v2)) {
        switch 128 {
        case 32767:
            if v2 {
                _ = len(v30)
                v47 = (v45 << ((197) & 63))
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
            } else {
                if (32767 > 1000) {
                    if (v2 || v2) {
                        v50 := GenS10{F0: 4294967295}
                        v51 := GenS10{F0: -391}
                        _ = v50
                        _ = v51
                    } else {
                        switch 2147483647 {
                        case -1:
                            fmt.Println(v45)
                        }
                        if (!(170 <= 246)) {
                            v3 += 65535
                        }
                        var v52 int64 = -1
                        _ = v52
                    }
                    if (v1 > (v48 | 18446744073709551615)) {
                        fmt.Println(v45, v47, v2)
                    }
                    for fi53 := 0; fi53 < 11; fi53++ {
                        v54 := GenS8{F0: -134, F1: 127, F2: 137}
                        for fi55 := 0; fi55 < 3; fi55++ {
                            v48 = (v48 | (v48 | 18446744073709551615))
                            if v2 {
                                if len(v30) > 0 {
                                    v30[0] = fi55
                                }
                                v2 = v2
                            } else {
                                v3 -= (v1 & v3)
                                v48 = v48
                            }
                            var v56 string = "\n\t"
                            _ = fi55
                            _ = v56
                        }
                        if ((!((!v2) && (!((v46 <= v46) || v2)))) || true) {
                            if false {
                                var v57 uint16 = 0
                                _ = v57
                            } else {
                                var v58 string = ("a" + "")
                                _ = v58
                            }
                            fmt.Println(v45, v45)
                            switch fi53 {
                            case -966:
                                switch ("foo bar" + ("a" + "a")) {
                                case "hello":
                                    v59 := GenS9{F0: (^(v3)), F1: int8(v3), F2: -620}
                                    _ = v59
                                default:
                                    for fi60 := 0; fi60 < 11; fi60++ {
                                        var v61 string = "\n\t"
                                        _ = fi60
                                        _ = v61
                                    }
                                }
                            case 128:
                                switch "hello" {
                                case "\n\t":
                                    v62 := GenS11{F0: (-(fi53))}
                                    _ = v62
                                default:
                                    if v2 {
                                        fmt.Println(v1)
                                        var v63 int8 = (v45 & v45)
                                        if ((42 <= 2) || (!true)) {
                                            var v64 uint8 = 9
                                            _ = v64
                                        } else {
                                            v65 := GenS10{F0: -2147483648}
                                            _ = v65
                                        }
                                        _ = v63
                                    } else {
                                        if v2 {
                                            v30 = append(v30, fi53)
                                        }
                                    }
                                }
                            case 147:
                                v45 += v47
                            default:
                                for fi66 := 0; fi66 < 13; fi66++ {
                                    var v67 int8 = v45
                                    if (!(!v2)) {
                                        var v68 uint64 = 18446744073709551615
                                        _ = v68
                                    }
                                    var v69 int64 = 0
                                    _ = fi66
                                    _ = v67
                                    _ = v69
                                }
                            }
                        }
                        _ = fi53
                        _ = v54
                    }
                } else {
                    for fi70 := 0; fi70 < 20; fi70++ {
                        v71 := GenS8{F0: 1, F1: 42, F2: 0}
                        _ = fi70
                        _ = v71
                    }
                    var v72 float64 = (float64(((1e10 / 1e10) * (3.14159 * float32((0.0 * 3.14159))))) / 1.0)
                    if v2 {
                        v45 -= v47
                        if (!v2) {
                            if (!((^(v46)) < 46)) {
                                switch 32768 {
                                case 585:
                                    var v73 int32 = int32(v1)
                                    _ = v73
                                case 861:
                                    v74 := GenS11{F0: 4294967295}
                                    _ = v74
                                case 32768:
                                    if len(v30) > 0 {
                                        v30[0] = -1
                                    }
                                default:
                                    v30 = append(v30, 155)
                                }
                                v2 = (v47 > int8(v1))
                                v30 = append(v30, -531)
                            } else {
                                _ = len(v30)
                                v72 += float64(3.14159)
                            }
                        } else {
                            v46 = -454
                            v30 = append(v30, 0)
                        }
                    } else {
                        switch -9223372036854775808 {
                        case -282:
                            fmt.Println(v72)
                        case 1:
                            var v75 int64 = 9223372036854775807
                            _ = v75
                        case 4294967295:
                            if len(v30) > 0 {
                                v30[0] = -9223372036854775808
                            }
                        }
                        fmt.Println(v3, v45)
                    }
                    _ = v72
                }
                fmt.Println(v2, v3)
                v76 := GenS5{F0: -731, F1: (v1 | 2), F2: 2147483647}
                _ = v76
            }
        }
    } else {
        fmt.Println(v3)
        for fi77 := 0; fi77 < 4; fi77++ {
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            for fi78 := 0; fi78 < 16; fi78++ {
                switch "foo bar" {
                case "":
                    if (v1 <= (v48 - (v1 & v48))) {
                        v79 := GenS12{F0: (v45 * (v45 % (((v45 + (^(v45)))) | 1))), F1: 824, F2: 2}
                        fmt.Println(v47)
                        _ = v79
                    }
                case "hello":
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                default:
                    v80 := GenS11{F0: (fi78 >> (((v48 << (((^(v3))) & 63))) & 63))}
                    _ = v80
                }
                _ = len(v30)
                switch fi78 {
                case 9223372036854775807:
                    v81 := GenS13{F0: -1, F1: 0}
                    v46 = 255
                    v46 = 217
                    _ = v81
                case 455:
                    v82 := GenS14{F0: 631}
                    _ = v82
                case 4294967295:
                    switch "a" {
                    case "\n\t":
                        var v83 uint32 = 230
                        _ = v83
                    case "":
                        v84 := GenS6{F0: 228, F1: 0}
                        _ = v84
                    default:
                        var v85 int = fi78
                        _ = v85
                    }
                }
                _ = fi78
            }
            v48 += 265
            _ = fi77
        }
        v3 -= (v48 ^ 703)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v39
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
