package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 bool }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 bool }
func genFn0(p0 uint32) int32 {
    v5 := []int{0, 32767}
    var v6 int64 = -1
    _ = p0
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 string) string {
    v7 := GenS1{F0: false, F1: 676}
    p0 = "foo bar"
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            if (false || true) {
                v10 := []int{fi8, (fi9 ^ fi8), fi9}
                var v11 uint16 = 255
                _ = v10
                _ = v11
            }
            _ = fi9
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        switch (fi12 % ((fi12) | 1)) {
        case 128:
            var v13 int8 = -1
            _ = v13
        case 9223372036854775807:
            p0 = (p0 + "")
        }
        var v14 bool = true
        fmt.Println(fi12)
        _ = fi12
        _ = v14
    }
    _ = p0
    _ = v7
    return p0
}

func main() {
    var v1 uint32 = 466
    var v2 uint32 = v1
    var v3 float32 = 1.5
    v4 := []uint32{(v1 ^ (v2 - (v1 | 636))), v1, (v2 >> ((599) & 63))}
    var v15 int64 = 0
    for fi16 := 0; fi16 < 3; fi16++ {
        if false {
            v17 := GenS1{F0: false, F1: 418}
            var v18 int8 = -128
            _ = v17
            _ = v18
        } else {
            v1 = v2
        }
        if false {
            for fi19 := 0; fi19 < 18; fi19++ {
                v3 += ((1.5 - v3) + v3)
                _ = fi19
            }
        }
        if true {
            if len(v4) > 0 {
                v4[0] = (v2 * v2)
            }
            if len(v4) > 0 {
                v4[0] = v2
            }
        } else {
            v20 := GenS1{F0: true, F1: 18446744073709551615}
            _ = v20
        }
        _ = fi16
    }
    var v21 uint32 = (v1 & (v2 >> ((2) & 63)))
    for fi22 := 0; fi22 < 6; fi22++ {
        if true {
            var v23 string = "a"
            v2 += v1
            v15 = (v15 + -705)
            _ = v23
        }
        v2 += (v21 - (v21 + v1))
        v2 = (v21 << ((835) & 63))
        _ = fi22
    }
    _ = len(v4)
    var v24 int64 = v15
    v25 := GenS1{F0: false, F1: 958}
    v26 := GenS1{F0: (!(1 >= 42)), F1: 33}
    switch "foo bar" {
    case "a":
        var v27 uint = 0
        _ = v27
    case "":
        if (!true) {
            var v28 int8 = -5
            _ = v28
        } else {
            v1 -= (v2 | (v1 | (v2 & (v1 * v2))))
        }
    case "hello":
        if (("a" + ("hello" + ("" + ("a" + "\n\t")))) < "hello") {
            fmt.Println(v15)
            var v29 uint64 = 4294967295
            _ = v29
        } else {
            fmt.Println(v2, v15, v3)
            _ = len(v4)
        }
    default:
        switch "hello" {
        case "":
            if (!true) {
                switch "\n\t" {
                case "\n\t":
                    v4 = append(v4, 2147483647)
                }
                switch "hello" {
                case "hello":
                    var v30 uint8 = 255
                    _ = v30
                case "\n\t":
                    switch 9223372036854775807 {
                    case 429:
                        for fi31 := 0; fi31 < 20; fi31++ {
                            if (567 <= 255) {
                                if (!(false || true)) {
                                    for fi32 := 0; fi32 < 11; fi32++ {
                                        switch fi32 {
                                        case 279:
                                            fmt.Println(fi32, v24)
                                            fmt.Println(v2)
                                            fmt.Println(fi31)
                                            var v33 float64 = 0.0
                                            v2 = (v21 << ((65535) & 63))
                                            if (uint(v21) == 255) {
                                                v34 := GenS1{F0: false, F1: 1}
                                                v35 := GenS1{F0: false, F1: 42}
                                                _ = v34
                                                _ = v35
                                            } else {
                                                var v36 int = (-(fi31))
                                                var v37 uint64 = 1
                                                if false {
                                                    var v38 uint32 = 65535
                                                    _ = v38
                                                }
                                                _ = v36
                                                _ = v37
                                            }
                                            var v39 bool = false
                                            switch "hello" {
                                            case "foo bar":
                                                v4 = append(v4, v2)
                                            case "\n\t":
                                                var v40 int8 = 42
                                                _ = v40
                                            default:
                                                v33 -= (-2.25 - v33)
                                            }
                                            var v41 uint8 = 0
                                            _ = v33
                                            _ = v39
                                            _ = v41
                                        case -1:
                                            if (65535 == 459) {
                                                if len(v4) > 0 {
                                                    v4[0] = v1
                                                }
                                                var v42 string = (("hello" + "foo bar") + "hello")
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v42
                                            } else {
                                                v3 += v3
                                            }
                                        }
                                        _ = fi32
                                    }
                                    v3 -= 1e10
                                    var v43 int64 = (^(v15))
                                    _ = v43
                                }
                                var v44 int64 = (v15 ^ v24)
                                if len(v4) > 0 {
                                    v4[0] = v2
                                }
                                _ = v44
                            } else {
                                for fi45 := 0; fi45 < 6; fi45++ {
                                    var v46 int64 = 492
                                    var v47 uint16 = 732
                                    var v48 float64 = -2.25
                                    _ = fi45
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                                for fi49 := 0; fi49 < 2; fi49++ {
                                    v3 = v3
                                    var v50 int32 = int32(v2)
                                    _ = fi49
                                    _ = v50
                                }
                            }
                            _ = fi31
                        }
                    case 32767:
                        switch -7 {
                        case -298:
                            v15 += -768
                        default:
                            var v51 int16 = 1000
                            _ = v51
                        }
                    }
                case "":
                    v21 = v21
                }
                v3 = (float32(0.0) - v3)
            }
        case "a":
            fmt.Println(v1, v1)
        }
    }
    v15 = v24
    for fi52 := 0; fi52 < 5; fi52++ {
        for fi53 := 0; fi53 < 4; fi53++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if false {
                switch "\n\t" {
                case "":
                    v4 = append(v4, v21)
                case "a":
                    switch fi53 {
                    case 0:
                        if (true && false) {
                            var v54 int8 = -1
                            _ = v54
                        } else {
                            v55 := GenS2{F0: 0, F1: 233, F2: (-1 != 0)}
                            v56 := GenS2{F0: 42, F1: 127, F2: (false && false)}
                            _ = v55
                            _ = v56
                        }
                    default:
                        _ = len(v4)
                    }
                default:
                    if (!false) {
                        v57 := GenS1{F0: (((("hello" + "") + "foo bar") <= "hello") || true), F1: 0}
                        _ = len(v4)
                        _ = v57
                    }
                }
                fmt.Println(v24)
                v1 = v21
            }
            if false {
                v21 += 65535
                v2 = v1
            } else {
                var v58 int16 = 1
                v15 = v24
                var v59 uint64 = 0
                _ = v58
                _ = v59
            }
            _ = fi53
        }
        if false {
            fmt.Println(fi52)
            for fi60 := 0; fi60 < 15; fi60++ {
                var v61 int8 = 100
                var v62 int32 = 0
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi60
                _ = v61
                _ = v62
            }
            if len(v4) > 0 {
                v4[0] = v21
            }
        } else {
            for fi63 := 0; fi63 < 10; fi63++ {
                if (((42 == 42) && (1 > 65535)) || (-1 <= -133)) {
                    switch (("" + "hello") + "a") {
                    case "foo bar":
                        v24 += v15
                        _ = len(v4)
                        if len(v4) > 0 {
                            v4[0] = v21
                        }
                    case "a":
                        switch ("" + "a") {
                        case "a":
                            var v64 int32 = 912
                            _ = v64
                        case "foo bar":
                            v1 += (v2 - (v21 - (v1 | v21)))
                        case "hello":
                            v15 += 9223372036854775807
                        }
                    case "":
                        v21 += v2
                    default:
                        var v65 string = ("a" + "\n\t")
                        _ = v65
                    }
                    fmt.Println(fi63, v2, v24)
                }
                fmt.Println(v15)
                v66 := GenS2{F0: 4294967295, F1: 235, F2: (false || ((false && true) || true))}
                if (216 == 127) {
                    if ((255 < 255) || true) {
                        switch 32768 {
                        case 404:
                            if true {
                                fmt.Println(v2, v2, v21)
                                var v67 int32 = genFn0((v1 << ((2) & 63)))
                                fmt.Println(fi52, fi63, v21)
                                _ = v67
                            }
                        case 128:
                            v68 := GenS3{F0: 65535, F1: 127}
                            v69 := GenS2{F0: 753, F1: 127, F2: true}
                            switch (("foo bar" + ((((("\n\t" + (("\n\t" + "a") + "foo bar")) + "foo bar") + "") + ("\n\t" + ("a" + "a"))) + (("hello" + ("" + ("hello" + ("a" + "\n\t")))) + "a"))) + "a") {
                            case "hello":
                                fmt.Println(v21)
                                v4 = append(v4, (v21 >> ((4294967295) & 63)))
                            case "":
                                var v70 int64 = (v24 ^ 13)
                                _ = v70
                            case "foo bar":
                                v15 = (v24 / (((v15 + v24)) | 1))
                                if true {
                                    v4 = append(v4, v2)
                                    v71 := GenS2{F0: 100, F1: 2, F2: true}
                                    _ = len(v4)
                                    _ = v71
                                } else {
                                    fmt.Println(v1)
                                }
                                v72 := GenS2{F0: 4294967295, F1: 89, F2: false}
                                fmt.Println(fi63, v3, v24)
                                _ = v72
                            }
                            _ = v68
                            _ = v69
                        case 869:
                            v73 := GenS4{F0: -53}
                            _ = v73
                        default:
                            _ = len(v4)
                        }
                    } else {
                        var v74 int64 = v15
                        switch 32768 {
                        case -2147483648:
                            switch (fi52 ^ 127) {
                            case 4294967295:
                                if (0 <= genFn0((v21 >> ((18446744073709551615) & 63)))) {
                                    if (true && false) {
                                        var v75 uint32 = (v1 << ((277) & 63))
                                        var v76 uint8 = 42
                                        var v77 int8 = -128
                                        _ = v75
                                        _ = v76
                                        _ = v77
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if (42 < 0) {
                                            var v78 uint16 = 68
                                            v79 := GenS2{F0: 497, F1: 40, F2: true}
                                            v78 += v78
                                            _ = len(v4)
                                            switch fi63 {
                                            case 246:
                                                var v80 int = fi63
                                                _ = v80
                                            }
                                            _ = v78
                                            _ = v79
                                        } else {
                                            v81 := GenS2{F0: 1, F1: 255, F2: (!false)}
                                            switch "" {
                                            case "hello":
                                                _ = len(v4)
                                            case "foo bar":
                                                fmt.Println(v1, v15, fi63)
                                            default:
                                                v15 = 9223372036854775807
                                            }
                                            _ = v81
                                        }
                                    }
                                    v74 -= (v74 & (v74 / (((v24 - v74)) | 1)))
                                } else {
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v21)
                                    case "foo bar":
                                        if (fi52 != fi63) {
                                            var v82 int = fi52
                                            _ = v82
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = len(v4)
                                        }
                                    default:
                                        var v83 int8 = 0
                                        _ = v83
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = v2
                                    }
                                }
                            case 954:
                                var v84 uint8 = 128
                                _ = v84
                            case 1:
                                if (610 < 65535) {
                                    switch 128 {
                                    case -354:
                                        fmt.Println(v2, v24)
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = v21
                                        }
                                    }
                                } else {
                                    v85 := GenS5{F0: v74, F1: true}
                                    var v86 float64 = 0.1
                                    _ = v85
                                    _ = v86
                                }
                            }
                        case 32768:
                            v24 += -243
                        case 498:
                            if (0.1 >= v3) {
                                fmt.Println(v2)
                            } else {
                                fmt.Println(fi52, v1)
                            }
                        default:
                            if false {
                                v87 := GenS5{F0: v74, F1: true}
                                _ = len(v4)
                                _ = v87
                            } else {
                                switch "foo bar" {
                                case "a":
                                    var v88 bool = (42 == 1)
                                    _ = v88
                                }
                            }
                        }
                        _ = v74
                    }
                } else {
                    fmt.Println(v15, v24, fi63)
                }
                if true {
                    fmt.Println(fi52)
                    var v89 int64 = (v24 * v15)
                    _ = v89
                }
                _ = fi63
                _ = v66
            }
            var v90 bool = (fi52 <= (-(fi52)))
            v24 = 0
            _ = v90
        }
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
