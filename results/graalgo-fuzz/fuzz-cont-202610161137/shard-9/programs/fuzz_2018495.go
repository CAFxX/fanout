package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int }

func main() {
    var v1 uint32 = 2
    var v2 float32 = 0.0
    v3 := []int{-1, -9223372036854775808, 1, 32767}
    var v4 int32 = -240
    if false {
        v3 = append(v3, 127)
    } else {
        if len(v3) > 0 {
            v3[0] = -1
        }
        _ = len(v3)
    }
    v4 = (v4 - -2147483648)
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 float32 = v2
        for fi7 := 0; fi7 < 8; fi7++ {
            if len(v3) > 0 {
                v3[0] = (^(fi5))
            }
            fmt.Println(fi7, fi5)
            switch "foo bar" {
            case "foo bar":
                var v8 int32 = v4
                _ = v8
            case "\n\t":
                v2 = float32((-2.25 / 1.5))
            }
            v6 += v2
            if true {
                if true {
                    var v9 uint16 = 996
                    var v10 string = ("foo bar" + "\n\t")
                    switch (fi7 - (fi7 | fi7)) {
                    case 511:
                        if (4294967295 > 535) {
                            _ = len(v3)
                            if (!(-743 == int(v1))) {
                                if (-1 < -1) {
                                    if true {
                                        v11 := GenS1{F0: -90}
                                        v12 := GenS1{F0: -1}
                                        _ = v11
                                        _ = v12
                                    } else {
                                        var v13 float64 = (((((0.0 * (-2.25 * ((1.5 + 1.5) + 1e10))) + (0.0 - (float64((v6 + 0.1)) / 0.1))) / 1.5) * (1e10 / ((-2.25 * ((0.0 / (1e10 - 0.1)) - 0.0)) - 1.5))) / 1.5)
                                        if ((v1 ^ v1) != v1) {
                                            v9 -= (^(v9))
                                        }
                                        fmt.Println(fi7, v2)
                                        if (fi5 <= fi5) {
                                            if (0 >= 654) {
                                                v14 := GenS1{F0: 0}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v14
                                            } else {
                                                switch (fi7 >> ((0) & 63)) {
                                                case 2147483647:
                                                    fmt.Println(v6, v13)
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = (^(fi7))
                                                }
                                            }
                                            v9 -= v9
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        _ = v13
                                    }
                                    v15 := GenS1{F0: -1}
                                    if len(v3) > 0 {
                                        v3[0] = (fi7 - (-(fi7)))
                                    }
                                    _ = v15
                                }
                            } else {
                                var v16 float64 = (3.14159 + float64(v6))
                                _ = v16
                            }
                        } else {
                            if (1e10 <= 1.5) {
                                fmt.Println(v10, v1, v6)
                                v17 := GenS1{F0: -50}
                                var v18 int32 = (v4 | (-(v4)))
                                switch (v10 + (v10 + v10)) {
                                case "\n\t":
                                    switch -687 {
                                    case 2147483647:
                                        if false {
                                            if (4294967295 <= 157) {
                                                switch (v10 + "a") {
                                                case "a":
                                                    switch v10 {
                                                    case "hello":
                                                        var v19 float32 = 3.14159
                                                        _ = v19
                                                    }
                                                default:
                                                    var v20 string = ((v10 + (v10 + (v10 + v10))) + v10)
                                                    _ = v20
                                                }
                                            }
                                            v1 += (v1 >> ((362) & 63))
                                            if len(v3) > 0 {
                                                v3[0] = (fi7 | fi7)
                                            }
                                            if (false && true) {
                                                switch ("" + (v10 + v10)) {
                                                case "":
                                                    fmt.Println(v18, v10, v1)
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = (fi7 + (fi5 << ((65535) & 63)))
                                                }
                                                v21 := GenS2{F0: 4294967295, F1: fi5}
                                                _ = v21
                                            }
                                        }
                                        var v22 uint8 = 0
                                        _ = v22
                                    }
                                case "hello":
                                    if true {
                                        if (0.1 == float64(0.1)) {
                                            fmt.Println(v18, fi7, v18)
                                            v23 := GenS2{F0: 42, F1: fi5}
                                            var v24 int64 = 1
                                            v10 = "\n\t"
                                            _ = v23
                                            _ = v24
                                        } else {
                                            v4 = v4
                                            _ = len(v3)
                                            v3 = append(v3, (fi7 + -914))
                                        }
                                        _ = len(v3)
                                    } else {
                                        switch -1 {
                                        case 4294967295:
                                            switch 42 {
                                            case 2147483647:
                                                fmt.Println(v4, v6, v2)
                                            }
                                        case 2147483647:
                                            v4 -= v4
                                        default:
                                            v18 -= (v18 & (v18 * (^(v18))))
                                        }
                                        if true {
                                            if false {
                                                v1 = 4294967295
                                            } else {
                                                var v25 int = (fi5 + fi5)
                                                _ = v25
                                            }
                                            var v26 uint32 = 0
                                            _ = v26
                                        }
                                    }
                                    v27 := GenS1{F0: -128}
                                    _ = v27
                                case "foo bar":
                                    v2 = v6
                                }
                                _ = v17
                                _ = v18
                            }
                        }
                    default:
                        var v28 string = (("\n\t" + (v10 + v10)) + ("a" + v10))
                        _ = v28
                    }
                    _ = v10
                    _ = v9
                }
                v1 = 668
            } else {
                fmt.Println(v6)
                switch fi5 {
                case 1:
                    if true {
                        switch (fi5 ^ (fi7 - fi5)) {
                        case 32767:
                            v6 += float32(0.0)
                        case -227:
                            var v29 uint32 = 42
                            _ = v29
                        case 4294967295:
                            switch "\n\t" {
                            case "\n\t":
                                if (-1 > 1000) {
                                    if (42 == 0) {
                                        var v30 int64 = 42
                                        var v31 uint16 = 255
                                        v32 := GenS2{F0: 26, F1: fi7}
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    } else {
                                        if true {
                                            v6 = (v6 / v2)
                                            switch "hello" {
                                            case "":
                                                if (100 == -1) {
                                                    var v33 int32 = (v4 >> ((65535) & 63))
                                                    if len(v3) > 0 {
                                                        v3[0] = 287
                                                    }
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "hello":
                                                        fmt.Println(v2)
                                                    case "foo bar":
                                                        v33 = v33
                                                    default:
                                                        if false {
                                                            v4 -= v33
                                                            v2 = 3.14159
                                                        }
                                                    }
                                                    _ = v33
                                                } else {
                                                    v2 -= (v2 + 3.14159)
                                                    var v34 int16 = 0
                                                    if false {
                                                        fmt.Println(v4)
                                                    } else {
                                                        switch "foo bar" {
                                                        case "a":
                                                            v6 -= 0.1
                                                        case "hello":
                                                            if (255 == 1) {
                                                                var v35 bool = true
                                                                v6 = (float32(float64(3.14159)) - 3.14159)
                                                                fmt.Println(fi7, v2)
                                                                _ = v35
                                                            }
                                                        case "foo bar":
                                                            var v36 uint = 1
                                                            _ = v36
                                                        default:
                                                            if (!false) {
                                                                fmt.Println(fi7)
                                                                fmt.Println(v1)
                                                                v6 -= (v2 / (v6 - (v2 * 3.14159)))
                                                            } else {
                                                                var v37 uint16 = 0
                                                                var v38 uint8 = 98
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                        }
                                                        fmt.Println(v34, fi7, fi5)
                                                    }
                                                    _ = v34
                                                }
                                            default:
                                                v39 := GenS3{F0: v4, F1: 947}
                                                _ = v39
                                            }
                                        }
                                        v40 := GenS4{F0: true, F1: 1}
                                        if true {
                                            var v41 string = "a"
                                            var v42 int32 = v4
                                            _ = v41
                                            _ = v42
                                        } else {
                                            var v43 int16 = -591
                                            _ = v43
                                        }
                                        _ = v40
                                    }
                                    v1 = 2147483648
                                    v1 = v1
                                    switch ("" + (("a" + ("\n\t" + "\n\t")) + "a")) {
                                    case "":
                                        var v44 uint8 = 42
                                        _ = v44
                                    default:
                                        fmt.Println(fi7, v1)
                                    }
                                }
                                v45 := GenS5{F0: 18446744073709551615}
                                _ = v45
                            case "foo bar":
                                if len(v3) > 0 {
                                    v3[0] = fi7
                                }
                            case "hello":
                                fmt.Println(v1, fi5, v1)
                            }
                        default:
                            if true {
                                var v46 uint8 = 0
                                _ = v46
                            } else {
                                v6 += v6
                                if (-891 >= -9223372036854775808) {
                                    fmt.Println(v2, v4)
                                    v1 -= v1
                                }
                            }
                        }
                        var v47 int32 = 1
                        v48 := GenS6{F0: v4}
                        _ = v47
                        _ = v48
                    } else {
                        switch (fi5 >> ((380) & 63)) {
                        case 259:
                            if false {
                                var v49 uint = 0
                                _ = v49
                            }
                        case 128:
                            var v50 int64 = int64(fi7)
                            _ = v50
                        case -1:
                            v6 = v6
                        default:
                            v6 += (v2 / v2)
                        }
                    }
                    v2 = float32(float64(((((v2 - 1.5) + (3.14159 - v2)) + v2) - ((v6 * v2) * v6))))
                case 32768:
                    v51 := GenS2{F0: 683, F1: fi5}
                    _ = v51
                }
            }
            var v52 int8 = -1
            _ = fi7
            _ = v52
        }
        v6 = (1e10 / -2.25)
        _ = fi5
        _ = v6
    }
    _ = len(v3)
    var v53 float64 = (((((float64((float32(1.5) - v2)) + (1.5 - float64(0.0))) + 1.5) * 0.1) - float64(v2)) - float64(v2))
    if (374 == 255) {
        if (true || false) {
            var v54 bool = (4294967295 < 1)
            _ = v54
        }
        v55 := GenS3{F0: v4, F1: 1000}
        _ = v55
    }
    if false {
        v4 = v4
    }
    var v56 uint = 2
    if true {
        fmt.Println(v53, v1)
        v57 := GenS1{F0: -128}
        switch 4294967295 {
        case 0:
            if (127 > uint8(v1)) {
                v58 := GenS4{F0: true, F1: 2}
                v1 -= (v1 >> ((v56) & 63))
                switch ("foo bar" + (("" + "") + "a")) {
                case "foo bar":
                    for fi59 := 0; fi59 < 6; fi59++ {
                        for fi60 := 0; fi60 < 20; fi60++ {
                            if false {
                                if false {
                                    var v61 string = ""
                                    v62 := GenS3{F0: v4, F1: -1}
                                    _ = v61
                                    _ = v62
                                }
                                var v63 uint8 = 128
                                v64 := GenS7{F0: fi60}
                                var v65 int32 = (v4 / (((^(v4))) | 1))
                                _ = v63
                                _ = v64
                                _ = v65
                            }
                            if true {
                                var v66 bool = (("" + "\n\t") > "")
                                _ = v66
                            }
                            _ = fi60
                        }
                        _ = fi59
                    }
                default:
                    v67 := GenS3{F0: v4, F1: -25}
                    _ = v67
                }
                _ = v58
            }
        case -559:
            for fi68 := 0; fi68 < 11; fi68++ {
                v2 = 3.14159
                var v69 uint16 = 255
                _ = fi68
                _ = v69
            }
        default:
            fmt.Println(v53, v53, v56)
        }
        _ = v57
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v56
}
