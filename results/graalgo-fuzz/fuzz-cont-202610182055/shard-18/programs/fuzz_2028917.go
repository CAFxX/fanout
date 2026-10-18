package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int; F1 uint; F2 uint64 }
type GenS4 struct { F0 int16; F1 int; F2 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 uint32; F1 string; F2 uint }
type GenS8 struct { F0 uint32; F1 uint8; F2 int32 }
func genFn0(p0 uint32) string {
    v7 := GenS1{F0: 2}
    var v8 int32 = 42
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := []int{fi9, (^(fi9)), (fi9 % ((fi9) | 1)), fi9}
        switch "\n\t" {
        case "a":
            var v11 string = ("\n\t" + ("foo bar" + "hello"))
            _ = v11
        default:
            var v12 float32 = (-2.25 + 1.5)
            _ = v12
        }
        if len(v10) > 0 {
            v10[0] = (fi9 | (fi9 ^ fi9))
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(p0, v8)
    _ = p0
    _ = v7
    _ = v8
    return ("hello" + "hello")
}

func main() {
    var v1 uint64 = 1
    var v2 int8 = -101
    v3 := make(map[uint32]string)
    var v4 int16 = 1
    var v5 int64 = -9223372036854775808
    var v6 bool = (!(!false))
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            v4 = (v4 + int16(v5))
            fmt.Println(v5, v5)
            switch 127 {
            case 32768:
                v6 = (v5 != 946)
            }
            _ = fi14
        }
        if true {
            v4 += v4
        } else {
            if (127 > int8(v1)) {
                if ((0.0 * 0.0) >= 3.14159) {
                    var v15 uint64 = v1
                    _ = v15
                }
            }
            v16 := []int{fi13, (fi13 ^ int(v1)), -2147483648, (fi13 * (fi13 - (^(fi13))))}
            v17 := GenS1{F0: 42}
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    var v18 int64 = v5
    v6 = true
    for fi19 := 0; fi19 < 17; fi19++ {
        v1 = v1
        v20 := GenS2{F0: 511, F1: 318}
        if ((v5 >> ((4294967295) & 63)) > 536) {
            switch "\n\t" {
            case "":
                if v6 {
                    v5 += (v18 / (((v18 | (v5 ^ v18))) | 1))
                    v6 = false
                    v21 := GenS1{F0: 1}
                    _ = v21
                }
            case "a":
                if false {
                    var v22 int = fi19
                    switch "foo bar" {
                    case "foo bar":
                        v23 := v3[751]
                        _ = v23
                    }
                    fmt.Println(v6, v4, v5)
                    _ = v22
                } else {
                    var v24 uint8 = 2
                    for fi25 := 0; fi25 < 8; fi25++ {
                        switch ("a" + "foo bar") {
                        case "hello":
                            v18 -= v18
                        case "a":
                            var v26 int = (fi25 % ((fi25) | 1))
                            _ = v26
                        case "":
                            fmt.Println(fi25, fi19)
                        }
                        _ = fi25
                    }
                    v3[2] = "a"
                    _ = v24
                }
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    v1 = v1
                default:
                    switch (fi19 / ((fi19) | 1)) {
                    case 397:
                        switch ("foo bar" + "a") {
                        case "":
                            var v27 uint = 65535
                            _ = v27
                        default:
                            switch "a" {
                            case "\n\t":
                                v18 -= v5
                            default:
                                for fi28 := 0; fi28 < 10; fi28++ {
                                    var v29 string = ("foo bar" + "foo bar")
                                    _ = fi28
                                    _ = v29
                                }
                            }
                        }
                    case -622:
                        if true {
                            v6 = (v5 == v5)
                            switch (("" + "\n\t") + (("" + (("a" + "foo bar") + ("hello" + "foo bar"))) + "foo bar")) {
                            case "":
                                if false {
                                    v30 := []int{}
                                    if ((!true) && false) {
                                        if ((v2 << ((65535) & 63)) != 127) {
                                            var v31 uint = 2
                                            _ = v31
                                        } else {
                                            v32 := v3[65535]
                                            if true {
                                                for fi33 := 0; fi33 < 14; fi33++ {
                                                    if false {
                                                        delete(v3, 2147483647)
                                                        var v34 int = (fi19 - (^(fi33)))
                                                        var v35 int16 = (v4 ^ (v4 / ((v4) | 1)))
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        if true {
                                                            if (v6 || (42 < 2)) {
                                                                if (v4 >= v4) {
                                                                    var v36 int16 = (v4 | 1)
                                                                    switch "" {
                                                                    case "hello":
                                                                        if (true || v6) {
                                                                            var v37 int16 = (v4 * int16(v1))
                                                                            _ = v37
                                                                        }
                                                                    case "a":
                                                                        v38 := v3[2147483647]
                                                                        _ = v38
                                                                    default:
                                                                        var v39 uint16 = 1
                                                                        _ = v39
                                                                    }
                                                                    if true {
                                                                        v40 := GenS2{F0: 935, F1: 404}
                                                                        switch "a" {
                                                                        case "a":
                                                                            v41 := v3[65535]
                                                                            _ = v41
                                                                        default:
                                                                            var v42 uint16 = uint16(v1)
                                                                            _ = v42
                                                                        }
                                                                        _ = v40
                                                                    } else {
                                                                        v43 := GenS3{F0: (fi19 - 42), F1: 65535, F2: 4294967295}
                                                                        var v44 int64 = (v18 << ((255) & 63))
                                                                        var v45 uint32 = 1
                                                                        v30 = append(v30, fi33)
                                                                        var v46 int8 = v2
                                                                        if (255 <= 65535) {
                                                                            v47 := GenS1{F0: 10}
                                                                            fmt.Println(v4, v2, fi33)
                                                                            _ = v47
                                                                        }
                                                                        _ = v43
                                                                        _ = v44
                                                                        _ = v45
                                                                        _ = v46
                                                                    }
                                                                    _ = v36
                                                                } else {
                                                                    delete(v3, 2147483648)
                                                                    switch (fi19 << ((0) & 63)) {
                                                                    case 32767:
                                                                        if ((v32 + "a") >= (v32 + (("hello" + "hello") + "foo bar"))) {
                                                                            var v48 uint64 = v1
                                                                            _ = v48
                                                                        } else {
                                                                            v2 = v2
                                                                        }
                                                                    case -2147483648:
                                                                        v49 := GenS4{F0: v4, F1: (fi33 * fi19), F2: 1}
                                                                        _ = v49
                                                                    case 127:
                                                                        v32 = "hello"
                                                                    default:
                                                                        if v6 {
                                                                            if (v2 == -128) {
                                                                                v32 = "a"
                                                                                v4 += v4
                                                                            }
                                                                            v4 = int16(v18)
                                                                            var v50 uint = 0
                                                                            _ = v50
                                                                        } else {
                                                                            fmt.Println(fi19, v2, v18)
                                                                            var v51 bool = false
                                                                            v52 := GenS5{F0: 350}
                                                                            _ = v51
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                    if len(v30) > 0 {
                                                                        v30[0] = (^(fi19))
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v4, v2, v1)
                                                        }
                                                        v53 := v3[1]
                                                        _ = v53
                                                    }
                                                    _ = fi33
                                                }
                                            }
                                            _ = v32
                                        }
                                    } else {
                                        v4 += (v4 - (v4 * v4))
                                        v1 = 2
                                    }
                                    _ = v30
                                }
                            case "hello":
                                v54 := GenS6{F0: (v2 & (v2 * v2)), F1: 131}
                                _ = v54
                            case "a":
                                v5 = int64(v1)
                            default:
                                var v55 bool = (!(255 >= 127))
                                _ = v55
                            }
                            if (962 == (-(v4))) {
                                fmt.Println(v6)
                            } else {
                                for fi56 := 0; fi56 < 7; fi56++ {
                                    fmt.Println(v5, v6)
                                    v57 := GenS4{F0: (v4 >> ((4294967295) & 63)), F1: (-(fi56)), F2: 192}
                                    _ = fi56
                                    _ = v57
                                }
                                v1 = v1
                            }
                        }
                    case 1:
                        fmt.Println(v2)
                    default:
                        for fi58 := 0; fi58 < 9; fi58++ {
                            if (134 != 41) {
                                fmt.Println(v1, v2)
                                fmt.Println(v4, fi58)
                                v1 -= v1
                            }
                            var v59 bool = true
                            _ = fi58
                            _ = v59
                        }
                    }
                }
            case "hello":
                switch (((("hello" + "\n\t") + "\n\t") + "\n\t") + "") {
                case "\n\t":
                    for fi60 := 0; fi60 < 3; fi60++ {
                        var v61 int = (-(fi60))
                        v62 := GenS7{F0: 4294967295, F1: "\n\t", F2: 2}
                        switch (v61 ^ fi60) {
                        case -1:
                            fmt.Println(v5, v4)
                        case -9223372036854775808:
                            v63 := GenS7{F0: 2147483648, F1: (("a" + "a") + ("hello" + ("\n\t" + "a"))), F2: 18446744073709551615}
                            _ = v63
                        case 32767:
                            fmt.Println(v5, v61)
                            v1 -= v1
                        default:
                            var v64 int64 = v18
                            _ = v64
                        }
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                case "foo bar":
                    for fi65 := 0; fi65 < 3; fi65++ {
                        fmt.Println(v4)
                        _ = fi65
                    }
                case "a":
                    var v66 int8 = 1
                    _ = v66
                default:
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        var v67 uint = 628
                        _ = v67
                    default:
                        var v68 int16 = v4
                        _ = v68
                    }
                }
            case "a":
                fmt.Println(v6)
            default:
                v69 := GenS3{F0: (fi19 + (-(fi19))), F1: 580, F2: 0}
                _ = v69
            }
        }
        _ = fi19
        _ = v20
    }
    v6 = v6
    v70 := GenS8{F0: 854, F1: 128, F2: 2147483647}
    _ = len(v3)
    for fi71 := 0; fi71 < 9; fi71++ {
        v4 -= v4
        _ = fi71
    }
    for fi72 := 0; fi72 < 19; fi72++ {
        v18 = -367
        var v73 int16 = v4
        _ = fi72
        _ = v73
    }
    var v74 float32 = (((0.0 * (0.0 - float32(1e10))) - 1e10) * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v74
}
