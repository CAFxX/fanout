package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 string; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 bool; F1 string; F2 int8 }
type GenS8 struct { F0 uint64; F1 uint32; F2 uint }
type GenS9 struct { F0 int64; F1 uint32 }
type GenS10 struct { F0 int16; F1 int16 }
type GenS11 struct { F0 uint32; F1 int8; F2 uint }
type GenS12 struct { F0 uint32; F1 bool; F2 int16 }
type GenS13 struct { F0 uint32; F1 uint8; F2 uint16 }

func main() {
    v1 := []int{-2147483648, 782, 9223372036854775807}
    var v2 float32 = 1.5
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    v5 := []string{"", ("hello" + "hello"), "hello", "hello"}
    switch ("a" + "a") {
    case "\n\t":
        v6 := GenS1{F0: 322, F1: 18446744073709551615}
        _ = v6
    case "":
        var v7 int64 = 9223372036854775807
        _ = v7
    default:
        var v8 bool = true
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        if (!true) {
            v10 := GenS1{F0: -660, F1: 18446744073709551615}
            if (!false) {
                if true {
                    var v11 float32 = 0.0
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 12; fi12++ {
                        var v13 int32 = 0
                        var v14 int16 = 582
                        var v15 int8 = 100
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                }
                for fi16 := 0; fi16 < 12; fi16++ {
                    if false {
                        fmt.Println(fi16, fi9)
                    } else {
                        var v17 float64 = 0.0
                        v18 := GenS1{F0: 1, F1: 4294967295}
                        _ = v17
                        _ = v18
                    }
                    var v19 float64 = (0.1 - 1e10)
                    _ = fi16
                    _ = v19
                }
            }
            var v20 uint8 = 127
            _ = v10
            _ = v20
        } else {
            if true {
                switch fi9 {
                case 32767:
                    switch fi9 {
                    case 1:
                        v2 = 0.0
                    case -9223372036854775808:
                        var v21 uint32 = 65535
                        _ = v21
                    default:
                        v2 += (1.5 / v2)
                    }
                case 0:
                    for fi22 := 0; fi22 < 19; fi22++ {
                        if ((!false) || true) {
                            v23 := GenS2{F0: false, F1: 65535, F2: 2147483648}
                            v24 := GenS3{F0: 340, F1: true}
                            _ = v23
                            _ = v24
                        } else {
                            v25 := GenS4{F0: ("\n\t" + "foo bar"), F1: 2147483647, F2: 255}
                            v26 := GenS5{F0: 128, F1: 0, F2: 1}
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS6{F0: "", F1: 18446744073709551615}
                        v28 := GenS7{F0: true, F1: "foo bar", F2: 100}
                        _ = fi22
                        _ = v27
                        _ = v28
                    }
                case -671:
                    v2 += (v2 * (((((((v2 / (v2 / (v2 * 0.0))) * v2) / (v2 / v2)) * 3.14159) - v2) + (v2 + (float32((((-2.25 - 0.1) - 3.14159) - 1.5)) - ((v2 + ((v2 / float32(0.1)) / (v2 * 0.0))) + 3.14159)))) / 0.1))
                }
            }
        }
        _ = fi9
    }
    v2 = 0.0
    if (6 < -128) {
        v29 := GenS5{F0: 2, F1: 1, F2: -128}
        var v30 uint8 = 77
        switch (("a" + ("" + "\n\t")) + "a") {
        case "a":
            switch 557 {
            case 128:
                for fi31 := 0; fi31 < 20; fi31++ {
                    v5 = append(v5, "hello")
                    delete(v3, (fi31 | 4294967295))
                    _ = fi31
                }
            case -927:
                _ = len(v3)
            default:
                _ = len(v3)
            }
        default:
            for fi32 := 0; fi32 < 7; fi32++ {
                delete(v3, (-(fi32)))
                var v33 int16 = 1
                var v34 int64 = 0
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        _ = v29
        _ = v30
    }
    switch "foo bar" {
    case "hello":
        for fi35 := 0; fi35 < 14; fi35++ {
            for fi36 := 0; fi36 < 2; fi36++ {
                if ((127 >= 127) && false) {
                    v1 = append(v1, (fi36 % ((fi36) | 1)))
                } else {
                    v37 := GenS8{F0: 4294967295, F1: 2147483648, F2: 57}
                    delete(v4, (fi35 & (fi35 | (fi36 + (fi36 % ((fi36) | 1))))))
                    v2 -= v2
                    _ = v37
                }
                switch (fi36 + 32767) {
                case 9223372036854775807:
                    if true {
                        v2 = v2
                        v38 := GenS9{F0: 9223372036854775807, F1: 4294967295}
                        v39 := GenS5{F0: 42, F1: 2147483648, F2: -128}
                        _ = v38
                        _ = v39
                    }
                }
                _ = fi36
            }
            for fi40 := 0; fi40 < 1; fi40++ {
                if false {
                    if true {
                        if (1 > 570) {
                            v2 -= ((v2 + v2) * 1e10)
                            fmt.Println(v2)
                        } else {
                            if (-120 > 42) {
                                if len(v1) > 0 {
                                    v1[0] = fi40
                                }
                                if true {
                                    switch ("\n\t" + ((("hello" + "hello") + "") + ("hello" + "\n\t"))) {
                                    case "hello":
                                        v41 := GenS10{F0: -32768, F1: 42}
                                        _ = v41
                                    }
                                }
                            } else {
                                v2 += v2
                                var v42 uint8 = 2
                                var v43 uint8 = v42
                                _ = v42
                                _ = v43
                            }
                        }
                        if false {
                            _ = len(v5)
                            switch fi35 {
                            case 127:
                                v2 -= v2
                            }
                            if (-1 != 1) {
                                var v44 uint16 = 579
                                v45 := GenS11{F0: 978, F1: 120, F2: 41}
                                var v46 uint32 = 42
                                v2 += (v2 / 1.0)
                                _ = v44
                                _ = v45
                                _ = v46
                            } else {
                                v2 -= v2
                                var v47 float64 = 3.14159
                                v48 := GenS1{F0: 42, F1: 355}
                                var v49 uint32 = 1
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                        }
                        switch "\n\t" {
                        case "a":
                            _ = len(v4)
                        case "\n\t":
                            var v50 uint8 = 251
                            _ = v50
                        case "hello":
                            var v51 int = fi35
                            _ = v51
                        }
                        var v52 bool = (42 == 0)
                        _ = v52
                    } else {
                        fmt.Println(fi35, v2, fi35)
                        fmt.Println(v2)
                        if (("" == "a") && ((-1 == 42) && false)) {
                            var v53 int = (fi35 << ((18446744073709551615) & 63))
                            _ = v53
                        } else {
                            var v54 uint16 = 0
                            fmt.Println(fi35, fi35, v2)
                            if (!((32767 >= 502) && false)) {
                                var v55 string = "foo bar"
                                if (false || (v2 <= 1.5)) {
                                    _ = len(v3)
                                    fmt.Println(v55, v54, v2)
                                }
                                if false {
                                    v2 = 0.1
                                    if (((((float64((v2 + v2)) + 3.14159) + 0.1) / 3.14159) + 1.5) <= (-2.25 - -2.25)) {
                                        var v56 uint8 = 0
                                        v2 = 1.5
                                        _ = v56
                                    } else {
                                        var v57 float64 = 0.1
                                        _ = v57
                                    }
                                    v2 = (v2 + v2)
                                } else {
                                    v58 := GenS12{F0: 227, F1: (255 <= 42), F2: 996}
                                    if true {
                                        if false {
                                            v2 = v2
                                            var v59 uint64 = 4294967295
                                            _ = v59
                                        } else {
                                            fmt.Println(fi35, v2, v54)
                                        }
                                        v54 = 136
                                        var v60 uint16 = (^(v54))
                                        _ = v60
                                    }
                                    v2 = 3.14159
                                    _ = v58
                                }
                                _ = v55
                            } else {
                                v54 += 42
                                v54 = v54
                            }
                            _ = v54
                        }
                    }
                } else {
                    v61 := GenS13{F0: 0, F1: 255, F2: 0}
                    _ = v61
                }
                var v62 int16 = -894
                var v63 int8 = 127
                _ = fi40
                _ = v62
                _ = v63
            }
            _ = fi35
        }
    case "foo bar":
        if (true || true) {
            if false {
                var v64 uint32 = 503
                _ = v64
            }
            fmt.Println(v2)
        } else {
            v5 = append(v5, ("hello" + ""))
        }
    default:
        if true {
            _ = len(v1)
            v2 = (v2 + 1.5)
            v2 = (0.1 - (v2 / ((((float32((3.14159 * -2.25)) * (v2 - v2)) - float32((-2.25 / (3.14159 + 1e10)))) / 1.0) - ((((0.1 - (0.0 * v2)) * v2) - (v2 - v2)) - v2))))
        }
    }
    if (0 != -128) {
        if (4294967295 == 18446744073709551615) {
            for fi65 := 0; fi65 < 3; fi65++ {
                if len(v5) > 0 {
                    v5[0] = ("a" + ((("" + "\n\t") + "a") + "foo bar"))
                }
                fmt.Println(v2, fi65)
                _ = fi65
            }
        }
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            if true {
                v2 = v2
            }
        }
    }
    switch 747 {
    case 119:
        var v66 uint = 909
        _ = v66
    case 786:
        for fi67 := 0; fi67 < 11; fi67++ {
            var v68 int32 = 0
            _ = fi67
            _ = v68
        }
    }
    v2 += ((v2 / 1.0) - v2)
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    var v69 uint = 356
    var v70 uint = (^(v69))
    if (100 > -1) {
        delete(v4, 2147483647)
        v2 -= v2
        var v71 uint16 = 821
        _ = v71
    }
    fmt.Println(v69, v70)
    switch "a" {
    case "\n\t":
        switch -237 {
        case 9223372036854775807:
            var v72 int8 = 42
            _ = v72
        }
    case "hello":
        for fi73 := 0; fi73 < 15; fi73++ {
            fmt.Println(fi73, fi73)
            for fi74 := 0; fi74 < 6; fi74++ {
                if ((true && false) || (false || false)) {
                    if (56 < -128) {
                        var v75 float64 = 1e10
                        _ = v75
                    } else {
                        var v76 int = (fi74 | (^(fi74)))
                        var v77 string = "hello"
                        _ = v76
                        _ = v77
                    }
                    if true {
                        var v78 int32 = -48
                        _ = v78
                    }
                } else {
                    _ = len(v4)
                    v1 = append(v1, 127)
                }
                _ = fi74
            }
            _ = fi73
        }
    case "foo bar":
        for fi79 := 0; fi79 < 15; fi79++ {
            var v80 uint32 = 2
            v81 := GenS9{F0: -595, F1: v80}
            v82 := GenS5{F0: 2, F1: uint32(v70), F2: 127}
            _ = fi79
            _ = v80
            _ = v81
            _ = v82
        }
    }
    switch 4294967295 {
    case 128:
        v70 += v70
    case -417:
        var v83 string = ((("\n\t" + "\n\t") + "") + "hello")
        _ = v83
    case 4294967295:
        v70 += v70
    default:
        for fi84 := 0; fi84 < 9; fi84++ {
            v85 := v4[(fi84 - -995)]
            v70 = v70
            _ = fi84
            _ = v85
        }
    }
    v70 += v70
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
}
