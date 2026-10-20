package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int8; F2 int32 }
type GenS3 struct { F0 uint; F1 int16; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int32; F1 int; F2 uint32 }

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]bool)
    var v3 uint32 = 172
    v4 := []uint32{(v3 | 2147483647), v3, (v3 >> ((4294967295) & 63))}
    v5 := []int{424, 0, 128}
    var v6 int16 = -606
    v6 += (v6 >> ((uint(v1)) & 63))
    var v7 int32 = -578
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(fi8)
        fmt.Println(v6)
        v9 := GenS1{F0: (-(v6))}
        _ = fi8
        _ = v9
    }
    var v10 int8 = 42
    v7 -= v7
    if (!false) {
        switch "hello" {
        case "foo bar":
            v6 += (^(v6))
        default:
            fmt.Println(v7, v3, v3)
        }
        fmt.Println(v1, v7)
    } else {
        if true {
            var v11 int32 = int32(v1)
            v12 := GenS1{F0: v6}
            _ = v11
            _ = v12
        }
        var v13 int32 = (v7 & (v7 ^ v7))
        v4 = append(v4, 499)
        _ = v13
    }
    var v14 float64 = 1.5
    if false {
        fmt.Println(v10)
    }
    var v15 uint16 = 424
    var v16 int64 = int64(v7)
    v7 += v7
    v17 := GenS1{F0: (^(v6))}
    if true {
        v18 := GenS1{F0: (v6 ^ v6)}
        var v19 int16 = v6
        v1 = v1
        _ = v18
        _ = v19
    } else {
        v2[v3] = false
        v20 := GenS2{F0: false, F1: v10, F2: (-(v7))}
        _ = v20
    }
    v1 += (v1 ^ (^(v1)))
    fmt.Println(v14, v15, v10)
    switch "a" {
    case "hello":
        switch 127 {
        case 32768:
            v21 := GenS1{F0: (v6 % ((v6) | 1))}
            _ = v21
        case -532:
            v22 := GenS2{F0: true, F1: (v10 / (((v10 * (v10 * (-(v10))))) | 1)), F2: (v7 | v7)}
            _ = v22
        }
    case "\n\t":
        switch (("\n\t" + ("\n\t" + "foo bar")) + ((((("\n\t" + ("a" + ("" + ""))) + "a") + ((((("a" + ("a" + "\n\t")) + "") + "foo bar") + "") + ("foo bar" + ""))) + "hello") + (("a" + ("\n\t" + ((("hello" + "a") + "\n\t") + (("foo bar" + "\n\t") + ("foo bar" + "\n\t"))))) + ""))) {
        case "":
            if ((^(v7)) != v7) {
                v1 = 277
                var v23 uint32 = (v3 >> ((1) & 63))
                var v24 float32 = (3.14159 / 1.0)
                _ = v23
                _ = v24
            }
        case "a":
            for fi25 := 0; fi25 < 12; fi25++ {
                for fi26 := 0; fi26 < 14; fi26++ {
                    var v27 float64 = v14
                    switch "" {
                    case "":
                        v15 -= 571
                    case "a":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        var v28 float32 = 0.0
                        _ = v28
                    }
                    fmt.Println(v10, v6, v14)
                    _ = fi26
                    _ = v27
                }
                if false {
                    for fi29 := 0; fi29 < 16; fi29++ {
                        switch "" {
                        case "hello":
                            v30 := v2[v3]
                            _ = v30
                        default:
                            v31 := v2[v3]
                            _ = v31
                        }
                        switch ("foo bar" + "a") {
                        case "":
                            var v32 uint16 = (v15 + (v15 - v15))
                            _ = v32
                        case "foo bar":
                            var v33 uint8 = 128
                            _ = v33
                        case "\n\t":
                            v14 -= v14
                        default:
                            switch (^(fi29)) {
                            case 127:
                                if true {
                                    switch fi29 {
                                    case 42:
                                        v34 := GenS3{F0: 332, F1: v6, F2: (fi25 & (fi29 * fi25))}
                                        _ = v34
                                    case 708:
                                        v35 := GenS3{F0: 132, F1: (v6 % ((v6) | 1)), F2: fi25}
                                        _ = v35
                                    default:
                                        fmt.Println(v16, v7, fi25)
                                    }
                                    fmt.Println(v10, v14, fi29)
                                } else {
                                    fmt.Println(v16, fi29)
                                }
                            case 2147483647:
                                if (-198 < (fi25 & fi29)) {
                                    switch (fi25 ^ (fi29 * (fi25 - -1))) {
                                    case 4294967295:
                                        v2[(^(v3))] = false
                                    }
                                    fmt.Println(v6, v1)
                                    if false {
                                        var v36 uint8 = 0
                                        if ((v1 % ((214) | 1)) == v1) {
                                            if (v36 < v36) {
                                                if (0 == fi25) {
                                                    var v37 int32 = (-(v7))
                                                    _ = v37
                                                } else {
                                                    if (((v36 ^ 0) >= (v36 - (v36 * v36))) || false) {
                                                        var v38 float32 = 1.5
                                                        _ = v38
                                                    }
                                                    v16 = (v16 / ((v16) | 1))
                                                    v39 := GenS4{F0: (v10 | 100)}
                                                    _ = v39
                                                }
                                                v6 -= (^(v6))
                                            }
                                            var v40 uint = 0
                                            _ = v40
                                        }
                                        _ = v36
                                    }
                                }
                            case 9223372036854775807:
                                v6 -= v6
                            default:
                                var v41 int16 = v6
                                _ = v41
                            }
                        }
                        _ = fi29
                    }
                    var v42 uint = 0
                    if (v15 >= (v15 | (v15 / ((v15) | 1)))) {
                        if (224 > 38) {
                            _ = len(v4)
                        } else {
                            v43 := GenS4{F0: 0}
                            for fi44 := 0; fi44 < 14; fi44++ {
                                v3 = 1
                                if len(v4) > 0 {
                                    v4[0] = v3
                                }
                                v15 += (v15 * (^(v15)))
                                _ = fi44
                            }
                            _ = v43
                        }
                        fmt.Println(fi25, v7, v16)
                    }
                    _ = v42
                } else {
                    fmt.Println(v1, v3)
                }
                _ = fi25
            }
        default:
            if true {
                var v45 int64 = -9223372036854775808
                fmt.Println(v1)
                switch "foo bar" {
                case "foo bar":
                    v3 -= 807
                default:
                    _ = len(v2)
                }
                _ = v45
            } else {
                v46 := GenS5{F0: false, F1: (^(v16))}
                switch "hello" {
                case "\n\t":
                    var v47 int = 9223372036854775807
                    _ = v47
                case "a":
                    var v48 uint16 = 345
                    _ = v48
                default:
                    if ((^(v3)) > 209) {
                        v14 = v14
                    } else {
                        v6 -= 0
                    }
                }
                _ = v46
            }
        }
    default:
        if true {
            fmt.Println(v3, v10)
            if (42 > 991) {
                for fi49 := 0; fi49 < 16; fi49++ {
                    var v50 int64 = 9223372036854775807
                    _ = fi49
                    _ = v50
                }
                v7 = v7
                if len(v4) > 0 {
                    v4[0] = v3
                }
            } else {
                for fi51 := 0; fi51 < 16; fi51++ {
                    if true {
                        var v52 float64 = v14
                        if (true && ((v1 == 18446744073709551615) || true)) {
                            v53 := v2[(v3 | 1)]
                            v54 := GenS6{F0: (v7 << ((973) & 63)), F1: 32767, F2: 4294967295}
                            _ = v53
                            _ = v54
                        } else {
                            v55 := GenS6{F0: (v7 * v7), F1: (fi51 % (((fi51 / ((2147483647) | 1))) | 1)), F2: v3}
                            _ = v55
                        }
                        _ = v52
                    } else {
                        if ((!true) && (false || ((v10 / ((127) | 1)) != (v10 & (v10 % (((v10 >> ((948) & 63))) | 1)))))) {
                            switch ("a" + "foo bar") {
                            case "":
                                v7 = 672
                            case "a":
                                var v56 float32 = float32(((v14 / v14) - 3.14159))
                                _ = v56
                            }
                            for fi57 := 0; fi57 < 8; fi57++ {
                                var v58 int8 = v10
                                var v59 int = (^(fi57))
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                            v15 += v15
                        } else {
                            var v60 uint16 = 65535
                            for fi61 := 0; fi61 < 16; fi61++ {
                                v60 -= uint16(v3)
                                v2[2] = (!((^(fi61)) != fi61))
                                _ = fi61
                            }
                            if (!(false || true)) {
                                var v62 string = "a"
                                var v63 int32 = (v7 + v7)
                                _ = v62
                                _ = v63
                            } else {
                                v64 := v2[(v3 & (v3 ^ (v3 | (v3 << ((uint(v15)) & 63)))))]
                                _ = v64
                            }
                            _ = v60
                        }
                    }
                    _ = fi51
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
