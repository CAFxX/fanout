package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }
type GenS3 struct { F0 uint; F1 int8; F2 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 int8 }

func main() {
    var v1 float32 = 0.1
    var v2 uint8 = 1
    var v3 uint8 = (v2 + (v2 % ((v2) | 1)))
    var v4 uint32 = 31
    var v5 float32 = (((-2.25 + 0.0) / float32(1.5)) + 1e10)
    v6 := []uint32{(v4 << ((125) & 63)), v4}
    v2 = (v3 * v3)
    if (18446744073709551615 >= 1) {
        v7 := GenS1{F0: 1}
        var v8 bool = true
        v6 = append(v6, (v4 << ((1) & 63)))
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            if len(v6) > 0 {
                v6[0] = (^(v4))
            }
            v5 = v1
            _ = fi10
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            if true {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v12 := GenS1{F0: 214}
                v4 -= 299
                _ = v12
            }
            _ = fi11
        }
        _ = fi9
    }
    var v13 uint = 607
    var v14 int32 = 700
    if (!false) {
        for fi15 := 0; fi15 < 3; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 uint64 = 1
                _ = fi16
                _ = v17
            }
            var v18 int8 = 127
            _ = fi15
            _ = v18
        }
        v14 = v14
        if (!true) {
            switch "hello" {
            case "":
                v14 = (v14 + 2147483647)
            case "hello":
                v19 := GenS2{F0: 1, F1: -765, F2: 509}
                _ = v19
            case "\n\t":
                for fi20 := 0; fi20 < 13; fi20++ {
                    if (!false) {
                        switch ("\n\t" + "") {
                        case "hello":
                            v4 = 1
                        case "a":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        v5 = (v1 / float32(((-2.25 + 0.1) + (3.14159 / 1.5))))
                    } else {
                        v21 := GenS1{F0: (v14 & (v14 ^ (v14 * v14)))}
                        v1 += 0.1
                        fmt.Println(v4, v2)
                        _ = v21
                    }
                    if true {
                        if false {
                            switch "foo bar" {
                            case "foo bar":
                                if (-9223372036854775808 == int64(v4)) {
                                    for fi22 := 0; fi22 < 17; fi22++ {
                                        var v23 uint = (v13 >> ((v13) & 63))
                                        v14 = (v14 - 2147483647)
                                        fmt.Println(fi20, v14)
                                        _ = fi22
                                        _ = v23
                                    }
                                    for fi24 := 0; fi24 < 14; fi24++ {
                                        var v25 uint8 = (v2 ^ (v3 % ((42) | 1)))
                                        v14 += v14
                                        _ = fi24
                                        _ = v25
                                    }
                                }
                            case "":
                                for fi26 := 0; fi26 < 20; fi26++ {
                                    v13 = v13
                                    switch (fi20 >> ((65535) & 63)) {
                                    case 242:
                                        fmt.Println(v4, v1, fi20)
                                    case -886:
                                        var v27 uint32 = v4
                                        _ = v27
                                    }
                                    v28 := GenS1{F0: (v14 - v14)}
                                    _ = fi26
                                    _ = v28
                                }
                            }
                        }
                        for fi29 := 0; fi29 < 17; fi29++ {
                            var v30 uint = v13
                            var v31 float32 = (v5 / v5)
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        var v32 float32 = float32(-2.25)
                        _ = v32
                    }
                    _ = fi20
                }
            default:
                var v33 int64 = -9223372036854775808
                _ = v33
            }
        }
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        if false {
            switch ((("" + "foo bar") + ("hello" + "hello")) + "hello") {
            case "hello":
                if (18446744073709551615 >= 1) {
                    if false {
                        v13 = v13
                    } else {
                        var v35 int32 = -2147483648
                        if ((3.14159 + (0.0 + 3.14159)) <= 3.14159) {
                            v2 = (v3 * v3)
                        } else {
                            switch (-(fi34)) {
                            case 32767:
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(fi34)
                                case "":
                                    var v36 int8 = 0
                                    _ = v36
                                default:
                                    v14 -= v14
                                }
                            case 10:
                                if (v4 != (v4 << ((463) & 63))) {
                                    v1 = ((0.0 + v1) - v5)
                                    var v37 float32 = v1
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v37
                                } else {
                                    for fi38 := 0; fi38 < 10; fi38++ {
                                        v39 := GenS2{F0: 4294967295, F1: 9223372036854775807, F2: 841}
                                        var v40 int16 = -1
                                        v41 := GenS2{F0: 18446744073709551615, F1: 34, F2: v14}
                                        _ = fi38
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                }
                            case 4294967295:
                                if len(v6) > 0 {
                                    v6[0] = v4
                                }
                            }
                            if (!true) {
                                v4 = 2147483647
                                switch -1 {
                                case 0:
                                    if true {
                                        v5 += v5
                                        v42 := GenS3{F0: 798, F1: 44, F2: 100}
                                        _ = v42
                                    } else {
                                        switch ((("a" + (("hello" + "hello") + "")) + "") + ("foo bar" + "a")) {
                                        case "":
                                            switch ("foo bar" + "foo bar") {
                                            case "":
                                                v43 := GenS4{F0: v4}
                                                _ = v43
                                            case "foo bar":
                                                if len(v6) > 0 {
                                                    v6[0] = (v4 / (((v4 / ((936) | 1))) | 1))
                                                }
                                            }
                                        case "\n\t":
                                            v44 := GenS1{F0: v14}
                                            _ = v44
                                        }
                                        for fi45 := 0; fi45 < 7; fi45++ {
                                            if false {
                                                if (uint16(v4) != 1) {
                                                    fmt.Println(v13)
                                                    v46 := GenS3{F0: (v13 & (v13 >> ((18446744073709551615) & 63))), F1: -1, F2: -74}
                                                    _ = v46
                                                } else {
                                                    if false {
                                                        fmt.Println(fi34, fi34)
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    } else {
                                                        if (-2.25 > (0.1 * (0.0 / ((float64(0.0) / 1e10) + ((3.14159 * -2.25) * (((float64(0.0) + 0.1) - (3.14159 - 3.14159)) - (-2.25 * 0.0))))))) {
                                                            switch -9223372036854775808 {
                                                            case 128:
                                                                switch (fi45 + (fi45 * 127)) {
                                                                case 370:
                                                                    v6 = append(v6, (v4 % (((v4 >> ((uint(v3)) & 63))) | 1)))
                                                                case -411:
                                                                    var v47 int32 = v35
                                                                    _ = v47
                                                                }
                                                            default:
                                                                switch (fi45 % ((32767) | 1)) {
                                                                case 32767:
                                                                    _ = len(v6)
                                                                default:
                                                                    var v48 uint16 = 160
                                                                    _ = v48
                                                                }
                                                            }
                                                        }
                                                        var v49 float64 = (0.0 * -2.25)
                                                        _ = v49
                                                    }
                                                    v50 := GenS1{F0: (v14 >> ((v13) & 63))}
                                                    _ = v50
                                                }
                                            }
                                            v51 := GenS3{F0: uint(v3), F1: 127, F2: -74}
                                            _ = fi45
                                            _ = v51
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v52 string = ("a" + (("hello" + "\n\t") + ""))
                                            _ = v52
                                        case "hello":
                                            switch 1 {
                                            case 0:
                                                var v53 uint32 = (v4 ^ v4)
                                                _ = v53
                                            case 928:
                                                var v54 int16 = 951
                                                _ = v54
                                            case 9223372036854775807:
                                                var v55 uint = (v13 >> ((v13) & 63))
                                                _ = v55
                                            default:
                                                var v56 float32 = 1.5
                                                _ = v56
                                            }
                                        case "foo bar":
                                            v35 -= v14
                                        }
                                    }
                                }
                            }
                            v5 = v1
                        }
                        fmt.Println(v13)
                        _ = v35
                    }
                    for fi57 := 0; fi57 < 13; fi57++ {
                        switch "\n\t" {
                        case "\n\t":
                            if (!true) {
                                fmt.Println(v2)
                            } else {
                                var v58 bool = false
                                var v59 int64 = 15
                                var v60 int16 = 241
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = v58
                                _ = v59
                                _ = v60
                            }
                        case "a":
                            v6 = append(v6, 984)
                        default:
                            var v61 int64 = 857
                            _ = v61
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi57
                    }
                } else {
                    for fi62 := 0; fi62 < 12; fi62++ {
                        if len(v6) > 0 {
                            v6[0] = 0
                        }
                        _ = fi62
                    }
                }
            case "":
                v63 := GenS5{F0: 2, F1: 59}
                _ = v63
            case "\n\t":
                v64 := GenS4{F0: v4}
                _ = v64
            }
        }
        _ = fi34
    }
    for fi65 := 0; fi65 < 20; fi65++ {
        v13 = v13
        v4 += (^(v4))
        _ = fi65
    }
    if (-945 < 9223372036854775807) {
        if false {
            v66 := GenS1{F0: (v14 - (^(v14)))}
            _ = v66
        }
        v3 -= v2
        var v67 float64 = 0.0
        _ = v67
    }
    v4 = v4
    var v68 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
}
