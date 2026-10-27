package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int16; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int64; F2 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int64; F2 int8 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 uint32, p1 uint8) uint8 {
    p1 += 1
    var v7 string = "a"
    p0 = uint32(p1)
    for fi8 := 0; fi8 < 13; fi8++ {
        if (-128 > 118) {
            v9 := []int{fi8, -2147483648, (fi8 + fi8), (fi8 + (fi8 % ((fi8) | 1)))}
            var v10 bool = (int32(p0) <= -2147483648)
            _ = v10
            _ = v9
        } else {
            switch fi8 {
            case 128:
                if true {
                    fmt.Println(p1)
                }
            default:
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 string = v7
                    v12 = ((v7 + "") + "\n\t")
                    if false {
                        var v13 string = "a"
                        switch "foo bar" {
                        case "":
                            fmt.Println(v7, p1)
                        }
                        _ = v13
                    } else {
                        p1 += 255
                    }
                    _ = fi11
                    _ = v12
                }
            }
            var v14 string = v7
            p1 += (p1 - 0)
            _ = v14
        }
        v15 := []int{(fi8 << ((255) & 63)), fi8, (fi8 ^ fi8), (fi8 >> ((945) & 63))}
        _ = fi8
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v7
    return (^(p1))
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint32 = (v1 % ((v1) | 1))
    var v3 uint64 = 1
    var v4 uint8 = 42
    var v5 uint8 = (v4 - (v4 - uint8(v3)))
    v6 := []string{"a", "foo bar", ("\n\t" + ("\n\t" + ""))}
    var v16 uint16 = 198
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v17 uint64 = (v3 * v3)
    var v18 uint64 = (v3 ^ (v3 & 1))
    switch "hello" {
    case "foo bar":
        v19 := GenS1{F0: false, F1: -647}
        _ = v19
    case "hello":
        switch ("foo bar" + ("hello" + "a")) {
        case "a":
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 float64 = (1e10 * 1.5)
                if ((v16 - (v16 ^ v16)) <= (v16 | (v16 << ((uint(v3)) & 63)))) {
                    v17 = uint64(v1)
                    var v22 uint16 = v16
                    for fi23 := 0; fi23 < 16; fi23++ {
                        var v24 uint = 18446744073709551615
                        switch (("foo bar" + "foo bar") + ("hello" + (("hello" + "") + ""))) {
                        case "\n\t":
                            v18 = v3
                        case "foo bar":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        default:
                            v21 = 0.1
                        }
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
        case "foo bar":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "":
            switch (("a" + "foo bar") + ("hello" + "\n\t")) {
            case "hello":
                v16 = v16
            }
        }
    default:
        var v25 float64 = 1.5
        _ = v25
    }
    v26 := GenS2{F0: 0, F1: 0, F2: uint32(v4)}
    var v27 int = 833
    v28 := GenS2{F0: 1, F1: 42, F2: v1}
    fmt.Println(v17, v1, v2)
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 uint32 = v1
            v32 := GenS3{F0: (v16 << ((uint(v16)) & 63)), F1: -934, F2: "\n\t"}
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi29
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if false {
            v6 = append(v6, "\n\t")
            for fi34 := 0; fi34 < 7; fi34++ {
                if (fi33 > fi33) {
                    switch (v27 % ((-1) | 1)) {
                    case 32767:
                        v1 += (v2 ^ (v2 >> ((255) & 63)))
                    case 1:
                        var v35 int16 = 0
                        _ = v35
                    case 144:
                        if ((4294967295 > v1) && true) {
                            v36 := GenS1{F0: (413 <= -2147483648), F1: (v27 / ((45) | 1))}
                            _ = v36
                        }
                    }
                    if len(v6) > 0 {
                        v6[0] = ""
                    }
                }
                _ = fi34
            }
        } else {
            switch (^(v27)) {
            case -402:
                v18 += uint64(v5)
            case -210:
                v18 = v3
            case 9223372036854775807:
                v6 = append(v6, ("foo bar" + "a"))
            default:
                if false {
                    if (-40 < 45) {
                        var v37 float64 = 0.1
                        var v38 uint = uint(v3)
                        _ = v37
                        _ = v38
                    } else {
                        if (!(((fi33 * (^(fi33))) <= v27) || (0.1 == 0.0))) {
                            fmt.Println(v18, v17, v4)
                            for fi39 := 0; fi39 < 5; fi39++ {
                                var v40 int = fi39
                                if (true && ("" != "hello")) {
                                    var v41 string = "foo bar"
                                    _ = v41
                                }
                                v4 = genFn0((v2 - (v1 - (^(v1)))), (^(v5)))
                                _ = fi39
                                _ = v40
                            }
                        } else {
                            if true {
                                v42 := GenS3{F0: v16, F1: -9223372036854775808, F2: ("a" + "hello")}
                                var v43 uint8 = uint8(v18)
                                _ = v42
                                _ = v43
                            }
                            v44 := GenS4{F0: 1}
                            _ = v44
                        }
                    }
                    var v45 uint = 255
                    v1 -= (^(v1))
                    _ = v45
                }
            }
            v3 = (v18 >> ((187) & 63))
            if (!true) {
                if ((^(v16)) <= 217) {
                    v16 = v16
                } else {
                    var v46 uint8 = (v5 * (v5 / ((v4) | 1)))
                    for fi47 := 0; fi47 < 6; fi47++ {
                        var v48 uint = 4294967295
                        v6 = append(v6, "\n\t")
                        _ = fi47
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 6; fi49++ {
                        fmt.Println(v3)
                        v27 -= -793
                        fmt.Println(v4, v18, v27)
                        if (255 <= 128) {
                            var v50 uint64 = v18
                            v51 := GenS5{F0: true}
                            if (4294967295 == 1) {
                                v46 = v4
                                v52 := GenS6{F0: v2}
                                _ = v52
                            }
                            _ = v50
                            _ = v51
                        } else {
                            v53 := GenS1{F0: (42 != 0), F1: fi33}
                            var v54 int16 = 42
                            _ = v53
                            _ = v54
                        }
                        v5 += (^(v4))
                        _ = fi49
                    }
                    _ = v46
                }
            }
        }
        if len(v6) > 0 {
            v6[0] = "foo bar"
        }
        _ = fi33
    }
    if ((v1 | v2) < 1) {
        _ = len(v6)
        var v55 float64 = (1.5 / ((1.5 - 0.1) / (1.5 / 1.5)))
        if (100 == 81) {
            fmt.Println(v55, v2, v1)
            if false {
                var v56 uint = 1
                v56 = 993
                _ = v56
            }
            for fi57 := 0; fi57 < 3; fi57++ {
                v1 = 312
                _ = fi57
            }
        } else {
            if (((^(v27)) >= 128) || (v55 >= (1e10 * (v55 + v55)))) {
                var v58 int8 = -16
                if (true || (true && (-9223372036854775808 > -745))) {
                    fmt.Println(v17, v16, v17)
                    v2 -= v1
                } else {
                    fmt.Println(v4, v27)
                    switch "a" {
                    case "a":
                        switch (v27 >> ((65535) & 63)) {
                        case 32767:
                            v1 += v2
                        case 2147483647:
                            v6 = append(v6, "foo bar")
                        case 911:
                            fmt.Println(v5, v3, v18)
                        default:
                            v3 = (v3 % (((^(v18))) | 1))
                        }
                    default:
                        var v59 uint16 = v16
                        _ = v59
                    }
                    for fi60 := 0; fi60 < 20; fi60++ {
                        switch (fi60 - (v27 - fi60)) {
                        case 4294967295:
                            for fi61 := 0; fi61 < 2; fi61++ {
                                v62 := GenS7{F0: 42, F1: 9223372036854775807, F2: 127}
                                if (2 == 911) {
                                    if ((v4 + (v5 ^ (v5 + (v5 * v4)))) != (^(v5))) {
                                        v5 -= v4
                                    }
                                    var v63 string = "hello"
                                    var v64 float64 = (-2.25 + (float64(3.14159) / float64((float32(0.0) - float32((v55 * v55))))))
                                    v65 := GenS4{F0: 1}
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                                v27 = fi60
                                var v66 uint32 = v1
                                _ = fi61
                                _ = v62
                                _ = v66
                            }
                        case 42:
                            var v67 uint = 2
                            _ = v67
                        case 32768:
                            if len(v6) > 0 {
                                v6[0] = "a"
                            }
                        }
                        for fi68 := 0; fi68 < 18; fi68++ {
                            if true {
                                fmt.Println(fi60, v17)
                                switch (fi60 / ((fi68) | 1)) {
                                case 4294967295:
                                    if ((fi68 >> ((0) & 63)) == -683) {
                                        v2 = (v2 * 2)
                                        var v69 uint16 = v16
                                        v27 = fi60
                                        _ = v69
                                    }
                                case 9223372036854775807:
                                    var v70 uint64 = 154
                                    _ = v70
                                case 0:
                                    fmt.Println(v55, v55, v58)
                                }
                            }
                            switch "" {
                            case "foo bar":
                                v58 = v58
                            case "":
                                switch "" {
                                case "hello":
                                    v5 = (^(v5))
                                case "foo bar":
                                    v71 := GenS8{F0: (v1 >> ((1) & 63))}
                                    _ = v71
                                case "\n\t":
                                    v6 = append(v6, "a")
                                }
                            case "\n\t":
                                v2 -= v2
                            default:
                                switch v27 {
                                case -176:
                                    var v72 uint8 = (^(v4))
                                    _ = v72
                                default:
                                    var v73 int16 = 0
                                    _ = v73
                                }
                            }
                            _ = len(v6)
                            v17 -= 42
                            var v74 int64 = 0
                            if (42 <= -13) {
                                var v75 uint = 65535
                                var v76 int8 = 8
                                _ = v75
                                _ = v76
                            } else {
                                var v77 string = (("a" + (((("a" + "foo bar") + "\n\t") + (("\n\t" + "\n\t") + ("hello" + "hello"))) + "")) + "")
                                v2 -= (v2 >> ((989) & 63))
                                var v78 int8 = 0
                                _ = v77
                                _ = v78
                            }
                            _ = fi68
                            _ = v74
                        }
                        if (!true) {
                            v1 += (v2 | (v2 - (v1 ^ (^(v1)))))
                            switch "" {
                            case "foo bar":
                                if false {
                                    for fi79 := 0; fi79 < 13; fi79++ {
                                        v17 -= v3
                                        v58 -= int8(v27)
                                        if false {
                                            var v80 int = fi60
                                            var v81 float32 = 0.1
                                            _ = v80
                                            _ = v81
                                        } else {
                                            if len(v6) > 0 {
                                                v6[0] = "foo bar"
                                            }
                                            var v82 uint8 = v5
                                            v1 = v2
                                            _ = v82
                                        }
                                        _ = fi79
                                    }
                                    var v83 int8 = v58
                                    _ = v83
                                }
                            case "\n\t":
                                var v84 int16 = 0
                                _ = v84
                            }
                            var v85 bool = true
                            _ = v85
                        } else {
                            for fi86 := 0; fi86 < 10; fi86++ {
                                switch -869 {
                                case -9223372036854775808:
                                    switch "a" {
                                    case "hello":
                                        v1 = (v1 + uint32(v3))
                                    }
                                case 128:
                                    if ((!(310 == v16)) && false) {
                                        if (!(-32768 > 42)) {
                                            var v87 int32 = 0
                                            v6 = append(v6, ("a" + "a"))
                                            v4 = v4
                                            _ = v87
                                        } else {
                                            fmt.Println(v27)
                                            fmt.Println(v55)
                                        }
                                    } else {
                                        v55 += v55
                                    }
                                case 42:
                                    var v88 int32 = int32(v58)
                                    _ = v88
                                default:
                                    v6 = append(v6, "foo bar")
                                }
                                _ = fi86
                            }
                        }
                        _ = fi60
                    }
                }
                for fi89 := 0; fi89 < 3; fi89++ {
                    var v90 int8 = v58
                    _ = fi89
                    _ = v90
                }
                _ = v58
            }
        }
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
