package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int16; F1 int; F2 uint16 }
type GenS12 struct { F0 string; F1 string; F2 int16 }

func main() {
    v1 := []uint32{2147483648, 2, 27, 2}
    v2 := []uint32{4294967295, 1, 2147483647, 175}
    v3 := []string{""}
    var v4 uint = 0
    v5 := []int{}
    var v6 uint16 = 42
    v4 += (v4 / ((v4) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v7 uint64 = 937
    var v8 float32 = 3.14159
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v7 = v7
            var v11 int8 = -128
            v12 := GenS1{F0: 255, F1: 1, F2: (v6 % (((^(v6))) | 1))}
            fmt.Println(v4, v11)
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS1{F0: 1, F1: 0, F2: (v6 >> ((v4) & 63))}
        _ = fi13
        _ = v14
    }
    if (float64((0.0 + (v8 / v8))) != ((0.0 + (-2.25 - (1.5 * 1e10))) + 1.5)) {
        v8 = ((-2.25 + -2.25) * 1e10)
        var v15 int = 128
        _ = v15
    } else {
        v7 = v7
        v6 += 32767
        if false {
            v16 := GenS1{F0: 200, F1: 838, F2: v6}
            fmt.Println(v6, v4, v8)
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        fmt.Println(v4, v7, v4)
        _ = len(v3)
        v7 = 42
        _ = fi17
    }
    switch "hello" {
    case "":
        for fi18 := 0; fi18 < 6; fi18++ {
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 float32 = (1.5 * float32(-2.25))
                if (1 != 0) {
                    if (true || true) {
                        v21 := GenS2{F0: 1}
                        v8 += float32(-2.25)
                        _ = v21
                    }
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    v22 := GenS1{F0: 0, F1: -364, F2: 778}
                    _ = v22
                } else {
                    if (int64(v6) != -9223372036854775808) {
                        switch 0 {
                        case 37:
                            switch "hello" {
                            case "hello":
                                fmt.Println(v20)
                            default:
                                v2 = append(v2, 2147483647)
                                switch ("" + "\n\t") {
                                case "\n\t":
                                    fmt.Println(v6)
                                case "foo bar":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    v23 := GenS3{F0: "foo bar", F1: (v6 & v6), F2: 1}
                                    _ = v23
                                }
                                fmt.Println(v6)
                            }
                        case -1:
                            if false {
                                _ = len(v3)
                            }
                        case -277:
                            v4 = 1
                        default:
                            switch ("hello" + ("" + ("hello" + "a"))) {
                            case "foo bar":
                                switch "a" {
                                case "foo bar":
                                    v7 -= v7
                                case "a":
                                    switch (fi19 * 1) {
                                    case -1:
                                        if (209 == 2) {
                                            fmt.Println(v20, v7, v20)
                                        } else {
                                            if ((!(0 <= fi19)) || true) {
                                                fmt.Println(v6)
                                                switch "hello" {
                                                case "\n\t":
                                                    v20 += v8
                                                case "hello":
                                                    v4 -= 4294967295
                                                case "foo bar":
                                                    var v24 string = ("" + "a")
                                                    _ = v24
                                                default:
                                                    switch "hello" {
                                                    case "a":
                                                        if true {
                                                            var v25 float64 = -2.25
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = v25
                                                        } else {
                                                            v26 := GenS4{F0: v6, F1: 255}
                                                            fmt.Println(fi18, fi19, v4)
                                                            _ = v26
                                                        }
                                                    case "foo bar":
                                                        if true {
                                                            var v27 int = fi18
                                                            _ = v27
                                                        }
                                                    case "\n\t":
                                                        v8 = 0.0
                                                    default:
                                                        v20 += v8
                                                    }
                                                }
                                            } else {
                                                v20 -= float32((1.5 + 3.14159))
                                                switch "" {
                                                case "":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case "foo bar":
                                                    var v28 uint16 = (v6 * v6)
                                                    _ = v28
                                                case "a":
                                                    _ = len(v2)
                                                    var v29 bool = (false || false)
                                                    if ((^(v7)) > v7) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    if (v7 <= v7) {
                                                        var v30 float64 = float64((v8 / 1.0))
                                                        _ = v30
                                                    }
                                                    _ = v29
                                                }
                                            }
                                        }
                                    case 2147483647:
                                        fmt.Println(fi18, v4)
                                    case 127:
                                        v8 += ((v8 * v8) + ((1.5 * 0.0) * v20))
                                    default:
                                        fmt.Println(v7, v8, fi19)
                                        if (127 == 100) {
                                            v31 := GenS5{F0: (v7 | (v7 + v7))}
                                            _ = v31
                                        }
                                        v32 := GenS1{F0: 1, F1: 1000, F2: uint16(v7)}
                                        _ = v32
                                    }
                                default:
                                    v33 := GenS6{F0: -128}
                                    _ = v33
                                }
                            case "":
                                var v34 uint = (^(v4))
                                _ = v34
                            case "hello":
                                var v35 uint32 = 471
                                _ = v35
                            default:
                                var v36 string = "hello"
                                _ = v36
                            }
                            var v37 int = (-(fi19))
                            _ = v37
                        }
                        _ = len(v5)
                    } else {
                        v38 := GenS1{F0: 26, F1: 1, F2: v6}
                        v6 = (v6 << ((uint(v6)) & 63))
                        _ = v38
                    }
                }
                v1 = append(v1, uint32(v6))
                switch -306 {
                case 128:
                    v8 = (v20 + (v20 * v20))
                case 42:
                    v39 := GenS7{F0: false, F1: false}
                    _ = v39
                default:
                    var v40 uint8 = 127
                    _ = v40
                }
                fmt.Println(v20, fi19)
                v41 := GenS2{F0: -32768}
                _ = fi19
                _ = v20
                _ = v41
            }
            v3 = append(v3, ("\n\t" + ""))
            for fi42 := 0; fi42 < 4; fi42++ {
                if false {
                    if len(v5) > 0 {
                        v5[0] = (fi42 & 32767)
                    }
                    switch int(v7) {
                    case 9223372036854775807:
                        v6 -= v6
                    case 32768:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        fmt.Println(fi18, v4, v4)
                    }
                    v3 = append(v3, ("foo bar" + ("a" + "foo bar")))
                    fmt.Println(v8, v4, v6)
                    if (("foo bar" == ((("" + "foo bar") + "\n\t") + "\n\t")) || true) {
                        v8 -= 0.0
                        _ = len(v5)
                        if (((3.14159 + v8) / (v8 - v8)) > float32(-2.25)) {
                            var v43 int = fi18
                            if true {
                                var v44 int32 = 114
                                _ = v44
                            }
                            _ = v43
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        if (true || true) {
                            v8 = float32(0.0)
                        } else {
                            v45 := GenS8{F0: true, F1: 42}
                            fmt.Println(v7)
                            fmt.Println(fi18, v4, v4)
                            _ = v45
                        }
                    }
                    switch (fi42 % (((^(fi42))) | 1)) {
                    case 243:
                        var v46 float64 = 1.5
                        _ = v46
                    case 32767:
                        v47 := GenS9{F0: 100}
                        _ = v47
                    case 9223372036854775807:
                        v5 = append(v5, (fi18 ^ 42))
                        if (!true) {
                            var v48 float64 = ((1e10 / ((-2.25 * float64(float32(0.1))) / 1.0)) - 0.0)
                            var v49 int32 = -520
                            var v50 int = fi42
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                        var v51 int = fi42
                        _ = v51
                    }
                    if ("foo bar" == "hello") {
                        v52 := GenS5{F0: 923}
                        var v53 uint16 = (v6 >> (((v4 / ((255) | 1))) & 63))
                        if ("\n\t" == ("" + ("hello" + "\n\t"))) {
                            if (v53 < (v53 % ((42) | 1))) {
                                v6 = (v53 | 193)
                                fmt.Println(v4, fi18, v7)
                                var v54 int64 = -423
                                _ = v54
                            } else {
                                switch fi18 {
                                case -923:
                                    var v55 uint8 = 1
                                    _ = v55
                                case -735:
                                    v5 = append(v5, (fi18 + fi18))
                                default:
                                    if (-128 < 42) {
                                        v6 = v6
                                    }
                                }
                            }
                            v56 := GenS10{F0: true}
                            _ = v56
                        }
                        v57 := GenS11{F0: 42, F1: -826, F2: (v53 << ((v4) & 63))}
                        _ = v52
                        _ = v53
                        _ = v57
                    }
                }
                v58 := GenS1{F0: 87, F1: 42, F2: (v6 | (v6 >> ((255) & 63)))}
                _ = fi42
                _ = v58
            }
            _ = fi18
        }
    case "foo bar":
        v8 -= (v8 + 3.14159)
    }
    if (true || true) {
        for fi59 := 0; fi59 < 15; fi59++ {
            if false {
                v7 = v7
            } else {
                var v60 float32 = 1e10
                switch "" {
                case "foo bar":
                    switch (^(fi59)) {
                    case -2147483648:
                        v61 := GenS10{F0: (v7 <= v7)}
                        _ = v61
                    case 32768:
                        if true {
                            fmt.Println(v4, v60)
                            fmt.Println(v8, fi59)
                        }
                    default:
                        for fi62 := 0; fi62 < 1; fi62++ {
                            var v63 float64 = 0.1
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            switch ("\n\t" + ("a" + "\n\t")) {
                            case "hello":
                                v60 += 3.14159
                            case "a":
                                var v64 uint64 = v7
                                _ = v64
                            default:
                                var v65 int8 = -9
                                _ = v65
                            }
                            _ = fi62
                            _ = v63
                        }
                    }
                }
                v5 = append(v5, (fi59 << ((v4) & 63)))
                _ = v60
            }
            _ = fi59
        }
        for fi66 := 0; fi66 < 20; fi66++ {
            for fi67 := 0; fi67 < 18; fi67++ {
                v5 = append(v5, fi67)
                fmt.Println(v6, v7, v4)
                _ = fi67
            }
            switch "\n\t" {
            case "":
                v6 += 65535
            case "\n\t":
                for fi68 := 0; fi68 < 5; fi68++ {
                    var v69 int64 = -930
                    _ = fi68
                    _ = v69
                }
            case "foo bar":
                var v70 uint = 4294967295
                _ = v70
            }
            fmt.Println(fi66, v7, fi66)
            _ = fi66
        }
        v8 = (((v8 / float32(1e10)) - (((v8 + v8) - v8) - ((((float32((((float64(-2.25) * (1.5 - 0.0)) * 0.1) * -2.25)) - -2.25) * (1.5 / 1.0)) * v8) + v8))) * 1e10)
    } else {
        if false {
            var v71 uint16 = v6
            var v72 uint64 = 4294967295
            _ = v71
            _ = v72
        }
        switch -970 {
        case -1:
            fmt.Println(v6)
        case -979:
            var v73 uint8 = 128
            _ = v73
        case -9223372036854775808:
            var v74 uint32 = 347
            _ = v74
        }
        var v75 int64 = 877
        _ = v75
    }
    v76 := GenS12{F0: "", F1: "hello", F2: -32768}
    fmt.Println(v7, v6, v6)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v2)
    if (!(false && ((v6 - (v6 & (v6 / ((v6) | 1)))) >= 42))) {
        if false {
            if (!(-1 == int(v7))) {
                fmt.Println(v8, v7, v8)
            }
        } else {
            switch 475 {
            case 9223372036854775807:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                var v77 uint16 = (v6 & 32767)
                _ = v77
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v76
    _ = v8
}
