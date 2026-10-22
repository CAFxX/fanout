package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int16 }
type GenS2 struct { F0 int16; F1 int64; F2 int16 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 uint8; F1 int; F2 uint }
type GenS9 struct { F0 uint32; F1 uint64; F2 uint8 }

func main() {
    v1 := []uint32{214, 42, 0}
    v2 := make(map[string]int)
    v3 := []int{0, 9223372036854775807, 4294967295, -1}
    v4 := make(map[uint32]uint32)
    var v5 int32 = -2147483648
    v5 -= (v5 + v5)
    for fi6 := 0; fi6 < 1; fi6++ {
        if true {
            fmt.Println(v5, fi6)
        }
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int64 = 42
            var v9 int8 = -120
            _ = fi7
            _ = v8
            _ = v9
        }
        v5 -= v5
        _ = fi6
    }
    if (true || false) {
        var v10 int8 = 90
        for fi11 := 0; fi11 < 8; fi11++ {
            v5 = v5
            var v12 uint16 = 42
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    if (!(1 < 664)) {
        fmt.Println(v5)
        var v13 int64 = 0
        _ = v13
    } else {
        var v14 int32 = v5
        for fi15 := 0; fi15 < 15; fi15++ {
            fmt.Println(fi15, fi15, v14)
            if false {
                if (true || true) {
                    v14 = v5
                    var v16 int32 = 42
                    _ = v16
                }
            }
            var v17 uint = 255
            _ = fi15
            _ = v17
        }
        v18 := v2["\n\t"]
        _ = v14
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            if true {
                v21 := v4[142]
                _ = v21
            } else {
                if true {
                    v5 += v5
                }
                v5 = v5
            }
            _ = fi20
        }
        switch "a" {
        case "a":
            if (568 != 32767) {
                if (((float32(0.0) / 1.0) == ((1e10 + 0.1) - (1e10 / 1.0))) || false) {
                    fmt.Println(v5)
                    switch 127 {
                    case 170:
                        var v22 int16 = -137
                        _ = v22
                    case 127:
                        var v23 int16 = 42
                        _ = v23
                    case 1:
                        fmt.Println(fi19, fi19)
                    }
                }
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                switch "" {
                case "foo bar":
                    for fi24 := 0; fi24 < 20; fi24++ {
                        if false {
                            if false {
                                v5 -= 1
                                v5 = v5
                                v5 -= 2147483647
                            }
                            v5 = (v5 - v5)
                        } else {
                            if false {
                                var v25 uint8 = 255
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(fi19, v5, fi19)
                                default:
                                    if false {
                                        var v26 uint = 4294967295
                                        _ = v26
                                    }
                                }
                                v5 = 93
                                if false {
                                    var v27 int16 = -35
                                    v27 += (v27 | v27)
                                    var v28 string = ""
                                    var v29 int32 = v5
                                    var v30 float64 = 0.0
                                    v28 = ""
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                } else {
                                    delete(v4, 4294967295)
                                    v1 = append(v1, 42)
                                }
                                _ = v25
                            }
                            fmt.Println(fi24, fi24)
                            switch fi19 {
                            case 129:
                                v5 = (v5 | -2147483648)
                            case 4294967295:
                                if true {
                                    v5 = (v5 ^ (v5 >> ((1) & 63)))
                                } else {
                                    v5 -= (v5 ^ -2147483648)
                                    var v31 int8 = -128
                                    var v32 uint = 65535
                                    _ = v31
                                    _ = v32
                                }
                            default:
                                v5 = v5
                            }
                        }
                        _ = fi24
                    }
                default:
                    for fi33 := 0; fi33 < 4; fi33++ {
                        switch ("foo bar" + "\n\t") {
                        case "":
                            fmt.Println(v5, v5)
                        }
                        _ = fi33
                    }
                }
            }
        case "hello":
            for fi34 := 0; fi34 < 17; fi34++ {
                v5 -= v5
                var v35 int8 = 0
                if (2 != 2) {
                    v35 = 1
                } else {
                    _ = len(v3)
                }
                _ = fi34
                _ = v35
            }
        case "\n\t":
            switch "" {
            case "hello":
                if len(v1) > 0 {
                    v1[0] = 452
                }
            case "a":
                v5 = (v5 + v5)
            case "\n\t":
                v5 = (v5 % ((v5) | 1))
            }
        default:
            v36 := GenS1{F0: -9223372036854775808, F1: 255, F2: -513}
            _ = v36
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi19
    }
    _ = len(v1)
    var v37 float32 = 1.5
    switch 32768 {
    case 42:
        var v38 int = 4294967295
        _ = v38
    case 644:
        if false {
            switch "\n\t" {
            case "":
                v39 := GenS2{F0: -168, F1: 42, F2: -32768}
                _ = v39
            default:
                if true {
                    v40 := GenS2{F0: 32767, F1: 988, F2: -994}
                    _ = v40
                }
            }
        } else {
            fmt.Println(v5)
        }
    }
    v41 := GenS2{F0: -347, F1: 720, F2: 21}
    if (float32(((float64(((1e10 - float32((-2.25 / 1.0))) / (v37 * v37))) - 0.1) / 1.5)) < ((1e10 * -2.25) + v37)) {
        v42 := GenS3{F0: 4294967295, F1: 2147483647, F2: 1}
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        if (2 < 65535) {
            var v43 uint16 = 1
            _ = v43
        } else {
            var v44 bool = ((-9223372036854775808 < 1) && (v5 > -1))
            var v45 float64 = 0.0
            _ = v44
            _ = v45
        }
        _ = v42
    }
    var v46 int64 = 9223372036854775807
    switch -1 {
    case 32768:
        v37 = v37
    case 447:
        if ((-128 <= 127) || ((102 == 0) || false)) {
            v46 += 478
            fmt.Println(v37, v46)
            var v47 uint = 18446744073709551615
            _ = v47
        } else {
            fmt.Println(v46)
            for fi48 := 0; fi48 < 7; fi48++ {
                v37 = v37
                fmt.Println(v37, fi48)
                if (-2.25 == v37) {
                    for fi49 := 0; fi49 < 3; fi49++ {
                        var v50 int32 = 1
                        switch "a" {
                        case "":
                            fmt.Println(v50, fi48)
                        case "a":
                            _ = len(v1)
                        case "\n\t":
                            if (2 > 127) {
                                switch "\n\t" {
                                case "hello":
                                    if (((("foo bar" + ("\n\t" + ("foo bar" + ("foo bar" + "")))) + "foo bar") + "\n\t") >= "a") {
                                        v51 := GenS3{F0: 0, F1: 955, F2: 65535}
                                        _ = v51
                                    } else {
                                        switch fi49 {
                                        case 760:
                                            if true {
                                                var v52 uint32 = 2
                                                _ = v52
                                            }
                                        case 9223372036854775807:
                                            v5 = -518
                                        }
                                        var v53 int = fi48
                                        if false {
                                            var v54 float32 = ((0.1 / ((v37 + v37) * v37)) * v37)
                                            if (v53 > fi49) {
                                                v55 := GenS2{F0: -832, F1: (^(v46)), F2: 311}
                                                _ = v55
                                            }
                                            fmt.Println(v46, v37, v54)
                                            _ = len(v1)
                                            _ = v54
                                        }
                                        var v56 int32 = (v5 / ((42) | 1))
                                        _ = v53
                                        _ = v56
                                    }
                                case "":
                                    v50 = v5
                                default:
                                    fmt.Println(v46, v5, fi48)
                                }
                                v3 = append(v3, (fi48 + fi49))
                            }
                        default:
                            var v57 float32 = (1.5 + ((((float32(-2.25) + v37) - (v37 * v37)) / v37) * ((v37 * (v37 - ((1.5 * v37) + (1.5 * 0.0)))) - 0.0)))
                            _ = v57
                        }
                        v50 -= v5
                        var v58 int8 = 30
                        if (fi48 == -355) {
                            v59 := GenS4{F0: v58, F1: 255}
                            switch "foo bar" {
                            case "foo bar":
                                v5 += (v5 * v5)
                            case "\n\t":
                                switch (^(fi49)) {
                                case 42:
                                    fmt.Println(v46)
                                    v60 := GenS5{F0: 1, F1: (fi48 - (fi48 - fi48))}
                                    _ = v60
                                case 127:
                                    if (("foo bar" + ("foo bar" + "a")) > "\n\t") {
                                        var v61 int8 = v58
                                        _ = v61
                                    } else {
                                        if true {
                                            v62 := GenS6{F0: (v58 * int8(fi48)), F1: 255}
                                            switch (fi48 | (fi48 ^ (fi49 & (fi49 % ((-936) | 1))))) {
                                            case 1:
                                                fmt.Println(v5, fi48, v58)
                                            }
                                            v4[4294967295] = 0
                                            _ = v62
                                        }
                                        v46 = 42
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = fi48
                                    }
                                }
                            default:
                                v63 := GenS6{F0: (v58 >> ((18446744073709551615) & 63)), F1: 754}
                                _ = v63
                            }
                            fmt.Println(v5)
                            var v64 float64 = 0.1
                            v64 = (float64(3.14159) + (v64 + 1e10))
                            switch "a" {
                            case "\n\t":
                                _ = len(v4)
                            case "foo bar":
                                var v65 uint8 = 51
                                _ = v65
                            case "hello":
                                if (2147483647 > (fi49 << ((65535) & 63))) {
                                    fmt.Println(v5)
                                }
                            default:
                                v58 = (v58 | 0)
                            }
                            _ = v59
                            _ = v64
                        } else {
                            _ = len(v3)
                            var v66 uint64 = 0
                            _ = v66
                        }
                        _ = fi49
                        _ = v50
                        _ = v58
                    }
                    v67 := GenS7{F0: (-(v46)), F1: v46}
                    v68 := GenS6{F0: 42, F1: 2}
                    _ = v67
                    _ = v68
                } else {
                    if false {
                        v69 := GenS8{F0: 255, F1: 0, F2: 321}
                        var v70 uint8 = 0
                        for fi71 := 0; fi71 < 5; fi71++ {
                            var v72 int32 = v5
                            var v73 int = 32767
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                        _ = v69
                        _ = v70
                    } else {
                        var v74 uint64 = 206
                        _ = len(v3)
                        v37 -= (v37 * 1.5)
                        _ = v74
                    }
                }
                _ = fi48
            }
            fmt.Println(v5)
        }
    case 127:
        for fi75 := 0; fi75 < 18; fi75++ {
            v5 += (v5 + v5)
            _ = fi75
        }
    default:
        fmt.Println(v5)
    }
    v46 = -585
    if ((0.1 + float32(-2.25)) <= v37) {
        v5 = v5
        for fi76 := 0; fi76 < 15; fi76++ {
            var v77 int8 = -1
            v37 += (((v37 * ((v37 * v37) * v37)) * float32(1e10)) - (v37 + (v37 + float32(((1e10 * (-2.25 * float64(0.0))) / float64(float32((1.5 / 3.14159))))))))
            _ = fi76
            _ = v77
        }
        fmt.Println(v46, v5, v37)
    }
    _ = len(v3)
    v78 := GenS9{F0: 599, F1: 0, F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v46
    _ = v5
    _ = v78
}
