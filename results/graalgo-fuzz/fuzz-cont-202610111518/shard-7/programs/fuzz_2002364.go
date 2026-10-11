package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint64; F1 int8; F2 int16 }
type GenS4 struct { F0 int16; F1 uint32; F2 int }
type GenS5 struct { F0 uint; F1 string; F2 bool }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 int32; F2 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 int8; F2 bool }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int64; F1 string; F2 int8 }
type GenS13 struct { F0 uint32 }

func main() {
    var v1 int = -1
    var v2 uint = 181
    v3 := []uint32{813, 2147483647, 524}
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    var v6 uint32 = 2
    if true {
        fmt.Println(v2, v1, v6)
        for fi7 := 0; fi7 < 11; fi7++ {
            if (!true) {
                v8 := GenS1{F0: 649}
                switch 696 {
                case 128:
                    fmt.Println(v6, fi7)
                case 32768:
                    if (1 > 2) {
                        v9 := v4[(fi7 ^ (fi7 * 9223372036854775807))]
                        _ = v9
                    }
                case 516:
                    for fi10 := 0; fi10 < 12; fi10++ {
                        var v11 int16 = 42
                        var v12 int64 = 9223372036854775807
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
                switch "hello" {
                case "a":
                    v1 = v1
                case "":
                    fmt.Println(v2, v6)
                default:
                    fmt.Println(v2)
                }
                _ = v8
            } else {
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := GenS1{F0: uint16(v6)}
                    _ = fi13
                    _ = v14
                }
                _ = len(v4)
            }
            _ = fi7
        }
    } else {
        if true {
            if (1.5 == -2.25) {
                for fi15 := 0; fi15 < 16; fi15++ {
                    v16 := GenS2{F0: 1, F1: 966, F2: 42}
                    _ = fi15
                    _ = v16
                }
                if len(v3) > 0 {
                    v3[0] = 718
                }
                switch "foo bar" {
                case "\n\t":
                    v17 := GenS1{F0: 192}
                    _ = v17
                case "a":
                    for fi18 := 0; fi18 < 18; fi18++ {
                        _ = len(v5)
                        for fi19 := 0; fi19 < 13; fi19++ {
                            var v20 bool = true
                            _ = fi19
                            _ = v20
                        }
                        v2 -= 2
                        _ = fi18
                    }
                }
            }
            var v21 int64 = 0
            _ = v21
        } else {
            switch "hello" {
            case "":
                var v22 uint16 = 466
                _ = v22
            }
            fmt.Println(v1, v2)
            fmt.Println(v1)
        }
        switch "hello" {
        case "foo bar":
            if (!(127 < 1)) {
                switch (^(v1)) {
                case -1:
                    var v23 int = v1
                    _ = v23
                case 127:
                    var v24 uint32 = (v6 / ((v6) | 1))
                    _ = v24
                case 385:
                    if (v1 != (-(v1))) {
                        if true {
                            var v25 float64 = 1e10
                            var v26 bool = true
                            _ = v25
                            _ = v26
                        } else {
                            if false {
                                fmt.Println(v2, v2, v2)
                            }
                            v1 -= (v1 / ((v1) | 1))
                        }
                        var v27 float32 = (1e10 - 0.0)
                        _ = v27
                    }
                default:
                    if (!(2 <= 49)) {
                        if (32767 == 442) {
                            v28 := GenS3{F0: 2, F1: 100, F2: 1}
                            if true {
                                v29 := GenS4{F0: 1, F1: v6, F2: v1}
                                fmt.Println(v1, v6, v1)
                                for fi30 := 0; fi30 < 6; fi30++ {
                                    var v31 string = ("foo bar" + "")
                                    v3 = append(v3, 715)
                                    _ = fi30
                                    _ = v31
                                }
                                _ = v29
                            }
                            for fi32 := 0; fi32 < 2; fi32++ {
                                fmt.Println(v6, v1, v1)
                                if (1 > 1) {
                                    v33 := GenS3{F0: 4294967295, F1: 6, F2: 1000}
                                    _ = v33
                                }
                                var v34 string = "foo bar"
                                _ = fi32
                                _ = v34
                            }
                            _ = v28
                        }
                    } else {
                        if (true && ((190 <= 26) && ((v6 + v6) > (v6 - (v6 * 42))))) {
                            v35 := GenS5{F0: 933, F1: ("\n\t" + ""), F2: true}
                            for fi36 := 0; fi36 < 4; fi36++ {
                                fmt.Println(fi36, v2)
                                fmt.Println(fi36, v6, v1)
                                switch ("foo bar" + ("\n\t" + "\n\t")) {
                                case "\n\t":
                                    var v37 int8 = 42
                                    _ = v37
                                case "a":
                                    v38 := GenS5{F0: v2, F1: ("foo bar" + "hello"), F2: (!true)}
                                    _ = v38
                                case "":
                                    v39 := v4[426]
                                    _ = v39
                                default:
                                    var v40 int64 = 0
                                    _ = v40
                                }
                                _ = fi36
                            }
                            _ = v35
                        } else {
                            for fi41 := 0; fi41 < 16; fi41++ {
                                v5["hello"] = v1
                                _ = fi41
                            }
                            _ = len(v3)
                        }
                    }
                }
                v42 := v4[v1]
                v43 := v4[v42]
                _ = v42
                _ = v43
            }
        case "\n\t":
            fmt.Println(v2, v6, v1)
        case "a":
            if (2 >= 4294967295) {
                for fi44 := 0; fi44 < 19; fi44++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi44
                }
                fmt.Println(v2, v2, v2)
                v2 = 65535
            } else {
                fmt.Println(v1)
                switch "\n\t" {
                case "foo bar":
                    v45 := GenS1{F0: 132}
                    _ = v45
                default:
                    var v46 uint64 = 43
                    _ = v46
                }
                v1 -= v1
            }
        }
    }
    var v47 string = "foo bar"
    for fi48 := 0; fi48 < 14; fi48++ {
        _ = len(v5)
        switch 0 {
        case 2147483647:
            v6 -= v6
        case 4294967295:
            var v49 int64 = 42
            _ = v49
        case -183:
            var v50 uint8 = 1
            _ = v50
        default:
            v2 = v2
        }
        for fi51 := 0; fi51 < 1; fi51++ {
            _ = len(v4)
            v52 := GenS6{F0: v6, F1: 389}
            fmt.Println(v6)
            _ = fi51
            _ = v52
        }
        _ = fi48
    }
    var v53 uint = v2
    var v54 float64 = 1.5
    delete(v4, (v1 ^ v1))
    v55 := GenS7{F0: v47}
    if (119 < 42) {
        if (-1 >= -1) {
            var v56 string = ""
            var v57 float64 = 0.0
            _ = v56
            _ = v57
        }
    } else {
        switch 2147483647 {
        case 0:
            if false {
                v54 += 1.5
                v58 := GenS7{F0: v47}
                _ = len(v5)
                _ = v58
            }
        case 1:
            if true {
                v1 -= v1
                var v59 int64 = 42
                _ = v59
            } else {
                _ = len(v4)
                switch "foo bar" {
                case "foo bar":
                    var v60 string = (("\n\t" + v47) + v47)
                    _ = v60
                case "a":
                    for fi61 := 0; fi61 < 1; fi61++ {
                        for fi62 := 0; fi62 < 17; fi62++ {
                            var v63 uint = (v2 + v53)
                            if (0 >= 0) {
                                if true {
                                    var v64 string = v47
                                    v65 := GenS5{F0: (v2 - (v2 << (((v2 ^ v2)) & 63))), F1: v47, F2: (0 == 42)}
                                    switch (((v64 + v47) + v64) + v47) {
                                    case "foo bar":
                                        v66 := GenS8{F0: v6, F1: 0, F2: 42}
                                        _ = v66
                                    case "a":
                                        var v67 int32 = 2147483647
                                        _ = v67
                                    default:
                                        var v68 uint16 = 1
                                        fmt.Println(fi61, fi62)
                                        if true {
                                            v47 = ("a" + ("foo bar" + "a"))
                                            v69 := GenS6{F0: (v6 | v6), F1: v53}
                                            v63 -= 18446744073709551615
                                            _ = v69
                                        }
                                        _ = v68
                                    }
                                    _ = v64
                                    _ = v65
                                }
                                v70 := GenS2{F0: v6, F1: 65535, F2: 1}
                                if false {
                                    v71 := GenS9{F0: int(v6)}
                                    var v72 int = fi61
                                    delete(v4, fi61)
                                    _ = v71
                                    _ = v72
                                }
                                _ = v70
                            } else {
                                fmt.Println(v47)
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi62
                            _ = v63
                        }
                        for fi73 := 0; fi73 < 2; fi73++ {
                            v54 = 1e10
                            fmt.Println(v53, v47)
                            _ = fi73
                        }
                        _ = fi61
                    }
                }
                v74 := v5[v47]
                _ = v74
            }
        case -9223372036854775808:
            if false {
                v3 = append(v3, v6)
                v75 := GenS10{F0: v1, F1: 0, F2: (v6 != 359)}
                _ = v75
            }
        default:
            var v76 int16 = -32768
            _ = v76
        }
        var v77 string = "\n\t"
        v78 := GenS11{F0: 686}
        _ = v77
        _ = v78
    }
    v53 = (v2 & 65535)
    if len(v3) > 0 {
        v3[0] = v6
    }
    if ((v54 / v54) >= ((v54 + (v54 - 0.1)) + (((0.1 - 3.14159) / (v54 * -2.25)) * v54))) {
        if (false && (1.5 > 0.1)) {
            if (!(int64(v6) <= -863)) {
                var v79 int16 = 1000
                v80 := GenS12{F0: -1, F1: v47, F2: int8(v6)}
                v81 := GenS8{F0: 42, F1: int32(v1), F2: v79}
                _ = v79
                _ = v80
                _ = v81
            } else {
                switch -9223372036854775808 {
                case 0:
                    var v82 uint8 = 194
                    _ = v82
                case -740:
                    for fi83 := 0; fi83 < 18; fi83++ {
                        for fi84 := 0; fi84 < 14; fi84++ {
                            v5["a"] = fi83
                            _ = fi84
                        }
                        v3 = append(v3, v6)
                        _ = fi83
                    }
                default:
                    var v85 uint8 = 2
                    _ = v85
                }
            }
            v53 = 18446744073709551615
            var v86 int32 = -562
            _ = v86
        }
    } else {
        fmt.Println(v47)
        _ = len(v4)
        fmt.Println(v1, v47, v1)
    }
    for fi87 := 0; fi87 < 18; fi87++ {
        var v88 float64 = -2.25
        for fi89 := 0; fi89 < 12; fi89++ {
            v54 = (v54 + v88)
            var v90 int8 = 0
            _ = fi89
            _ = v90
        }
        v91 := GenS13{F0: (v6 * (v6 - 2147483647))}
        _ = fi87
        _ = v88
        _ = v91
    }
    v1 = v1
    var v92 uint64 = uint64(v6)
    if false {
        var v93 int32 = 1
        switch "foo bar" {
        case "a":
            v93 += -883
        }
        fmt.Println(v1, v92)
        _ = v93
    } else {
        var v94 uint16 = 1
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v6
    _ = v92
}
