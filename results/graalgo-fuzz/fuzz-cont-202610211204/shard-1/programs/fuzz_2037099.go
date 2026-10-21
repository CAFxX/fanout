package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int8; F1 int32 }

func main() {
    var v1 uint = 65535
    var v2 uint64 = 42
    v3 := []string{"foo bar", (("hello" + (("" + "hello") + "hello")) + ("foo bar" + ("hello" + "foo bar"))), ("foo bar" + "")}
    v4 := make(map[uint32]bool)
    var v5 string = ""
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            v2 = 0
            var v8 float32 = float32(0.0)
            _ = fi7
            _ = v8
        }
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        v9 := GenS1{F0: 0, F1: 744}
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        switch (fi10 % ((fi10) | 1)) {
        case -273:
            v2 = v2
        case 32767:
            switch fi10 {
            case 1:
                var v11 int = fi10
                _ = v11
            case -9223372036854775808:
                v12 := GenS1{F0: 42, F1: 0}
                _ = v12
            }
        case -33:
            v5 = (v5 + "a")
        }
        v13 := GenS2{F0: 89, F1: v2}
        if true {
            var v14 int = fi10
            v5 = v5
            switch v14 {
            case 1:
                for fi15 := 0; fi15 < 10; fi15++ {
                    var v16 int8 = -1
                    _ = fi15
                    _ = v16
                }
            }
            _ = v14
        }
        _ = fi10
        _ = v13
    }
    if ((0.1 / ((0.1 + float64(0.0)) / 1.0)) >= ((3.14159 * 0.1) + 1e10)) {
        v1 -= (v1 * (v1 * v1))
        if len(v3) > 0 {
            v3[0] = (v5 + ("hello" + ("a" + "a")))
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            v4[677] = (!(!false))
            var v18 bool = (-39 >= 100)
            var v19 int64 = 705
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        v5 = (v5 + "a")
        v20 := GenS2{F0: 828, F1: (v2 << ((v1) & 63))}
        var v21 int16 = 98
        _ = v20
        _ = v21
    }
    var v22 uint8 = 188
    switch v5 {
    case "foo bar":
        var v23 int = -9223372036854775808
        _ = v23
    default:
        delete(v4, 2147483647)
    }
    v2 -= (v2 * (v2 * v2))
    var v24 uint32 = 779
    v1 += v1
    v2 += uint64(v22)
    if (v1 != 0) {
        v25 := GenS3{F0: v1}
        fmt.Println(v5, v1, v22)
        _ = len(v4)
        _ = v25
    } else {
        v26 := GenS1{F0: 65535, F1: (v24 << ((v1) & 63))}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v26
    }
    v24 -= v24
    for fi27 := 0; fi27 < 6; fi27++ {
        if (false && false) {
            v28 := GenS4{F0: (^(v1)), F1: (-1 <= -218)}
            _ = v28
        } else {
            for fi29 := 0; fi29 < 19; fi29++ {
                v30 := v4[0]
                v31 := GenS5{F0: 42, F1: 9223372036854775807, F2: (^(v2))}
                v24 -= v24
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        var v32 uint16 = 381
        _ = fi27
        _ = v32
    }
    var v33 int = 128
    v1 -= v1
    if false {
        for fi34 := 0; fi34 < 19; fi34++ {
            if ((v22 ^ v22) > (v22 + (v22 + 127))) {
                for fi35 := 0; fi35 < 12; fi35++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi35
                }
                v4[(v24 ^ v24)] = (!true)
                v5 = v5
            } else {
                var v36 uint16 = 255
                _ = v36
            }
            _ = fi34
        }
        if true {
            v1 -= 647
            fmt.Println(v22)
        }
        for fi37 := 0; fi37 < 3; fi37++ {
            v24 -= (^(v24))
            v1 += (v1 ^ (v1 | v1))
            fmt.Println(v24, fi37)
            _ = fi37
        }
    } else {
        if ((v22 << ((v1) & 63)) < (v22 % ((144) | 1))) {
            var v38 int32 = 299
            var v39 float32 = ((1e10 + float32(1.5)) - (0.1 * 1.5))
            _ = v38
            _ = v39
        } else {
            v40 := GenS1{F0: 2, F1: 2147483647}
            _ = v40
        }
        switch (v5 + v5) {
        case "":
            fmt.Println(v24)
        case "hello":
            v33 -= v33
        case "a":
            if false {
                if (549 <= 65535) {
                    switch "foo bar" {
                    case "":
                        if true {
                            var v41 int = v33
                            var v42 uint32 = 65535
                            if len(v3) > 0 {
                                v3[0] = v5
                            }
                            _ = v41
                            _ = v42
                        } else {
                            v43 := GenS6{F0: false, F1: -1}
                            _ = v43
                        }
                    case "hello":
                        v1 += 581
                    case "foo bar":
                        for fi44 := 0; fi44 < 4; fi44++ {
                            v45 := GenS6{F0: false, F1: 1}
                            _ = fi44
                            _ = v45
                        }
                    }
                    for fi46 := 0; fi46 < 16; fi46++ {
                        if true {
                            v2 = 141
                            var v47 int16 = 685
                            _ = v47
                        } else {
                            switch (fi46 % ((fi46) | 1)) {
                            case 1:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 128:
                                v48 := GenS2{F0: v24, F1: v2}
                                _ = v48
                            case -9223372036854775808:
                                v49 := GenS7{F0: (v33 << (((v1 * v1)) & 63)), F1: 1}
                                _ = v49
                            default:
                                if (!(false && ((false || (892 > 0)) && ((v2 <= (v2 & v2)) || true)))) {
                                    switch (v33 ^ 639) {
                                    case -364:
                                        var v50 uint64 = uint64(v1)
                                        _ = v50
                                    case 0:
                                        v2 = v2
                                    case 4294967295:
                                        _ = len(v3)
                                    default:
                                        v24 = uint32(v22)
                                    }
                                } else {
                                    switch 0 {
                                    case -391:
                                        v3 = append(v3, "\n\t")
                                    case 59:
                                        _ = len(v3)
                                    case 2147483647:
                                        var v51 float64 = 1e10
                                        _ = v51
                                    default:
                                        v52 := GenS6{F0: ((v2 * 18446744073709551615) <= v2), F1: -32768}
                                        _ = v52
                                    }
                                }
                            }
                            fmt.Println(v33)
                            for fi53 := 0; fi53 < 3; fi53++ {
                                var v54 bool = false
                                v55 := GenS1{F0: uint32(v22), F1: 2147483647}
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                        }
                        var v56 uint32 = v24
                        _ = fi46
                        _ = v56
                    }
                    if false {
                        var v57 uint = 18446744073709551615
                        v58 := GenS8{F0: -687}
                        switch v5 {
                        case "hello":
                            fmt.Println(v1)
                        case "":
                            switch v5 {
                            case "":
                                v22 = (v22 * v22)
                            case "hello":
                                if (false && true) {
                                    v59 := GenS2{F0: v24, F1: 106}
                                    v1 = v1
                                    v60 := GenS9{F0: -128, F1: -899}
                                    _ = v59
                                    _ = v60
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v4[(v24 << ((v1) & 63))] = ((0.1 + 0.1) != float32(((1.5 / float64((float32((((0.1 - ((-2.25 + (float64(1.5) + -2.25)) / 1.5)) + (1.5 - 0.0)) - (float64(((0.0 - float32(-2.25)) / -2.25)) + (0.1 / (1e10 / float64(1.5)))))) / (float32((1.5 + 1.5)) * (3.14159 / 1.0))))) / (float64(3.14159) / float64(float32(1e10))))))
                                }
                            case "a":
                                if true {
                                    fmt.Println(v57, v57, v22)
                                }
                            default:
                                delete(v4, 65535)
                            }
                        case "a":
                            delete(v4, v24)
                        }
                        _ = v57
                        _ = v58
                    }
                }
                v2 -= (v2 << (((^(v1))) & 63))
                switch v33 {
                case 266:
                    if true {
                        v61 := GenS3{F0: (v1 & (v1 * uint(v22)))}
                        v1 = (v1 >> ((629) & 63))
                        _ = v61
                    }
                case 128:
                    for fi62 := 0; fi62 < 5; fi62++ {
                        _ = len(v3)
                        _ = fi62
                    }
                default:
                    v24 += v24
                }
            }
        default:
            if (!false) {
                v1 = (v1 ^ (v1 / (((v1 | v1)) | 1)))
            } else {
                var v63 string = v5
                _ = v63
            }
        }
        var v64 uint8 = 2
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
