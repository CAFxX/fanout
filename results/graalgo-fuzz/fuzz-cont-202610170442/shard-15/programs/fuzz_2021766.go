package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS7 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS8 struct { F0 int8; F1 uint32; F2 bool }
func genFn0(p0 bool, p1 int) int8 {
    fmt.Println(p0)
    if true {
        fmt.Println(p0)
    } else {
        v7 := GenS1{F0: 104, F1: 65535, F2: -470}
        switch (^(p1)) {
        case 0:
            v8 := GenS2{F0: 444, F1: 36, F2: -2147483648}
            _ = v8
        case 32768:
            if false {
                v9 := GenS2{F0: 2, F1: 127, F2: -2147483648}
                fmt.Println(p1, p1)
                _ = v9
            }
        default:
            switch "foo bar" {
            case "\n\t":
                if ("\n\t" == "a") {
                    if true {
                        v10 := GenS3{F0: 65535, F1: 1}
                        var v11 float64 = float64((0.1 + (1e10 - 1e10)))
                        _ = v10
                        _ = v11
                    } else {
                        var v12 float32 = 3.14159
                        p0 = false
                        _ = v12
                    }
                    if (!true) {
                        if (((("\n\t" + "hello") + "\n\t") >= "\n\t") && (p0 || (p0 && p0))) {
                            v13 := []int{(p1 * (p1 * 128)), p1, (-(p1)), p1}
                            p1 += (p1 - p1)
                            v14 := GenS3{F0: 18446744073709551615, F1: 255}
                            _ = v13
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 19; fi15++ {
                            v16 := GenS2{F0: 2, F1: 128, F2: 2147483647}
                            if (42 != 65535) {
                                var v17 uint8 = 208
                                switch ("" + (("hello" + "a") + ("a" + "foo bar"))) {
                                case "a":
                                    v17 -= v17
                                default:
                                    var v18 float32 = (1e10 + 1.5)
                                    _ = v18
                                }
                                p0 = (p0 || (((v17 | 127) != 1) && (!p0)))
                                _ = v17
                            } else {
                                fmt.Println(fi15, fi15)
                                switch "" {
                                case "hello":
                                    var v19 int64 = 1
                                    _ = v19
                                default:
                                    if p0 {
                                        p1 -= p1
                                        p1 = fi15
                                        v20 := []int{}
                                        _ = v20
                                    } else {
                                        if (0 >= 127) {
                                            switch "a" {
                                            case "hello":
                                                for fi21 := 0; fi21 < 7; fi21++ {
                                                    if (100 < 46) {
                                                        p1 = 999
                                                        v22 := []int{2147483647, 127, (^(fi21))}
                                                        _ = v22
                                                    }
                                                    _ = fi21
                                                }
                                            default:
                                                switch "" {
                                                case "hello":
                                                    var v23 int8 = 42
                                                    _ = v23
                                                case "foo bar":
                                                    if (1 != 700) {
                                                        switch 844 {
                                                        case 9223372036854775807:
                                                            fmt.Println(fi15, p0)
                                                        }
                                                        var v24 int = 1
                                                        v25 := GenS4{F0: -2147483648, F1: 1}
                                                        _ = v24
                                                        _ = v25
                                                    } else {
                                                        v26 := []int{fi15, p1, (p1 ^ p1)}
                                                        switch p1 {
                                                        case -719:
                                                            if len(v26) > 0 {
                                                                v26[0] = (p1 % ((fi15) | 1))
                                                            }
                                                        default:
                                                            var v27 float64 = 0.0
                                                            _ = v27
                                                        }
                                                        _ = v26
                                                    }
                                                default:
                                                    p0 = ((("hello" + "\n\t") + ("hello" + (("foo bar" + (("foo bar" + "") + ("" + "foo bar"))) + "a"))) >= ((("a" + "a") + ("hello" + ("" + "\n\t"))) + "a"))
                                                }
                                            }
                                        } else {
                                            for fi28 := 0; fi28 < 7; fi28++ {
                                                switch ("\n\t" + ("a" + "foo bar")) {
                                                case "hello":
                                                    v29 := []int{127}
                                                    _ = v29
                                                case "foo bar":
                                                    if false {
                                                        var v30 int = 737
                                                        switch 42 {
                                                        case -1:
                                                            var v31 uint8 = 100
                                                            _ = v31
                                                        case 2147483647:
                                                            p1 += (fi15 >> ((110) & 63))
                                                        case 32767:
                                                            p1 += (p1 >> ((593) & 63))
                                                        default:
                                                            var v32 uint16 = 2
                                                            _ = v32
                                                        }
                                                        _ = v30
                                                    } else {
                                                        fmt.Println(fi15, fi28, fi15)
                                                        v33 := []int{-297, fi28, (fi15 - 42), (p1 << ((2) & 63))}
                                                        _ = v33
                                                    }
                                                default:
                                                    p0 = p0
                                                }
                                                var v34 float64 = (1.5 + (1e10 - (0.1 - ((0.0 * (float64(float32(-2.25)) - (-2.25 + -2.25))) / 0.1))))
                                                v34 -= (v34 - v34)
                                                _ = fi28
                                                _ = v34
                                            }
                                            v35 := GenS4{F0: 0, F1: 0}
                                            _ = v35
                                        }
                                        p0 = p0
                                    }
                                }
                                var v36 uint64 = 1
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 1; fi37++ {
                                var v38 float64 = 0.0
                                fmt.Println(fi37, v38)
                                _ = fi37
                                _ = v38
                            }
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        v39 := GenS1{F0: 445, F1: 760, F2: -515}
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 12; fi40++ {
                        var v41 uint16 = 42
                        _ = fi40
                        _ = v41
                    }
                }
            default:
                p1 -= 128
            }
        }
        var v42 string = "a"
        _ = v42
        _ = v7
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        var v44 uint8 = 1
        var v45 uint64 = 2
        _ = fi43
        _ = v44
        _ = v45
    }
    _ = p0
    _ = p1
    return 100
}

func main() {
    var v1 int = -870
    var v2 uint64 = 434
    var v3 int32 = -193
    var v4 uint32 = 1
    v5 := []uint32{(v4 / ((0) | 1)), 539, v4}
    var v6 int16 = 549
    for fi46 := 0; fi46 < 14; fi46++ {
        v1 = (fi46 * (v1 - 1))
        var v47 int16 = 1
        _ = fi46
        _ = v47
    }
    if (!false) {
        fmt.Println(v4, v6)
        if ((v6 % ((v6) | 1)) > -226) {
            var v48 int32 = (-(v3))
            switch (v1 >> ((4294967295) & 63)) {
            case -2147483648:
                var v49 int16 = 1
                _ = v49
            case 32767:
                v1 += v1
            case 127:
                if (!(!("" <= "\n\t"))) {
                    v4 -= (v4 + 1)
                    var v50 int8 = 1
                    _ = v50
                } else {
                    fmt.Println(v3, v1, v1)
                    var v51 uint32 = v4
                    _ = v51
                }
            }
            _ = v48
        } else {
            v52 := GenS5{F0: -9223372036854775808, F1: v1}
            var v53 uint8 = 128
            _ = v52
            _ = v53
        }
    }
    v54 := GenS1{F0: 0, F1: 42, F2: 0}
    var v55 int32 = int32(v4)
    v56 := GenS6{F0: 128, F1: 1, F2: v4}
    switch (v1 | (v1 - 32768)) {
    case -2147483648:
        var v57 int64 = -9223372036854775808
        _ = v57
    case 464:
        switch v1 {
        case -319:
            for fi58 := 0; fi58 < 2; fi58++ {
                v2 -= v2
                for fi59 := 0; fi59 < 7; fi59++ {
                    var v60 uint32 = (v4 | (v4 * (v4 >> ((1) & 63))))
                    fmt.Println(v4)
                    v4 += 412
                    _ = fi59
                    _ = v60
                }
                if true {
                    var v61 bool = true
                    fmt.Println(v4, v61)
                    _ = v61
                } else {
                    if (!true) {
                        switch "\n\t" {
                        case "foo bar":
                            v55 = (v3 - v55)
                        case "\n\t":
                            _ = len(v5)
                        case "a":
                            fmt.Println(v2)
                        default:
                            if true {
                                var v62 bool = (0 >= 255)
                                _ = v62
                            }
                        }
                        for fi63 := 0; fi63 < 9; fi63++ {
                            v55 = v55
                            _ = fi63
                        }
                        fmt.Println(v2)
                    } else {
                        var v64 int = fi58
                        var v65 uint16 = 846
                        switch "\n\t" {
                        case "hello":
                            var v66 uint = 914
                            _ = v66
                        default:
                            var v67 int16 = 1
                            _ = v67
                        }
                        _ = v64
                        _ = v65
                    }
                    var v68 bool = ((1.5 / 1.0) != 3.14159)
                    v68 = (221 > 384)
                    _ = v68
                }
                _ = fi58
            }
        case -9223372036854775808:
            fmt.Println(v4)
        case 1:
            switch v1 {
            case -786:
                for fi69 := 0; fi69 < 4; fi69++ {
                    v55 = v55
                    var v70 bool = ((v1 + v1) != (fi69 + fi69))
                    _ = fi69
                    _ = v70
                }
            case 2147483647:
                v71 := GenS7{F0: 128, F1: (v55 - (v55 - 0)), F2: (v2 * (v2 | v2))}
                _ = v71
            }
        }
    case 9223372036854775807:
        var v72 int32 = (v55 * (^(v55)))
        _ = v72
    }
    v2 -= v2
    if true {
        for fi73 := 0; fi73 < 1; fi73++ {
            if true {
                for fi74 := 0; fi74 < 3; fi74++ {
                    var v75 uint64 = 265
                    v76 := GenS7{F0: 166, F1: 1, F2: v2}
                    v3 += v3
                    _ = fi74
                    _ = v75
                    _ = v76
                }
                switch "hello" {
                case "hello":
                    var v77 int16 = (v6 / ((-1) | 1))
                    _ = v77
                case "a":
                    switch (fi73 << ((65535) & 63)) {
                    case 8:
                        var v78 int = (v1 | -9223372036854775808)
                        _ = v78
                    case -495:
                        v79 := GenS3{F0: 733, F1: 0}
                        _ = v79
                    }
                }
            }
            _ = len(v5)
            v4 -= v4
            _ = fi73
        }
    } else {
        var v80 int8 = 127
        _ = v80
    }
    v3 = -1
    v81 := GenS8{F0: 0, F1: v4, F2: false}
    for fi82 := 0; fi82 < 5; fi82++ {
        var v83 uint32 = v4
        var v84 uint8 = 128
        v83 = (v83 ^ 146)
        _ = fi82
        _ = v83
        _ = v84
    }
    switch v1 {
    case 42:
        switch "\n\t" {
        case "":
            var v85 uint = 7
            _ = v85
        }
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v6
    _ = v81
}
