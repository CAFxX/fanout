package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint8; F1 string; F2 uint16 }
type GenS3 struct { F0 uint; F1 int64; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS7 struct { F0 int16; F1 uint; F2 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 uint16; F2 uint64 }
type GenS11 struct { F0 uint8; F1 int64 }
type GenS12 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint; F1 uint; F2 int }
type GenS15 struct { F0 bool; F1 int8; F2 uint64 }
type GenS16 struct { F0 int32; F1 int8; F2 uint }
type GenS17 struct { F0 string }
func genFn0() bool {
    switch 104 {
    case 0:
        var v4 int8 = 42
        _ = v4
    case 32767:
        v5 := GenS1{F0: -32768, F1: true}
        _ = v5
    case 355:
        var v6 int16 = 42
        _ = v6
    default:
        var v7 int16 = 32767
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        if ("\n\t" == ("a" + "\n\t")) {
            var v9 int8 = 42
            v10 := GenS2{F0: 1, F1: "a", F2: 737}
            v9 += -51
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    return false
}
func genFn1(p0 int32, p1 uint64, p2 bool) int8 {
    var v11 bool = p2
    var v12 int16 = 120
    var v13 int = 1
    switch ("hello" + "a") {
    case "":
        var v14 float64 = 1.5
        _ = v14
    case "a":
        if false {
            fmt.Println(v11)
            if (false || true) {
                v12 = 168
                if true {
                    fmt.Println(p0, v11, v12)
                } else {
                    var v15 float64 = 0.1
                    v16 := GenS2{F0: 42, F1: "foo bar", F2: 2}
                    _ = v15
                    _ = v16
                }
            }
        } else {
            var v17 string = ""
            _ = v17
        }
    case "foo bar":
        v18 := []int{v13}
        _ = v18
    default:
        p2 = v11
    }
    v19 := GenS2{F0: 2, F1: "\n\t", F2: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    return int8(v12)
}

func main() {
    var v1 uint8 = 158
    v2 := []int{-575, 1, 1}
    v3 := make(map[int]int)
    v20 := GenS3{F0: 0, F1: -56, F2: 120}
    for fi21 := 0; fi21 < 11; fi21++ {
        if true {
            switch (-(fi21)) {
            case -528:
                _ = len(v3)
            case 128:
                if true {
                    fmt.Println(v1, fi21)
                    fmt.Println(fi21, fi21)
                    v3[127] = -1
                } else {
                    v1 = 2
                }
            default:
                var v22 int8 = 127
                _ = v22
            }
            switch fi21 {
            case -555:
                var v23 string = "foo bar"
                _ = v23
            case 0:
                v24 := GenS4{F0: v1, F1: 332, F2: (v1 / (((v1 << ((422) & 63))) | 1))}
                _ = v24
            case -9223372036854775808:
                var v25 float64 = ((1.5 * 1e10) * 3.14159)
                _ = v25
            }
            for fi26 := 0; fi26 < 8; fi26++ {
                v1 += v1
                v1 += 42
                _ = fi26
            }
        } else {
            if false {
                v27 := v3[fi21]
                var v28 uint8 = (v1 | v1)
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                var v30 uint = 65535
                v1 = v1
                var v31 int16 = 32767
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi21
    }
    v1 += v1
    var v32 uint64 = 4294967295
    if (0.0 > (1e10 + 0.0)) {
        v3[-9223372036854775808] = -9223372036854775808
    }
    fmt.Println(v32)
    v33 := GenS4{F0: 42, F1: 255, F2: v1}
    v34 := GenS2{F0: v1, F1: "a", F2: 32767}
    v35 := GenS2{F0: v1, F1: "foo bar", F2: 357}
    v36 := GenS2{F0: v1, F1: "\n\t", F2: 32767}
    var v37 int8 = -128
    fmt.Println(v32, v1, v37)
    if (127 <= v1) {
        v37 = (v37 | v37)
        v32 = v32
        switch "\n\t" {
        case "":
            var v38 bool = false
            _ = v38
        case "a":
            if false {
                v32 = v32
            } else {
                var v39 uint32 = 4294967295
                v40 := GenS4{F0: v1, F1: 32767, F2: (v1 - v1)}
                _ = v39
                _ = v40
            }
        case "\n\t":
            v41 := GenS4{F0: 127, F1: 65535, F2: 128}
            _ = v41
        }
    }
    v42 := GenS5{F0: 42, F1: 9223372036854775807}
    for fi43 := 0; fi43 < 8; fi43++ {
        var v44 int8 = (v37 - (v37 + genFn1(1, v32, false)))
        fmt.Println(v32, v32)
        v45 := GenS1{F0: 0, F1: (false || (!(2 == 1)))}
        _ = fi43
        _ = v44
        _ = v45
    }
    switch "" {
    case "hello":
        if (!true) {
            switch -1 {
            case 653:
                v32 = v32
            case 32767:
                fmt.Println(v32)
            case -9223372036854775808:
                fmt.Println(v32, v1, v37)
            }
            if false {
                if (((v1 * (v1 / ((v1) | 1))) > v1) || true) {
                    switch -390 {
                    case -2147483648:
                        _ = len(v3)
                    }
                    v46 := GenS5{F0: 0, F1: 128}
                    fmt.Println(v37, v37)
                    _ = v46
                }
            }
        } else {
            v47 := GenS6{F0: 42, F1: 0, F2: 1}
            v48 := GenS4{F0: (v1 % ((2) | 1)), F1: 533, F2: v1}
            if true {
                v32 = (v32 / ((v32) | 1))
                if true {
                    switch "a" {
                    case "\n\t":
                        _ = len(v3)
                    default:
                        switch "a" {
                        case "foo bar":
                            var v49 int32 = 1
                            _ = v49
                        case "hello":
                            fmt.Println(v37)
                        case "\n\t":
                            v37 = v37
                        default:
                            v50 := GenS7{F0: 0, F1: 65535, F2: -1}
                            _ = v50
                        }
                    }
                    fmt.Println(v1, v37, v1)
                    v1 -= 30
                } else {
                    _ = len(v2)
                    switch "foo bar" {
                    case "a":
                        switch "" {
                        case "":
                            if true {
                                if (340 != 9223372036854775807) {
                                    if (((1 == 710) && false) && (!true)) {
                                        if false {
                                            fmt.Println(v37)
                                        }
                                        fmt.Println(v1)
                                    }
                                    if false {
                                        v32 = 0
                                    }
                                    v51 := GenS6{F0: 2, F1: 42, F2: 2147483648}
                                    _ = v51
                                } else {
                                    var v52 uint8 = 184
                                    for fi53 := 0; fi53 < 10; fi53++ {
                                        if false {
                                            v54 := GenS8{F0: 979}
                                            _ = v54
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = (fi53 - -9223372036854775808)
                                        }
                                        v55 := GenS5{F0: -2147483648, F1: (fi53 >> ((251) & 63))}
                                        _ = fi53
                                        _ = v55
                                    }
                                    _ = v52
                                }
                                if ((3.14159 + (1.5 + 0.1)) >= float64(-2.25)) {
                                    v3[-193] = 320
                                    v32 = (v32 / ((48) | 1))
                                } else {
                                    for fi56 := 0; fi56 < 10; fi56++ {
                                        var v57 uint8 = 42
                                        var v58 uint = 1
                                        if ((!false) || (251 > 65535)) {
                                            if (2 == (^(v58))) {
                                                for fi59 := 0; fi59 < 3; fi59++ {
                                                    var v60 float32 = 0.0
                                                    var v61 int8 = genFn1(448, (v32 << ((v58) & 63)), false)
                                                    v37 = genFn1(-956, 0, true)
                                                    v3[32767] = (-(fi59))
                                                    _ = fi59
                                                    _ = v60
                                                    _ = v61
                                                }
                                                var v62 uint64 = v32
                                                for fi63 := 0; fi63 < 10; fi63++ {
                                                    var v64 int = fi56
                                                    if (968 > v58) {
                                                        var v65 uint8 = v1
                                                        fmt.Println(v37, v37)
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v3[v64] = -846
                                                        case "hello":
                                                            fmt.Println(fi63, v64, v62)
                                                        }
                                                        _ = v65
                                                    }
                                                    v1 += 1
                                                    fmt.Println(v1, v57, v57)
                                                    _ = fi63
                                                    _ = v64
                                                }
                                                _ = v62
                                            }
                                            for fi66 := 0; fi66 < 8; fi66++ {
                                                v67 := GenS9{F0: fi56}
                                                v68 := GenS1{F0: 32767, F1: (346 > 4294967295)}
                                                _ = fi66
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                        _ = fi56
                                        _ = v57
                                        _ = v58
                                    }
                                    v37 = v37
                                    switch -769 {
                                    case 42:
                                        var v69 uint8 = 255
                                        _ = v69
                                    case 848:
                                        v70 := GenS9{F0: 0}
                                        _ = v70
                                    case 0:
                                        for fi71 := 0; fi71 < 6; fi71++ {
                                            v72 := GenS10{F0: int(v37), F1: 2, F2: 2}
                                            if false {
                                                switch "foo bar" {
                                                case "":
                                                    v1 = v1
                                                case "\n\t":
                                                    v3[(fi71 * fi71)] = (^(fi71))
                                                case "hello":
                                                    _ = len(v2)
                                                }
                                                for fi73 := 0; fi73 < 8; fi73++ {
                                                    v37 = (v37 + genFn1(1, (v32 << ((945) & 63)), false))
                                                    _ = len(v3)
                                                    if (false && false) {
                                                        v37 += (v37 - v37)
                                                        fmt.Println(v37)
                                                        var v74 int16 = -32768
                                                        switch fi71 {
                                                        case -9223372036854775808:
                                                            switch ("hello" + "a") {
                                                            case "foo bar":
                                                                if true {
                                                                    fmt.Println(v74, v37)
                                                                }
                                                            }
                                                        }
                                                        _ = v74
                                                    } else {
                                                        if (((0.1 * -2.25) * 0.0) < 0.0) {
                                                            v37 = (v37 << ((42) & 63))
                                                            if false {
                                                                v3[-246] = fi73
                                                                var v75 uint32 = 402
                                                                if ((0.1 * (((0.0 * (0.0 + 0.1)) / 1e10) * (0.1 * 0.0))) != (float64(-2.25) + (1e10 * float64(1.5)))) {
                                                                    fmt.Println(v75)
                                                                    v75 = (v75 & (v75 * (v75 * v75)))
                                                                    v76 := GenS11{F0: v1, F1: -843}
                                                                    switch ((("\n\t" + "foo bar") + ("" + "\n\t")) + "\n\t") {
                                                                    case "a":
                                                                        var v77 int32 = 295
                                                                        _ = v77
                                                                    case "foo bar":
                                                                        fmt.Println(v32)
                                                                    case "":
                                                                        v1 = v1
                                                                    }
                                                                    _ = v76
                                                                } else {
                                                                    v75 -= uint32(v32)
                                                                    var v78 uint = 551
                                                                    _ = v78
                                                                }
                                                                fmt.Println(fi73, v32, v37)
                                                                fmt.Println(v1, fi73, fi71)
                                                                if (!(!(true || true))) {
                                                                    if true {
                                                                        switch -399 {
                                                                        case 2147483647:
                                                                            v32 = (v32 | v32)
                                                                        case 32768:
                                                                            v37 += v37
                                                                        case 127:
                                                                            fmt.Println(fi73, v75)
                                                                        default:
                                                                            v79 := GenS12{F0: -529, F1: v75, F2: 2}
                                                                            _ = v79
                                                                        }
                                                                    }
                                                                    v80 := GenS10{F0: (fi71 | 127), F1: 32767, F2: 636}
                                                                    _ = v80
                                                                }
                                                                _ = v75
                                                            } else {
                                                                v1 += 3
                                                                v81 := GenS13{F0: (!false)}
                                                                _ = len(v3)
                                                                _ = v81
                                                            }
                                                        } else {
                                                            v1 += (v1 + v1)
                                                        }
                                                    }
                                                    var v82 float64 = 1e10
                                                    v83 := GenS1{F0: -1, F1: false}
                                                    _ = fi73
                                                    _ = v82
                                                    _ = v83
                                                }
                                                switch "a" {
                                                case "hello":
                                                    v84 := GenS14{F0: 255, F1: 18446744073709551615, F2: fi71}
                                                    _ = v84
                                                default:
                                                    var v85 uint32 = 2147483648
                                                    _ = v85
                                                }
                                            }
                                            var v86 uint16 = 65535
                                            _ = fi71
                                            _ = v72
                                            _ = v86
                                        }
                                    }
                                }
                                v87 := GenS15{F0: (65535 != 42), F1: v37, F2: 18446744073709551615}
                                _ = v87
                            }
                        case "hello":
                            var v88 uint64 = 560
                            _ = v88
                        case "foo bar":
                            v32 = (v32 / ((v32) | 1))
                        default:
                            for fi89 := 0; fi89 < 18; fi89++ {
                                v37 += v37
                                _ = fi89
                            }
                        }
                    case "foo bar":
                        v90 := GenS2{F0: v1, F1: "", F2: 42}
                        _ = v90
                    case "hello":
                        var v91 float32 = (-2.25 / 1e10)
                        _ = v91
                    default:
                        _ = len(v3)
                    }
                }
            } else {
                if false {
                    _ = len(v3)
                }
                v92 := GenS2{F0: 230, F1: (("\n\t" + "\n\t") + ("a" + "foo bar")), F2: 2}
                v93 := GenS16{F0: 0, F1: (v37 * v37), F2: 355}
                _ = v92
                _ = v93
            }
            _ = v47
            _ = v48
        }
    case "a":
        fmt.Println(v37, v32, v32)
    case "foo bar":
        v94 := GenS17{F0: "\n\t"}
        _ = v94
    default:
        if true {
            _ = len(v3)
        }
    }
    var v95 int8 = genFn1(443, (^(v32)), false)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v42
    _ = v95
}
