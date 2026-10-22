package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS7 struct { F0 int; F1 uint8; F2 uint }
type GenS8 struct { F0 uint8; F1 int16; F2 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8; F1 uint }
type GenS11 struct { F0 bool; F1 int8; F2 uint16 }
type GenS12 struct { F0 uint8; F1 int16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    v4 := make(map[int]string)
    v5 := GenS1{F0: 57, F1: 803, F2: true}
    var v6 uint16 = 2
    v7 := GenS1{F0: 208, F1: 127, F2: (!(!true))}
    if (1 >= 100) {
        fmt.Println(v6)
        switch -2147483648 {
        case 4294967295:
            var v8 int16 = 0
            _ = v8
        case -1:
            if ((146 >= 1000) && false) {
                v6 -= (v6 / ((v6) | 1))
            }
        default:
            var v9 int16 = 42
            _ = v9
        }
    } else {
        var v10 uint16 = (v6 ^ 2)
        _ = v10
    }
    var v11 float32 = (0.0 * ((0.0 * 0.0) * 1.5))
    fmt.Println(v11, v11)
    if true {
        var v12 int64 = 1
        _ = v12
    }
    v11 += (v11 / v11)
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            if true {
                if (false || false) {
                    if true {
                        fmt.Println(v6)
                        var v15 uint = 2
                        _ = v15
                    }
                    v3 = append(v3, 2)
                    var v16 uint = 965
                    if false {
                        v6 -= (v6 * uint16(v16))
                        switch "" {
                        case "a":
                            if false {
                                v17 := GenS2{F0: 128}
                                v18 := GenS1{F0: 128, F1: (fi13 / ((fi14) | 1)), F2: false}
                                if (-598 == (fi13 % ((fi13) | 1))) {
                                    v6 = 32767
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        var v19 uint8 = 128
                                        if ((83 > -68) && true) {
                                            v11 += ((0.0 * v11) / float32(float64(((v11 - 0.0) / 1.5))))
                                            v20 := GenS3{F0: 2, F1: 164}
                                            if len(v3) > 0 {
                                                v3[0] = 512
                                            }
                                            _ = v20
                                        } else {
                                            var v21 uint64 = 42
                                            _ = v21
                                        }
                                        var v22 float64 = 0.1
                                        switch (fi13 * fi14) {
                                        case -503:
                                            v23 := GenS2{F0: 2}
                                            _ = v23
                                        case 0:
                                            switch "" {
                                            case "\n\t":
                                                var v24 string = "hello"
                                                _ = v24
                                            case "hello":
                                                v19 += uint8(v16)
                                            case "foo bar":
                                                v25 := GenS2{F0: v19}
                                                _ = v25
                                            }
                                        case -821:
                                            v26 := GenS2{F0: 128}
                                            _ = v26
                                        default:
                                            switch (fi13 & 42) {
                                            case 1:
                                                switch "foo bar" {
                                                case "a":
                                                    var v27 uint64 = 235
                                                    _ = v27
                                                case "\n\t":
                                                    if (-81 <= 541) {
                                                        v28 := GenS4{F0: 53, F1: -9223372036854775808}
                                                        v2[946] = "hello"
                                                        v29 := GenS5{F0: 32767, F1: 1000}
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        v11 += float32(3.14159)
                                                        v16 += 814
                                                    }
                                                default:
                                                    var v30 uint = v16
                                                    fmt.Println(v16)
                                                    _ = v30
                                                }
                                                fmt.Println(v11, v19, v6)
                                                switch "\n\t" {
                                                case "a":
                                                    var v31 uint32 = 2147483647
                                                    _ = v31
                                                case "hello":
                                                    if true {
                                                        var v32 int16 = 1
                                                        switch fi13 {
                                                        case -238:
                                                            v33 := GenS2{F0: uint8(v16)}
                                                            _ = v33
                                                        case 32767:
                                                            v6 += 2
                                                        case -2147483648:
                                                            v34 := GenS6{F0: 472, F1: v6, F2: 585}
                                                            _ = v34
                                                        }
                                                        _ = v32
                                                    } else {
                                                        fmt.Println(fi13)
                                                    }
                                                case "\n\t":
                                                    var v35 string = "\n\t"
                                                    _ = v35
                                                default:
                                                    var v36 uint8 = (v19 + (^(v19)))
                                                    _ = v36
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = 672
                                                }
                                            case -299:
                                                v37 := GenS7{F0: -173, F1: v19, F2: v16}
                                                _ = v37
                                            case -166:
                                                v19 -= 0
                                            default:
                                                v19 = 1
                                            }
                                        }
                                        _ = v19
                                        _ = v22
                                    case "\n\t":
                                        if true {
                                            var v38 int16 = -998
                                            var v39 uint8 = 86
                                            _ = v38
                                            _ = v39
                                        } else {
                                            if false {
                                                var v40 int = (fi14 ^ (fi14 >> (((v16 | 4294967295)) & 63)))
                                                v6 = 255
                                                _ = v40
                                            }
                                        }
                                    case "a":
                                        v11 = (v11 - v11)
                                    }
                                    if false {
                                        v41 := GenS1{F0: uint8(v16), F1: 262, F2: (!false)}
                                        _ = v41
                                    } else {
                                        v42 := GenS6{F0: 42, F1: 552, F2: 2}
                                        _ = v42
                                    }
                                    v43 := GenS8{F0: 128, F1: 110, F2: 42}
                                    v11 = (3.14159 / v11)
                                    if true {
                                        if len(v3) > 0 {
                                            v3[0] = 4294967295
                                        }
                                    } else {
                                        fmt.Println(fi14, fi14, fi13)
                                        switch "a" {
                                        case "\n\t":
                                            var v44 uint32 = 2147483647
                                            _ = v44
                                        default:
                                            if (!true) {
                                                v11 = 0.1
                                                fmt.Println(v16)
                                                switch fi14 {
                                                case 0:
                                                    if len(v3) > 0 {
                                                        v3[0] = 42
                                                    }
                                                case 2147483647:
                                                    var v45 int16 = 0
                                                    _ = v45
                                                case 4294967295:
                                                    v6 -= (v6 ^ v6)
                                                }
                                            }
                                        }
                                        v6 = v6
                                    }
                                    fmt.Println(v11, v11, v6)
                                    v11 = v11
                                    _ = v43
                                }
                                _ = v17
                                _ = v18
                            }
                        case "":
                            v46 := GenS9{F0: 920}
                            _ = v46
                        case "foo bar":
                            var v47 string = "foo bar"
                            _ = v47
                        default:
                            fmt.Println(v6)
                        }
                        if false {
                            if true {
                                v48 := GenS10{F0: 42, F1: (^(v16))}
                                v11 = (((v11 + v11) - 3.14159) * (v11 + 0.1))
                                _ = v48
                            }
                            v4[-1] = "hello"
                            if false {
                                v11 += v11
                            }
                        } else {
                            v16 = 1
                            v6 = (v6 % ((v6) | 1))
                            var v49 int = (^(fi14))
                            _ = v49
                        }
                    }
                    _ = v16
                }
                v11 -= (v11 + float32(0.0))
            }
            var v50 int32 = -2147483648
            if true {
                switch 32768 {
                case 127:
                    v51 := GenS3{F0: 4294967295, F1: 42}
                    _ = v51
                case -949:
                    v52 := GenS11{F0: ((("a" + "\n\t") != "a") && false), F1: 45, F2: v6}
                    _ = v52
                case 1:
                    var v53 float32 = (((float32(0.0) + 1.5) + 0.0) / 1.5)
                    _ = v53
                default:
                    v54 := GenS4{F0: 0, F1: 42}
                    _ = v54
                }
            }
            _ = fi14
            _ = v50
        }
        v55 := GenS7{F0: fi13, F1: 0, F2: 65535}
        var v56 uint32 = 65535
        _ = fi13
        _ = v55
        _ = v56
    }
    switch "hello" {
    case "foo bar":
        if false {
            _ = len(v3)
            for fi57 := 0; fi57 < 3; fi57++ {
                var v58 float64 = ((0.0 * (0.0 - 3.14159)) / float64(v11))
                if true {
                    var v59 float64 = v58
                    var v60 uint64 = 941
                    fmt.Println(v6)
                    _ = v59
                    _ = v60
                } else {
                    if true {
                        for fi61 := 0; fi61 < 3; fi61++ {
                            v58 += (float64((-2.25 + v11)) + (v58 / float64(v11)))
                            v62 := GenS12{F0: 42, F1: -1}
                            _ = fi61
                            _ = v62
                        }
                        for fi63 := 0; fi63 < 18; fi63++ {
                            v6 -= v6
                            v64 := v4[fi63]
                            _ = fi63
                            _ = v64
                        }
                    }
                    v58 = -2.25
                    var v65 int64 = 1
                    _ = v65
                }
                _ = fi57
                _ = v58
            }
        } else {
            v66 := GenS9{F0: 891}
            _ = v66
        }
    case "hello":
        v1["a"] = 32768
    default:
        v6 -= v6
    }
    v11 -= v11
    switch -1 {
    case -772:
        v11 = v11
    case 128:
        var v67 int8 = -1
        _ = v67
    case 2147483647:
        for fi68 := 0; fi68 < 11; fi68++ {
            var v69 int = fi68
            _ = fi68
            _ = v69
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v6 = 640
    fmt.Println(v6)
    var v70 float64 = ((1.5 / ((0.1 - (float64(v11) * 3.14159)) - (1.5 / 1.0))) - 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v70
}
