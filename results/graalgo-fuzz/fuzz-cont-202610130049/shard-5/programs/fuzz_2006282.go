package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 int16 }
type GenS3 struct { F0 int32; F1 uint32; F2 string }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int8 }
func genFn0() int32 {
    var v4 uint16 = 65535
    v4 = (v4 * 65535)
    v5 := GenS1{F0: 1, F1: 493, F2: 42}
    v6 := GenS1{F0: -54, F1: 2, F2: 0}
    for fi7 := 0; fi7 < 9; fi7++ {
        if false {
            for fi8 := 0; fi8 < 12; fi8++ {
                v4 += (v4 >> ((201) & 63))
                _ = fi8
            }
        } else {
            v9 := []int{(fi7 | (fi7 | 32768)), 4294967295, -1}
            _ = v9
        }
        _ = fi7
    }
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []string{"a"}
    var v2 uint64 = 0
    v3 := make(map[int]uint32)
    fmt.Println(v2)
    if ((true && false) && (!(0 > 255))) {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: 101, F1: 471, F2: (v2 | v2)}
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := GenS1{F0: -13, F1: 0, F2: 2}
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint = 18446744073709551615
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    } else {
        fmt.Println(v2)
        var v16 uint = 1
        v16 += v16
        _ = v16
    }
    if false {
        var v17 int32 = 199
        if true {
            if true {
                _ = len(v1)
                var v18 int32 = (v17 - (v17 / ((genFn0()) | 1)))
                fmt.Println(v18)
                _ = v18
            } else {
                v19 := GenS1{F0: -64, F1: 1, F2: v2}
                _ = v19
            }
            var v20 bool = (956 == 0)
            switch "a" {
            case "a":
                v17 += v17
            case "":
                v21 := GenS1{F0: 0, F1: (^(v2)), F2: 1}
                _ = v21
            default:
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            }
            _ = v20
        } else {
            var v22 int32 = -724
            _ = v22
        }
        _ = v17
    } else {
        v23 := GenS2{F0: 0, F1: 255, F2: 42}
        _ = v23
    }
    var v24 uint64 = 18446744073709551615
    switch 896 {
    case -2147483648:
        _ = len(v1)
    default:
        for fi25 := 0; fi25 < 4; fi25++ {
            if true {
                for fi26 := 0; fi26 < 17; fi26++ {
                    v27 := GenS2{F0: 598, F1: 1, F2: 32767}
                    v28 := GenS1{F0: -7, F1: (v2 | v24), F2: v2}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            } else {
                for fi29 := 0; fi29 < 18; fi29++ {
                    if (236 == 0) {
                        var v30 float64 = -2.25
                        fmt.Println(v30)
                        v31 := GenS2{F0: genFn0(), F1: 32767, F2: -760}
                        _ = v30
                        _ = v31
                    }
                    _ = fi29
                }
                if false {
                    if (!true) {
                        switch (fi25 ^ (fi25 ^ fi25)) {
                        case 127:
                            fmt.Println(v24, v2, fi25)
                        case 545:
                            var v32 int = (^(fi25))
                            _ = v32
                        case -9223372036854775808:
                            _ = len(v3)
                        default:
                            if (!false) {
                                var v33 uint8 = 68
                                _ = v33
                            }
                        }
                    } else {
                        if true {
                            var v34 uint8 = 1
                            _ = v34
                        }
                        switch "hello" {
                        case "\n\t":
                            if ((true || (!(1 < 42))) && false) {
                                v24 = (v24 % ((42) | 1))
                                fmt.Println(fi25)
                            }
                        }
                    }
                    v2 = 18446744073709551615
                } else {
                    var v35 int32 = 1
                    if (!(57 != 1)) {
                        var v36 int = -2147483648
                        for fi37 := 0; fi37 < 11; fi37++ {
                            var v38 int32 = (v35 ^ (^(v35)))
                            v39 := GenS2{F0: 42, F1: 32767, F2: -32768}
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        _ = v36
                    } else {
                        for fi40 := 0; fi40 < 8; fi40++ {
                            v2 = (v24 + (^(v2)))
                            v41 := GenS3{F0: v35, F1: 2147483647, F2: (("a" + "a") + "a")}
                            fmt.Println(v35, v2)
                            _ = fi40
                            _ = v41
                        }
                        for fi42 := 0; fi42 < 18; fi42++ {
                            var v43 int16 = -809
                            _ = fi42
                            _ = v43
                        }
                    }
                    var v44 uint = 1
                    _ = v35
                    _ = v44
                }
                if len(v1) > 0 {
                    v1[0] = ("hello" + ("hello" + "\n\t"))
                }
            }
            var v45 int64 = 0
            v3[(fi25 % ((fi25) | 1))] = 590
            _ = fi25
            _ = v45
        }
    }
    var v46 uint = 966
    _ = len(v3)
    for fi47 := 0; fi47 < 13; fi47++ {
        var v48 float32 = ((float32(1e10) / 1.0) * 0.1)
        _ = fi47
        _ = v48
    }
    if (v46 > v46) {
        for fi49 := 0; fi49 < 2; fi49++ {
            v50 := GenS2{F0: 264, F1: 774, F2: 314}
            v24 += (v2 << (((v46 % ((uint(v2)) | 1))) & 63))
            _ = fi49
            _ = v50
        }
        switch (("" + "foo bar") + "a") {
        case "\n\t":
            if false {
                if true {
                    v3[127] = 2147483647
                    var v51 uint8 = 128
                    v46 -= 1
                    _ = v51
                } else {
                    if (0.0 < (-2.25 - -2.25)) {
                        v52 := GenS2{F0: 0, F1: 459, F2: 42}
                        if (2147483648 > 627) {
                            v2 = 0
                        } else {
                            for fi53 := 0; fi53 < 13; fi53++ {
                                _ = len(v1)
                                var v54 float32 = float32(((((-2.25 - (3.14159 / (0.0 - ((1.5 - 1e10) + 3.14159)))) - float64(3.14159)) - (1e10 + 1.5)) - (0.1 * 0.1)))
                                _ = fi53
                                _ = v54
                            }
                            switch 2147483647 {
                            case 127:
                                if false {
                                    fmt.Println(v46)
                                }
                            case 0:
                                var v55 int64 = 41
                                _ = v55
                            }
                            var v56 uint = (^(v46))
                            _ = v56
                        }
                        _ = v52
                    } else {
                        var v57 uint16 = 65535
                        v46 -= v46
                        _ = v57
                    }
                    fmt.Println(v2, v46)
                    v58 := GenS2{F0: 0, F1: 255, F2: 980}
                    _ = v58
                }
                v24 -= (v24 | v24)
            }
        case "hello":
            v59 := GenS4{F0: 255, F1: 42}
            _ = v59
        case "foo bar":
            if false {
                v2 -= (v24 << ((952) & 63))
            }
        default:
            for fi60 := 0; fi60 < 8; fi60++ {
                switch "" {
                case "\n\t":
                    v2 = (v24 >> ((65535) & 63))
                case "hello":
                    v61 := GenS5{F0: int64(v46)}
                    _ = v61
                case "a":
                    if true {
                        _ = len(v3)
                        for fi62 := 0; fi62 < 10; fi62++ {
                            fmt.Println(fi62, v24)
                            v46 += (v46 + 956)
                            v63 := GenS1{F0: 0, F1: (v24 * 2), F2: (^(v24))}
                            _ = fi62
                            _ = v63
                        }
                    } else {
                        v2 = 1
                        if true {
                            v24 += (v2 * v24)
                        } else {
                            v2 += (v2 | v2)
                            for fi64 := 0; fi64 < 10; fi64++ {
                                fmt.Println(fi60, fi64)
                                v46 = (v46 + v46)
                                var v65 uint32 = 4294967295
                                _ = fi64
                                _ = v65
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                default:
                    if (v2 <= 0) {
                        v2 = v24
                    }
                }
                if (42 <= 1) {
                    if true {
                        v24 -= v2
                    } else {
                        v24 = (v24 | (v24 % (((v24 >> (((v46 >> ((v46) & 63))) & 63))) | 1)))
                        var v66 uint64 = (v2 & (^(v24)))
                        _ = v66
                    }
                    switch -821 {
                    case 9223372036854775807:
                        if (1 > 127) {
                            var v67 bool = (1 <= genFn0())
                            if true {
                                fmt.Println(v46)
                                for fi68 := 0; fi68 < 17; fi68++ {
                                    switch (fi60 ^ (fi68 + fi60)) {
                                    case 9223372036854775807:
                                        _ = len(v1)
                                    case 32768:
                                        var v69 int32 = -1
                                        _ = v69
                                    }
                                    if (255 == (^(v46))) {
                                        var v70 int64 = 42
                                        var v71 int8 = 42
                                        _ = v70
                                        _ = v71
                                    }
                                    v2 += (^(v2))
                                    _ = fi68
                                }
                                fmt.Println(v46, v24)
                            }
                            fmt.Println(v67)
                            _ = v67
                        }
                    case 42:
                        var v72 uint64 = (v24 ^ (v24 * (v2 % (((v24 | (v24 * (v2 % ((v2) | 1))))) | 1))))
                        _ = v72
                    default:
                        if (!(0 > 0)) {
                            v46 = (v46 >> ((18446744073709551615) & 63))
                            var v73 int32 = genFn0()
                            _ = v73
                        }
                    }
                    for fi74 := 0; fi74 < 3; fi74++ {
                        var v75 bool = true
                        v2 = (v2 & (v2 * v24))
                        _ = fi74
                        _ = v75
                    }
                } else {
                    v24 -= (v2 | (v24 / ((786) | 1)))
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi60
            }
        }
        v76 := GenS6{F0: 127, F1: (v24 - 4294967295)}
        _ = v76
    }
    var v77 float64 = float64(0.1)
    var v78 int = 0
    for fi79 := 0; fi79 < 4; fi79++ {
        switch fi79 {
        case 4294967295:
            v80 := GenS5{F0: -1}
            _ = v80
        case 56:
            switch 0 {
            case 0:
                var v81 int16 = 32767
                _ = v81
            case 4294967295:
                if ((^(v46)) <= (^(v46))) {
                    var v82 uint16 = 716
                    switch v78 {
                    case 994:
                        var v83 bool = true
                        _ = v83
                    }
                    switch (v78 / ((4294967295) | 1)) {
                    case 159:
                        if true {
                            v24 += (v2 | v24)
                            v84 := GenS7{F0: 100}
                            _ = v84
                        }
                    case 42:
                        fmt.Println(v78, v78)
                    }
                    _ = v82
                } else {
                    switch (v78 << ((v46) & 63)) {
                    case 4294967295:
                        for fi85 := 0; fi85 < 3; fi85++ {
                            var v86 uint = (v46 / ((uint(v24)) | 1))
                            _ = fi85
                            _ = v86
                        }
                    }
                    if true {
                        for fi87 := 0; fi87 < 11; fi87++ {
                            var v88 int8 = 46
                            switch 306 {
                            case 2147483647:
                                var v89 int16 = 42
                                _ = v89
                            case -378:
                                var v90 uint = (v46 + v46)
                                _ = v90
                            case 0:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                            }
                            _ = fi87
                            _ = v88
                        }
                    }
                    for fi91 := 0; fi91 < 11; fi91++ {
                        fmt.Println(fi79)
                        _ = fi91
                    }
                }
            case -537:
                if false {
                    v78 -= (fi79 / ((9223372036854775807) | 1))
                    switch "" {
                    case "hello":
                        if (189 < uint8(v46)) {
                            if (65535 >= 2) {
                                if (1e10 >= ((((1.5 + 3.14159) - (1.5 / 0.1)) + ((0.0 - (1e10 / 1e10)) - 0.0)) - 3.14159)) {
                                    var v92 int16 = -648
                                    v93 := GenS7{F0: 42}
                                    _ = v92
                                    _ = v93
                                }
                            } else {
                                for fi94 := 0; fi94 < 15; fi94++ {
                                    v78 = fi79
                                    v46 = (^(v46))
                                    var v95 int = fi79
                                    v24 += (v2 - 4294967295)
                                    switch ("foo bar" + (("\n\t" + "hello") + "")) {
                                    case "a":
                                        var v96 float32 = 0.0
                                        _ = v96
                                    case "\n\t":
                                        var v97 int = fi79
                                        var v98 uint64 = v2
                                        var v99 int16 = 47
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    case "foo bar":
                                        switch 0 {
                                        case 972:
                                            v95 = fi79
                                            if (true || true) {
                                                if false {
                                                    if true {
                                                        v100 := GenS2{F0: 2147483647, F1: 241, F2: 32767}
                                                        var v101 uint8 = 255
                                                        _ = v100
                                                        _ = v101
                                                    }
                                                } else {
                                                    v24 = (^(v24))
                                                    v102 := GenS1{F0: 100, F1: (v2 * 1), F2: v2}
                                                    _ = v102
                                                }
                                            } else {
                                                var v103 string = ("hello" + "foo bar")
                                                _ = v103
                                            }
                                        }
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    _ = fi94
                                    _ = v95
                                }
                                fmt.Println(v77, v77, v78)
                            }
                        } else {
                            switch ("a" + "foo bar") {
                            case "\n\t":
                                fmt.Println(v78, v46)
                            case "foo bar":
                                var v104 int64 = 1
                                _ = v104
                            }
                            if false {
                                var v105 uint32 = 2147483647
                                var v106 int32 = -1
                                _ = v105
                                _ = v106
                            }
                            v107 := v3[(v78 % ((fi79) | 1))]
                            _ = v107
                        }
                    case "foo bar":
                        for fi108 := 0; fi108 < 13; fi108++ {
                            var v109 uint64 = v24
                            _ = fi108
                            _ = v109
                        }
                    case "a":
                        fmt.Println(v78, v24, v24)
                    }
                    for fi110 := 0; fi110 < 11; fi110++ {
                        var v111 uint8 = 128
                        v112 := v3[0]
                        v3[fi110] = 452
                        _ = fi110
                        _ = v111
                        _ = v112
                    }
                } else {
                    if false {
                        var v113 uint64 = (v24 * (v24 / (((v2 * v2)) | 1)))
                        v77 += -2.25
                        var v114 uint16 = uint16(v24)
                        _ = v113
                        _ = v114
                    }
                    if true {
                        v24 = v2
                    } else {
                        var v115 bool = ((((v77 / ((v77 * v77) / float64(1.5))) - v77) * (v77 / v77)) >= v77)
                        var v116 int = fi79
                        _ = v115
                        _ = v116
                    }
                }
            }
        case 2147483647:
            if (0 != uint64(v46)) {
                v3[fi79] = 362
                var v117 uint64 = (v24 + 1)
                _ = len(v1)
                _ = v117
            } else {
                var v118 uint32 = 4294967295
                if (3.14159 == 3.14159) {
                    var v119 uint16 = 426
                    v119 -= v119
                    _ = v119
                }
                var v120 uint = v46
                _ = v118
                _ = v120
            }
        }
        var v121 int8 = 0
        switch (v78 / (((v78 ^ (-(fi79)))) | 1)) {
        case -992:
            v1 = append(v1, ("a" + "hello"))
        default:
            _ = len(v3)
        }
        _ = fi79
        _ = v121
    }
    v77 = 0.1
    var v122 int = (v78 + (^(v78)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v122
    _ = v2
    _ = v24
    _ = v3
    _ = v46
    _ = v77
    _ = v78
}
