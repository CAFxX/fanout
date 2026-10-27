package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 bool; F2 uint }
type GenS12 struct { F0 int; F1 uint16; F2 int16 }

func main() {
    var v1 int8 = 127
    v2 := []int{395}
    var v3 int8 = 78
    var v4 uint16 = 255
    v5 := make(map[int]int)
    v4 += 2
    var v6 float64 = 0.0
    v6 += v6
    switch "a" {
    case "\n\t":
        switch "foo bar" {
        case "foo bar":
            v7 := GenS1{F0: int(v3), F1: 32767, F2: 814}
            _ = v7
        }
    case "":
        switch "" {
        case "":
            var v8 float64 = float64(1e10)
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 1; fi9++ {
                fmt.Println(v4, fi9, v3)
                _ = fi9
            }
        case "a":
            v4 -= (v4 | (^(v4)))
        default:
            if (1 <= 9223372036854775807) {
                _ = len(v5)
            } else {
                switch "\n\t" {
                case "a":
                    v6 = 3.14159
                case "foo bar":
                    for fi10 := 0; fi10 < 4; fi10++ {
                        if true {
                            fmt.Println(v3)
                            if true {
                                switch "hello" {
                                case "":
                                    _ = len(v5)
                                case "\n\t":
                                    v6 -= v6
                                case "hello":
                                    var v11 int8 = (^(v3))
                                    _ = v11
                                }
                                v1 += v3
                                switch (fi10 * fi10) {
                                case 32767:
                                    if false {
                                        var v12 int16 = -1
                                        v6 = v6
                                        if (!(!true)) {
                                            var v13 int8 = v3
                                            v14 := GenS1{F0: (fi10 * fi10), F1: fi10, F2: -166}
                                            for fi15 := 0; fi15 < 14; fi15++ {
                                                if (!(42 > -128)) {
                                                    fmt.Println(fi10, v3)
                                                    v16 := v5[4294967295]
                                                    if false {
                                                        v17 := GenS2{F0: (v4 | (v4 % (((v4 / (((v4 | 842)) | 1))) | 1))), F1: 18446744073709551615, F2: 42}
                                                        var v18 bool = (!false)
                                                        var v19 int64 = 684
                                                        _ = v17
                                                        _ = v18
                                                        _ = v19
                                                    } else {
                                                        switch fi10 {
                                                        case -162:
                                                            v20 := GenS1{F0: 127, F1: 878, F2: -1}
                                                            _ = v20
                                                        case 32767:
                                                            v6 += v6
                                                        case 9223372036854775807:
                                                            v16 += fi10
                                                            var v21 string = ("foo bar" + "a")
                                                            v22 := GenS1{F0: (^(fi10)), F1: v16, F2: 1}
                                                            switch (("hello" + (("a" + ("hello" + v21)) + "hello")) + "a") {
                                                            case "a":
                                                                var v23 uint8 = 0
                                                                _ = v23
                                                            default:
                                                                v6 = ((v6 * v6) - v6)
                                                            }
                                                            switch (^(v16)) {
                                                            case 42:
                                                                delete(v5, 579)
                                                            default:
                                                                if (v21 == "a") {
                                                                    if false {
                                                                        v12 -= v12
                                                                        if len(v2) > 0 {
                                                                            v2[0] = -2147483648
                                                                        }
                                                                        v16 = 42
                                                                    } else {
                                                                        var v24 int = (fi15 % (((-(fi15))) | 1))
                                                                        var v25 uint = 4294967295
                                                                        _ = v24
                                                                        _ = v25
                                                                    }
                                                                    fmt.Println(v12)
                                                                    v26 := v5[(fi10 / (((^(fi15))) | 1))]
                                                                    _ = v26
                                                                }
                                                            }
                                                            _ = v21
                                                            _ = v22
                                                        }
                                                        v27 := GenS3{F0: 42, F1: 1}
                                                        v28 := GenS3{F0: 479, F1: 42}
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                    _ = v16
                                                } else {
                                                    if true {
                                                        var v29 int16 = (v12 * (v12 % (((v12 ^ v12)) | 1)))
                                                        if ("\n\t" != ("" + "\n\t")) {
                                                            fmt.Println(fi10, v4, v29)
                                                            v6 -= v6
                                                            var v30 float64 = (v6 + v6)
                                                            v30 -= v6
                                                            _ = v30
                                                        } else {
                                                            var v31 uint16 = v4
                                                            _ = v31
                                                        }
                                                        if (!false) {
                                                            _ = len(v5)
                                                            v32 := GenS4{F0: "foo bar"}
                                                            switch 1 {
                                                            case 127:
                                                                var v33 bool = true
                                                                _ = v33
                                                            default:
                                                                var v34 uint64 = 0
                                                                _ = v34
                                                            }
                                                            var v35 bool = (int16(v3) < -967)
                                                            v35 = ((v4 + v4) == (v4 << ((65535) & 63)))
                                                            _ = v32
                                                            _ = v35
                                                        }
                                                        _ = v29
                                                    }
                                                }
                                                var v36 uint32 = 42
                                                v37 := GenS2{F0: (v4 | (v4 - (v4 | v4))), F1: 1, F2: 1}
                                                var v38 string = ("\n\t" + "hello")
                                                _ = fi15
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            }
                                            _ = v13
                                            _ = v14
                                        }
                                        _ = v12
                                    } else {
                                        v6 = ((v6 / 1e10) / float64(1.5))
                                        if false {
                                            var v39 int8 = 1
                                            _ = v39
                                        }
                                    }
                                case -413:
                                    v6 -= v6
                                default:
                                    v3 = v1
                                }
                            } else {
                                for fi40 := 0; fi40 < 15; fi40++ {
                                    v41 := GenS5{F0: 1, F1: 2}
                                    var v42 float64 = float64(0.0)
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                            }
                            switch (^(fi10)) {
                            case 128:
                                var v43 int32 = -26
                                _ = v43
                            case 32768:
                                for fi44 := 0; fi44 < 4; fi44++ {
                                    fmt.Println(v3)
                                    switch fi10 {
                                    case 4294967295:
                                        v45 := GenS3{F0: 4294967295, F1: 2147483647}
                                        _ = v45
                                    case 90:
                                        var v46 int = 0
                                        _ = v46
                                    }
                                    v3 = (v1 + v1)
                                    _ = fi44
                                }
                            case -503:
                                for fi47 := 0; fi47 < 10; fi47++ {
                                    _ = len(v2)
                                    v1 += v3
                                    switch fi10 {
                                    case -2147483648:
                                        v1 = (v1 << ((2) & 63))
                                    case -934:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case 4294967295:
                                        fmt.Println(fi10, fi47)
                                    default:
                                        fmt.Println(fi10)
                                    }
                                    v48 := v5[(^(fi10))]
                                    if (65535 <= 2) {
                                        delete(v5, fi10)
                                    } else {
                                        fmt.Println(v4, v48)
                                        if true {
                                            v49 := GenS6{F0: (v3 & v1)}
                                            _ = v49
                                        }
                                        v3 = int8(v48)
                                        if true {
                                            fmt.Println(v3, fi47)
                                            _ = len(v5)
                                            var v50 uint8 = 127
                                            _ = v50
                                        }
                                        switch ("hello" + "\n\t") {
                                        case "hello":
                                            if len(v2) > 0 {
                                                v2[0] = (fi47 + fi10)
                                            }
                                        case "\n\t":
                                            v51 := GenS7{F0: 2147483647, F1: (v3 | (v3 - v3))}
                                            _ = v51
                                        case "":
                                            var v52 bool = (true || (!false))
                                            _ = v52
                                        default:
                                            var v53 float32 = float32((0.1 + float64(-2.25)))
                                            _ = v53
                                        }
                                    }
                                    _ = fi47
                                    _ = v48
                                }
                            default:
                                v3 = (-(v1))
                            }
                        }
                        _ = fi10
                    }
                }
                switch ("hello" + ("\n\t" + ("a" + "\n\t"))) {
                case "foo bar":
                    v54 := GenS8{F0: 65535, F1: -9223372036854775808}
                    _ = v54
                case "\n\t":
                    var v55 int8 = (-(v3))
                    _ = v55
                case "":
                    if (555 <= 2147483647) {
                        v56 := GenS9{F0: 2147483648, F1: v1, F2: 2}
                        if ((2 > 255) || (-2.25 == v6)) {
                            var v57 string = "hello"
                            if (158 >= 0) {
                                v6 -= -2.25
                            } else {
                                if (v1 >= v3) {
                                    v58 := GenS10{F0: v1}
                                    if (42 >= -9223372036854775808) {
                                        v59 := GenS4{F0: "foo bar"}
                                        v6 -= (0.1 + float64(0.1))
                                        if false {
                                            fmt.Println(v1, v6, v1)
                                            for fi60 := 0; fi60 < 13; fi60++ {
                                                v57 = v57
                                                for fi61 := 0; fi61 < 16; fi61++ {
                                                    switch ("hello" + v57) {
                                                    case "foo bar":
                                                        v62 := GenS11{F0: 1, F1: (0 != 194), F2: 18446744073709551615}
                                                        _ = v62
                                                    default:
                                                        var v63 uint16 = (v4 >> ((221) & 63))
                                                        _ = v63
                                                    }
                                                    _ = fi61
                                                }
                                                _ = fi60
                                            }
                                        }
                                        _ = v59
                                    } else {
                                        fmt.Println(v57, v6)
                                    }
                                    fmt.Println(v3, v1)
                                    _ = v58
                                }
                            }
                            v6 = v6
                            _ = v57
                        } else {
                            for fi64 := 0; fi64 < 15; fi64++ {
                                var v65 uint = 4294967295
                                v66 := GenS12{F0: (fi64 - 9223372036854775807), F1: (v4 % ((1) | 1)), F2: -32768}
                                _ = fi64
                                _ = v65
                                _ = v66
                            }
                        }
                        for fi67 := 0; fi67 < 3; fi67++ {
                            fmt.Println(v1, v3)
                            if false {
                                fmt.Println(fi67, v1, fi67)
                                v68 := GenS12{F0: fi67, F1: (^(v4)), F2: int16(v4)}
                                _ = v68
                            }
                            _ = fi67
                        }
                        _ = v56
                    } else {
                        v4 = 127
                    }
                }
            }
        }
    case "a":
        var v69 string = ("\n\t" + "")
        _ = v69
    default:
        for fi70 := 0; fi70 < 11; fi70++ {
            var v71 uint = 2
            _ = fi70
            _ = v71
        }
    }
    v2 = append(v2, -9223372036854775808)
    for fi72 := 0; fi72 < 19; fi72++ {
        for fi73 := 0; fi73 < 15; fi73++ {
            switch ("foo bar" + "hello") {
            case "foo bar":
                v1 -= -128
            }
            _ = fi73
        }
        _ = fi72
    }
    _ = len(v5)
    switch (("\n\t" + "hello") + "\n\t") {
    case "a":
        v3 = (v3 | (-(v3)))
    case "\n\t":
        var v74 uint32 = 346
        _ = v74
    case "":
        _ = len(v2)
    }
    switch (("foo bar" + "") + ("foo bar" + ("\n\t" + ("\n\t" + "a")))) {
    case "\n\t":
        if false {
            var v75 int16 = -32768
            _ = v75
        } else {
            v4 -= v4
            v6 += float64(((-2.25 * 1e10) / 1.5))
            _ = len(v5)
        }
    }
    v76 := v5[-741]
    v6 = v6
    if ((-32768 <= 32767) || (3.14159 >= (1.5 / float64(-2.25)))) {
        switch "foo bar" {
        case "hello":
            v77 := v5[v76]
            _ = v77
        case "foo bar":
            v78 := GenS4{F0: (("" + "a") + "hello")}
            _ = v78
        }
        var v79 float64 = (v6 / v6)
        if false {
            switch "a" {
            case "":
                for fi80 := 0; fi80 < 20; fi80++ {
                    if false {
                        v6 = v79
                        var v81 int64 = 9223372036854775807
                        var v82 int = 0
                        _ = v81
                        _ = v82
                    }
                    if len(v2) > 0 {
                        v2[0] = int(v4)
                    }
                    var v83 int32 = -843
                    _ = fi80
                    _ = v83
                }
            case "foo bar":
                fmt.Println(v4)
            case "hello":
                v6 = (3.14159 * v6)
            }
            if (0 >= 51) {
                var v84 float64 = (v6 * v6)
                _ = v84
            }
            for fi85 := 0; fi85 < 17; fi85++ {
                var v86 string = ("a" + "foo bar")
                switch (^(fi85)) {
                case 493:
                    fmt.Println(v4, v1)
                case 32767:
                    for fi87 := 0; fi87 < 11; fi87++ {
                        var v88 float64 = 1.5
                        var v89 int8 = v3
                        if ("hello" > "a") {
                            var v90 float32 = 1e10
                            _ = v90
                        }
                        _ = fi87
                        _ = v88
                        _ = v89
                    }
                }
                v4 = v4
                _ = fi85
                _ = v86
            }
        }
        _ = v79
    }
    var v91 bool = false
    if (929 == 0) {
        v1 = (-(v1))
        switch "foo bar" {
        case "foo bar":
            for fi92 := 0; fi92 < 8; fi92++ {
                v76 = v76
                _ = fi92
            }
        case "\n\t":
            v93 := GenS5{F0: 127, F1: 65535}
            _ = v93
        case "a":
            for fi94 := 0; fi94 < 4; fi94++ {
                v6 += 3.14159
                _ = fi94
            }
        default:
            var v95 float32 = float32(v6)
            _ = v95
        }
        var v96 int64 = 1
        _ = v96
    } else {
        var v97 string = "\n\t"
        v6 = v6
        if ((v76 * v76) == 128) {
            v98 := GenS10{F0: (v1 | 0)}
            v2 = append(v2, v76)
            v97 = v97
            _ = v98
        }
        _ = v97
    }
    v76 = (^(v76))
    for fi99 := 0; fi99 < 8; fi99++ {
        fmt.Println(v3, v1)
        var v100 int = -172
        switch "hello" {
        case "\n\t":
            var v101 bool = (!(int32(fi99) != 2147483647))
            _ = v101
        case "":
            var v102 float64 = (float64(0.1) - float64(0.1))
            _ = v102
        case "a":
            for fi103 := 0; fi103 < 13; fi103++ {
                var v104 float32 = (0.1 * 0.1)
                var v105 int = -9223372036854775808
                _ = fi103
                _ = v104
                _ = v105
            }
        }
        _ = fi99
        _ = v100
    }
    v2 = append(v2, (^(v76)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v91
}
