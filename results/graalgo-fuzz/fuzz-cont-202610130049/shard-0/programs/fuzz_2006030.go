package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint8; F2 string }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint8; F1 uint32; F2 string }
func genFn0(p0 int64, p1 int, p2 uint8) uint16 {
    v7 := []int{(p1 % ((32767) | 1)), p1}
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint16 = 37
            v10 = v10
            var v11 int32 = -2147483648
            v11 = 971
            var v12 uint64 = 18446744073709551615
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        if true {
            var v13 uint = uint(p2)
            var v14 uint = 255
            var v15 int8 = -128
            _ = v13
            _ = v14
            _ = v15
        } else {
            p0 = (p0 / ((963) | 1))
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 int8 = 67
    var v2 int8 = v1
    var v3 bool = false
    var v4 uint32 = 191
    v5 := make(map[int]bool)
    var v6 int = -1
    v3 = (2 != 18446744073709551615)
    v16 := []int{0, (v6 / ((v6) | 1)), -865}
    var v17 float64 = (((float64((-2.25 - 0.0)) * 3.14159) / 1e10) - -2.25)
    if v3 {
        for fi18 := 0; fi18 < 4; fi18++ {
            if (!false) {
                _ = len(v16)
                v3 = v3
            } else {
                if (899 > -32768) {
                    fmt.Println(v6, fi18, v2)
                }
            }
            for fi19 := 0; fi19 < 16; fi19++ {
                v20 := GenS1{F0: fi18}
                v3 = false
                delete(v5, (fi19 >> ((1) & 63)))
                var v21 float32 = 1e10
                if (((18446744073709551615 >= 42) || v3) && v3) {
                    v21 -= 0.0
                    switch "" {
                    case "\n\t":
                        if (65535 <= 18446744073709551615) {
                            var v22 string = "hello"
                            _ = v22
                        } else {
                            switch 32768 {
                            case -15:
                                if (!(255 != 2)) {
                                    var v23 string = "foo bar"
                                    if v3 {
                                        v23 = "\n\t"
                                    }
                                    if v3 {
                                        v23 = v23
                                        if v3 {
                                            var v24 int16 = 0
                                            if (!(!false)) {
                                                if v3 {
                                                    if false {
                                                        v25 := GenS1{F0: 944}
                                                        _ = v25
                                                    }
                                                    var v26 float32 = v21
                                                    _ = v26
                                                } else {
                                                    fmt.Println(fi19)
                                                    v23 = v23
                                                    if (0.0 >= float64((v21 + (v21 / (1e10 / v21))))) {
                                                        if v3 {
                                                            var v27 uint32 = (^(v4))
                                                            var v28 int64 = 42
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                        var v29 int32 = -1
                                                        if false {
                                                            switch -1 {
                                                            case 42:
                                                                switch v23 {
                                                                case "a":
                                                                    switch v23 {
                                                                    case "":
                                                                        if (v23 < (v23 + ("" + (((v23 + (v23 + ("foo bar" + "a"))) + v23) + v23)))) {
                                                                            fmt.Println(v3)
                                                                            var v30 int32 = (v29 - v29)
                                                                            var v31 uint64 = 1
                                                                            fmt.Println(v17, v24, v6)
                                                                            v23 = v23
                                                                            _ = v30
                                                                            _ = v31
                                                                        }
                                                                    }
                                                                default:
                                                                    if (-78 < v1) {
                                                                        var v32 int8 = v2
                                                                        v33 := v5[v6]
                                                                        if true {
                                                                            var v34 uint8 = 127
                                                                            switch (v6 * 32767) {
                                                                            case 9223372036854775807:
                                                                                var v35 int16 = -1
                                                                                _ = v35
                                                                            case 32767:
                                                                                switch (v23 + v23) {
                                                                                case "\n\t":
                                                                                    v24 += (v24 << ((99) & 63))
                                                                                }
                                                                                v36 := GenS2{F0: (^(v29)), F1: 2147483647}
                                                                                _ = v36
                                                                            case -1:
                                                                                var v37 int16 = 965
                                                                                _ = v37
                                                                            default:
                                                                                var v38 int64 = -1
                                                                                _ = v38
                                                                            }
                                                                            v39 := GenS3{F0: 1}
                                                                            v40 := GenS1{F0: v6}
                                                                            v1 -= (v1 * (v2 & v2))
                                                                            fmt.Println(v1, v3, v6)
                                                                            fmt.Println(v34)
                                                                            if (1.5 != (v21 / -2.25)) {
                                                                                switch v6 {
                                                                                case -1:
                                                                                    var v41 int = 42
                                                                                    _ = v41
                                                                                case -9223372036854775808:
                                                                                    if (127 == -38) {
                                                                                        fmt.Println(v21, v21, v34)
                                                                                        if v3 {
                                                                                            var v42 bool = true
                                                                                            v4 += (v4 ^ 199)
                                                                                            v43 := GenS4{F0: (v6 | fi19), F1: (v34 % (((v34 * v34)) | 1)), F2: ("foo bar" + (((v23 + (v23 + (v23 + v23))) + (v23 + (v23 + v23))) + (v23 + "foo bar")))}
                                                                                            _ = v42
                                                                                            _ = v43
                                                                                        } else {
                                                                                            var v44 string = "a"
                                                                                            if (((v23 + "foo bar") == "") && v3) {
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    var v45 int32 = (v29 - (v29 >> ((18446744073709551615) & 63)))
                                                                                                    _ = v45
                                                                                                case "":
                                                                                                    if v3 {
                                                                                                        switch (fi19 ^ 9223372036854775807) {
                                                                                                        case -63:
                                                                                                            v46 := GenS5{F0: (v24 & (v24 + 0)), F1: 1}
                                                                                                            _ = v46
                                                                                                        case -9223372036854775808:
                                                                                                            var v47 uint = 242
                                                                                                            _ = v47
                                                                                                        case 2147483647:
                                                                                                            _ = len(v5)
                                                                                                        }
                                                                                                        v2 = v1
                                                                                                        var v48 uint64 = 775
                                                                                                        _ = v48
                                                                                                    } else {
                                                                                                        if (0.0 > v17) {
                                                                                                            v34 += 127
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v49 int32 = (-(v29))
                                                                                                if v33 {
                                                                                                    v32 = (v1 ^ v2)
                                                                                                }
                                                                                                _ = v49
                                                                                            }
                                                                                            v50 := GenS5{F0: v24, F1: 9223372036854775807}
                                                                                            _ = v44
                                                                                            _ = v50
                                                                                        }
                                                                                        v51 := GenS4{F0: fi19, F1: (v34 - (v34 | 255)), F2: v23}
                                                                                        _ = v51
                                                                                    } else {
                                                                                        if ((771 == 0) && true) {
                                                                                            switch (v23 + v23) {
                                                                                            case "a":
                                                                                                v21 = -2.25
                                                                                            case "\n\t":
                                                                                                fmt.Println(fi19)
                                                                                            default:
                                                                                                if true {
                                                                                                    fmt.Println(v29, v24)
                                                                                                    var v52 int8 = v32
                                                                                                    _ = v52
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v29, v17, v3)
                                                                                        }
                                                                                        var v53 float32 = (v21 - ((((0.1 * v21) / (((v21 + 0.1) - (v21 * (((v21 / v21) / 0.1) * (v21 + v21)))) * v21)) * v21) - 0.1))
                                                                                        if (fi19 < (fi18 - fi19)) {
                                                                                            if len(v16) > 0 {
                                                                                                v16[0] = fi18
                                                                                            }
                                                                                            if (2 == 0) {
                                                                                                v1 += -1
                                                                                                switch "hello" {
                                                                                                case "hello":
                                                                                                    switch "" {
                                                                                                    case "\n\t":
                                                                                                        v17 -= (1e10 + v17)
                                                                                                    case "hello":
                                                                                                        if ((v29 % ((0) | 1)) > v29) {
                                                                                                            var v54 int = (fi18 ^ 128)
                                                                                                            v6 -= (fi18 * (fi19 ^ fi18))
                                                                                                            fmt.Println(v2, v2)
                                                                                                            _ = v54
                                                                                                        } else {
                                                                                                            v4 = (v4 >> ((65535) & 63))
                                                                                                        }
                                                                                                    case "a":
                                                                                                        var v55 bool = (v4 >= 2)
                                                                                                        _ = v55
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                v56 := GenS1{F0: 9223372036854775807}
                                                                                                var v57 int16 = (v24 - (v24 % ((1) | 1)))
                                                                                                _ = v56
                                                                                                _ = v57
                                                                                            }
                                                                                        }
                                                                                        _ = v53
                                                                                    }
                                                                                case 127:
                                                                                    var v58 uint64 = 516
                                                                                    _ = v58
                                                                                }
                                                                            }
                                                                            _ = v34
                                                                            _ = v39
                                                                            _ = v40
                                                                        } else {
                                                                            if (!v33) {
                                                                                fmt.Println(v29, fi19)
                                                                                if len(v16) > 0 {
                                                                                    v16[0] = (v6 % ((v6) | 1))
                                                                                }
                                                                                v59 := v5[(fi18 << ((534) & 63))]
                                                                                fmt.Println(v2)
                                                                                if ((!false) || ((v4 % ((v4) | 1)) <= (v4 | (^(v4))))) {
                                                                                    v23 = ""
                                                                                    v16 = append(v16, (fi19 << ((0) & 63)))
                                                                                    switch (v23 + v23) {
                                                                                    case "":
                                                                                        var v60 int16 = -1
                                                                                        _ = v60
                                                                                    }
                                                                                    v1 -= (v1 ^ v2)
                                                                                }
                                                                                _ = v59
                                                                            } else {
                                                                                var v61 bool = (v29 > (v29 % ((-2147483648) | 1)))
                                                                                var v62 uint8 = 24
                                                                                _ = v61
                                                                                _ = v62
                                                                            }
                                                                            v21 += 0.0
                                                                        }
                                                                        if v3 {
                                                                            v3 = false
                                                                            fmt.Println(v1, fi19, v3)
                                                                        }
                                                                        _ = v32
                                                                        _ = v33
                                                                    } else {
                                                                        fmt.Println(v23, v29, v3)
                                                                    }
                                                                }
                                                            case -1:
                                                                var v63 int64 = 1
                                                                _ = v63
                                                            default:
                                                                if false {
                                                                    if v3 {
                                                                        if (((1e10 / v17) / v17) == v17) {
                                                                            if v3 {
                                                                                var v64 int16 = (v24 % ((v24) | 1))
                                                                                if (0 < v64) {
                                                                                    fmt.Println(v4, v29, v4)
                                                                                    fmt.Println(v3, fi19, v4)
                                                                                    if len(v16) > 0 {
                                                                                        v16[0] = 9223372036854775807
                                                                                    }
                                                                                } else {
                                                                                    var v65 string = v23
                                                                                    v66 := GenS3{F0: -9223372036854775808}
                                                                                    switch v23 {
                                                                                    case "foo bar":
                                                                                        var v67 uint64 = 1
                                                                                        _ = v67
                                                                                    case "a":
                                                                                        var v68 uint32 = v4
                                                                                        v69 := GenS6{F0: (v24 % ((-32768) | 1)), F1: 32767}
                                                                                        if false {
                                                                                            fmt.Println(v4)
                                                                                            var v70 uint64 = 0
                                                                                            _ = v70
                                                                                        } else {
                                                                                            if ((fi19 + fi18) <= fi18) {
                                                                                                _ = len(v5)
                                                                                                fmt.Println(v3, v65)
                                                                                            }
                                                                                            v68 = 374
                                                                                        }
                                                                                        _ = v68
                                                                                        _ = v69
                                                                                    case "hello":
                                                                                        if false {
                                                                                            if false {
                                                                                                if false {
                                                                                                    fmt.Println(v21, v4, v21)
                                                                                                    v29 -= 0
                                                                                                    fmt.Println(v2)
                                                                                                    var v71 float32 = 0.1
                                                                                                    if true {
                                                                                                        v72 := v5[fi18]
                                                                                                        _ = v72
                                                                                                    }
                                                                                                    _ = v71
                                                                                                }
                                                                                            } else {
                                                                                                if (!(true || true)) {
                                                                                                    if v3 {
                                                                                                        if (((v23 + v23) + "hello") <= v65) {
                                                                                                            switch "" {
                                                                                                            case "hello":
                                                                                                                var v73 float64 = v17
                                                                                                                _ = v73
                                                                                                            case "a":
                                                                                                                fmt.Println(v64, v17)
                                                                                                            case "foo bar":
                                                                                                                var v74 uint8 = 249
                                                                                                                _ = v74
                                                                                                            default:
                                                                                                                v75 := GenS6{F0: (v24 + v24), F1: (v24 - v24)}
                                                                                                                _ = v75
                                                                                                            }
                                                                                                            if (!v3) {
                                                                                                                v65 = v65
                                                                                                                v65 = v23
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    v76 := GenS7{F0: (fi19 | 42)}
                                                                                                    if (v3 && true) {
                                                                                                        v3 = (!v3)
                                                                                                    }
                                                                                                    _ = v76
                                                                                                }
                                                                                                var v77 float32 = (v21 / v21)
                                                                                                _ = v77
                                                                                            }
                                                                                            fmt.Println(v6, v1, v2)
                                                                                            fmt.Println(v4, v4)
                                                                                        } else {
                                                                                            v4 = 42
                                                                                            fmt.Println(v24, v3, v4)
                                                                                        }
                                                                                    }
                                                                                    _ = v65
                                                                                    _ = v66
                                                                                }
                                                                                if len(v16) >= 2 {
                                                                                    v16 = v16[1:len(v16)]
                                                                                }
                                                                                _ = v64
                                                                            }
                                                                            fmt.Println(v24)
                                                                            v23 = (v23 + v23)
                                                                            v78 := GenS4{F0: (-(fi19)), F1: 2, F2: v23}
                                                                            _ = v78
                                                                        } else {
                                                                            switch (v23 + "foo bar") {
                                                                            case "a":
                                                                                var v79 int32 = 2147483647
                                                                                _ = v79
                                                                            }
                                                                        }
                                                                        switch fi19 {
                                                                        case -848:
                                                                            delete(v5, fi19)
                                                                        case -801:
                                                                            var v80 bool = ((float64(v21) - 0.0) != (0.1 / v17))
                                                                            _ = v80
                                                                        default:
                                                                            fmt.Println(v2, v23)
                                                                        }
                                                                        v81 := GenS8{F0: 42, F1: (v2 - (v2 + v2))}
                                                                        _ = v81
                                                                    }
                                                                } else {
                                                                    var v82 string = ("" + v23)
                                                                    switch (fi19 ^ v6) {
                                                                    case 650:
                                                                        if (!(40 < 101)) {
                                                                            v23 = ""
                                                                        } else {
                                                                            var v83 float32 = v21
                                                                            if len(v16) > 0 {
                                                                                v16[0] = (fi19 << ((65535) & 63))
                                                                            }
                                                                            var v84 float64 = v17
                                                                            _ = v83
                                                                            _ = v84
                                                                        }
                                                                    case -9223372036854775808:
                                                                        var v85 int32 = v29
                                                                        _ = v85
                                                                    }
                                                                    _ = v82
                                                                }
                                                            }
                                                            fmt.Println(v21, fi19)
                                                            switch (fi18 >> ((1) & 63)) {
                                                            case 4294967295:
                                                                if v3 {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        _ = len(v5)
                                                                    case "":
                                                                        if v3 {
                                                                            fmt.Println(v6, v23)
                                                                            if v3 {
                                                                                var v86 int = fi18
                                                                                fmt.Println(fi18, fi19, v3)
                                                                                v17 = (v17 * v17)
                                                                                if len(v16) >= 2 {
                                                                                    v16 = v16[1:len(v16)]
                                                                                }
                                                                                if ((false || v3) || v3) {
                                                                                    if (9 < int64(fi19)) {
                                                                                        v6 = (-(v86))
                                                                                        v87 := GenS1{F0: 9223372036854775807}
                                                                                        var v88 uint8 = 127
                                                                                        _ = v87
                                                                                        _ = v88
                                                                                    } else {
                                                                                        if true {
                                                                                            var v89 bool = v3
                                                                                            var v90 float32 = float32((((v17 - (v17 + (v17 * ((v17 / -2.25) / ((float64(v21) * v17) + v17))))) * v17) / v17))
                                                                                            _ = v89
                                                                                            _ = v90
                                                                                        } else {
                                                                                            v86 = (fi18 + fi18)
                                                                                            if v3 {
                                                                                                if len(v16) >= 2 {
                                                                                                    v16 = v16[1:len(v16)]
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v91 int16 = -32768
                                                                                        _ = v91
                                                                                    }
                                                                                    v1 = v1
                                                                                } else {
                                                                                    v16 = append(v16, fi19)
                                                                                }
                                                                                _ = v86
                                                                            }
                                                                            v92 := GenS1{F0: (^(fi19))}
                                                                            v5[-2147483648] = v3
                                                                            var v93 int16 = (^(v24))
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                    case "\n\t":
                                                                        var v94 int8 = v1
                                                                        _ = v94
                                                                    default:
                                                                        if v3 {
                                                                            if true {
                                                                                var v95 uint = 42
                                                                                if (!(!(!(33 == 0)))) {
                                                                                    if v3 {
                                                                                        delete(v5, (fi19 ^ 42))
                                                                                    } else {
                                                                                        var v96 int = (^(fi19))
                                                                                        if false {
                                                                                            v97 := GenS4{F0: (fi18 | v96), F1: 42, F2: (((((v23 + "\n\t") + (v23 + ("" + v23))) + v23) + ("hello" + (("\n\t" + v23) + (v23 + "foo bar")))) + (v23 + v23))}
                                                                                            _ = len(v5)
                                                                                            v17 = v17
                                                                                            _ = v97
                                                                                        } else {
                                                                                            if (-1 == 810) {
                                                                                                v98 := GenS6{F0: 992, F1: 32767}
                                                                                                v24 = 1000
                                                                                                _ = v98
                                                                                            } else {
                                                                                                switch 0 {
                                                                                                case 128:
                                                                                                    if false {
                                                                                                        v16 = append(v16, 42)
                                                                                                        switch v96 {
                                                                                                        case 1:
                                                                                                            v2 = (v1 + v2)
                                                                                                        }
                                                                                                    } else {
                                                                                                        fmt.Println(v24)
                                                                                                    }
                                                                                                    delete(v5, 0)
                                                                                                case -2147483648:
                                                                                                    v3 = false
                                                                                                default:
                                                                                                    switch "\n\t" {
                                                                                                    case "a":
                                                                                                        var v99 bool = (v3 || ((v95 >= v95) && v3))
                                                                                                        _ = v99
                                                                                                    }
                                                                                                }
                                                                                                if false {
                                                                                                    if false {
                                                                                                        v29 += v29
                                                                                                        switch "foo bar" {
                                                                                                        case "a":
                                                                                                            if len(v16) >= 2 {
                                                                                                                v16 = v16[1:len(v16)]
                                                                                                            }
                                                                                                        case "":
                                                                                                            switch "\n\t" {
                                                                                                            case "":
                                                                                                                if ((float32(v17) * v21) != (v21 - v21)) {
                                                                                                                    v21 = v21
                                                                                                                    if v3 {
                                                                                                                        var v100 uint = v95
                                                                                                                        _ = v100
                                                                                                                    }
                                                                                                                    fmt.Println(v17, v23, v29)
                                                                                                                } else {
                                                                                                                    v5[fi18] = (!v3)
                                                                                                                }
                                                                                                            default:
                                                                                                                if ((-(v24)) < -1) {
                                                                                                                    v6 = 192
                                                                                                                    v5[fi19] = v3
                                                                                                                }
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            if (uint16(v4) <= genFn0(42, (v96 | v6), 127)) {
                                                                                                                switch 42 {
                                                                                                                case 2147483647:
                                                                                                                    fmt.Println(v23)
                                                                                                                default:
                                                                                                                    fmt.Println(fi19, v6)
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v101 int32 = (v29 / ((int32(v4)) | 1))
                                                                                                                v23 = "\n\t"
                                                                                                                _ = v101
                                                                                                            }
                                                                                                        default:
                                                                                                            var v102 float32 = float32((float64((v21 / v21)) / 3.14159))
                                                                                                            _ = v102
                                                                                                        }
                                                                                                        var v103 int8 = v1
                                                                                                        _ = v103
                                                                                                    } else {
                                                                                                        var v104 int16 = v24
                                                                                                        v5[fi18] = v3
                                                                                                        _ = v104
                                                                                                    }
                                                                                                } else {
                                                                                                    v4 = v4
                                                                                                    delete(v5, v96)
                                                                                                }
                                                                                                v105 := v5[v96]
                                                                                                _ = v105
                                                                                            }
                                                                                        }
                                                                                        switch v23 {
                                                                                        case "a":
                                                                                            if v3 {
                                                                                                var v106 bool = (v29 <= v29)
                                                                                                switch "a" {
                                                                                                case "a":
                                                                                                    var v107 int32 = (v29 & v29)
                                                                                                    var v108 int = fi19
                                                                                                    switch (v96 & (-(v108))) {
                                                                                                    case -2147483648:
                                                                                                        switch v108 {
                                                                                                        case 32767:
                                                                                                            v4 = v4
                                                                                                        case -9223372036854775808:
                                                                                                            v16 = append(v16, (-(fi19)))
                                                                                                        default:
                                                                                                            var v109 bool = v3
                                                                                                            _ = v109
                                                                                                        }
                                                                                                        switch 1 {
                                                                                                        case 9223372036854775807:
                                                                                                            if ((42 != 644) && v106) {
                                                                                                                _ = len(v5)
                                                                                                                if v3 {
                                                                                                                    v110 := GenS3{F0: -642}
                                                                                                                    if v106 {
                                                                                                                        if (138 < 145) {
                                                                                                                            var v111 uint32 = (v4 % (((v4 + 927)) | 1))
                                                                                                                            _ = len(v16)
                                                                                                                            if v3 {
                                                                                                                                fmt.Println(v95, v1)
                                                                                                                            }
                                                                                                                            _ = v111
                                                                                                                        }
                                                                                                                        if len(v16) >= 2 {
                                                                                                                            v16 = v16[1:len(v16)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v110
                                                                                                                }
                                                                                                            }
                                                                                                        case -2147483648:
                                                                                                            v5[(fi19 ^ v108)] = (v3 || (4294967295 != 481))
                                                                                                            v24 = (v24 | v24)
                                                                                                            if true {
                                                                                                                var v112 int32 = v107
                                                                                                                _ = v112
                                                                                                            } else {
                                                                                                                v113 := v5[fi18]
                                                                                                                v106 = v113
                                                                                                                if len(v16) > 0 {
                                                                                                                    v16[0] = 622
                                                                                                                }
                                                                                                                if ("\n\t" == "a") {
                                                                                                                    fmt.Println(v6)
                                                                                                                } else {
                                                                                                                    delete(v5, (fi18 * (v6 & 4294967295)))
                                                                                                                }
                                                                                                                fmt.Println(v23, v106)
                                                                                                                _ = v113
                                                                                                            }
                                                                                                        }
                                                                                                        var v114 float64 = ((v17 - v17) - ((float64(float32(3.14159)) * v17) + ((((1e10 / v17) - ((v17 / v17) - float64(((float32(v17) / v21) - float32(v17))))) / v17) * (v17 + -2.25))))
                                                                                                        v115 := GenS8{F0: 255, F1: 0}
                                                                                                        var v116 bool = ((v4 + v4) == v4)
                                                                                                        v117 := GenS2{F0: v29, F1: v29}
                                                                                                        _ = v114
                                                                                                        _ = v115
                                                                                                        _ = v116
                                                                                                        _ = v117
                                                                                                    case 173:
                                                                                                        v3 = (v24 == (v24 - 0))
                                                                                                    case 0:
                                                                                                        switch "hello" {
                                                                                                        case "\n\t":
                                                                                                            delete(v5, v96)
                                                                                                            if true {
                                                                                                                v29 = v107
                                                                                                                _ = len(v5)
                                                                                                            } else {
                                                                                                                v21 += v21
                                                                                                            }
                                                                                                        default:
                                                                                                            _ = len(v5)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v107
                                                                                                    _ = v108
                                                                                                case "\n\t":
                                                                                                    v21 -= v21
                                                                                                case "hello":
                                                                                                    v29 += (v29 | (^(v29)))
                                                                                                default:
                                                                                                    switch 32767 {
                                                                                                    case 9223372036854775807:
                                                                                                        v3 = (!true)
                                                                                                    case 0:
                                                                                                        var v118 int8 = v2
                                                                                                        _ = v118
                                                                                                    }
                                                                                                }
                                                                                                _ = v106
                                                                                            } else {
                                                                                                if (((v29 / ((v29) | 1)) > v29) || true) {
                                                                                                    v119 := v5[(fi18 / ((v6) | 1))]
                                                                                                    v120 := GenS1{F0: fi19}
                                                                                                    fmt.Println(v17, v119, v119)
                                                                                                    if ((18446744073709551615 < 1) && v3) {
                                                                                                        v17 += (v17 - ((v17 * v17) - ((float64((((1e10 / v21) - (v21 * (v21 / v21))) / (3.14159 - v21))) / (v17 * float64(v21))) + 0.0)))
                                                                                                        fmt.Println(v24, v96, v29)
                                                                                                    }
                                                                                                    var v121 int8 = (v2 * (-(v1)))
                                                                                                    _ = v119
                                                                                                    _ = v120
                                                                                                    _ = v121
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            switch fi18 {
                                                                                            case -459:
                                                                                                if ((v29 - v29) >= -2147483648) {
                                                                                                    fmt.Println(v24)
                                                                                                } else {
                                                                                                    v122 := v5[(fi18 ^ (fi18 & (fi18 / ((fi19) | 1))))]
                                                                                                    fmt.Println(v17, v6)
                                                                                                    v17 = float64(((v21 * (float32(v17) / ((v21 * v21) / 3.14159))) / 1.5))
                                                                                                    _ = v122
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v96
                                                                                    }
                                                                                    fmt.Println(v29, v95)
                                                                                    fmt.Println(fi19, fi18, v29)
                                                                                }
                                                                                fmt.Println(v17)
                                                                                _ = v95
                                                                            } else {
                                                                                _ = len(v5)
                                                                                if (42 >= 42) {
                                                                                    var v123 int32 = -2147483648
                                                                                    v6 += v6
                                                                                    _ = v123
                                                                                }
                                                                            }
                                                                            if len(v16) > 0 {
                                                                                v16[0] = -1
                                                                            }
                                                                        } else {
                                                                            v16 = append(v16, fi18)
                                                                            var v124 int64 = -1
                                                                            var v125 float64 = v17
                                                                            v126 := GenS3{F0: v124}
                                                                            _ = v124
                                                                            _ = v125
                                                                            _ = v126
                                                                        }
                                                                    }
                                                                    v1 -= 127
                                                                    var v127 int8 = v2
                                                                    _ = v127
                                                                }
                                                            }
                                                        }
                                                        _ = v29
                                                    } else {
                                                        var v128 float32 = v21
                                                        _ = v128
                                                    }
                                                }
                                            } else {
                                                fmt.Println(fi18, v6)
                                                var v129 uint16 = genFn0(9223372036854775807, -9223372036854775808, 99)
                                                _ = v129
                                            }
                                            switch (("a" + (((v23 + "\n\t") + v23) + (v23 + ""))) + v23) {
                                            case "foo bar":
                                                v130 := GenS6{F0: -32768, F1: v24}
                                                _ = v130
                                            case "hello":
                                                switch 1 {
                                                case -237:
                                                    if v3 {
                                                        var v131 bool = false
                                                        v21 -= (v21 / v21)
                                                        var v132 uint16 = 906
                                                        _ = v131
                                                        _ = v132
                                                    } else {
                                                        var v133 int8 = (v2 - v2)
                                                        _ = v133
                                                    }
                                                case 133:
                                                    var v134 uint64 = 0
                                                    _ = v134
                                                }
                                            default:
                                                var v135 int = 641
                                                _ = v135
                                            }
                                            fmt.Println(v3, v21, v17)
                                            v23 = v23
                                            _ = v24
                                        } else {
                                            if v3 {
                                                if (32767 < 99) {
                                                    v2 = 1
                                                    v2 = v1
                                                } else {
                                                    var v136 uint32 = v4
                                                    _ = v136
                                                }
                                                v137 := GenS1{F0: fi18}
                                                v3 = (!(v3 || ((v21 + v21) > v21)))
                                                _ = v137
                                            }
                                            v17 += (0.1 + 0.1)
                                            fmt.Println(v17, v21, v21)
                                        }
                                    } else {
                                        var v138 int32 = 423
                                        v5[42] = true
                                        v2 -= -1
                                        v17 += ((float64(v21) * (v17 - v17)) / (v17 * float64(v21)))
                                        _ = v138
                                    }
                                    _ = v23
                                }
                            case 128:
                                var v139 uint16 = 32767
                                _ = v139
                            case 32768:
                                switch "\n\t" {
                                case "a":
                                    v3 = (v3 && (228 > 894))
                                case "foo bar":
                                    var v140 int = (fi18 << ((255) & 63))
                                    _ = v140
                                case "hello":
                                    v2 = v1
                                }
                            }
                        }
                    case "":
                        var v141 uint32 = (v4 / ((1) | 1))
                        if (42 < 0) {
                            v6 = fi18
                        }
                        _ = v141
                    default:
                        fmt.Println(v21)
                    }
                    if len(v16) > 0 {
                        v16[0] = fi18
                    }
                    v6 -= fi18
                    if false {
                        var v142 uint32 = (v4 | v4)
                        switch 179 {
                        case 0:
                            v3 = v3
                            fmt.Println(fi19)
                        }
                        _ = v142
                    } else {
                        v17 = (float64(v21) + ((v17 + 1.5) / (v17 / v17)))
                        if v3 {
                            v5[v6] = v3
                            v1 += (v2 % (((v1 - 1)) | 1))
                            if v3 {
                                var v143 uint32 = 414
                                _ = v143
                            }
                            v2 = v2
                        }
                        var v144 int32 = -848
                        _ = v144
                    }
                }
                var v145 bool = (!(!true))
                var v146 string = ("" + "a")
                _ = fi19
                _ = v145
                _ = v146
                _ = v20
                _ = v21
            }
            switch -9223372036854775808 {
            case -530:
                v4 += v4
            case 128:
                if len(v16) > 0 {
                    v16[0] = 4294967295
                }
            }
            _ = fi18
        }
        v3 = (v17 < (3.14159 / 1.0))
    } else {
        fmt.Println(v6, v6)
    }
    v147 := GenS9{F0: 255, F1: v4, F2: "foo bar"}
    fmt.Println(v3)
    for fi148 := 0; fi148 < 9; fi148++ {
        fmt.Println(v1)
        _ = fi148
    }
    var v149 uint = 2
    var v150 float64 = v17
    fmt.Println(v2)
    switch ("" + "hello") {
    case "a":
        switch "foo bar" {
        case "a":
            _ = len(v5)
        }
    case "":
        v17 = (v150 - (v150 * 3.14159))
    default:
        if (1 >= genFn0(1, (v6 | (v6 << ((255) & 63))), 128)) {
            var v151 uint32 = (v4 << (((v149 ^ v149)) & 63))
            _ = v151
        }
    }
    v150 -= v150
    fmt.Println("v1:", v1)
    fmt.Println("v149:", v149)
    fmt.Println("v150:", v150)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v147
    _ = v149
    _ = v150
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
