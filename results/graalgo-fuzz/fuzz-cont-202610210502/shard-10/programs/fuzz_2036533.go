package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 string; F1 bool; F2 int32 }
type GenS4 struct { F0 int8; F1 int32; F2 int }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 uint32; F2 uint64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64; F1 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint8; F1 bool }
type GenS14 struct { F0 uint64; F1 bool; F2 int64 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 string; F1 int16; F2 int64 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int32; F1 uint32; F2 uint32 }
func genFn0(p0 uint32, p1 int64, p2 uint8) int8 {
    v4 := []int{-9223372036854775808}
    v5 := GenS1{F0: p2, F1: 852}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -128
}

func main() {
    v1 := []string{"foo bar", "foo bar"}
    v2 := make(map[string]bool)
    var v3 uint = 255
    var v6 int64 = -1
    var v7 uint64 = 940
    switch "hello" {
    case "a":
        v8 := GenS2{F0: 1, F1: v7}
        _ = v8
    case "foo bar":
        if false {
            v6 = v6
        }
    case "hello":
        var v9 int = 0
        _ = v9
    default:
        var v10 uint16 = 54
        _ = v10
    }
    var v11 float64 = 0.0
    v12 := GenS3{F0: (("a" + ((("hello" + "foo bar") + "") + (("" + "hello") + ""))) + "\n\t"), F1: true, F2: -1}
    v7 = 207
    switch 503 {
    case 32767:
        v13 := GenS4{F0: 127, F1: -1, F2: 127}
        _ = v13
    }
    v6 -= v6
    switch "" {
    case "\n\t":
        delete(v2, ("foo bar" + "\n\t"))
    default:
        fmt.Println(v11, v7, v3)
    }
    v7 += (v7 & (v7 * (v7 / ((v7) | 1))))
    var v14 int8 = 1
    if true {
        delete(v2, "\n\t")
        v1 = append(v1, ("\n\t" + (("" + "foo bar") + "\n\t")))
    } else {
        v14 -= (v14 << (((^(v3))) & 63))
        if false {
            if false {
                v7 = v7
                var v15 float32 = 3.14159
                _ = v15
            } else {
                if ((-2.25 - v11) < 0.1) {
                    var v16 bool = ((-(v14)) == (v14 * int8(v7)))
                    fmt.Println(v3)
                    _ = v16
                } else {
                    switch 32767 {
                    case 32768:
                        if (180 < 127) {
                            for fi17 := 0; fi17 < 14; fi17++ {
                                if ((1 <= (v7 << (((v3 * (v3 / ((18446744073709551615) | 1)))) & 63))) || false) {
                                    v11 += v11
                                }
                                _ = fi17
                            }
                            v18 := v2[("\n\t" + "\n\t")]
                            _ = v18
                        }
                    default:
                        if (!(true && (!(0.1 < float32(v11))))) {
                            if (true && false) {
                                v7 -= (v7 / ((501) | 1))
                                var v19 float32 = 0.1
                                v6 -= -9223372036854775808
                                _ = v19
                            } else {
                                for fi20 := 0; fi20 < 8; fi20++ {
                                    v11 += -2.25
                                    v6 += -9223372036854775808
                                    _ = fi20
                                }
                                for fi21 := 0; fi21 < 15; fi21++ {
                                    v22 := GenS5{F0: v14, F1: 0}
                                    var v23 uint16 = 741
                                    v6 -= v6
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                                for fi24 := 0; fi24 < 3; fi24++ {
                                    for fi25 := 0; fi25 < 16; fi25++ {
                                        delete(v2, "hello")
                                        var v26 bool = true
                                        v7 -= v7
                                        _ = fi25
                                        _ = v26
                                    }
                                    v14 = v14
                                    for fi27 := 0; fi27 < 20; fi27++ {
                                        v28 := GenS5{F0: (v14 - (v14 % ((v14) | 1))), F1: 255}
                                        _ = fi27
                                        _ = v28
                                    }
                                    _ = fi24
                                }
                            }
                            for fi29 := 0; fi29 < 11; fi29++ {
                                var v30 float64 = ((v11 / ((3.14159 / v11) / v11)) + (((0.1 - 0.1) / (-2.25 - (-2.25 * v11))) * float64(1e10)))
                                _ = fi29
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 14; fi31++ {
                                if false {
                                    switch "a" {
                                    case "hello":
                                        if (42 > 2147483648) {
                                            if false {
                                                fmt.Println(fi31)
                                            }
                                            var v32 int32 = -11
                                            _ = v32
                                        } else {
                                            fmt.Println(fi31, v7, v6)
                                        }
                                    case "foo bar":
                                        for fi33 := 0; fi33 < 6; fi33++ {
                                            _ = len(v1)
                                            var v34 uint64 = v7
                                            var v35 float64 = (0.1 - -2.25)
                                            _ = fi33
                                            _ = v34
                                            _ = v35
                                        }
                                    default:
                                        var v36 float32 = 3.14159
                                        _ = v36
                                    }
                                    fmt.Println(fi31)
                                    v1 = append(v1, "\n\t")
                                } else {
                                    v14 += (^(v14))
                                    fmt.Println(v11, fi31)
                                    for fi37 := 0; fi37 < 14; fi37++ {
                                        var v38 uint32 = 462
                                        if (v38 >= v38) {
                                            var v39 uint64 = (v7 - 4294967295)
                                            v40 := GenS6{F0: 214}
                                            _ = v39
                                            _ = v40
                                        } else {
                                            v3 = v3
                                            fmt.Println(v6, v3, fi37)
                                            v6 = 9223372036854775807
                                            if (32767 >= -32768) {
                                                if false {
                                                    if true {
                                                        var v41 uint64 = (v7 / (((v7 >> (((v3 | v3)) & 63))) | 1))
                                                        var v42 int64 = -1
                                                        switch 32768 {
                                                        case 128:
                                                            v43 := GenS6{F0: uint8(v3)}
                                                            _ = v43
                                                        case -1:
                                                            v42 -= 9223372036854775807
                                                        }
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    if true {
                                                        var v44 uint32 = v38
                                                        fmt.Println(v6)
                                                        fmt.Println(fi31, v3)
                                                        _ = v44
                                                    }
                                                }
                                            } else {
                                                switch "\n\t" {
                                                case "a":
                                                    _ = len(v2)
                                                case "":
                                                    var v45 uint8 = 128
                                                    _ = v45
                                                default:
                                                    var v46 string = "a"
                                                    _ = v46
                                                }
                                                switch ("\n\t" + ("a" + "hello")) {
                                                case "\n\t":
                                                    v47 := GenS5{F0: (v14 % (((v14 - genFn0(v38, -364, 1))) | 1)), F1: 255}
                                                    _ = v47
                                                case "foo bar":
                                                    v3 += v3
                                                default:
                                                    var v48 uint16 = 804
                                                    _ = v48
                                                }
                                                fmt.Println(v38, v38)
                                            }
                                            v49 := GenS7{F0: (!false)}
                                            v1 = append(v1, ("foo bar" + ("a" + "a")))
                                            if len(v1) > 0 {
                                                v1[0] = "foo bar"
                                            }
                                            _ = v49
                                        }
                                        _ = fi37
                                        _ = v38
                                    }
                                }
                                _ = fi31
                            }
                        } else {
                            fmt.Println(v11, v7, v3)
                            v1 = append(v1, "a")
                            var v50 float64 = v11
                            _ = v50
                        }
                    }
                    v6 = (v6 / ((v6) | 1))
                    v3 = (v3 << (((v3 + (^(v3)))) & 63))
                }
            }
            if (!(!(513 <= 2147483648))) {
                var v51 uint64 = v7
                var v52 uint8 = 181
                if false {
                    switch ("" + "\n\t") {
                    case "hello":
                        v6 = (v6 ^ 1)
                    case "a":
                        if ((1.5 + -2.25) == (0.0 / 1e10)) {
                            for fi53 := 0; fi53 < 20; fi53++ {
                                if ("foo bar" < "a") {
                                    v6 -= -9223372036854775808
                                    if false {
                                        var v54 float32 = 1e10
                                        _ = v54
                                    } else {
                                        if true {
                                            if false {
                                                switch (fi53 ^ 4294967295) {
                                                case 1:
                                                    for fi55 := 0; fi55 < 14; fi55++ {
                                                        if (!true) {
                                                            fmt.Println(fi55)
                                                        }
                                                        v11 -= v11
                                                        _ = fi55
                                                    }
                                                case -2147483648:
                                                    switch "a" {
                                                    case "":
                                                        if true {
                                                            for fi56 := 0; fi56 < 13; fi56++ {
                                                                if true {
                                                                    var v57 uint32 = uint32(v7)
                                                                    if (v6 <= (-(v6))) {
                                                                        if true {
                                                                            v14 = v14
                                                                        }
                                                                    } else {
                                                                        var v58 int8 = (v14 & v14)
                                                                        _ = len(v1)
                                                                        v51 += (v7 & uint64(v57))
                                                                        _ = v58
                                                                    }
                                                                    _ = v57
                                                                } else {
                                                                    v59 := GenS4{F0: (v14 % ((genFn0(497, v6, (v52 ^ (^(v52))))) | 1)), F1: 785, F2: fi53}
                                                                    v2["hello"] = false
                                                                    _ = v59
                                                                }
                                                                _ = fi56
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        } else {
                                                            _ = len(v1)
                                                        }
                                                    case "\n\t":
                                                        v60 := GenS8{F0: fi53, F1: 2147483647, F2: (v7 - 2)}
                                                        _ = v60
                                                    default:
                                                        fmt.Println(v7)
                                                    }
                                                default:
                                                    v14 = (v14 << (((v3 / ((v3) | 1))) & 63))
                                                }
                                            } else {
                                                fmt.Println(v6, v51, v3)
                                                if (false && true) {
                                                    var v61 bool = (false && (0 != 65535))
                                                    for fi62 := 0; fi62 < 1; fi62++ {
                                                        fmt.Println(fi53, v11)
                                                        _ = len(v1)
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v63 uint = v3
                                                            _ = v63
                                                        case "hello":
                                                            var v64 uint16 = 3
                                                            fmt.Println(fi53, v6)
                                                            fmt.Println(v14, fi62, v7)
                                                            var v65 float64 = v11
                                                            v66 := GenS3{F0: ("\n\t" + "a"), F1: (((25 < uint32(v3)) && ((!v61) && (v6 != (v6 / ((v6) | 1))))) || ((v6 ^ 42) >= v6)), F2: 0}
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        case "":
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v67 int32 = 0
                                                                _ = v67
                                                            case "\n\t":
                                                                v1 = append(v1, ("hello" + "foo bar"))
                                                            case "":
                                                                if false {
                                                                    if ((3.14159 * 3.14159) >= 3.14159) {
                                                                        v6 = v6
                                                                        v68 := GenS9{F0: v14}
                                                                        _ = v68
                                                                    } else {
                                                                        v1 = append(v1, ("hello" + "hello"))
                                                                    }
                                                                }
                                                                if false {
                                                                    v52 -= v52
                                                                    fmt.Println(v7)
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        if (!false) {
                                                                            fmt.Println(v52)
                                                                        }
                                                                    case "\n\t":
                                                                        v69 := GenS1{F0: uint8(v51), F1: 788}
                                                                        _ = v69
                                                                    }
                                                                }
                                                                var v70 bool = true
                                                                _ = v70
                                                            default:
                                                                var v71 uint = 42
                                                                _ = v71
                                                            }
                                                        }
                                                        _ = fi62
                                                    }
                                                    _ = v61
                                                }
                                                _ = len(v1)
                                            }
                                            for fi72 := 0; fi72 < 15; fi72++ {
                                                _ = len(v1)
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                                _ = fi72
                                            }
                                            var v73 int8 = (v14 % (((^(v14))) | 1))
                                            _ = v73
                                        } else {
                                            if false {
                                                if true {
                                                    v51 += (v51 | 42)
                                                }
                                            }
                                            v2["\n\t"] = (((1e10 - (0.1 * ((1.5 + 1.5) - float32(1e10)))) - (0.0 + ((-2.25 / 1e10) + float32((((v11 * ((v11 * 1e10) * (v11 * (0.0 - ((((1.5 - ((v11 / ((v11 / v11) / float64((-2.25 + (((0.0 + 0.1) / 3.14159) + (1e10 * 0.1)))))) * v11)) + v11) * ((((v11 / 0.1) - 1.5) - v11) + float64(3.14159))) * (((v11 + float64((1.5 - ((1.5 / 1e10) - (((0.1 - 3.14159) / 1e10) * ((3.14159 * 3.14159) * 0.0)))))) * v11) + (v11 * ((v11 - v11) / 1.5)))))))) - (((v11 + (((v11 / -2.25) + (v11 * (0.0 * 0.1))) * 1.5)) / v11) - (float64(-2.25) / v11))) * ((v11 - ((v11 + (v11 + v11)) + (3.14159 + v11))) + 0.1)))))) < -2.25)
                                            switch fi53 {
                                            case -9223372036854775808:
                                                v74 := GenS8{F0: 1, F1: 979, F2: v51}
                                                _ = v74
                                            default:
                                                var v75 float32 = -2.25
                                                _ = v75
                                            }
                                        }
                                    }
                                }
                                _ = fi53
                            }
                            if (42 > 417) {
                                if ((4294967295 < 2147483647) || true) {
                                    v76 := GenS9{F0: int8(v52)}
                                    _ = v76
                                } else {
                                    for fi77 := 0; fi77 < 7; fi77++ {
                                        v2[""] = true
                                        if false {
                                            for fi78 := 0; fi78 < 10; fi78++ {
                                                var v79 uint32 = 2147483648
                                                v80 := GenS10{F0: (v52 ^ v52)}
                                                v7 = (v7 & v51)
                                                _ = fi78
                                                _ = v79
                                                _ = v80
                                            }
                                        } else {
                                            for fi81 := 0; fi81 < 20; fi81++ {
                                                _ = len(v1)
                                                fmt.Println(v14, v11, v52)
                                                var v82 float64 = 0.1
                                                v82 += v11
                                                v83 := GenS10{F0: 0}
                                                _ = fi81
                                                _ = v82
                                                _ = v83
                                            }
                                            if (v6 == 1) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                fmt.Println(v14)
                                            }
                                            fmt.Println(v6, v51, v6)
                                        }
                                        _ = fi77
                                    }
                                }
                                v84 := GenS1{F0: v52, F1: 32767}
                                _ = v84
                            } else {
                                v14 = int8(v52)
                                v85 := v2["hello"]
                                _ = v85
                            }
                        } else {
                            if (0 == 42) {
                                if true {
                                    fmt.Println(v11, v11)
                                } else {
                                    v86 := GenS7{F0: (v7 < 0)}
                                    if (!false) {
                                        if (false && true) {
                                            var v87 float64 = (float64(0.1) / -2.25)
                                            if (!(4294967295 == 32767)) {
                                                v51 = (^(v51))
                                                if false {
                                                    fmt.Println(v14, v6)
                                                    var v88 int64 = v6
                                                    _ = v88
                                                } else {
                                                    delete(v2, "hello")
                                                    _ = len(v1)
                                                }
                                            }
                                            _ = v87
                                        } else {
                                            var v89 float64 = v11
                                            v3 = (v3 | v3)
                                            v52 = uint8(v3)
                                            _ = v89
                                        }
                                        var v90 bool = (-9223372036854775808 < (v6 << ((v3) & 63)))
                                        delete(v2, ("hello" + ("" + "hello")))
                                        _ = v90
                                    }
                                    _ = v86
                                }
                            }
                            if false {
                                var v91 uint16 = 280
                                _ = v91
                            } else {
                                if true {
                                    v92 := GenS5{F0: v14, F1: 0}
                                    v93 := GenS11{F0: 1, F1: ("hello" + ("" + ""))}
                                    var v94 int8 = -1
                                    _ = v92
                                    _ = v93
                                    _ = v94
                                }
                                var v95 uint16 = 938
                                fmt.Println(v95)
                                _ = v95
                            }
                            if (32767 <= 42) {
                                for fi96 := 0; fi96 < 7; fi96++ {
                                    v97 := GenS11{F0: v7, F1: ""}
                                    _ = fi96
                                    _ = v97
                                }
                                v98 := GenS12{F0: v3}
                                _ = v98
                            }
                        }
                    case "foo bar":
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "hello":
                                switch 1 {
                                case 127:
                                    switch -223 {
                                    case 42:
                                        for fi99 := 0; fi99 < 15; fi99++ {
                                            var v100 uint = 2
                                            _ = fi99
                                            _ = v100
                                        }
                                    case 128:
                                        for fi101 := 0; fi101 < 16; fi101++ {
                                            for fi102 := 0; fi102 < 16; fi102++ {
                                                var v103 string = ""
                                                _ = fi102
                                                _ = v103
                                            }
                                            _ = fi101
                                        }
                                    }
                                case -161:
                                    switch "a" {
                                    case "hello":
                                        v52 -= (^(v52))
                                    case "foo bar":
                                        var v104 bool = true
                                        _ = v104
                                    }
                                case 740:
                                    var v105 float64 = 1e10
                                    _ = v105
                                default:
                                    v14 -= v14
                                }
                            case "":
                                if (!((v51 >> (((v3 * v3)) & 63)) == (v7 % ((uint64(v52)) | 1)))) {
                                    v3 = 751
                                    if (false || true) {
                                        var v106 uint64 = v51
                                        var v107 int64 = v6
                                        if false {
                                            for fi108 := 0; fi108 < 3; fi108++ {
                                                v109 := GenS13{F0: (v52 - v52), F1: ("a" == "foo bar")}
                                                fmt.Println(v3, v107)
                                                v110 := GenS7{F0: true}
                                                _ = fi108
                                                _ = v109
                                                _ = v110
                                            }
                                            v111 := v2["foo bar"]
                                            fmt.Println(v14)
                                            _ = v111
                                        } else {
                                            v6 = -217
                                            v51 -= 1
                                            for fi112 := 0; fi112 < 12; fi112++ {
                                                v113 := GenS14{F0: (^(v7)), F1: (-720 != -1), F2: 42}
                                                if len(v1) > 0 {
                                                    v1[0] = "\n\t"
                                                }
                                                if true {
                                                    if len(v1) > 0 {
                                                        v1[0] = ("a" + "hello")
                                                    }
                                                } else {
                                                    v114 := GenS12{F0: (^(v3))}
                                                    v115 := v2[((((("hello" + ((("\n\t" + ("" + "a")) + "") + (("foo bar" + ("foo bar" + "hello")) + "\n\t"))) + "") + "foo bar") + "hello") + "\n\t")]
                                                    fmt.Println(v52, v11)
                                                    _ = v114
                                                    _ = v115
                                                }
                                                _ = fi112
                                                _ = v113
                                            }
                                        }
                                        _ = v106
                                        _ = v107
                                    }
                                    v14 = v14
                                } else {
                                    v3 = v3
                                    v14 += v14
                                    fmt.Println(v6, v51)
                                }
                            }
                            var v116 uint = 925
                            _ = v116
                        }
                    }
                    switch "" {
                    case "a":
                        v51 = v51
                    case "hello":
                        v2["hello"] = true
                    case "foo bar":
                        v2["a"] = false
                    }
                }
                _ = v51
                _ = v52
            }
        }
    }
    if true {
        for fi117 := 0; fi117 < 19; fi117++ {
            for fi118 := 0; fi118 < 18; fi118++ {
                v7 -= v7
                _ = fi118
            }
            for fi119 := 0; fi119 < 5; fi119++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (0 > -2147483648) {
                    if ((!false) && (false && false)) {
                        switch (fi117 * (fi119 >> ((v3) & 63))) {
                        case -9223372036854775808:
                            _ = len(v1)
                        case 4294967295:
                            _ = len(v1)
                            switch "" {
                            case "\n\t":
                                if (853 <= -1) {
                                    v120 := GenS15{F0: (v6 % (((v6 % ((v6) | 1))) | 1))}
                                    if false {
                                        v3 += v3
                                        v121 := GenS16{F0: ("" + "foo bar"), F1: 1000, F2: (v6 % (((v6 + v6)) | 1))}
                                        _ = v121
                                    }
                                    _ = v120
                                }
                            case "":
                                v122 := GenS13{F0: 179, F1: false}
                                _ = v122
                            case "a":
                                v123 := GenS15{F0: (v6 << (((^(v3))) & 63))}
                                _ = v123
                            default:
                                if false {
                                    if true {
                                        fmt.Println(v3, v14)
                                    }
                                    fmt.Println(fi119, v11, v7)
                                    switch fi119 {
                                    case -960:
                                        var v124 float64 = v11
                                        _ = v124
                                    case -36:
                                        var v125 float32 = -2.25
                                        _ = v125
                                    default:
                                        if len(v1) > 0 {
                                            v1[0] = "foo bar"
                                        }
                                    }
                                    v3 += v3
                                    v6 -= v6
                                    var v126 float32 = 0.1
                                    switch "hello" {
                                    case "a":
                                        v126 -= (float32(0.1) / v126)
                                    }
                                    switch (-(fi117)) {
                                    case 9223372036854775807:
                                        v127 := GenS15{F0: (v6 / ((-401) | 1))}
                                        _ = v127
                                    default:
                                        v11 = v11
                                    }
                                    _ = v126
                                }
                            }
                            _ = len(v2)
                            var v128 int = (fi117 ^ fi119)
                            var v129 uint8 = 128
                            _ = v128
                            _ = v129
                        case 32767:
                            if (v3 <= (v3 * 18446744073709551615)) {
                                v3 = (v3 / ((v3) | 1))
                                var v130 int16 = 172
                                v14 -= v14
                                _ = v130
                            } else {
                                v11 += v11
                                var v131 int = fi117
                                delete(v2, "\n\t")
                                _ = v131
                            }
                        }
                        v14 += v14
                    }
                    if false {
                        if (18446744073709551615 <= (v7 ^ v7)) {
                            v1 = append(v1, ("" + "hello"))
                            fmt.Println(v6, fi117, v6)
                        }
                        var v132 uint16 = 306
                        v6 += 387
                        _ = v132
                    }
                    _ = len(v2)
                }
                var v133 uint = (v3 - v3)
                fmt.Println(fi117)
                _ = fi119
                _ = v133
            }
            if (!false) {
                v6 -= v6
                _ = len(v2)
            }
            _ = fi117
        }
        var v134 bool = false
        _ = v134
    } else {
        switch "a" {
        case "a":
            v3 += v3
        case "foo bar":
            if true {
                v135 := GenS1{F0: 0, F1: 390}
                var v136 int32 = int32(v6)
                _ = v135
                _ = v136
            }
        }
        switch 213 {
        case 32768:
            v7 += 42
        }
    }
    fmt.Println(v11, v14, v3)
    v137 := v2[""]
    fmt.Println(v11, v7, v6)
    switch 700 {
    case 42:
        for fi138 := 0; fi138 < 4; fi138++ {
            fmt.Println(v7)
            fmt.Println(v7)
            var v139 float32 = 0.1
            _ = fi138
            _ = v139
        }
    case 577:
        switch "a" {
        case "foo bar":
            v3 += v3
        case "":
            if true {
                switch "a" {
                case "foo bar":
                    _ = len(v1)
                case "\n\t":
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                case "":
                    v140 := GenS17{F0: uint16(v7)}
                    _ = v140
                default:
                    var v141 uint16 = 32767
                    _ = v141
                }
            } else {
                for fi142 := 0; fi142 < 20; fi142++ {
                    if v137 {
                        v11 = v11
                        if (!(!(209 > 1))) {
                            v7 += (v7 % ((2) | 1))
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v137 = (fi142 <= int(v6))
                        switch "\n\t" {
                        case "a":
                            switch ("foo bar" + "foo bar") {
                            case "foo bar":
                                switch (-(fi142)) {
                                case -9223372036854775808:
                                    if (v137 && ((((v6 > v6) || (false && v137)) || v137) || true)) {
                                        var v143 uint8 = 42
                                        v14 = 127
                                        var v144 int64 = v6
                                        _ = v143
                                        _ = v144
                                    } else {
                                        v137 = v137
                                    }
                                default:
                                    if v137 {
                                        for fi145 := 0; fi145 < 17; fi145++ {
                                            var v146 uint64 = (v7 / ((v7) | 1))
                                            if ((97 >= 218) && true) {
                                                v147 := GenS11{F0: v146, F1: ""}
                                                _ = v147
                                            }
                                            if (true && ("foo bar" > "foo bar")) {
                                                if v137 {
                                                    v148 := GenS18{F0: -2147483648, F1: 86, F2: 769}
                                                    v149 := GenS4{F0: -55, F1: -2147483648, F2: fi142}
                                                    _ = v148
                                                    _ = v149
                                                }
                                                var v150 int64 = (v6 - v6)
                                                _ = v150
                                            }
                                            _ = fi145
                                            _ = v146
                                        }
                                        fmt.Println(v6)
                                    } else {
                                        var v151 uint = (v3 - v3)
                                        v151 += v151
                                        v3 -= v3
                                        _ = v151
                                    }
                                }
                            case "":
                                v6 -= 9223372036854775807
                            case "\n\t":
                                v152 := GenS18{F0: 0, F1: 496, F2: 2}
                                _ = v152
                            default:
                                v6 = (v6 & 808)
                            }
                        }
                    }
                    _ = fi142
                }
            }
        }
    case 903:
        var v153 uint = (v3 - v3)
        _ = v153
    }
    v14 = (v14 % (((v14 - v14)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v137:", v137)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v137
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
