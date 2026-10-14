package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint8 }
type GenS3 struct { F0 bool; F1 int8; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 bool; F1 int64; F2 int8 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 string; F1 int16; F2 int64 }
type GenS8 struct { F0 int8 }
func genFn0() string {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int = fi6
        fmt.Println(v7, fi6)
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 0
    if true {
        for fi9 := 0; fi9 < 4; fi9++ {
            switch fi9 {
            case 1:
                if true {
                    var v10 string = ""
                    _ = v10
                }
            case 32768:
                var v11 int = fi9
                _ = v11
            default:
                var v12 int16 = -986
                _ = v12
            }
            var v13 uint = 978
            for fi14 := 0; fi14 < 11; fi14++ {
                v13 += v13
                switch fi9 {
                case -1:
                    v15 := GenS1{F0: 0, F1: uint32(v13)}
                    _ = v15
                case -224:
                    v8 += v8
                case 343:
                    v16 := GenS1{F0: -575, F1: 1}
                    _ = v16
                }
                switch "a" {
                case "hello":
                    if ((100 > -126) || (false || false)) {
                        v13 += (v13 * (v13 ^ v13))
                        switch fi14 {
                        case 2147483647:
                            if false {
                                v13 += v13
                            } else {
                                switch 2147483647 {
                                case -2147483648:
                                    var v17 float32 = 1.5
                                    var v18 bool = (!(!false))
                                    v8 = (v8 / (((v8 + v8)) | 1))
                                    _ = v17
                                    _ = v18
                                default:
                                    v13 = 18446744073709551615
                                }
                                var v19 float32 = float32(-2.25)
                                if (true || true) {
                                    fmt.Println(fi9)
                                    var v20 bool = (42 >= 1)
                                    _ = v20
                                } else {
                                    var v21 bool = (!false)
                                    v22 := []int{}
                                    _ = v21
                                    _ = v22
                                }
                                v13 = (v13 << (((^(v13))) & 63))
                                if false {
                                    switch fi9 {
                                    case 128:
                                        v23 := GenS2{F0: 42, F1: 398, F2: 208}
                                        _ = v23
                                    case -948:
                                        var v24 int64 = -9223372036854775808
                                        _ = v24
                                    default:
                                        var v25 string = "a"
                                        _ = v25
                                    }
                                    var v26 string = (("foo bar" + "a") + "a")
                                    switch "a" {
                                    case "\n\t":
                                        var v27 string = v26
                                        _ = v27
                                    case "":
                                        if false {
                                            if false {
                                                fmt.Println(v19)
                                                v28 := GenS3{F0: true, F1: int8(v13), F2: int16(fi14)}
                                                if (false && ((v13 & v13) >= 18446744073709551615)) {
                                                    if true {
                                                        v29 := GenS4{F0: 65535, F1: v8}
                                                        _ = v29
                                                    } else {
                                                        var v30 int64 = -9223372036854775808
                                                        v31 := []int{fi9, fi14, (fi14 * (fi9 * 2147483647)), fi9}
                                                        v8 += v8
                                                        fmt.Println(v19, fi9)
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                }
                                                _ = v28
                                            }
                                        }
                                    case "a":
                                        var v32 int32 = 754
                                        _ = v32
                                    }
                                    v33 := []int{(fi14 << (((v13 + (v13 % (((v13 - v13)) | 1)))) & 63)), (^(fi9))}
                                    _ = v26
                                    _ = v33
                                }
                                _ = v19
                            }
                            v13 = (v13 + (^(v13)))
                            var v34 bool = ((v8 + (^(v8))) >= 128)
                            _ = v34
                        default:
                            v35 := GenS4{F0: 0, F1: v8}
                            _ = v35
                        }
                        var v36 string = "hello"
                        var v37 float32 = 0.0
                        _ = v36
                        _ = v37
                    } else {
                        v38 := []int{-1, fi9, fi14, 128}
                        _ = v38
                    }
                default:
                    fmt.Println(fi9, fi14)
                }
                v39 := GenS3{F0: (!(false || true)), F1: 42, F2: -32768}
                var v40 uint16 = 2
                _ = fi14
                _ = v39
                _ = v40
            }
            _ = fi9
            _ = v13
        }
    } else {
        v41 := GenS3{F0: (1 >= -2147483648), F1: -1, F2: -219}
        var v42 string = ("\n\t" + "hello")
        var v43 uint64 = 2
        _ = v41
        _ = v42
        _ = v43
    }
    v8 = (^(v8))
    _ = v8
    return ("hello" + ("" + ("\n\t" + "")))
}
func genFn1(p0 uint16, p1 int16) int64 {
    var v44 int64 = -230
    v44 -= v44
    _ = p0
    _ = p1
    _ = v44
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    v3 := []uint32{0, 960, 2147483647}
    var v4 uint8 = 42
    v5 := make(map[string]int)
    var v45 float32 = -2.25
    if ((1.5 != (0.0 - 0.0)) || (42 > 1)) {
        switch 858 {
        case -9223372036854775808:
            v46 := GenS5{F0: true, F1: -130, F2: 71}
            _ = v46
        }
        switch "a" {
        case "":
            v45 = 1e10
        case "foo bar":
            v3 = append(v3, 149)
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            switch 32767 {
            case 686:
                switch 9223372036854775807 {
                case 0:
                    switch (("\n\t" + "hello") + (("hello" + "") + "\n\t")) {
                    case "foo bar":
                        v4 = v4
                    case "":
                        v45 += 0.1
                    default:
                        var v47 uint8 = 42
                        _ = v47
                    }
                case 9223372036854775807:
                    if true {
                        switch 32768 {
                        case 9223372036854775807:
                            v48 := GenS4{F0: 0, F1: (v4 * v4)}
                            _ = v48
                        case 127:
                            if (!false) {
                                if true {
                                    v45 -= (0.0 + v45)
                                    var v49 int64 = 1
                                    fmt.Println(v49)
                                    _ = v49
                                }
                                v45 = (v45 * v45)
                            }
                        case 0:
                            v45 = (3.14159 / -2.25)
                        default:
                            v45 += v45
                        }
                        fmt.Println(v4, v45)
                        v50 := GenS6{F0: 100, F1: 1}
                        _ = v50
                    }
                case -776:
                    switch -742 {
                    case -1:
                        var v51 uint32 = 268
                        _ = v51
                    default:
                        var v52 uint8 = v4
                        _ = v52
                    }
                }
            case -775:
                fmt.Println(v45, v45)
            case 32768:
                if true {
                    v2[-9223372036854775808] = 127
                    for fi53 := 0; fi53 < 16; fi53++ {
                        var v54 string = "foo bar"
                        v55 := GenS2{F0: 42, F1: 46, F2: (^(v4))}
                        var v56 int8 = 0
                        _ = fi53
                        _ = v54
                        _ = v55
                        _ = v56
                    }
                } else {
                    if (1.5 == float64(v45)) {
                        var v57 int16 = 381
                        _ = v57
                    } else {
                        v45 = ((v45 / v45) + v45)
                    }
                }
            default:
                var v58 uint32 = 664
                _ = v58
            }
        }
    } else {
        var v59 int16 = -266
        v60 := v2[127]
        _ = v59
        _ = v60
    }
    fmt.Println(v4, v45)
    for fi61 := 0; fi61 < 17; fi61++ {
        v4 = v4
        v4 += v4
        if (false || ((3.14159 + 1e10) == -2.25)) {
            fmt.Println(fi61)
        } else {
            v45 = ((0.1 - 0.0) * 1.5)
            for fi62 := 0; fi62 < 2; fi62++ {
                if true {
                    var v63 int16 = -465
                    _ = len(v3)
                    v63 += 46
                    v3 = append(v3, 814)
                    _ = v63
                }
                var v64 uint32 = 278
                var v65 int64 = 644
                _ = fi62
                _ = v64
                _ = v65
            }
        }
        _ = fi61
    }
    fmt.Println(v45, v4)
    switch "a" {
    case "\n\t":
        v45 = (v45 * v45)
    case "hello":
        for fi66 := 0; fi66 < 13; fi66++ {
            if (false && ((!(18446744073709551615 != 42)) && false)) {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                v45 += (((v45 - 3.14159) / v45) - 1.5)
                var v67 int8 = -125
                _ = v67
            } else {
                var v68 uint = 395
                var v69 bool = true
                _ = v68
                _ = v69
            }
            var v70 int64 = -1
            _ = fi66
            _ = v70
        }
    case "":
        var v71 uint32 = 42
        _ = v71
    default:
        v4 = v4
    }
    fmt.Println(v4, v4)
    if false {
        v72 := GenS7{F0: ((((("" + "hello") + "foo bar") + "\n\t") + "foo bar") + "foo bar"), F1: 32767, F2: 0}
        _ = v72
    }
    switch "" {
    case "hello":
        for fi73 := 0; fi73 < 5; fi73++ {
            v74 := GenS1{F0: 1, F1: 65535}
            _ = fi73
            _ = v74
        }
    case "":
        v75 := GenS8{F0: 42}
        _ = v75
    case "\n\t":
        var v76 int16 = 1000
        _ = v76
    default:
        v45 = v45
    }
    delete(v5, "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
