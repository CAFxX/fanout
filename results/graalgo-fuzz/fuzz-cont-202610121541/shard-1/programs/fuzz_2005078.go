package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 string; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint64; F1 uint64; F2 string }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 uint32; F1 int8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS12 struct { F0 uint16; F1 bool }
type GenS13 struct { F0 uint32; F1 uint32 }
type GenS14 struct { F0 string; F1 int }
func genFn0(p0 bool) int32 {
    var v5 uint16 = 65535
    p0 = true
    _ = p0
    _ = v5
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"a", "\n\t", "hello", "foo bar"}
    var v3 int8 = -26
    v4 := []uint32{}
    var v6 int = 4294967295
    if false {
        if ("foo bar" >= "a") {
            if false {
                var v7 uint16 = 1
                v8 := GenS1{F0: v7, F1: 42, F2: v7}
                var v9 uint = 290
                _ = v7
                _ = v8
                _ = v9
            } else {
                fmt.Println(v3, v6)
            }
        }
    } else {
        if (600 <= 4294967295) {
            if false {
                v3 = (v3 | -128)
                var v10 int64 = 1
                _ = v10
            } else {
                v3 += -84
                if false {
                    for fi11 := 0; fi11 < 13; fi11++ {
                        v4 = append(v4, 65535)
                        var v12 float64 = 0.1
                        v6 -= v6
                        _ = fi11
                        _ = v12
                    }
                }
            }
            v13 := GenS1{F0: 284, F1: -32768, F2: 1}
            switch int(v3) {
            case -9223372036854775808:
                fmt.Println(v3, v3)
            case 145:
                if false {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        var v15 uint16 = 255
                        _ = fi14
                        _ = v15
                    }
                    var v16 uint = 441
                    _ = v16
                }
            case -1:
                v3 = (v3 - v3)
            }
            _ = v13
        } else {
            for fi17 := 0; fi17 < 17; fi17++ {
                for fi18 := 0; fi18 < 4; fi18++ {
                    delete(v1, "hello")
                    _ = fi18
                }
                _ = fi17
            }
        }
        if (!true) {
            switch v6 {
            case -952:
                if true {
                    if false {
                        if (((v6 <= v6) && ((!false) && true)) || (false || false)) {
                            v19 := GenS2{F0: 1, F1: 32767, F2: 2}
                            v3 += v3
                            var v20 uint8 = 128
                            _ = v19
                            _ = v20
                        } else {
                            for fi21 := 0; fi21 < 17; fi21++ {
                                switch ("hello" + "\n\t") {
                                case "":
                                    v3 = v3
                                }
                                _ = fi21
                            }
                        }
                        for fi22 := 0; fi22 < 15; fi22++ {
                            _ = len(v4)
                            v23 := GenS2{F0: 1, F1: -85, F2: 4294967295}
                            _ = fi22
                            _ = v23
                        }
                    }
                }
            case 9223372036854775807:
                if false {
                    switch (("hello" + "hello") + "hello") {
                    case "":
                        v6 = v6
                    case "\n\t":
                        v24 := v1["hello"]
                        _ = v24
                    case "hello":
                        if true {
                            v3 -= 100
                            v25 := GenS3{F0: 65535}
                            _ = v25
                        }
                    default:
                        if (!(1e10 >= -2.25)) {
                            if true {
                                delete(v1, "foo bar")
                            } else {
                                if ((!true) && false) {
                                    var v26 uint = 0
                                    _ = v26
                                }
                                v6 -= (v6 + v6)
                                var v27 int8 = v3
                                _ = v27
                            }
                        } else {
                            v6 = -2147483648
                            v28 := GenS4{F0: 0}
                            if true {
                                var v29 uint8 = 194
                                v6 -= (v6 % ((-9223372036854775808) | 1))
                                _ = v29
                            } else {
                                v30 := GenS3{F0: 65535}
                                v6 = (v6 << ((968) & 63))
                                fmt.Println(v3, v3)
                                _ = v30
                            }
                            _ = v28
                        }
                    }
                } else {
                    v6 -= -1
                }
            case 4294967295:
                if (false && (1e10 == 0.1)) {
                    v3 -= 100
                    for fi31 := 0; fi31 < 1; fi31++ {
                        if (255 == 984) {
                            switch "foo bar" {
                            case "":
                                for fi32 := 0; fi32 < 19; fi32++ {
                                    switch "foo bar" {
                                    case "":
                                        v33 := v1[("hello" + (("a" + "\n\t") + ("" + "")))]
                                        _ = v33
                                    case "hello":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "foo bar":
                                        switch ("hello" + "foo bar") {
                                        case "foo bar":
                                            v34 := GenS2{F0: 2147483648, F1: -11, F2: 219}
                                            _ = v34
                                        case "":
                                            var v35 int16 = -530
                                            _ = v35
                                        case "hello":
                                            fmt.Println(v6, v6)
                                        default:
                                            fmt.Println(v6, fi32, v6)
                                        }
                                        v36 := GenS4{F0: 9223372036854775807}
                                        _ = v36
                                    default:
                                        fmt.Println(fi31, v6)
                                    }
                                    switch fi31 {
                                    case -2147483648:
                                        fmt.Println(v3, v6, fi32)
                                        fmt.Println(fi32)
                                    case 9223372036854775807:
                                        var v37 float64 = float64(-2.25)
                                        _ = v37
                                    case -9223372036854775808:
                                        fmt.Println(v3, fi32)
                                    default:
                                        fmt.Println(fi32)
                                        var v38 int64 = 0
                                        _ = v38
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        v6 = (v6 ^ fi32)
                                    case "\n\t":
                                        v6 = v6
                                    default:
                                        fmt.Println(fi31)
                                    }
                                    switch (fi32 ^ fi32) {
                                    case 2147483647:
                                        v39 := GenS2{F0: 702, F1: 1000, F2: 2}
                                        _ = v39
                                    }
                                    _ = fi32
                                }
                            case "a":
                                switch 32767 {
                                case -2147483648:
                                    if true {
                                        if (854 == 0) {
                                            for fi40 := 0; fi40 < 16; fi40++ {
                                                fmt.Println(v6)
                                                v41 := GenS5{F0: 0, F1: ("hello" + ("a" + "foo bar")), F2: 1}
                                                v42 := GenS5{F0: 42, F1: "\n\t", F2: 2147483647}
                                                var v43 int16 = 32767
                                                if (!true) {
                                                    v3 = (v3 ^ -107)
                                                    _ = len(v1)
                                                    _ = len(v4)
                                                } else {
                                                    fmt.Println(v3)
                                                }
                                                _ = fi40
                                                _ = v41
                                                _ = v42
                                                _ = v43
                                            }
                                            fmt.Println(fi31)
                                        }
                                        v3 -= 127
                                        if false {
                                            if ((1e10 + 1e10) != 0.0) {
                                                var v44 uint8 = 36
                                                _ = v44
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = 0
                                                }
                                                v3 = 15
                                                v45 := GenS2{F0: 2147483647, F1: 584, F2: 4294967295}
                                                _ = v45
                                            }
                                            for fi46 := 0; fi46 < 5; fi46++ {
                                                var v47 int16 = 42
                                                _ = fi46
                                                _ = v47
                                            }
                                        }
                                    }
                                }
                            default:
                                v3 = 0
                            }
                            switch (v6 ^ v6) {
                            case 128:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        } else {
                            for fi48 := 0; fi48 < 18; fi48++ {
                                v4 = append(v4, 1)
                                fmt.Println(fi31, fi48)
                                _ = fi48
                            }
                        }
                        switch fi31 {
                        case 309:
                            v6 = -2147483648
                        default:
                            v3 = (v3 & v3)
                        }
                        _ = fi31
                    }
                } else {
                    v3 = 100
                }
            }
        } else {
            if (((865 <= 472) && (!false)) && true) {
                switch ("\n\t" + "") {
                case "foo bar":
                    if true {
                        v3 -= (v3 | (^(v3)))
                        var v49 float32 = 1.5
                        _ = v49
                    }
                case "a":
                    for fi50 := 0; fi50 < 14; fi50++ {
                        if true {
                            if false {
                                v6 = v6
                                var v51 uint32 = 800
                                _ = v51
                            }
                            var v52 uint = 221
                            v6 -= (v6 + v6)
                            _ = v52
                        } else {
                            if len(v4) > 0 {
                                v4[0] = 42
                            }
                        }
                        _ = fi50
                    }
                }
                for fi53 := 0; fi53 < 10; fi53++ {
                    for fi54 := 0; fi54 < 1; fi54++ {
                        v3 = v3
                        var v55 uint = 0
                        _ = fi54
                        _ = v55
                    }
                    fmt.Println(v6, v3, v3)
                    _ = fi53
                }
                if true {
                    v6 -= v6
                    v1["a"] = v6
                } else {
                    v3 = 110
                }
            }
        }
        var v56 float32 = 0.0
        _ = v56
    }
    if (2147483648 == 0) {
        v57 := GenS6{F0: 1, F1: 689}
        for fi58 := 0; fi58 < 18; fi58++ {
            var v59 uint16 = 335
            fmt.Println(v59)
            _ = fi58
            _ = v59
        }
        v2 = append(v2, "\n\t")
        _ = v57
    }
    for fi60 := 0; fi60 < 15; fi60++ {
        v3 += v3
        v6 += v6
        switch "a" {
        case "a":
            if ((-39 < v3) || true) {
                fmt.Println(fi60, v6)
                v61 := GenS7{F0: 902, F1: 2, F2: ("hello" + "")}
                var v62 int32 = -492
                _ = v61
                _ = v62
            }
        }
        _ = fi60
    }
    switch "\n\t" {
    case "a":
        v63 := GenS8{F0: 0, F1: 42}
        _ = v63
    case "\n\t":
        var v64 float32 = (0.0 + float32(3.14159))
        _ = v64
    }
    v65 := GenS9{F0: 190, F1: 100}
    for fi66 := 0; fi66 < 10; fi66++ {
        var v67 int64 = 274
        _ = fi66
        _ = v67
    }
    if ("\n\t" != ("a" + (("\n\t" + (("\n\t" + "foo bar") + "a")) + "\n\t"))) {
        v6 -= (^(v6))
        v68 := GenS10{F0: "a"}
        for fi69 := 0; fi69 < 17; fi69++ {
            for fi70 := 0; fi70 < 11; fi70++ {
                if true {
                    if true {
                        v3 = -34
                        if (false && false) {
                            if (447 >= 648) {
                                v6 = 2147483647
                                v3 = v3
                                v3 = (^(v3))
                            }
                            if true {
                                var v71 uint8 = 2
                                if true {
                                    var v72 int16 = -1
                                    v73 := GenS11{F0: v72, F1: (v71 >> ((18446744073709551615) & 63)), F2: 2}
                                    _ = len(v1)
                                    if false {
                                        switch fi70 {
                                        case 0:
                                            switch "hello" {
                                            case "foo bar":
                                                v72 = (v72 - v72)
                                            case "\n\t":
                                                v72 += int16(v71)
                                            case "hello":
                                                v3 = v3
                                                v72 = v72
                                            }
                                        case 32767:
                                            if true {
                                                v3 -= (v3 ^ (v3 << ((2) & 63)))
                                                fmt.Println(v3, v3, fi70)
                                                v1[("foo bar" + "a")] = fi69
                                            }
                                        case 1:
                                            var v74 float32 = 1e10
                                            var v75 float32 = v74
                                            v76 := v1["a"]
                                            var v77 uint8 = v71
                                            v78 := GenS12{F0: 2, F1: false}
                                            v4 = append(v4, 0)
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                            _ = v78
                                        default:
                                            var v79 bool = ((fi69 ^ 42) != v6)
                                            _ = v79
                                        }
                                        var v80 int = fi70
                                        _ = v80
                                    }
                                    v81 := GenS11{F0: v72, F1: (v71 * 14), F2: 4294967295}
                                    var v82 int64 = -259
                                    var v83 int = v6
                                    _ = v72
                                    _ = v73
                                    _ = v81
                                    _ = v82
                                    _ = v83
                                } else {
                                    v2 = append(v2, (("hello" + ("" + "\n\t")) + "\n\t"))
                                }
                                fmt.Println(fi69)
                                _ = v71
                            }
                            v84 := v1["hello"]
                            _ = v84
                        }
                    }
                    var v85 uint32 = 4294967295
                    v3 = (v3 / ((v3) | 1))
                    _ = v85
                } else {
                    v86 := GenS13{F0: 935, F1: 2147483647}
                    _ = v86
                }
                fmt.Println(v6, fi70, v3)
                v2 = append(v2, "foo bar")
                _ = fi70
            }
            v3 += -13
            _ = fi69
        }
        _ = v68
    }
    var v87 int8 = -127
    var v88 int32 = genFn0(true)
    v89 := GenS14{F0: "foo bar", F1: (v6 - v6)}
    v88 -= (v88 << ((18446744073709551615) & 63))
    v4 = append(v4, 2147483648)
    var v90 int64 = 143
    fmt.Println(v6, v6)
    v91 := v1[""]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v65
    _ = v87
    _ = v88
    _ = v89
    _ = v90
    _ = v91
}
