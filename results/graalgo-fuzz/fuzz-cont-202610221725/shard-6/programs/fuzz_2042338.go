package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint64; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int; F2 uint16 }
type GenS6 struct { F0 int32; F1 int16; F2 uint8 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int16 }
func genFn0(p0 int64, p1 uint, p2 uint) int {
    if false {
        if (!true) {
            v5 := []int{106, -9223372036854775808, 28, 32768}
            fmt.Println(p0)
            _ = v5
        } else {
            var v6 int16 = 0
            var v7 float32 = (((0.1 / 1e10) * -2.25) / ((0.1 / ((0.1 + (3.14159 / float32(0.1))) + 3.14159)) / ((1e10 / float32((0.0 + 3.14159))) + (0.1 / 0.1))))
            p1 = 42
            _ = v6
            _ = v7
        }
        v8 := []int{}
        for fi9 := 0; fi9 < 12; fi9++ {
            switch (fi9 / ((fi9) | 1)) {
            case 42:
                for fi10 := 0; fi10 < 9; fi10++ {
                    fmt.Println(fi9)
                    fmt.Println(fi9, p0, fi10)
                    var v11 uint8 = 128
                    _ = fi10
                    _ = v11
                }
            default:
                switch "hello" {
                case "\n\t":
                    for fi12 := 0; fi12 < 1; fi12++ {
                        var v13 string = "\n\t"
                        var v14 uint = 18446744073709551615
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case "":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                case "hello":
                    var v15 uint16 = 1
                    _ = v15
                default:
                    if (false && ((^(p2)) < p1)) {
                        v16 := GenS1{F0: (!false), F1: uint16(p2)}
                        p0 = 1
                        _ = v16
                    }
                }
            }
            _ = fi9
        }
        _ = v8
    } else {
        if true {
            fmt.Println(p2, p2, p1)
            v17 := []int{-2147483648}
            if true {
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 string = ("hello" + (((("a" + "") + (("" + (((("foo bar" + (("" + (("foo bar" + "hello") + "a")) + "foo bar")) + "hello") + (("" + ("\n\t" + ("\n\t" + ("a" + "")))) + "foo bar")) + ("" + (("" + ((("a" + "") + "\n\t") + "foo bar")) + "a")))) + "\n\t")) + "foo bar") + (("foo bar" + "a") + "foo bar")))
                    _ = fi18
                    _ = v19
                }
                var v20 int32 = -2147483648
                _ = v20
            } else {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                p2 -= 1
                p2 += 168
            }
            _ = v17
        }
        fmt.Println(p1, p0, p0)
    }
    if true {
        var v21 int32 = -458
        var v22 float64 = float64(0.0)
        p0 -= (p0 + p0)
        _ = v21
        _ = v22
    } else {
        p2 += p1
    }
    v23 := []int{-182}
    if true {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        v24 := GenS2{F0: 9223372036854775807, F1: 4294967295, F2: (((((3.14159 * (1.5 + 1.5)) * 0.1) / ((1.5 / -2.25) - (1e10 - 1.5))) + 0.1) >= (((-2.25 * ((1.5 + 3.14159) - (0.1 - 0.0))) * 3.14159) + ((((float64(1e10) * 1.5) * 1.5) / (((((-2.25 - -2.25) + 1e10) / -2.25) - 1.5) - (1e10 / 0.1))) / 1.0)))}
        _ = v24
    }
    if (2 <= 1) {
        p0 = (p0 | (p0 << ((p2) & 63)))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 537
}

func main() {
    var v1 uint16 = 800
    v2 := make(map[int]string)
    var v3 float32 = 1.5
    v4 := make(map[string]bool)
    v25 := []int{4294967295, -9223372036854775808, 42}
    v1 += 32767
    if false {
        switch "foo bar" {
        case "a":
            var v26 float64 = 0.0
            _ = v26
        case "\n\t":
            for fi27 := 0; fi27 < 14; fi27++ {
                if (genFn0(-307, 41, 869) < int(v1)) {
                    switch (fi27 + fi27) {
                    case 32767:
                        v28 := GenS3{F0: 124}
                        _ = v28
                    case 2147483647:
                        if (42 <= 2) {
                            for fi29 := 0; fi29 < 11; fi29++ {
                                v30 := GenS2{F0: 618, F1: 2, F2: true}
                                _ = fi29
                                _ = v30
                            }
                            v3 = v3
                        }
                    default:
                        v31 := GenS4{F0: 1, F1: (^(v1))}
                        _ = v31
                    }
                    v32 := GenS1{F0: false, F1: v1}
                    var v33 uint64 = 1
                    _ = v32
                    _ = v33
                } else {
                    if len(v25) > 0 {
                        v25[0] = 4294967295
                    }
                }
                fmt.Println(fi27, v1, fi27)
                _ = fi27
            }
        default:
            var v34 int32 = -944
            _ = v34
        }
    }
    if true {
        _ = len(v25)
        switch -248 {
        case 2147483647:
            for fi35 := 0; fi35 < 9; fi35++ {
                _ = len(v4)
                switch "hello" {
                case "a":
                    var v36 uint64 = 4294967295
                    _ = v36
                case "\n\t":
                    if false {
                        switch "foo bar" {
                        case "\n\t":
                            v37 := GenS1{F0: false, F1: 255}
                            _ = v37
                        case "a":
                            if len(v25) > 0 {
                                v25[0] = (fi35 << ((42) & 63))
                            }
                        default:
                            v3 += (v3 * v3)
                        }
                        switch fi35 {
                        case 1:
                            if false {
                                fmt.Println(fi35, v1)
                                if (112 <= 42) {
                                    switch (fi35 % (((fi35 << ((uint(v1)) & 63))) | 1)) {
                                    case 128:
                                        v3 = float32(0.1)
                                    default:
                                        for fi38 := 0; fi38 < 16; fi38++ {
                                            var v39 uint = 42
                                            _ = fi38
                                            _ = v39
                                        }
                                    }
                                } else {
                                    for fi40 := 0; fi40 < 19; fi40++ {
                                        if (false && true) {
                                            v3 += v3
                                        } else {
                                            var v41 int16 = -32768
                                            _ = v41
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v42 int64 = 42
                                            _ = v42
                                        case "":
                                            v43 := GenS4{F0: -9223372036854775808, F1: (v1 + v1)}
                                            _ = v43
                                        }
                                        _ = fi40
                                    }
                                    var v44 float64 = (3.14159 - 0.1)
                                    if true {
                                        fmt.Println(v3, fi35)
                                        v44 += float64(float32(v44))
                                        switch "" {
                                        case "hello":
                                            switch (^(fi35)) {
                                            case 154:
                                                var v45 int16 = -709
                                                _ = v45
                                            case 418:
                                                if false {
                                                    v46 := GenS5{F0: 42, F1: (fi35 - (-(fi35))), F2: (v1 >> ((18446744073709551615) & 63))}
                                                    var v47 string = "a"
                                                    for fi48 := 0; fi48 < 5; fi48++ {
                                                        if true {
                                                            fmt.Println(fi48)
                                                            v2[128] = "\n\t"
                                                        } else {
                                                            var v49 int8 = 42
                                                            _ = v49
                                                        }
                                                        _ = fi48
                                                    }
                                                    _ = v46
                                                    _ = v47
                                                }
                                            }
                                        case "a":
                                            v44 -= (v44 / v44)
                                        case "":
                                            v44 += (-2.25 * (v44 * 1e10))
                                        default:
                                            switch "\n\t" {
                                            case "hello":
                                                var v50 uint8 = 62
                                                _ = v50
                                            case "a":
                                                v3 = (-2.25 * v3)
                                            case "\n\t":
                                                v3 = 0.0
                                            default:
                                                v3 += v3
                                            }
                                        }
                                    } else {
                                        if (true && false) {
                                            v3 -= (-2.25 + (0.0 - (1.5 * (((v3 + v3) - v3) + -2.25))))
                                            if (-74 < 24) {
                                                if len(v25) >= 2 {
                                                    v25 = v25[1:len(v25)]
                                                }
                                                if (v1 == (^(v1))) {
                                                    for fi51 := 0; fi51 < 17; fi51++ {
                                                        v52 := GenS6{F0: 758, F1: -1, F2: 42}
                                                        v53 := GenS7{F0: 0, F1: (false || true)}
                                                        fmt.Println(v44, v3, fi51)
                                                        _ = fi51
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    switch (("\n\t" + ((("" + ("\n\t" + "")) + "foo bar") + ("foo bar" + "a"))) + ("\n\t" + ("" + ("" + "hello")))) {
                                                    case "\n\t":
                                                        v54 := GenS8{F0: 690}
                                                        _ = v54
                                                    }
                                                } else {
                                                    _ = len(v25)
                                                    v55 := GenS8{F0: 831}
                                                    _ = v55
                                                }
                                            } else {
                                                v3 = ((3.14159 / 3.14159) / (v3 / (0.1 / float32(v44))))
                                            }
                                            if (("foo bar" + "\n\t") == "") {
                                                var v56 uint = 11
                                                _ = v56
                                            } else {
                                                v3 = float32((0.0 - v44))
                                                v57 := GenS6{F0: -1, F1: 1, F2: 42}
                                                var v58 uint16 = v1
                                                _ = v57
                                                _ = v58
                                            }
                                        }
                                    }
                                    _ = v44
                                }
                                var v59 int = fi35
                                _ = v59
                            }
                        case 588:
                            v25 = append(v25, (fi35 << ((1) & 63)))
                        case 32768:
                            switch 1 {
                            case 128:
                                v3 = v3
                            }
                        }
                        var v60 uint32 = 2147483647
                        _ = v60
                    } else {
                        if (!(22 <= 1)) {
                            v61 := GenS9{F0: 191, F1: 1, F2: -1}
                            v1 = (^(v1))
                            fmt.Println(fi35)
                            _ = v61
                        }
                        fmt.Println(v3)
                        v62 := v2[32767]
                        _ = v62
                    }
                case "hello":
                    delete(v4, "")
                }
                _ = fi35
            }
        case 979:
            var v63 int16 = -262
            _ = v63
        default:
            var v64 float64 = 3.14159
            _ = v64
        }
    }
    var v65 int32 = -616
    var v66 int32 = v65
    fmt.Println(v66)
    switch genFn0(42, 65535, 411) {
    case 9223372036854775807:
        if false {
            fmt.Println(v66, v3)
        }
    case 4294967295:
        for fi67 := 0; fi67 < 10; fi67++ {
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            var v68 uint32 = 523
            if false {
                delete(v2, (fi67 + (-(fi67))))
                v69 := GenS10{F0: 1}
                for fi70 := 0; fi70 < 11; fi70++ {
                    switch (fi67 ^ -2147483648) {
                    case 334:
                        v1 = v1
                        v3 -= 1.5
                        v71 := GenS11{F0: 11}
                        _ = v71
                    case 32767:
                        v25 = append(v25, (^(fi67)))
                        var v72 int8 = -1
                        _ = v72
                    default:
                        var v73 int8 = -91
                        _ = v73
                    }
                    if (!true) {
                        fmt.Println(v68, v68, v68)
                        v74 := GenS9{F0: 1, F1: 1, F2: -50}
                        _ = v74
                    } else {
                        v75 := GenS5{F0: v68, F1: (fi70 | fi70), F2: (v1 >> ((868) & 63))}
                        fmt.Println(v1)
                        var v76 int32 = v65
                        _ = v75
                        _ = v76
                    }
                    _ = fi70
                }
                _ = v69
            } else {
                var v77 bool = true
                _ = v77
            }
            _ = fi67
            _ = v68
        }
    case 503:
        v66 = (-(v66))
    default:
        switch 9223372036854775807 {
        case -1:
            var v78 int8 = 127
            _ = v78
        case 4294967295:
            var v79 int64 = -9223372036854775808
            _ = v79
        }
    }
    fmt.Println(v1, v65)
    if true {
        v3 = v3
        switch "\n\t" {
        case "foo bar":
            v25 = append(v25, 4294967295)
        case "a":
            v3 = v3
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v65
    _ = v66
}
