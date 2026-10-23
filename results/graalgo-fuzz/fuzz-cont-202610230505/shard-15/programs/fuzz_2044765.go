package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int }
type GenS2 struct { F0 int64; F1 int8; F2 int16 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 int8; F2 uint64 }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 uint16; F1 int8; F2 uint }
type GenS9 struct { F0 uint }
func genFn0(p0 int16) int32 {
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(fi6, fi6)
        for fi7 := 0; fi7 < 4; fi7++ {
            p0 += (p0 >> ((0) & 63))
            _ = fi7
        }
        v8 := GenS1{F0: 42, F1: 0, F2: fi6}
        _ = fi6
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        switch 42 {
        case 9223372036854775807:
            fmt.Println(p0)
        case -345:
            p0 = 1000
        case 850:
            v10 := GenS1{F0: 0, F1: 166, F2: (fi9 >> ((2) & 63))}
            _ = v10
        }
        p0 = (p0 % ((p0) | 1))
        _ = fi9
    }
    p0 -= (p0 << ((18446744073709551615) & 63))
    _ = p0
    return -927
}

func main() {
    var v1 uint64 = 385
    var v2 int = 466
    var v3 int16 = -554
    var v4 int16 = (v3 & (-(v3)))
    var v5 uint = 42
    if (2147483647 > 998) {
        for fi11 := 0; fi11 < 5; fi11++ {
            v1 += v1
            var v12 float32 = float32(1e10)
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: -160, F1: 13, F2: v2}
        _ = v13
    } else {
        switch ("\n\t" + "a") {
        case "\n\t":
            if (v3 > 42) {
                for fi14 := 0; fi14 < 14; fi14++ {
                    v4 = (v3 + (^(v4)))
                    if (("hello" > "foo bar") || false) {
                        v5 = 1
                    }
                    if true {
                        var v15 float64 = 3.14159
                        v2 += v2
                        if false {
                            v16 := GenS1{F0: 1, F1: 57, F2: fi14}
                            for fi17 := 0; fi17 < 10; fi17++ {
                                var v18 uint = (^(v5))
                                _ = fi17
                                _ = v18
                            }
                            _ = v16
                        } else {
                            v19 := []int{}
                            var v20 uint16 = 2
                            _ = v19
                            _ = v20
                        }
                        _ = v15
                    } else {
                        v21 := []int{fi14, int(v5), (fi14 + v2)}
                        switch "" {
                        case "foo bar":
                            if (1 > -1) {
                                v22 := GenS2{F0: 426, F1: -128, F2: (v3 | v3)}
                                var v23 float64 = 1.5
                                for fi24 := 0; fi24 < 17; fi24++ {
                                    var v25 int = fi24
                                    _ = fi24
                                    _ = v25
                                }
                                _ = v22
                                _ = v23
                            }
                        case "a":
                            v5 = 42
                        }
                        _ = v21
                    }
                    _ = fi14
                }
                for fi26 := 0; fi26 < 10; fi26++ {
                    for fi27 := 0; fi27 < 11; fi27++ {
                        var v28 int64 = -1
                        var v29 bool = false
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    v1 -= 1
                    for fi30 := 0; fi30 < 2; fi30++ {
                        v3 += 1000
                        if ((-2.25 - 1.5) < float64(1e10)) {
                            v31 := []int{v2, (fi26 - fi30)}
                            var v32 int8 = 0
                            _ = v31
                            _ = v32
                        } else {
                            v3 = 1
                        }
                        var v33 float32 = float32((1.5 / (0.0 - (1.5 / 3.14159))))
                        v34 := []int{(fi26 * (-(fi26)))}
                        v2 = (fi26 >> (((^(v5))) & 63))
                        _ = fi30
                        _ = v33
                        _ = v34
                    }
                    _ = fi26
                }
                for fi35 := 0; fi35 < 9; fi35++ {
                    for fi36 := 0; fi36 < 15; fi36++ {
                        var v37 int32 = 66
                        v38 := []int{(v2 | (fi36 >> ((v5) & 63))), fi36}
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                    _ = fi35
                }
            } else {
                for fi39 := 0; fi39 < 6; fi39++ {
                    v1 -= (^(v1))
                    if true {
                        v2 = (v2 ^ fi39)
                    } else {
                        v3 -= (v3 - 248)
                        v40 := []int{-2147483648, v2}
                        switch (("hello" + "foo bar") + "\n\t") {
                        case "hello":
                            fmt.Println(v3)
                        case "a":
                            for fi41 := 0; fi41 < 12; fi41++ {
                                var v42 uint = v5
                                if true {
                                    var v43 uint8 = 240
                                    switch (fi39 << ((0) & 63)) {
                                    case 127:
                                        v44 := GenS3{F0: ("foo bar" + "foo bar"), F1: (-2.25 > 1.5)}
                                        _ = v44
                                    case -1:
                                        v42 -= (v42 - (^(v5)))
                                        if len(v40) >= 2 {
                                            v40 = v40[1:len(v40)]
                                        }
                                        v45 := GenS1{F0: 1, F1: v43, F2: fi39}
                                        if len(v40) > 0 {
                                            v40[0] = -2147483648
                                        }
                                        _ = v45
                                    default:
                                        fmt.Println(fi39, v43, v43)
                                    }
                                    _ = v43
                                }
                                _ = fi41
                                _ = v42
                            }
                        default:
                            if len(v40) >= 2 {
                                v40 = v40[1:len(v40)]
                            }
                        }
                        _ = v40
                    }
                    v46 := GenS3{F0: "foo bar", F1: true}
                    _ = fi39
                    _ = v46
                }
                v1 = (v1 | (v1 << ((2) & 63)))
            }
        case "hello":
            v47 := GenS3{F0: (("a" + "") + "foo bar"), F1: (!(false && (!false)))}
            _ = v47
        case "a":
            v1 += (v1 >> (((v5 * v5)) & 63))
        default:
            if ((((1.5 + 1.5) * 0.0) / 0.1) < (1.5 - 3.14159)) {
                if (((v1 / ((4294967295) | 1)) != (v1 << ((811) & 63))) || true) {
                    var v48 float64 = 0.1
                    for fi49 := 0; fi49 < 14; fi49++ {
                        v50 := []int{4294967295, 42}
                        for fi51 := 0; fi51 < 20; fi51++ {
                            v52 := GenS3{F0: "", F1: false}
                            var v53 int32 = 230
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        switch "foo bar" {
                        case "foo bar":
                            var v54 string = ("\n\t" + "a")
                            _ = v54
                        case "hello":
                            for fi55 := 0; fi55 < 7; fi55++ {
                                switch "" {
                                case "foo bar":
                                    v48 += 1e10
                                    v50 = append(v50, fi55)
                                    v5 -= v5
                                case "a":
                                    v56 := GenS2{F0: 42, F1: 0, F2: v4}
                                    _ = v56
                                case "\n\t":
                                    if false {
                                        var v57 int = (fi55 % ((v2) | 1))
                                        _ = v57
                                    }
                                default:
                                    v58 := GenS1{F0: -2147483648, F1: 0, F2: fi55}
                                    _ = v58
                                }
                                if ((0.1 != float32(0.1)) && false) {
                                    _ = len(v50)
                                } else {
                                    _ = len(v50)
                                    if true {
                                        v50 = append(v50, fi55)
                                        var v59 int32 = 473
                                        fmt.Println(v4, fi55)
                                        _ = v59
                                    }
                                }
                                v50 = append(v50, (-(fi49)))
                                _ = fi55
                            }
                        default:
                            fmt.Println(v4, v1, v4)
                        }
                        _ = fi49
                        _ = v50
                    }
                    var v60 uint8 = 2
                    _ = v48
                    _ = v60
                } else {
                    if true {
                        v61 := []int{v2}
                        if (0 <= -741) {
                            fmt.Println(v1, v5)
                            if ((((v2 << ((v5) & 63)) > (v2 * (v2 / ((v2) | 1)))) || ((1.5 + 1e10) >= 1.5)) && false) {
                                v61 = append(v61, (v2 / ((v2) | 1)))
                                fmt.Println(v3, v5, v3)
                                if true {
                                    if (-2.25 == float32(((((float64(0.1) / 1.0) / (-2.25 + 0.1)) + ((1.5 * 3.14159) + 0.0)) - 0.1))) {
                                        switch v2 {
                                        case 45:
                                            for fi62 := 0; fi62 < 19; fi62++ {
                                                v5 -= (v5 / ((0) | 1))
                                                fmt.Println(v5)
                                                _ = fi62
                                            }
                                        }
                                    }
                                } else {
                                    fmt.Println(v2)
                                }
                            }
                        } else {
                            v63 := GenS2{F0: int64(v1), F1: 1, F2: 1000}
                            v64 := GenS4{F0: 1, F1: 2, F2: v5}
                            _ = v63
                            _ = v64
                        }
                        _ = v61
                    }
                    var v65 int8 = 127
                    _ = v65
                }
                v66 := GenS5{F0: v2}
                _ = v66
            }
        }
    }
    if true {
        if true {
            var v67 uint = v5
            if (2 != 686) {
                var v68 int64 = -1
                var v69 int16 = v4
                for fi70 := 0; fi70 < 3; fi70++ {
                    fmt.Println(fi70, v67, v5)
                    var v71 int8 = 100
                    for fi72 := 0; fi72 < 16; fi72++ {
                        v68 += (^(v68))
                        var v73 uint32 = 65535
                        v74 := GenS6{F0: "foo bar", F1: v71, F2: (v1 + (v1 % (((^(v1))) | 1)))}
                        _ = fi72
                        _ = v73
                        _ = v74
                    }
                    _ = fi70
                    _ = v71
                }
                _ = v68
                _ = v69
            } else {
                v75 := []int{}
                var v76 float64 = 3.14159
                _ = v75
                _ = v76
            }
            _ = v67
        } else {
            for fi77 := 0; fi77 < 11; fi77++ {
                v5 += 4294967295
                var v78 float64 = 1e10
                _ = fi77
                _ = v78
            }
        }
    } else {
        v79 := GenS7{F0: 0, F1: 42}
        _ = v79
    }
    var v80 int8 = 1
    v81 := GenS8{F0: 787, F1: v80, F2: (v5 * v5)}
    var v82 int16 = 42
    if (0.1 != ((0.1 + ((0.0 * (0.1 + 0.1)) / 1.5)) / 3.14159)) {
        var v83 float64 = 1.5
        var v84 string = ""
        var v85 uint8 = 0
        _ = v83
        _ = v84
        _ = v85
    }
    if true {
        if (!false) {
            switch (v2 << ((859) & 63)) {
            case -258:
                v2 += v2
            default:
                if true {
                    if (!true) {
                        v86 := GenS8{F0: 0, F1: v80, F2: 255}
                        _ = v86
                    }
                    v87 := GenS1{F0: 0, F1: 42, F2: (-(v2))}
                    _ = v87
                } else {
                    if (0 <= 127) {
                        var v88 int8 = v80
                        var v89 bool = (2 != 308)
                        var v90 int16 = (^(v3))
                        _ = v88
                        _ = v89
                        _ = v90
                    }
                    v4 = v4
                    v91 := []int{}
                    _ = v91
                }
            }
            var v92 float64 = 3.14159
            _ = v92
        }
    } else {
        var v93 uint32 = 2
        _ = v93
    }
    v94 := []int{}
    switch (("foo bar" + "hello") + (("\n\t" + "a") + "hello")) {
    case "\n\t":
        v5 += v5
    case "hello":
        var v95 uint = 218
        _ = v95
    case "":
        var v96 uint16 = 1
        _ = v96
    default:
        if (-9223372036854775808 > -733) {
            var v97 bool = ((-1 != 667) || false)
            if false {
                var v98 int = v2
                v99 := GenS6{F0: "a", F1: v80, F2: (v1 + v1)}
                _ = v98
                _ = v99
            } else {
                var v100 string = "foo bar"
                _ = v100
            }
            v5 -= 202
            _ = v97
        } else {
            fmt.Println(v80, v4, v5)
        }
    }
    var v101 uint32 = 2147483648
    _ = len(v94)
    for fi102 := 0; fi102 < 7; fi102++ {
        for fi103 := 0; fi103 < 17; fi103++ {
            fmt.Println(fi103, v3, v3)
            _ = fi103
        }
        v5 -= (v5 * (v5 & 18446744073709551615))
        var v104 float32 = 1e10
        _ = fi102
        _ = v104
    }
    switch v2 {
    case -9223372036854775808:
        var v105 uint16 = 32767
        _ = v105
    case -525:
        var v106 float32 = 1e10
        _ = v106
    case -2147483648:
        var v107 int = 128
        _ = v107
    }
    if (!((true || true) && ((v101 | (^(v101))) <= (v101 * v101)))) {
        for fi108 := 0; fi108 < 9; fi108++ {
            _ = len(v94)
            _ = fi108
        }
    } else {
        if ((v101 % ((195) | 1)) <= (v101 + uint32(v5))) {
            v101 = (v101 + 45)
            fmt.Println(v5)
        }
        switch 42 {
        case 2147483647:
            var v109 uint16 = 0
            _ = v109
        default:
            var v110 uint64 = v1
            _ = v110
        }
        var v111 uint32 = v101
        _ = v111
    }
    var v112 int64 = 252
    v113 := GenS9{F0: v5}
    if len(v94) > 0 {
        v94[0] = v2
    }
    var v114 float64 = ((float64((float32(float64(((-2.25 * 1.5) / 1.0))) + 1e10)) - (0.1 - -2.25)) + ((0.0 - float64((float32(float64((((3.14159 / 1.5) + 1.5) - float32(float64(1e10))))) / float32((3.14159 + (float64(float32(float64(((((float32(1.5) + ((0.0 + 1e10) * 0.0)) * (float32(((3.14159 / -2.25) - 0.0)) - float32(1e10))) + ((3.14159 * 0.1) * ((-2.25 * (-2.25 / -2.25)) - 1.5))) - 1e10)))) / 1.5)))))) * -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v112:", v112)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    fmt.Println("v82:", v82)
    fmt.Println("slicelen v94:", len(v94), v94)
    _ = v1
    _ = v101
    _ = v112
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v81
    _ = v82
    _ = v94
}
