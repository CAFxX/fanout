package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint64; F2 int }
type GenS6 struct { F0 int16; F1 string; F2 uint16 }
type GenS7 struct { F0 int16; F1 string; F2 int64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 465
    var v3 float32 = 1e10
    v4 := []string{"a", "hello", "\n\t"}
    var v5 uint16 = 0
    switch (("\n\t" + "") + "a") {
    case "hello":
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 int = (fi6 ^ -2147483648)
            v3 = (v3 / v3)
            _ = fi6
            _ = v7
        }
    default:
        v8 := GenS1{F0: "hello", F1: "", F2: 128}
        _ = v8
    }
    switch 128 {
    case 4294967295:
        for fi9 := 0; fi9 < 2; fi9++ {
            v3 = -2.25
            _ = fi9
        }
    }
    if (0 < 421) {
        if (!(!((0 <= 18446744073709551615) && (!true)))) {
            fmt.Println(v3, v5)
        }
        v3 = v3
        v10 := v1[65535]
        _ = v10
    }
    if false {
        var v11 int = 2147483647
        for fi12 := 0; fi12 < 6; fi12++ {
            if true {
                v3 += (v3 * v3)
                var v13 string = ""
                _ = v13
            } else {
                for fi14 := 0; fi14 < 11; fi14++ {
                    switch ("hello" + "a") {
                    case "hello":
                        var v15 int = fi14
                        _ = v15
                    case "\n\t":
                        v16 := GenS2{F0: 888, F1: true, F2: 665}
                        _ = v16
                    case "a":
                        var v17 uint16 = 836
                        _ = v17
                    }
                    switch "\n\t" {
                    case "\n\t":
                        if (((1000 <= 1) || (2 < 65535)) || false) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if ((!(!false)) || (!(v3 < v3))) {
                                switch v11 {
                                case 1:
                                    v18 := GenS2{F0: 4294967295, F1: (34 >= -29), F2: 0}
                                    _ = v18
                                case -9223372036854775808:
                                    var v19 bool = false
                                    _ = v19
                                case 8:
                                    v5 += v5
                                default:
                                    _ = len(v1)
                                }
                                v4 = append(v4, "foo bar")
                                if false {
                                    v3 = v3
                                    if true {
                                        v2 = v2
                                        v20 := GenS2{F0: 1, F1: ((504 > 65535) || false), F2: 0}
                                        if true {
                                            v4 = append(v4, "a")
                                            switch ("\n\t" + "foo bar") {
                                            case "\n\t":
                                                v2 -= (^(v2))
                                            default:
                                                switch (^(fi12)) {
                                                case 0:
                                                    v2 = (v2 + (v2 ^ (v2 * (v2 + 9223372036854775807))))
                                                case -2147483648:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                default:
                                                    v11 = 127
                                                }
                                            }
                                            var v21 uint = 18446744073709551615
                                            _ = v21
                                        } else {
                                            v11 = 128
                                        }
                                        _ = v20
                                    }
                                    v4 = append(v4, "foo bar")
                                    fmt.Println(v2)
                                } else {
                                    v11 += -567
                                    switch ((("\n\t" + "\n\t") + "") + "hello") {
                                    case "a":
                                        v22 := GenS1{F0: "", F1: "\n\t", F2: 242}
                                        _ = v22
                                    case "\n\t":
                                        if len(v4) > 0 {
                                            v4[0] = "\n\t"
                                        }
                                    }
                                    _ = len(v4)
                                    v3 += v3
                                }
                            } else {
                                var v23 uint16 = v5
                                v24 := GenS1{F0: "", F1: "foo bar", F2: 0}
                                v25 := GenS1{F0: "foo bar", F1: ("hello" + "\n\t"), F2: 246}
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        } else {
                            if (v5 == 65535) {
                                if (!((8 != (v2 / (((-(v2))) | 1))) || (v5 != v5))) {
                                    v26 := GenS1{F0: "", F1: "hello", F2: 0}
                                    if ((!(!false)) || true) {
                                        if (1 > 2) {
                                            var v27 bool = (0.0 < (-2.25 / 0.1))
                                            switch "a" {
                                            case "hello":
                                                v4 = append(v4, ("hello" + "a"))
                                            case "a":
                                                delete(v1, 608)
                                            case "\n\t":
                                                var v28 int = (^(v11))
                                                _ = v28
                                            }
                                            if (v5 >= v5) {
                                                v5 += 2
                                                fmt.Println(v27)
                                                v27 = ((("\n\t" + ("foo bar" + "hello")) + "hello") == "hello")
                                            } else {
                                                v3 = ((3.14159 * ((v3 - float32(-2.25)) - v3)) - v3)
                                                fmt.Println(fi12, v27)
                                                var v29 uint = 4294967295
                                                v30 := GenS2{F0: 0, F1: ((1 < 18446744073709551615) || (v27 || v27)), F2: 1}
                                                if len(v4) > 0 {
                                                    v4[0] = "hello"
                                                }
                                                _ = v29
                                                _ = v30
                                            }
                                            var v31 bool = v27
                                            _ = v27
                                            _ = v31
                                        } else {
                                            fmt.Println(v3)
                                            v32 := GenS3{F0: 0, F1: 692}
                                            if (v11 >= (fi12 ^ -2147483648)) {
                                                switch "" {
                                                case "\n\t":
                                                    if false {
                                                        if ((1 > 18446744073709551615) || ((v5 * v5) == 1)) {
                                                            var v33 uint8 = 255
                                                            var v34 int8 = int8(v33)
                                                            switch ((("hello" + "") + "hello") + (("a" + "a") + "\n\t")) {
                                                            case "":
                                                                v33 = v33
                                                            default:
                                                                v35 := v1[2147483648]
                                                                _ = v35
                                                            }
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                        var v36 int = (fi14 % ((fi12) | 1))
                                                        switch "hello" {
                                                        case "":
                                                            v11 = (fi14 - fi12)
                                                        case "hello":
                                                            v3 = (((v3 + (0.0 * 0.0)) / ((v3 + float32(0.0)) * float32(((0.0 + (1.5 * -2.25)) * 0.0)))) * 1e10)
                                                        case "foo bar":
                                                            v3 = float32(0.1)
                                                        default:
                                                            v37 := GenS3{F0: 42, F1: 0}
                                                            _ = v37
                                                        }
                                                        _ = v36
                                                    }
                                                }
                                            }
                                            _ = v32
                                        }
                                        if (true && false) {
                                            fmt.Println(fi14)
                                        }
                                        v38 := GenS2{F0: 0, F1: true, F2: 2}
                                        _ = v38
                                    }
                                    fmt.Println(v2, v11, v2)
                                    _ = v26
                                }
                                var v39 float64 = (1.5 - float64(-2.25))
                                _ = v39
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v5 = v5
                            }
                            if false {
                                v5 -= v5
                                v11 = fi14
                            } else {
                                switch (fi12 + v11) {
                                case -1:
                                    if true {
                                        v11 = -133
                                        var v40 uint16 = v5
                                        if false {
                                            fmt.Println(fi12, v2)
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = ("" + "a")
                                            }
                                        }
                                        _ = v40
                                    } else {
                                        var v41 uint64 = 352
                                        if (v2 == (-(v2))) {
                                            var v42 uint = 18446744073709551615
                                            _ = v42
                                        }
                                        _ = v41
                                    }
                                case 1:
                                    v11 += v11
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = ("a" + "hello")
                                    }
                                }
                            }
                        }
                    }
                    v11 -= 127
                    _ = fi14
                }
                var v43 uint = 18446744073709551615
                v2 -= 9223372036854775807
                _ = v43
            }
            _ = fi12
        }
        _ = v11
    } else {
        if false {
            delete(v1, 275)
            v44 := GenS1{F0: "\n\t", F1: "foo bar", F2: 127}
            _ = v44
        } else {
            v45 := GenS3{F0: 578, F1: 4294967295}
            var v46 uint16 = v5
            for fi47 := 0; fi47 < 14; fi47++ {
                if (!(127 > 177)) {
                    var v48 int8 = -28
                    v49 := GenS4{F0: (^(fi47))}
                    var v50 int64 = (^(v2))
                    _ = v48
                    _ = v49
                    _ = v50
                }
                for fi51 := 0; fi51 < 4; fi51++ {
                    v5 = v46
                    v52 := GenS3{F0: 42, F1: 4294967295}
                    _ = fi51
                    _ = v52
                }
                _ = fi47
            }
            _ = v45
            _ = v46
        }
        var v53 uint64 = 602
        _ = v53
    }
    v54 := GenS3{F0: uint32(v5), F1: 585}
    v3 = 1e10
    if (384 >= 815) {
        var v55 string = ((("a" + "") + "foo bar") + "")
        if true {
            for fi56 := 0; fi56 < 18; fi56++ {
                switch 969 {
                case -601:
                    if (!true) {
                        fmt.Println(fi56)
                        fmt.Println(fi56, v55)
                        if true {
                            v3 = float32(1e10)
                            switch (fi56 + (fi56 % ((262) | 1))) {
                            case 9223372036854775807:
                                fmt.Println(v55)
                            }
                        }
                    } else {
                        _ = len(v1)
                        if false {
                            v57 := GenS1{F0: "foo bar", F1: v55, F2: 0}
                            if true {
                                v58 := GenS4{F0: int(v2)}
                                _ = v58
                            } else {
                                v55 = ("\n\t" + (v55 + ("hello" + (v55 + v55))))
                            }
                            v4 = append(v4, ("a" + (v55 + "")))
                            _ = v57
                        } else {
                            if true {
                                switch (^(fi56)) {
                                case 255:
                                    if (!(true || (0 == 95))) {
                                        if false {
                                            v59 := GenS5{F0: 543, F1: 220, F2: (fi56 + fi56)}
                                            fmt.Println(v55, v55)
                                            v55 = "a"
                                            _ = v59
                                        }
                                        v60 := GenS6{F0: -32768, F1: (v55 + "\n\t"), F2: v5}
                                        _ = v60
                                    } else {
                                        if true {
                                            for fi61 := 0; fi61 < 17; fi61++ {
                                                var v62 uint8 = 221
                                                v63 := GenS3{F0: 2147483647, F1: 4294967295}
                                                _ = fi61
                                                _ = v62
                                                _ = v63
                                            }
                                        } else {
                                            var v64 int32 = 143
                                            _ = v64
                                        }
                                    }
                                }
                                if false {
                                    v65 := GenS7{F0: 716, F1: v55, F2: v2}
                                    if (517 <= (v2 % (((^(v2))) | 1))) {
                                        var v66 int32 = -1
                                        for fi67 := 0; fi67 < 20; fi67++ {
                                            var v68 float32 = v3
                                            _ = fi67
                                            _ = v68
                                        }
                                        for fi69 := 0; fi69 < 16; fi69++ {
                                            if true {
                                                var v70 string = v55
                                                _ = v70
                                            }
                                            fmt.Println(v55)
                                            fmt.Println(v55)
                                            _ = fi69
                                        }
                                        _ = v66
                                    } else {
                                        fmt.Println(v55, v3, v55)
                                    }
                                    _ = v65
                                } else {
                                    v3 = (v3 * (v3 + (float32(1e10) + (float32((1.5 / 1.0)) - v3))))
                                    var v71 int16 = -412
                                    v71 = (v71 | (v71 ^ v71))
                                    _ = v71
                                }
                            }
                        }
                    }
                }
                _ = fi56
            }
        }
        _ = v55
    }
    v2 += v2
    v3 = (v3 / v3)
    v72 := GenS7{F0: 0, F1: "\n\t", F2: (^(v2))}
    v4 = append(v4, ((("a" + "hello") + ("\n\t" + "")) + ("a" + "foo bar")))
    var v73 string = "foo bar"
    for fi74 := 0; fi74 < 2; fi74++ {
        switch fi74 {
        case -72:
            v73 = v73
        case 32768:
            fmt.Println(v5)
        case 4294967295:
            for fi75 := 0; fi75 < 20; fi75++ {
                v2 = v2
                fmt.Println(fi74, v5)
                v76 := GenS8{F0: (fi74 >> ((934) & 63))}
                switch (^(fi74)) {
                case 127:
                    if true {
                        v73 = v73
                        var v77 uint64 = 42
                        v77 += (^(v77))
                        _ = v77
                    }
                case 318:
                    if true {
                        var v78 int64 = v2
                        v79 := GenS9{F0: (!false)}
                        _ = v78
                        _ = v79
                    }
                case -290:
                    if true {
                        fmt.Println(fi74)
                        var v80 int8 = 100
                        if len(v4) > 0 {
                            v4[0] = ((v73 + (("" + v73) + "")) + v73)
                        }
                        v2 = v2
                        _ = len(v4)
                        _ = v80
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                }
                v5 = (v5 + v5)
                var v81 int8 = 42
                _ = fi75
                _ = v76
                _ = v81
            }
        default:
            fmt.Println(v5)
        }
        switch ((v73 + v73) + (v73 + v73)) {
        case "\n\t":
            v5 -= (^(v5))
        case "a":
            fmt.Println(fi74)
        case "foo bar":
            var v82 int = fi74
            _ = v82
        }
        switch v73 {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi74
    }
    var v83 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v72
    _ = v73
    _ = v83
}
