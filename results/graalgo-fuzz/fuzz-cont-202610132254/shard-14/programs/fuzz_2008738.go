package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 string; F2 uint32 }
type GenS6 struct { F0 int; F1 string }

func main() {
    var v1 int8 = 100
    var v2 int64 = -9223372036854775808
    var v3 int = 42
    var v4 uint8 = 63
    fmt.Println(v1, v1, v3)
    v5 := []int{}
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 18446744073709551615}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 2}
    if (105 > 32767) {
        switch (("\n\t" + "a") + (("" + "foo bar") + "hello")) {
        case "hello":
            _ = len(v5)
        case "foo bar":
            if (int32(v4) < -324) {
                fmt.Println(v2)
                v1 = -128
                switch (v3 - 916) {
                case 1:
                    for fi9 := 0; fi9 < 4; fi9++ {
                        for fi10 := 0; fi10 < 20; fi10++ {
                            fmt.Println(fi10, v4, fi9)
                            v4 -= v4
                            _ = fi10
                        }
                        v11 := GenS1{F0: 2}
                        _ = fi9
                        _ = v11
                    }
                case -9223372036854775808:
                    for fi12 := 0; fi12 < 11; fi12++ {
                        fmt.Println(v3)
                        v4 = v4
                        _ = fi12
                    }
                case 32768:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        for fi14 := 0; fi14 < 9; fi14++ {
                            if false {
                                if (836 == 696) {
                                    var v15 int = fi13
                                    switch -2147483648 {
                                    case 2147483647:
                                        v4 -= 2
                                    case 127:
                                        if false {
                                            if (!true) {
                                                fmt.Println(v4, v4, fi14)
                                                var v16 float32 = (0.0 - 1.5)
                                                v3 = (v15 + (-(fi14)))
                                                if (42 > 4294967295) {
                                                    v17 := GenS2{F0: ((v1 * v1) > v1)}
                                                    switch "" {
                                                    case "\n\t":
                                                        fmt.Println(fi14, v1, v3)
                                                    case "":
                                                        fmt.Println(v16)
                                                    case "a":
                                                        if true {
                                                            v18 := GenS2{F0: ((18446744073709551615 <= 18446744073709551615) && (true && true))}
                                                            _ = v18
                                                        }
                                                    }
                                                    _ = v17
                                                }
                                                _ = v16
                                            } else {
                                                var v19 int16 = -32768
                                                var v20 int = fi14
                                                var v21 uint8 = v4
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                            }
                                            if false {
                                                v3 -= (v15 & v15)
                                                v5 = append(v5, -865)
                                            }
                                        } else {
                                            fmt.Println(v15, v4)
                                            if ((v2 | (-(v2))) <= 348) {
                                                var v22 uint64 = 4294967295
                                                var v23 uint8 = 1
                                                _ = v22
                                                _ = v23
                                            } else {
                                                if true {
                                                    if len(v5) > 0 {
                                                        v5[0] = (fi13 * (fi13 - fi13))
                                                    }
                                                    if (0.1 != -2.25) {
                                                        switch v15 {
                                                        case -1:
                                                            v4 += (v4 / (((v4 << ((4294967295) & 63))) | 1))
                                                        default:
                                                            var v24 float32 = ((1e10 / float32(-2.25)) * (-2.25 - 0.1))
                                                            _ = v24
                                                        }
                                                        fmt.Println(v4)
                                                    }
                                                } else {
                                                    v25 := GenS3{F0: 110, F1: (^(v4)), F2: -1}
                                                    if (((-(v2)) == v2) && false) {
                                                        if len(v5) > 0 {
                                                            v5[0] = -1
                                                        }
                                                        switch ("" + "\n\t") {
                                                        case "hello":
                                                            v2 += (v2 * v2)
                                                        default:
                                                            v4 -= v4
                                                        }
                                                    }
                                                    v1 = (v1 * 42)
                                                    _ = v25
                                                }
                                                v26 := GenS3{F0: 824, F1: v4, F2: 2147483647}
                                                _ = v26
                                            }
                                        }
                                    default:
                                        v4 += v4
                                    }
                                    switch v3 {
                                    case 127:
                                        v27 := GenS3{F0: 18446744073709551615, F1: (v4 >> ((0) & 63)), F2: 2147483647}
                                        _ = v27
                                    case 0:
                                        if (!true) {
                                            v1 += v1
                                        } else {
                                            fmt.Println(v3, v4)
                                        }
                                    }
                                    _ = v15
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        v1 = (v1 << ((255) & 63))
                                        _ = len(v5)
                                        var v28 uint32 = 123
                                        _ = v28
                                    case "":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case "a":
                                        v29 := GenS2{F0: ((true && ("a" < "a")) && (false && true))}
                                        _ = v29
                                    default:
                                        fmt.Println(fi13, v1, v2)
                                    }
                                    if true {
                                        switch "" {
                                        case "hello":
                                            v30 := GenS1{F0: 4294967295}
                                            _ = v30
                                        case "\n\t":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        v31 := GenS1{F0: 963}
                                        _ = v31
                                    }
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                fmt.Println(v1, v3, fi14)
                            } else {
                                var v32 float32 = (float32(0.0) * 3.14159)
                                v5 = append(v5, 127)
                                _ = v32
                            }
                            switch fi13 {
                            case 32768:
                                v5 = append(v5, 128)
                            case 128:
                                if len(v5) > 0 {
                                    v5[0] = -648
                                }
                            }
                            var v33 int32 = -272
                            _ = fi14
                            _ = v33
                        }
                        fmt.Println(v3, v3, v2)
                        fmt.Println(v1)
                        _ = fi13
                    }
                default:
                    for fi34 := 0; fi34 < 20; fi34++ {
                        v4 += 42
                        if len(v5) > 0 {
                            v5[0] = (^(fi34))
                        }
                        _ = fi34
                    }
                }
            } else {
                var v35 int8 = 6
                _ = v35
            }
        }
        var v36 uint8 = 159
        _ = v36
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        for fi38 := 0; fi38 < 6; fi38++ {
            switch v3 {
            case -544:
                v1 = (v1 ^ -1)
            default:
                switch "hello" {
                case "":
                    v39 := GenS4{F0: 707}
                    _ = v39
                case "a":
                    switch "a" {
                    case "hello":
                        v3 += fi38
                    }
                case "hello":
                    var v40 string = ""
                    _ = v40
                }
            }
            var v41 int = (fi37 / ((v3) | 1))
            _ = fi38
            _ = v41
        }
        if (0.0 != 0.0) {
            v2 -= int64(v4)
            if (((^(v1)) > 42) && (!true)) {
                v5 = append(v5, -2147483648)
            } else {
                v4 = (v4 << ((18446744073709551615) & 63))
            }
            var v42 float64 = 1.5
            _ = v42
        } else {
            if len(v5) > 0 {
                v5[0] = -2147483648
            }
            v1 = (v1 / (((v1 * 79)) | 1))
            if ((true || (597 == 395)) && true) {
                switch v3 {
                case 265:
                    switch "hello" {
                    case "":
                        var v43 uint64 = 399
                        _ = v43
                    case "hello":
                        if len(v5) > 0 {
                            v5[0] = (v3 % ((fi37) | 1))
                        }
                    }
                default:
                    for fi44 := 0; fi44 < 3; fi44++ {
                        switch (fi44 & (fi37 + (fi37 - int(v2)))) {
                        case 1:
                            v3 += (^(fi44))
                        case 791:
                            if true {
                                _ = len(v5)
                                var v45 uint32 = 0
                                _ = v45
                            }
                        }
                        switch fi44 {
                        case 36:
                            var v46 uint32 = 65535
                            v47 := GenS1{F0: 42}
                            _ = v46
                            _ = v47
                        case 415:
                            var v48 bool = false
                            _ = v48
                        }
                        v2 += v2
                        if true {
                            v49 := GenS3{F0: 4294967295, F1: v4, F2: 2147483647}
                            _ = v49
                        }
                        var v50 string = "foo bar"
                        var v51 int32 = -2147483648
                        var v52 uint16 = 1
                        var v53 float64 = 0.0
                        v2 = 42
                        v3 += 303
                        v54 := GenS4{F0: v52}
                        _ = fi44
                        _ = v50
                        _ = v51
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                }
                v4 -= (^(v4))
                var v55 string = "hello"
                _ = v55
            }
        }
        _ = fi37
    }
    var v56 uint64 = 0
    var v57 float32 = (0.0 / 0.1)
    if (1 < v4) {
        var v58 uint8 = (v4 - (^(v4)))
        _ = v58
    } else {
        switch "hello" {
        case "":
            var v59 uint8 = v4
            _ = v59
        case "hello":
            v1 = (v1 - -1)
        }
    }
    var v60 uint32 = 2147483647
    v3 = -630
    if (v4 != (v4 & 255)) {
        v61 := GenS4{F0: 2}
        v60 -= (v60 / ((v60) | 1))
        fmt.Println(v2, v4)
        _ = v61
    } else {
        if (592 < (v3 / ((v3) | 1))) {
            for fi62 := 0; fi62 < 12; fi62++ {
                _ = len(v5)
                v5 = append(v5, int(v2))
                v1 -= int8(v4)
                _ = fi62
            }
            v63 := GenS4{F0: 42}
            _ = v63
        } else {
            v56 += v56
        }
    }
    for fi64 := 0; fi64 < 3; fi64++ {
        v65 := GenS5{F0: 18446744073709551615, F1: ("" + ("foo bar" + "")), F2: (v60 + v60)}
        _ = fi64
        _ = v65
    }
    var v66 int8 = (v1 | v1)
    if (false && true) {
        var v67 uint8 = (v4 ^ v4)
        _ = v67
    } else {
        if true {
            if ((v2 % (((v2 / ((v2) | 1))) | 1)) >= int64(v66)) {
                var v68 uint64 = 42
                v69 := GenS6{F0: v3, F1: (("" + ("\n\t" + "a")) + "")}
                v4 = (v4 << ((42) & 63))
                _ = v68
                _ = v69
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v60:", v60)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v60
    _ = v66
    _ = v8
}
