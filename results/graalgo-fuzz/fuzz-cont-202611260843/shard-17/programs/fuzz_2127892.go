package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint64; F2 string }
type GenS4 struct { F0 int32; F1 int8; F2 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 int8; F2 int8 }
type GenS7 struct { F0 string; F1 uint64; F2 int16 }
type GenS8 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS9 struct { F0 bool; F1 string; F2 uint64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int16; F1 uint32 }
type GenS12 struct { F0 string; F1 uint32; F2 uint8 }
type GenS13 struct { F0 uint32; F1 int8; F2 uint64 }
func genFn0(p0 bool, p1 uint32, p2 int64) uint64 {
    var v5 bool = (p0 && (((((!p0) && (-9223372036854775808 <= 4294967295)) || p0) || (p0 && true)) || true))
    var v6 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 232
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -9223372036854775808
    var v3 int = 4294967295
    v4 := []int{(v3 * 127), v3, v3}
    var v7 int64 = (v2 & (^(v2)))
    if (0 != v7) {
        if len(v4) > 0 {
            v4[0] = (v3 ^ (v3 * v3))
        }
        var v8 int8 = -7
        _ = v8
    } else {
        v9 := GenS1{F0: 42}
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 uint8 = 170
                v13 := GenS1{F0: 1}
                v12 += (^(v12))
                _ = fi11
                _ = v12
                _ = v13
            }
            if (true && (175 != 4294967295)) {
                var v14 uint16 = 253
                for fi15 := 0; fi15 < 17; fi15++ {
                    if (-4 != -851) {
                        v14 -= 65535
                        if (!(159 >= 67)) {
                            _ = len(v4)
                        }
                        fmt.Println(v14, v2)
                        v16 := GenS1{F0: 0}
                        var v17 int16 = -32768
                        if (0.1 != (float32(1.5) + 1e10)) {
                            v7 += v2
                            v17 -= 747
                            v7 = (v2 + v7)
                        }
                        _ = v16
                        _ = v17
                    } else {
                        v7 += -641
                        var v18 int = v3
                        var v19 float64 = 0.0
                        _ = v18
                        _ = v19
                    }
                    switch (fi10 - (fi10 % (((v3 / ((v3) | 1))) | 1))) {
                    case 42:
                        v20 := GenS1{F0: 1}
                        _ = v20
                    case 32768:
                        if len(v4) > 0 {
                            v4[0] = (fi10 << ((4294967295) & 63))
                        }
                    }
                    switch (v3 - 42) {
                    case 738:
                        if (-266 <= int64(fi15)) {
                            v21 := GenS2{F0: v14}
                            if len(v4) > 0 {
                                v4[0] = fi10
                            }
                            v2 -= (v7 / ((-81) | 1))
                            v3 = fi15
                            if false {
                                v14 = (v14 / (((v14 >> ((4294967295) & 63))) | 1))
                                fmt.Println(v14)
                                v22 := GenS1{F0: -1}
                                _ = v22
                            } else {
                                v4 = append(v4, (fi10 / ((fi15) | 1)))
                            }
                            switch (fi10 + v3) {
                            case 42:
                                v23 := GenS1{F0: -1}
                                _ = v23
                            case -2147483648:
                                v2 = 9223372036854775807
                            default:
                                if ((-1 != 779) && false) {
                                    fmt.Println(v2)
                                    var v24 bool = false
                                    switch (fi10 - (fi10 << ((0) & 63))) {
                                    case 398:
                                        var v25 float64 = 1e10
                                        _ = v25
                                    case 0:
                                        v24 = v24
                                    }
                                    _ = v24
                                } else {
                                    v26 := GenS2{F0: (v14 + 819)}
                                    _ = v26
                                }
                            }
                            if (!(!true)) {
                                var v27 string = ("a" + "\n\t")
                                var v28 uint = 18446744073709551615
                                var v29 float32 = 0.1
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                v2 = 1
                                switch (fi10 ^ fi10) {
                                case 0:
                                    var v30 uint16 = (v14 ^ v14)
                                    _ = v30
                                case 32768:
                                    v7 += v2
                                default:
                                    if false {
                                        fmt.Println(fi15, fi10)
                                        v7 -= v2
                                        var v31 int16 = 42
                                        v14 += 255
                                        v32 := v1[0]
                                        _ = len(v4)
                                        v14 = (v14 + 420)
                                        _ = v31
                                        _ = v32
                                    }
                                    if (!true) {
                                        v14 = v14
                                        var v33 float64 = 0.0
                                        delete(v1, 2147483647)
                                        _ = v33
                                    } else {
                                        v3 += v3
                                    }
                                    v2 += v2
                                    var v34 bool = (0.0 == 1e10)
                                    _ = v34
                                }
                            }
                            _ = v21
                        } else {
                            fmt.Println(v2, fi10, v3)
                        }
                    case 42:
                        v14 -= (v14 + 6)
                    case 1:
                        if true {
                            fmt.Println(v7)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            if len(v4) > 0 {
                                v4[0] = 2147483647
                            }
                            var v35 int = 1
                            v36 := GenS1{F0: 103}
                            switch -511 {
                            case 177:
                                switch "foo bar" {
                                case "foo bar":
                                    v3 = (fi10 - v3)
                                case "\n\t":
                                    v1[0] = -136
                                case "hello":
                                    fmt.Println(fi10, fi10, v7)
                                default:
                                    v37 := GenS1{F0: 1}
                                    _ = v37
                                }
                            case 9223372036854775807:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 0:
                                if false {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                } else {
                                    var v38 uint = 18446744073709551615
                                    var v39 uint = v38
                                    switch "a" {
                                    case "hello":
                                        if (false || (v3 < 9223372036854775807)) {
                                            fmt.Println(v39, v35, v3)
                                        } else {
                                            if (false || ((v7 & (v7 | (v7 ^ (^(v7))))) <= (v7 - (v2 << (((v39 % ((v38) | 1))) & 63))))) {
                                                fmt.Println(fi10, v35, v35)
                                                v4 = append(v4, fi15)
                                            }
                                        }
                                    case "a":
                                        fmt.Println(v38)
                                    case "\n\t":
                                        if ((v14 * 605) < v14) {
                                            switch fi10 {
                                            case -569:
                                                if true {
                                                    v40 := GenS3{F0: false, F1: 0, F2: "foo bar"}
                                                    v41 := GenS4{F0: -79, F1: 0, F2: "hello"}
                                                    _ = v40
                                                    _ = v41
                                                } else {
                                                    if (-999 <= 0) {
                                                        v42 := GenS5{F0: 1}
                                                        v7 += (v2 & v2)
                                                        var v43 uint64 = 42
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        v3 = fi15
                                                        v3 = 4294967295
                                                    }
                                                    _ = len(v4)
                                                    if (865 >= 588) {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        switch fi15 {
                                                        case 499:
                                                            var v44 uint16 = v14
                                                            _ = v44
                                                        case 32768:
                                                            switch "\n\t" {
                                                            case "hello":
                                                                fmt.Println(fi15, v14, fi10)
                                                            case "a":
                                                                var v45 uint32 = 449
                                                                _ = v45
                                                            default:
                                                                fmt.Println(fi10, v35, v7)
                                                                var v46 uint8 = 42
                                                                v2 = v2
                                                                _ = v46
                                                            }
                                                        default:
                                                            if false {
                                                                v35 += v35
                                                            } else {
                                                                fmt.Println(fi10, v38)
                                                                if true {
                                                                    if (!true) {
                                                                        switch (v35 / (((v35 >> ((v39) & 63))) | 1)) {
                                                                        case 2147483647:
                                                                            v14 = 1
                                                                        case -9223372036854775808:
                                                                            v2 += (v7 >> (((^(v39))) & 63))
                                                                            v47 := GenS3{F0: (42 > 1), F1: 330, F2: "hello"}
                                                                            fmt.Println(v35, fi15, v39)
                                                                            v48 := GenS1{F0: 100}
                                                                            switch v35 {
                                                                            case -561:
                                                                                switch "" {
                                                                                case "hello":
                                                                                    var v49 float64 = 0.0
                                                                                    _ = v49
                                                                                case "\n\t":
                                                                                    v50 := GenS6{F0: 821, F1: -1, F2: -1}
                                                                                    _ = v50
                                                                                case "foo bar":
                                                                                    switch ("hello" + ("\n\t" + "\n\t")) {
                                                                                    case "hello":
                                                                                        if ((false || false) && true) {
                                                                                            var v51 bool = ((fi10 > v3) || ((v38 & 4294967295) > (^(v39))))
                                                                                            _ = v51
                                                                                        } else {
                                                                                            if ((("" + "hello") + "hello") > "foo bar") {
                                                                                                var v52 int16 = 461
                                                                                                fmt.Println(v52, v14, fi15)
                                                                                                v35 = 32768
                                                                                                _ = v52
                                                                                            }
                                                                                            if true {
                                                                                                _ = len(v4)
                                                                                                if ("hello" <= "hello") {
                                                                                                    var v53 float64 = 1e10
                                                                                                    var v54 int8 = 127
                                                                                                    _ = len(v1)
                                                                                                    v35 = 4294967295
                                                                                                    if len(v4) >= 2 {
                                                                                                        v4 = v4[1:len(v4)]
                                                                                                    }
                                                                                                    var v55 uint16 = (v14 ^ (v14 - v14))
                                                                                                    v35 += fi10
                                                                                                    _ = v53
                                                                                                    _ = v54
                                                                                                    _ = v55
                                                                                                }
                                                                                                if true {
                                                                                                    v2 += int64(fi10)
                                                                                                    delete(v1, 2)
                                                                                                    var v56 int64 = (v7 << (((v38 * (^(v38)))) & 63))
                                                                                                    v3 = 127
                                                                                                    fmt.Println(fi10, v56, v7)
                                                                                                    _ = v56
                                                                                                } else {
                                                                                                    v57 := GenS7{F0: ("hello" + ("hello" + "\n\t")), F1: 2, F2: -32768}
                                                                                                    fmt.Println(v14, v38, v39)
                                                                                                    _ = v57
                                                                                                }
                                                                                            } else {
                                                                                                _ = len(v1)
                                                                                                v58 := GenS2{F0: (^(v14))}
                                                                                                _ = v58
                                                                                            }
                                                                                        }
                                                                                    case "":
                                                                                        var v59 string = "\n\t"
                                                                                        _ = v59
                                                                                    case "foo bar":
                                                                                        v1[2147483647] = (fi15 >> ((18446744073709551615) & 63))
                                                                                    default:
                                                                                        var v60 float64 = 1.5
                                                                                        _ = v60
                                                                                    }
                                                                                default:
                                                                                    var v61 uint64 = 0
                                                                                    _ = v61
                                                                                }
                                                                            }
                                                                            _ = v47
                                                                            _ = v48
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v7, v35, fi10)
                                                                    }
                                                                    if len(v4) > 0 {
                                                                        v4[0] = (fi15 << ((uint(v14)) & 63))
                                                                    }
                                                                } else {
                                                                    fmt.Println(v7, v3)
                                                                    v7 = v2
                                                                }
                                                                v14 = 804
                                                            }
                                                            v38 = (v39 % ((4294967295) | 1))
                                                            switch "" {
                                                            case "":
                                                                if (v2 != v2) {
                                                                    fmt.Println(v7, v35)
                                                                } else {
                                                                    fmt.Println(v38, v2)
                                                                    if false {
                                                                        var v62 float64 = (3.14159 / ((-2.25 + 1.5) + (0.1 / (((0.0 - 0.1) * (-2.25 * 1e10)) + 0.1))))
                                                                        v63 := GenS1{F0: -1}
                                                                        _ = v62
                                                                        _ = v63
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                var v64 int = v35
                                                                _ = v64
                                                            case "a":
                                                                v65 := v1[472]
                                                                _ = v65
                                                            default:
                                                                v66 := GenS8{F0: 127, F1: (v14 * (v14 ^ v14)), F2: v14}
                                                                _ = v66
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            fmt.Println(v2, v2, v35)
                                            _ = len(v1)
                                        } else {
                                            if (!("\n\t" == ("foo bar" + "\n\t"))) {
                                                fmt.Println(v3, v39, v39)
                                                var v67 uint32 = 176
                                                var v68 uint16 = (^(v14))
                                                v69 := GenS8{F0: 127, F1: v68, F2: 880}
                                                var v70 float32 = 0.0
                                                if len(v4) > 0 {
                                                    v4[0] = fi10
                                                }
                                                _ = v67
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                            } else {
                                                if (!(2147483647 >= 42)) {
                                                    switch ("a" + ("\n\t" + "\n\t")) {
                                                    case "hello":
                                                        v71 := GenS9{F0: (true || ((-(v3)) < (v3 ^ (fi10 ^ fi15)))), F1: (("a" + "") + ""), F2: 42}
                                                        _ = v71
                                                    case "foo bar":
                                                        fmt.Println(v38, v35, v38)
                                                    default:
                                                        fmt.Println(v38)
                                                    }
                                                    delete(v1, 65535)
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi10 | v3)
                                                    }
                                                }
                                                v7 -= -592
                                            }
                                            if false {
                                                switch fi10 {
                                                case 127:
                                                    var v72 uint8 = 127
                                                    _ = v72
                                                case 9223372036854775807:
                                                    switch (v3 | (v3 ^ 752)) {
                                                    case 32767:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    case 0:
                                                        switch (v3 + (fi15 * -9223372036854775808)) {
                                                        case 4294967295:
                                                            var v73 int8 = 42
                                                            delete(v1, 642)
                                                            var v74 float32 = (float32(float64(0.1)) * 1e10)
                                                            _ = v73
                                                            _ = v74
                                                        default:
                                                            v39 = 0
                                                        }
                                                    case 32768:
                                                        v3 += v3
                                                    default:
                                                        if false {
                                                            v1[2] = (v35 - fi10)
                                                            delete(v1, 535)
                                                            v35 -= 128
                                                        } else {
                                                            v75 := GenS10{F0: genFn0((!false), 4294967295, v7)}
                                                            switch "" {
                                                            case "":
                                                                var v76 uint = v38
                                                                _ = v76
                                                            default:
                                                                v77 := GenS8{F0: -121, F1: (^(v14)), F2: v14}
                                                                _ = v77
                                                            }
                                                            _ = v75
                                                        }
                                                    }
                                                }
                                                v78 := GenS5{F0: 127}
                                                delete(v1, 2147483648)
                                                _ = v78
                                            }
                                            var v79 uint = (v39 >> (((v38 & (v39 | (v39 + uint(v14))))) & 63))
                                            _ = v79
                                        }
                                    }
                                    _ = v38
                                    _ = v39
                                }
                                fmt.Println(v14, v7, fi10)
                                v14 = v14
                                v2 = (v2 ^ int64(v14))
                                v7 -= 0
                            default:
                                var v80 int8 = 100
                                _ = v80
                            }
                            _ = len(v4)
                            var v81 int8 = 127
                            var v82 uint = 264
                            _ = v35
                            _ = v36
                            _ = v81
                            _ = v82
                        }
                    }
                    _ = fi15
                }
                fmt.Println(v7, fi10)
                _ = v14
            }
            if (("hello" + "a") >= "\n\t") {
                var v83 int32 = -1
                _ = v83
            }
            _ = fi10
        }
        _ = v9
    }
    if (1.5 < float32((-2.25 - 0.1))) {
        for fi84 := 0; fi84 < 12; fi84++ {
            if false {
                if true {
                    v2 = (v2 << ((885) & 63))
                    v85 := GenS8{F0: -1, F1: 1, F2: 65535}
                    switch 324 {
                    case 42:
                        v86 := GenS7{F0: (("\n\t" + ("foo bar" + "a")) + ("\n\t" + "foo bar")), F1: 0, F2: -32768}
                        _ = v86
                    case -1:
                        for fi87 := 0; fi87 < 2; fi87++ {
                            v88 := GenS11{F0: 814, F1: 42}
                            if (-1 > 32767) {
                                var v89 uint8 = 255
                                var v90 uint32 = 945
                                var v91 int64 = 9223372036854775807
                                var v92 int = (fi87 * fi84)
                                fmt.Println(v89, v91, v91)
                                delete(v1, v90)
                                _ = v89
                                _ = v90
                                _ = v91
                                _ = v92
                            }
                            _ = fi87
                            _ = v88
                        }
                    case 4294967295:
                        v1[2] = fi84
                    default:
                        v93 := GenS3{F0: true, F1: 0, F2: ("hello" + "\n\t")}
                        _ = v93
                    }
                    _ = v85
                }
                v7 += (v2 << ((65535) & 63))
                v94 := v1[42]
                _ = v94
            } else {
                if len(v4) > 0 {
                    v4[0] = -445
                }
            }
            fmt.Println(v2)
            v2 += (v7 >> ((0) & 63))
            _ = fi84
        }
        v4 = append(v4, (^(v3)))
        if (42 <= -1) {
            var v95 bool = ((977 <= 2) && (!(-2.25 != (1e10 * 1.5))))
            v95 = false
            _ = v95
        } else {
            v96 := GenS5{F0: 128}
            var v97 uint32 = 943
            _ = v96
            _ = v97
        }
    } else {
        var v98 uint16 = 53
        if false {
            v99 := GenS12{F0: "foo bar", F1: 198, F2: 97}
            var v100 uint8 = 117
            _ = v100
            _ = v99
        } else {
            for fi101 := 0; fi101 < 16; fi101++ {
                if false {
                    v102 := GenS6{F0: 65535, F1: 42, F2: 127}
                    v98 += 255
                    v103 := GenS1{F0: -1}
                    _ = v102
                    _ = v103
                }
                v7 = v7
                fmt.Println(fi101)
                _ = fi101
            }
        }
        v2 -= (v7 ^ 215)
        _ = v98
    }
    if false {
        v104 := GenS10{F0: 240}
        var v105 bool = true
        _ = v104
        _ = v105
    } else {
        var v106 bool = true
        _ = v106
    }
    fmt.Println(v2, v3)
    for fi107 := 0; fi107 < 9; fi107++ {
        fmt.Println(fi107)
        v3 = (fi107 * (fi107 + (fi107 * v3)))
        fmt.Println(v3, v2, fi107)
        _ = fi107
    }
    var v108 float32 = float32(3.14159)
    fmt.Println(v7, v3)
    v109 := GenS13{F0: 189, F1: 0, F2: 710}
    var v110 int64 = v2
    _ = len(v4)
    v108 += v108
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if false {
            var v111 float64 = (1e10 * ((0.0 + float64(0.1)) - 0.1))
            fmt.Println(v110)
            _ = v111
        } else {
            var v112 uint64 = 683
            _ = v112
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v108:", v108)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
