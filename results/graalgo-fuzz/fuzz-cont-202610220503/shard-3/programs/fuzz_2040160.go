package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int16; F1 uint8; F2 int8 }
type GenS5 struct { F0 int16; F1 bool; F2 uint8 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 int32; F1 uint }
type GenS8 struct { F0 bool; F1 int; F2 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32; F1 int16 }
type GenS11 struct { F0 int; F1 uint16 }
type GenS12 struct { F0 uint32; F1 int16; F2 int16 }
type GenS13 struct { F0 uint16; F1 int64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int64; F1 uint32; F2 uint }
type GenS17 struct { F0 string; F1 uint32; F2 int16 }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 uint; F1 uint8; F2 int64 }
type GenS20 struct { F0 bool; F1 uint32 }
type GenS21 struct { F0 uint8; F1 string; F2 int32 }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 uint64; F1 int32 }
type GenS24 struct { F0 bool; F1 string }
type GenS25 struct { F0 int16; F1 int8; F2 uint64 }
type GenS26 struct { F0 uint16 }
type GenS27 struct { F0 string }
type GenS28 struct { F0 uint64; F1 int; F2 uint32 }
type GenS29 struct { F0 uint32; F1 bool }
type GenS30 struct { F0 bool }
type GenS31 struct { F0 uint16; F1 bool; F2 uint8 }
func genFn0() uint8 {
    var v4 float32 = (1e10 + (1.5 / 0.1))
    if (!(914 == 32767)) {
        if (1 <= 42) {
            v5 := []int{945, 0, 32768}
            _ = len(v5)
            _ = v5
        } else {
            v6 := []int{646, 2147483647}
            _ = v6
        }
        for fi7 := 0; fi7 < 8; fi7++ {
            v4 += -2.25
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := GenS1{F0: 18446744073709551615, F1: 0}
                _ = fi8
                _ = v9
            }
            var v10 uint32 = 65535
            _ = fi7
            _ = v10
        }
        var v11 float64 = (float64(v4) + (3.14159 / float64(v4)))
        _ = v11
    }
    var v12 int16 = -32768
    _ = v12
    _ = v4
    return 73
}

func main() {
    var v1 float64 = (1.5 * 3.14159)
    var v2 int8 = -17
    var v3 bool = true
    var v13 uint64 = 479
    if true {
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 int8 = 100
                v17 := GenS1{F0: (v13 % ((v13) | 1)), F1: uint(v13)}
                if ((fi14 & 4294967295) != (fi15 + (fi14 << ((4294967295) & 63)))) {
                    v18 := GenS2{F0: 166, F1: ("foo bar" + "foo bar")}
                    if (((v13 - (^(v13))) >= (v13 >> ((4294967295) & 63))) && (267 < 32767)) {
                        if v3 {
                            v19 := GenS3{F0: 4294967295, F1: ("foo bar" + "hello")}
                            v2 = v2
                            if true {
                                v16 = (v2 - 100)
                            }
                            fmt.Println(v2)
                            var v20 uint8 = 127
                            _ = v19
                            _ = v20
                        }
                    } else {
                        v21 := GenS4{F0: 32767, F1: 255, F2: (v2 * v16)}
                        _ = v21
                    }
                    v22 := []int{int(v2), 459, (-(fi15))}
                    var v23 int64 = 94
                    switch "foo bar" {
                    case "\n\t":
                        v16 -= 100
                    case "hello":
                        v3 = (!v3)
                    case "foo bar":
                        _ = len(v22)
                    default:
                        fmt.Println(v2, v23, v23)
                    }
                    _ = v18
                    _ = v22
                    _ = v23
                } else {
                    switch "" {
                    case "foo bar":
                        v24 := []int{(fi14 * (fi15 - 4294967295))}
                        _ = v24
                    case "\n\t":
                        v25 := GenS4{F0: -809, F1: genFn0(), F2: (v16 - v2)}
                        _ = v25
                    }
                    fmt.Println(v3, v2)
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            v26 := GenS3{F0: 42, F1: "\n\t"}
            _ = fi14
            _ = v26
        }
        if v3 {
            v27 := []int{2147483647, 42, -987}
            v27 = append(v27, 2147483647)
            _ = v27
        } else {
            fmt.Println(v13, v1, v1)
            v28 := GenS5{F0: 623, F1: false, F2: 42}
            v2 = (v2 * v2)
            _ = v28
        }
    } else {
        v2 -= v2
    }
    switch "foo bar" {
    case "foo bar":
        v29 := []int{}
        _ = v29
    default:
        var v30 int = 2147483647
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        v1 = v1
        _ = fi31
    }
    v32 := []int{-624, 2147483647}
    for fi33 := 0; fi33 < 8; fi33++ {
        if v3 {
            v34 := GenS4{F0: 42, F1: 2, F2: -41}
            if (42 != 255) {
                for fi35 := 0; fi35 < 11; fi35++ {
                    v36 := GenS6{F0: 1, F1: 778}
                    v1 = -2.25
                    v2 = v2
                    _ = fi35
                    _ = v36
                }
            }
            _ = v34
        } else {
            v1 = v1
            var v37 uint = 87
            fmt.Println(v37, v3)
            _ = v37
        }
        if v3 {
            switch ("" + "\n\t") {
            case "a":
                var v38 uint32 = 65535
                _ = v38
            default:
                v3 = (false && v3)
            }
        }
        _ = fi33
    }
    var v39 float64 = float64(((1e10 * 0.1) + ((float32(v1) + 1.5) + 3.14159)))
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    switch "" {
    case "\n\t":
        if len(v32) > 0 {
            v32[0] = 201
        }
    case "":
        for fi40 := 0; fi40 < 11; fi40++ {
            if (-2.25 <= -2.25) {
                for fi41 := 0; fi41 < 11; fi41++ {
                    switch (("foo bar" + "\n\t") + "") {
                    case "hello":
                        if true {
                            v42 := GenS7{F0: 0, F1: 297}
                            fmt.Println(v3, v13, fi40)
                            _ = v42
                        } else {
                            if ((0.1 + ((v1 - ((float64(((float32(float64(((3.14159 + (0.1 + 0.1)) * 0.1))) + (0.1 - float32(float64(float32(v1))))) * 1e10)) / v1) / v1)) + v39)) > v1) {
                                v2 = (v2 % ((35) | 1))
                                v43 := GenS6{F0: 0, F1: 65535}
                                v39 = 0.0
                                _ = v43
                            }
                            var v44 int8 = v2
                            v45 := GenS8{F0: v3, F1: fi40, F2: ("\n\t" + "")}
                            _ = v44
                            _ = v45
                        }
                    case "\n\t":
                        _ = len(v32)
                    default:
                        var v46 bool = (0 > genFn0())
                        _ = v46
                    }
                    if len(v32) >= 2 {
                        v32 = v32[1:len(v32)]
                    }
                    _ = fi41
                }
                v32 = append(v32, -9223372036854775808)
                v47 := GenS9{F0: 0}
                _ = v47
            } else {
                v13 -= (v13 ^ (v13 + v13))
                if len(v32) > 0 {
                    v32[0] = fi40
                }
                if false {
                    var v48 uint = 65535
                    for fi49 := 0; fi49 < 7; fi49++ {
                        if (((((((((("hello" + "a") + "hello") + "\n\t") + "a") + ("foo bar" + "\n\t")) != "hello") || (1 > 1)) || v3) && v3) && (false || v3)) {
                            fmt.Println(v48, v48, v2)
                        } else {
                            var v50 string = "a"
                            switch 918 {
                            case -2147483648:
                                if (!(v3 || (2147483647 == 65535))) {
                                    switch fi49 {
                                    case 127:
                                        if false {
                                            v2 = -1
                                            v39 = (3.14159 + (v1 + ((float64((3.14159 - 0.1)) / 0.1) * (((-2.25 / 1.0) * ((0.0 / v39) - float64(0.1))) * 1e10))))
                                            v48 = (v48 - v48)
                                            v51 := GenS6{F0: 1, F1: 478}
                                            if (v50 == v50) {
                                                if v3 {
                                                    if v3 {
                                                        if (!(v48 > 255)) {
                                                            _ = len(v32)
                                                            v2 = (v2 + -48)
                                                            var v52 uint32 = 42
                                                            _ = v52
                                                        }
                                                        fmt.Println(v13, v39)
                                                        if (-1 > 1) {
                                                            fmt.Println(v48)
                                                            switch 4294967295 {
                                                            case 128:
                                                                v32 = append(v32, (-(fi49)))
                                                            case 534:
                                                                v53 := GenS8{F0: (v3 || true), F1: (^(fi40)), F2: v50}
                                                                _ = v53
                                                            case 508:
                                                                fmt.Println(v1, v50)
                                                            }
                                                        }
                                                    } else {
                                                        if v3 {
                                                            switch v50 {
                                                            case "foo bar":
                                                                fmt.Println(v39)
                                                                v54 := GenS8{F0: (-297 == int32(v2)), F1: fi49, F2: (v50 + v50)}
                                                                var v55 uint8 = 128
                                                                _ = v54
                                                                _ = v55
                                                            case "\n\t":
                                                                var v56 uint = v48
                                                                _ = v56
                                                            case "":
                                                                v50 = v50
                                                            default:
                                                                if len(v32) >= 2 {
                                                                    v32 = v32[1:len(v32)]
                                                                }
                                                            }
                                                        } else {
                                                            if v3 {
                                                                if (false || true) {
                                                                    fmt.Println(v3, v1)
                                                                    var v57 int = fi40
                                                                    _ = v57
                                                                } else {
                                                                    switch fi40 {
                                                                    case 793:
                                                                        v58 := GenS3{F0: 42, F1: v50}
                                                                        _ = v58
                                                                    case 127:
                                                                        var v59 uint64 = (v13 * 469)
                                                                        _ = v59
                                                                    default:
                                                                        if len(v32) > 0 {
                                                                            v32[0] = fi49
                                                                        }
                                                                    }
                                                                }
                                                                if (2147483647 == 2147483647) {
                                                                    var v60 uint = (v48 >> ((v48) & 63))
                                                                    if (int32(v13) > -1) {
                                                                        v61 := GenS3{F0: 65535, F1: ("foo bar" + v50)}
                                                                        if v3 {
                                                                            var v62 uint8 = genFn0()
                                                                            var v63 int8 = -55
                                                                            v62 = genFn0()
                                                                            _ = v62
                                                                            _ = v63
                                                                        } else {
                                                                            v1 -= v39
                                                                            v32 = append(v32, (fi40 ^ 32768))
                                                                            fmt.Println(v3, v13, v48)
                                                                        }
                                                                        _ = v61
                                                                    }
                                                                    switch v50 {
                                                                    case "":
                                                                        v64 := GenS10{F0: 42, F1: 1000}
                                                                        _ = v64
                                                                    case "foo bar":
                                                                        v60 = v60
                                                                    case "\n\t":
                                                                        var v65 int8 = v2
                                                                        _ = v65
                                                                    default:
                                                                        fmt.Println(v50, v50)
                                                                    }
                                                                    _ = v60
                                                                }
                                                                var v66 uint32 = 2
                                                                _ = v66
                                                            } else {
                                                                v2 = (v2 % (((v2 * 0)) | 1))
                                                                if (false || ((v3 && (!false)) && (!(188 < v13)))) {
                                                                    if (128 <= 2) {
                                                                        var v67 string = v50
                                                                        _ = len(v32)
                                                                        switch fi49 {
                                                                        case -680:
                                                                            switch (fi49 | fi40) {
                                                                            case -409:
                                                                                if len(v32) > 0 {
                                                                                    v32[0] = fi49
                                                                                }
                                                                            case 32768:
                                                                                if len(v32) > 0 {
                                                                                    v32[0] = (fi49 >> ((v48) & 63))
                                                                                }
                                                                            }
                                                                        case -2147483648:
                                                                            if len(v32) >= 2 {
                                                                                v32 = v32[1:len(v32)]
                                                                            }
                                                                        case 32768:
                                                                            var v68 uint8 = 127
                                                                            _ = v68
                                                                        default:
                                                                            _ = len(v32)
                                                                        }
                                                                        _ = v67
                                                                    } else {
                                                                        if len(v32) >= 2 {
                                                                            v32 = v32[1:len(v32)]
                                                                        }
                                                                        var v69 uint = (v48 << (((v48 | v48)) & 63))
                                                                        _ = v69
                                                                    }
                                                                    if (v48 == (^(v48))) {
                                                                        if len(v32) >= 2 {
                                                                            v32 = v32[1:len(v32)]
                                                                        }
                                                                    }
                                                                    var v70 uint64 = v13
                                                                    _ = v70
                                                                }
                                                            }
                                                            v32 = append(v32, 127)
                                                            v2 = 0
                                                        }
                                                        if len(v32) >= 2 {
                                                            v32 = v32[1:len(v32)]
                                                        }
                                                        v50 = (v50 + ("\n\t" + v50))
                                                    }
                                                } else {
                                                    if (3.14159 <= 3.14159) {
                                                        if v3 {
                                                            fmt.Println(v1)
                                                        } else {
                                                            switch (v50 + "a") {
                                                            case "hello":
                                                                var v71 int16 = 1
                                                                _ = v71
                                                            case "":
                                                                v2 = (v2 % ((v2) | 1))
                                                                v32 = append(v32, 9223372036854775807)
                                                                v50 = (("" + (("hello" + ("\n\t" + (((v50 + "") + "a") + (v50 + v50)))) + v50)) + ((v50 + "hello") + v50))
                                                            }
                                                            if true {
                                                                fmt.Println(v1, v48, v13)
                                                                if (0 > -32768) {
                                                                    var v72 uint64 = 562
                                                                    if (v3 || v3) {
                                                                        v13 = v72
                                                                        var v73 float64 = -2.25
                                                                        _ = v73
                                                                    } else {
                                                                        v2 = 0
                                                                        if true {
                                                                            v39 += 1.5
                                                                            if (v3 && (3.14159 >= 0.1)) {
                                                                                v1 += 1.5
                                                                                fmt.Println(fi40, v72, fi40)
                                                                            }
                                                                        } else {
                                                                            if (!(0 <= 42)) {
                                                                                if false {
                                                                                    switch fi40 {
                                                                                    case 42:
                                                                                        v74 := GenS6{F0: 2147483647, F1: 319}
                                                                                        _ = v74
                                                                                    }
                                                                                    var v75 uint64 = v13
                                                                                    switch 2147483647 {
                                                                                    case 0:
                                                                                        switch ("a" + (v50 + (v50 + v50))) {
                                                                                        case "":
                                                                                            v48 = (^(v48))
                                                                                        }
                                                                                    case 32768:
                                                                                        v75 = (^(v72))
                                                                                    default:
                                                                                        v76 := GenS11{F0: fi40, F1: 112}
                                                                                        _ = v76
                                                                                    }
                                                                                    _ = v75
                                                                                } else {
                                                                                    v32 = append(v32, 127)
                                                                                }
                                                                                fmt.Println(v39, v50, v48)
                                                                                fmt.Println(fi49, v39)
                                                                            } else {
                                                                                v32 = append(v32, (fi40 | fi40))
                                                                                var v77 string = (v50 + v50)
                                                                                v3 = v3
                                                                                _ = v77
                                                                            }
                                                                            v78 := GenS10{F0: 960, F1: 1}
                                                                            if len(v32) > 0 {
                                                                                v32[0] = (fi40 | 127)
                                                                            }
                                                                            _ = v78
                                                                        }
                                                                        v32 = append(v32, (fi40 | (fi49 + fi40)))
                                                                        v48 += v48
                                                                        var v79 uint8 = 2
                                                                        var v80 int16 = 42
                                                                        _ = v79
                                                                        _ = v80
                                                                    }
                                                                    _ = v72
                                                                } else {
                                                                    v48 += 4294967295
                                                                }
                                                                if v3 {
                                                                    v81 := GenS12{F0: 2147483647, F1: 1000, F2: 949}
                                                                    var v82 uint = v48
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                            }
                                                            var v83 uint64 = (v13 ^ (v13 ^ (v13 % (((v13 >> ((v48) & 63))) | 1))))
                                                            _ = v83
                                                        }
                                                    }
                                                    v84 := GenS13{F0: 127, F1: 42}
                                                    if (-2147483648 <= -2147483648) {
                                                        v85 := GenS3{F0: v48, F1: v50}
                                                        v3 = (-1 != int32(v2))
                                                        if v3 {
                                                            if (-1 <= -459) {
                                                                v32 = append(v32, -989)
                                                                _ = len(v32)
                                                                var v86 int8 = v2
                                                                _ = v86
                                                            } else {
                                                                switch 2147483647 {
                                                                case 0:
                                                                    v2 = v2
                                                                case 32768:
                                                                    v32 = append(v32, fi40)
                                                                default:
                                                                    var v87 int = (fi40 + 32768)
                                                                    var v88 string = "a"
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            }
                                                            var v89 uint32 = 170
                                                            _ = v89
                                                        } else {
                                                            var v90 bool = ((255 != v48) || v3)
                                                            _ = v90
                                                        }
                                                        var v91 float32 = 1e10
                                                        v1 = 1.5
                                                        if v3 {
                                                            if (fi40 > fi49) {
                                                                var v92 bool = (!v3)
                                                                var v93 bool = (2147483647 > (fi40 % (((fi49 + (fi49 % ((-228) | 1)))) | 1)))
                                                                _ = v92
                                                                _ = v93
                                                            } else {
                                                                var v94 float32 = v91
                                                                fmt.Println(v91, fi49, v91)
                                                                _ = v94
                                                            }
                                                        }
                                                        _ = v85
                                                        _ = v91
                                                    } else {
                                                        fmt.Println(v1)
                                                        var v95 string = ""
                                                        var v96 int16 = 32767
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    var v97 int32 = 411
                                                    _ = v84
                                                    _ = v97
                                                }
                                                v3 = (!(0.0 == 0.1))
                                                var v98 uint32 = 65535
                                                switch (fi49 - (fi49 * (fi40 & (-(fi40))))) {
                                                case 0:
                                                    switch fi49 {
                                                    case 1:
                                                        v13 -= 1
                                                    case 485:
                                                        fmt.Println(fi49, v50)
                                                    }
                                                case -2147483648:
                                                    if v3 {
                                                        if (238 != 68) {
                                                            v50 = ""
                                                            v2 = v2
                                                            fmt.Println(v2)
                                                        }
                                                        fmt.Println(v98, fi49)
                                                        v50 = "\n\t"
                                                    }
                                                }
                                                fmt.Println(v3, v48)
                                                _ = v98
                                            } else {
                                                fmt.Println(v39, v2)
                                                if len(v32) > 0 {
                                                    v32[0] = (fi40 / ((fi40) | 1))
                                                }
                                            }
                                            _ = v51
                                        }
                                    case 128:
                                        var v99 uint32 = 23
                                        _ = v99
                                    }
                                } else {
                                    _ = len(v32)
                                    var v100 int64 = 0
                                    v48 = (v48 | (v48 + 255))
                                    if (fi40 < (fi40 << ((18446744073709551615) & 63))) {
                                        v101 := GenS14{F0: 1000}
                                        _ = v101
                                    }
                                    var v102 float32 = 0.0
                                    if len(v32) >= 2 {
                                        v32 = v32[1:len(v32)]
                                    }
                                    _ = v100
                                    _ = v102
                                }
                            }
                            var v103 float64 = (0.1 / v1)
                            if len(v32) > 0 {
                                v32[0] = fi49
                            }
                            _ = v103
                            _ = v50
                        }
                        var v104 string = "foo bar"
                        v32 = append(v32, (fi40 % ((fi49) | 1)))
                        _ = fi49
                        _ = v104
                    }
                    _ = v48
                }
            }
            _ = fi40
        }
    case "foo bar":
        v105 := GenS2{F0: 255, F1: ""}
        _ = v105
    default:
        for fi106 := 0; fi106 < 14; fi106++ {
            if true {
                for fi107 := 0; fi107 < 17; fi107++ {
                    v39 += v1
                    v108 := GenS14{F0: -594}
                    fmt.Println(v13)
                    _ = fi107
                    _ = v108
                }
                v39 -= 1.5
            }
            for fi109 := 0; fi109 < 4; fi109++ {
                var v110 bool = v3
                if v3 {
                    if true {
                        if (934 == 2) {
                            switch -1 {
                            case 4294967295:
                                v111 := GenS5{F0: 97, F1: (!false), F2: 72}
                                _ = v111
                            case 42:
                                var v112 uint = 1
                                _ = v112
                            case -9223372036854775808:
                                v1 -= v39
                            }
                            var v113 float32 = (-2.25 * 0.0)
                            var v114 uint16 = 379
                            _ = v113
                            _ = v114
                        }
                        switch "hello" {
                        case "a":
                            v3 = v110
                        case "hello":
                            var v115 int32 = 42
                            _ = v115
                        default:
                            if (!v3) {
                                v13 = v13
                                switch "hello" {
                                case "hello":
                                    if true {
                                        v1 -= 1.5
                                    }
                                    if (v3 && ((v13 ^ v13) == (^(v13)))) {
                                        v116 := GenS9{F0: -9223372036854775808}
                                        switch "" {
                                        case "foo bar":
                                            if len(v32) >= 2 {
                                                v32 = v32[1:len(v32)]
                                            }
                                        default:
                                            v2 = (v2 >> ((0) & 63))
                                        }
                                        v13 = 647
                                        _ = v116
                                    } else {
                                        if (v2 <= -118) {
                                            fmt.Println(v2)
                                            switch "foo bar" {
                                            case "foo bar":
                                                if len(v32) >= 2 {
                                                    v32 = v32[1:len(v32)]
                                                }
                                            case "a":
                                                _ = len(v32)
                                            case "":
                                                var v117 uint16 = 334
                                                _ = v117
                                            default:
                                                var v118 int32 = 444
                                                if v110 {
                                                    var v119 bool = v110
                                                    var v120 uint16 = 0
                                                    _ = v119
                                                    _ = v120
                                                } else {
                                                    v121 := GenS15{F0: (((!v110) && false) && (146 <= 128))}
                                                    switch "foo bar" {
                                                    case "a":
                                                        v122 := GenS10{F0: (v118 >> ((2) & 63)), F1: 965}
                                                        _ = v122
                                                    case "hello":
                                                        fmt.Println(fi106, v13)
                                                    case "":
                                                        fmt.Println(v39, fi106, fi109)
                                                    default:
                                                        v39 = (v39 - ((0.0 + 0.0) + v39))
                                                    }
                                                    _ = v121
                                                }
                                                _ = v118
                                            }
                                        }
                                        var v123 uint = 1
                                        _ = v123
                                    }
                                case "":
                                    v124 := GenS8{F0: ((false && (!(22 != (^(v2))))) || (!(v110 || v110))), F1: 1, F2: ""}
                                    _ = v124
                                }
                            } else {
                                var v125 int8 = (^(v2))
                                _ = v125
                            }
                        }
                        switch fi109 {
                        case -1:
                            if ((float64((((1e10 / 1e10) + ((0.1 - ((-2.25 + float32(v1)) / 0.1)) - -2.25)) / float32(v39))) / v1) != -2.25) {
                                v32 = append(v32, fi106)
                                v3 = v3
                            } else {
                                if len(v32) >= 2 {
                                    v32 = v32[1:len(v32)]
                                }
                            }
                        default:
                            if v110 {
                                v32 = append(v32, fi106)
                            }
                        }
                    } else {
                        var v126 uint = 255
                        v39 -= float64(1.5)
                        if v3 {
                            v127 := GenS13{F0: 32767, F1: -599}
                            v110 = (32767 == 42)
                            if (fi109 >= fi106) {
                                if (v3 || false) {
                                    v128 := GenS16{F0: 1, F1: 2, F2: (v126 / (((v126 + v126)) | 1))}
                                    v129 := GenS17{F0: ("" + "foo bar"), F1: 65535, F2: -32768}
                                    _ = v128
                                    _ = v129
                                } else {
                                    fmt.Println(fi109, fi109)
                                    v126 = v126
                                    var v130 int32 = -1
                                    _ = v130
                                }
                            } else {
                                v131 := GenS6{F0: 1, F1: 32767}
                                _ = v131
                            }
                            _ = v127
                        } else {
                            var v132 uint = v126
                            _ = v132
                        }
                        v3 = (!v110)
                        fmt.Println(v1, v126)
                        _ = v126
                    }
                    switch ("foo bar" + "") {
                    case "hello":
                        switch "\n\t" {
                        case "\n\t":
                            v133 := GenS18{F0: 0}
                            _ = v133
                        case "hello":
                            v134 := GenS11{F0: 9223372036854775807, F1: 242}
                            _ = v134
                        case "a":
                            v110 = v110
                            switch fi106 {
                            case 1:
                                v135 := GenS5{F0: 1, F1: ("hello" >= ("hello" + (("a" + "\n\t") + (("foo bar" + "a") + "")))), F2: 2}
                                _ = v135
                            }
                        default:
                            var v136 bool = (255 == 255)
                            _ = v136
                        }
                    }
                } else {
                    if v110 {
                        if len(v32) >= 2 {
                            v32 = v32[1:len(v32)]
                        }
                        switch "hello" {
                        case "hello":
                            v32 = append(v32, fi109)
                        case "foo bar":
                            fmt.Println(v3, v39, v1)
                        default:
                            if false {
                                var v137 int32 = 1
                                switch int(v137) {
                                case -1:
                                    v110 = false
                                case 87:
                                    var v138 string = "hello"
                                    _ = v138
                                default:
                                    var v139 int8 = (v2 << ((0) & 63))
                                    _ = v139
                                }
                                _ = v137
                            } else {
                                var v140 string = "foo bar"
                                var v141 int64 = 1
                                v140 = v140
                                _ = v140
                                _ = v141
                            }
                        }
                        v142 := GenS11{F0: (fi109 & (fi109 >> ((717) & 63))), F1: 0}
                        if true {
                            if true {
                                if len(v32) > 0 {
                                    v32[0] = fi109
                                }
                            } else {
                                _ = len(v32)
                            }
                        } else {
                            v2 = (^(v2))
                            var v143 int32 = -1
                            _ = len(v32)
                            var v144 int16 = 1000
                            _ = v143
                            _ = v144
                        }
                        _ = v142
                    }
                    if (v3 || v110) {
                        v2 += int8(v13)
                        var v145 int32 = -2147483648
                        var v146 int = fi106
                        var v147 int = 447
                        v110 = v3
                        _ = v145
                        _ = v146
                        _ = v147
                    } else {
                        v148 := GenS19{F0: 606, F1: 127, F2: 42}
                        _ = v148
                    }
                    var v149 int = 1
                    switch fi106 {
                    case -2147483648:
                        var v150 int64 = 42
                        _ = v150
                    case 9223372036854775807:
                        v2 = v2
                    case 4294967295:
                        var v151 uint16 = 661
                        _ = v151
                    default:
                        if v110 {
                            if (!(!false)) {
                                var v152 uint8 = 215
                                v1 = float64((0.1 - 0.1))
                                _ = v152
                            } else {
                                var v153 float64 = 3.14159
                                v154 := GenS1{F0: 890, F1: 4294967295}
                                _ = v153
                                _ = v154
                            }
                            switch v149 {
                            case 9223372036854775807:
                                v2 -= -54
                            }
                        } else {
                            v13 += 42
                            v155 := GenS15{F0: (true && (!v110))}
                            if false {
                                switch "hello" {
                                case "a":
                                    var v156 uint32 = 0
                                    _ = v156
                                case "foo bar":
                                    var v157 float32 = 3.14159
                                    _ = v157
                                case "\n\t":
                                    if (1 <= 1) {
                                        v149 += (^(fi106))
                                        v110 = (0 >= 42)
                                    } else {
                                        if v3 {
                                            switch 2147483647 {
                                            case 1:
                                                if len(v32) > 0 {
                                                    v32[0] = 1
                                                }
                                            default:
                                                if v110 {
                                                    if len(v32) >= 2 {
                                                        v32 = v32[1:len(v32)]
                                                    }
                                                    v1 -= v1
                                                }
                                            }
                                        } else {
                                            var v158 uint = 65535
                                            v149 = (-(fi109))
                                            _ = len(v32)
                                            v159 := GenS20{F0: v110, F1: 292}
                                            _ = v158
                                            _ = v159
                                        }
                                        var v160 uint = 0
                                        _ = v160
                                    }
                                default:
                                    var v161 uint16 = 2
                                    _ = v161
                                }
                                if v110 {
                                    if (2147483647 > int32(v13)) {
                                        var v162 int8 = (v2 | (v2 & v2))
                                        v39 -= (v39 * v1)
                                        fmt.Println(fi109, fi109)
                                        _ = v162
                                    }
                                    v149 += fi106
                                }
                                if v110 {
                                    if ((((v39 + 3.14159) / (((float64(float32((v1 * v1))) - v39) / float64((3.14159 * (((-2.25 / 1.5) + (float32(1.5) / float32(((((3.14159 * -2.25) + (v1 * (v39 + 1.5))) / 3.14159) + v1)))) * (3.14159 - float32(v39)))))) + v1)) - 0.1) > (v39 * (((v39 / (((float64(1.5) - (v39 / v1)) - v1) * float64(0.1))) * v39) + v1))) {
                                        var v163 uint8 = 255
                                        if v110 {
                                            v1 -= v1
                                            v13 -= v13
                                            if len(v32) >= 2 {
                                                v32 = v32[1:len(v32)]
                                            }
                                            v39 = v1
                                            switch "" {
                                            case "hello":
                                                switch v149 {
                                                case -239:
                                                    v13 = (^(v13))
                                                case 42:
                                                    switch (fi109 | v149) {
                                                    case -85:
                                                        var v164 bool = (v110 && ((!(1e10 < 1.5)) && (!v110)))
                                                        _ = v164
                                                    }
                                                case 9223372036854775807:
                                                    var v165 float32 = (-2.25 + ((((0.0 + float32((v39 + 1e10))) * ((0.1 / 0.1) + (0.0 * (0.1 - 3.14159)))) - (float32(v1) - 1.5)) * 0.0))
                                                    _ = v165
                                                default:
                                                    switch int(v2) {
                                                    case 128:
                                                        v2 -= 127
                                                    case 212:
                                                        fmt.Println(v163)
                                                    case -87:
                                                        if v3 {
                                                            var v166 uint32 = 849
                                                            v167 := GenS15{F0: v110}
                                                            fmt.Println(v2, v3, v149)
                                                            _ = v166
                                                            _ = v167
                                                        }
                                                    default:
                                                        v168 := GenS21{F0: (v163 * (v163 & v163)), F1: "a", F2: 1}
                                                        _ = v168
                                                    }
                                                    v149 = 9223372036854775807
                                                }
                                            case "foo bar":
                                                var v169 bool = (-1 <= 492)
                                                _ = v169
                                            }
                                        }
                                        v110 = ((v163 & (v163 & (v163 + 18))) != v163)
                                        v32 = append(v32, (fi109 & fi106))
                                        _ = v163
                                    }
                                }
                            } else {
                                fmt.Println(fi106, v39, v3)
                                var v170 int16 = 32767
                                _ = v170
                            }
                            _ = v155
                        }
                    }
                    v39 += float64(0.1)
                    _ = v149
                }
                v171 := GenS22{F0: 1}
                _ = fi109
                _ = v110
                _ = v171
            }
            for fi172 := 0; fi172 < 20; fi172++ {
                _ = len(v32)
                _ = fi172
            }
            _ = fi106
        }
    }
    v3 = (3.14159 >= 1e10)
    for fi173 := 0; fi173 < 10; fi173++ {
        v1 += v39
        if v3 {
            var v174 uint16 = 255
            if ("" == "\n\t") {
                switch (fi173 + (fi173 >> ((4294967295) & 63))) {
                case -135:
                    var v175 int16 = 32767
                    _ = v175
                case 9223372036854775807:
                    v13 = (v13 << ((42) & 63))
                case 128:
                    for fi176 := 0; fi176 < 7; fi176++ {
                        if v3 {
                            fmt.Println(fi173, fi176, v3)
                            _ = len(v32)
                            if true {
                                if v3 {
                                    v3 = false
                                    if ((^(v2)) >= v2) {
                                        v177 := GenS23{F0: v13, F1: -2147483648}
                                        v13 = v13
                                        _ = v177
                                    }
                                    if ((^(v13)) > (^(v13))) {
                                        fmt.Println(v13, v39)
                                    }
                                } else {
                                    if ((v13 == (v13 * v13)) && true) {
                                        v178 := GenS11{F0: (fi176 & fi176), F1: (v174 << ((0) & 63))}
                                        switch fi176 {
                                        case -1:
                                            switch 32768 {
                                            case 0:
                                                fmt.Println(v13, v1)
                                            default:
                                                var v179 int16 = 32767
                                                if true {
                                                    v180 := GenS4{F0: (v179 >> ((18446744073709551615) & 63)), F1: 19, F2: (v2 & v2)}
                                                    var v181 bool = v3
                                                    v181 = v181
                                                    _ = v180
                                                    _ = v181
                                                } else {
                                                    var v182 string = "a"
                                                    v39 += (v39 - v39)
                                                    v183 := GenS24{F0: ((true || (!false)) || v3), F1: v182}
                                                    _ = v182
                                                    _ = v183
                                                }
                                                v13 = 881
                                                if v3 {
                                                    if (0.1 <= 0.0) {
                                                        var v184 uint16 = v174
                                                        var v185 string = ("" + (("foo bar" + "\n\t") + "a"))
                                                        var v186 int32 = 42
                                                        _ = v184
                                                        _ = v185
                                                        _ = v186
                                                    } else {
                                                        var v187 float64 = (v39 + v39)
                                                        v179 = (-(v179))
                                                        _ = v187
                                                    }
                                                    if len(v32) > 0 {
                                                        v32[0] = 0
                                                    }
                                                } else {
                                                    var v188 uint16 = v174
                                                    _ = v188
                                                }
                                                _ = v179
                                            }
                                        case 627:
                                            v39 = (float64(float32(v1)) + -2.25)
                                        default:
                                            var v189 int = fi176
                                            _ = v189
                                        }
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v39, fi173)
                                        default:
                                            v190 := GenS25{F0: -1, F1: v2, F2: 18446744073709551615}
                                            _ = v190
                                        }
                                        _ = v178
                                    } else {
                                        if len(v32) >= 2 {
                                            v32 = v32[1:len(v32)]
                                        }
                                    }
                                    v174 = 580
                                    if (!(!true)) {
                                        if v3 {
                                            v39 = v39
                                            var v191 float64 = ((v39 / (0.0 / ((float64((-2.25 / float32(v39))) / v39) / (v1 - v1)))) - (v39 - (v1 / 1.5)))
                                            v3 = (0.1 >= v191)
                                            _ = v191
                                        } else {
                                            fmt.Println(v39, v39, v3)
                                            switch "\n\t" {
                                            case "a":
                                                var v192 uint8 = 0
                                                _ = v192
                                            case "\n\t":
                                                if len(v32) >= 2 {
                                                    v32 = v32[1:len(v32)]
                                                }
                                            case "hello":
                                                var v193 uint32 = 65535
                                                _ = v193
                                            default:
                                                var v194 uint8 = 255
                                                _ = v194
                                            }
                                            fmt.Println(v39)
                                            v32 = append(v32, fi173)
                                            var v195 uint8 = genFn0()
                                            _ = len(v32)
                                            _ = len(v32)
                                            _ = v195
                                        }
                                        fmt.Println(fi176, v13, v3)
                                        if (true && v3) {
                                            v196 := GenS26{F0: v174}
                                            var v197 uint64 = v13
                                            _ = v196
                                            _ = v197
                                        } else {
                                            if (false || (972 < 2147483647)) {
                                                var v198 uint64 = 0
                                                var v199 int8 = 0
                                                var v200 uint = 42
                                                _ = v198
                                                _ = v199
                                                _ = v200
                                            }
                                            v201 := GenS27{F0: ""}
                                            _ = v201
                                        }
                                    } else {
                                        if (-1 >= 1) {
                                            v39 = float64((1.5 * -2.25))
                                        }
                                        v13 -= v13
                                        if v3 {
                                            v202 := GenS8{F0: false, F1: -1, F2: "hello"}
                                            switch ("a" + ("hello" + (("" + "") + "foo bar"))) {
                                            case "\n\t":
                                                var v203 int64 = 0
                                                _ = v203
                                            case "foo bar":
                                                if false {
                                                    fmt.Println(v174)
                                                    v204 := GenS28{F0: 760, F1: (-(fi176)), F2: 2147483647}
                                                    var v205 float64 = v1
                                                    _ = v204
                                                    _ = v205
                                                } else {
                                                    var v206 int64 = -1
                                                    if len(v32) >= 2 {
                                                        v32 = v32[1:len(v32)]
                                                    }
                                                    if (v3 || v3) {
                                                        var v207 int64 = -736
                                                        if (!((fi173 % ((fi176) | 1)) <= (fi173 ^ fi176))) {
                                                            switch "hello" {
                                                            case "a":
                                                                switch fi176 {
                                                                case 4294967295:
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        var v208 bool = ((3.14159 - -2.25) < float32((((v1 + ((v1 + (((float64(3.14159) * (v39 * v1)) - v39) - v39)) + (v1 / (((v39 * 1e10) * 0.1) - (((((v1 + v39) / float64(0.1)) + 0.0) * 3.14159) / (3.14159 / v39)))))) / v1) / 3.14159)))
                                                                        _ = v208
                                                                    default:
                                                                        v207 += 9223372036854775807
                                                                    }
                                                                default:
                                                                    v32 = append(v32, (fi176 - fi173))
                                                                }
                                                            case "foo bar":
                                                                if len(v32) >= 2 {
                                                                    v32 = v32[1:len(v32)]
                                                                }
                                                            default:
                                                                fmt.Println(fi173, fi173, fi176)
                                                            }
                                                            v1 += v1
                                                            if ((-1 == 376) && (0 < 2147483647)) {
                                                                fmt.Println(v13, v1, v3)
                                                                var v209 float32 = (1e10 / (1e10 + 0.0))
                                                                v206 = (v207 >> ((65535) & 63))
                                                                _ = v209
                                                            }
                                                        } else {
                                                            v210 := GenS29{F0: 1, F1: v3}
                                                            v3 = (v3 || v3)
                                                            _ = v210
                                                        }
                                                        var v211 int8 = v2
                                                        _ = v207
                                                        _ = v211
                                                    }
                                                    _ = len(v32)
                                                    _ = v206
                                                }
                                            }
                                            v212 := GenS28{F0: v13, F1: (fi173 & (fi173 | 651)), F2: 1}
                                            _ = v202
                                            _ = v212
                                        }
                                        var v213 uint16 = (^(v174))
                                        v214 := GenS30{F0: (!(((0.0 - 1.5) + float32(v1)) <= 1.5))}
                                        _ = v213
                                        _ = v214
                                    }
                                }
                            }
                        }
                        _ = fi176
                    }
                default:
                    var v215 uint = 2
                    _ = v215
                }
            } else {
                var v216 uint8 = 255
                v32 = append(v32, (fi173 & fi173))
                switch fi173 {
                case 9223372036854775807:
                    if len(v32) > 0 {
                        v32[0] = fi173
                    }
                case 955:
                    if len(v32) > 0 {
                        v32[0] = fi173
                    }
                default:
                    var v217 uint64 = v13
                    _ = v217
                }
                _ = v216
            }
            fmt.Println(v39)
            _ = v174
        }
        var v218 string = "hello"
        _ = fi173
        _ = v218
    }
    if len(v32) > 0 {
        v32[0] = 127
    }
    v219 := GenS31{F0: 128, F1: ("hello" == "\n\t"), F2: 127}
    for fi220 := 0; fi220 < 8; fi220++ {
        switch (fi220 - (fi220 - (-(fi220)))) {
        case 2147483647:
            var v221 int32 = 0
            _ = v221
        case -9223372036854775808:
            v2 = v2
        case 127:
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
        }
        _ = fi220
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v13
    _ = v2
    _ = v219
    _ = v3
    _ = v32
    _ = v39
}
