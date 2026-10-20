package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 bool }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
type GenS4 struct { F0 uint; F1 int16; F2 uint }
type GenS5 struct { F0 int64; F1 string; F2 uint64 }
type GenS6 struct { F0 int64; F1 uint32; F2 uint64 }
func genFn0(p0 bool) uint16 {
    var v7 uint32 = 42
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: -1}
        fmt.Println(p0, v7, fi8)
        _ = fi8
        _ = v9
    }
    var v10 uint16 = 1
    _ = p0
    _ = v10
    _ = v7
    return 2
}
func genFn1(p0 uint64, p1 int16) uint {
    if false {
        var v11 int64 = 1
        if (42 < 100) {
            v12 := []int{128, 1, -842}
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            v13 := GenS2{F0: (p0 << ((65535) & 63)), F1: 845, F2: (true || false)}
            _ = v12
            _ = v13
        } else {
            v14 := []int{-9223372036854775808, 32768}
            p1 = (-(p1))
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 uint8 = 0
                if true {
                    p1 = -73
                    v17 := GenS2{F0: (p0 * p0), F1: 1, F2: (((0.1 * (0.1 * (-2.25 * (3.14159 + (-2.25 * float64(float32(float64((1.5 - 0.1))))))))) * 0.0) < 3.14159)}
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        _ = v11
    } else {
        p0 += p0
        var v18 float64 = (float64(float32(((3.14159 - (0.1 + 0.0)) + ((1e10 * 0.0) / 3.14159)))) + (1e10 * 1.5))
        _ = v18
    }
    p1 = 32767
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 float64 = (-2.25 * 1.5)
    var v3 uint64 = 646
    var v4 uint64 = v3
    var v5 uint = 194
    var v6 int16 = -121
    v5 -= v5
    if false {
        var v19 uint32 = 223
        fmt.Println(v19)
        _ = v19
    }
    v5 += (v5 + genFn1(v3, (-(v6))))
    if true {
        v20 := GenS1{F0: -1}
        var v21 uint16 = 963
        v5 = v5
        _ = v20
        _ = v21
    }
    if false {
        var v22 string = "foo bar"
        switch ("\n\t" + v22) {
        case "foo bar":
            if false {
                var v23 uint16 = 147
                if ((1.5 >= v2) && false) {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            var v26 string = v22
                            _ = fi25
                            _ = v26
                        }
                        var v27 float64 = v2
                        _ = fi24
                        _ = v27
                    }
                    switch 127 {
                    case 1:
                        v5 = (v5 << (((v5 % ((42) | 1))) & 63))
                    case 180:
                        switch -9223372036854775808 {
                        case -9223372036854775808:
                            for fi28 := 0; fi28 < 18; fi28++ {
                                for fi29 := 0; fi29 < 15; fi29++ {
                                    var v30 float64 = (v2 + 0.1)
                                    if false {
                                        var v31 uint64 = v3
                                        var v32 uint = v5
                                        _ = v31
                                        _ = v32
                                    } else {
                                        v33 := GenS2{F0: (v4 ^ v3), F1: (v5 - 625), F2: (false || false)}
                                        v1 = append(v1, 4294967295)
                                        _ = v33
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                                v34 := GenS1{F0: 588}
                                _ = fi28
                                _ = v34
                            }
                        case 623:
                            var v35 int8 = 42
                            _ = v35
                        default:
                            fmt.Println(v23, v3, v4)
                        }
                    case -2147483648:
                        for fi36 := 0; fi36 < 17; fi36++ {
                            switch (("a" + v22) + v22) {
                            case "foo bar":
                                if ("a" <= (v22 + (v22 + ""))) {
                                    v5 = (^(v5))
                                }
                            }
                            if len(v1) > 0 {
                                v1[0] = 372
                            }
                            _ = fi36
                        }
                    default:
                        var v37 string = "\n\t"
                        _ = v37
                    }
                }
                _ = v23
            } else {
                var v38 int8 = 42
                var v39 uint8 = 0
                _ = v38
                _ = v39
            }
        case "hello":
            switch "a" {
            case "":
                fmt.Println(v22)
            case "a":
                fmt.Println(v22)
            case "foo bar":
                switch 4294967295 {
                case 127:
                    if len(v1) > 0 {
                        v1[0] = 82
                    }
                case -2147483648:
                    for fi40 := 0; fi40 < 9; fi40++ {
                        var v41 int8 = -13
                        switch 0 {
                        case 4294967295:
                            v6 = (v6 + int16(v41))
                        }
                        var v42 string = (v22 + v22)
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                case 2147483647:
                    _ = len(v1)
                default:
                    v43 := GenS3{F0: 922, F1: v5, F2: 9223372036854775807}
                    _ = v43
                }
            default:
                if ((v6 & (v6 + v6)) != v6) {
                    var v44 uint8 = 78
                    if (!(!(!(0 == -128)))) {
                        v5 = v5
                        var v45 bool = false
                        v44 = v44
                        _ = v45
                    } else {
                        fmt.Println(v44, v2, v2)
                        var v46 uint8 = v44
                        _ = v46
                    }
                    _ = v44
                }
            }
        case "\n\t":
            v47 := GenS1{F0: 42}
            _ = v47
        default:
            var v48 uint = v5
            _ = v48
        }
        _ = v22
    } else {
        for fi49 := 0; fi49 < 3; fi49++ {
            v2 -= ((1e10 - float64((1e10 * 3.14159))) / v2)
            if (false || (32 == 1)) {
                switch 32768 {
                case 42:
                    if true {
                        if false {
                            var v50 uint = (^(v5))
                            v51 := GenS3{F0: 2, F1: v5, F2: -756}
                            var v52 float32 = float32(v2)
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                        for fi53 := 0; fi53 < 20; fi53++ {
                            v3 = 4294967295
                            fmt.Println(v3, v5)
                            v6 -= 162
                            if ((!(2147483647 == 295)) && false) {
                                var v54 uint64 = (v4 * (v3 / ((v4) | 1)))
                                switch "a" {
                                case "hello":
                                    var v55 int8 = 20
                                    _ = v55
                                case "foo bar":
                                    switch fi53 {
                                    case 42:
                                        v1 = append(v1, 42)
                                    default:
                                        v54 = v3
                                    }
                                    fmt.Println(v54)
                                    v56 := GenS4{F0: 1, F1: v6, F2: v5}
                                    _ = v56
                                case "a":
                                    switch -909 {
                                    case -621:
                                        v57 := GenS5{F0: 0, F1: "", F2: (v54 * 0)}
                                        _ = v57
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                default:
                                    v5 -= genFn1(v3, 174)
                                }
                                _ = v54
                            } else {
                                var v58 string = "foo bar"
                                v2 += (1.5 / v2)
                                var v59 bool = (v3 == (v3 * 614))
                                _ = v58
                                _ = v59
                            }
                            _ = fi53
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        v4 -= 18446744073709551615
                    }
                }
                if len(v1) > 0 {
                    v1[0] = 274
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v1 = append(v1, 836)
            _ = fi49
        }
        v6 += int16(v4)
    }
    fmt.Println(v2, v3)
    v60 := GenS6{F0: 1, F1: 925, F2: v4}
    if false {
        var v61 int8 = -128
        v61 += v61
        _ = v61
    } else {
        for fi62 := 0; fi62 < 12; fi62++ {
            var v63 uint32 = 65535
            v4 = (^(v3))
            for fi64 := 0; fi64 < 20; fi64++ {
                v4 = uint64(v63)
                v6 = (v6 | (v6 | (^(v6))))
                var v65 float64 = (float64(3.14159) - v2)
                _ = fi64
                _ = v65
            }
            _ = fi62
            _ = v63
        }
    }
    var v66 uint16 = 656
    if (!(42 < 0)) {
        var v67 float32 = (0.1 + float32(v2))
        if true {
            fmt.Println(v3, v2)
        } else {
            if len(v1) > 0 {
                v1[0] = 748
            }
        }
        _ = v67
    }
    var v68 uint8 = 42
    var v69 int8 = 127
    var v70 float32 = 1.5
    v68 -= v68
    v6 += v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v66
    _ = v68
    _ = v69
    _ = v70
}
