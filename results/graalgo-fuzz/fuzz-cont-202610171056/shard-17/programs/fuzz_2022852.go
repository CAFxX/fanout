package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int16 }
type GenS2 struct { F0 int8; F1 uint; F2 int8 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int32; F2 int64 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 uint; F1 uint16; F2 uint }
type GenS8 struct { F0 int; F1 bool; F2 int }
type GenS9 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint32; F1 uint32 }
type GenS13 struct { F0 int32; F1 uint32 }
func genFn0() uint16 {
    var v4 int32 = 1
    v5 := []int{32767, 42}
    if (!(("hello" + ("\n\t" + ("\n\t" + ("hello" + "")))) != "")) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v4 += 2147483647
        v5 = append(v5, 426)
    }
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{}
    var v3 int8 = -23
    if false {
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 uint64 = 331
            _ = fi6
            _ = v7
        }
    } else {
        var v8 int16 = 1000
        v3 = (v3 - 1)
        _ = v8
    }
    switch "hello" {
    case "":
        for fi9 := 0; fi9 < 8; fi9++ {
            switch "\n\t" {
            case "a":
                var v10 int8 = v3
                _ = v10
            case "foo bar":
                var v11 string = "foo bar"
                _ = v11
            case "\n\t":
                v12 := GenS1{F0: 621, F1: 0, F2: 801}
                _ = v12
            default:
                fmt.Println(v3)
            }
            _ = fi9
        }
    default:
        v13 := GenS1{F0: 2147483647, F1: (v3 / ((1) | 1)), F2: -32768}
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            delete(v1, fi15)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch fi15 {
            case 914:
                switch -2147483648 {
                case -1:
                    v16 := v1[fi14]
                    _ = v16
                default:
                    delete(v1, (-(fi14)))
                    if (!((("foo bar" + "hello") + "") != "\n\t")) {
                        v2 = append(v2, (fi15 ^ (^(fi14))))
                        if (!(true && true)) {
                            fmt.Println(fi15, fi15)
                        } else {
                            v17 := GenS1{F0: -280, F1: (v3 >> ((905) & 63)), F2: 42}
                            if true {
                                if ((v3 ^ v3) > -1) {
                                    v3 -= (v3 ^ v3)
                                    v3 += 100
                                    var v18 uint64 = 321
                                    _ = v18
                                } else {
                                    v3 = 127
                                }
                            }
                            _ = v17
                        }
                        v3 = v3
                    }
                    if false {
                        fmt.Println(v3)
                        if (354 < -2147483648) {
                            if true {
                                _ = len(v1)
                            }
                            if len(v2) > 0 {
                                v2[0] = fi14
                            }
                            switch fi14 {
                            case 32767:
                                v2 = append(v2, fi15)
                            case 128:
                                if (((1 > -1) || false) || true) {
                                    v3 -= 1
                                    var v19 uint32 = 0
                                    var v20 float64 = (1e10 + float64((1.5 - float32(1.5))))
                                    _ = v19
                                    _ = v20
                                }
                            case -9223372036854775808:
                                v21 := GenS1{F0: -2147483648, F1: v3, F2: 1000}
                                var v22 float32 = 0.0
                                delete(v1, 32767)
                                switch fi14 {
                                case 42:
                                    var v23 uint64 = 0
                                    _ = v23
                                }
                                _ = v21
                                _ = v22
                            default:
                                switch (fi14 ^ (fi14 & fi14)) {
                                case 251:
                                    v24 := GenS1{F0: 1, F1: v3, F2: -192}
                                    if (((0.1 * (0.0 - ((float64(1e10) + float64(1e10)) / 1e10))) / -2.25) > float64(-2.25)) {
                                        var v25 uint = 42
                                        if (0.1 >= 3.14159) {
                                            switch (-(fi14)) {
                                            case 42:
                                                v26 := GenS2{F0: -1, F1: v25, F2: (v3 % ((100) | 1))}
                                                _ = v26
                                            }
                                            var v27 uint64 = 4294967295
                                            if true {
                                                fmt.Println(v3, v25, v27)
                                            } else {
                                                _ = len(v2)
                                                fmt.Println(fi15, v3)
                                            }
                                            v27 += v27
                                            if ((3.14159 * 3.14159) > -2.25) {
                                                if true {
                                                    fmt.Println(v3)
                                                } else {
                                                    var v28 float32 = (3.14159 * ((((1.5 - (((3.14159 - -2.25) / 1e10) / 1.0)) / ((0.1 - (1e10 - (1.5 - (0.1 * 0.1)))) / (float32(0.1) / 0.1))) / -2.25) / 1.0))
                                                    _ = v28
                                                }
                                                if (true && false) {
                                                    var v29 int16 = 1000
                                                    if true {
                                                        v30 := GenS3{F0: genFn0(), F1: -598, F2: 128}
                                                        _ = v30
                                                    }
                                                    _ = v29
                                                } else {
                                                    v3 = 101
                                                    var v31 uint32 = 4294967295
                                                    _ = v31
                                                }
                                            } else {
                                                v3 = v3
                                                if false {
                                                    v32 := GenS4{F0: genFn0(), F1: -636, F2: 42}
                                                    switch "a" {
                                                    case "\n\t":
                                                        v33 := GenS5{F0: 0, F1: 42}
                                                        _ = v33
                                                    case "a":
                                                        v25 = v25
                                                    }
                                                    _ = v32
                                                } else {
                                                    var v34 uint = (v25 | v25)
                                                    v2 = append(v2, -1)
                                                    var v35 uint = v25
                                                    _ = v34
                                                    _ = v35
                                                }
                                                var v36 float32 = 3.14159
                                                v37 := v1[(fi14 & (fi15 ^ 1))]
                                                _ = v36
                                                _ = v37
                                            }
                                            _ = v27
                                        } else {
                                            v38 := GenS2{F0: v3, F1: v25, F2: 1}
                                            _ = v38
                                        }
                                        _ = v25
                                    }
                                    _ = v24
                                case 62:
                                    var v39 int16 = 1000
                                    _ = v39
                                }
                            }
                        } else {
                            v40 := GenS5{F0: -246, F1: 2}
                            v2 = append(v2, (-(fi15)))
                            v41 := GenS2{F0: (v3 | 42), F1: 42, F2: -1}
                            v3 += v3
                            _ = v40
                            _ = v41
                        }
                        v42 := GenS6{F0: 126, F1: -523}
                        _ = v42
                    } else {
                        v3 = (^(v3))
                    }
                }
            case -9223372036854775808:
                v43 := GenS7{F0: 42, F1: 777, F2: 587}
                _ = v43
            case 768:
                v3 += 0
            default:
                fmt.Println(fi14, v3, v3)
                v44 := GenS5{F0: 1, F1: 2}
                _ = v44
            }
            _ = fi15
        }
        switch ("hello" + "\n\t") {
        case "":
            if true {
                switch fi14 {
                case 0:
                    switch "hello" {
                    case "hello":
                        fmt.Println(v3, v3)
                    case "\n\t":
                        v3 = (v3 >> ((2) & 63))
                    }
                }
                var v45 float32 = float32(1e10)
                _ = v45
            }
        case "hello":
            for fi46 := 0; fi46 < 11; fi46++ {
                v3 += v3
                _ = fi46
            }
        case "a":
            v47 := v1[(-(fi14))]
            _ = v47
        }
        v2 = append(v2, fi14)
        _ = fi14
    }
    if true {
        if false {
            var v48 float64 = 1.5
            for fi49 := 0; fi49 < 9; fi49++ {
                v3 += (v3 << ((2) & 63))
                if false {
                    v2 = append(v2, (fi49 >> ((411) & 63)))
                    var v50 int8 = (v3 >> ((2) & 63))
                    for fi51 := 0; fi51 < 18; fi51++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi51
                    }
                    _ = v50
                } else {
                    v52 := GenS8{F0: (fi49 & (-(fi49))), F1: true, F2: 4294967295}
                    if (!((v48 > ((v48 - ((float64(0.0) - (-2.25 + float64(0.0))) + 0.0)) - float64((-2.25 + 1.5)))) || true)) {
                        for fi53 := 0; fi53 < 2; fi53++ {
                            v3 += (v3 ^ (v3 - (v3 >> ((157) & 63))))
                            var v54 uint8 = 128
                            _ = len(v2)
                            _ = fi53
                            _ = v54
                        }
                        switch "foo bar" {
                        case "a":
                            var v55 uint64 = 0
                            _ = v55
                        case "foo bar":
                            delete(v1, fi49)
                        }
                        switch "hello" {
                        case "foo bar":
                            var v56 uint32 = 4294967295
                            _ = v56
                        case "a":
                            var v57 uint16 = 42
                            _ = v57
                        case "":
                            v58 := GenS9{F0: 44, F1: 42, F2: 154}
                            _ = v58
                        }
                    } else {
                        fmt.Println(fi49, v48, fi49)
                        v48 -= (0.1 + -2.25)
                    }
                    var v59 int8 = -128
                    _ = v52
                    _ = v59
                }
                v3 -= v3
                _ = fi49
            }
            _ = v48
        } else {
            var v60 int16 = 0
            v60 = 32767
            _ = v60
        }
        var v61 bool = (462 != -2147483648)
        switch ("foo bar" + "a") {
        case "hello":
            v62 := GenS5{F0: -321, F1: 0}
            _ = v62
        case "\n\t":
            var v63 uint = 18446744073709551615
            _ = v63
        default:
            v61 = false
        }
        _ = v61
    } else {
        v3 = v3
        var v64 uint32 = 2147483647
        var v65 uint8 = 2
        _ = v64
        _ = v65
    }
    v3 += v3
    v66 := GenS6{F0: 0, F1: -577}
    v3 = 83
    if (!true) {
        var v67 uint = 255
        for fi68 := 0; fi68 < 6; fi68++ {
            if true {
                switch fi68 {
                case 42:
                    v3 += v3
                case 9223372036854775807:
                    var v69 int16 = 42
                    _ = v69
                }
            }
            for fi70 := 0; fi70 < 20; fi70++ {
                var v71 string = ("" + (("a" + "hello") + "hello"))
                v71 = (v71 + v71)
                _ = fi70
                _ = v71
            }
            _ = fi68
        }
        switch (((("" + "a") + "a") + "") + "") {
        case "hello":
            for fi72 := 0; fi72 < 10; fi72++ {
                v3 = v3
                v3 += v3
                for fi73 := 0; fi73 < 7; fi73++ {
                    var v74 int8 = v3
                    switch (fi72 % (((fi72 * fi73)) | 1)) {
                    case 32767:
                        var v75 uint64 = 764
                        _ = v75
                    case -9223372036854775808:
                        var v76 bool = true
                        _ = v76
                    case 42:
                        var v77 int64 = 42
                        _ = v77
                    }
                    if ((0.1 + ((0.1 * 0.0) - 1.5)) == 0.1) {
                        var v78 int8 = (v74 / ((v74) | 1))
                        _ = v78
                    } else {
                        if (!((3.14159 / ((3.14159 + (0.0 - -2.25)) * 1.5)) >= float64(3.14159))) {
                            if (255 <= 4294967295) {
                                v79 := v1[(fi73 >> ((979) & 63))]
                                v80 := GenS10{F0: 764}
                                if len(v2) > 0 {
                                    v2[0] = -924
                                }
                                _ = v79
                                _ = v80
                            }
                            if false {
                                v81 := GenS11{F0: 121}
                                _ = v81
                            }
                            v82 := GenS1{F0: 2147483647, F1: v3, F2: -674}
                            _ = v82
                        } else {
                            v83 := GenS12{F0: 65535, F1: 723}
                            _ = v83
                        }
                    }
                    _ = fi73
                    _ = v74
                }
                _ = fi72
            }
        case "\n\t":
            _ = len(v1)
        case "":
            if (true || false) {
                v3 += v3
                if (-2147483648 != -36) {
                    switch ("a" + ("foo bar" + (("foo bar" + (("\n\t" + ("a" + "")) + ("a" + "a"))) + "a"))) {
                    case "":
                        if true {
                            var v84 float64 = 3.14159
                            _ = v84
                        }
                    case "a":
                        for fi85 := 0; fi85 < 13; fi85++ {
                            var v86 int = 9223372036854775807
                            _ = fi85
                            _ = v86
                        }
                    default:
                        var v87 uint32 = 660
                        _ = v87
                    }
                    fmt.Println(v3)
                }
                v88 := GenS13{F0: -1, F1: 2}
                _ = v88
            } else {
                v67 -= 2
            }
        default:
            v89 := GenS3{F0: 930, F1: -9223372036854775808, F2: 127}
            _ = v89
        }
        _ = v67
    }
    for fi90 := 0; fi90 < 15; fi90++ {
        for fi91 := 0; fi91 < 19; fi91++ {
            var v92 int16 = 1
            switch "\n\t" {
            case "hello":
                fmt.Println(v92, v3, v3)
            case "\n\t":
                fmt.Println(v3)
            case "a":
                var v93 bool = ((v92 | v92) > (^(v92)))
                _ = v93
            }
            v94 := GenS3{F0: 42, F1: -9223372036854775808, F2: 1}
            _ = fi91
            _ = v92
            _ = v94
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi95 := 0; fi95 < 7; fi95++ {
            switch (fi95 ^ fi95) {
            case 180:
                delete(v1, (fi90 + (fi95 % (((-(fi95))) | 1))))
            case 633:
                var v96 float32 = 1e10
                _ = v96
            }
            fmt.Println(fi95, fi90)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi95
        }
        _ = fi90
    }
    var v97 uint = 619
    switch 127 {
    case 302:
        if (2147483647 > -1) {
            var v98 uint16 = 1
            for fi99 := 0; fi99 < 3; fi99++ {
                delete(v1, fi99)
                _ = fi99
            }
            _ = v98
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v66
    _ = v97
}
