package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 string; F1 bool; F2 uint8 }

func main() {
    var v1 uint64 = 0
    v2 := []string{"foo bar", "\n\t", "a", "foo bar"}
    v3 := make(map[uint32]uint32)
    var v4 int16 = 42
    v5 := make(map[uint32]int)
    v6 := []int{416, 4294967295, -9223372036854775808}
    fmt.Println(v1)
    var v7 uint16 = 426
    var v8 float32 = -2.25
    var v9 uint = 155
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 170 {
    case 32767:
        var v10 uint16 = 0
        _ = v10
    case -2147483648:
        var v11 uint8 = 99
        _ = v11
    }
    if (v7 >= 42) {
        v8 = 0.0
    }
    _ = len(v6)
    v4 = (v4 * -32768)
    for fi12 := 0; fi12 < 15; fi12++ {
        _ = len(v6)
        _ = fi12
    }
    v13 := GenS1{F0: 100, F1: v9, F2: (^(v1))}
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(fi14, v9, v9)
        var v15 bool = false
        var v16 uint8 = 0
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        var v17 int = 2147483647
        for fi18 := 0; fi18 < 13; fi18++ {
            delete(v5, 912)
            _ = fi18
        }
        _ = v17
    } else {
        if false {
            for fi19 := 0; fi19 < 19; fi19++ {
                if (true && ((v4 ^ v4) <= (v4 + v4))) {
                    switch ("hello" + "") {
                    case "hello":
                        v1 += 188
                    case "foo bar":
                        for fi20 := 0; fi20 < 12; fi20++ {
                            v1 += 4294967295
                            _ = fi20
                        }
                    case "":
                        if true {
                            for fi21 := 0; fi21 < 10; fi21++ {
                                fmt.Println(v9)
                                v22 := v3[174]
                                if true {
                                    if false {
                                        var v23 uint64 = v1
                                        v7 += v7
                                        v5[9] = 32768
                                        _ = v23
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        v4 = v4
                                    case "\n\t":
                                        v24 := v3[v22]
                                        _ = v24
                                    default:
                                        if ((496 >= 4294967295) || (0 == v9)) {
                                            var v25 float64 = float64(v8)
                                            var v26 uint64 = uint64(v22)
                                            var v27 uint32 = (v22 * uint32(v26))
                                            if true {
                                                switch fi19 {
                                                case -2147483648:
                                                    v28 := GenS1{F0: 42, F1: (v9 * v9), F2: v1}
                                                    _ = v28
                                                }
                                                v26 = v26
                                                v29 := GenS1{F0: int8(fi21), F1: 2, F2: 2}
                                                var v30 int = (-(fi21))
                                                var v31 float32 = v8
                                                v22 = v27
                                                var v32 int = (fi19 + v30)
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                            if (-1 <= 19) {
                                                v1 += (v26 ^ 4294967295)
                                                if true {
                                                    var v33 uint8 = 42
                                                    v25 = (v25 / 1.0)
                                                    v34 := GenS2{F0: (v1 << ((v9) & 63))}
                                                    if ((true || false) && false) {
                                                        v25 -= (((float64(v8) + v25) / 1.5) + (float64((v8 * v8)) / v25))
                                                        v35 := v3[uint32(v9)]
                                                        fmt.Println(v1)
                                                        _ = v35
                                                    } else {
                                                        _ = len(v5)
                                                    }
                                                    _ = v33
                                                    _ = v34
                                                }
                                            }
                                            v36 := v5[(v27 | 26)]
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                            _ = v36
                                        }
                                    }
                                    v37 := v5[65535]
                                    v4 = (v4 ^ -245)
                                    var v38 float32 = (v8 / v8)
                                    switch fi19 {
                                    case 128:
                                        switch "a" {
                                        case "foo bar":
                                            v2 = append(v2, "\n\t")
                                        }
                                        if true {
                                            var v39 string = "hello"
                                            _ = v39
                                        }
                                        if (!(v1 == (v1 - uint64(v7)))) {
                                            var v40 uint64 = (v1 ^ 4294967295)
                                            if ((v38 / 3.14159) != v8) {
                                                v41 := GenS1{F0: -1, F1: 0, F2: 42}
                                                var v42 uint64 = v1
                                                _ = v41
                                                _ = v42
                                            }
                                            _ = v40
                                        }
                                    case 32767:
                                        v43 := GenS3{F0: (v1 / ((1) | 1)), F1: ("foo bar" + "\n\t")}
                                        _ = v43
                                    case 42:
                                        v44 := v5[(^(v22))]
                                        var v45 int = -757
                                        _ = v44
                                        _ = v45
                                    default:
                                        switch (v37 << ((1) & 63)) {
                                        case -1:
                                            v46 := GenS3{F0: (^(v1)), F1: (("a" + "a") + "hello")}
                                            _ = v46
                                        case 558:
                                            v4 += v4
                                        case 447:
                                            var v47 int64 = -1
                                            _ = v47
                                        default:
                                            if false {
                                                v48 := GenS4{F0: -1, F1: uint16(v9)}
                                                if ((true && (!true)) || (((v22 * v22) < uint32(v7)) && (false && ((v7 / (((^(v7))) | 1)) == v7)))) {
                                                    var v49 uint64 = v1
                                                    v9 = v9
                                                    v7 -= (v7 % ((v7) | 1))
                                                    _ = v49
                                                }
                                                v8 -= (v38 - 1.5)
                                                if len(v2) > 0 {
                                                    v2[0] = (("\n\t" + "foo bar") + "hello")
                                                }
                                                _ = v48
                                            }
                                        }
                                    }
                                    _ = len(v5)
                                    _ = v37
                                    _ = v38
                                } else {
                                    v50 := GenS5{F0: 100, F1: int64(v7)}
                                    v51 := GenS6{F0: ("" + "a"), F1: false, F2: 0}
                                    v7 = 2
                                    _ = v50
                                    _ = v51
                                }
                                _ = fi21
                                _ = v22
                            }
                            fmt.Println(v9, v1)
                        }
                    default:
                        if (42 != 100) {
                            var v52 int64 = -281
                            var v53 uint64 = v1
                            if false {
                                if (1 < 128) {
                                    var v54 int = 2147483647
                                    v4 -= 0
                                    for fi55 := 0; fi55 < 19; fi55++ {
                                        fmt.Println(fi55)
                                        _ = fi55
                                    }
                                    _ = v54
                                }
                                v52 -= 1
                            } else {
                                var v56 uint16 = 65535
                                if (false && true) {
                                    fmt.Println(v1)
                                    v57 := GenS5{F0: 127, F1: v52}
                                    v4 -= v4
                                    _ = v57
                                }
                                _ = v56
                            }
                            _ = v52
                            _ = v53
                        } else {
                            v2 = append(v2, ("" + "\n\t"))
                        }
                    }
                    v58 := GenS6{F0: "", F1: ((false && false) || false), F2: 113}
                    v1 = v1
                    _ = v58
                } else {
                    for fi59 := 0; fi59 < 15; fi59++ {
                        v4 -= (v4 % (((v4 | (v4 >> (((v9 + (v9 + uint(v7)))) & 63)))) | 1))
                        _ = len(v2)
                        var v60 uint = v9
                        _ = fi59
                        _ = v60
                    }
                }
                _ = fi19
            }
            v6 = append(v6, 0)
        }
    }
    for fi61 := 0; fi61 < 8; fi61++ {
        var v62 uint = v9
        _ = len(v3)
        fmt.Println(v8)
        _ = fi61
        _ = v62
    }
    var v63 string = "hello"
    var v64 uint16 = v7
    v64 = v64
    for fi65 := 0; fi65 < 8; fi65++ {
        var v66 int8 = -1
        var v67 uint16 = 104
        switch fi65 {
        case 9223372036854775807:
            if (false && true) {
                if (true || true) {
                    if (v67 == 65535) {
                        var v68 int8 = (v66 - 31)
                        v69 := GenS5{F0: v68, F1: 0}
                        if (true && true) {
                            var v70 int8 = (v66 % ((v66) | 1))
                            fmt.Println(v63, v67, v67)
                            _ = v70
                        } else {
                            v71 := GenS3{F0: (v1 ^ v1), F1: v63}
                            fmt.Println(v66)
                            _ = v71
                        }
                        _ = v68
                        _ = v69
                    }
                } else {
                    v3[841] = 552
                    v2 = append(v2, "a")
                    v66 = v66
                }
                var v72 uint16 = v7
                _ = v72
            }
        default:
            fmt.Println(v8)
        }
        _ = fi65
        _ = v66
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v7
    _ = v8
    _ = v9
}
