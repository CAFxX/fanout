package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 uint32 = 324
    v2 := []int{-259, 32768}
    v3 := make(map[int]bool)
    var v4 uint64 = 0
    v5 := make(map[uint32]string)
    v6 := []string{"hello", ""}
    var v7 uint = 42
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int8 = 1
        _ = fi8
        _ = v9
    }
    var v10 float32 = 0.1
    if ((-985 > 338) && true) {
        switch "\n\t" {
        case "":
            v11 := GenS1{F0: 42, F1: 2147483647, F2: 127}
            _ = v11
        case "foo bar":
            v12 := GenS1{F0: 255, F1: -1, F2: 100}
            _ = v12
        case "hello":
            if ("a" != "a") {
                v13 := GenS2{F0: v4, F1: 255, F2: (^(v7))}
                v1 = (v1 * (v1 & (v1 % (((v1 * v1)) | 1))))
                for fi14 := 0; fi14 < 7; fi14++ {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 float64 = ((3.14159 - float64(1.5)) / 1.0)
                        v4 -= v4
                        _ = fi15
                        _ = v16
                    }
                    if (-1 > 9223372036854775807) {
                        var v17 int = 32768
                        _ = v17
                    }
                    v4 -= (v4 & 355)
                    _ = fi14
                }
                _ = v13
            } else {
                var v18 int16 = 42
                _ = v18
            }
        }
    }
    switch "" {
    case "foo bar":
        if true {
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
        } else {
            v19 := GenS3{F0: 32768, F1: 32767}
            v10 = float32(-2.25)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v19
        }
    case "":
        switch "" {
        case "hello":
            v20 := GenS4{F0: ((2 <= (v4 >> (((v7 + v7)) & 63))) && false), F1: 0}
            _ = v20
        case "":
            if (-563 > -683) {
                v21 := GenS4{F0: (641 < 42), F1: 196}
                v22 := GenS1{F0: 695, F1: -1, F2: int8(v1)}
                _ = v21
                _ = v22
            } else {
                switch -609 {
                case 32768:
                    var v23 uint8 = 1
                    _ = v23
                case -9223372036854775808:
                    if (1 < 36) {
                        var v24 uint32 = (v1 | v1)
                        if true {
                            v1 -= uint32(v4)
                            fmt.Println(v1)
                            var v25 uint = v7
                            _ = v25
                        } else {
                            switch "" {
                            case "\n\t":
                                fmt.Println(v1, v7, v4)
                            case "hello":
                                v4 += v4
                            default:
                                var v26 uint32 = 160
                                _ = v26
                            }
                        }
                        if true {
                            if true {
                                v10 = v10
                                var v27 uint32 = (^(v24))
                                _ = v27
                            }
                            var v28 int64 = 648
                            var v29 float32 = (v10 / v10)
                            _ = v28
                            _ = v29
                        }
                        _ = v24
                    } else {
                        if (v4 <= (v4 >> ((0) & 63))) {
                            if false {
                                var v30 uint32 = (v1 / (((v1 | 2147483647)) | 1))
                                _ = v30
                            } else {
                                fmt.Println(v1, v7, v7)
                            }
                        } else {
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v10, v7, v7)
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                            default:
                                v10 += ((float32(float64(v10)) * (v10 + (v10 * float32((0.1 - ((-2.25 - float64(v10)) / 1.0)))))) + (((1e10 + v10) - 0.1) - 1.5))
                            }
                        }
                        if true {
                            switch "\n\t" {
                            case "foo bar":
                                var v31 int32 = -1
                                _ = v31
                            case "a":
                                switch 713 {
                                case 1:
                                    v10 = (float32(1e10) + v10)
                                case 9223372036854775807:
                                    for fi32 := 0; fi32 < 9; fi32++ {
                                        fmt.Println(v7, v1)
                                        _ = fi32
                                    }
                                case 2147483647:
                                    for fi33 := 0; fi33 < 11; fi33++ {
                                        fmt.Println(fi33)
                                        v34 := GenS5{F0: -2147483648, F1: 1}
                                        _ = fi33
                                        _ = v34
                                    }
                                }
                            default:
                                if true {
                                    var v35 bool = (!(32767 == 2))
                                    _ = v35
                                } else {
                                    v36 := GenS6{F0: 41}
                                    _ = v36
                                }
                            }
                            switch 388 {
                            case 4294967295:
                                v7 += uint(v1)
                            case 128:
                                var v37 uint = uint(v1)
                                _ = v37
                            }
                            v38 := GenS7{F0: 794}
                            _ = v38
                        } else {
                            var v39 uint8 = 60
                            _ = v39
                        }
                        v3[283] = (v7 == (v7 * (v7 ^ 42)))
                    }
                case -1:
                    v1 += (^(v1))
                }
            }
        case "foo bar":
            _ = len(v5)
        }
    }
    switch -1 {
    case 32768:
        fmt.Println(v7, v10, v10)
    case 2147483647:
        fmt.Println(v10, v4)
    case -1:
        _ = len(v6)
    default:
        for fi40 := 0; fi40 < 5; fi40++ {
            fmt.Println(v7)
            var v41 bool = true
            fmt.Println(v7)
            _ = fi40
            _ = v41
        }
    }
    switch "\n\t" {
    case "\n\t":
        var v42 uint64 = 2
        _ = v42
    case "hello":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "a":
        v43 := v3[42]
        _ = v43
    default:
        if ((688 >= 32767) || (127 > 9)) {
            switch 1 {
            case 32767:
                var v44 uint32 = (v1 ^ 0)
                _ = v44
            case 1:
                var v45 uint32 = (v1 | v1)
                _ = v45
            case 528:
                v46 := GenS5{F0: int32(v7), F1: -2147483648}
                _ = v46
            default:
                var v47 int64 = 9223372036854775807
                _ = v47
            }
            v4 = v4
            fmt.Println(v1, v4)
        } else {
            _ = len(v2)
            v7 += (^(v7))
            v1 = 70
        }
    }
    v4 = v4
    v48 := GenS1{F0: 2, F1: 2147483647, F2: 42}
    var v49 int = -689
    if (v4 > v4) {
        var v50 float32 = (v10 * v10)
        var v51 uint64 = (v4 % ((1) | 1))
        _ = v50
        _ = v51
    } else {
        if (false && (v7 > 1)) {
            if (false || true) {
                var v52 uint16 = 0
                fmt.Println(v52, v4, v7)
                fmt.Println(v52, v7, v49)
                _ = v52
            }
            v5[2147483648] = ""
            var v53 bool = false
            _ = v53
        }
        if (v4 <= (v4 % ((v4) | 1))) {
            for fi54 := 0; fi54 < 7; fi54++ {
                var v55 int64 = 9223372036854775807
                _ = fi54
                _ = v55
            }
            if true {
                v4 = (v4 >> ((v7) & 63))
            }
            if (v10 < v10) {
                v10 = 1e10
                v4 += uint64(v7)
                v1 = 324
            }
        } else {
            switch ("a" + (("a" + ("\n\t" + (("" + "\n\t") + ("" + (("hello" + ("\n\t" + "hello")) + ""))))) + ("" + ("foo bar" + (("" + "\n\t") + "foo bar"))))) {
            case "\n\t":
                for fi56 := 0; fi56 < 9; fi56++ {
                    switch "hello" {
                    case "a":
                        v10 = (((v10 * 0.1) + 0.1) - (1e10 + 1e10))
                    case "\n\t":
                        var v57 uint32 = v1
                        _ = v57
                    }
                    fmt.Println(v1, v1, v49)
                    _ = fi56
                }
            case "foo bar":
                v49 -= (v49 & 834)
            }
        }
    }
    for fi58 := 0; fi58 < 13; fi58++ {
        v10 += (v10 / v10)
        _ = fi58
    }
    v7 += (v7 | 2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
