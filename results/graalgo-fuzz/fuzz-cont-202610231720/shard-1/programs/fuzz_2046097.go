package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint32; F1 int32; F2 bool }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 uint64; F1 uint16; F2 string }
type GenS9 struct { F0 string; F1 uint; F2 uint8 }

func main() {
    var v1 int8 = -105
    var v2 string = "hello"
    var v3 int8 = v1
    var v4 uint32 = 2
    for fi5 := 0; fi5 < 18; fi5++ {
        if (!true) {
            fmt.Println(fi5)
        } else {
            v6 := GenS1{F0: 1, F1: (127 == (v3 | -40))}
            v7 := []int{2147483647, (fi5 * fi5)}
            var v8 uint32 = v4
            _ = v6
            _ = v7
            _ = v8
        }
        v3 = v1
        _ = fi5
    }
    v9 := []int{843}
    var v10 float64 = 0.1
    for fi11 := 0; fi11 < 7; fi11++ {
        if (112 == 42) {
            var v12 uint = 4294967295
            var v13 string = v2
            _ = v12
            _ = v13
        }
        v4 += v4
        v14 := GenS1{F0: 494, F1: false}
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        _ = len(v9)
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS1{F0: 716, F1: false}
            if true {
                var v18 string = ((v2 + ("hello" + (v2 + ("\n\t" + ("\n\t" + "hello"))))) + v2)
                fmt.Println(v3, fi15, fi16)
                _ = v18
            } else {
                var v19 int8 = 76
                fmt.Println(fi16, v3)
                _ = v19
            }
            fmt.Println(v3, v4, v2)
            switch "hello" {
            case "\n\t":
                var v20 uint = 42
                _ = v20
            }
            if true {
                fmt.Println(v2, v4, v3)
                var v21 float64 = (v10 + (v10 + 0.1))
                _ = v21
            }
            v4 -= 4294967295
            fmt.Println(v10, fi16, fi16)
            if len(v9) > 0 {
                v9[0] = fi16
            }
            _ = fi16
            _ = v17
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            v10 = v10
            switch (fi15 >> ((1) & 63)) {
            case 32767:
                if (!true) {
                    _ = len(v9)
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        if (!(458 == 563)) {
                            var v23 uint16 = 244
                            v3 = -128
                            _ = v23
                        }
                    case "hello":
                        var v24 uint = 494
                        _ = v24
                    case "foo bar":
                        var v25 string = "hello"
                        _ = v25
                    }
                }
            case 409:
                if ((v3 - -128) == -113) {
                    if len(v9) > 0 {
                        v9[0] = 1
                    }
                    v10 = 0.1
                    _ = len(v9)
                }
            case -1:
                switch 1 {
                case -2147483648:
                    switch (fi22 ^ 4294967295) {
                    case 127:
                        fmt.Println(v2, fi22, fi15)
                    case 4294967295:
                        var v26 uint16 = 255
                        _ = v26
                    case 2147483647:
                        if len(v9) > 0 {
                            v9[0] = (^(fi15))
                        }
                        switch fi22 {
                        case 858:
                            var v27 float32 = 1.5
                            _ = v27
                        case 4294967295:
                            v28 := GenS2{F0: 565, F1: 1}
                            _ = v28
                        }
                    default:
                        if false {
                            switch fi15 {
                            case 32768:
                                v3 = 44
                            case -1:
                                v4 += v4
                            case -247:
                                fmt.Println(v3, v3, v1)
                            default:
                                if (false && false) {
                                    _ = len(v9)
                                } else {
                                    if true {
                                        v4 = 82
                                        var v29 int64 = -9223372036854775808
                                        _ = v29
                                    }
                                    v30 := GenS1{F0: 4294967295, F1: true}
                                    _ = v30
                                }
                            }
                            v31 := GenS2{F0: 0, F1: 14}
                            _ = v31
                        } else {
                            v9 = append(v9, fi15)
                            v4 = 1
                            switch v2 {
                            case "foo bar":
                                fmt.Println(v4, fi15, v3)
                            case "hello":
                                v3 = v1
                            case "a":
                                switch v2 {
                                case "":
                                    v1 = (v3 | (v3 % ((0) | 1)))
                                case "\n\t":
                                    v3 = -1
                                }
                                var v32 uint16 = uint16(v4)
                                _ = v32
                            }
                        }
                    }
                default:
                    if false {
                        switch (fi22 << ((1) & 63)) {
                        case -2147483648:
                            if len(v9) > 0 {
                                v9[0] = -707
                            }
                        case 42:
                            v4 -= v4
                        case 1:
                            _ = len(v9)
                        }
                        switch (v2 + v2) {
                        case "":
                            var v33 uint32 = v4
                            _ = v33
                        case "foo bar":
                            v4 -= 482
                        case "hello":
                            v1 = (-(v3))
                        default:
                            fmt.Println(v2, fi22)
                        }
                        var v34 uint8 = 255
                        _ = v34
                    }
                    v3 = (v1 & (v3 * 127))
                    switch (-(fi22)) {
                    case 128:
                        switch 128 {
                        case 128:
                            _ = len(v9)
                        case 42:
                            v4 = v4
                        case -28:
                            if (65535 > (^(v4))) {
                                _ = len(v9)
                                v35 := GenS3{F0: ("a" + ("a" + v2)), F1: v2}
                                _ = v35
                            }
                        }
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    fmt.Println(v10, v2, fi15)
                    v9 = append(v9, (fi22 + (fi15 ^ -2147483648)))
                }
            }
            switch fi15 {
            case 32767:
                var v36 uint32 = (v4 / ((4294967295) | 1))
                _ = v36
            case 128:
                fmt.Println(fi22, v2, fi15)
            case 224:
                v4 = (^(v4))
            default:
                switch fi22 {
                case -1:
                    v9 = append(v9, (^(fi22)))
                case 9223372036854775807:
                    var v37 int32 = -1
                    _ = v37
                case -9223372036854775808:
                    var v38 int16 = 1
                    _ = v38
                default:
                    switch fi22 {
                    case -996:
                        var v39 int16 = 937
                        _ = v39
                    case -9223372036854775808:
                        var v40 uint16 = 32767
                        _ = v40
                    }
                }
                v41 := GenS4{F0: 4294967295, F1: 42, F2: 4294967295}
                v42 := GenS2{F0: 1, F1: 1}
                switch (-(fi15)) {
                case -2147483648:
                    v4 -= (^(v4))
                default:
                    var v43 int8 = v3
                    var v44 uint8 = 127
                    _ = v43
                    _ = v44
                }
                _ = v41
                _ = v42
            }
            switch (v2 + ("foo bar" + "\n\t")) {
            case "\n\t":
                if true {
                    var v45 uint32 = (v4 | v4)
                    var v46 bool = true
                    _ = v45
                    _ = v46
                } else {
                    var v47 int8 = 42
                    v1 += 0
                    if len(v9) > 0 {
                        v9[0] = fi22
                    }
                    _ = v47
                }
            case "hello":
                var v48 int16 = 32767
                _ = v48
            case "":
                _ = len(v9)
            }
            var v49 bool = false
            _ = fi22
            _ = v49
        }
        _ = fi15
    }
    _ = len(v9)
    switch 1 {
    case 9223372036854775807:
        for fi50 := 0; fi50 < 19; fi50++ {
            var v51 uint64 = 1
            v52 := GenS1{F0: v51, F1: (!(0 == int16(v1)))}
            _ = fi50
            _ = v51
            _ = v52
        }
    }
    var v53 uint32 = (v4 * v4)
    var v54 int16 = 32767
    for fi55 := 0; fi55 < 3; fi55++ {
        if len(v9) > 0 {
            v9[0] = (fi55 - -1)
        }
        if false {
            if ((!(false && (!true))) || false) {
                fmt.Println(v10)
                fmt.Println(v1)
                v56 := GenS5{F0: 487, F1: 1, F2: (v3 / ((0) | 1))}
                _ = v56
            }
            if len(v9) > 0 {
                v9[0] = 32767
            }
        } else {
            v57 := GenS6{F0: v4, F1: 2147483647, F2: false}
            if len(v9) > 0 {
                v9[0] = int(v53)
            }
            v1 += v3
            _ = v57
        }
        _ = fi55
    }
    if false {
        switch 127 {
        case 128:
            v4 = v53
        case -9223372036854775808:
            v58 := GenS7{F0: 2, F1: 42}
            _ = v58
        case 1:
            switch -479 {
            case -2147483648:
                v53 = (^(v4))
            default:
                var v59 uint16 = 95
                _ = v59
            }
        default:
            var v60 int32 = -2147483648
            _ = v60
        }
        for fi61 := 0; fi61 < 15; fi61++ {
            for fi62 := 0; fi62 < 10; fi62++ {
                v1 = (v3 - -128)
                fmt.Println(v2)
                _ = fi62
            }
            if (true || (v4 >= v4)) {
                if (2 != 2) {
                    switch (fi61 | 896) {
                    case 128:
                        v53 += v53
                    case 69:
                        fmt.Println(v1, v10)
                    case 32767:
                        switch fi61 {
                        case 9223372036854775807:
                            for fi63 := 0; fi63 < 16; fi63++ {
                                v4 -= (^(v53))
                                _ = fi63
                            }
                        case -918:
                            var v64 int = fi61
                            _ = v64
                        default:
                            var v65 int16 = 1
                            _ = v65
                        }
                    default:
                        for fi66 := 0; fi66 < 10; fi66++ {
                            if len(v9) > 0 {
                                v9[0] = 457
                            }
                            var v67 uint32 = v4
                            _ = fi66
                            _ = v67
                        }
                    }
                }
            }
            _ = fi61
        }
        v68 := GenS2{F0: 42, F1: 16}
        _ = v68
    }
    switch -509 {
    case -486:
        for fi69 := 0; fi69 < 5; fi69++ {
            if (v53 < 4294967295) {
                if true {
                    fmt.Println(v3, v1)
                    for fi70 := 0; fi70 < 16; fi70++ {
                        fmt.Println(fi70, v53)
                        _ = fi70
                    }
                }
                v4 = v4
            }
            _ = fi69
        }
    default:
        var v71 uint8 = 255
        _ = v71
    }
    v72 := GenS8{F0: 220, F1: 734, F2: v2}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    if (779 != 395) {
        v4 = v53
        if ((v53 + 2147483647) >= 2147483648) {
            var v73 int32 = -1
            switch 2147483647 {
            case 32767:
                v73 -= (^(v73))
            case 2147483647:
                var v74 int = 9223372036854775807
                _ = v74
            case -679:
                v4 = 424
            }
            switch "\n\t" {
            case "foo bar":
                var v75 uint64 = 42
                _ = v75
            default:
                fmt.Println(v2)
            }
            _ = v73
        } else {
            if true {
                for fi76 := 0; fi76 < 16; fi76++ {
                    v77 := GenS2{F0: 750, F1: 18446744073709551615}
                    v2 = v2
                    _ = fi76
                    _ = v77
                }
                v54 -= 526
                for fi78 := 0; fi78 < 16; fi78++ {
                    v79 := GenS1{F0: 4294967295, F1: false}
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = fi78
                    _ = v79
                }
            }
            var v80 uint32 = v4
            fmt.Println(v1, v80)
            _ = v80
        }
        var v81 uint64 = 42
        _ = v81
    } else {
        fmt.Println(v53, v4)
        for fi82 := 0; fi82 < 6; fi82++ {
            v83 := GenS9{F0: "a", F1: 499, F2: 125}
            var v84 uint16 = 0
            v2 = v2
            _ = fi82
            _ = v83
            _ = v84
        }
    }
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v72
    _ = v9
}
