package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint8; F1 string; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int16 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]int)
    var v3 float32 = -2.25
    var v4 int = 74
    var v5 int16 = -78
    var v6 uint64 = 483
    v7 := []int{v4, v4}
    if (1 <= -1) {
        v4 = v4
        switch "foo bar" {
        case "\n\t":
            var v8 int8 = 50
            _ = v8
        case "a":
            v9 := GenS1{F0: 100}
            _ = v9
        case "":
            v10 := GenS1{F0: 42}
            _ = v10
        default:
            fmt.Println(v6, v4, v6)
        }
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            switch "\n\t" {
            case "":
                if len(v7) > 0 {
                    v7[0] = v4
                }
            case "\n\t":
                var v13 int64 = 800
                _ = v13
            case "foo bar":
                var v14 string = (("a" + "\n\t") + ("foo bar" + "\n\t"))
                _ = v14
            default:
                if true {
                    v15 := v2[fi12]
                    _ = v15
                } else {
                    fmt.Println(v3, v3, fi11)
                    switch "" {
                    case "":
                        fmt.Println(fi11, v3, v1)
                    default:
                        switch (fi11 - (fi11 + -747)) {
                        case 883:
                            if len(v7) > 0 {
                                v7[0] = 128
                            }
                        default:
                            _ = len(v7)
                        }
                    }
                }
            }
            v16 := GenS2{F0: 0, F1: 2, F2: -441}
            v17 := GenS3{F0: 255, F1: 454}
            _ = fi12
            _ = v16
            _ = v17
        }
        if false {
            var v18 uint = 1
            switch "foo bar" {
            case "\n\t":
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 uint = (v18 / (((v18 + v18)) | 1))
                    _ = fi19
                    _ = v20
                }
            case "":
                var v21 string = ("\n\t" + "foo bar")
                _ = v21
            case "foo bar":
                var v22 float32 = v3
                _ = v22
            }
            v1 += v1
            _ = v18
        } else {
            if false {
                v5 += v5
                v23 := GenS4{F0: 0, F1: ("foo bar" + "foo bar"), F2: "foo bar"}
                for fi24 := 0; fi24 < 11; fi24++ {
                    v4 += fi11
                    _ = fi24
                }
                _ = v23
            }
        }
        _ = fi11
    }
    delete(v2, (v4 ^ (-(v4))))
    v25 := GenS2{F0: 1, F1: 255, F2: 0}
    switch ("a" + "\n\t") {
    case "a":
        fmt.Println(v6, v4)
    case "foo bar":
        if ("hello" < ("hello" + "hello")) {
            switch v4 {
            case 9223372036854775807:
                var v26 int16 = -226
                _ = v26
            }
        } else {
            var v27 float64 = 1e10
            var v28 uint = 65535
            _ = v27
            _ = v28
        }
    case "\n\t":
        v1 = float64(float32(float64(v3)))
    }
    switch v4 {
    case 9223372036854775807:
        v3 = v3
    case 1:
        if (1 < 7) {
            fmt.Println(v3, v4)
            var v29 int64 = -9223372036854775808
            _ = v29
        } else {
            for fi30 := 0; fi30 < 17; fi30++ {
                fmt.Println(v5, fi30, v1)
                _ = fi30
            }
            if true {
                v31 := GenS1{F0: 127}
                v2[v4] = (^(v4))
                var v32 int64 = -9223372036854775808
                _ = v31
                _ = v32
            } else {
                if false {
                    var v33 uint8 = 127
                    _ = v33
                } else {
                    v34 := GenS4{F0: 2, F1: ("" + "a"), F2: "\n\t"}
                    v1 -= ((v1 - v1) * 1e10)
                    _ = v34
                }
            }
        }
    default:
        for fi35 := 0; fi35 < 19; fi35++ {
            var v36 uint = 0
            v3 = v3
            _ = fi35
            _ = v36
        }
    }
    v1 = (-2.25 * (v1 / v1))
    switch 0 {
    case 0:
        v37 := GenS3{F0: 18446744073709551615, F1: 541}
        _ = v37
    case 557:
        if len(v7) > 0 {
            v7[0] = v4
        }
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        v3 -= ((float32(v1) / 1e10) + 0.0)
        v39 := GenS4{F0: 127, F1: ("foo bar" + ("\n\t" + (("\n\t" + "") + ""))), F2: "\n\t"}
        _ = fi38
        _ = v39
    }
    if true {
        for fi40 := 0; fi40 < 16; fi40++ {
            v5 = (v5 % ((v5) | 1))
            if (65535 == 2147483648) {
                var v41 int8 = 60
                fmt.Println(v1, v41, v3)
                _ = v41
            } else {
                if false {
                    var v42 uint = uint(v6)
                    if true {
                        var v43 int8 = 127
                        v6 += 4294967295
                        switch v4 {
                        case 127:
                            fmt.Println(v6, v5, fi40)
                        case 42:
                            if true {
                                var v44 int16 = (v5 / ((v5) | 1))
                                _ = v44
                            }
                        case 0:
                            for fi45 := 0; fi45 < 5; fi45++ {
                                fmt.Println(v3, v43)
                                _ = fi45
                            }
                        }
                        _ = v43
                    }
                    _ = v42
                }
            }
            var v46 int64 = 42
            _ = fi40
            _ = v46
        }
        for fi47 := 0; fi47 < 8; fi47++ {
            fmt.Println(v5)
            _ = fi47
        }
        v48 := GenS3{F0: 255, F1: 1}
        _ = v48
    } else {
        for fi49 := 0; fi49 < 6; fi49++ {
            switch fi49 {
            case -2147483648:
                v50 := GenS3{F0: 1, F1: 777}
                _ = v50
            case 1:
                switch v4 {
                case 9223372036854775807:
                    v51 := GenS2{F0: 435, F1: 42, F2: -186}
                    _ = v51
                case -613:
                    switch -2147483648 {
                    case 42:
                        v4 += (v4 & (fi49 % (((v4 / ((1) | 1))) | 1)))
                    }
                case 128:
                    if (0.0 <= v1) {
                        fmt.Println(v6, v4)
                        fmt.Println(fi49)
                        fmt.Println(fi49, v4, v3)
                    } else {
                        v4 = (fi49 % (((v4 - v4)) | 1))
                        v5 -= v5
                        switch "foo bar" {
                        case "foo bar":
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        case "hello":
                            v4 = fi49
                        case "":
                            v3 += (((float32((v1 + v1)) / (-2.25 - v3)) + v3) * (1.5 + 1e10))
                        }
                    }
                }
            default:
                if true {
                    if (true && true) {
                        v2[(fi49 + (fi49 << ((365) & 63)))] = 9223372036854775807
                    } else {
                        v1 -= (1.5 - v1)
                        delete(v2, v4)
                    }
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v5, v4)
                    case "":
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            var v52 uint8 = 226
                            _ = v52
                        case "":
                            var v53 float32 = ((v3 - (1e10 - v3)) / v3)
                            _ = v53
                        default:
                            v3 = 3.14159
                        }
                    case "hello":
                        for fi54 := 0; fi54 < 8; fi54++ {
                            _ = len(v2)
                            if (!false) {
                                v55 := v2[(v4 | 32767)]
                                _ = v55
                            }
                            v3 += (0.0 / (v3 * v3))
                            var v56 float32 = float32(v1)
                            v6 += (v6 - v6)
                            _ = fi54
                            _ = v56
                        }
                    }
                } else {
                    for fi57 := 0; fi57 < 14; fi57++ {
                        v58 := GenS5{F0: 65535}
                        _ = len(v2)
                        v6 += 0
                        _ = fi57
                        _ = v58
                    }
                    switch (fi49 % ((v4) | 1)) {
                    case -1:
                        var v59 int8 = 42
                        _ = v59
                    case 128:
                        var v60 int16 = (v5 / ((42) | 1))
                        _ = v60
                    case 2147483647:
                        v61 := GenS5{F0: 199}
                        _ = v61
                    }
                }
            }
            v62 := GenS3{F0: 255, F1: 255}
            _ = fi49
            _ = v62
        }
        if (42 == 489) {
            var v63 uint16 = 2
            _ = v63
        } else {
            var v64 int64 = 9223372036854775807
            _ = v64
        }
    }
    var v65 uint16 = 42
    v66 := GenS6{F0: false}
    if true {
        switch (("" + "\n\t") + "hello") {
        case "foo bar":
            var v67 int8 = 1
            _ = v67
        case "\n\t":
            if (("hello" + "a") <= ("a" + "\n\t")) {
                delete(v2, 0)
            }
        case "hello":
            _ = len(v7)
        default:
            fmt.Println(v5, v3)
        }
        v68 := GenS7{F0: (^(v6)), F1: (-(v5))}
        _ = v68
    } else {
        delete(v2, v4)
        var v69 bool = (v5 >= v5)
        _ = v69
    }
    switch v4 {
    case 32768:
        if false {
            if (127 >= 22) {
                for fi70 := 0; fi70 < 20; fi70++ {
                    v65 += (^(v65))
                    switch (v4 * (fi70 & -9223372036854775808)) {
                    case -1:
                        var v71 uint16 = v65
                        _ = v71
                    case 1:
                        switch "foo bar" {
                        case "\n\t":
                            v2[fi70] = (-(fi70))
                        case "hello":
                            v4 = fi70
                        case "a":
                            var v72 uint = 255
                            _ = v72
                        }
                    default:
                        fmt.Println(v4, v65, v5)
                    }
                    switch v4 {
                    case 32768:
                        v6 += (^(v6))
                    case 0:
                        v1 = float64(v3)
                    case -147:
                        for fi73 := 0; fi73 < 2; fi73++ {
                            var v74 float64 = 3.14159
                            if len(v7) > 0 {
                                v7[0] = -531
                            }
                            _ = fi73
                            _ = v74
                        }
                    default:
                        var v75 int16 = -478
                        _ = v75
                    }
                    _ = fi70
                }
                for fi76 := 0; fi76 < 17; fi76++ {
                    switch (^(fi76)) {
                    case 32768:
                        switch "foo bar" {
                        case "\n\t":
                            v65 += v65
                        case "foo bar":
                            var v77 float32 = 0.1
                            _ = v77
                        case "a":
                            v4 += v4
                        default:
                            var v78 uint16 = v65
                            _ = v78
                        }
                    case -9223372036854775808:
                        v4 = fi76
                    case 959:
                        var v79 uint8 = 255
                        _ = v79
                    }
                    var v80 uint8 = 57
                    fmt.Println(v3)
                    _ = fi76
                    _ = v80
                }
            } else {
                if len(v7) > 0 {
                    v7[0] = (-(v4))
                }
                switch ("\n\t" + "hello") {
                case "\n\t":
                    switch "hello" {
                    case "hello":
                        if false {
                            v5 = (v5 & 32767)
                        }
                    default:
                        var v81 int = -1
                        _ = v81
                    }
                case "foo bar":
                    v3 = v3
                default:
                    var v82 uint64 = v6
                    _ = v82
                }
            }
        }
    case -532:
        for fi83 := 0; fi83 < 3; fi83++ {
            for fi84 := 0; fi84 < 18; fi84++ {
                v65 += v65
                var v85 int16 = 1
                _ = fi84
                _ = v85
            }
            v65 -= (v65 << ((435) & 63))
            var v86 uint64 = v6
            _ = fi83
            _ = v86
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v7
}
