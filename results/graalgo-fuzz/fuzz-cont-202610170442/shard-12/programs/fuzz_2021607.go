package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int64; F2 bool }
type GenS4 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint; F2 int32 }
type GenS6 struct { F0 int16; F1 uint16 }

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{552}
    v3 := []uint32{1, 4294967295, 2147483647}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 := GenS1{F0: -1, F1: 97, F2: 735}
    switch "a" {
    case "":
        var v5 float64 = (0.1 / v1)
        _ = v5
    case "foo bar":
        v6 := GenS2{F0: 0}
        _ = v6
    case "a":
        if false {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            if (1 >= 0) {
                if ((-389 == 1) || (!(!false))) {
                    var v7 uint32 = 966
                    for fi8 := 0; fi8 < 13; fi8++ {
                        var v9 int8 = 127
                        var v10 uint16 = 111
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    _ = v7
                }
            }
        } else {
            var v11 int64 = -9223372036854775808
            var v12 uint32 = 842
            if len(v3) > 0 {
                v3[0] = v12
            }
            _ = v11
            _ = v12
        }
    default:
        if (-1 >= 42) {
            var v13 uint32 = 42
            v1 -= 3.14159
            _ = v13
        } else {
            v1 -= (0.1 * v1)
            if (0 >= 2) {
                v1 -= v1
                if (!((65535 == 4) && false)) {
                    var v14 bool = ((!true) && (!true))
                    for fi15 := 0; fi15 < 3; fi15++ {
                        _ = len(v2)
                        for fi16 := 0; fi16 < 15; fi16++ {
                            if (!v14) {
                                _ = len(v3)
                                v14 = false
                                if v14 {
                                    var v17 uint64 = 2
                                    switch "" {
                                    case "":
                                        var v18 bool = v14
                                        _ = v18
                                    case "a":
                                        var v19 bool = (true && v14)
                                        _ = v19
                                    case "\n\t":
                                        switch (fi16 % ((fi16) | 1)) {
                                        case -1:
                                            if (352 <= 65535) {
                                                if (true && (false && ((v14 || v14) || (4294967295 < 4294967295)))) {
                                                    var v20 int16 = 1000
                                                    var v21 int8 = -128
                                                    fmt.Println(v17)
                                                    var v22 int64 = 42
                                                    if v14 {
                                                        var v23 uint64 = (v17 / (((v17 & v17)) | 1))
                                                        _ = v23
                                                    }
                                                    _ = v20
                                                    _ = v21
                                                    _ = v22
                                                } else {
                                                    v24 := GenS3{F0: 152, F1: -311, F2: (127 == 1)}
                                                    var v25 uint16 = 65535
                                                    if (1 < 42) {
                                                        if (v25 > 1) {
                                                            var v26 string = "hello"
                                                            if v14 {
                                                                v27 := GenS1{F0: (fi15 >> ((255) & 63)), F1: 42, F2: (v25 / ((v25) | 1))}
                                                                _ = v27
                                                            }
                                                            v2 = append(v2, 0)
                                                            if false {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            fmt.Println(v14)
                                                            _ = v26
                                                        } else {
                                                            var v28 int64 = 0
                                                            _ = v28
                                                        }
                                                        v2 = append(v2, 2)
                                                        v14 = false
                                                    }
                                                    fmt.Println(fi15, fi15, fi16)
                                                    switch (fi15 << ((42) & 63)) {
                                                    case 0:
                                                        switch (fi15 + 32767) {
                                                        case 128:
                                                            v29 := GenS4{F0: 2, F1: 368, F2: 4294967295}
                                                            _ = v29
                                                        case 32767:
                                                            switch -2147483648 {
                                                            case -248:
                                                                var v30 int32 = 570
                                                                _ = v30
                                                            case 32767:
                                                                v14 = (18446744073709551615 < 455)
                                                            case 42:
                                                                v31 := GenS3{F0: 24, F1: -9223372036854775808, F2: v14}
                                                                _ = v31
                                                            default:
                                                                var v32 int = fi16
                                                                _ = v32
                                                            }
                                                        case -454:
                                                            var v33 float64 = float64((3.14159 + (-2.25 / ((1e10 / float32(1.5)) + 3.14159))))
                                                            _ = v33
                                                        }
                                                    case -514:
                                                        var v34 string = ("\n\t" + ("" + "foo bar"))
                                                        _ = v34
                                                    }
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v35 float64 = float64(3.14159)
                                                        _ = v35
                                                    case "hello":
                                                        v36 := GenS5{F0: -718, F1: 255, F2: -1}
                                                        _ = v36
                                                    case "":
                                                        v1 = v1
                                                    }
                                                    switch ((((("\n\t" + "\n\t") + ("" + ("foo bar" + ("foo bar" + "foo bar")))) + "") + "hello") + (("a" + "foo bar") + "a")) {
                                                    case "\n\t":
                                                        var v37 uint16 = v25
                                                        _ = v37
                                                    case "a":
                                                        var v38 uint = 0
                                                        _ = v38
                                                    case "hello":
                                                        v14 = (1 != v25)
                                                    default:
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v1 = ((1e10 / v1) * float64(((1.5 / (1e10 * ((float32(v1) + -2.25) - -2.25))) * 0.0)))
                                                        case "":
                                                            switch "foo bar" {
                                                            case "":
                                                                fmt.Println(fi16)
                                                            case "\n\t":
                                                                v14 = true
                                                                if (!((255 <= 2) && v14)) {
                                                                    var v39 uint64 = (v17 >> ((18446744073709551615) & 63))
                                                                    _ = v39
                                                                }
                                                                var v40 uint16 = v25
                                                                _ = v40
                                                            default:
                                                                if (!((!(!(-1 > 967))) || v14)) {
                                                                    if true {
                                                                        fmt.Println(fi15)
                                                                        var v41 uint64 = 490
                                                                        _ = v41
                                                                    } else {
                                                                        if (42 == -2147483648) {
                                                                            if v14 {
                                                                                var v42 uint64 = (^(v17))
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                _ = len(v3)
                                                                                _ = v42
                                                                            }
                                                                            if (1 <= -128) {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v17 += v17
                                                                            if ((0.0 * (0.1 / 1e10)) != 1e10) {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            } else {
                                                                                v43 := GenS5{F0: -1, F1: 74, F2: -666}
                                                                                _ = v43
                                                                            }
                                                                            v2 = append(v2, 2147483648)
                                                                        }
                                                                        if v14 {
                                                                            var v44 int8 = 42
                                                                            var v45 uint64 = v17
                                                                            _ = v44
                                                                            _ = v45
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            var v46 int64 = -1
                                                            _ = v46
                                                        }
                                                    }
                                                    _ = v24
                                                    _ = v25
                                                }
                                                fmt.Println(v17, fi15)
                                                v47 := GenS6{F0: 0, F1: 42}
                                                _ = v47
                                            } else {
                                                v14 = false
                                                switch 127 {
                                                case 989:
                                                    var v48 float64 = v1
                                                    v1 -= float64(0.0)
                                                    _ = v48
                                                case 929:
                                                    var v49 uint32 = 1
                                                    var v50 int8 = -128
                                                    _ = v49
                                                    _ = v50
                                                case -160:
                                                    var v51 int64 = 0
                                                    _ = v51
                                                }
                                                v52 := GenS3{F0: 127, F1: -9223372036854775808, F2: (v14 || v14)}
                                                _ = v52
                                            }
                                        case 880:
                                            v53 := GenS3{F0: 0, F1: 42, F2: false}
                                            _ = v53
                                        default:
                                            v17 = v17
                                        }
                                    default:
                                        v3 = append(v3, 2147483647)
                                    }
                                    _ = v17
                                }
                            }
                            v54 := GenS1{F0: 0, F1: -74, F2: 32767}
                            v1 -= ((v1 - (v1 + 3.14159)) + (1e10 * v1))
                            _ = fi16
                            _ = v54
                        }
                        _ = fi15
                    }
                    v1 += v1
                    _ = v14
                } else {
                    v1 = 1e10
                    switch "hello" {
                    case "hello":
                        v2 = append(v2, 2147483648)
                    }
                    var v55 int16 = 17
                    _ = v55
                }
            }
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi56 := 0; fi56 < 1; fi56++ {
        if (706 >= 0) {
            var v57 uint64 = 18446744073709551615
            var v58 uint = 0
            _ = v57
            _ = v58
        } else {
            if len(v3) > 0 {
                v3[0] = 340
            }
        }
        _ = fi56
    }
    for fi59 := 0; fi59 < 8; fi59++ {
        var v60 int16 = -1
        switch (-(fi59)) {
        case 42:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        switch "" {
        case "foo bar":
            var v61 float64 = float64(float32((v1 / (v1 + v1))))
            _ = v61
        case "\n\t":
            v60 = (^(v60))
        default:
            var v62 int32 = 2147483647
            _ = v62
        }
        _ = fi59
        _ = v60
    }
    for fi63 := 0; fi63 < 15; fi63++ {
        v1 = ((1e10 + ((v1 - ((v1 - 3.14159) / 1.0)) - 0.0)) / v1)
        for fi64 := 0; fi64 < 19; fi64++ {
            if false {
                var v65 uint16 = 874
                _ = v65
            }
            _ = fi64
        }
        _ = fi63
    }
    v1 = v1
    v1 -= (0.1 * float64(0.0))
    switch "a" {
    case "\n\t":
        fmt.Println(v1)
    case "a":
        if false {
            for fi66 := 0; fi66 < 16; fi66++ {
                v1 = 3.14159
                switch 32768 {
                case -195:
                    switch "a" {
                    case "foo bar":
                        for fi67 := 0; fi67 < 1; fi67++ {
                            var v68 uint16 = 1
                            _ = fi67
                            _ = v68
                        }
                    }
                case -517:
                    for fi69 := 0; fi69 < 10; fi69++ {
                        switch fi69 {
                        case -1:
                            v1 = v1
                            v70 := GenS4{F0: 565, F1: 198, F2: 1}
                            var v71 float64 = (float64(0.1) - v1)
                            _ = v70
                            _ = v71
                        case 0:
                            switch "\n\t" {
                            case "a":
                                _ = len(v3)
                            default:
                                if false {
                                    fmt.Println(fi69, fi66, v1)
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                }
                            }
                        default:
                            if (!false) {
                                switch (fi69 * (fi69 | (fi66 >> ((521) & 63)))) {
                                case 9223372036854775807:
                                    v1 -= 0.0
                                case 32768:
                                    fmt.Println(fi69, v1)
                                }
                                v3 = append(v3, 42)
                            }
                            var v72 int32 = 1
                            _ = v72
                        }
                        v1 = v1
                        _ = fi69
                    }
                default:
                    var v73 int64 = -9223372036854775808
                    _ = v73
                }
                var v74 uint32 = 652
                _ = fi66
                _ = v74
            }
            v75 := GenS5{F0: 485, F1: 0, F2: 2147483647}
            _ = v75
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
