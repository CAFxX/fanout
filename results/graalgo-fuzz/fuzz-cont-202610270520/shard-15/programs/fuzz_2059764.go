package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint8; F2 int64 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 string }
func genFn0() uint16 {
    switch "" {
    case "foo bar":
        v6 := []int{54, 4294967295}
        _ = v6
    case "":
        var v7 int16 = 32767
        _ = v7
    default:
        if true {
            var v8 float64 = 0.1
            _ = v8
        } else {
            var v9 uint8 = 42
            if (((1e10 + 1e10) + 0.0) > ((((-2.25 * (0.1 / ((3.14159 * 3.14159) * ((0.0 + (1.5 - 3.14159)) + 0.0)))) - 0.0) / -2.25) - (3.14159 - (((-2.25 * (1e10 / 3.14159)) / float32((0.1 * 0.1))) - (1.5 - 1.5))))) {
                v9 = 254
                for fi10 := 0; fi10 < 1; fi10++ {
                    v9 = v9
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v12 := GenS1{F0: 96}
                        if (int8(fi11) != 39) {
                            fmt.Println(fi10, v9, fi11)
                        }
                        v13 := GenS1{F0: 42}
                        if (!false) {
                            v9 = (^(v9))
                            v14 := GenS1{F0: 79}
                            if (!false) {
                                v9 = v9
                                v15 := GenS1{F0: 2}
                                _ = v15
                            }
                            _ = v14
                        } else {
                            v9 -= 12
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    var v16 string = ""
                    _ = fi10
                    _ = v16
                }
                v17 := GenS1{F0: 32767}
                _ = v17
            }
            v18 := GenS2{F0: (v9 ^ (v9 & (^(v9))))}
            _ = v18
            _ = v9
        }
    }
    var v19 int8 = 100
    v20 := GenS2{F0: 255}
    _ = v19
    _ = v20
    return 119
}
func genFn1(p0 int, p1 int64, p2 int32) int32 {
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := []int{}
        v22 = append(v22, (p0 % ((4294967295) | 1)))
        _ = fi21
        _ = v22
    }
    var v23 uint16 = 42
    v24 := GenS1{F0: v23}
    v23 = (v23 >> ((141) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    return p2
}

func main() {
    var v1 float64 = 0.1
    var v2 int64 = 647
    v3 := []int{}
    var v4 string = ""
    var v5 int32 = 1
    v1 = (-2.25 * 3.14159)
    v25 := GenS1{F0: 1}
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := GenS1{F0: 932}
        for fi28 := 0; fi28 < 1; fi28++ {
            if len(v3) > 0 {
                v3[0] = fi26
            }
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v29 int64 = v2
                _ = len(v3)
                fmt.Println(fi28, v5)
                _ = v29
            }
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        v4 = ("foo bar" + "hello")
        v31 := GenS2{F0: 55}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        v1 = (0.1 / 3.14159)
        v33 := GenS1{F0: 0}
        _ = fi32
        _ = v33
    }
    if (912 != 255) {
        for fi34 := 0; fi34 < 11; fi34++ {
            v2 -= v2
            switch (fi34 << ((42) & 63)) {
            case 4294967295:
                var v35 uint64 = 2
                _ = v35
            case 177:
                if false {
                    fmt.Println(v4, v2)
                    for fi36 := 0; fi36 < 12; fi36++ {
                        v5 = (v5 / (((v5 >> ((0) & 63))) | 1))
                        v1 = float64(0.0)
                        var v37 int64 = v2
                        fmt.Println(v37, v37, v1)
                        switch fi34 {
                        case 4294967295:
                            if true {
                                if (!false) {
                                    v4 = ""
                                    var v38 int = -1
                                    if (0 != 18446744073709551615) {
                                        v37 += (-(v2))
                                        v39 := GenS3{F0: 0, F1: 42, F2: v2}
                                        _ = v39
                                    } else {
                                        if false {
                                            fmt.Println(v4, v38)
                                            v3 = append(v3, v38)
                                            switch (v38 / ((-2147483648) | 1)) {
                                            case 42:
                                                _ = len(v3)
                                            case 32768:
                                                var v40 uint16 = 0
                                                _ = v40
                                            case -9223372036854775808:
                                                switch -2147483648 {
                                                case -201:
                                                    var v41 float64 = v1
                                                    _ = v41
                                                default:
                                                    var v42 uint64 = 4294967295
                                                    _ = v42
                                                }
                                            default:
                                                var v43 uint = 38
                                                _ = v43
                                            }
                                        } else {
                                            v37 = -1
                                            if len(v3) > 0 {
                                                v3[0] = fi34
                                            }
                                            _ = len(v3)
                                        }
                                        v37 -= 0
                                    }
                                    _ = v38
                                } else {
                                    _ = len(v3)
                                    v5 -= 42
                                    if false {
                                        v37 += v2
                                        fmt.Println(fi36)
                                    }
                                }
                                v5 -= v5
                            } else {
                                v44 := GenS4{F0: "hello", F1: 255}
                                if true {
                                    v2 -= -1
                                    v2 = v37
                                } else {
                                    switch (fi34 | fi34) {
                                    case 1:
                                        var v45 int32 = (-(v5))
                                        _ = v45
                                    default:
                                        fmt.Println(v37, v4)
                                    }
                                    switch (fi34 ^ (^(fi34))) {
                                    case 128:
                                        switch fi34 {
                                        case -1:
                                            fmt.Println(v37, v1)
                                        case -9223372036854775808:
                                            var v46 bool = false
                                            _ = v46
                                        default:
                                            v5 += (v5 >> ((443) & 63))
                                        }
                                    }
                                }
                                var v47 uint64 = 1
                                _ = v44
                                _ = v47
                            }
                        case 127:
                            if (false && (0 >= 127)) {
                                switch "a" {
                                case "\n\t":
                                    v4 = v4
                                case "hello":
                                    if true {
                                        if false {
                                            _ = len(v3)
                                        } else {
                                            if (genFn0() > genFn0()) {
                                                v5 -= v5
                                                var v48 int64 = v2
                                                var v49 int16 = 32767
                                                v5 += genFn1((-(fi36)), 42, (v5 ^ genFn1((fi36 * (fi34 & (fi34 & (fi34 ^ (fi34 % (((fi36 << ((42) & 63))) | 1)))))), v37, 2147483647)))
                                                _ = v48
                                                _ = v49
                                            } else {
                                                switch v4 {
                                                case "hello":
                                                    var v50 int64 = v37
                                                    _ = v50
                                                case "":
                                                    fmt.Println(fi34)
                                                case "foo bar":
                                                    v1 -= ((((1e10 - v1) / v1) - 0.0) / (v1 / v1))
                                                default:
                                                    var v51 int16 = 1
                                                    _ = v51
                                                }
                                                fmt.Println(v5, v37, v5)
                                                v1 += 1.5
                                            }
                                            v2 += -1
                                        }
                                        if true {
                                            var v52 int64 = 0
                                            v3 = append(v3, (fi36 >> ((500) & 63)))
                                            _ = v52
                                        } else {
                                            v5 += v5
                                        }
                                        fmt.Println(v2, v2, fi34)
                                    } else {
                                        var v53 float32 = (0.0 / 3.14159)
                                        var v54 int16 = -798
                                        var v55 int32 = (^(v5))
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                }
                                fmt.Println(v5, fi36)
                                v5 -= v5
                                if (1 == 127) {
                                    if false {
                                        _ = len(v3)
                                    }
                                }
                            } else {
                                v56 := GenS4{F0: ((v4 + "\n\t") + v4), F1: 255}
                                var v57 int = fi34
                                _ = v56
                                _ = v57
                            }
                        default:
                            if (!true) {
                                var v58 float32 = 1.5
                                v59 := GenS4{F0: ((v4 + "\n\t") + v4), F1: 65535}
                                _ = v58
                                _ = v59
                            }
                        }
                        _ = fi36
                        _ = v37
                    }
                    var v60 int16 = 42
                    _ = v60
                } else {
                    for fi61 := 0; fi61 < 9; fi61++ {
                        v4 = "foo bar"
                        if false {
                            v1 += v1
                            if (481 == 4294967295) {
                                var v62 uint = 94
                                v2 -= (v2 % ((v2) | 1))
                                _ = v62
                            } else {
                                v2 = 0
                                _ = len(v3)
                                if (v1 <= v1) {
                                    v63 := GenS1{F0: 480}
                                    fmt.Println(v1, v2, v2)
                                    if len(v3) > 0 {
                                        v3[0] = fi61
                                    }
                                    v64 := GenS1{F0: 1}
                                    _ = v63
                                    _ = v64
                                }
                            }
                        }
                        fmt.Println(v5)
                        var v65 bool = (!false)
                        _ = fi61
                        _ = v65
                    }
                    switch (v4 + (("hello" + v4) + v4)) {
                    case "foo bar":
                        if (!true) {
                            var v66 int32 = 339
                            _ = v66
                        } else {
                            v3 = append(v3, (fi34 / ((127) | 1)))
                            var v67 string = ""
                            _ = v67
                        }
                    case "\n\t":
                        var v68 int8 = 1
                        _ = v68
                    case "a":
                        if len(v3) > 0 {
                            v3[0] = (fi34 << ((4294967295) & 63))
                        }
                    default:
                        if (!(true && true)) {
                            var v69 string = "hello"
                            v70 := GenS1{F0: 2}
                            _ = v69
                            _ = v70
                        } else {
                            for fi71 := 0; fi71 < 6; fi71++ {
                                var v72 int16 = 965
                                _ = fi71
                                _ = v72
                            }
                        }
                    }
                }
            case 32767:
                var v73 uint32 = 217
                _ = v73
            }
            _ = fi34
        }
        if false {
            if true {
                if len(v3) > 0 {
                    v3[0] = 61
                }
            } else {
                switch 181 {
                case 128:
                    v74 := GenS4{F0: v4, F1: 203}
                    _ = v74
                default:
                    v75 := GenS1{F0: 255}
                    _ = v75
                }
                var v76 bool = (2 >= 4294967295)
                v77 := GenS3{F0: (v2 % ((v2) | 1)), F1: 127, F2: 0}
                _ = v76
                _ = v77
            }
            fmt.Println(v5, v4, v2)
        } else {
            var v78 float32 = float32(0.0)
            v2 += 9223372036854775807
            _ = v78
        }
        var v79 uint64 = 2
        _ = v79
    }
    for fi80 := 0; fi80 < 12; fi80++ {
        if (255 >= 2) {
            if len(v3) > 0 {
                v3[0] = fi80
            }
        }
        if (!(true && true)) {
            switch v4 {
            case "a":
                var v81 int32 = (v5 + v5)
                _ = v81
            default:
                var v82 uint32 = 4294967295
                _ = v82
            }
            for fi83 := 0; fi83 < 11; fi83++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi83
            }
            for fi84 := 0; fi84 < 16; fi84++ {
                if true {
                    switch (("" + (v4 + v4)) + v4) {
                    case "\n\t":
                        v4 = (((v4 + (((v4 + "foo bar") + "hello") + "a")) + ("hello" + v4)) + v4)
                    case "hello":
                        fmt.Println(fi84, v1, v2)
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = (^(fi80))
                        }
                    }
                    var v85 float64 = ((((v1 - v1) * v1) / v1) / (1.5 * 1.5))
                    _ = v85
                }
                var v86 int = fi80
                var v87 uint64 = 4294967295
                _ = fi84
                _ = v86
                _ = v87
            }
        } else {
            switch fi80 {
            case -9223372036854775808:
                v2 = v2
            case -950:
                if (!false) {
                    for fi88 := 0; fi88 < 15; fi88++ {
                        v4 = (((("" + "a") + (((v4 + "\n\t") + ("foo bar" + "\n\t")) + ((((v4 + (((("hello" + v4) + "a") + v4) + (v4 + v4))) + v4) + v4) + ""))) + ("foo bar" + v4)) + ("\n\t" + (v4 + v4)))
                        if (true && false) {
                            switch v4 {
                            case "":
                                switch (v4 + ("a" + v4)) {
                                case "\n\t":
                                    var v89 uint = 0
                                    _ = v89
                                case "foo bar":
                                    v90 := GenS5{F0: "a"}
                                    _ = v90
                                default:
                                    var v91 float32 = 0.0
                                    _ = v91
                                }
                            }
                        } else {
                            var v92 uint64 = 42
                            _ = v92
                        }
                        if false {
                            if false {
                                _ = len(v3)
                                _ = len(v3)
                                var v93 uint = 4294967295
                                _ = v93
                            }
                        } else {
                            v4 = ((v4 + ((v4 + "hello") + v4)) + v4)
                            v2 += (v2 ^ v2)
                        }
                        v1 += float64(1.5)
                        _ = fi88
                    }
                    if len(v3) > 0 {
                        v3[0] = fi80
                    }
                }
            case 32767:
                if len(v3) > 0 {
                    v3[0] = fi80
                }
            }
            var v94 float32 = 0.0
            _ = v94
        }
        _ = fi80
    }
    for fi95 := 0; fi95 < 4; fi95++ {
        fmt.Println(v5)
        _ = fi95
    }
    if true {
        if true {
            v3 = append(v3, 98)
            v2 = 1
            var v96 int64 = v2
            _ = v96
        }
        v1 = 1.5
    } else {
        if len(v3) > 0 {
            v3[0] = 511
        }
        v97 := GenS3{F0: 923, F1: 0, F2: 1}
        fmt.Println(v4, v5)
        _ = v97
    }
    switch 8 {
    case 0:
        var v98 uint8 = 42
        _ = v98
    case 245:
        v99 := GenS4{F0: ("foo bar" + "\n\t"), F1: 255}
        _ = v99
    default:
        v4 = v4
    }
    v3 = append(v3, -613)
    var v100 bool = (v2 <= (v2 + v2))
    switch 32768 {
    case 42:
        var v101 int64 = (v2 >> ((4294967295) & 63))
        _ = v101
    case 32768:
        v5 -= (v5 + (v5 ^ genFn1(1, (v2 / ((-1) | 1)), (v5 * 2147483647))))
    case 1:
        if (!(((1e10 + (0.1 / 1e10)) + (1e10 / 1.0)) >= (3.14159 * (((3.14159 + -2.25) / (float32(1e10) / 1e10)) * 1.5)))) {
            if true {
                if v100 {
                    v100 = v100
                    if (float64(-2.25) < v1) {
                        var v102 uint = 0
                        fmt.Println(v5, v102, v102)
                        fmt.Println(v4, v102, v100)
                        _ = v102
                    }
                } else {
                    var v103 int32 = (v5 - (v5 * v5))
                    for fi104 := 0; fi104 < 4; fi104++ {
                        switch v4 {
                        case "":
                            v2 = 1
                        case "hello":
                            var v105 int16 = -200
                            _ = v105
                        }
                        v106 := GenS5{F0: v4}
                        var v107 int64 = v2
                        _ = fi104
                        _ = v106
                        _ = v107
                    }
                    _ = v103
                }
                if (127 > 8) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v4 = ("" + v4)
                } else {
                    var v108 uint64 = 42
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    v100 = (v100 && v100)
                    _ = v108
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v109 uint8 = 128
            _ = v109
        }
    default:
        for fi110 := 0; fi110 < 1; fi110++ {
            v5 += (v5 | (^(v5)))
            _ = fi110
        }
    }
    var v111 float64 = ((v1 / (v1 + 3.14159)) - 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v100
    _ = v111
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
