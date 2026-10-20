package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
type GenS2 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 uint; F1 uint64; F2 uint }
type GenS8 struct { F0 uint }

func main() {
    var v1 uint = 4294967295
    var v2 int64 = 42
    v3 := make(map[string]int)
    v4 := []int{}
    if false {
        v4 = append(v4, 1)
        fmt.Println(v2)
    } else {
        switch ("hello" + "hello") {
        case "hello":
            for fi5 := 0; fi5 < 12; fi5++ {
                for fi6 := 0; fi6 < 15; fi6++ {
                    _ = len(v4)
                    _ = len(v3)
                    _ = fi6
                }
                var v7 float32 = float32((3.14159 + (float64(0.1) - -2.25)))
                _ = fi5
                _ = v7
            }
        case "a":
            v1 = v1
        case "\n\t":
            if (394 <= 0) {
                switch int(v1) {
                case -9223372036854775808:
                    var v8 int64 = (v2 & (v2 * v2))
                    _ = v8
                case 62:
                    v1 += v1
                default:
                    if len(v4) > 0 {
                        v4[0] = 32768
                    }
                }
                for fi9 := 0; fi9 < 19; fi9++ {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        switch "" {
                        case "hello":
                            v2 += -1
                        case "\n\t":
                            fmt.Println(v2, fi10)
                        }
                        var v11 int32 = -2147483648
                        _ = fi10
                        _ = v11
                    }
                    switch fi9 {
                    case 42:
                        if len(v4) > 0 {
                            v4[0] = fi9
                        }
                    default:
                        switch "hello" {
                        case "hello":
                            switch ((("foo bar" + ("foo bar" + "foo bar")) + "foo bar") + ("a" + "\n\t")) {
                            case "\n\t":
                                if true {
                                    for fi12 := 0; fi12 < 10; fi12++ {
                                        var v13 uint16 = 42
                                        _ = fi12
                                        _ = v13
                                    }
                                    v14 := GenS1{F0: 132, F1: -806, F2: ("\n\t" + (("" + ("a" + "hello")) + "foo bar"))}
                                    _ = v14
                                }
                            case "hello":
                                var v15 int64 = (v2 | (^(v2)))
                                _ = v15
                            }
                        case "\n\t":
                            var v16 float64 = (1e10 * float64(float32((0.1 * 1e10))))
                            _ = v16
                        default:
                            if ((0.0 + (1.5 + (1.5 - 3.14159))) < 0.0) {
                                if (215 < 255) {
                                    if false {
                                        if true {
                                            var v17 int64 = v2
                                            v18 := v3["foo bar"]
                                            switch v18 {
                                            case 32767:
                                                switch ("a" + "foo bar") {
                                                case "a":
                                                    if false {
                                                        for fi19 := 0; fi19 < 19; fi19++ {
                                                            var v20 uint = (v1 % ((v1) | 1))
                                                            _ = fi19
                                                            _ = v20
                                                        }
                                                        if (((true && ((false || (!(uint8(v1) >= 183))) || false)) && false) && true) {
                                                            v21 := GenS1{F0: 1, F1: (-(fi9)), F2: "a"}
                                                            var v22 uint32 = 65535
                                                            var v23 uint64 = 42
                                                            _ = v21
                                                            _ = v22
                                                            _ = v23
                                                        }
                                                    }
                                                case "":
                                                    switch 42 {
                                                    case 127:
                                                        fmt.Println(v2)
                                                    }
                                                case "hello":
                                                    for fi24 := 0; fi24 < 17; fi24++ {
                                                        var v25 uint16 = 65535
                                                        v2 = (^(v2))
                                                        _ = fi24
                                                        _ = v25
                                                    }
                                                default:
                                                    if (!true) {
                                                        var v26 uint64 = 2
                                                        v27 := GenS2{F0: ((-2.25 <= (3.14159 / ((1e10 + 3.14159) - 1e10))) || (1000 >= 833)), F1: 86, F2: 1}
                                                        _ = v26
                                                        _ = v27
                                                    } else {
                                                        if false {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            delete(v3, "a")
                                                            v18 = fi9
                                                        }
                                                        var v28 int32 = 42
                                                        _ = v28
                                                    }
                                                }
                                            }
                                            _ = v17
                                            _ = v18
                                        } else {
                                            if (0 <= -1) {
                                                v2 = (v2 - (^(v2)))
                                                v2 += (v2 << (((v1 << ((v1) & 63))) & 63))
                                            } else {
                                                for fi29 := 0; fi29 < 8; fi29++ {
                                                    var v30 float64 = (0.0 * 3.14159)
                                                    var v31 uint32 = 42
                                                    _ = fi29
                                                    _ = v30
                                                    _ = v31
                                                }
                                                if (0 != 1) {
                                                    v1 -= (v1 | 4294967295)
                                                } else {
                                                    if (362 != 371) {
                                                        for fi32 := 0; fi32 < 6; fi32++ {
                                                            switch fi9 {
                                                            case 127:
                                                                v1 = v1
                                                            case -2147483648:
                                                                switch "a" {
                                                                case "a":
                                                                    v3["foo bar"] = (fi32 >> ((65535) & 63))
                                                                case "":
                                                                    var v33 int16 = 42
                                                                    fmt.Println(v1, v2, v33)
                                                                    _ = v33
                                                                default:
                                                                    v2 = -9223372036854775808
                                                                }
                                                            }
                                                            _ = fi32
                                                        }
                                                        v2 = v2
                                                        var v34 string = "foo bar"
                                                        _ = v34
                                                    } else {
                                                        switch fi9 {
                                                        case 191:
                                                            var v35 uint64 = 0
                                                            _ = v35
                                                        }
                                                        if ("" >= ("" + "hello")) {
                                                            v36 := GenS3{F0: 52, F1: 32767}
                                                            _ = v36
                                                        } else {
                                                            switch ((("a" + "") + ("a" + "a")) + ("" + (("\n\t" + (("" + ("a" + "a")) + "\n\t")) + ("hello" + ("" + (("" + "\n\t") + "")))))) {
                                                            case "":
                                                                var v37 uint16 = 1
                                                                _ = v37
                                                            case "foo bar":
                                                                var v38 int = 0
                                                                _ = v38
                                                            case "hello":
                                                                var v39 float32 = 1e10
                                                                _ = v39
                                                            default:
                                                                for fi40 := 0; fi40 < 8; fi40++ {
                                                                    if (!true) {
                                                                        if true {
                                                                            v41 := v3[("" + "hello")]
                                                                            if (true || (!(-348 < -2147483648))) {
                                                                                switch 2147483647 {
                                                                                case -9223372036854775808:
                                                                                    v42 := GenS4{F0: 1}
                                                                                    _ = v42
                                                                                case -2147483648:
                                                                                    if true {
                                                                                        var v43 int64 = v2
                                                                                        v3[""] = (fi9 | v41)
                                                                                        var v44 uint64 = 0
                                                                                        _ = v43
                                                                                        _ = v44
                                                                                    }
                                                                                default:
                                                                                    v1 = (v1 / ((490) | 1))
                                                                                }
                                                                            } else {
                                                                                v1 = (^(v1))
                                                                            }
                                                                            _ = v41
                                                                        }
                                                                        var v45 uint8 = 255
                                                                        fmt.Println(v45)
                                                                        var v46 uint16 = 65535
                                                                        _ = v45
                                                                        _ = v46
                                                                    } else {
                                                                        var v47 int32 = 1
                                                                        _ = v47
                                                                    }
                                                                    v48 := GenS5{F0: ""}
                                                                    _ = fi40
                                                                    _ = v48
                                                                }
                                                            }
                                                        }
                                                        var v49 float64 = (0.1 / 1.5)
                                                        _ = v49
                                                    }
                                                    switch "a" {
                                                    case "a":
                                                        if len(v4) > 0 {
                                                            v4[0] = 42
                                                        }
                                                    case "hello":
                                                        var v50 int32 = 1
                                                        _ = v50
                                                    default:
                                                        v51 := v3["hello"]
                                                        _ = v51
                                                    }
                                                }
                                            }
                                        }
                                        var v52 int32 = -2147483648
                                        _ = len(v3)
                                        _ = v52
                                    } else {
                                        for fi53 := 0; fi53 < 19; fi53++ {
                                            fmt.Println(v1, v2)
                                            if len(v4) > 0 {
                                                v4[0] = (fi53 & 1)
                                            }
                                            v2 = (v2 & -9223372036854775808)
                                            _ = fi53
                                        }
                                        var v54 bool = false
                                        _ = v54
                                    }
                                } else {
                                    for fi55 := 0; fi55 < 4; fi55++ {
                                        var v56 string = "\n\t"
                                        fmt.Println(fi9)
                                        var v57 uint8 = 1
                                        _ = fi55
                                        _ = v56
                                        _ = v57
                                    }
                                }
                            } else {
                                switch (fi9 >> ((v1) & 63)) {
                                case 42:
                                    fmt.Println(v2, fi9)
                                case 32768:
                                    var v58 int16 = 0
                                    _ = v58
                                }
                                v1 += (^(v1))
                            }
                        }
                    }
                    var v59 uint8 = 128
                    _ = fi9
                    _ = v59
                }
                for fi60 := 0; fi60 < 14; fi60++ {
                    if true {
                        for fi61 := 0; fi61 < 17; fi61++ {
                            fmt.Println(v1, v2, fi61)
                            switch "foo bar" {
                            case "hello":
                                _ = len(v4)
                            case "":
                                delete(v3, "\n\t")
                            case "foo bar":
                                v3["foo bar"] = 1
                            default:
                                v1 += v1
                            }
                            _ = fi61
                        }
                        for fi62 := 0; fi62 < 15; fi62++ {
                            v63 := GenS4{F0: (fi60 ^ fi62)}
                            v2 = v2
                            _ = fi62
                            _ = v63
                        }
                    } else {
                        fmt.Println(fi60, v2)
                        v64 := v3["a"]
                        _ = v64
                    }
                    var v65 uint8 = 42
                    fmt.Println(fi60, v2)
                    _ = fi60
                    _ = v65
                }
            }
        default:
            v1 += v1
        }
        v2 = v2
        v4 = append(v4, 2147483647)
    }
    switch 42 {
    case 1:
        if (int32(v1) <= 0) {
            v66 := GenS6{F0: 1, F1: 0}
            v67 := GenS6{F0: 603, F1: -697}
            if true {
                if (0 == 0) {
                    var v68 string = ("hello" + "\n\t")
                    v4 = append(v4, 32767)
                    _ = v68
                }
                if (1 >= 152) {
                    v4 = append(v4, -9223372036854775808)
                    var v69 int = -20
                    v1 -= (v1 | 255)
                    _ = v69
                }
            } else {
                _ = len(v3)
                if false {
                    var v70 int = int(v1)
                    _ = len(v4)
                    _ = v70
                } else {
                    fmt.Println(v1, v2)
                    if len(v4) > 0 {
                        v4[0] = 154
                    }
                }
            }
            _ = v66
            _ = v67
        }
    default:
        v71 := GenS3{F0: 127, F1: 32767}
        _ = v71
    }
    var v72 uint8 = 115
    var v73 uint8 = (v72 >> ((v1) & 63))
    var v74 int64 = v2
    fmt.Println(v73)
    v1 = (v1 << ((v1) & 63))
    for fi75 := 0; fi75 < 7; fi75++ {
        var v76 uint32 = 364
        v72 += (v73 % (((^(v73))) | 1))
        v1 += v1
        _ = fi75
        _ = v76
    }
    for fi77 := 0; fi77 < 13; fi77++ {
        if ("a" != "hello") {
            v78 := GenS7{F0: (^(v1)), F1: 743, F2: v1}
            _ = v78
        } else {
            fmt.Println(v2)
            v1 -= (v1 & 490)
        }
        if ((false && (32767 > 42)) || (false && (!false))) {
            var v79 string = "a"
            if ((!((100 <= 103) || ((v1 > v1) || (233 != 1)))) || false) {
                for fi80 := 0; fi80 < 10; fi80++ {
                    if len(v4) > 0 {
                        v4[0] = fi77
                    }
                    _ = fi80
                }
            } else {
                v1 = (v1 % ((2) | 1))
                for fi81 := 0; fi81 < 9; fi81++ {
                    fmt.Println(fi77)
                    _ = fi81
                }
            }
            _ = v79
        } else {
            v72 += (^(v73))
        }
        var v82 int32 = 661
        _ = fi77
        _ = v82
    }
    for fi83 := 0; fi83 < 19; fi83++ {
        v74 = v2
        if (false || false) {
            fmt.Println(v73)
            v84 := GenS7{F0: v1, F1: 1, F2: (^(v1))}
            v2 = (-(v74))
            _ = v84
        }
        v1 = v1
        _ = fi83
    }
    fmt.Println(v72, v73, v73)
    v85 := GenS8{F0: v1}
    v72 = v73
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v72
    _ = v73
    _ = v74
    _ = v85
}
