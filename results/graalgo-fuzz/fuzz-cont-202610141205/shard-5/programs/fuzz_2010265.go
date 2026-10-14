package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS7 struct { F0 int8; F1 int8; F2 string }
func genFn0() bool {
    v5 := GenS1{F0: -405, F1: 127, F2: (!true)}
    v6 := []int{644, -2147483648, 0, 42}
    var v7 uint = 910
    for fi8 := 0; fi8 < 6; fi8++ {
        v6 = append(v6, -581)
        v9 := GenS1{F0: 9223372036854775807, F1: 127, F2: false}
        _ = fi8
        _ = v9
    }
    fmt.Println(v7)
    _ = v5
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 int16 = 42
    var v2 int32 = -2147483648
    var v3 int8 = -128
    v4 := make(map[int]uint32)
    fmt.Println(v3, v3)
    v10 := []int{-266, 4294967295, -392, 42}
    v1 += 549
    for fi11 := 0; fi11 < 7; fi11++ {
        if (true && false) {
            if (true && (250 < 1)) {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v13 := GenS1{F0: -1, F1: 1, F2: (v2 == v2)}
                    _ = fi12
                    _ = v13
                }
            } else {
                switch ("\n\t" + "") {
                case "foo bar":
                    v1 += (v1 | v1)
                }
                if (179 != 0) {
                    if false {
                        if (881 == 387) {
                            v2 = (v2 ^ (^(v2)))
                            v3 += v3
                            v14 := GenS1{F0: int64(v3), F1: 3, F2: (true || false)}
                            _ = v14
                        }
                        v1 -= 0
                    } else {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        v2 += v2
                    }
                    if true {
                        v15 := GenS1{F0: 0, F1: 128, F2: (42 < 4294967295)}
                        v16 := GenS1{F0: 198, F1: 255, F2: ((float64(((0.0 - (0.1 - 1e10)) * -2.25)) * 1e10) != 3.14159)}
                        _ = v15
                        _ = v16
                    }
                    if (!true) {
                        for fi17 := 0; fi17 < 6; fi17++ {
                            var v18 uint8 = 48
                            _ = fi17
                            _ = v18
                        }
                        var v19 uint32 = 1
                        if true {
                            for fi20 := 0; fi20 < 13; fi20++ {
                                v21 := GenS2{F0: int32(v1), F1: 0}
                                v22 := v4[907]
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        }
                        _ = v19
                    }
                }
                if (false && (1 >= 549)) {
                    if false {
                        v2 -= int32(fi11)
                    } else {
                        var v23 uint32 = 913
                        v1 += v1
                        _ = v23
                    }
                } else {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        if (true || (false || (215 >= 1))) {
                            v25 := v4[fi24]
                            _ = v25
                        } else {
                            v26 := GenS2{F0: (v2 | 42), F1: 0}
                            switch "foo bar" {
                            case "hello":
                                delete(v4, (fi24 + (fi24 & fi11)))
                            default:
                                v3 -= (v3 ^ -1)
                                v1 = v1
                                var v27 float64 = ((0.0 + -2.25) / 1.0)
                                _ = len(v4)
                                _ = v27
                            }
                            _ = v26
                        }
                        var v28 uint64 = 4294967295
                        switch "\n\t" {
                        case "foo bar":
                            v29 := GenS2{F0: (v2 >> ((2) & 63)), F1: v28}
                            _ = v29
                        case "\n\t":
                            v28 += v28
                        case "":
                            var v30 bool = false
                            _ = v30
                        default:
                            if (789 < 1) {
                                switch "foo bar" {
                                case "hello":
                                    var v31 int16 = (v1 * (v1 % (((^(v1))) | 1)))
                                    _ = v31
                                case "":
                                    v1 = int16(fi24)
                                case "a":
                                    if true {
                                        v3 = v3
                                        var v32 int16 = (-(v1))
                                        _ = v32
                                    } else {
                                        var v33 int32 = (v2 | v2)
                                        v34 := GenS3{F0: fi11, F1: (v3 / ((v3) | 1))}
                                        _ = v33
                                        _ = v34
                                    }
                                default:
                                    var v35 float64 = 1e10
                                    var v36 float32 = 0.1
                                    v37 := GenS1{F0: 42, F1: 119, F2: (114 < 42)}
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                            } else {
                                switch -744 {
                                case -743:
                                    v38 := GenS4{F0: 65, F1: 65535, F2: ((("a" + "\n\t") + "foo bar") + ("" + ""))}
                                    _ = v38
                                case 281:
                                    v28 = v28
                                case 128:
                                    v39 := GenS4{F0: 127, F1: 839, F2: "\n\t"}
                                    _ = v39
                                default:
                                    var v40 int = (fi11 & (fi24 / (((fi11 + 32768)) | 1)))
                                    _ = v40
                                }
                            }
                        }
                        _ = fi24
                        _ = v28
                    }
                    v2 -= (-(v2))
                    for fi41 := 0; fi41 < 12; fi41++ {
                        v42 := GenS5{F0: (v3 / ((v3) | 1))}
                        v1 -= (v1 & 32767)
                        if true {
                            fmt.Println(fi11, v3, v1)
                        } else {
                            var v43 int16 = int16(fi41)
                            _ = v43
                        }
                        _ = fi41
                        _ = v42
                    }
                }
            }
            for fi44 := 0; fi44 < 9; fi44++ {
                if ("foo bar" != "hello") {
                    v10 = append(v10, 42)
                    _ = len(v4)
                    v45 := GenS4{F0: 24, F1: 4294967295, F2: "hello"}
                    v46 := GenS6{F0: 0, F1: 0, F2: (^(v3))}
                    if (false && false) {
                        v47 := GenS1{F0: 42, F1: 42, F2: ((v1 - (v1 - (v1 % ((v1) | 1)))) == 1000)}
                        v3 -= v3
                        _ = v47
                    } else {
                        v48 := GenS5{F0: 0}
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        v1 = 1
                        _ = v48
                    }
                    _ = v45
                    _ = v46
                } else {
                    v49 := GenS6{F0: 1, F1: 128, F2: v3}
                    var v50 float32 = ((0.1 + ((0.1 - 3.14159) / 0.1)) / (0.0 - 1e10))
                    _ = v49
                    _ = v50
                }
                _ = fi44
            }
        } else {
            v3 -= v3
            if false {
                for fi51 := 0; fi51 < 6; fi51++ {
                    v1 += (v1 * (v1 & v1))
                    if false {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        var v52 int64 = -1
                        if true {
                            switch (^(fi51)) {
                            case 1:
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            case -847:
                                v1 = (^(v1))
                            case 128:
                                var v53 int = (fi51 << ((2) & 63))
                                _ = v53
                            default:
                                v4[fi51] = 2
                            }
                            fmt.Println(fi51, v2, v3)
                            v54 := GenS7{F0: 0, F1: (v3 - (v3 ^ 97)), F2: "foo bar"}
                            _ = v54
                        }
                        _ = v52
                    }
                    var v55 float64 = 1e10
                    _ = fi51
                    _ = v55
                }
            }
            for fi56 := 0; fi56 < 7; fi56++ {
                v57 := GenS5{F0: 100}
                v2 += (v2 << ((2) & 63))
                _ = fi56
                _ = v57
            }
        }
        if true {
            fmt.Println(v2, v2, v2)
        }
        _ = fi11
    }
    switch -308 {
    case -9223372036854775808:
        v58 := GenS1{F0: 0, F1: 0, F2: true}
        _ = v58
    case 4294967295:
        v1 -= -240
    }
    switch 974 {
    case 1:
        var v59 int64 = 1
        _ = v59
    case 0:
        var v60 string = "foo bar"
        _ = v60
    case 386:
        v2 -= v2
    default:
        var v61 float32 = 3.14159
        _ = v61
    }
    v62 := GenS6{F0: 127, F1: 189, F2: -1}
    for fi63 := 0; fi63 < 4; fi63++ {
        var v64 uint = 132
        _ = fi63
        _ = v64
    }
    if (!true) {
        v4[-660] = 65535
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v62
}
