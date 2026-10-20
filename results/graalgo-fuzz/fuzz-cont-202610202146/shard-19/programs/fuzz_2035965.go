package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int8; F1 string; F2 int16 }
type GenS4 struct { F0 int8; F1 int32; F2 uint }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 string; F1 uint16; F2 bool }

func main() {
    var v1 uint64 = 645
    var v2 int8 = -1
    var v3 uint8 = 128
    var v4 int8 = (v2 * (v2 % ((42) | 1)))
    var v5 int = 4294967295
    v6 := make(map[uint32]bool)
    var v7 int8 = (v2 - v2)
    var v8 float64 = 1.5
    if true {
        v4 += -1
        v9 := GenS1{F0: v1}
        _ = v9
    } else {
        v10 := GenS2{F0: -1, F1: 423}
        var v11 int32 = 1
        _ = v10
        _ = v11
    }
    var v12 int8 = v4
    _ = len(v6)
    var v13 uint32 = 1
    v14 := GenS2{F0: 9223372036854775807, F1: 9223372036854775807}
    var v15 uint64 = 2
    var v16 uint16 = uint16(v1)
    if true {
        if (!false) {
            var v17 int64 = 427
            _ = v17
        } else {
            for fi18 := 0; fi18 < 13; fi18++ {
                if (0 > v1) {
                    v16 = (v16 * 0)
                    v15 = (^(v1))
                }
                for fi19 := 0; fi19 < 20; fi19++ {
                    switch fi19 {
                    case 128:
                        if ((v13 - (v13 + (^(v13)))) <= 0) {
                            v15 += 1
                            var v20 bool = (uint(v15) <= 18446744073709551615)
                            _ = v20
                        } else {
                            v2 -= (v7 | (v7 + (v2 * (v2 >> ((uint(v13)) & 63)))))
                        }
                    case 898:
                        var v21 uint = 761
                        _ = v21
                    case 9223372036854775807:
                        switch fi19 {
                        case -9223372036854775808:
                            switch "hello" {
                            case "\n\t":
                                var v22 int16 = 1000
                                _ = v22
                            }
                        default:
                            fmt.Println(v15, v7)
                        }
                    }
                    if false {
                        if (true || (!("\n\t" > ""))) {
                            v5 += fi18
                        } else {
                            v8 = v8
                            if true {
                                v23 := []int{}
                                var v24 string = ""
                                v13 += v13
                                var v25 uint32 = 2
                                v25 = v13
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                if ((((float64((3.14159 * (1e10 * (1.5 / (-2.25 * 1e10))))) - v8) / ((v8 + 3.14159) * (v8 / -2.25))) / v8) != ((0.0 * (3.14159 * v8)) * v8)) {
                                    if false {
                                        fmt.Println(v8)
                                        v6[uint32(v3)] = true
                                        if true {
                                            var v26 uint = 760
                                            _ = v26
                                        }
                                        switch "hello" {
                                        case "hello":
                                            v7 -= v7
                                        default:
                                            fmt.Println(v5, v8, v16)
                                        }
                                        var v27 int32 = 846
                                        _ = v27
                                    }
                                    v5 = 2147483647
                                    switch "hello" {
                                    case "hello":
                                        v28 := GenS2{F0: -351, F1: 0}
                                        _ = v28
                                    default:
                                        v8 += (v8 * 3.14159)
                                    }
                                } else {
                                    var v29 uint = 1
                                    v29 = v29
                                    _ = v29
                                }
                                switch fi19 {
                                case 0:
                                    var v30 float32 = (((0.0 / 3.14159) / 1.5) / ((1.5 - 0.1) / float32(v8)))
                                    _ = v30
                                case 127:
                                    v13 = uint32(v16)
                                case 32767:
                                    v31 := GenS2{F0: -9223372036854775808, F1: 1}
                                    _ = v31
                                default:
                                    var v32 float32 = 0.0
                                    _ = v32
                                }
                                if false {
                                    v33 := GenS1{F0: (v1 / (((v1 * (v15 + v15))) | 1))}
                                    v16 = (^(v16))
                                    _ = v33
                                } else {
                                    if (((((((v8 - v8) * 1e10) + 1e10) > 0.1) || true) && false) && true) {
                                        v7 += int8(v3)
                                        var v34 int32 = 42
                                        var v35 uint16 = (v16 ^ (v16 - v16))
                                        _ = v34
                                        _ = v35
                                    }
                                }
                            }
                            switch ("\n\t" + "a") {
                            case "a":
                                v36 := GenS3{F0: 100, F1: ("foo bar" + (("foo bar" + "a") + ("\n\t" + "\n\t"))), F2: 1000}
                                v37 := []int{(fi18 * fi18), -1}
                                _ = v36
                                _ = v37
                            case "hello":
                                var v38 int64 = -9223372036854775808
                                _ = v38
                            case "\n\t":
                                if (0 > 42) {
                                    var v39 float64 = (((v8 - ((float64(-2.25) / v8) * (1.5 - (-2.25 - v8)))) - (v8 + 3.14159)) - float64(float32((1.5 + (-2.25 * 0.0)))))
                                    var v40 int64 = -9223372036854775808
                                    v41 := GenS1{F0: v15}
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                } else {
                                    switch -404 {
                                    case 32767:
                                        v8 += v8
                                    default:
                                        v7 = 53
                                    }
                                }
                            }
                        }
                        if ("" >= ("\n\t" + "foo bar")) {
                            v4 = int8(v13)
                        } else {
                            var v42 float64 = v8
                            var v43 int8 = -1
                            _ = v42
                            _ = v43
                        }
                        var v44 uint = 42
                        if true {
                            v7 += (v7 - v7)
                            v13 = (^(v13))
                            var v45 uint16 = v16
                            _ = v45
                        }
                        v7 -= (v7 - -1)
                        _ = v44
                    } else {
                        if (65535 == 1) {
                            v46 := GenS4{F0: 1, F1: 745, F2: uint(v16)}
                            _ = v46
                        }
                    }
                    if (9223372036854775807 == -207) {
                        var v47 uint = 333
                        var v48 bool = (-32768 >= -1)
                        _ = v47
                        _ = v48
                    } else {
                        v49 := []int{}
                        v8 = float64(1e10)
                        _ = v49
                    }
                    fmt.Println(v8, v13, v13)
                    _ = fi19
                }
                for fi50 := 0; fi50 < 3; fi50++ {
                    v16 = v16
                    var v51 int8 = v7
                    _ = fi50
                    _ = v51
                }
                _ = fi18
            }
            v52 := GenS5{F0: (v1 + 42), F1: (v1 & 42)}
            v53 := GenS6{F0: -2147483648, F1: 32767}
            _ = v52
            _ = v53
        }
        _ = len(v6)
    } else {
        switch v5 {
        case 2147483647:
            for fi54 := 0; fi54 < 6; fi54++ {
                var v55 uint8 = (^(v3))
                v4 = (v12 & (v4 >> ((1) & 63)))
                v3 = (^(v3))
                _ = fi54
                _ = v55
            }
        case 128:
            if ((v16 <= v16) && true) {
                switch "\n\t" {
                case "":
                    fmt.Println(v2, v12)
                case "a":
                    var v56 string = (("a" + "foo bar") + "foo bar")
                    _ = v56
                default:
                    switch "" {
                    case "\n\t":
                        var v57 int8 = v7
                        _ = v57
                    case "foo bar":
                        v58 := v6[uint32(v16)]
                        _ = v58
                    case "hello":
                        if false {
                            if true {
                                v2 -= (v4 << ((18446744073709551615) & 63))
                                if true {
                                    switch v5 {
                                    case 32767:
                                        var v59 int8 = (v2 + 120)
                                        _ = v59
                                    case -51:
                                        delete(v6, (v13 + (v13 % (((v13 ^ v13)) | 1))))
                                    }
                                    v60 := []int{v5, v5, (v5 << ((65535) & 63))}
                                    _ = v60
                                } else {
                                    for fi61 := 0; fi61 < 16; fi61++ {
                                        v16 += v16
                                        v62 := GenS1{F0: v15}
                                        _ = fi61
                                        _ = v62
                                    }
                                    v3 += 0
                                }
                            }
                        }
                    default:
                        var v63 uint = 255
                        _ = v63
                    }
                }
            } else {
                v64 := GenS7{F0: ("a" + (("foo bar" + "a") + ("\n\t" + "a"))), F1: v16, F2: (!true)}
                var v65 string = (("foo bar" + (("a" + ("\n\t" + "")) + ("\n\t" + "foo bar"))) + "\n\t")
                _ = v64
                _ = v65
            }
        case 803:
            fmt.Println(v3, v4)
        }
        v4 += 42
        v6[65535] = false
    }
    v66 := []int{(-(v5)), v5, (v5 / ((4294967295) | 1))}
    var v67 uint8 = v3
    for fi68 := 0; fi68 < 8; fi68++ {
        v1 += (v15 | v1)
        _ = fi68
    }
    v69 := GenS3{F0: v2, F1: "", F2: 1}
    v70 := v6[(v13 & (v13 >> ((18446744073709551615) & 63)))]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v69
    _ = v7
    _ = v70
    _ = v8
}
