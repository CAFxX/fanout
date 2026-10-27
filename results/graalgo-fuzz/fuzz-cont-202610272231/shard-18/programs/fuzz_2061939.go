package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int64; F2 int32 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint8; F1 string; F2 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 uint16 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    var v4 uint8 = 255
    v5 := []int{882, 42}
    v6 := make(map[string]int)
    if ((9223372036854775807 == 4) && true) {
        v7 := GenS1{F0: 1}
        if true {
            var v8 int8 = -51
            if (((!(2 < 4294967295)) || true) && false) {
                v1 -= 473
                var v9 uint = 1
                v4 = 66
                _ = v9
            } else {
                v4 = 255
            }
            v1 -= (v1 + (v1 << ((42) & 63)))
            _ = v8
        } else {
            v1 -= v1
        }
        _ = v7
    } else {
        switch "a" {
        case "foo bar":
            if (((1.5 * 0.1) + ((float64(1e10) + 1.5) + (0.1 / ((1.5 - 0.1) / -2.25)))) < (float64(1.5) / 1.0)) {
                fmt.Println(v4)
            }
        case "\n\t":
            var v10 bool = true
            _ = v10
        case "":
            v1 += uint64(v4)
        default:
            var v11 bool = false
            _ = v11
        }
    }
    if true {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 float32 = -2.25
            var v14 int64 = -1
            if ((v4 + 0) == (v4 ^ v4)) {
                for fi15 := 0; fi15 < 7; fi15++ {
                    var v16 float32 = v13
                    v5 = append(v5, (fi15 >> ((967) & 63)))
                    fmt.Println(v4, v14)
                    v17 := GenS2{F0: 76, F1: (v14 + (v14 % (((-(v14))) | 1))), F2: 2147483647}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v13 += 0.0
                v6[("a" + ("hello" + "\n\t"))] = (fi12 & (fi12 >> ((39) & 63)))
            } else {
                v18 := GenS3{F0: 1, F1: "\n\t"}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    v1 -= 1
                    v20 := GenS1{F0: -32768}
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            switch fi21 {
            case -773:
                var v22 int16 = -903
                _ = v22
            case 0:
                switch "foo bar" {
                case "a":
                    v4 -= v4
                case "foo bar":
                    if true {
                        var v23 int8 = -38
                        v4 -= (v4 * 0)
                        v23 = (v23 - (-(v23)))
                        _ = v23
                    } else {
                        fmt.Println(v1, fi21)
                        var v24 bool = (false || (0 == 4294967295))
                        _ = v24
                    }
                case "\n\t":
                    v4 = (^(v4))
                default:
                    fmt.Println(v4, v4)
                }
            default:
                var v25 int16 = 492
                _ = v25
            }
            switch (((("hello" + "foo bar") + "hello") + "a") + "\n\t") {
            case "a":
                var v26 uint16 = 42
                _ = v26
            }
            _ = fi21
        }
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        switch "\n\t" {
        case "foo bar":
            for fi28 := 0; fi28 < 12; fi28++ {
                if (false && false) {
                    if len(v5) > 0 {
                        v5[0] = fi27
                    }
                    v29 := v6[((("" + "foo bar") + ("" + "hello")) + ("hello" + "\n\t"))]
                    switch "\n\t" {
                    case "":
                        if (!false) {
                            if true {
                                _ = len(v3)
                            } else {
                                v30 := GenS1{F0: 620}
                                switch "" {
                                case "a":
                                    v4 = (v4 << ((360) & 63))
                                }
                                v4 -= v4
                                if true {
                                    fmt.Println(v1, fi27)
                                }
                                _ = v30
                            }
                            var v31 int16 = 477
                            _ = v31
                        }
                    case "foo bar":
                        fmt.Println(fi28)
                    default:
                        if ((false && false) || (-303 > -927)) {
                            fmt.Println(fi27)
                            v32 := GenS4{F0: "", F1: 42}
                            v4 = 127
                            _ = v32
                        }
                    }
                    _ = v29
                }
                _ = fi28
            }
        case "hello":
            _ = len(v5)
        case "":
            var v33 float64 = (float64(float32(0.0)) - -2.25)
            _ = v33
        default:
            var v34 float64 = (0.0 - (-2.25 + float64(float32((3.14159 - (((0.0 - 0.1) / float64(((-2.25 - 0.1) + (1e10 * (0.0 - -2.25))))) + -2.25))))))
            _ = v34
        }
        switch 9223372036854775807 {
        case 0:
            switch fi27 {
            case 4294967295:
                if true {
                    switch "hello" {
                    case "a":
                        if true {
                            for fi35 := 0; fi35 < 9; fi35++ {
                                if ((false || false) && false) {
                                    var v36 uint8 = 46
                                    _ = len(v5)
                                    _ = v36
                                }
                                fmt.Println(fi35, fi35, fi35)
                                _ = fi35
                            }
                            for fi37 := 0; fi37 < 13; fi37++ {
                                var v38 uint16 = 736
                                v4 += (v4 * 255)
                                var v39 uint8 = 128
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                            switch (^(fi27)) {
                            case 136:
                                v40 := GenS2{F0: 0, F1: 1, F2: -2147483648}
                                _ = v40
                            case 9223372036854775807:
                                switch (-(fi27)) {
                                case 32768:
                                    if true {
                                        var v41 uint64 = uint64(v4)
                                        _ = v41
                                    } else {
                                        var v42 int16 = 460
                                        v5 = append(v5, 32768)
                                        v42 = (v42 - 42)
                                        _ = v42
                                    }
                                case 32767:
                                    v1 -= 613
                                default:
                                    v1 += v1
                                }
                            case 555:
                                var v43 uint = 65535
                                _ = v43
                            default:
                                var v44 int = (-(fi27))
                                _ = v44
                            }
                        } else {
                            var v45 int32 = 471
                            var v46 int32 = v45
                            if (!((fi27 << ((1) & 63)) == -351)) {
                                v46 += 2147483647
                                var v47 int32 = (v46 + (v45 / (((v46 * v45)) | 1)))
                                _ = v47
                            }
                            _ = v45
                            _ = v46
                        }
                    case "hello":
                        v5 = append(v5, (-(fi27)))
                    case "foo bar":
                        switch fi27 {
                        case 864:
                            for fi48 := 0; fi48 < 2; fi48++ {
                                v1 = (v1 >> ((255) & 63))
                                if (!(!false)) {
                                    switch "foo bar" {
                                    case "a":
                                        switch (fi27 << ((451) & 63)) {
                                        case 1:
                                            var v49 uint16 = 255
                                            _ = v49
                                        case 4294967295:
                                            var v50 uint32 = 1
                                            _ = v50
                                        case 0:
                                            if true {
                                                if false {
                                                    switch 0 {
                                                    case 9223372036854775807:
                                                        if ((!true) && false) {
                                                            v4 = (v4 >> ((1) & 63))
                                                            var v51 bool = (!false)
                                                            _ = v51
                                                        }
                                                    case 4294967295:
                                                        v1 -= v1
                                                    default:
                                                        var v52 float32 = (-2.25 - 0.1)
                                                        _ = v52
                                                    }
                                                    v53 := v2["hello"]
                                                    _ = v53
                                                }
                                                if true {
                                                    switch -2147483648 {
                                                    case 32768:
                                                        fmt.Println(v1, v4)
                                                    case 0:
                                                        v54 := GenS2{F0: 1, F1: -9223372036854775808, F2: -838}
                                                        _ = v54
                                                    default:
                                                        fmt.Println(fi48)
                                                    }
                                                }
                                            } else {
                                                v4 += v4
                                            }
                                        default:
                                            switch (fi27 * (fi48 << ((65535) & 63))) {
                                            case 152:
                                                _ = len(v5)
                                            case -204:
                                                fmt.Println(fi48)
                                            }
                                        }
                                    }
                                    v55 := GenS5{F0: v4, F1: "foo bar", F2: (v4 << ((2) & 63))}
                                    var v56 int64 = 9223372036854775807
                                    _ = v55
                                    _ = v56
                                } else {
                                    v57 := GenS4{F0: "", F1: 127}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v57
                                }
                                var v58 float64 = 0.0
                                _ = fi48
                                _ = v58
                            }
                        case 127:
                            if false {
                                var v59 string = ""
                                if (0 <= 436) {
                                    _ = len(v3)
                                    v59 = ((("\n\t" + v59) + v59) + v59)
                                    switch ((("hello" + v59) + v59) + ("foo bar" + v59)) {
                                    case "":
                                        fmt.Println(fi27)
                                    case "foo bar":
                                        var v60 float64 = (1e10 + 0.0)
                                        _ = v60
                                    case "hello":
                                        v1 = (v1 - 18446744073709551615)
                                    }
                                } else {
                                    fmt.Println(v1, v1)
                                }
                                _ = v59
                            } else {
                                if (v1 == v1) {
                                    var v61 float32 = (1.5 + -2.25)
                                    _ = v61
                                } else {
                                    fmt.Println(v4)
                                    for fi62 := 0; fi62 < 20; fi62++ {
                                        if ((1.5 * (float32(-2.25) + (3.14159 + float32((3.14159 - (0.0 + -2.25)))))) == 0.0) {
                                            v4 = 167
                                            v63 := GenS3{F0: 172, F1: "a"}
                                            _ = v63
                                        }
                                        _ = fi62
                                    }
                                }
                            }
                        default:
                            var v64 uint16 = 82
                            _ = v64
                        }
                    default:
                        switch (fi27 | 2147483647) {
                        case -1:
                            var v65 uint64 = v1
                            _ = v65
                        case 42:
                            v3[2147483647] = (fi27 << ((4294967295) & 63))
                        }
                    }
                    v1 += v1
                }
            case -928:
                v66 := GenS4{F0: (("\n\t" + "a") + "a"), F1: -128}
                _ = v66
            case 456:
                v67 := GenS5{F0: 42, F1: "", F2: v4}
                _ = v67
            default:
                if true {
                    switch int(v1) {
                    case -842:
                        v1 += (v1 / (((v1 & 176)) | 1))
                    }
                    var v68 int64 = -1
                    _ = v68
                }
            }
        case -2147483648:
            switch (fi27 / ((4294967295) | 1)) {
            case 335:
                v4 = v4
            }
        case 9223372036854775807:
            delete(v2, (("foo bar" + (("" + "foo bar") + "\n\t")) + ""))
        }
        _ = len(v5)
        _ = fi27
    }
    for fi69 := 0; fi69 < 18; fi69++ {
        _ = len(v3)
        _ = fi69
    }
    switch "hello" {
    case "\n\t":
        for fi70 := 0; fi70 < 4; fi70++ {
            v1 = v1
            fmt.Println(fi70, v4)
            _ = fi70
        }
    }
    var v71 float64 = (((1.5 / 3.14159) + (((3.14159 - (((((float64(-2.25) - 3.14159) * 0.1) - 1e10) * float64(-2.25)) + (1.5 - float64(((float32(-2.25) / 3.14159) + -2.25))))) + ((1.5 - ((float64(3.14159) + float64(0.0)) * -2.25)) - float64(1e10))) / 1.0)) / 1e10)
    switch ("" + "foo bar") {
    case "hello":
        v1 -= v1
    default:
        switch "\n\t" {
        case "foo bar":
            var v72 int16 = 0
            _ = v72
        default:
            v73 := GenS1{F0: -32768}
            _ = v73
        }
    }
    delete(v2, (("foo bar" + "foo bar") + "a"))
    v71 -= ((float64(((float32((v71 * v71)) + 1.5) + 1e10)) * 3.14159) + float64((((-2.25 / float32(-2.25)) / 1e10) + float32(0.1))))
    if (-530 == -32768) {
        v74 := GenS6{F0: 9223372036854775807}
        switch "hello" {
        case "foo bar":
            switch 252 {
            case -1:
                fmt.Println(v71, v1)
            }
        }
        _ = len(v3)
        _ = v74
    }
    for fi75 := 0; fi75 < 8; fi75++ {
        v3[42] = -1
        _ = fi75
    }
    var v76 float64 = v71
    if ((v71 < v71) && false) {
        var v77 uint = 42
        var v78 uint8 = v4
        _ = v77
        _ = v78
    } else {
        v79 := GenS5{F0: v4, F1: "", F2: v4}
        if (!false) {
            v1 = (v1 + 4294967295)
        }
        switch ("a" + "") {
        case "":
            for fi80 := 0; fi80 < 20; fi80++ {
                var v81 int16 = -32768
                _ = fi80
                _ = v81
            }
        case "hello":
            if len(v5) > 0 {
                v5[0] = 1
            }
        default:
            var v82 int64 = 245
            _ = v82
        }
        _ = v79
    }
    if (-9223372036854775808 > 783) {
        v83 := GenS7{F0: "", F1: 65535}
        _ = v83
    }
    var v84 uint = 408
    fmt.Println(v84)
    for fi85 := 0; fi85 < 16; fi85++ {
        v76 -= float64((3.14159 - (((-2.25 * (0.0 + 3.14159)) * 1e10) - (-2.25 + 0.1))))
        _ = fi85
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v71:", v71)
    fmt.Println("v76:", v76)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v76
    _ = v84
}
