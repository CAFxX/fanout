package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint32; F1 int32; F2 int64 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 int; F1 uint32; F2 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64 }
func genFn0() int16 {
    v4 := []int{127, 1}
    var v5 uint8 = 1
    v5 -= v5
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(fi6, v5)
        _ = fi6
    }
    switch "hello" {
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        v7 := GenS1{F0: 127}
        _ = v7
    default:
        var v8 uint32 = 2147483647
        _ = v8
    }
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 uint = 25
    v2 := []int{42, 1}
    v3 := []string{((("hello" + "\n\t") + ("\n\t" + (("foo bar" + "foo bar") + "\n\t"))) + "a"), ("hello" + (("hello" + "\n\t") + "")), (("" + "\n\t") + ""), "foo bar"}
    for fi9 := 0; fi9 < 18; fi9++ {
        if false {
            v10 := GenS1{F0: 0}
            v1 += (v1 & (v1 | 846))
            _ = v10
        } else {
            v1 += v1
            fmt.Println(fi9, fi9)
        }
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = fi9
    }
    var v11 int16 = 1
    v12 := GenS1{F0: 0}
    v2 = append(v2, -1)
    var v13 bool = false
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(fi14, v1, v13)
        for fi15 := 0; fi15 < 5; fi15++ {
            if false {
                if ((!v13) && (v13 && v13)) {
                    var v16 string = ""
                    _ = v16
                } else {
                    if true {
                        v17 := GenS1{F0: 255}
                        v18 := GenS1{F0: 2}
                        _ = len(v2)
                        _ = v17
                        _ = v18
                    } else {
                        if v13 {
                            switch "" {
                            case "hello":
                                if (v13 || false) {
                                    v19 := GenS2{F0: 1, F1: 946}
                                    v1 += v1
                                    switch "\n\t" {
                                    case "":
                                        _ = len(v3)
                                        if (!true) {
                                            fmt.Println(v1, fi14, fi14)
                                            v1 = 1
                                        } else {
                                            var v20 float32 = 0.0
                                            _ = v20
                                        }
                                        v1 -= v1
                                    default:
                                        fmt.Println(v13)
                                        if (-9223372036854775808 != -782) {
                                            switch ("" + ("" + "foo bar")) {
                                            case "foo bar":
                                                _ = len(v2)
                                            case "hello":
                                                v2 = append(v2, fi15)
                                            default:
                                                v11 = (v11 << ((v1) & 63))
                                            }
                                            fmt.Println(v1)
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                fmt.Println(v1)
                                            case "hello":
                                                if v13 {
                                                    if (43 == 77) {
                                                        switch fi14 {
                                                        case 335:
                                                            v21 := GenS1{F0: 255}
                                                            _ = v21
                                                        }
                                                        var v22 int = (fi14 ^ fi15)
                                                        _ = v22
                                                    }
                                                    v11 = int16(v1)
                                                    if (true || (!(2147483647 >= 42))) {
                                                        v23 := GenS3{F0: v11, F1: 65535, F2: (v11 | (v11 / (((v11 & (v11 - genFn0()))) | 1)))}
                                                        v1 = 0
                                                        var v24 bool = v13
                                                        _ = v23
                                                        _ = v24
                                                    } else {
                                                        var v25 uint = 1
                                                        v25 = (^(v1))
                                                        v1 = v25
                                                        _ = v25
                                                    }
                                                }
                                            default:
                                                var v26 uint32 = 2147483647
                                                _ = v26
                                            }
                                            if v13 {
                                                v11 = (v11 * v11)
                                                v27 := GenS3{F0: v11, F1: 42, F2: genFn0()}
                                                _ = v27
                                            } else {
                                                var v28 bool = false
                                                _ = v28
                                            }
                                            v29 := GenS4{F0: 255, F1: v1}
                                            _ = v29
                                        }
                                    }
                                    _ = v19
                                } else {
                                    v1 = (v1 + v1)
                                    if (255 < 0) {
                                        fmt.Println(fi15, fi15)
                                    }
                                    v11 -= -309
                                }
                            case "\n\t":
                                var v30 int = (fi15 + fi15)
                                _ = v30
                            }
                            fmt.Println(v11, v13)
                        } else {
                            v3 = append(v3, "foo bar")
                            if v13 {
                                if (v13 && (v13 && (42 == 255))) {
                                    var v31 float32 = (0.0 - 3.14159)
                                    _ = v31
                                }
                                var v32 int32 = 42
                                v32 = 1
                                _ = v32
                            } else {
                                var v33 uint32 = 952
                                _ = v33
                            }
                        }
                        v11 = v11
                    }
                }
            } else {
                v13 = v13
                if (-208 <= 1) {
                    v11 -= v11
                    var v34 float64 = 1e10
                    switch (("" + ("foo bar" + "a")) + ("\n\t" + "a")) {
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                    case "":
                        var v35 uint = v1
                        _ = v35
                    case "hello":
                        v1 = (v1 ^ (v1 | (^(v1))))
                    }
                    _ = v34
                }
            }
            if v13 {
                switch 2147483647 {
                case 1:
                    if false {
                        v13 = v13
                        if false {
                            v36 := GenS4{F0: 128, F1: (v1 ^ v1)}
                            _ = v36
                        }
                    } else {
                        if v13 {
                            v37 := GenS2{F0: 32767, F1: 0}
                            v2 = append(v2, (fi14 + fi15))
                            _ = v37
                        } else {
                            v1 -= (v1 % ((1) | 1))
                            if v13 {
                                if (v13 && v13) {
                                    v38 := GenS5{F0: 2, F1: 1, F2: -1}
                                    _ = v38
                                } else {
                                    v11 -= (v11 ^ 32767)
                                    if (127 > 61) {
                                        fmt.Println(v11, v11, v11)
                                        var v39 string = ("a" + "foo bar")
                                        var v40 uint32 = 58
                                        _ = v39
                                        _ = v40
                                    } else {
                                        var v41 int32 = -354
                                        var v42 uint32 = 845
                                        v43 := GenS5{F0: v42, F1: -432, F2: 9223372036854775807}
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    }
                                }
                                switch 0 {
                                case 1:
                                    var v44 uint16 = 65535
                                    v13 = (int16(fi15) == v11)
                                    _ = v44
                                case 4294967295:
                                    v45 := GenS6{F0: 42, F1: v13}
                                    _ = v45
                                case 32768:
                                    var v46 int64 = 1
                                    _ = v46
                                }
                                v11 -= (v11 ^ 42)
                            }
                            var v47 int16 = (-(v11))
                            _ = v47
                        }
                        if (!(!v13)) {
                            v2 = append(v2, fi15)
                            switch (fi14 / ((fi14) | 1)) {
                            case -597:
                                v13 = (!v13)
                            default:
                                var v48 int8 = 42
                                _ = v48
                            }
                        } else {
                            switch (fi15 / (((fi14 + fi15)) | 1)) {
                            case -2147483648:
                                var v49 int32 = -218
                                _ = v49
                            case 9223372036854775807:
                                v50 := GenS5{F0: 2147483647, F1: 1, F2: 42}
                                _ = v50
                            case 42:
                                v51 := GenS2{F0: 42, F1: 351}
                                _ = v51
                            }
                            if false {
                                switch (("hello" + "") + ("\n\t" + (("" + "a") + "a"))) {
                                case "hello":
                                    if (3.14159 >= 1.5) {
                                        if (1 >= 77) {
                                            v13 = v13
                                        } else {
                                            switch "hello" {
                                            case "hello":
                                                var v52 int64 = -1
                                                _ = v52
                                            case "foo bar":
                                                v53 := GenS5{F0: 304, F1: -1, F2: -9223372036854775808}
                                                _ = v53
                                            case "\n\t":
                                                v2 = append(v2, fi14)
                                            default:
                                                fmt.Println(v11)
                                            }
                                            var v54 float64 = ((0.1 * (1e10 / 1e10)) * (0.1 - ((((1.5 / 3.14159) + float64((1e10 + 3.14159))) - 0.1) - float64(1.5))))
                                            v11 += (v11 + (v11 | genFn0()))
                                            _ = v54
                                        }
                                        v11 -= 1
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        v55 := GenS5{F0: 220, F1: -2147483648, F2: 1}
                                        v2 = append(v2, fi14)
                                        var v56 int = (fi14 + (fi15 / ((fi15) | 1)))
                                        _ = v55
                                        _ = v56
                                    }
                                case "":
                                    v57 := GenS5{F0: 1, F1: 2147483647, F2: 950}
                                    _ = v57
                                case "foo bar":
                                    var v58 int64 = 9223372036854775807
                                    _ = v58
                                }
                                var v59 string = "foo bar"
                                switch fi14 {
                                case 9223372036854775807:
                                    _ = len(v3)
                                }
                                _ = v59
                            } else {
                                var v60 int8 = 42
                                var v61 uint64 = 2
                                _ = v60
                                _ = v61
                            }
                            fmt.Println(fi14)
                        }
                    }
                    if v13 {
                        fmt.Println(v13)
                        var v62 int64 = -9223372036854775808
                        _ = v62
                    } else {
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        switch fi15 {
                        case 32768:
                            switch fi14 {
                            case 32768:
                                fmt.Println(fi15, v13)
                                var v63 uint = (v1 + (v1 - (^(v1))))
                                _ = v63
                            case 1:
                                if (v13 && (v13 || (v1 > 776))) {
                                    v64 := GenS7{F0: fi15, F1: 2, F2: 140}
                                    var v65 int = (fi15 >> (((v1 >> ((18446744073709551615) & 63))) & 63))
                                    _ = len(v3)
                                    _ = v64
                                    _ = v65
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = 17
                                    }
                                    if v13 {
                                        v66 := GenS1{F0: 153}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v67 uint8 = 0
                                        _ = v66
                                        _ = v67
                                    } else {
                                        fmt.Println(v11, fi14, fi14)
                                    }
                                    switch ("\n\t" + "") {
                                    case "\n\t":
                                        var v68 uint64 = 495
                                        _ = v68
                                    }
                                }
                            case 485:
                                fmt.Println(fi14, fi14, fi14)
                            default:
                                var v69 int64 = 1
                                _ = v69
                            }
                        case 807:
                            v70 := GenS7{F0: (fi14 | fi15), F1: 4294967295, F2: 244}
                            _ = v70
                        case 4294967295:
                            switch int(v11) {
                            case 205:
                                switch 127 {
                                case 9223372036854775807:
                                    v1 += (v1 % (((v1 / ((v1) | 1))) | 1))
                                    if ((^(fi14)) <= 128) {
                                        var v71 bool = (1 < -1)
                                        _ = v71
                                    }
                                case 909:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case -9223372036854775808:
                                    fmt.Println(fi15, v13)
                                }
                            case -2147483648:
                                if len(v2) > 0 {
                                    v2[0] = (fi14 / (((^(fi14))) | 1))
                                }
                            case -499:
                                if len(v2) > 0 {
                                    v2[0] = fi14
                                }
                            }
                        }
                    }
                default:
                    if v13 {
                        v72 := GenS5{F0: 65535, F1: -852, F2: 42}
                        _ = v72
                    } else {
                        if true {
                            v13 = (!true)
                            if true {
                                var v73 float64 = 3.14159
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                                if (genFn0() >= (v11 * v11)) {
                                    v74 := GenS4{F0: 0, F1: 255}
                                    _ = v74
                                } else {
                                    v75 := GenS8{F0: 1}
                                    if v13 {
                                        v76 := GenS7{F0: -856, F1: 170, F2: 2}
                                        _ = v76
                                    } else {
                                        var v77 int8 = -128
                                        var v78 uint = v1
                                        _ = v77
                                        _ = v78
                                    }
                                    fmt.Println(v1)
                                    _ = v75
                                }
                                _ = v73
                            } else {
                                if v13 {
                                    var v79 uint16 = 243
                                    v79 = 32767
                                    _ = v79
                                }
                                v3 = append(v3, "")
                                var v80 int64 = -1
                                fmt.Println(v11, v1, fi15)
                                v80 += 274
                                _ = v80
                            }
                        }
                        v11 -= (v11 % (((v11 ^ (v11 >> ((v1) & 63)))) | 1))
                        if (36 != 222) {
                            v81 := GenS5{F0: 4294967295, F1: -377, F2: -973}
                            v11 = v11
                            var v82 float32 = (0.1 + (3.14159 / 1.0))
                            _ = v81
                            _ = v82
                        }
                    }
                    var v83 float32 = ((1e10 - float32(-2.25)) / 1e10)
                    _ = v83
                }
                v13 = ((!((v13 && (fi14 > (-(fi14)))) || false)) && (65535 != 998))
                if (((3.14159 + 1e10) >= 1e10) || true) {
                    fmt.Println(v13)
                    var v84 bool = v13
                    _ = v84
                } else {
                    if ((float64((float32((float64(1.5) * ((1.5 * 1.5) * 1e10))) * (-2.25 / float32(3.14159)))) / 3.14159) >= 1.5) {
                        var v85 uint16 = 255
                        v86 := GenS8{F0: 1}
                        _ = v85
                        _ = v86
                    }
                }
                var v87 string = "foo bar"
                fmt.Println(v13, fi15, fi14)
                _ = v87
            } else {
                v13 = (fi15 <= (fi14 + fi15))
            }
            var v88 uint64 = 974
            _ = fi15
            _ = v88
        }
        _ = fi14
    }
    if (0 > 0) {
        fmt.Println(v1, v1, v13)
        var v89 int8 = 1
        if v13 {
            var v90 uint64 = 0
            v1 -= v1
            var v91 uint64 = v90
            _ = v90
            _ = v91
        }
        _ = v89
    }
    v92 := GenS5{F0: 1, F1: 59, F2: -990}
    switch 4294967295 {
    case -1:
        var v93 uint = (v1 ^ v1)
        _ = v93
    }
    switch -9223372036854775808 {
    case 128:
        var v94 int64 = -9223372036854775808
        _ = v94
    case -389:
        var v95 float32 = 0.1
        _ = v95
    case 698:
        if (87 == 18) {
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            v1 -= v1
        } else {
            switch 4294967295 {
            case -191:
                fmt.Println(v13, v11, v11)
            case 42:
                for fi96 := 0; fi96 < 2; fi96++ {
                    var v97 uint = v1
                    _ = fi96
                    _ = v97
                }
            }
            v98 := GenS9{F0: 1}
            _ = v98
        }
    }
    if (false || (v1 < 0)) {
        var v99 bool = true
        _ = len(v2)
        var v100 int64 = 42
        _ = v100
        _ = v99
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v92
}
