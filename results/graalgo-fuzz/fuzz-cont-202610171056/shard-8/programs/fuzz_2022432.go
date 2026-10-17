package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int8; F1 int64; F2 uint64 }
type GenS7 struct { F0 string; F1 int; F2 uint8 }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 int16; F1 string }
type GenS11 struct { F0 string; F1 int8 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int16; F1 uint32 }
type GenS15 struct { F0 uint; F1 string; F2 int64 }
type GenS16 struct { F0 int32; F1 bool; F2 uint64 }
type GenS17 struct { F0 uint8; F1 int }
func genFn0(p0 uint16) int32 {
    v6 := GenS1{F0: 42}
    var v7 uint8 = 42
    var v8 bool = ("\n\t" >= "")
    var v9 int16 = 42
    switch 42 {
    case 9223372036854775807:
        v9 += 0
    case 0:
        if ((false && false) && (959 == uint(p0))) {
            p0 += p0
        }
    case -9223372036854775808:
        var v10 float64 = 1e10
        _ = v10
    default:
        var v11 float64 = 0.1
        _ = v11
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 int32, p1 uint, p2 string) uint8 {
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 uint8 = 127
        fmt.Println(p1, p1, fi12)
        _ = fi12
        _ = v13
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 38
}

func main() {
    var v1 float64 = 3.14159
    var v2 int64 = -1
    v3 := []int{1}
    v4 := make(map[uint32]string)
    var v5 uint8 = 42
    if (!((0.1 / ((-2.25 / float32(3.14159)) + ((((1e10 - 1.5) - 0.0) * -2.25) - (-2.25 / 1.0)))) >= float32(v1))) {
        v2 -= (v2 - -9223372036854775808)
        var v14 int = 32767
        v1 = v1
        _ = v14
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 int64 = v2
            v17 := GenS1{F0: 1}
            _ = fi15
            _ = v16
            _ = v17
        }
        if (!true) {
            if (v2 <= (v2 - v2)) {
                for fi18 := 0; fi18 < 18; fi18++ {
                    v19 := GenS2{F0: v2}
                    if (-80 == 2147483647) {
                        var v20 int64 = (v2 / ((-13) | 1))
                        _ = len(v3)
                        _ = v20
                    }
                    _ = fi18
                    _ = v19
                }
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 int32 = 2147483647
                    if (2 > 42) {
                        fmt.Println(v1, fi21)
                        v2 = (^(v2))
                        fmt.Println(v22, v22, v22)
                    } else {
                        var v23 uint16 = 2
                        var v24 uint = 65535
                        if ((fi21 & (fi21 / ((fi21) | 1))) <= fi21) {
                            for fi25 := 0; fi25 < 17; fi25++ {
                                if true {
                                    var v26 uint8 = (v5 % (((v5 & (v5 << (((v24 + 18446744073709551615)) & 63)))) | 1))
                                    v2 += v2
                                    _ = v26
                                } else {
                                    fmt.Println(fi25, v2)
                                    var v27 int64 = 0
                                    _ = v27
                                }
                                switch (-(fi25)) {
                                case 9223372036854775807:
                                    switch "hello" {
                                    case "hello":
                                        v28 := GenS3{F0: fi21}
                                        _ = v28
                                    case "\n\t":
                                        fmt.Println(v22, v2)
                                    case "foo bar":
                                        delete(v4, 899)
                                        _ = len(v3)
                                    default:
                                        var v29 uint16 = (v23 / (((v23 & (v23 + 666))) | 1))
                                        _ = v29
                                    }
                                case -2147483648:
                                    v22 = (v22 & (^(v22)))
                                    v30 := GenS4{F0: v23, F1: 4294967295, F2: v5}
                                    if true {
                                        switch (fi25 & fi25) {
                                        case -82:
                                            var v31 bool = true
                                            _ = v31
                                        case 32767:
                                            if ((v23 >= (^(v23))) && (false && ((27 < 65535) || false))) {
                                                var v32 int64 = v2
                                                v33 := GenS5{F0: int8(v23), F1: (v5 >> ((uint(v23)) & 63))}
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi21, fi21, v24)
                                        v24 = 4294967295
                                        v23 += v23
                                    }
                                    _ = v30
                                case 42:
                                    v34 := GenS6{F0: 42, F1: (-(v2)), F2: 2}
                                    _ = v34
                                }
                                var v35 int8 = 19
                                _ = fi25
                                _ = v35
                            }
                            for fi36 := 0; fi36 < 20; fi36++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v37 int64 = int64(fi21)
                                _ = fi36
                                _ = v37
                            }
                        } else {
                            delete(v4, 4294967295)
                        }
                        _ = v23
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                }
                var v38 bool = false
                _ = v38
            } else {
                _ = len(v4)
            }
            var v39 string = "foo bar"
            _ = v39
        } else {
            v2 = v2
        }
        var v40 int = -512
        _ = v40
    }
    if true {
        if (-9223372036854775808 < 2147483647) {
            switch 4294967295 {
            case -9223372036854775808:
                if true {
                    v41 := GenS1{F0: 436}
                    switch "\n\t" {
                    case "":
                        switch "" {
                        case "hello":
                            v42 := GenS7{F0: ("a" + "a"), F1: 0, F2: genFn1(genFn0(843), 955, (("hello" + "\n\t") + (("a" + "\n\t") + ("foo bar" + "hello"))))}
                            _ = v42
                        case "\n\t":
                            v5 = 208
                        default:
                            if false {
                                delete(v4, 872)
                                for fi43 := 0; fi43 < 9; fi43++ {
                                    switch "" {
                                    case "foo bar":
                                        v2 += -564
                                    default:
                                        if true {
                                            v2 += (v2 + v2)
                                            delete(v4, 348)
                                            v44 := GenS8{F0: 4294967295, F1: 18446744073709551615}
                                            _ = v44
                                        }
                                    }
                                    _ = fi43
                                }
                                v2 = v2
                            }
                        }
                    case "\n\t":
                        if false {
                            if false {
                                if false {
                                    fmt.Println(v5, v1)
                                    if (1 >= 0) {
                                        var v45 int8 = 127
                                        _ = v45
                                    } else {
                                        switch -9223372036854775808 {
                                        case 2147483647:
                                            fmt.Println(v5, v5)
                                        default:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    }
                                    for fi46 := 0; fi46 < 2; fi46++ {
                                        _ = len(v4)
                                        v2 += v2
                                        v1 = float64(1.5)
                                        _ = fi46
                                    }
                                }
                                var v47 int8 = 42
                                _ = v47
                            } else {
                                v5 -= (v5 / ((v5) | 1))
                                fmt.Println(v2)
                            }
                            v48 := GenS3{F0: -1}
                            v2 = 1
                            _ = v48
                        } else {
                            for fi49 := 0; fi49 < 13; fi49++ {
                                for fi50 := 0; fi50 < 9; fi50++ {
                                    var v51 int32 = 42
                                    v1 += v1
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi50
                                    _ = v51
                                }
                                var v52 uint64 = uint64(v5)
                                _ = fi49
                                _ = v52
                            }
                            v53 := GenS2{F0: (v2 ^ v2)}
                            var v54 bool = (1 <= (v5 << ((773) & 63)))
                            _ = v53
                            _ = v54
                        }
                    case "foo bar":
                        for fi55 := 0; fi55 < 4; fi55++ {
                            switch (fi55 & (^(fi55))) {
                            case -1:
                                if (v1 == v1) {
                                    if (65535 < 1) {
                                        for fi56 := 0; fi56 < 8; fi56++ {
                                            switch 128 {
                                            case 2147483647:
                                                v5 -= (v5 * 2)
                                            }
                                            v57 := GenS9{F0: v2, F1: 65535}
                                            _ = fi56
                                            _ = v57
                                        }
                                        v58 := GenS10{F0: -42, F1: ("foo bar" + ("foo bar" + ""))}
                                        if false {
                                            _ = len(v3)
                                            v59 := GenS2{F0: v2}
                                            v60 := GenS7{F0: ("a" + "a"), F1: fi55, F2: (v5 + (v5 / ((genFn1(2147483647, 255, "a")) | 1)))}
                                            _ = v59
                                            _ = v60
                                        }
                                        _ = v58
                                    }
                                } else {
                                    v5 -= (^(v5))
                                }
                            default:
                                switch (fi55 >> ((196) & 63)) {
                                case 4294967295:
                                    var v61 uint64 = 4294967295
                                    _ = v61
                                }
                            }
                            _ = fi55
                        }
                    default:
                        v5 -= (^(v5))
                    }
                    v62 := GenS11{F0: "a", F1: 1}
                    _ = v41
                    _ = v62
                } else {
                    v63 := GenS6{F0: 0, F1: 0, F2: 64}
                    _ = v63
                }
            case 32767:
                for fi64 := 0; fi64 < 14; fi64++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v65 uint64 = 482
                    _ = fi64
                    _ = v65
                }
            default:
                v66 := GenS4{F0: 412, F1: 42, F2: v5}
                _ = v66
            }
            v1 -= (1e10 - v1)
        } else {
            switch "" {
            case "a":
                var v67 int = 9223372036854775807
                _ = v67
            case "hello":
                v1 = v1
            default:
                switch "hello" {
                case "":
                    if true {
                        var v68 int16 = 32767
                        var v69 int16 = -32768
                        _ = v68
                        _ = v69
                    }
                case "hello":
                    for fi70 := 0; fi70 < 4; fi70++ {
                        v5 -= v5
                        var v71 uint16 = 0
                        _ = fi70
                        _ = v71
                    }
                case "foo bar":
                    v5 = v5
                default:
                    fmt.Println(v2)
                }
            }
            _ = len(v4)
        }
        var v72 int = -414
        fmt.Println(v1)
        _ = v72
    }
    v73 := GenS7{F0: "", F1: 127, F2: v5}
    v74 := v4[1]
    v2 = v2
    for fi75 := 0; fi75 < 6; fi75++ {
        var v76 uint64 = 0
        if false {
            if ((true && (1 >= v2)) && ((!true) && true)) {
                for fi77 := 0; fi77 < 13; fi77++ {
                    if true {
                        _ = len(v3)
                        if true {
                            fmt.Println(v2, v2, v2)
                            fmt.Println(v76)
                            v76 = v76
                            switch "\n\t" {
                            case "\n\t":
                                if false {
                                    v76 += 4294967295
                                    var v78 int = -806
                                    _ = v78
                                }
                            default:
                                if false {
                                    v2 = v2
                                    var v79 uint32 = 324
                                    _ = v79
                                }
                            }
                        }
                    } else {
                        v76 -= (v76 ^ (v76 >> ((65535) & 63)))
                    }
                    var v80 uint8 = genFn1(0, 0, v74)
                    var v81 string = v74
                    _ = fi77
                    _ = v80
                    _ = v81
                }
                v82 := GenS12{F0: 574}
                var v83 bool = (1 >= 2147483647)
                _ = v82
                _ = v83
            } else {
                if (1e10 <= 0.1) {
                    for fi84 := 0; fi84 < 20; fi84++ {
                        v76 = v76
                        v85 := GenS13{F0: 583}
                        v76 = 18446744073709551615
                        _ = len(v3)
                        var v86 int32 = -242
                        _ = fi84
                        _ = v85
                        _ = v86
                    }
                    v5 = (v5 | v5)
                    var v87 uint = 10
                    _ = v87
                } else {
                    v1 -= (1e10 + v1)
                }
                for fi88 := 0; fi88 < 6; fi88++ {
                    switch 28 {
                    case 993:
                        var v89 int8 = -4
                        _ = v89
                    case -2147483648:
                        if true {
                            var v90 uint8 = (v5 / ((87) | 1))
                            v76 += 18446744073709551615
                            _ = v90
                        }
                    case -177:
                        fmt.Println(v1, v76, v2)
                    }
                    _ = fi88
                }
            }
            switch fi75 {
            case 260:
                _ = len(v4)
            case 518:
                switch "hello" {
                case "":
                    v3 = append(v3, (fi75 >> ((686) & 63)))
                case "foo bar":
                    v91 := GenS14{F0: 521, F1: 65535}
                    _ = v91
                default:
                    fmt.Println(v76, fi75, v5)
                }
            case 0:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v92 := GenS3{F0: -9223372036854775808}
                _ = v92
            }
        } else {
            if (488 != 18446744073709551615) {
                fmt.Println(v76, v1, v76)
            } else {
                if (((v5 <= genFn1(0, 2, "foo bar")) || false) && (true && true)) {
                    for fi93 := 0; fi93 < 4; fi93++ {
                        v2 = v2
                        _ = fi93
                    }
                }
                switch fi75 {
                case 0:
                    var v94 float64 = float64(0.0)
                    _ = v94
                case 113:
                    var v95 float32 = (((1e10 - 1.5) * 3.14159) - -2.25)
                    _ = v95
                default:
                    v2 += -489
                }
                v96 := GenS8{F0: (v76 - 2), F1: 675}
                _ = v96
            }
            if (((v1 + (0.0 - v1)) - v1) >= (v1 / float64((0.1 - 0.0)))) {
                if (65535 > 718) {
                    if ((65535 <= 65535) && true) {
                        switch "\n\t" {
                        case "":
                            switch fi75 {
                            case 32768:
                                v2 += -9223372036854775808
                            case 301:
                                for fi97 := 0; fi97 < 19; fi97++ {
                                    if false {
                                        if (65535 == 2147483648) {
                                            var v98 float32 = 1.5
                                            switch 4294967295 {
                                            case 798:
                                                v99 := GenS6{F0: -1, F1: v2, F2: v76}
                                                _ = v99
                                            case 94:
                                                v98 += v98
                                            case 42:
                                                switch (^(fi75)) {
                                                case 2147483647:
                                                    v2 -= v2
                                                }
                                            default:
                                                switch int(v76) {
                                                case -727:
                                                    fmt.Println(v2, v74, fi97)
                                                default:
                                                    v100 := v4[65535]
                                                    _ = v100
                                                }
                                            }
                                            var v101 int = (fi75 ^ (fi75 % ((4294967295) | 1)))
                                            _ = v101
                                            _ = v98
                                        } else {
                                            fmt.Println(v1, v1)
                                        }
                                        var v102 int = fi97
                                        var v103 uint = 42
                                        _ = v102
                                        _ = v103
                                    } else {
                                        var v104 float32 = 3.14159
                                        if (1.5 <= 0.0) {
                                            switch (fi75 - fi75) {
                                            case 32767:
                                                v105 := GenS1{F0: 32767}
                                                _ = v105
                                            }
                                        } else {
                                            v2 += (v2 | v2)
                                        }
                                        _ = v104
                                    }
                                    v5 -= (v5 * v5)
                                    v106 := GenS8{F0: v76, F1: 211}
                                    _ = fi97
                                    _ = v106
                                }
                            default:
                                for fi107 := 0; fi107 < 17; fi107++ {
                                    if len(v3) > 0 {
                                        v3[0] = fi107
                                    }
                                    _ = fi107
                                }
                            }
                        }
                        v74 = ""
                    }
                    var v108 int8 = 100
                    _ = v108
                } else {
                    v5 += v5
                    var v109 int64 = (-(v2))
                    _ = v109
                }
                v76 = (v76 % ((562) | 1))
            }
            fmt.Println(v74, v5)
        }
        v1 -= 0.0
        _ = fi75
        _ = v76
    }
    for fi110 := 0; fi110 < 18; fi110++ {
        v1 = 0.0
        v111 := GenS15{F0: 18446744073709551615, F1: ("a" + v74), F2: (v2 - (v2 ^ v2))}
        var v112 uint16 = 0
        _ = fi110
        _ = v111
        _ = v112
    }
    fmt.Println(v5)
    var v113 float64 = v1
    v2 = v2
    if (2 <= 849) {
        _ = len(v3)
    } else {
        var v114 int16 = -32768
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v115 := GenS16{F0: -1, F1: true, F2: 2}
        _ = v114
        _ = v115
    }
    v2 += v2
    if (-52 != 89) {
        var v116 int32 = 2147483647
        _ = v116
    }
    var v117 float64 = float64(-2.25)
    v118 := GenS17{F0: (^(v5)), F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v113
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v73
    _ = v74
}
