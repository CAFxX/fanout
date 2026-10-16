package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16; F1 bool; F2 int32 }
type GenS10 struct { F0 int8; F1 int; F2 uint }

func main() {
    var v1 uint8 = 127
    var v2 int = 32767
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    var v5 int16 = -1
    v2 += (v2 | (v2 ^ v2))
    if false {
        fmt.Println(v1, v5)
    }
    v1 -= 2
    switch v2 {
    case -9223372036854775808:
        var v6 uint = 255
        _ = v6
    case 2147483647:
        v2 = (v2 + 413)
    case 32768:
        v7 := GenS1{F0: v1, F1: (^(v1))}
        _ = v7
    }
    if (v2 == -850) {
        switch 4294967295 {
        case -2147483648:
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 int32 = -2147483648
                v2 = fi8
                switch "hello" {
                case "\n\t":
                    var v10 int64 = 42
                    _ = v10
                case "foo bar":
                    var v11 uint64 = 217
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        case 1:
            v12 := GenS1{F0: (v1 & 255), F1: v1}
            _ = v12
        }
    }
    v5 += (v5 - (v5 << ((2) & 63)))
    if true {
        for fi13 := 0; fi13 < 6; fi13++ {
            delete(v3, 42)
            _ = fi13
        }
        fmt.Println(v2)
        switch ("" + "a") {
        case "hello":
            for fi14 := 0; fi14 < 11; fi14++ {
                if (float64(0.0) >= 0.1) {
                    if (1 == 373) {
                        v15 := []int{(v2 | (v2 + -337)), (fi14 + fi14)}
                        for fi16 := 0; fi16 < 20; fi16++ {
                            v5 += 452
                            _ = fi16
                        }
                        _ = v15
                    } else {
                        var v17 int16 = v5
                        switch ("" + "foo bar") {
                        case "\n\t":
                            if (float32(1e10) < (float32(float64((0.0 / 0.1))) - float32((1e10 * 1.5)))) {
                                for fi18 := 0; fi18 < 6; fi18++ {
                                    v5 -= (v5 & v17)
                                    v5 -= v5
                                    if false {
                                        v2 -= 0
                                        v19 := GenS1{F0: v1, F1: v1}
                                        _ = v19
                                    }
                                    var v20 string = "foo bar"
                                    v21 := GenS2{F0: (("\n\t" + "\n\t") + "hello"), F1: 166}
                                    _ = fi18
                                    _ = v20
                                    _ = v21
                                }
                                var v22 float64 = 0.0
                                _ = v22
                            }
                        case "hello":
                            if (2 != 0) {
                                v5 += int16(v1)
                            } else {
                                v23 := GenS3{F0: 838}
                                v24 := GenS2{F0: "hello", F1: 42}
                                _ = v23
                                _ = v24
                            }
                        case "foo bar":
                            v25 := GenS3{F0: 1}
                            _ = v25
                        default:
                            for fi26 := 0; fi26 < 18; fi26++ {
                                _ = len(v3)
                                if (770 > 4294967295) {
                                    fmt.Println(fi14, v1, fi26)
                                }
                                v27 := []int{}
                                if true {
                                    fmt.Println(fi26, v17, fi26)
                                } else {
                                    if len(v27) > 0 {
                                        v27[0] = (^(fi14))
                                    }
                                }
                                _ = fi26
                                _ = v27
                            }
                        }
                        _ = v17
                    }
                    fmt.Println(v2, v5, fi14)
                } else {
                    _ = len(v4)
                    v28 := GenS3{F0: 2}
                    for fi29 := 0; fi29 < 18; fi29++ {
                        var v30 string = "hello"
                        var v31 uint64 = 4294967295
                        var v32 int = 313
                        _ = fi29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    _ = v28
                }
                var v33 uint16 = 32767
                _ = fi14
                _ = v33
            }
        case "foo bar":
            if false {
                for fi34 := 0; fi34 < 10; fi34++ {
                    var v35 int16 = v5
                    _ = fi34
                    _ = v35
                }
            }
        case "a":
            for fi36 := 0; fi36 < 6; fi36++ {
                if (660 < 0) {
                    for fi37 := 0; fi37 < 13; fi37++ {
                        var v38 int32 = 42
                        v5 += (v5 + 0)
                        if false {
                            _ = len(v3)
                            v2 = 128
                        } else {
                            switch (-(fi37)) {
                            case 2147483647:
                                _ = len(v3)
                            case 1:
                                v4[1] = (361 > 175)
                            default:
                                v39 := GenS4{F0: 42}
                                _ = v39
                            }
                            _ = len(v4)
                        }
                        _ = fi37
                        _ = v38
                    }
                    v2 += 9
                    v40 := GenS3{F0: 839}
                    _ = v40
                }
                if true {
                    v41 := GenS5{F0: fi36, F1: 0}
                    _ = len(v4)
                    var v42 int8 = 100
                    _ = v41
                    _ = v42
                } else {
                    switch "hello" {
                    case "":
                        if (-401 >= 9223372036854775807) {
                            v2 -= fi36
                            fmt.Println(v2, v2)
                            var v43 int64 = -9223372036854775808
                            _ = v43
                        }
                    case "hello":
                        if ((!(("a" + "a") == ("a" + ""))) && (!false)) {
                            fmt.Println(v5, v2)
                            var v44 int64 = 9223372036854775807
                            delete(v3, 42)
                            _ = v44
                        } else {
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(fi36, fi36)
                            default:
                                fmt.Println(v1)
                            }
                        }
                    default:
                        fmt.Println(v2, v5)
                    }
                    v45 := GenS5{F0: v2, F1: 195}
                    for fi46 := 0; fi46 < 12; fi46++ {
                        if (9223372036854775807 > 0) {
                            switch v2 {
                            case -685:
                                fmt.Println(v2, fi36)
                            case 1:
                                fmt.Println(v2, v2, v5)
                            case 447:
                                v47 := []int{v2, (fi36 & v2), (fi36 ^ fi46)}
                                var v48 bool = true
                                _ = len(v47)
                                _ = v47
                                _ = v48
                            default:
                                v2 -= (fi36 * v2)
                            }
                            fmt.Println(v1, fi36)
                            v49 := GenS6{F0: 1, F1: 0}
                            _ = v49
                        }
                        v5 -= -1
                        fmt.Println(v1)
                        if (255 <= 42) {
                            if false {
                                switch "\n\t" {
                                case "a":
                                    switch fi46 {
                                    case 704:
                                        if (fi46 == fi46) {
                                            v5 = v5
                                            var v50 string = ("" + (("\n\t" + (((("" + ("" + ("\n\t" + ("hello" + "a")))) + "foo bar") + "hello") + "foo bar")) + ("\n\t" + "\n\t")))
                                            _ = v50
                                        } else {
                                            fmt.Println(v1, v2, v1)
                                        }
                                    }
                                case "foo bar":
                                    var v51 string = ("hello" + ("" + ""))
                                    _ = v51
                                case "hello":
                                    var v52 uint64 = 1
                                    _ = v52
                                }
                                switch "" {
                                case "foo bar":
                                    v1 = (v1 << ((uint(v1)) & 63))
                                case "hello":
                                    var v53 int32 = -1
                                    _ = v53
                                case "\n\t":
                                    if true {
                                        v3[2] = (fi36 / ((fi36) | 1))
                                        if (1 < 42) {
                                            switch (v2 >> ((411) & 63)) {
                                            case -527:
                                                var v54 float32 = 0.1
                                                _ = v54
                                            }
                                            fmt.Println(fi46, v2, v5)
                                            var v55 int16 = (v5 / ((v5) | 1))
                                            _ = v55
                                        }
                                    }
                                }
                                v1 += 42
                            }
                            var v56 int8 = -1
                            _ = v56
                        } else {
                            v5 += (v5 ^ v5)
                            v5 -= v5
                        }
                        if (-945 <= -1) {
                            var v57 uint16 = 2
                            _ = v57
                        } else {
                            v58 := GenS7{F0: 0}
                            v59 := []int{(fi36 | v2), -980, fi36, (-(v2))}
                            fmt.Println(v1, fi36)
                            v2 -= int(v1)
                            _ = v58
                            _ = v59
                        }
                        _ = fi46
                    }
                    _ = v45
                }
                _ = fi36
            }
        }
    } else {
        fmt.Println(v2, v5)
        _ = len(v3)
    }
    switch (v2 << ((1) & 63)) {
    case 206:
        v5 -= v5
    case -68:
        v60 := GenS8{F0: -1}
        _ = v60
    case 0:
        for fi61 := 0; fi61 < 10; fi61++ {
            v62 := GenS9{F0: (v5 * 42), F1: (((!true) || true) && false), F2: -2147483648}
            v1 = (v1 - (v1 * 127))
            _ = fi61
            _ = v62
        }
    }
    var v63 uint = 266
    v64 := GenS10{F0: 0, F1: v2, F2: v63}
    v1 = v1
    v65 := GenS8{F0: 115}
    var v66 string = "hello"
    if true {
        for fi67 := 0; fi67 < 17; fi67++ {
            if ((v1 % ((v1) | 1)) > 0) {
                var v68 uint8 = (v1 * v1)
                for fi69 := 0; fi69 < 18; fi69++ {
                    v70 := v3[569]
                    fmt.Println(v1)
                    _ = fi69
                    _ = v70
                }
                _ = v68
            } else {
                v71 := GenS9{F0: (v5 | (v5 * 1)), F1: (!true), F2: -1}
                fmt.Println(v1)
                v5 += v5
                _ = v71
            }
            v5 -= (v5 % ((-957) | 1))
            fmt.Println(v1)
            _ = fi67
        }
        if (true && (!(-98 < -1))) {
            switch 32768 {
            case 128:
                switch v66 {
                case "hello":
                    v72 := v3[uint32(v63)]
                    _ = v72
                case "":
                    var v73 uint64 = 2
                    _ = v73
                case "\n\t":
                    for fi74 := 0; fi74 < 15; fi74++ {
                        if false {
                            for fi75 := 0; fi75 < 2; fi75++ {
                                var v76 int32 = 42
                                v77 := []int{fi75, (fi74 >> ((4294967295) & 63))}
                                _ = fi75
                                _ = v76
                                _ = v77
                            }
                        }
                        _ = fi74
                    }
                }
            }
        } else {
            if (-2.25 == -2.25) {
                var v78 int = 128
                _ = v78
            }
            v63 = (v63 - (^(v63)))
        }
    }
    v5 -= (v5 * v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v64
    _ = v65
    _ = v66
}
