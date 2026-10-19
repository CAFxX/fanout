package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 int; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint16; F2 string }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 int16; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    v3 := []int{417}
    var v4 int16 = 0
    fmt.Println(v4)
    fmt.Println(v4)
    if false {
        v4 = -1
    } else {
        if false {
            if false {
                v4 += v4
                if false {
                    v4 = (-(v4))
                    switch ("a" + "foo bar") {
                    case "a":
                        v5 := GenS1{F0: (!true), F1: "", F2: 42}
                        _ = v5
                    case "foo bar":
                        for fi6 := 0; fi6 < 13; fi6++ {
                            if false {
                                for fi7 := 0; fi7 < 1; fi7++ {
                                    var v8 float64 = 1.5
                                    _ = fi7
                                    _ = v8
                                }
                            }
                            _ = fi6
                        }
                    case "\n\t":
                        v3 = append(v3, 0)
                    default:
                        v4 = v4
                    }
                }
                var v9 uint8 = 42
                _ = v9
            } else {
                var v10 int = 32768
                v11 := GenS2{F0: 676, F1: 0}
                _ = v10
                _ = v11
            }
        }
    }
    if (9223372036854775807 != 858) {
        if true {
            v4 -= v4
        }
        v12 := GenS1{F0: (!true), F1: ("" + "hello"), F2: 65535}
        if (65535 == 1) {
            for fi13 := 0; fi13 < 20; fi13++ {
                v4 = v4
                if false {
                    var v14 uint = 42
                    fmt.Println(v4, fi13, fi13)
                    var v15 uint16 = 65535
                    _ = v14
                    _ = v15
                } else {
                    var v16 uint8 = 177
                    v4 -= (v4 - v4)
                    fmt.Println(fi13, v16)
                    _ = v16
                }
                if (494 == 42) {
                    v17 := GenS1{F0: (false && false), F1: (((("a" + "\n\t") + "\n\t") + "foo bar") + "\n\t"), F2: 42}
                    _ = v17
                }
                _ = fi13
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (127 < 469) {
                if false {
                    _ = len(v3)
                    var v18 uint64 = 1
                    switch "a" {
                    case "\n\t":
                        var v19 int8 = 43
                        _ = v19
                    }
                    _ = v18
                }
            }
        } else {
            switch 9223372036854775807 {
            case 42:
                var v20 float32 = 0.1
                _ = v20
            default:
                fmt.Println(v4)
            }
            v4 += (v4 + 95)
            for fi21 := 0; fi21 < 19; fi21++ {
                v4 -= v4
                _ = fi21
            }
        }
        _ = v12
    } else {
        v22 := GenS3{F0: 128, F1: 303}
        v23 := GenS2{F0: (v4 / ((v4) | 1)), F1: 42}
        if true {
            switch (("hello" + "") + "\n\t") {
            case "\n\t":
                v24 := GenS4{F0: (^(v4))}
                _ = v24
            case "":
                for fi25 := 0; fi25 < 5; fi25++ {
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v27 := GenS4{F0: v4}
                        v28 := GenS5{F0: (v4 ^ 1), F1: (fi25 & fi26), F2: 1}
                        var v29 string = ((("foo bar" + "hello") + "a") + "a")
                        v30 := GenS6{F0: 0, F1: 615, F2: (v29 + v29)}
                        if (false || true) {
                            v31 := GenS1{F0: true, F1: v29, F2: 2}
                            _ = v31
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    _ = fi25
                }
            }
            v4 = v4
            for fi32 := 0; fi32 < 8; fi32++ {
                switch "" {
                case "\n\t":
                    switch ("" + "hello") {
                    case "hello":
                        switch 127 {
                        case 128:
                            if (!(!false)) {
                                if true {
                                    fmt.Println(v4, fi32)
                                }
                                if true {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                switch "foo bar" {
                                case "":
                                    v4 += v4
                                case "foo bar":
                                    if ((1e10 == 1.5) || (1 != -215)) {
                                        fmt.Println(fi32)
                                    }
                                }
                            } else {
                                if (25 > 127) {
                                    if (!false) {
                                        v33 := GenS5{F0: (^(v4)), F1: fi32, F2: 0}
                                        _ = v33
                                    }
                                    v34 := GenS1{F0: (!(127 > fi32)), F1: "foo bar", F2: 32767}
                                    _ = len(v3)
                                    _ = v34
                                }
                                if true {
                                    delete(v2, "hello")
                                } else {
                                    for fi35 := 0; fi35 < 18; fi35++ {
                                        v4 -= (v4 >> ((0) & 63))
                                        var v36 float32 = 0.0
                                        fmt.Println(v36, v4, v36)
                                        _ = fi35
                                        _ = v36
                                    }
                                }
                            }
                        case 0:
                            var v37 uint = 255
                            _ = v37
                        case 9223372036854775807:
                            var v38 int16 = v4
                            _ = v38
                        default:
                            var v39 bool = (65535 < 0)
                            _ = v39
                        }
                    }
                default:
                    switch "" {
                    case "\n\t":
                        v4 -= (v4 % ((v4) | 1))
                    case "a":
                        var v40 int8 = 42
                        _ = v40
                    }
                }
                if (0.1 > 1e10) {
                    v4 = v4
                    var v41 int = fi32
                    fmt.Println(v4, fi32, v4)
                    _ = v41
                } else {
                    v1[291] = ""
                    var v42 int64 = 1
                    if false {
                        switch fi32 {
                        case 32767:
                            for fi43 := 0; fi43 < 7; fi43++ {
                                if (!true) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v42 = 1
                                } else {
                                    if true {
                                        v3 = append(v3, 0)
                                        v4 = (v4 % (((-(v4))) | 1))
                                        if false {
                                            v4 += v4
                                        } else {
                                            v42 = v42
                                            v4 -= 32767
                                            v44 := GenS6{F0: 763, F1: 0, F2: ("hello" + "a")}
                                            if true {
                                                fmt.Println(fi43, fi32)
                                                if false {
                                                    var v45 int64 = (-(v42))
                                                    switch (fi43 >> ((1) & 63)) {
                                                    case 9223372036854775807:
                                                        if false {
                                                            var v46 uint32 = 2147483648
                                                            _ = v46
                                                        }
                                                    case -236:
                                                        var v47 uint64 = 0
                                                        _ = v47
                                                    default:
                                                        switch ("a" + ("foo bar" + (("\n\t" + "\n\t") + ("a" + "hello")))) {
                                                        case "foo bar":
                                                            var v48 int16 = (v4 ^ int16(fi43))
                                                            _ = v48
                                                        case "a":
                                                            var v49 int = (^(fi43))
                                                            _ = v49
                                                        case "hello":
                                                            var v50 uint16 = 370
                                                            _ = v50
                                                        }
                                                    }
                                                    switch "a" {
                                                    case "a":
                                                        v51 := GenS7{F0: 210, F1: 127}
                                                        _ = v51
                                                    }
                                                    _ = len(v3)
                                                    _ = v45
                                                } else {
                                                    fmt.Println(fi32)
                                                    v4 = v4
                                                    if false {
                                                        fmt.Println(fi43, fi32)
                                                        fmt.Println(fi43, v42, v4)
                                                    } else {
                                                        v4 -= v4
                                                        fmt.Println(fi43, v42)
                                                        var v52 uint8 = 128
                                                        _ = v52
                                                    }
                                                }
                                                var v53 int16 = (v4 << ((173) & 63))
                                                _ = v53
                                            }
                                            _ = v44
                                        }
                                    }
                                    v54 := GenS8{F0: 1}
                                    _ = v54
                                }
                                fmt.Println(fi43, fi32, fi43)
                                fmt.Println(fi43, fi43, v42)
                                _ = fi43
                            }
                        }
                        fmt.Println(v42)
                    }
                    _ = v42
                }
                _ = fi32
            }
        }
        _ = v22
        _ = v23
    }
    v55 := GenS5{F0: (-(v4)), F1: -9223372036854775808, F2: 113}
    v56 := GenS9{F0: (127 != 162), F1: (0.1 < (-2.25 + (((1.5 + 0.0) * 0.1) * -2.25)))}
    if (139 == 92) {
        if len(v3) > 0 {
            v3[0] = 623
        }
    }
    var v57 int = 931
    fmt.Println(v4, v57)
    switch "" {
    case "hello":
        v58 := GenS10{F0: (v4 % ((v4) | 1)), F1: 164, F2: 4294967295}
        _ = v58
    case "a":
        switch -9223372036854775808 {
        case -492:
            if (false && false) {
                if true {
                    v4 = 42
                } else {
                    for fi59 := 0; fi59 < 1; fi59++ {
                        for fi60 := 0; fi60 < 8; fi60++ {
                            v4 = (v4 / ((v4) | 1))
                            switch "\n\t" {
                            case "a":
                                var v61 float32 = 0.0
                                _ = v61
                            case "\n\t":
                                v62 := GenS4{F0: v4}
                                _ = v62
                            default:
                                fmt.Println(v57, v57)
                            }
                            _ = fi60
                        }
                        _ = len(v3)
                        v4 += v4
                        _ = fi59
                    }
                }
            }
        case -591:
            v63 := GenS9{F0: (false || false), F1: (42 != v57)}
            _ = v63
        default:
            v57 = 1
        }
    case "\n\t":
        fmt.Println(v4, v4)
    default:
        var v64 string = "a"
        _ = v64
    }
    var v65 bool = true
    for fi66 := 0; fi66 < 18; fi66++ {
        fmt.Println(v57)
        delete(v1, 42)
        _ = fi66
    }
    delete(v1, 2)
    var v67 uint64 = 979
    if true {
        switch "\n\t" {
        case "a":
            for fi68 := 0; fi68 < 15; fi68++ {
                if len(v3) > 0 {
                    v3[0] = fi68
                }
                fmt.Println(v67, fi68)
                fmt.Println(fi68)
                _ = fi68
            }
        case "":
            if (v65 || false) {
                _ = len(v3)
                fmt.Println(v57)
            }
        default:
            var v69 bool = (true && (-327 <= 2147483647))
            _ = v69
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
    _ = v65
    _ = v67
}
