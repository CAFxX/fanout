package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int8; F2 int8 }
type GenS7 struct { F0 uint32; F1 uint8 }
func genFn0() int16 {
    var v7 int = 32767
    v8 := []int{(v7 * 4294967295), v7, (v7 ^ 42), 4294967295}
    var v9 int = (v7 * (v7 >> ((2) & 63)))
    v7 = 42
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []uint32{11, 440, 2147483648, 306}
    var v2 float32 = (1.5 / (1.5 / 0.1))
    var v3 uint8 = 2
    var v4 int16 = 32767
    var v5 int8 = 1
    v6 := []uint32{532, 1}
    if (!true) {
        for fi10 := 0; fi10 < 2; fi10++ {
            v1 = append(v1, 536)
            _ = fi10
        }
        for fi11 := 0; fi11 < 16; fi11++ {
            v2 = 0.0
            v12 := GenS1{F0: (v5 | (v5 + v5)), F1: (v4 % ((genFn0()) | 1)), F2: (v3 * (v3 | v3))}
            if (!false) {
                if true {
                    v2 = (v2 * (float32(-2.25) + v2))
                }
            }
            _ = fi11
            _ = v12
        }
        v3 -= v3
    }
    v5 += (^(v5))
    v5 += v5
    var v13 bool = ((0.1 >= v2) && (-647 < -159))
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 string = ""
            if false {
                var v17 uint32 = 2
                var v18 float32 = (v2 - 3.14159)
                var v19 bool = v13
                var v20 int = 128
                _ = v17
                _ = v18
                _ = v19
                _ = v20
            }
            if (!v13) {
                switch v16 {
                case "":
                    fmt.Println(v4)
                    var v21 string = (v16 + "a")
                    v22 := GenS2{F0: 2, F1: (v4 | (-(v4)))}
                    _ = v21
                    _ = v22
                case "a":
                    var v23 uint32 = 1
                    _ = v23
                default:
                    switch ("foo bar" + "a") {
                    case "":
                        v24 := GenS2{F0: 847, F1: v4}
                        _ = v24
                    case "foo bar":
                        if v13 {
                            v2 -= 0.1
                            if false {
                                v6 = append(v6, 2147483648)
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if v13 {
                                    v16 = v16
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    switch fi15 {
                                    case 0:
                                        if (v13 || ((435 >= 42) && (260 >= -240))) {
                                            v4 = v4
                                            fmt.Println(v13)
                                            var v25 float32 = float32(1e10)
                                            _ = v25
                                        }
                                    }
                                    var v26 float64 = ((((float64(float32(3.14159)) + (0.0 - float64(3.14159))) * (0.0 - ((float64(v2) / 0.1) * ((float64(3.14159) - float64(1.5)) - float64(float32((-2.25 / (3.14159 * (0.1 / -2.25))))))))) - ((1e10 * (1.5 / 1.5)) + 1e10)) - 1e10)
                                    v3 = (^(v3))
                                    if (2 == 255) {
                                        fmt.Println(v13, v4, fi15)
                                        if len(v6) > 0 {
                                            v6[0] = 433
                                        }
                                    }
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v26
                                } else {
                                    v27 := GenS2{F0: 18446744073709551615, F1: (v4 / ((v4) | 1))}
                                    _ = v27
                                }
                            } else {
                                if (v4 < v4) {
                                    fmt.Println(v4, v16)
                                    if false {
                                        v28 := GenS3{F0: 953, F1: v16, F2: -472}
                                        _ = v28
                                    } else {
                                        fmt.Println(v4)
                                        var v29 int32 = 731
                                        v3 -= v3
                                        _ = v29
                                    }
                                } else {
                                    switch "" {
                                    case "":
                                        var v30 uint16 = 1
                                        _ = v30
                                    case "foo bar":
                                        v6 = append(v6, 2)
                                    default:
                                        var v31 int32 = 2147483647
                                        _ = v31
                                    }
                                    var v32 int64 = 1
                                    _ = v32
                                }
                            }
                            var v33 int8 = v5
                            _ = v33
                        }
                    case "\n\t":
                        fmt.Println(v16, v2, v13)
                    }
                }
                _ = len(v1)
                var v34 float32 = 1e10
                _ = v34
            }
            v16 = (v16 + "\n\t")
            v35 := GenS3{F0: 255, F1: (v16 + "a"), F2: -2147483648}
            _ = fi15
            _ = v16
            _ = v35
        }
        v3 += v3
        _ = fi14
    }
    if (42 >= 348) {
        var v36 int32 = int32(v3)
        _ = v36
    }
    switch ("foo bar" + "hello") {
    case "foo bar":
        for fi37 := 0; fi37 < 14; fi37++ {
            if (!((v4 >> ((65535) & 63)) < v4)) {
                fmt.Println(v5, v5)
            }
            _ = fi37
        }
    case "\n\t":
        for fi38 := 0; fi38 < 13; fi38++ {
            if (2147483647 >= 0) {
                var v39 uint8 = (^(v3))
                fmt.Println(v4, fi38)
                _ = v39
            }
            _ = fi38
        }
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        if (!(!(0 >= fi40))) {
            var v41 float32 = (((1.5 / v2) * 3.14159) - v2)
            if v13 {
                var v42 int16 = v4
                if ("hello" == "hello") {
                    v43 := GenS4{F0: false}
                    _ = v43
                } else {
                    v3 += v3
                    var v44 uint8 = v3
                    _ = v44
                }
                _ = v42
            } else {
                v5 += v5
            }
            _ = v41
        }
        _ = fi40
    }
    fmt.Println(v3, v3)
    v45 := GenS4{F0: v13}
    var v46 uint64 = 0
    v47 := GenS2{F0: (^(v46)), F1: v4}
    _ = len(v6)
    if (int32(v3) <= 2147483647) {
        fmt.Println(v4)
        var v48 int = 42
        v49 := GenS4{F0: (65535 > 28)}
        _ = v48
        _ = v49
    } else {
        switch ("a" + "a") {
        case "\n\t":
            var v50 uint64 = (v46 | (v46 ^ v46))
            _ = v50
        case "hello":
            v51 := GenS5{F0: false}
            _ = v51
        }
        if (42 > int64(v3)) {
            v3 = v3
            v3 = 130
            switch -1 {
            case 557:
                if ((^(v46)) != 2) {
                    switch "a" {
                    case "a":
                        var v52 int8 = (v5 * v5)
                        _ = v52
                    case "hello":
                        switch "hello" {
                        case "hello":
                            v1 = append(v1, 2147483647)
                        case "foo bar":
                            var v53 uint32 = 0
                            _ = v53
                        case "":
                            var v54 uint16 = 736
                            _ = v54
                        default:
                            var v55 int8 = -73
                            _ = v55
                        }
                    case "foo bar":
                        var v56 int16 = genFn0()
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 19; fi57++ {
                        v13 = ("hello" != ("foo bar" + ("hello" + ("\n\t" + ""))))
                        _ = fi57
                    }
                }
            }
        }
    }
    switch -9223372036854775808 {
    case 660:
        fmt.Println(v3, v13)
    case 128:
        switch 920 {
        case 128:
            for fi58 := 0; fi58 < 16; fi58++ {
                fmt.Println(v13, v2)
                switch "foo bar" {
                case "foo bar":
                    var v59 int16 = v4
                    _ = v59
                case "\n\t":
                    for fi60 := 0; fi60 < 18; fi60++ {
                        var v61 uint32 = 988
                        _ = fi60
                        _ = v61
                    }
                case "hello":
                    if (255 >= 2) {
                        for fi62 := 0; fi62 < 3; fi62++ {
                            var v63 uint = 18446744073709551615
                            _ = fi62
                            _ = v63
                        }
                        var v64 int64 = 182
                        var v65 uint32 = 2147483647
                        _ = v64
                        _ = v65
                    }
                default:
                    if (-1 == 42) {
                        if (true && true) {
                            if v13 {
                                fmt.Println(v5, v13, v3)
                                v66 := GenS2{F0: v46, F1: genFn0()}
                                _ = v66
                            } else {
                                if v13 {
                                    v46 += (v46 % ((42) | 1))
                                    v67 := GenS2{F0: (v46 - v46), F1: v4}
                                    var v68 int64 = -9223372036854775808
                                    _ = v67
                                    _ = v68
                                } else {
                                    v13 = (v13 || v13)
                                    v46 = 326
                                    if (true && (!false)) {
                                        v69 := GenS5{F0: (v5 > 127)}
                                        var v70 int = 128
                                        _ = v69
                                        _ = v70
                                    }
                                }
                            }
                            if v13 {
                                var v71 string = "hello"
                                if len(v1) > 0 {
                                    v1[0] = 4294967295
                                }
                                var v72 int64 = int64(v46)
                                _ = v71
                                _ = v72
                            } else {
                                if ((v3 << ((928) & 63)) <= v3) {
                                    v73 := GenS5{F0: (0 < -1)}
                                    var v74 uint = 857
                                    for fi75 := 0; fi75 < 5; fi75++ {
                                        v76 := GenS5{F0: (v13 && ((^(v74)) > (v74 - (v74 | (v74 >> (((v74 - (v74 * v74))) & 63))))))}
                                        v77 := GenS2{F0: (v46 + v46), F1: v4}
                                        if ((v46 << (((v74 | (v74 | v74))) & 63)) > (^(v46))) {
                                            v78 := GenS1{F0: v5, F1: v4, F2: (v3 % (((v3 % (((^(v3))) | 1))) | 1))}
                                            if len(v1) > 0 {
                                                v1[0] = 2147483648
                                            }
                                            v2 += v2
                                            _ = v78
                                        } else {
                                            var v79 uint32 = uint32(v46)
                                            if len(v6) > 0 {
                                                v6[0] = (v79 ^ (v79 % ((v79) | 1)))
                                            }
                                            _ = v79
                                        }
                                        _ = fi75
                                        _ = v76
                                        _ = v77
                                    }
                                    _ = v73
                                    _ = v74
                                } else {
                                    fmt.Println(v46, fi58, v46)
                                    v46 -= v46
                                }
                            }
                        } else {
                            for fi80 := 0; fi80 < 17; fi80++ {
                                if v13 {
                                    v3 -= (v3 * v3)
                                    if len(v6) > 0 {
                                        v6[0] = 4294967295
                                    }
                                    v1 = append(v1, 2147483648)
                                    v81 := GenS5{F0: (2147483647 > -2147483648)}
                                    _ = v81
                                }
                                v13 = (2 > 42)
                                v2 += v2
                                _ = fi80
                            }
                            v4 = v4
                            if (817 == 0) {
                                var v82 uint = 18446744073709551615
                                for fi83 := 0; fi83 < 20; fi83++ {
                                    v5 = (v5 >> (((^(v82))) & 63))
                                    var v84 int32 = -573
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v3)
                                    case "":
                                        v85 := GenS5{F0: (787 != 2147483647)}
                                        _ = v85
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "foo bar":
                                            v84 = v84
                                        case "":
                                            v3 = v3
                                        default:
                                            var v86 int16 = v4
                                            _ = v86
                                        }
                                    default:
                                        v13 = (v13 && false)
                                    }
                                    _ = fi83
                                    _ = v84
                                }
                                switch 4294967295 {
                                case -869:
                                    var v87 int8 = v5
                                    _ = v87
                                case 98:
                                    fmt.Println(v46, v46, v82)
                                case -1:
                                    if (v46 == (^(v46))) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        if v13 {
                                            v88 := GenS6{F0: ("" + ("\n\t" + "hello")), F1: -31, F2: 42}
                                            _ = v88
                                        }
                                        var v89 int16 = v4
                                        _ = v89
                                    }
                                }
                                _ = v82
                            }
                        }
                        fmt.Println(v46)
                    }
                }
                _ = fi58
            }
        case 968:
            if true {
                if v13 {
                    for fi90 := 0; fi90 < 15; fi90++ {
                        var v91 int64 = 9223372036854775807
                        _ = fi90
                        _ = v91
                    }
                    fmt.Println(v13)
                } else {
                    for fi92 := 0; fi92 < 15; fi92++ {
                        fmt.Println(v2, v2, v13)
                        var v93 int32 = -610
                        _ = fi92
                        _ = v93
                    }
                    switch 42 {
                    case -9223372036854775808:
                        if v13 {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 280
                            }
                        }
                    case -85:
                        for fi94 := 0; fi94 < 7; fi94++ {
                            v4 -= -666
                            var v95 int32 = 0
                            _ = fi94
                            _ = v95
                        }
                    default:
                        for fi96 := 0; fi96 < 16; fi96++ {
                            v5 = (v5 - v5)
                            _ = fi96
                        }
                    }
                    if ((((float32(1.5) + 1.5) * (v2 + 0.1)) - v2) != (v2 / -2.25)) {
                        for fi97 := 0; fi97 < 19; fi97++ {
                            v46 -= v46
                            for fi98 := 0; fi98 < 8; fi98++ {
                                v6 = append(v6, 0)
                                if len(v1) > 0 {
                                    v1[0] = 404
                                }
                                _ = fi98
                            }
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                            _ = fi97
                        }
                        if (!((true || (false || v13)) || v13)) {
                            v13 = v13
                            v2 = v2
                            v5 = int8(v4)
                        } else {
                            v3 += (^(v3))
                            v99 := GenS7{F0: 4294967295, F1: v3}
                            v46 += (v46 + 175)
                            _ = v99
                        }
                    } else {
                        var v100 uint16 = 721
                        var v101 uint = 65535
                        _ = v100
                        _ = v101
                    }
                }
                var v102 uint64 = 4294967295
                _ = v102
            }
        case 9223372036854775807:
            switch 0 {
            case -466:
                var v103 bool = ((0 <= 0) && true)
                _ = v103
            }
        }
    case -1:
        v46 -= v46
    default:
        var v104 string = "hello"
        _ = v104
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
