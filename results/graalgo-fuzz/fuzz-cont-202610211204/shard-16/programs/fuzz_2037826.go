package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 uint8 }

func main() {
    v1 := []uint32{}
    v2 := []uint32{1, 4294967295, 2, 382}
    v3 := []int{0, 127, 4294967295, 9223372036854775807}
    var v4 int = 32768
    if ("" < "") {
        v4 -= (v4 << ((18446744073709551615) & 63))
        v5 := GenS1{F0: "foo bar"}
        _ = v5
    } else {
        if (true || false) {
            var v6 float64 = float64(float32(3.14159))
            switch (v4 | (v4 - v4)) {
            case 9223372036854775807:
                switch v4 {
                case -1:
                    for fi7 := 0; fi7 < 12; fi7++ {
                        var v8 uint32 = 1
                        _ = fi7
                        _ = v8
                    }
                case 469:
                    v6 -= (v6 / v6)
                case 127:
                    for fi9 := 0; fi9 < 15; fi9++ {
                        v6 += v6
                        fmt.Println(v6)
                        v6 = 1.5
                        _ = fi9
                    }
                }
            case -290:
                v4 -= (v4 / ((v4) | 1))
            case 32767:
                if ((v6 * (1.5 + (3.14159 + (0.1 * -2.25)))) >= v6) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = len(v3)
                } else {
                    v6 = float64(0.1)
                    v6 += 1e10
                    var v10 bool = false
                    _ = v10
                }
            default:
                v11 := GenS1{F0: "foo bar"}
                _ = v11
            }
            switch (v4 >> ((65535) & 63)) {
            case -46:
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 bool = (!false)
                    _ = fi12
                    _ = v13
                }
            case 720:
                switch v4 {
                case -1:
                    switch "foo bar" {
                    case "":
                        fmt.Println(v6, v4)
                    case "hello":
                        var v14 uint64 = 4294967295
                        _ = v14
                    case "a":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case -932:
                    v3 = append(v3, (v4 + v4))
                case -188:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            _ = v6
        }
        fmt.Println(v4)
        if (2 > 1) {
            if ((!(float64(float32(1.5)) <= (-2.25 / 0.1))) || (-510 < -611)) {
                if ((-1 > -2147483648) && false) {
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v4)
                    case "":
                        if true {
                            if true {
                                var v15 int32 = 42
                                var v16 string = "a"
                                if true {
                                    switch (("" + v16) + v16) {
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = 1
                                        }
                                    case "":
                                        var v17 int64 = -547
                                        _ = v17
                                    }
                                    if (2 > 18446744073709551615) {
                                        v15 += -2147483648
                                    } else {
                                        var v18 int8 = 52
                                        var v19 float32 = 3.14159
                                        _ = v18
                                        _ = v19
                                    }
                                    v4 = v4
                                } else {
                                    fmt.Println(v15, v4, v16)
                                }
                                _ = v15
                                _ = v16
                            }
                            var v20 float64 = 1e10
                            _ = v20
                        } else {
                            fmt.Println(v4)
                        }
                    }
                }
                v21 := GenS2{F0: "foo bar", F1: 169, F2: (!("" != "a"))}
                _ = v21
            } else {
                v4 = (v4 >> ((65535) & 63))
                v4 -= (v4 ^ v4)
                if (-1 < -32768) {
                    if len(v3) > 0 {
                        v3[0] = 632
                    }
                    var v22 bool = (626 >= 168)
                    switch ("foo bar" + "") {
                    case "foo bar":
                        switch ("hello" + "") {
                        case "a":
                            var v23 string = "\n\t"
                            _ = v23
                        default:
                            _ = len(v1)
                        }
                    }
                    _ = v22
                } else {
                    v24 := GenS3{F0: 1, F1: 174}
                    v25 := GenS1{F0: ""}
                    _ = v24
                    _ = v25
                }
            }
            if (42 != -385) {
                v4 += 127
            } else {
                v26 := GenS4{F0: 1}
                v4 = v4
                _ = v26
            }
        }
    }
    fmt.Println(v4)
    fmt.Println(v4)
    fmt.Println(v4)
    _ = len(v2)
    switch "" {
    case "a":
        var v27 bool = false
        _ = v27
    }
    switch v4 {
    case 32767:
        var v28 bool = (!(42 != 177))
        _ = v28
    case 231:
        for fi29 := 0; fi29 < 10; fi29++ {
            v4 += -2147483648
            var v30 uint32 = 4294967295
            _ = fi29
            _ = v30
        }
    case 128:
        if (!((!(v4 != v4)) || (!false))) {
            v4 = v4
            switch ("hello" + "hello") {
            case "\n\t":
                fmt.Println(v4)
            }
        } else {
            v4 = (v4 ^ v4)
            v31 := GenS2{F0: ((("a" + ("foo bar" + "\n\t")) + "a") + "foo bar"), F1: 42, F2: ((("" + ("foo bar" + "hello")) + "a") <= ("" + "hello"))}
            _ = v31
        }
    default:
        if (127 < 100) {
            v4 = -1
            for fi32 := 0; fi32 < 8; fi32++ {
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 float32 = 0.1
                    v35 := GenS1{F0: "foo bar"}
                    v36 := GenS4{F0: 653}
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                }
                v37 := GenS3{F0: 1000, F1: 128}
                _ = fi32
                _ = v37
            }
        } else {
            if (!(!false)) {
                var v38 uint16 = 228
                _ = v38
            } else {
                switch "\n\t" {
                case "a":
                    v4 -= (v4 ^ (-(v4)))
                default:
                    switch (v4 << ((18446744073709551615) & 63)) {
                    case 32768:
                        fmt.Println(v4)
                    case 42:
                        if true {
                            fmt.Println(v4)
                        } else {
                            if (!(!(false || (157 >= 221)))) {
                                v4 -= 127
                                if false {
                                    for fi39 := 0; fi39 < 4; fi39++ {
                                        for fi40 := 0; fi40 < 8; fi40++ {
                                            v2 = append(v2, 0)
                                            if len(v1) > 0 {
                                                v1[0] = 353
                                            }
                                            switch fi40 {
                                            case -2147483648:
                                                var v41 int32 = 1
                                                _ = v41
                                            case 344:
                                                var v42 uint = 42
                                                _ = v42
                                            }
                                            _ = fi40
                                        }
                                        v4 -= fi39
                                        _ = fi39
                                    }
                                    var v43 float32 = -2.25
                                    v43 += v43
                                    _ = v43
                                } else {
                                    if ((!true) || false) {
                                        v4 = -301
                                        if (65535 >= 414) {
                                            v44 := GenS3{F0: 1000, F1: 255}
                                            for fi45 := 0; fi45 < 8; fi45++ {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                for fi46 := 0; fi46 < 12; fi46++ {
                                                    switch (fi45 | fi45) {
                                                    case 127:
                                                        var v47 int8 = 0
                                                        _ = v47
                                                    default:
                                                        var v48 string = ""
                                                        _ = v48
                                                    }
                                                    var v49 int32 = 2147483647
                                                    if false {
                                                        _ = len(v1)
                                                        if (0 == -9223372036854775808) {
                                                            var v50 uint = 121
                                                            _ = v50
                                                        } else {
                                                            var v51 float64 = 1.5
                                                            var v52 uint64 = 1
                                                            _ = len(v3)
                                                            fmt.Println(v51, v52, v4)
                                                            fmt.Println(v51)
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                    }
                                                    _ = fi46
                                                    _ = v49
                                                }
                                                for fi53 := 0; fi53 < 18; fi53++ {
                                                    v2 = append(v2, 65535)
                                                    if (-9223372036854775808 <= 42) {
                                                        v3 = append(v3, (fi45 - (-(fi45))))
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v4 = fi53
                                                    _ = fi53
                                                }
                                                _ = fi45
                                            }
                                            if false {
                                                v3 = append(v3, v4)
                                                v54 := GenS4{F0: 964}
                                                var v55 float64 = float64(-2.25)
                                                _ = v54
                                                _ = v55
                                            }
                                            _ = v44
                                        }
                                    }
                                    switch (v4 * v4) {
                                    case -9223372036854775808:
                                        fmt.Println(v4)
                                    }
                                }
                            } else {
                                for fi56 := 0; fi56 < 3; fi56++ {
                                    for fi57 := 0; fi57 < 2; fi57++ {
                                        fmt.Println(v4)
                                        var v58 int16 = 0
                                        v59 := GenS1{F0: ("\n\t" + ("hello" + "hello"))}
                                        v60 := GenS2{F0: "\n\t", F1: 127, F2: (4294967295 == 440)}
                                        fmt.Println(fi57, fi57, fi57)
                                        _ = fi57
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                    }
                                    _ = fi56
                                }
                            }
                            switch v4 {
                            case 1:
                                switch v4 {
                                case 1:
                                    fmt.Println(v4)
                                case -122:
                                    v61 := GenS5{F0: "a", F1: ("" + "hello")}
                                    _ = v61
                                default:
                                    if (-1 <= -1) {
                                        var v62 int64 = 9223372036854775807
                                        _ = v62
                                    }
                                }
                            case -2147483648:
                                switch v4 {
                                case 727:
                                    if true {
                                        v4 = (v4 + (v4 << ((65535) & 63)))
                                        if (350 <= 42) {
                                            v4 += v4
                                        }
                                        if true {
                                            _ = len(v2)
                                            var v63 int32 = -1
                                            if (2147483648 == 65535) {
                                                for fi64 := 0; fi64 < 10; fi64++ {
                                                    v65 := GenS6{F0: -319}
                                                    fmt.Println(v63)
                                                    _ = fi64
                                                    _ = v65
                                                }
                                                switch "a" {
                                                case "\n\t":
                                                    if (42 < 0) {
                                                        fmt.Println(v63)
                                                        v4 = 0
                                                    } else {
                                                        v4 -= -9223372036854775808
                                                        for fi66 := 0; fi66 < 11; fi66++ {
                                                            var v67 float64 = 0.0
                                                            _ = fi66
                                                            _ = v67
                                                        }
                                                    }
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = 65535
                                                }
                                            } else {
                                                switch "a" {
                                                case "foo bar":
                                                    v4 = (v4 + 32768)
                                                case "a":
                                                    if false {
                                                        fmt.Println(v63, v63)
                                                        fmt.Println(v4)
                                                        for fi68 := 0; fi68 < 12; fi68++ {
                                                            v69 := GenS1{F0: "a"}
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            _ = fi68
                                                            _ = v69
                                                        }
                                                    }
                                                case "hello":
                                                    var v70 int64 = 778
                                                    _ = v70
                                                }
                                                _ = len(v1)
                                                fmt.Println(v63, v63)
                                            }
                                            _ = v63
                                        }
                                    } else {
                                        if false {
                                            fmt.Println(v4)
                                            v4 = (^(v4))
                                        } else {
                                            v1 = append(v1, 4294967295)
                                            var v71 int32 = -394
                                            _ = v71
                                        }
                                        v72 := GenS4{F0: 0}
                                        _ = v72
                                    }
                                case -433:
                                    v73 := GenS3{F0: 1, F1: 26}
                                    _ = v73
                                case 128:
                                    var v74 uint8 = 195
                                    _ = v74
                                default:
                                    var v75 int = v4
                                    _ = v75
                                }
                            case 39:
                                var v76 string = "hello"
                                _ = v76
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case 127:
                        var v77 float32 = -2.25
                        _ = v77
                    default:
                        fmt.Println(v4)
                    }
                }
            }
            if (!true) {
                _ = len(v1)
                v4 = (v4 * v4)
            }
            if true {
                var v78 float32 = 0.0
                v78 = 1e10
                for fi79 := 0; fi79 < 3; fi79++ {
                    fmt.Println(v78, v4)
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    _ = fi79
                }
                _ = v78
            } else {
                switch "foo bar" {
                case "":
                    v80 := GenS3{F0: 987, F1: 42}
                    _ = v80
                default:
                    v81 := GenS7{F0: -1, F1: 176}
                    _ = v81
                }
            }
        }
    }
    var v82 float64 = float64(1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v82
}
