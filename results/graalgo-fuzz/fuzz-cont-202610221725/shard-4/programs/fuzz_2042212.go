package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int32; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 uint8; F1 uint; F2 int64 }
type GenS8 struct { F0 uint64; F1 bool; F2 int16 }
func genFn0(p0 uint) int32 {
    v4 := []int{int(p0)}
    var v5 bool = true
    var v6 int32 = 2147483647
    if false {
        if len(v4) > 0 {
            v4[0] = -937
        }
        var v7 uint64 = 810
        switch ((("a" + ("hello" + (("\n\t" + "hello") + "a"))) + "\n\t") + ("hello" + "foo bar")) {
        case "\n\t":
            v7 = v7
        default:
            fmt.Println(v7, p0, v7)
        }
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        if v5 {
            fmt.Println(fi8)
        } else {
            var v9 int8 = -1
            _ = v9
        }
        switch "foo bar" {
        case "foo bar":
            fmt.Println(p0, p0, v5)
        case "\n\t":
            if (v5 || v5) {
                if ((!v5) || true) {
                    fmt.Println(v5, v5)
                }
            } else {
                if v5 {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi10
                    }
                    v11 := GenS1{F0: -1, F1: (fi8 | fi8)}
                    _ = v11
                } else {
                    v12 := GenS2{F0: 599}
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v14 := GenS1{F0: fi13, F1: 0}
                        if ((fi13 | (fi13 % ((32768) | 1))) == 42) {
                            v4 = append(v4, 127)
                        } else {
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                            var v15 int8 = -41
                            v16 := GenS3{F0: v5}
                            if (29 == 128) {
                                v6 += (v6 ^ (v6 >> ((p0) & 63)))
                                v6 += -1
                                var v17 int8 = 0
                                _ = v17
                            }
                            v18 := GenS4{F0: -2147483648, F1: v6, F2: 38}
                            v6 -= v6
                            v5 = (fi13 != (fi8 % ((fi8) | 1)))
                            _ = v15
                            _ = v16
                            _ = v18
                        }
                        v19 := GenS5{F0: 360}
                        _ = fi13
                        _ = v14
                        _ = v19
                    }
                    _ = v12
                }
                v6 = v6
            }
        case "":
            fmt.Println(fi8, v5)
        default:
            fmt.Println(p0)
        }
        _ = fi8
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return v6
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 255
    v3 := make(map[int]uint32)
    v20 := GenS4{F0: 0, F1: genFn0(18446744073709551615), F2: 1}
    for fi21 := 0; fi21 < 14; fi21++ {
        switch (fi21 * 32768) {
        case 42:
            switch 128 {
            case 42:
                var v22 string = "foo bar"
                _ = v22
            case 216:
                switch "" {
                case "foo bar":
                    v2 = 42
                case "\n\t":
                    for fi23 := 0; fi23 < 7; fi23++ {
                        switch "" {
                        case "a":
                            var v24 uint32 = 65535
                            _ = v24
                        case "hello":
                            v25 := GenS4{F0: (fi23 >> ((0) & 63)), F1: -1, F2: 42}
                            _ = v25
                        case "foo bar":
                            v2 += v2
                        }
                        var v26 uint64 = 413
                        v26 += (^(v26))
                        v26 -= v26
                        v27 := GenS2{F0: (v26 * 42)}
                        _ = fi23
                        _ = v26
                        _ = v27
                    }
                }
            default:
                var v28 uint32 = 2147483648
                _ = v28
            }
        }
        if false {
            if false {
                for fi29 := 0; fi29 < 2; fi29++ {
                    if false {
                        fmt.Println(fi29)
                        v30 := []int{}
                        _ = v30
                    }
                    _ = fi29
                }
                v2 += v2
            } else {
                switch "a" {
                case "a":
                    var v31 bool = false
                    _ = v31
                case "foo bar":
                    if true {
                        for fi32 := 0; fi32 < 8; fi32++ {
                            var v33 uint64 = 18446744073709551615
                            v34 := GenS2{F0: v33}
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    } else {
                        v35 := GenS3{F0: false}
                        _ = v35
                    }
                default:
                    switch (fi21 ^ 9223372036854775807) {
                    case 128:
                        v2 -= 2
                    default:
                        var v36 uint8 = (v2 - (v2 + v2))
                        _ = v36
                    }
                }
                if false {
                    var v37 int8 = -128
                    _ = v37
                } else {
                    switch "foo bar" {
                    case "a":
                        v38 := []int{(fi21 >> ((730) & 63)), -374, (fi21 << ((65535) & 63))}
                        _ = v38
                    default:
                        if (fi21 >= (-(fi21))) {
                            v2 = v2
                        } else {
                            switch 42 {
                            case -2147483648:
                                v2 = (v2 / (((^(v2))) | 1))
                            default:
                                var v39 string = ("foo bar" + ("hello" + ("\n\t" + "hello")))
                                _ = v39
                            }
                        }
                    }
                    if true {
                        for fi40 := 0; fi40 < 11; fi40++ {
                            v41 := GenS6{F0: -1, F1: 342}
                            if (false && ((-9223372036854775808 >= 258) && (577 <= 2147483647))) {
                                var v42 string = ""
                                _ = v42
                            }
                            _ = fi40
                            _ = v41
                        }
                    } else {
                        v2 = v2
                        for fi43 := 0; fi43 < 11; fi43++ {
                            var v44 uint64 = uint64(v2)
                            _ = fi43
                            _ = v44
                        }
                        switch ("a" + "\n\t") {
                        case "hello":
                            if true {
                                v2 = v2
                                for fi45 := 0; fi45 < 1; fi45++ {
                                    var v46 uint64 = 4294967295
                                    var v47 int8 = 42
                                    v48 := []int{}
                                    if true {
                                        v49 := v3[fi45]
                                        v49 = v49
                                        _ = v49
                                    } else {
                                        var v50 float32 = ((float32((0.1 - 1.5)) + 0.0) - 1e10)
                                        _ = v50
                                    }
                                    _ = fi45
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                                v51 := GenS1{F0: (fi21 ^ (-(fi21))), F1: (fi21 - (fi21 % ((fi21) | 1)))}
                                _ = v51
                            }
                        }
                    }
                    switch "" {
                    case "hello":
                        fmt.Println(fi21)
                    case "foo bar":
                        var v52 uint8 = v2
                        _ = v52
                    }
                }
            }
        }
        _ = fi21
    }
    fmt.Println(v2)
    switch "foo bar" {
    case "\n\t":
        switch "" {
        case "foo bar":
            v2 = v2
        case "a":
            var v53 uint = 1
            _ = v53
        case "\n\t":
            v2 -= (v2 % ((2) | 1))
        }
    default:
        switch "foo bar" {
        case "foo bar":
            v2 += (v2 ^ 2)
        case "\n\t":
            v2 -= (^(v2))
        case "":
            switch -1 {
            case 4294967295:
                if (!false) {
                    v2 = (^(v2))
                    var v54 string = ("hello" + ("foo bar" + ((("hello" + "hello") + "hello") + "")))
                    v54 = v54
                    _ = v54
                }
            case 661:
                var v55 int64 = 0
                _ = v55
            default:
                for fi56 := 0; fi56 < 11; fi56++ {
                    v2 += v2
                    _ = fi56
                }
            }
        }
    }
    v2 += (v2 & (^(v2)))
    fmt.Println(v2)
    v57 := GenS4{F0: 32767, F1: 0, F2: -71}
    fmt.Println(v2)
    v58 := GenS6{F0: -1, F1: 681}
    fmt.Println(v2)
    v2 = v2
    if false {
        for fi59 := 0; fi59 < 15; fi59++ {
            if (0.0 > (1e10 * -2.25)) {
                fmt.Println(v2, fi59)
                var v60 uint = 633
                var v61 float64 = 0.1
                _ = v60
                _ = v61
            } else {
                for fi62 := 0; fi62 < 13; fi62++ {
                    v63 := []int{}
                    var v64 int16 = 9
                    fmt.Println(fi62, fi62, fi59)
                    _ = fi62
                    _ = v63
                    _ = v64
                }
            }
            v2 = (v2 >> ((255) & 63))
            v2 = (^(v2))
            _ = fi59
        }
        if ("hello" < "foo bar") {
            v2 = 0
            v65 := GenS1{F0: -196, F1: 1}
            var v66 int8 = 115
            _ = v65
            _ = v66
        } else {
            var v67 uint32 = 380
            v68 := GenS7{F0: v2, F1: 42, F2: 455}
            _ = v67
            _ = v68
        }
    } else {
        v2 = v2
        for fi69 := 0; fi69 < 14; fi69++ {
            if (757 > 1) {
                switch 909 {
                case 388:
                    var v70 int = (fi69 + fi69)
                    _ = v70
                case 9223372036854775807:
                    v2 -= v2
                }
            } else {
                var v71 float64 = float64(1.5)
                v72 := GenS5{F0: 174}
                fmt.Println(fi69)
                _ = v71
                _ = v72
            }
            switch fi69 {
            case 0:
                if false {
                    v2 += (v2 | 127)
                    v2 = (v2 / ((v2) | 1))
                } else {
                    if true {
                        v73 := []int{}
                        _ = v73
                    }
                }
            default:
                fmt.Println(fi69, fi69)
            }
            var v74 int16 = -1
            _ = fi69
            _ = v74
        }
        var v75 int32 = -2147483648
        _ = v75
    }
    v76 := GenS8{F0: 42, F1: (!true), F2: 0}
    switch "" {
    case "a":
        var v77 int32 = -1
        _ = v77
    default:
        var v78 float32 = (((3.14159 / float32((-2.25 - 1e10))) + ((3.14159 + float32(((3.14159 * (0.1 * 0.1)) - (0.0 - 1e10)))) * float32(0.0))) - float32((1e10 / 3.14159)))
        _ = v78
    }
    if true {
        var v79 float32 = float32((float64((-2.25 - 3.14159)) - 0.1))
        var v80 float32 = ((0.0 * (v79 + v79)) / v79)
        v81 := GenS4{F0: 4294967295, F1: -1, F2: 0}
        _ = v79
        _ = v80
        _ = v81
    } else {
        fmt.Println(v2)
    }
    if ((!(32767 == 255)) || ((1.5 * -2.25) < float32(0.0))) {
        for fi82 := 0; fi82 < 1; fi82++ {
            var v83 int16 = -1
            var v84 uint8 = v2
            if false {
                delete(v3, fi82)
                v85 := []int{}
                _ = v85
            }
            _ = fi82
            _ = v83
            _ = v84
        }
        v2 += v2
        v2 = (v2 * 255)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v57
    _ = v58
    _ = v76
}
