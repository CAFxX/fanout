package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 int64; F1 int8; F2 int64 }
type GenS7 struct { F0 int16; F1 uint8; F2 int32 }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int16; F1 int }
type GenS13 struct { F0 string; F1 uint8 }
type GenS14 struct { F0 uint32 }
func genFn0(p0 uint, p1 int32) uint32 {
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: 127, F1: 2}
        var v8 uint8 = 128
        if true {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint64 = 4294967295
                p1 -= (^(p1))
                var v11 string = "hello"
                v12 := []int{-832, fi6, fi9, fi9}
                var v13 uint16 = 0
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
                _ = v13
            }
            if ((v8 % (((v8 << (((^(p0))) & 63))) | 1)) == v8) {
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 uint32 = 2147483647
                    v16 := GenS1{F0: 18, F1: 158}
                    if true {
                        var v17 float64 = (1.5 / -2.25)
                        switch (fi14 - (-(fi6))) {
                        case -254:
                            switch ((("hello" + "\n\t") + "") + "\n\t") {
                            case "\n\t":
                                p1 = p1
                            case "foo bar":
                                v18 := []int{fi6, fi14, (-(fi14))}
                                _ = v18
                            }
                            var v19 float32 = (1.5 / (3.14159 / (0.1 / 1.0)))
                            v8 -= 127
                            _ = v19
                        default:
                            var v20 uint32 = (v15 + uint32(p0))
                            _ = v20
                        }
                        _ = v17
                    } else {
                        if false {
                            if false {
                                v21 := GenS2{F0: 362, F1: fi6, F2: 0}
                                var v22 string = ("" + ((("a" + ("foo bar" + ("hello" + "hello"))) + "hello") + "a"))
                                _ = v21
                                _ = v22
                            }
                            if (1.5 < 1.5) {
                                if (!(695 == 2)) {
                                    v8 = (^(v8))
                                    var v23 int = (fi6 % (((fi6 & (fi14 - (fi6 | -644)))) | 1))
                                    _ = v23
                                } else {
                                    var v24 uint8 = (v8 / ((1) | 1))
                                    _ = v24
                                }
                                v8 -= 29
                                v25 := GenS1{F0: 42, F1: 787}
                                _ = v25
                            } else {
                                v8 = uint8(v15)
                                p1 = p1
                            }
                            v26 := GenS1{F0: -1, F1: 499}
                            _ = v26
                        }
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v27 uint = 1
    v28 := GenS2{F0: -603, F1: 32768, F2: 671}
    for fi29 := 0; fi29 < 1; fi29++ {
        v30 := []int{fi29, -383, 9223372036854775807, fi29}
        for fi31 := 0; fi31 < 8; fi31++ {
            p1 += (p1 / ((-1) | 1))
            if (false || (!true)) {
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
            }
            switch "" {
            case "foo bar":
                v32 := GenS2{F0: 0, F1: fi31, F2: 32767}
                _ = v32
            case "":
                var v33 bool = false
                _ = v33
            default:
                switch ("foo bar" + ("hello" + "\n\t")) {
                case "foo bar":
                    if true {
                        v34 := GenS2{F0: 568, F1: fi29, F2: 1}
                        _ = v34
                    } else {
                        p0 -= v27
                        fmt.Println(fi31)
                    }
                case "a":
                    switch ("" + "\n\t") {
                    case "\n\t":
                        if (1 > -846) {
                            _ = len(v30)
                        }
                    case "foo bar":
                        p0 = (p0 / (((v27 << (((v27 & p0)) & 63))) | 1))
                    }
                default:
                    v35 := GenS1{F0: 0, F1: 4294967295}
                    _ = v35
                }
            }
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return 145
}
func genFn1(p0 int16) uint8 {
    switch -141 {
    case 42:
        switch "hello" {
        case "foo bar":
            fmt.Println(p0)
        case "\n\t":
            v36 := []int{}
            _ = v36
        default:
            switch 128 {
            case 128:
                if true {
                    for fi37 := 0; fi37 < 17; fi37++ {
                        var v38 bool = (43 > 42)
                        if (p0 < -22) {
                            v39 := GenS3{F0: 4294967295, F1: "foo bar"}
                            fmt.Println(p0)
                            _ = v39
                        }
                        _ = fi37
                        _ = v38
                    }
                    var v40 uint = 1
                    var v41 uint64 = 1
                    _ = v40
                    _ = v41
                } else {
                    p0 = p0
                }
            default:
                v42 := GenS4{F0: 0}
                _ = v42
            }
        }
    case -374:
        for fi43 := 0; fi43 < 18; fi43++ {
            switch fi43 {
            case 42:
                if true {
                    var v44 uint16 = 0
                    v45 := []int{}
                    v44 = (v44 - 162)
                    _ = v44
                    _ = v45
                } else {
                    for fi46 := 0; fi46 < 2; fi46++ {
                        v47 := []int{fi46, fi46, fi43}
                        _ = fi46
                        _ = v47
                    }
                }
            }
            var v48 int8 = -128
            if (false && (1 == 4294967295)) {
                v49 := GenS3{F0: 4294967295, F1: "foo bar"}
                var v50 int16 = p0
                for fi51 := 0; fi51 < 20; fi51++ {
                    fmt.Println(p0)
                    _ = fi51
                }
                _ = v49
                _ = v50
            } else {
                v48 = (v48 & (v48 * v48))
            }
            _ = fi43
            _ = v48
        }
    default:
        p0 = (p0 / (((p0 - p0)) | 1))
    }
    fmt.Println(p0)
    if true {
        var v52 int8 = 51
        _ = v52
    }
    v53 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = v53
    return 255
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 617
    var v3 bool = false
    v4 := []string{"hello"}
    var v5 uint32 = 288
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    var v54 int32 = v1
    if (1 == 18446744073709551615) {
        if false {
            for fi55 := 0; fi55 < 16; fi55++ {
                if v3 {
                    v56 := GenS2{F0: 0, F1: fi55, F2: (v2 + uint16(v5))}
                    _ = v56
                }
                v57 := GenS4{F0: uint64(v2)}
                for fi58 := 0; fi58 < 14; fi58++ {
                    switch "\n\t" {
                    case "\n\t":
                        v3 = (!(354 < (fi55 + fi55)))
                    case "a":
                        v4 = append(v4, "hello")
                    case "foo bar":
                        switch "" {
                        case "hello":
                            if (true || false) {
                                var v59 int64 = -845
                                switch "foo bar" {
                                case "foo bar":
                                    v54 = int32(v59)
                                case "hello":
                                    switch fi58 {
                                    case 0:
                                        v60 := GenS5{F0: 805, F1: (v2 + (v2 + (v2 / (((v2 & (v2 / (((v2 * v2)) | 1)))) | 1))))}
                                        _ = v60
                                    case 746:
                                        var v61 int8 = 0
                                        _ = v61
                                    case -1:
                                        v3 = v3
                                    default:
                                        var v62 string = ""
                                        _ = v62
                                    }
                                case "\n\t":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    switch "foo bar" {
                                    case "\n\t":
                                        v59 += (-(v59))
                                    default:
                                        v63 := GenS6{F0: v59, F1: 42, F2: -9223372036854775808}
                                        _ = v63
                                    }
                                }
                                switch (^(fi55)) {
                                case 522:
                                    v2 = (v2 | 42)
                                case -589:
                                    v2 = 42
                                }
                                v4 = append(v4, "foo bar")
                                _ = v59
                            } else {
                                switch ("a" + "foo bar") {
                                case "foo bar":
                                    v3 = (99 == 97)
                                case "\n\t":
                                    var v64 bool = ((!true) || false)
                                    _ = v64
                                case "":
                                    v65 := GenS7{F0: -1, F1: 255, F2: (^(v54))}
                                    _ = v65
                                default:
                                    v4 = append(v4, "")
                                }
                            }
                        case "":
                            if (3.14159 >= -2.25) {
                                v1 = (v54 - 7)
                                switch "hello" {
                                case "hello":
                                    var v66 int32 = (v1 % (((^(v1))) | 1))
                                    _ = v66
                                case "\n\t":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    switch (fi58 / ((fi58) | 1)) {
                                    case 874:
                                        switch 839 {
                                        case 0:
                                            var v67 int16 = 151
                                            _ = v67
                                        case 4294967295:
                                            var v68 int = -9223372036854775808
                                            _ = v68
                                        case -2147483648:
                                            var v69 float32 = 0.1
                                            _ = v69
                                        default:
                                            v4 = append(v4, "a")
                                        }
                                    case 42:
                                        fmt.Println(fi55, v2, v54)
                                    default:
                                        v1 = v1
                                    }
                                }
                                if (0.0 == 3.14159) {
                                    v1 = v54
                                } else {
                                    v70 := GenS8{F0: -9223372036854775808, F1: -9223372036854775808}
                                    v2 = v2
                                    _ = v70
                                }
                                v2 += 2
                            }
                        case "foo bar":
                            v5 = 65535
                        }
                    default:
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                    }
                    if (1 != int64(v5)) {
                        var v71 uint16 = 1
                        if v3 {
                            v72 := GenS9{F0: (3.14159 < 1.5), F1: 42, F2: 42}
                            _ = len(v4)
                            _ = v72
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            switch fi55 {
                            case 2147483647:
                                if len(v4) > 0 {
                                    v4[0] = "foo bar"
                                }
                            case 128:
                                var v73 string = ""
                                _ = v73
                            }
                        }
                        _ = len(v4)
                        _ = v71
                    }
                    _ = fi58
                }
                _ = fi55
                _ = v57
            }
        } else {
            switch "hello" {
            case "":
                switch "hello" {
                case "\n\t":
                    var v74 uint8 = 65
                    _ = v74
                case "a":
                    v54 -= (-(v54))
                case "foo bar":
                    for fi75 := 0; fi75 < 4; fi75++ {
                        v4 = append(v4, "hello")
                        _ = fi75
                    }
                default:
                    switch "foo bar" {
                    case "a":
                        for fi76 := 0; fi76 < 10; fi76++ {
                            fmt.Println(v5, v3, v1)
                            _ = fi76
                        }
                    case "\n\t":
                        v77 := GenS5{F0: 42, F1: 42}
                        _ = v77
                    case "":
                        v2 = (v2 + (v2 & 65535))
                    default:
                        for fi78 := 0; fi78 < 6; fi78++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v79 float64 = (1.5 * ((0.0 / 3.14159) / 3.14159))
                            _ = fi78
                            _ = v79
                        }
                    }
                }
            default:
                for fi80 := 0; fi80 < 6; fi80++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v81 string = "foo bar"
                    _ = fi80
                    _ = v81
                }
            }
            switch ("foo bar" + "") {
            case "a":
                var v82 int32 = (v1 * v1)
                _ = v82
            }
            var v83 uint32 = uint32(v2)
            _ = v83
        }
    }
    for fi84 := 0; fi84 < 10; fi84++ {
        var v85 int = (fi84 - 1)
        if (!(false && v3)) {
            for fi86 := 0; fi86 < 7; fi86++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 -= v2
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v5 = v5
                _ = fi86
            }
        }
        var v87 string = "foo bar"
        _ = fi84
        _ = v85
        _ = v87
    }
    if true {
        v88 := GenS5{F0: 730, F1: 277}
        for fi89 := 0; fi89 < 12; fi89++ {
            v90 := GenS10{F0: v3}
            v54 += v1
            v54 = (v54 ^ (v1 / (((v1 << ((313) & 63))) | 1)))
            _ = fi89
            _ = v90
        }
        _ = v88
    } else {
        var v91 int = 933
        fmt.Println(v2, v5)
        var v92 uint = 65535
        _ = v91
        _ = v92
    }
    var v93 int8 = 48
    if v3 {
        if true {
            v54 -= v54
        } else {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            v54 += (v1 / ((v54) | 1))
        }
    } else {
        var v94 uint8 = 127
        var v95 float64 = (-2.25 - ((float64(3.14159) * (0.0 - 0.1)) - ((-2.25 + 0.1) + 3.14159)))
        if true {
            v96 := GenS11{F0: (^(v2))}
            var v97 string = "foo bar"
            _ = v96
            _ = v97
        }
        _ = v94
        _ = v95
    }
    fmt.Println(v1, v3)
    v3 = ((2 <= 4294967295) || true)
    for fi98 := 0; fi98 < 1; fi98++ {
        for fi99 := 0; fi99 < 6; fi99++ {
            if true {
                switch fi99 {
                case 32768:
                    v3 = true
                case 1:
                    var v100 int8 = (v93 >> ((679) & 63))
                    if v3 {
                        v101 := GenS12{F0: 1000, F1: 732}
                        _ = v101
                    } else {
                        v102 := GenS13{F0: "", F1: 34}
                        var v103 int64 = -9223372036854775808
                        var v104 uint8 = 127
                        _ = v102
                        _ = v103
                        _ = v104
                    }
                    var v105 int8 = int8(v5)
                    _ = v100
                    _ = v105
                default:
                    v5 = v5
                }
            } else {
                var v106 int64 = -9223372036854775808
                v107 := GenS14{F0: genFn0(953, v54)}
                _ = v106
                _ = v107
            }
            _ = fi99
        }
        var v108 int64 = -962
        if true {
            var v109 int64 = 9223372036854775807
            var v110 uint8 = 0
            _ = v109
            _ = v110
        }
        _ = fi98
        _ = v108
    }
    v111 := GenS10{F0: v3}
    fmt.Println(v5, v93)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v93
}
