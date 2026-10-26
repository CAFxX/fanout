package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint16; F1 uint; F2 int16 }
type GenS7 struct { F0 int; F1 bool; F2 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 uint8; F1 bool }
func genFn0(p0 int8) uint {
    v5 := GenS1{F0: "hello", F1: 113}
    var v6 int16 = -499
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            v6 -= (-(v6))
            fmt.Println(fi7)
            _ = fi8
        }
        _ = fi7
    }
    if false {
        p0 += p0
    } else {
        if false {
            var v9 bool = ((-9223372036854775808 >= 2147483647) && (!false))
            _ = v9
        } else {
            v10 := []int{9223372036854775807, -780, 4294967295, 127}
            _ = v10
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 32
}
func genFn1(p0 uint8) bool {
    if false {
        v11 := []int{}
        var v12 int16 = 179
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: ("a" + "\n\t"), F1: 255}
    _ = p0
    _ = v13
    return true
}

func main() {
    var v1 int = -105
    var v2 float32 = (3.14159 - 0.0)
    var v3 bool = (!false)
    var v4 uint16 = 2
    v3 = (135 >= 18446744073709551615)
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 uint8 = 101
        v15 += v15
        _ = fi14
        _ = v15
    }
    var v16 float64 = ((0.1 + 0.0) / (-2.25 * (-2.25 + 1.5)))
    var v17 int = v1
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            v20 := []int{(^(fi19)), v17, (-(v1))}
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = fi19
            _ = v20
        }
        v21 := []int{v17, 677}
        if true {
            if v3 {
                v16 = v16
            } else {
                v22 := GenS1{F0: (("foo bar" + ("\n\t" + ("hello" + (((("foo bar" + "hello") + ("hello" + "foo bar")) + "") + "foo bar")))) + ""), F1: 1}
                switch "foo bar" {
                case "hello":
                    var v23 uint = 42
                    _ = v23
                case "foo bar":
                    var v24 float32 = v2
                    _ = v24
                case "":
                    switch 907 {
                    case 32767:
                        fmt.Println(fi18)
                    case 42:
                        if v3 {
                            switch ("hello" + "foo bar") {
                            case "a":
                                for fi25 := 0; fi25 < 3; fi25++ {
                                    switch "hello" {
                                    case "hello":
                                        v26 := GenS1{F0: "", F1: 412}
                                        _ = v26
                                    case "":
                                        v16 += 1.5
                                    case "foo bar":
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                    }
                                    _ = fi25
                                }
                            case "foo bar":
                                if (852 > 4294967295) {
                                    if ((42 >= -80) || true) {
                                        var v27 uint = 509
                                        v2 -= v2
                                        if (v16 > float64(-2.25)) {
                                            v21 = append(v21, (v17 * 128))
                                        } else {
                                            fmt.Println(fi18)
                                            v2 = (0.0 - (v2 - v2))
                                        }
                                        _ = v27
                                    } else {
                                        for fi28 := 0; fi28 < 11; fi28++ {
                                            var v29 int8 = 42
                                            v1 = -1
                                            if (!(-2147483648 == -167)) {
                                                if v3 {
                                                    switch ("a" + "") {
                                                    case "\n\t":
                                                        v30 := GenS2{F0: 127, F1: 42}
                                                        _ = v30
                                                    default:
                                                        fmt.Println(v1, v16)
                                                    }
                                                    v21 = append(v21, -841)
                                                } else {
                                                    v29 = (v29 % ((127) | 1))
                                                }
                                                if (v2 > ((1e10 + float32(((v16 + ((float64(((0.1 + v2) + (v2 * ((1e10 / (((((v2 / v2) + v2) - (((v2 * v2) - (v2 + float32((v16 - v16)))) - v2)) - v2) / ((v2 + (1.5 + 1e10)) + (1e10 * (v2 * 3.14159))))) / v2)))) * 0.0) * v16)) + v16))) + v2)) {
                                                    switch (-(fi18)) {
                                                    case 32767:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            fmt.Println(v3, v16)
                                                        case "foo bar":
                                                            switch "a" {
                                                            case "hello":
                                                                v29 = -128
                                                            case "a":
                                                                var v31 float64 = v16
                                                                _ = v31
                                                            }
                                                        }
                                                    default:
                                                        switch (fi28 & fi18) {
                                                        case 35:
                                                            v21 = append(v21, 127)
                                                        case -1:
                                                            v29 = (v29 ^ v29)
                                                        }
                                                    }
                                                } else {
                                                    switch (v1 << ((858) & 63)) {
                                                    case 128:
                                                        if len(v21) > 0 {
                                                            v21[0] = int(v4)
                                                        }
                                                    case 32767:
                                                        v32 := GenS3{F0: (v29 >= (v29 << ((65535) & 63)))}
                                                        var v33 uint64 = 4294967295
                                                        if len(v21) > 0 {
                                                            v21[0] = (v1 & v1)
                                                        }
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    fmt.Println(v17)
                                                    if len(v21) > 0 {
                                                        v21[0] = v17
                                                    }
                                                }
                                            } else {
                                                var v34 string = "hello"
                                                if (v3 || (128 < 1)) {
                                                    switch v34 {
                                                    case "\n\t":
                                                        v21 = append(v21, v1)
                                                    case "foo bar":
                                                        v21 = append(v21, (fi18 >> ((633) & 63)))
                                                    case "hello":
                                                        var v35 float32 = v2
                                                        if v3 {
                                                            v3 = v3
                                                            if v3 {
                                                                fmt.Println(v17, v3, v34)
                                                                v3 = true
                                                                v36 := GenS4{F0: ((("a" + v34) + v34) + v34)}
                                                                var v37 uint64 = 0
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                            if true {
                                                                v21 = append(v21, (fi28 & 0))
                                                                v2 = 3.14159
                                                                fmt.Println(v35)
                                                                var v38 string = "foo bar"
                                                                if (v3 && (-9223372036854775808 <= 9223372036854775807)) {
                                                                    if true {
                                                                        fmt.Println(v1, v3, v35)
                                                                        switch fi18 {
                                                                        case 32768:
                                                                            switch 2147483647 {
                                                                            case 0:
                                                                                var v39 int = fi28
                                                                                _ = v39
                                                                            case 127:
                                                                                fmt.Println(v35, v1)
                                                                            case -910:
                                                                                v17 -= 9223372036854775807
                                                                            default:
                                                                                v17 = -1
                                                                            }
                                                                        case 4294967295:
                                                                            v16 += v16
                                                                        }
                                                                    } else {
                                                                        if ((genFn0((v29 >> ((501) & 63))) > 833) && v3) {
                                                                            v38 = v34
                                                                            fmt.Println(v17)
                                                                            switch (fi18 % ((2147483647) | 1)) {
                                                                            case 749:
                                                                                var v40 uint64 = 0
                                                                                _ = v40
                                                                            }
                                                                        }
                                                                        v41 := GenS5{F0: 0, F1: v4}
                                                                        _ = v41
                                                                    }
                                                                    v35 = (float32(float64((v2 - 0.0))) / v35)
                                                                } else {
                                                                    if (65535 != 4294967295) {
                                                                        fmt.Println(v4, v17)
                                                                        var v42 uint8 = 0
                                                                        fmt.Println(v3, v3, v17)
                                                                        _ = v42
                                                                    }
                                                                    v43 := GenS1{F0: "hello", F1: 2}
                                                                    v17 += v1
                                                                    switch (v38 + ("a" + v34)) {
                                                                    case "":
                                                                        var v44 uint32 = 2147483648
                                                                        _ = v44
                                                                    case "hello":
                                                                        fmt.Println(fi18)
                                                                        v45 := GenS1{F0: v34, F1: genFn0((v29 + -1))}
                                                                        _ = v45
                                                                    default:
                                                                        v29 = (v29 + v29)
                                                                    }
                                                                    _ = v43
                                                                }
                                                                if (1 != 255) {
                                                                    if len(v21) > 0 {
                                                                        v21[0] = (fi28 ^ (fi28 ^ fi18))
                                                                    }
                                                                    var v46 uint64 = 2
                                                                    _ = v46
                                                                }
                                                                _ = v38
                                                            }
                                                            switch ("" + v34) {
                                                            case "\n\t":
                                                                v35 += ((v35 / ((-2.25 * 0.0) + v35)) / 1.5)
                                                            case "":
                                                                var v47 int8 = -12
                                                                _ = v47
                                                            }
                                                            v48 := GenS3{F0: (v29 < (v29 | -128))}
                                                            v49 := GenS6{F0: 2, F1: 42, F2: 1000}
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                        v4 = 32767
                                                        _ = v35
                                                    }
                                                }
                                                switch ("hello" + (v34 + v34)) {
                                                case "":
                                                    v50 := GenS4{F0: v34}
                                                    _ = v50
                                                case "hello":
                                                    var v51 uint16 = v4
                                                    v4 = (v51 >> ((610) & 63))
                                                    _ = v51
                                                case "\n\t":
                                                    switch "hello" {
                                                    case "a":
                                                        v1 = (fi28 >> ((2) & 63))
                                                    default:
                                                        fmt.Println(v2)
                                                    }
                                                }
                                                var v52 uint = 42
                                                _ = v34
                                                _ = v52
                                            }
                                            if (v4 != v4) {
                                                v4 = (v4 << ((861) & 63))
                                                if len(v21) > 0 {
                                                    v21[0] = v17
                                                }
                                                var v53 float32 = 0.0
                                                v53 = v53
                                                switch "" {
                                                case "\n\t":
                                                    var v54 uint32 = 357
                                                    _ = v54
                                                case "foo bar":
                                                    v16 -= float64(v53)
                                                default:
                                                    _ = len(v21)
                                                }
                                                _ = v53
                                            } else {
                                                v55 := GenS5{F0: 761, F1: v4}
                                                v2 = v2
                                                v16 = 1.5
                                                _ = v55
                                            }
                                            fmt.Println(fi28, v16)
                                            _ = fi28
                                            _ = v29
                                        }
                                    }
                                    if len(v21) > 0 {
                                        v21[0] = (v1 & (v17 | fi18))
                                    }
                                } else {
                                    v21 = append(v21, -9223372036854775808)
                                    switch v17 {
                                    case -983:
                                        if (v3 && (-1 <= -9223372036854775808)) {
                                            if v3 {
                                                fmt.Println(v1, v4)
                                            } else {
                                                v17 += fi18
                                                var v56 float64 = v16
                                                _ = v56
                                            }
                                        } else {
                                            for fi57 := 0; fi57 < 9; fi57++ {
                                                v58 := GenS1{F0: "", F1: 255}
                                                _ = fi57
                                                _ = v58
                                            }
                                            if (-127 > -106) {
                                                v16 -= (3.14159 / v16)
                                                if (!(1 == -116)) {
                                                    v1 += fi18
                                                    v16 -= 0.1
                                                    switch -9223372036854775808 {
                                                    case 32767:
                                                        fmt.Println(v3)
                                                    case -416:
                                                        fmt.Println(fi18)
                                                    case 32768:
                                                        var v59 int8 = -128
                                                        _ = v59
                                                    }
                                                } else {
                                                    for fi60 := 0; fi60 < 18; fi60++ {
                                                        if len(v21) > 0 {
                                                            v21[0] = (v1 - (^(v1)))
                                                        }
                                                        if v3 {
                                                            v17 += 1
                                                            v1 = fi60
                                                        } else {
                                                            v4 -= 65535
                                                        }
                                                        _ = fi60
                                                    }
                                                    v1 -= (fi18 >> ((948) & 63))
                                                    v4 = v4
                                                }
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                v61 := GenS7{F0: v17, F1: true, F2: 0}
                                                _ = v61
                                            case "a":
                                                for fi62 := 0; fi62 < 1; fi62++ {
                                                    var v63 int64 = 1
                                                    _ = fi62
                                                    _ = v63
                                                }
                                            case "":
                                                v21 = append(v21, (^(fi18)))
                                            default:
                                                v64 := GenS1{F0: ("hello" + ("" + ("" + "a"))), F1: 2}
                                                _ = v64
                                            }
                                        }
                                    }
                                }
                            default:
                                var v65 uint16 = (v4 / ((921) | 1))
                                _ = v65
                            }
                            fmt.Println(v3, v17, v17)
                            fmt.Println(v2, v1)
                        }
                    case 127:
                        if v3 {
                            v17 = v17
                            fmt.Println(fi18)
                        }
                    default:
                        for fi66 := 0; fi66 < 7; fi66++ {
                            switch v1 {
                            case 0:
                                switch "hello" {
                                case "foo bar":
                                    switch "a" {
                                    case "hello":
                                        switch v17 {
                                        case 489:
                                            var v67 int32 = 1
                                            _ = v67
                                        default:
                                            fmt.Println(v1, v16)
                                        }
                                    case "":
                                        v17 = 800
                                    }
                                case "\n\t":
                                    fmt.Println(v17, fi66, fi66)
                                case "a":
                                    v68 := GenS8{F0: 42}
                                    _ = v68
                                default:
                                    var v69 int32 = 2147483647
                                    fmt.Println(v1)
                                    v3 = (("\n\t" + ("a" + "foo bar")) != "hello")
                                    _ = v69
                                }
                            case 238:
                                fmt.Println(v2, v16, v2)
                                v2 = 1e10
                                var v70 int32 = 435
                                v4 += 32767
                                v71 := GenS4{F0: ""}
                                var v72 int16 = 1
                                var v73 string = ("foo bar" + "\n\t")
                                _ = v70
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                            fmt.Println(fi66, fi18, v4)
                            v74 := GenS9{F0: ((v16 * 0.0) == v16), F1: v4}
                            _ = fi66
                            _ = v74
                        }
                    }
                default:
                    v4 = v4
                }
                _ = v22
            }
            for fi75 := 0; fi75 < 4; fi75++ {
                v2 = v2
                _ = fi75
            }
        } else {
            v76 := GenS5{F0: 1000, F1: (v4 | v4)}
            if v3 {
                _ = len(v21)
            } else {
                for fi77 := 0; fi77 < 18; fi77++ {
                    v16 -= ((0.1 / -2.25) * 0.0)
                    v17 += v1
                    _ = fi77
                }
            }
            _ = v76
        }
        _ = fi18
        _ = v21
    }
    for fi78 := 0; fi78 < 15; fi78++ {
        fmt.Println(v17, fi78)
        if (v4 < (^(v4))) {
            v79 := GenS6{F0: (v4 + 0), F1: 2, F2: 873}
            _ = v79
        }
        v80 := GenS10{F0: "foo bar", F1: 160}
        _ = fi78
        _ = v80
    }
    var v81 float64 = v16
    for fi82 := 0; fi82 < 5; fi82++ {
        if false {
            if v3 {
                var v83 int64 = -1
                _ = v83
            }
            var v84 uint16 = 0
            var v85 int8 = -22
            _ = v84
            _ = v85
        }
        if false {
            for fi86 := 0; fi86 < 10; fi86++ {
                switch 128 {
                case -2147483648:
                    var v87 bool = (2 == uint32(v4))
                    _ = v87
                case 32768:
                    var v88 float64 = ((v81 + v16) + v16)
                    _ = v88
                }
                _ = fi86
            }
            if (!(!(v3 || false))) {
                var v89 float32 = 1.5
                v90 := GenS8{F0: 366}
                _ = v89
                _ = v90
            } else {
                v2 -= (v2 - v2)
                v91 := []int{}
                var v92 int64 = 9223372036854775807
                _ = v91
                _ = v92
            }
            if (2 >= 42) {
                var v93 int = (v17 / (((fi82 % (((fi82 ^ 2147483647)) | 1))) | 1))
                var v94 string = ("a" + "foo bar")
                _ = v93
                _ = v94
            } else {
                var v95 string = "foo bar"
                if (!(-1 > 3)) {
                    v96 := []int{32767, (fi82 / ((v17) | 1)), int(v4)}
                    v81 += ((3.14159 / v81) + (v81 + 1.5))
                    _ = v96
                } else {
                    switch (v1 + int(v4)) {
                    case -2147483648:
                        v97 := []int{(-(v17)), 32768, v1, v1}
                        _ = v97
                    case 4294967295:
                        switch (v95 + ((("hello" + "a") + ("a" + (v95 + ""))) + v95)) {
                        case "hello":
                            v98 := GenS11{F0: 97, F1: (1 >= 0)}
                            _ = v98
                        case "\n\t":
                            var v99 string = "hello"
                            _ = v99
                        }
                    case 9223372036854775807:
                        var v100 uint8 = 0
                        _ = v100
                    }
                }
                _ = v95
            }
        }
        _ = fi82
    }
    for fi101 := 0; fi101 < 4; fi101++ {
        v1 += (^(fi101))
        if ((!v3) || v3) {
            for fi102 := 0; fi102 < 12; fi102++ {
                var v103 int8 = int8(fi102)
                if (v103 < (v103 & v103)) {
                    var v104 uint8 = 160
                    _ = v104
                }
                fmt.Println(v4, fi102)
                var v105 int16 = 42
                _ = fi102
                _ = v103
                _ = v105
            }
        }
        var v106 int8 = 127
        _ = fi101
        _ = v106
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v81
}
