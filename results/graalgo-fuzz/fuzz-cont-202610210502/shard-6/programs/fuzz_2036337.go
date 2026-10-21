package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 0
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    switch "hello" {
    case "foo bar":
        var v5 uint64 = 0
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 3; fi6++ {
            _ = len(v1)
            _ = fi6
        }
    case "\n\t":
        if (!true) {
            v7 := GenS1{F0: 1, F1: 4294967295, F2: 2}
            v2 -= v2
            if (uint64(v2) > 2) {
                v2 += v2
            }
            _ = v7
        } else {
            fmt.Println(v2)
        }
    default:
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS2{F0: (^(v2))}
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS2{F0: 1}
    var v11 float64 = 1e10
    for fi12 := 0; fi12 < 14; fi12++ {
        v11 += ((1e10 * -2.25) / (v11 * 3.14159))
        var v13 int8 = -17
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := v4[42]
        v2 = v2
        for fi16 := 0; fi16 < 14; fi16++ {
            if true {
                switch ((("hello" + "hello") + "") + "foo bar") {
                case "\n\t":
                    v1[(fi16 * -1)] = ((false || (v11 >= (((v11 - (-2.25 + v11)) / 1.0) / v11))) && (0.1 >= v11))
                case "a":
                    var v17 int64 = 0
                    _ = v17
                case "":
                    if true {
                        v18 := v4[v15]
                        var v19 string = ("hello" + "\n\t")
                        _ = v18
                        _ = v19
                    } else {
                        var v20 int16 = -358
                        var v21 uint = 813
                        _ = v20
                        _ = v21
                    }
                default:
                    var v22 uint8 = (v2 & v2)
                    _ = v22
                }
            } else {
                if ((v2 >> ((937) & 63)) >= (v2 << ((uint(v2)) & 63))) {
                    fmt.Println(v15, fi16, v15)
                    if true {
                        v23 := []int{(^(fi14)), fi14}
                        v24 := GenS1{F0: uint16(v2), F1: 0, F2: 18446744073709551615}
                        _ = v23
                        _ = v24
                    } else {
                        v25 := GenS2{F0: (v2 << ((65535) & 63))}
                        _ = v25
                    }
                    var v26 int32 = -1
                    _ = v26
                }
                delete(v3, fi16)
                var v27 uint = 226
                v28 := []int{fi14, fi16}
                _ = v27
                _ = v28
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    v11 = v11
    _ = len(v3)
    var v29 int32 = 42
    v30 := GenS1{F0: 32767, F1: 735, F2: 0}
    v29 = 618
    v31 := []int{1, 9223372036854775807}
    for fi32 := 0; fi32 < 12; fi32++ {
        switch "hello" {
        case "foo bar":
            var v33 bool = (2 > 559)
            _ = v33
        case "hello":
            if (!(true && true)) {
                var v34 float64 = ((float64(0.1) * 0.1) + v11)
                v35 := GenS3{F0: "hello", F1: (fi32 ^ (fi32 ^ fi32))}
                _ = v34
                _ = v35
            } else {
                var v36 int = fi32
                _ = v36
            }
        }
        v29 += int32(v2)
        _ = fi32
    }
    switch ("" + ("foo bar" + "a")) {
    case "":
        if true {
            fmt.Println(v11, v11, v11)
            v37 := GenS4{F0: 4294967295, F1: 681}
            _ = v37
        } else {
            if true {
                switch "hello" {
                case "foo bar":
                    v2 = 245
                case "a":
                    fmt.Println(v11, v11, v2)
                case "hello":
                    var v38 int8 = -1
                    _ = v38
                default:
                    switch 1 {
                    case -837:
                        switch ("hello" + "hello") {
                        case "":
                            var v39 int8 = -3
                            _ = v39
                        case "hello":
                            switch "" {
                            case "hello":
                                var v40 float32 = -2.25
                                _ = v40
                            case "a":
                                if true {
                                    v2 += 42
                                }
                            }
                        case "\n\t":
                            var v41 float32 = 3.14159
                            _ = v41
                        default:
                            for fi42 := 0; fi42 < 14; fi42++ {
                                v1[(fi42 * 2147483647)] = (1 == -9223372036854775808)
                                var v43 float64 = float64(1e10)
                                _ = fi42
                                _ = v43
                            }
                        }
                    default:
                        switch "\n\t" {
                        case "hello":
                            v11 -= (v11 / 1.5)
                        }
                    }
                }
                v44 := v1[-2147483648]
                v11 = (-2.25 + v11)
                _ = v44
            } else {
                if true {
                    if ((("foo bar" + "a") + "a") == "\n\t") {
                        v29 = (^(v29))
                        switch "a" {
                        case "\n\t":
                            for fi45 := 0; fi45 < 1; fi45++ {
                                if false {
                                    if (true && (true && (false && true))) {
                                        v11 = (v11 - -2.25)
                                        for fi46 := 0; fi46 < 7; fi46++ {
                                            v11 = v11
                                            v47 := GenS2{F0: v2}
                                            _ = fi46
                                            _ = v47
                                        }
                                        for fi48 := 0; fi48 < 3; fi48++ {
                                            var v49 int32 = (v29 << ((652) & 63))
                                            if true {
                                                v50 := v4[369]
                                                _ = v50
                                            }
                                            if (-2.25 > float32((v11 / (float64((1.5 * 0.1)) * (v11 / v11))))) {
                                                switch ("\n\t" + "foo bar") {
                                                case "a":
                                                    var v51 float64 = (v11 * (v11 + (v11 + (float64(-2.25) * float64(0.0)))))
                                                    _ = v51
                                                case "\n\t":
                                                    var v52 string = "foo bar"
                                                    _ = v52
                                                default:
                                                    v11 = 0.1
                                                }
                                                v53 := GenS5{F0: ((0.1 * -2.25) >= 0.1)}
                                                _ = v53
                                            } else {
                                                v29 = (v29 ^ v49)
                                                switch -904 {
                                                case 127:
                                                    var v54 uint64 = 334
                                                    _ = v54
                                                }
                                                var v55 string = "hello"
                                                v49 += 0
                                                if true {
                                                    switch "a" {
                                                    case "":
                                                        v56 := GenS3{F0: v55, F1: (fi45 ^ 9223372036854775807)}
                                                        _ = v56
                                                    case "\n\t":
                                                        v57 := GenS2{F0: v2}
                                                        _ = v57
                                                    default:
                                                        v58 := GenS6{F0: -9223372036854775808, F1: 9223372036854775807}
                                                        _ = v58
                                                    }
                                                    if true {
                                                        v11 = v11
                                                        if true {
                                                            if len(v31) > 0 {
                                                                v31[0] = (fi48 << ((4294967295) & 63))
                                                            }
                                                        }
                                                    }
                                                    var v59 uint64 = 0
                                                    _ = v59
                                                }
                                                _ = v55
                                            }
                                            _ = fi48
                                            _ = v49
                                        }
                                    } else {
                                        v29 += (v29 - v29)
                                        if len(v31) >= 2 {
                                            v31 = v31[1:len(v31)]
                                        }
                                        v60 := GenS7{F0: 2147483647}
                                        _ = v60
                                    }
                                }
                                for fi61 := 0; fi61 < 14; fi61++ {
                                    v29 += (v29 & (-(v29)))
                                    if (255 >= 65535) {
                                        delete(v4, 357)
                                    } else {
                                        var v62 uint16 = 32767
                                        _ = v62
                                    }
                                    _ = fi61
                                }
                                _ = fi45
                            }
                        default:
                            var v63 int16 = 1
                            _ = v63
                        }
                    } else {
                        if true {
                            if true {
                                v64 := v1[-257]
                                _ = v64
                            }
                        } else {
                            v31 = append(v31, 42)
                            var v65 uint = 677
                            v11 -= 1.5
                            _ = v65
                        }
                        if len(v31) > 0 {
                            v31[0] = 0
                        }
                        if (0.0 < ((((-2.25 + (0.0 - 0.1)) * 0.1) / ((float32(v11) / (-2.25 * (0.1 + -2.25))) * 0.0)) / 1.0)) {
                            switch 88 {
                            case 1:
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                            case -348:
                                v66 := GenS4{F0: 2147483648, F1: 32767}
                                _ = v66
                            case 128:
                                var v67 string = "hello"
                                _ = v67
                            }
                            if (!true) {
                                if false {
                                    fmt.Println(v11)
                                } else {
                                    v11 -= v11
                                }
                                fmt.Println(v29, v11)
                            } else {
                                switch -663 {
                                case -709:
                                    if true {
                                        var v68 int = -9223372036854775808
                                        v2 += 127
                                        _ = v68
                                    }
                                case 237:
                                    for fi69 := 0; fi69 < 17; fi69++ {
                                        var v70 uint8 = 200
                                        for fi71 := 0; fi71 < 18; fi71++ {
                                            var v72 float64 = ((1.5 + ((v11 * -2.25) - 0.1)) * (v11 * -2.25))
                                            _ = fi71
                                            _ = v72
                                        }
                                        _ = fi69
                                        _ = v70
                                    }
                                case 0:
                                    var v73 uint = 0
                                    _ = v73
                                default:
                                    if len(v31) >= 2 {
                                        v31 = v31[1:len(v31)]
                                    }
                                }
                            }
                        } else {
                            fmt.Println(v29, v2, v11)
                        }
                    }
                }
            }
            switch -366 {
            case 588:
                for fi74 := 0; fi74 < 16; fi74++ {
                    v2 += 42
                    fmt.Println(fi74, v2, v2)
                    switch "hello" {
                    case "hello":
                        for fi75 := 0; fi75 < 14; fi75++ {
                            var v76 uint8 = 128
                            _ = fi75
                            _ = v76
                        }
                    }
                    _ = fi74
                }
            case -1:
                var v77 uint8 = (v2 * (v2 + (v2 ^ v2)))
                _ = v77
            case -46:
                v3[-14] = 4294967295
            }
        }
    }
    v78 := GenS8{F0: 2147483648, F1: -6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v78
}
