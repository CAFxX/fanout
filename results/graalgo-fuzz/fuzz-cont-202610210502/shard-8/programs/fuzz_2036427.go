package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS5 struct { F0 bool; F1 int32; F2 int64 }

func main() {
    var v1 int32 = 1
    v2 := []int{2147483647, 0, 4294967295, -503}
    v3 := make(map[int]uint32)
    v4 := []string{}
    v1 += (v1 << ((42) & 63))
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 float32 = -2.25
            var v8 uint16 = 65535
            _ = fi6
            _ = v7
            _ = v8
        }
        switch ("\n\t" + "hello") {
        case "hello":
            if false {
                switch ("\n\t" + "hello") {
                case "":
                    var v9 int64 = 9223372036854775807
                    _ = v9
                case "\n\t":
                    var v10 uint64 = 42
                    _ = v10
                }
                var v11 uint32 = 503
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := GenS1{F0: 411, F1: 255, F2: (v11 / ((v11) | 1))}
                    switch "hello" {
                    case "a":
                        switch (^(fi12)) {
                        case -568:
                            v14 := GenS1{F0: (v11 >> ((2) & 63)), F1: 0, F2: (v11 << ((2) & 63))}
                            _ = v14
                        case -304:
                            if len(v2) > 0 {
                                v2[0] = fi5
                            }
                        case -205:
                            v15 := GenS1{F0: (v11 - v11), F1: 255, F2: (v11 - (v11 + v11))}
                            _ = v15
                        default:
                            v16 := GenS1{F0: v11, F1: 128, F2: v11}
                            _ = v16
                        }
                    case "foo bar":
                        var v17 string = "a"
                        _ = v17
                    case "hello":
                        v18 := GenS2{F0: 0, F1: 366, F2: -9223372036854775808}
                        _ = v18
                    default:
                        if true {
                            var v19 uint8 = 0
                            var v20 int64 = -54
                            var v21 uint16 = 516
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            fmt.Println(fi12, fi5, fi5)
                        }
                    }
                    var v22 float64 = (0.1 - 1e10)
                    switch "a" {
                    case "":
                        if (!(-617 <= -721)) {
                            v23 := GenS2{F0: 42, F1: 414, F2: 42}
                            switch 127 {
                            case 0:
                                var v24 float32 = 0.0
                                _ = v24
                            case 32768:
                                switch fi12 {
                                case -9223372036854775808:
                                    var v25 int = (fi12 * int(v11))
                                    _ = v25
                                }
                            default:
                                v11 = v11
                            }
                            switch fi12 {
                            case 9223372036854775807:
                                if false {
                                    switch (("" + "foo bar") + ("" + ("\n\t" + "foo bar"))) {
                                    case "a":
                                        switch fi12 {
                                        case 9223372036854775807:
                                            if false {
                                                v26 := GenS1{F0: (v11 + (^(v11))), F1: 1, F2: (v11 >> ((255) & 63))}
                                                _ = v26
                                            } else {
                                                fmt.Println(v1, fi5)
                                                v27 := GenS1{F0: (v11 % ((v11) | 1)), F1: 1, F2: 2147483648}
                                                fmt.Println(fi5, v22, v1)
                                                _ = v27
                                            }
                                        default:
                                            var v28 int8 = 8
                                            _ = v28
                                        }
                                        var v29 int32 = (v1 >> ((634) & 63))
                                        v11 -= v11
                                        _ = v29
                                    case "hello":
                                        v22 = 0.0
                                    case "\n\t":
                                        var v30 uint = 255
                                        _ = v30
                                    default:
                                        var v31 uint16 = 32767
                                        _ = v31
                                    }
                                    delete(v3, 860)
                                }
                            case -2147483648:
                                v32 := GenS1{F0: (^(v11)), F1: uint8(v11), F2: v11}
                                _ = v32
                            default:
                                switch (fi12 * -821) {
                                case 32768:
                                    fmt.Println(v22, v1, v11)
                                case 32767:
                                    v11 = 2147483648
                                }
                            }
                            _ = v23
                        } else {
                            v11 = 529
                            v11 += 2147483647
                        }
                    }
                    _ = fi12
                    _ = v13
                    _ = v22
                }
                _ = v11
            }
        }
        if true {
            if (!(true || (0 >= 1))) {
                for fi33 := 0; fi33 < 15; fi33++ {
                    v34 := GenS2{F0: 0, F1: 4294967295, F2: 1}
                    fmt.Println(fi33, v1)
                    _ = fi33
                    _ = v34
                }
                var v35 uint64 = 974
                _ = v35
            } else {
                switch "a" {
                case "foo bar":
                    for fi36 := 0; fi36 < 5; fi36++ {
                        v4 = append(v4, "foo bar")
                        _ = fi36
                    }
                }
                var v37 uint32 = 2
                _ = v37
            }
            v1 -= (v1 + 91)
        } else {
            switch fi5 {
            case -9223372036854775808:
                fmt.Println(fi5)
            case 9223372036854775807:
                if false {
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        switch "\n\t" {
                        case "hello":
                            v1 = (v1 >> ((42) & 63))
                        case "\n\t":
                            if true {
                                var v38 int16 = -277
                                var v39 uint32 = 39
                                var v40 uint64 = 2
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                        default:
                            var v41 int8 = 0
                            _ = v41
                        }
                    case "hello":
                        fmt.Println(fi5)
                    default:
                        v1 += (-(v1))
                    }
                }
            }
        }
        _ = fi5
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        switch "hello" {
        case "foo bar":
            var v43 string = (("" + "foo bar") + "hello")
            _ = v43
        case "":
            v1 += v1
        case "hello":
            var v44 uint32 = 847
            _ = v44
        default:
            if false {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v1 -= (v1 | -1)
                }
            } else {
                v1 += (-(v1))
                switch "" {
                case "":
                    var v45 int64 = 9223372036854775807
                    _ = v45
                default:
                    fmt.Println(fi42, v1)
                }
            }
        }
        _ = fi42
    }
    v1 -= (v1 % ((v1) | 1))
    if true {
        v1 = (v1 | -1)
        var v46 int64 = 1
        _ = v46
    } else {
        for fi47 := 0; fi47 < 11; fi47++ {
            v1 = v1
            if (3.14159 < 1e10) {
                if true {
                    var v48 uint = 0
                    fmt.Println(v48)
                    v48 = 255
                    _ = v48
                } else {
                    v1 = (-(v1))
                }
            }
            _ = fi47
        }
        v1 = -830
        var v49 uint32 = 392
        _ = v49
    }
    fmt.Println(v1)
    if (true && false) {
        for fi50 := 0; fi50 < 9; fi50++ {
            var v51 uint32 = 4294967295
            _ = fi50
            _ = v51
        }
        if ((!(3.14159 > 0.1)) || (367 < 62)) {
            var v52 int16 = 42
            for fi53 := 0; fi53 < 8; fi53++ {
                if ((42 > 1) && ("hello" == "foo bar")) {
                    fmt.Println(v52, v1, fi53)
                } else {
                    v54 := GenS3{F0: 18446744073709551615}
                    var v55 int = (-(fi53))
                    _ = v54
                    _ = v55
                }
                fmt.Println(fi53, v1, fi53)
                v56 := GenS4{F0: 957, F1: 255, F2: v52}
                _ = fi53
                _ = v56
            }
            _ = v52
        }
        for fi57 := 0; fi57 < 8; fi57++ {
            fmt.Println(fi57)
            switch (("foo bar" + "\n\t") + "hello") {
            case "hello":
                if (364 >= 42) {
                    var v58 bool = (!(!(42 != 65535)))
                    v59 := GenS4{F0: 587, F1: 65535, F2: 781}
                    fmt.Println(v1, fi57, v58)
                    _ = v58
                    _ = v59
                } else {
                    v3[(^(fi57))] = 575
                    delete(v3, fi57)
                }
            }
            _ = fi57
        }
    } else {
        switch "foo bar" {
        case "":
            var v60 bool = false
            _ = v60
        case "hello":
            var v61 bool = false
            _ = v61
        default:
            var v62 float32 = 3.14159
            _ = v62
        }
        switch "foo bar" {
        case "a":
            if ("\n\t" > "foo bar") {
                if (true && false) {
                    var v63 int16 = -487
                    v63 += v63
                    fmt.Println(v1, v63)
                    _ = v63
                } else {
                    var v64 float32 = 3.14159
                    if false {
                        var v65 int64 = -300
                        switch 4294967295 {
                        case 9223372036854775807:
                            for fi66 := 0; fi66 < 11; fi66++ {
                                fmt.Println(v64)
                                switch (fi66 - (^(fi66))) {
                                case -856:
                                    fmt.Println(fi66, v65)
                                case 1:
                                    v64 += 1e10
                                default:
                                    fmt.Println(v65, fi66, v64)
                                }
                                v67 := GenS4{F0: 1, F1: 261, F2: -1}
                                _ = fi66
                                _ = v67
                            }
                        case 1:
                            v1 = v1
                        case 594:
                            var v68 int64 = (^(v65))
                            _ = v68
                        }
                        _ = v65
                    }
                    if (-2.25 < 1.5) {
                        if (255 < 255) {
                            v64 = v64
                        } else {
                            v1 = (v1 - (v1 * (v1 & v1)))
                            switch ((("foo bar" + "") + ("\n\t" + "hello")) + "a") {
                            case "a":
                                v69 := GenS5{F0: (-128 > -128), F1: v1, F2: -1}
                                _ = v69
                            case "foo bar":
                                v64 = v64
                            case "hello":
                                if true {
                                    switch (("hello" + "\n\t") + "a") {
                                    case "\n\t":
                                        var v70 string = ("a" + "a")
                                        _ = v70
                                    case "hello":
                                        v64 += ((1e10 * v64) - (3.14159 - 1e10))
                                    }
                                    var v71 bool = ((0 <= 4294967295) && true)
                                    _ = v71
                                } else {
                                    v1 += (v1 * (v1 >> ((18446744073709551615) & 63)))
                                    var v72 float64 = (1e10 + ((1e10 * (0.0 / 1.0)) - float64(v64)))
                                    var v73 float32 = ((1.5 * (1e10 * v64)) / ((float32((v72 * v72)) / (((((v64 + (v64 / float32((3.14159 + v72)))) / float32(v72)) / 3.14159) / 3.14159) - (v64 / (0.1 / 3.14159)))) + v64))
                                    _ = v72
                                    _ = v73
                                }
                            }
                        }
                        var v74 int = -2147483648
                        v75 := GenS1{F0: 0, F1: 252, F2: 4294967295}
                        _ = v74
                        _ = v75
                    } else {
                        switch 0 {
                        case 128:
                            _ = len(v3)
                        }
                    }
                    _ = v64
                }
            }
        case "":
            var v76 float32 = 0.1
            _ = v76
        }
        v77 := v3[127]
        _ = v77
    }
    switch 127 {
    case 0:
        for fi78 := 0; fi78 < 16; fi78++ {
            for fi79 := 0; fi79 < 2; fi79++ {
                var v80 uint8 = 244
                v80 = (v80 - (v80 * 1))
                var v81 uint = 42
                _ = fi79
                _ = v80
                _ = v81
            }
            for fi82 := 0; fi82 < 1; fi82++ {
                v83 := v3[-163]
                v83 -= (v83 * 65535)
                v3[fi78] = (v83 >> ((42) & 63))
                v1 += (-(v1))
                _ = fi82
                _ = v83
            }
            var v84 uint8 = 42
            _ = fi78
            _ = v84
        }
    case 4294967295:
        v1 += v1
    }
    var v85 uint16 = 491
    var v86 uint64 = 325
    v86 += 42
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v85
    _ = v86
}
