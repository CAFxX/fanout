package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int32; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 uint16; F2 string }
type GenS7 struct { F0 int16; F1 int64; F2 int }
type GenS8 struct { F0 string }
func genFn0(p0 int32, p1 int32, p2 uint64) bool {
    fmt.Println(p1, p2)
    var v7 int32 = -298
    switch "hello" {
    case "":
        v8 := GenS1{F0: "\n\t", F1: 32768}
        _ = v8
    case "\n\t":
        p1 -= -259
    case "foo bar":
        v9 := GenS2{F0: 255, F1: 800, F2: 42}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}
func genFn1(p0 uint, p1 int64, p2 uint32) int16 {
    var v10 uint32 = (p2 + 2147483648)
    fmt.Println(v10, p1, p2)
    var v11 float64 = ((0.0 + (-2.25 / 1e10)) - 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 uint = 42
    v2 := []int{-992, 758, 1, 0}
    var v3 bool = (!(true || (4294967295 < 0)))
    v4 := make(map[uint32]uint32)
    var v5 int = 606
    var v6 int = v5
    delete(v4, 358)
    for fi12 := 0; fi12 < 20; fi12++ {
        if (541 != 42) {
            var v13 uint32 = uint32(v1)
            v5 = (fi12 + v5)
            _ = v13
        } else {
            fmt.Println(v1, v5, v6)
            switch (("" + ("hello" + "a")) + "\n\t") {
            case "":
                _ = len(v2)
            default:
                v4[730] = 1
            }
        }
        switch (("\n\t" + "hello") + "foo bar") {
        case "a":
            switch "a" {
            case "\n\t":
                v1 -= v1
            default:
                var v14 float32 = (1.5 + 3.14159)
                _ = v14
            }
        case "\n\t":
            var v15 uint64 = 355
            _ = v15
        case "foo bar":
            v16 := GenS1{F0: ("hello" + "a"), F1: (v6 << ((109) & 63))}
            _ = v16
        }
        _ = fi12
    }
    switch v5 {
    case -9223372036854775808:
        switch (("\n\t" + ("\n\t" + ("foo bar" + ""))) + "foo bar") {
        case "":
            fmt.Println(v3, v3)
        default:
            if (65535 < 1) {
                var v17 float64 = (0.0 / 1.0)
                v18 := GenS1{F0: "a", F1: (v5 ^ v6)}
                delete(v4, 4294967295)
                _ = v17
                _ = v18
            }
        }
    case 127:
        var v19 int64 = 42
        _ = v19
    case -583:
        delete(v4, 0)
    }
    v20 := GenS1{F0: ("\n\t" + "a"), F1: v6}
    if v3 {
        var v21 int32 = 44
        if (v3 && (v3 && true)) {
            if ((v3 && ((float32(((-2.25 + float64(0.0)) * 1.5)) * ((1e10 + 3.14159) / 1.0)) != -2.25)) && v3) {
                var v22 int64 = 9223372036854775807
                _ = v22
            } else {
                fmt.Println(v1, v3, v6)
            }
        }
        _ = v21
    }
    var v23 uint = 1
    v23 -= 18446744073709551615
    switch (v6 << (((v23 ^ v1)) & 63)) {
    case 4294967295:
        switch "hello" {
        case "a":
            if v3 {
                if true {
                    fmt.Println(v5)
                    v24 := GenS3{F0: 42}
                    _ = v24
                }
            }
        case "foo bar":
            v2 = append(v2, (-(v5)))
        case "":
            var v25 int8 = int8(v23)
            _ = v25
        }
    case 128:
        for fi26 := 0; fi26 < 13; fi26++ {
            delete(v4, 65535)
            fmt.Println(v1, v6)
            _ = fi26
        }
    case 233:
        if (v3 && v3) {
            if (genFn1(v23, 230, 2147483648) != 1000) {
                v27 := GenS3{F0: (v23 ^ v23)}
                v6 = 2147483647
                _ = v27
            }
        }
    default:
        var v28 int = (-(v5))
        _ = v28
    }
    if v3 {
        switch "a" {
        case "foo bar":
            fmt.Println(v6)
        case "hello":
            switch v5 {
            case 34:
                for fi29 := 0; fi29 < 3; fi29++ {
                    fmt.Println(fi29, v5)
                    var v30 uint = 2
                    var v31 float32 = (1.5 * (0.0 - 1.5))
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            case 2147483647:
                var v32 int8 = -1
                _ = v32
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        if true {
            v1 = v23
            if v3 {
                for fi34 := 0; fi34 < 5; fi34++ {
                    v35 := GenS3{F0: v23}
                    fmt.Println(fi33, v5, v5)
                    switch fi34 {
                    case 91:
                        var v36 bool = (127 > 1)
                        _ = v36
                    case 9223372036854775807:
                        if len(v2) > 0 {
                            v2[0] = -9223372036854775808
                        }
                    default:
                        switch ("" + "a") {
                        case "a":
                            if v3 {
                                v4[739] = 2147483647
                                var v37 string = "a"
                                switch v37 {
                                case "foo bar":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "\n\t":
                                    fmt.Println(v1)
                                    _ = len(v4)
                                    fmt.Println(v6, v6)
                                case "hello":
                                    if v3 {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        var v38 int8 = -128
                                        switch ("foo bar" + v37) {
                                        case "a":
                                            if (34 >= 1) {
                                                v4[2147483647] = 1
                                                var v39 float32 = 0.0
                                                v40 := GenS4{F0: 87, F1: -1, F2: -1}
                                                switch (^(fi33)) {
                                                case 9223372036854775807:
                                                    if false {
                                                        if true {
                                                            var v41 int8 = int8(fi34)
                                                            _ = v41
                                                        }
                                                    } else {
                                                        var v42 int = 32768
                                                        _ = v42
                                                    }
                                                case 42:
                                                    if v3 {
                                                        v23 = 42
                                                        v2 = append(v2, 0)
                                                    } else {
                                                        switch 128 {
                                                        case 408:
                                                            v4[2] = 100
                                                            var v43 bool = v3
                                                            delete(v4, 2147483648)
                                                            v37 = v37
                                                            fmt.Println(v38)
                                                            switch (fi33 + 128) {
                                                            case 42:
                                                                switch v5 {
                                                                case 970:
                                                                    var v44 int8 = -106
                                                                    _ = v44
                                                                case -2147483648:
                                                                    v37 = v37
                                                                }
                                                            case 127:
                                                                switch v37 {
                                                                case "a":
                                                                    v4[1] = 42
                                                                case "hello":
                                                                    var v45 int = (fi34 / (((-(fi34))) | 1))
                                                                    _ = v45
                                                                case "foo bar":
                                                                    var v46 int8 = v38
                                                                    _ = v46
                                                                default:
                                                                    v47 := v4[263]
                                                                    _ = v47
                                                                }
                                                            case 32768:
                                                                var v48 string = (v37 + v37)
                                                                _ = v48
                                                            }
                                                            _ = len(v2)
                                                            var v49 bool = (!false)
                                                            v4[103] = 732
                                                            var v50 int32 = -161
                                                            fmt.Println(fi34)
                                                            _ = v43
                                                            _ = v49
                                                            _ = v50
                                                        case -2147483648:
                                                            var v51 int64 = 42
                                                            fmt.Println(fi33, v5)
                                                            _ = v51
                                                        case 42:
                                                            _ = len(v2)
                                                            _ = len(v4)
                                                            v52 := GenS1{F0: "foo bar", F1: (fi33 >> (((v23 - v23)) & 63))}
                                                            _ = v52
                                                        default:
                                                            v1 -= (v23 / (((v1 >> ((v1) & 63))) | 1))
                                                        }
                                                        fmt.Println(v23, v1, v5)
                                                        var v53 uint16 = 42
                                                        _ = v53
                                                    }
                                                }
                                                switch (v6 | (v5 & 32767)) {
                                                case 4294967295:
                                                    var v54 float32 = (v39 - 0.1)
                                                    _ = v54
                                                }
                                                v55 := GenS2{F0: 255, F1: 42, F2: 716}
                                                _ = v39
                                                _ = v40
                                                _ = v55
                                            }
                                        }
                                        if false {
                                            v3 = (0.1 > 3.14159)
                                        }
                                        _ = v38
                                    }
                                }
                                _ = v37
                            } else {
                                switch (fi33 - (v5 << ((174) & 63))) {
                                case 866:
                                    if ("foo bar" <= ("hello" + "")) {
                                        var v56 int64 = -1
                                        _ = v56
                                    }
                                case 4294967295:
                                    v6 += -209
                                case 237:
                                    fmt.Println(v6, v23)
                                }
                            }
                        case "foo bar":
                            v1 -= (v1 << ((v1) & 63))
                        case "hello":
                            var v57 int = (fi34 >> ((v1) & 63))
                            v58 := GenS2{F0: 0, F1: 1, F2: 0}
                            _ = v57
                            _ = v58
                        default:
                            if (615 < 32767) {
                                v59 := GenS4{F0: 231, F1: 19, F2: 887}
                                v23 -= v23
                                _ = v59
                            } else {
                                if v3 {
                                    var v60 uint64 = 407
                                    v61 := GenS1{F0: ("\n\t" + (("a" + "a") + ("foo bar" + ""))), F1: v5}
                                    if false {
                                        switch "\n\t" {
                                        case "a":
                                            var v62 float32 = 0.0
                                            _ = v62
                                        }
                                        if v3 {
                                            if false {
                                                v1 -= (v1 - 255)
                                                var v63 uint32 = 4294967295
                                                var v64 int32 = 2147483647
                                                v65 := GenS2{F0: 0, F1: (v63 >> (((v23 * (v23 | (^(v1))))) & 63)), F2: int64(v63)}
                                                var v66 float32 = ((1e10 * -2.25) - 0.1)
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            }
                                            v67 := GenS1{F0: "\n\t", F1: fi34}
                                            switch v6 {
                                            case 790:
                                                fmt.Println(fi34, fi33)
                                            case 375:
                                                v68 := GenS1{F0: "a", F1: (v5 | (fi34 * fi33))}
                                                _ = v68
                                            }
                                            var v69 int64 = 0
                                            _ = v67
                                            _ = v69
                                        }
                                    } else {
                                        v70 := GenS5{F0: 0}
                                        _ = v70
                                    }
                                    switch (("a" + ("" + "foo bar")) + "hello") {
                                    case "":
                                        var v71 uint64 = v60
                                        _ = v71
                                    case "hello":
                                        v72 := GenS6{F0: (v60 | (v60 ^ v60)), F1: 42, F2: "a"}
                                        _ = v72
                                    }
                                    _ = v60
                                    _ = v61
                                }
                                fmt.Println(v3, v6, v5)
                            }
                        }
                    }
                    _ = fi34
                    _ = v35
                }
                var v73 string = "hello"
                _ = len(v4)
                _ = v73
            } else {
                var v74 int16 = 1000
                var v75 int64 = 42
                switch (v5 * int(v74)) {
                case -190:
                    v1 += (^(v1))
                case 804:
                    switch "foo bar" {
                    case "":
                        var v76 float32 = 0.1
                        _ = v76
                    case "a":
                        v3 = v3
                    default:
                        switch "a" {
                        case "foo bar":
                            var v77 int8 = 127
                            _ = v77
                        case "hello":
                            switch (("foo bar" + ("" + ("a" + "foo bar"))) + "a") {
                            case "foo bar":
                                for fi78 := 0; fi78 < 10; fi78++ {
                                    fmt.Println(v3, fi78, v75)
                                    if v3 {
                                        v2 = append(v2, v6)
                                        v79 := GenS2{F0: 0, F1: 65535, F2: v75}
                                        v4[42] = 4294967295
                                        if len(v2) > 0 {
                                            v2[0] = int(v74)
                                        }
                                        v23 += v1
                                        _ = v79
                                    }
                                    _ = fi78
                                }
                            case "":
                                var v80 int8 = -128
                                _ = v80
                            case "hello":
                                switch "hello" {
                                case "foo bar":
                                    v81 := GenS7{F0: -939, F1: int64(v74), F2: 0}
                                    _ = v81
                                }
                            default:
                                v23 = (v1 % (((v1 >> ((v23) & 63))) | 1))
                            }
                        default:
                            if v3 {
                                var v82 uint8 = 42
                                _ = v82
                            } else {
                                var v83 uint = (v1 % ((4294967295) | 1))
                                _ = v83
                            }
                        }
                    }
                case 42:
                    v5 = (fi33 - 32767)
                default:
                    v84 := GenS8{F0: (("hello" + "") + (((("foo bar" + ("\n\t" + "\n\t")) + "hello") + "\n\t") + "hello"))}
                    _ = v84
                }
                _ = v74
                _ = v75
            }
            v85 := GenS5{F0: -399}
            _ = v85
        }
        _ = fi33
    }
    v6 += v5
    v1 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
