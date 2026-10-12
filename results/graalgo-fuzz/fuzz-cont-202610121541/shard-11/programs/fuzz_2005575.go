package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 uint8; F1 string; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int; F2 int8 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int8; F1 uint16; F2 uint }
type GenS7 struct { F0 uint16; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-970, 32768}
    v3 := make(map[uint32]uint32)
    var v4 int16 = 32767
    var v5 int8 = 100
    switch -146 {
    case 1:
        var v6 int = -9223372036854775808
        _ = v6
    case 32768:
        fmt.Println(v5)
    case 128:
        for fi7 := 0; fi7 < 18; fi7++ {
            delete(v3, 42)
            _ = fi7
        }
    }
    if ((false || false) && true) {
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 int64 = 1
                var v11 float64 = 3.14159
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        switch 333 {
        case -1:
            switch 4294967295 {
            case 2147483647:
                var v12 string = ""
                _ = v12
            case -769:
                fmt.Println(v4)
            case 32768:
                var v13 uint = 18446744073709551615
                _ = v13
            }
        case -9223372036854775808:
            v4 = 42
        case -829:
            var v14 int16 = (v4 >> ((65535) & 63))
            _ = v14
        }
    } else {
        fmt.Println(v4, v5)
        if true {
            v3[2] = 476
            switch "a" {
            case "":
                fmt.Println(v5, v4)
            case "foo bar":
                var v15 float32 = 1.5
                _ = v15
            case "\n\t":
                v4 = v4
            }
            var v16 uint64 = 613
            _ = v16
        } else {
            var v17 int8 = v5
            fmt.Println(v5)
            v18 := GenS1{F0: 65535, F1: "", F2: -895}
            _ = v17
            _ = v18
        }
        var v19 uint16 = 2
        _ = v19
    }
    v20 := GenS2{F0: 214, F1: "hello", F2: 2}
    var v21 float32 = 3.14159
    if true {
        switch ((("\n\t" + (("hello" + ("hello" + (("" + ("hello" + ("a" + "hello"))) + "hello"))) + ("\n\t" + "\n\t"))) + "foo bar") + "a") {
        case "hello":
            fmt.Println(v4, v5)
        }
        if (("\n\t" + "") >= "foo bar") {
            if true {
                switch int(v5) {
                case 708:
                    var v22 uint64 = 260
                    _ = v22
                case 42:
                    var v23 uint = 507
                    _ = v23
                default:
                    switch ("a" + "") {
                    case "a":
                        v2 = append(v2, 673)
                    case "hello":
                        v2 = append(v2, 1)
                    case "foo bar":
                        var v24 float64 = ((1e10 / 0.1) + 0.1)
                        _ = v24
                    default:
                        for fi25 := 0; fi25 < 9; fi25++ {
                            if (84 == 197) {
                                v5 = v5
                                var v26 uint32 = 494
                                for fi27 := 0; fi27 < 17; fi27++ {
                                    v21 -= ((v21 - (v21 - (v21 / v21))) * v21)
                                    fmt.Println(fi25, v26, fi25)
                                    _ = fi27
                                }
                                _ = v26
                            } else {
                                var v28 uint16 = 0
                                var v29 int32 = 42
                                var v30 float32 = v21
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            var v31 float32 = v21
                            _ = fi25
                            _ = v31
                        }
                    }
                }
                var v32 int16 = -1
                fmt.Println(v5)
                _ = v32
            } else {
                for fi33 := 0; fi33 < 20; fi33++ {
                    v4 += v4
                    _ = fi33
                }
            }
            v4 = v4
        } else {
            if false {
                if (true && (true || true)) {
                    var v34 uint = 942
                    _ = v34
                }
            }
        }
    }
    _ = len(v2)
    v5 -= 42
    v35 := GenS3{F0: 4294967295}
    if true {
        v21 += v21
        v36 := GenS4{F0: 128, F1: 2147483647, F2: -1}
        v37 := GenS5{F0: true, F1: -482}
        _ = v36
        _ = v37
    }
    if true {
        switch 188 {
        case 127:
            for fi38 := 0; fi38 < 10; fi38++ {
                v39 := GenS6{F0: v5, F1: 42, F2: 2}
                _ = fi38
                _ = v39
            }
        case -1:
            v40 := GenS3{F0: 65535}
            _ = v40
        case -9223372036854775808:
            if true {
                v21 += ((0.0 * (v21 / ((v21 + ((float32(1e10) + v21) - (v21 + (float32(3.14159) + 3.14159)))) - (v21 - (v21 - v21))))) - v21)
            } else {
                v4 = v4
                var v41 uint64 = 876
                v41 += (v41 >> ((4294967295) & 63))
                _ = v41
            }
        }
        switch "hello" {
        case "a":
            var v42 int = -1
            _ = v42
        case "\n\t":
            v21 = v21
        case "hello":
            _ = len(v1)
        default:
            v5 = v5
        }
    }
    var v43 uint = 23
    if false {
        switch 42 {
        case 501:
            fmt.Println(v5, v43, v43)
        case -453:
            v44 := v3[4294967295]
            _ = v44
        }
        if true {
            fmt.Println(v43)
            v21 -= v21
            if false {
                v45 := GenS6{F0: v5, F1: 32767, F2: 2}
                _ = v45
            } else {
                v21 -= v21
                var v46 uint32 = 146
                if ((!(181 != 0)) && (false && (((float64(((0.0 - ((v21 * 0.0) + v21)) + 0.1)) + 0.1) * 0.0) != 0.0))) {
                    if (1 >= 42) {
                        v3[v46] = (v46 & v46)
                    }
                    switch "" {
                    case "":
                        if true {
                            fmt.Println(v46, v4)
                            for fi47 := 0; fi47 < 2; fi47++ {
                                v5 = (-(v5))
                                _ = fi47
                            }
                            v48 := GenS2{F0: 42, F1: "hello", F2: uint8(v43)}
                            _ = v48
                        }
                    case "hello":
                        fmt.Println(v4, v5)
                    }
                }
                _ = v46
            }
        } else {
            for fi49 := 0; fi49 < 7; fi49++ {
                if (225 >= 0) {
                    fmt.Println(v43, v5)
                    var v50 int32 = -450
                    var v51 float64 = 0.1
                    _ = v50
                    _ = v51
                }
                v52 := GenS1{F0: 836, F1: "a", F2: 1}
                _ = fi49
                _ = v52
            }
            if false {
                v53 := GenS7{F0: 642, F1: 289}
                var v54 int64 = int64(v4)
                var v55 float32 = v21
                _ = v53
                _ = v54
                _ = v55
            } else {
                v56 := GenS4{F0: 32768, F1: 42, F2: (v5 | v5)}
                switch (("" + (("foo bar" + "hello") + ((("foo bar" + "a") + "foo bar") + "\n\t"))) + "hello") {
                case "":
                    fmt.Println(v21, v21, v43)
                default:
                    for fi57 := 0; fi57 < 10; fi57++ {
                        switch "\n\t" {
                        case "a":
                            if (!(-2.25 == 0.0)) {
                                var v58 uint64 = 462
                                _ = v58
                            }
                        }
                        _ = fi57
                    }
                }
                _ = v56
            }
            switch 215 {
            case 9223372036854775807:
                if false {
                    v4 += v4
                    fmt.Println(v21, v21, v5)
                } else {
                    var v59 int16 = (v4 << ((v43) & 63))
                    if false {
                        if len(v2) > 0 {
                            v2[0] = 127
                        }
                    } else {
                        for fi60 := 0; fi60 < 7; fi60++ {
                            _ = len(v3)
                            if (-2.25 < 1.5) {
                                if (float64(float32(1e10)) == 0.0) {
                                    if (false || (uint32(v43) >= 65535)) {
                                        for fi61 := 0; fi61 < 6; fi61++ {
                                            if ((^(v5)) > (v5 / ((v5) | 1))) {
                                                v2 = append(v2, 0)
                                                v4 = (v4 << (((v43 / (((v43 >> (((v43 / (((v43 * v43)) | 1))) & 63))) | 1))) & 63))
                                                fmt.Println(fi60)
                                            } else {
                                                var v62 int16 = v4
                                                v21 -= (0.0 + (v21 - v21))
                                                _ = v62
                                            }
                                            v2 = append(v2, (fi60 << ((v43) & 63)))
                                            v43 -= v43
                                            _ = fi61
                                        }
                                        for fi63 := 0; fi63 < 12; fi63++ {
                                            var v64 uint = (v43 | (v43 * v43))
                                            _ = fi63
                                            _ = v64
                                        }
                                        fmt.Println(v43)
                                    }
                                    for fi65 := 0; fi65 < 18; fi65++ {
                                        fmt.Println(v5, fi60)
                                        var v66 int8 = int8(v43)
                                        v4 = (-(v4))
                                        _ = fi65
                                        _ = v66
                                    }
                                    switch ("a" + "") {
                                    case "a":
                                        v21 = float32(0.0)
                                    case "foo bar":
                                        v5 -= v5
                                    case "hello":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        v4 -= (v4 >> (((v43 % ((v43) | 1))) & 63))
                                    }
                                }
                                fmt.Println(v21, v43)
                            }
                            _ = fi60
                        }
                    }
                    _ = v59
                }
            case 128:
                for fi67 := 0; fi67 < 7; fi67++ {
                    var v68 uint16 = 1
                    fmt.Println(v5, v5)
                    _ = fi67
                    _ = v68
                }
            case -2147483648:
                var v69 uint32 = 0
                _ = v69
            default:
                if true {
                    if ((v43 >= (v43 | (v43 ^ v43))) || true) {
                        v1[2147483647] = ("\n\t" + "hello")
                        v4 += v4
                    }
                    switch "\n\t" {
                    case "a":
                        v70 := GenS6{F0: v5, F1: 65535, F2: (v43 & 222)}
                        _ = v70
                    case "foo bar":
                        v21 += v21
                    case "hello":
                        if (!true) {
                            v21 = (v21 - v21)
                            var v71 int16 = (v4 << (((v43 ^ (v43 ^ (v43 | v43)))) & 63))
                            _ = v71
                        }
                    }
                } else {
                    if ((true || true) || true) {
                        v72 := GenS6{F0: (^(v5)), F1: 1, F2: (v43 ^ (v43 * 255))}
                        _ = v72
                    } else {
                        var v73 uint = (^(v43))
                        _ = v73
                    }
                    var v74 int16 = v4
                    _ = v74
                }
            }
        }
    }
    if false {
        var v75 float32 = (v21 * v21)
        _ = v75
    }
    var v76 uint32 = 85
    var v77 int64 = 1
    fmt.Println(v77)
    var v78 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v5
    _ = v76
    _ = v77
    _ = v78
}
