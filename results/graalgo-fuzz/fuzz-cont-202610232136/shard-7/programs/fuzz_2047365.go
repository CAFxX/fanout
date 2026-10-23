package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int32; F1 int8; F2 int32 }
type GenS5 struct { F0 bool; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 uint32 = 2
    var v2 uint = 0
    v3 := []string{}
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: -128, F1: "foo bar"}
    v6 := GenS2{F0: -32768}
    v1 -= v1
    for fi7 := 0; fi7 < 7; fi7++ {
        switch (^(fi7)) {
        case -1:
            if (3.14159 < 1.5) {
                v2 = v2
            }
        case -701:
            v2 -= 701
        case 0:
            if true {
                v8 := GenS2{F0: 1}
                _ = v8
            }
        default:
            if false {
                var v9 uint = v2
                _ = v9
            }
        }
        if false {
            v1 = v1
            var v10 int16 = -32768
            v2 = v2
            _ = v10
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS3{F0: 9223372036854775807, F1: (2 <= 42)}
            _ = fi11
            _ = v12
        }
        _ = fi7
    }
    if (0 < 138) {
        if false {
            switch 0 {
            case 32768:
                var v13 bool = ((!false) && true)
                _ = v13
            case 42:
                if true {
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v2 += (v2 % ((0) | 1))
                        var v14 int64 = 0
                        _ = v14
                    } else {
                        fmt.Println(v1, v2)
                        _ = len(v3)
                        var v15 float32 = -2.25
                        _ = v15
                    }
                    var v16 int = -1
                    v2 = (^(v2))
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        for fi18 := 0; fi18 < 9; fi18++ {
                            v19 := GenS4{F0: -2147483648, F1: -128, F2: 51}
                            v1 -= (v1 >> ((v2) & 63))
                            _ = fi18
                            _ = v19
                        }
                        _ = fi17
                    }
                }
            case -379:
                _ = len(v3)
            default:
                if ((true || false) && false) {
                    if (42 > 2) {
                        if true {
                            v2 = v2
                            fmt.Println(v2, v1)
                        }
                    } else {
                        var v20 int32 = 2147483647
                        _ = v20
                    }
                } else {
                    if (18446744073709551615 >= 2) {
                        for fi21 := 0; fi21 < 8; fi21++ {
                            if true {
                                v22 := GenS5{F0: (!true), F1: int8(v2), F2: 45}
                                _ = v22
                            } else {
                                switch ("\n\t" + "\n\t") {
                                case "hello":
                                    switch ("\n\t" + ("\n\t" + "")) {
                                    case "\n\t":
                                        v23 := GenS1{F0: 127, F1: "foo bar"}
                                        _ = v23
                                    case "foo bar":
                                        if (true || (-32768 < 0)) {
                                            if (!false) {
                                                var v24 uint8 = 51
                                                if (true || false) {
                                                    if len(v3) > 0 {
                                                        v3[0] = "a"
                                                    }
                                                } else {
                                                    fmt.Println(v1, fi21)
                                                }
                                                _ = v24
                                            }
                                        }
                                    }
                                case "":
                                    for fi25 := 0; fi25 < 14; fi25++ {
                                        var v26 uint16 = 1
                                        _ = fi25
                                        _ = v26
                                    }
                                default:
                                    v1 = (v1 / ((v1) | 1))
                                }
                                switch "a" {
                                case "\n\t":
                                    v1 -= 2147483647
                                case "foo bar":
                                    v1 -= 861
                                }
                            }
                            v2 += v2
                            if false {
                                if (4294967295 <= 11) {
                                    if (false || false) {
                                        v1 -= v1
                                        for fi27 := 0; fi27 < 6; fi27++ {
                                            fmt.Println(v1, v2)
                                            var v28 bool = false
                                            _ = fi27
                                            _ = v28
                                        }
                                        if true {
                                            var v29 int64 = 868
                                            v1 -= v1
                                            for fi30 := 0; fi30 < 10; fi30++ {
                                                var v31 float32 = -2.25
                                                fmt.Println(v1, fi30, fi21)
                                                if (true || (!false)) {
                                                    fmt.Println(fi21, fi30)
                                                    var v32 int32 = -667
                                                    var v33 uint64 = 973
                                                    _ = v32
                                                    _ = v33
                                                }
                                                _ = fi30
                                                _ = v31
                                            }
                                            _ = v29
                                        }
                                    } else {
                                        var v34 uint64 = 42
                                        for fi35 := 0; fi35 < 3; fi35++ {
                                            v34 = (v34 + (v34 * 18446744073709551615))
                                            v2 = (v2 * v2)
                                            if true {
                                                var v36 uint64 = (v34 ^ v34)
                                                if (337 == 1) {
                                                    var v37 bool = false
                                                    _ = v37
                                                } else {
                                                    var v38 uint32 = (v1 + (v1 / (((v1 & (v1 | (v1 << ((v2) & 63))))) | 1)))
                                                    var v39 float32 = float32(((1.5 / 3.14159) + 1.5))
                                                    _ = v38
                                                    _ = v39
                                                }
                                                var v40 int8 = 100
                                                _ = v36
                                                _ = v40
                                            }
                                            _ = fi35
                                        }
                                        var v41 uint64 = 18446744073709551615
                                        _ = v34
                                        _ = v41
                                    }
                                    if false {
                                        var v42 float64 = 1e10
                                        _ = v42
                                    } else {
                                        _ = len(v3)
                                    }
                                } else {
                                    v4[v1] = (-(fi21))
                                    v43 := GenS2{F0: -772}
                                    switch (fi21 >> ((v2) & 63)) {
                                    case 271:
                                        switch fi21 {
                                        case 32768:
                                            v1 = (v1 % ((870) | 1))
                                        }
                                    default:
                                        _ = len(v3)
                                    }
                                    _ = v43
                                }
                                fmt.Println(fi21, v1, v2)
                            } else {
                                var v44 int = (^(fi21))
                                if true {
                                    if false {
                                        var v45 int64 = 1
                                        v45 = (v45 << ((246) & 63))
                                        _ = v45
                                    } else {
                                        fmt.Println(fi21, fi21, v2)
                                        v44 = (-(v44))
                                        var v46 int = -190
                                        _ = v46
                                    }
                                    for fi47 := 0; fi47 < 7; fi47++ {
                                        v48 := GenS2{F0: 0}
                                        if (18446744073709551615 != 645) {
                                            fmt.Println(v44, v44)
                                        } else {
                                            if false {
                                                v2 = (v2 * 695)
                                                var v49 float64 = float64(((0.1 * 3.14159) - 1.5))
                                                _ = v49
                                            } else {
                                                v2 = v2
                                                v50 := GenS2{F0: int16(v2)}
                                                _ = v50
                                            }
                                            var v51 float64 = 0.1
                                            _ = v51
                                        }
                                        v1 = (^(v1))
                                        _ = fi47
                                        _ = v48
                                    }
                                    var v52 int8 = -1
                                    _ = v52
                                } else {
                                    for fi53 := 0; fi53 < 9; fi53++ {
                                        var v54 float32 = float32(0.1)
                                        v54 = ((v54 / 0.1) / (1e10 - ((v54 / 1.0) + (float32(0.0) - v54))))
                                        _ = fi53
                                        _ = v54
                                    }
                                    for fi55 := 0; fi55 < 15; fi55++ {
                                        v4[v1] = (fi55 ^ (v44 << ((920) & 63)))
                                        _ = fi55
                                    }
                                    var v56 int32 = 203
                                    _ = v56
                                }
                                _ = v44
                            }
                            _ = fi21
                        }
                    } else {
                        v2 = (v2 / ((v2) | 1))
                        var v57 uint = (^(v2))
                        if true {
                            v57 = (v57 / (((v57 << (((v57 / ((uint(v1)) | 1))) & 63))) | 1))
                            if (!(((false && (!(!(2 > 65535)))) && true) && true)) {
                                var v58 float64 = 1e10
                                for fi59 := 0; fi59 < 7; fi59++ {
                                    v1 -= (v1 / (((v1 >> ((247) & 63))) | 1))
                                    v60 := GenS6{F0: v1}
                                    _ = fi59
                                    _ = v60
                                }
                                v58 = (v58 + 3.14159)
                                _ = v58
                            }
                        } else {
                            if (0 > 42) {
                                v61 := v4[v1]
                                v57 = 677
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                                _ = v61
                            } else {
                                v2 -= 506
                            }
                        }
                        _ = v57
                    }
                }
            }
            switch "foo bar" {
            case "\n\t":
                if (!(122 > 127)) {
                    switch 9223372036854775807 {
                    case -683:
                        for fi62 := 0; fi62 < 1; fi62++ {
                            v1 = (^(v1))
                            _ = fi62
                        }
                    case 32767:
                        var v63 int32 = -1
                        _ = v63
                    case 127:
                        var v64 float32 = -2.25
                        _ = v64
                    default:
                        v65 := GenS7{F0: 2}
                        _ = v65
                    }
                }
            case "hello":
                var v66 uint8 = 83
                _ = v66
            case "foo bar":
                var v67 int16 = -32768
                _ = v67
            }
        }
    }
    var v68 int = 4294967295
    switch v68 {
    case 261:
        switch v68 {
        case -9223372036854775808:
            var v69 uint = 562
            _ = v69
        case 364:
            var v70 int8 = -92
            _ = v70
        default:
            var v71 uint8 = 248
            _ = v71
        }
    case -825:
        delete(v4, v1)
    }
    for fi72 := 0; fi72 < 15; fi72++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v68, fi72, v68)
        if (false || (!(2 != 255))) {
            if false {
                v2 = (v2 << ((65535) & 63))
            }
            if (0.0 == 3.14159) {
                switch "" {
                case "\n\t":
                    v2 -= v2
                case "foo bar":
                    if (!(!false)) {
                        var v73 bool = false
                        _ = v73
                    } else {
                        for fi74 := 0; fi74 < 4; fi74++ {
                            fmt.Println(v1, v68, fi74)
                            if (((-2.25 * float64(-2.25)) + (0.1 - (float64((float32(-2.25) * (-2.25 / ((3.14159 / 0.1) + (0.0 * (float32(float64(0.0)) + (0.0 - float32(0.1)))))))) / (float64((-2.25 / (0.1 * (((0.1 - float32((0.1 / 1.0))) - (float32((((float64(((1e10 + 1.5) / 1.5)) - 1.5) / ((3.14159 - ((1e10 / 1.0) * 1e10)) / -2.25)) / float64((float32((0.1 - 0.0)) + -2.25)))) - 0.1)) + (3.14159 - (1.5 / 1.0)))))) - -2.25)))) > -2.25) {
                                var v75 int = (-(fi72))
                                _ = v75
                            } else {
                                v1 -= (v1 | (v1 - v1))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v76 uint32 = v1
                                _ = v76
                            }
                            _ = fi74
                        }
                    }
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                v3 = append(v3, "")
                v2 = v2
            }
            v77 := v4[(v1 | (v1 ^ v1))]
            _ = v77
        }
        _ = fi72
    }
    _ = len(v4)
    v1 = (^(v1))
    var v78 uint32 = 2147483648
    var v79 string = "hello"
    fmt.Println(v2, v2, v2)
    v78 -= 457
    for fi80 := 0; fi80 < 4; fi80++ {
        var v81 string = v79
        _ = fi80
        _ = v81
    }
    var v82 uint32 = (v1 >> ((v2) & 63))
    switch ("\n\t" + v79) {
    case "":
        v1 = 550
    case "\n\t":
        v83 := GenS3{F0: 595, F1: false}
        _ = v83
    }
    var v84 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v68:", v68)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v82:", v82)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v78
    _ = v79
    _ = v82
    _ = v84
}
