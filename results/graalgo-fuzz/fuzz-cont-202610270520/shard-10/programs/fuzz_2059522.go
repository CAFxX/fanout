package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 string; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint16; F2 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8 }
func genFn0() int32 {
    v4 := []int{}
    v5 := GenS1{F0: (!(866 < 42)), F1: 100}
    _ = v4
    _ = v5
    return -489
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -1
    var v3 uint8 = 127
    if (127 >= 315) {
        var v6 float32 = (1e10 / (3.14159 * 0.1))
        for fi7 := 0; fi7 < 3; fi7++ {
            v6 -= ((v6 * 0.1) / 3.14159)
            _ = fi7
        }
        _ = v6
    } else {
        v3 = (v3 ^ 24)
    }
    switch "" {
    case "a":
        v3 += 198
    case "":
        var v8 int = -20
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v3 -= 128
        var v10 int8 = 42
        _ = fi9
        _ = v10
    }
    if (false || false) {
        v11 := v1[65535]
        if true {
            v2 = v2
            v3 = v3
            fmt.Println(v3)
        } else {
            fmt.Println(v11)
        }
        var v12 float64 = 0.0
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v2 += (v2 % ((v2) | 1))
        _ = fi13
    }
    switch ("" + (("a" + "hello") + "")) {
    case "\n\t":
        var v14 int8 = 0
        _ = v14
    case "hello":
        var v15 uint = 560
        _ = v15
    }
    v2 += (v2 & 2147483647)
    var v16 bool = (false || true)
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 int32 = v2
            switch 42 {
            case 2147483647:
                var v20 bool = (0 == 47)
                _ = v20
            case -790:
                v3 -= 127
            case -2147483648:
                v21 := GenS2{F0: 32767}
                _ = v21
            }
            v19 = (v19 | (^(v19)))
            _ = fi18
            _ = v19
        }
        v22 := []int{(^(fi17))}
        v1[42] = "a"
        _ = fi17
        _ = v22
    }
    var v23 uint16 = 1
    if v16 {
        var v24 uint32 = 509
        _ = v24
    }
    fmt.Println(v2, v23, v23)
    switch 4294967295 {
    case 0:
        if (-119 >= -41) {
            v25 := []int{42, 32768}
            fmt.Println(v23, v3, v16)
            v26 := GenS3{F0: 2147483648, F1: "\n\t", F2: (v3 * uint8(v23))}
            _ = v25
            _ = v26
        } else {
            for fi27 := 0; fi27 < 9; fi27++ {
                if false {
                    if v16 {
                        v28 := GenS3{F0: 988, F1: ("" + ("a" + "a")), F2: (v3 / ((v3) | 1))}
                        switch "" {
                        case "foo bar":
                            if (42 <= 421) {
                                v2 += (v2 ^ v2)
                            } else {
                                if false {
                                    switch "hello" {
                                    case "a":
                                        for fi29 := 0; fi29 < 17; fi29++ {
                                            v16 = true
                                            var v30 float32 = 0.0
                                            if v16 {
                                                var v31 int64 = 42
                                                var v32 int32 = v2
                                                v2 = (v2 ^ (v32 % ((v2) | 1)))
                                                _ = v31
                                                _ = v32
                                            }
                                            if ((!v16) || (!(int64(v23) > 0))) {
                                                v3 = 1
                                                v33 := GenS1{F0: ("hello" >= ""), F1: 64}
                                                _ = v33
                                            }
                                            v16 = ((0.1 + float32((0.0 / float64(v30)))) != ((float32((3.14159 / 1.5)) + v30) / 1.0))
                                            _ = fi29
                                            _ = v30
                                        }
                                    default:
                                        v34 := GenS1{F0: true, F1: 42}
                                        _ = v34
                                    }
                                    var v35 int = fi27
                                    _ = v35
                                } else {
                                    fmt.Println(v23)
                                    fmt.Println(fi27)
                                }
                                var v36 uint8 = v3
                                v1[2147483647] = "foo bar"
                                _ = v36
                            }
                        case "":
                            var v37 uint = uint(v23)
                            _ = v37
                        default:
                            for fi38 := 0; fi38 < 18; fi38++ {
                                v16 = (!((100 == 127) && (v16 && v16)))
                                fmt.Println(v16, v3, v16)
                                if (-128 > 0) {
                                    var v39 int32 = genFn0()
                                    v16 = v16
                                    var v40 float64 = 3.14159
                                    _ = v39
                                    _ = v40
                                }
                                _ = fi38
                            }
                        }
                        _ = v28
                    } else {
                        v16 = (!v16)
                        for fi41 := 0; fi41 < 12; fi41++ {
                            v42 := GenS2{F0: -298}
                            _ = fi41
                            _ = v42
                        }
                    }
                } else {
                    v43 := GenS1{F0: v16, F1: -18}
                    v44 := GenS4{F0: 941, F1: v23, F2: v16}
                    _ = v43
                    _ = v44
                }
                _ = fi27
            }
        }
    case -1:
        if (!v16) {
            if ((9223372036854775807 > 349) || (v16 || true)) {
                for fi45 := 0; fi45 < 3; fi45++ {
                    if v16 {
                        switch 4294967295 {
                        case -878:
                            v3 = 128
                        case 128:
                            v46 := GenS2{F0: -32768}
                            _ = v46
                        }
                    }
                    var v47 string = "hello"
                    _ = fi45
                    _ = v47
                }
                delete(v1, 2147483648)
                if false {
                    for fi48 := 0; fi48 < 18; fi48++ {
                        var v49 uint8 = (v3 / ((v3) | 1))
                        _ = fi48
                        _ = v49
                    }
                    fmt.Println(v3, v23, v16)
                    var v50 int8 = 15
                    _ = v50
                } else {
                    var v51 string = "a"
                    var v52 bool = v16
                    if v16 {
                        if false {
                            if v52 {
                                v53 := GenS5{F0: 1}
                                _ = v53
                            } else {
                                var v54 int32 = (v2 & v2)
                                v52 = (2147483647 <= 524)
                                v55 := GenS6{F0: 255}
                                _ = v54
                                _ = v55
                            }
                            fmt.Println(v16, v16)
                            v16 = v52
                        } else {
                            v56 := GenS3{F0: 106, F1: (((v51 + ((((v51 + "a") + ("hello" + ("foo bar" + ((v51 + "hello") + (v51 + "hello"))))) + v51) + v51)) + ("" + (v51 + ""))) + v51), F2: 248}
                            switch "\n\t" {
                            case "hello":
                                v51 = v51
                            case "\n\t":
                                if v16 {
                                    var v57 int16 = -32768
                                    fmt.Println(v23)
                                    for fi58 := 0; fi58 < 12; fi58++ {
                                        if (1.5 == (3.14159 + 1e10)) {
                                            v52 = (!v52)
                                            v59 := GenS7{F0: -454}
                                            var v60 float64 = float64(0.1)
                                            _ = v59
                                            _ = v60
                                        } else {
                                            if (v52 && v52) {
                                                if ((v16 && ((!v16) && (v23 <= 422))) || (2 < 225)) {
                                                    if (v16 && (!false)) {
                                                        if v16 {
                                                            fmt.Println(v2, v23)
                                                            if v16 {
                                                                fmt.Println(v3)
                                                            }
                                                        }
                                                        for fi61 := 0; fi61 < 9; fi61++ {
                                                            fmt.Println(fi61, v52)
                                                            v62 := GenS8{F0: (v23 - (v23 >> ((255) & 63))), F1: 1, F2: 42}
                                                            if true {
                                                                var v63 float32 = 3.14159
                                                                _ = v63
                                                            }
                                                            var v64 int16 = v57
                                                            _ = fi61
                                                            _ = v62
                                                            _ = v64
                                                        }
                                                    } else {
                                                        fmt.Println(v52, fi58, fi58)
                                                        fmt.Println(v57, v52)
                                                        var v65 int64 = 536
                                                        _ = v65
                                                    }
                                                    for fi66 := 0; fi66 < 10; fi66++ {
                                                        var v67 uint16 = 42
                                                        var v68 int16 = (v57 | -564)
                                                        _ = fi66
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                } else {
                                                    v3 -= v3
                                                    var v69 float64 = 1e10
                                                    _ = v69
                                                }
                                                var v70 int64 = 0
                                                _ = v70
                                            } else {
                                                for fi71 := 0; fi71 < 2; fi71++ {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if ((!false) || false) {
                                                            v51 = "\n\t"
                                                        } else {
                                                            var v72 int8 = 100
                                                            v2 -= (v2 % (((v2 ^ v2)) | 1))
                                                            var v73 uint32 = 42
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    case "a":
                                                        switch fi58 {
                                                        case -443:
                                                            if true {
                                                                var v74 int = -9223372036854775808
                                                                v75 := GenS1{F0: true, F1: 101}
                                                                var v76 uint16 = (v23 | v23)
                                                                v57 = 1000
                                                                var v77 uint16 = 543
                                                                _ = v74
                                                                _ = v75
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                        case 4294967295:
                                                            v78 := []int{}
                                                            _ = v78
                                                        case -600:
                                                            if true {
                                                                v79 := GenS8{F0: 2, F1: (v57 & (v57 << ((2) & 63))), F2: 42}
                                                                if v52 {
                                                                    v2 += (v2 & (-(v2)))
                                                                    var v80 int16 = -1
                                                                    _ = v80
                                                                }
                                                                v57 = v57
                                                                v81 := GenS9{F0: 1}
                                                                _ = v79
                                                                _ = v81
                                                            } else {
                                                                v16 = ((42 > (v57 - (v57 % ((1) | 1)))) && false)
                                                                var v82 float64 = ((1.5 / ((1.5 / ((float64(3.14159) + -2.25) - 0.1)) / 3.14159)) - float64((0.0 * (float32(1e10) - 0.1))))
                                                                v23 = (v23 + v23)
                                                                _ = v82
                                                            }
                                                            var v83 string = ("foo bar" + (v51 + ("hello" + ((v51 + ("\n\t" + v51)) + v51))))
                                                            var v84 bool = false
                                                            _ = v83
                                                            _ = v84
                                                        default:
                                                            v51 = (v51 + v51)
                                                        }
                                                    default:
                                                        var v85 uint8 = v3
                                                        _ = v85
                                                    }
                                                    v16 = v16
                                                    v86 := GenS3{F0: 42, F1: v51, F2: v3}
                                                    _ = fi71
                                                    _ = v86
                                                }
                                            }
                                        }
                                        v2 = (v2 << ((uint(v3)) & 63))
                                        _ = fi58
                                    }
                                    _ = v57
                                } else {
                                    v1[1] = ""
                                    v87 := []int{32768, 343, -72, 128}
                                    _ = v87
                                }
                            case "":
                                var v88 int32 = (v2 << ((238) & 63))
                                _ = v88
                            default:
                                if (-563 < 4294967295) {
                                    v2 = v2
                                } else {
                                    v52 = v16
                                    if ((-9223372036854775808 >= 9223372036854775807) && false) {
                                        var v89 float64 = 1e10
                                        _ = v89
                                    }
                                    v90 := v1[2147483647]
                                    _ = v90
                                }
                            }
                            switch v51 {
                            case "a":
                                var v91 string = "foo bar"
                                _ = v91
                            case "foo bar":
                                v92 := GenS10{F0: (v2 % (((^(v2))) | 1))}
                                _ = v92
                            case "hello":
                                if v52 {
                                    for fi93 := 0; fi93 < 11; fi93++ {
                                        switch fi93 {
                                        case -137:
                                            v2 += (v2 | 466)
                                        case -1:
                                            v94 := []int{fi93, fi93, -805, fi93}
                                            _ = v94
                                        case 32767:
                                            v23 -= (v23 % (((v23 ^ (^(v23)))) | 1))
                                        default:
                                            var v95 int32 = (-(v2))
                                            _ = v95
                                        }
                                        var v96 int16 = 32767
                                        _ = fi93
                                        _ = v96
                                    }
                                    v16 = v52
                                    v97 := []int{-2147483648, 88, 4294967295}
                                    _ = v97
                                } else {
                                    var v98 float64 = (-2.25 / 1.5)
                                    v99 := []int{9223372036854775807, -314}
                                    var v100 uint64 = uint64(v23)
                                    _ = v100
                                    _ = v98
                                    _ = v99
                                }
                            default:
                                v101 := []int{42, -2147483648, 42}
                                _ = v101
                            }
                            _ = v56
                        }
                        switch -46 {
                        case 9223372036854775807:
                            fmt.Println(v16, v51, v3)
                        default:
                            for fi102 := 0; fi102 < 1; fi102++ {
                                switch (fi102 - (fi102 ^ fi102)) {
                                case 32768:
                                    v23 = uint16(v3)
                                case -9223372036854775808:
                                    if (!(1 <= -207)) {
                                        fmt.Println(v52, v2)
                                    } else {
                                        fmt.Println(v23, v3)
                                        v103 := GenS11{F0: uint8(v23)}
                                        _ = v103
                                    }
                                }
                                _ = fi102
                            }
                        }
                    }
                    _ = v51
                    _ = v52
                }
            }
            var v104 uint64 = 0
            _ = v104
        } else {
            v1[2] = "a"
            v105 := v1[65535]
            var v106 uint8 = 255
            _ = v105
            _ = v106
        }
    }
    v107 := GenS7{F0: 1}
    var v108 uint16 = 669
    v109 := GenS9{F0: 42}
    if false {
        var v110 uint16 = (v108 % (((v108 << ((0) & 63))) | 1))
        fmt.Println(v23)
        _ = len(v1)
        _ = v110
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v108:", v108)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v107
    _ = v108
    _ = v109
    _ = v16
    _ = v2
    _ = v23
    _ = v3
}
