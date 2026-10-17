package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 bool; F1 int16; F2 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32 }
func genFn0(p0 string, p1 int64, p2 int8) uint8 {
    v4 := []int{1, -526}
    p0 = p0
    switch p0 {
    case "hello":
        v5 := GenS1{F0: ((p2 << ((1) & 63)) > -1), F1: 2147483647}
        _ = v5
    case "a":
        for fi6 := 0; fi6 < 19; fi6++ {
            switch p0 {
            case "a":
                var v7 bool = (42 > -1)
                _ = v7
            case "":
                v8 := GenS1{F0: (true || false), F1: 1}
                _ = v8
            }
            switch 137 {
            case 32767:
                v9 := GenS1{F0: true, F1: 368}
                _ = v9
            }
            fmt.Println(p1, fi6, fi6)
            _ = fi6
        }
    case "foo bar":
        switch p0 {
        case "a":
            p2 -= (p2 + -1)
        }
    default:
        if len(v4) > 0 {
            v4[0] = 32768
        }
    }
    p2 -= (p2 >> ((255) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 134
}
func genFn1(p0 uint) string {
    var v10 uint64 = 291
    v10 = (^(v10))
    var v11 int64 = 9223372036854775807
    _ = p0
    _ = v10
    _ = v11
    return ("hello" + "\n\t")
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]uint32)
    var v3 uint64 = 42
    fmt.Println(v1, v3)
    if true {
        switch "foo bar" {
        case "foo bar":
            var v12 int = -551
            _ = v12
        case "\n\t":
            if true {
                switch ((("hello" + ("a" + "\n\t")) + (((("foo bar" + "hello") + "hello") + "foo bar") + ("" + "hello"))) + "hello") {
                case "a":
                    var v13 uint64 = 0
                    _ = v13
                case "":
                    switch (("hello" + ("a" + "\n\t")) + "hello") {
                    case "foo bar":
                        var v14 float32 = v1
                        _ = v14
                    default:
                        var v15 uint64 = v3
                        _ = v15
                    }
                case "hello":
                    v3 = 311
                }
                switch "\n\t" {
                case "foo bar":
                    v16 := []int{}
                    _ = v16
                default:
                    var v17 float32 = (v1 / v1)
                    _ = v17
                }
                if (false || (0 >= 255)) {
                    for fi18 := 0; fi18 < 13; fi18++ {
                        v3 = v3
                        v19 := GenS1{F0: true, F1: 1}
                        v3 += v3
                        _ = fi18
                        _ = v19
                    }
                    if false {
                        v20 := GenS1{F0: false, F1: 2147483647}
                        _ = v20
                    } else {
                        switch ("a" + "\n\t") {
                        case "\n\t":
                            fmt.Println(v3, v3)
                        case "":
                            switch 1 {
                            case 4294967295:
                                var v21 uint16 = 42
                                _ = v21
                            case -9223372036854775808:
                                if (-128 > 1) {
                                    for fi22 := 0; fi22 < 7; fi22++ {
                                        _ = len(v2)
                                        _ = fi22
                                    }
                                } else {
                                    if true {
                                        v23 := GenS2{F0: true, F1: -673, F2: 0}
                                        _ = v23
                                    } else {
                                        for fi24 := 0; fi24 < 15; fi24++ {
                                            fmt.Println(v1, v3, fi24)
                                            _ = fi24
                                        }
                                        v25 := GenS2{F0: false, F1: 520, F2: 255}
                                        for fi26 := 0; fi26 < 12; fi26++ {
                                            delete(v2, (((("a" + "a") + "") + "hello") + ("\n\t" + "hello")))
                                            _ = fi26
                                        }
                                        _ = v25
                                    }
                                }
                            case 1:
                                var v27 float32 = v1
                                _ = v27
                            default:
                                v28 := v2[""]
                                _ = v28
                            }
                        case "a":
                            for fi29 := 0; fi29 < 9; fi29++ {
                                switch (^(fi29)) {
                                case 9223372036854775807:
                                    delete(v2, "a")
                                }
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    var v31 float64 = 0.0
                                    var v32 uint32 = 323
                                    if false {
                                        fmt.Println(fi29)
                                        var v33 uint = 255
                                        v34 := []int{}
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v35 := GenS2{F0: false, F1: int16(v3), F2: 0}
                                        v36 := v2["\n\t"]
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                                v3 = (^(v3))
                                _ = fi29
                            }
                        }
                    }
                }
            } else {
                fmt.Println(v1, v1)
                v37 := []int{-9223372036854775808, -2147483648}
                var v38 int = -407
                _ = v37
                _ = v38
            }
        case "a":
            v1 = (0.0 - ((v1 * (v1 + 1e10)) - v1))
        }
        var v39 uint64 = v3
        _ = v39
    } else {
        v40 := []int{9223372036854775807, 128, -192, 32768}
        _ = v40
    }
    v1 = (v1 + v1)
    v41 := GenS1{F0: (uint(v3) <= 0), F1: 42}
    v3 = (v3 ^ 0)
    if (42 <= 9223372036854775807) {
        switch 734 {
        case 2147483647:
            if true {
                var v42 int16 = 445
                var v43 int16 = v42
                _ = v42
                _ = v43
            }
        case -9223372036854775808:
            fmt.Println(v1, v1)
        case 1:
            var v44 float64 = -2.25
            _ = v44
        default:
            var v45 string = "\n\t"
            _ = v45
        }
        if true {
            fmt.Println(v3, v1)
            if false {
                var v46 int16 = 32767
                v47 := v2["foo bar"]
                _ = v46
                _ = v47
            } else {
                for fi48 := 0; fi48 < 5; fi48++ {
                    _ = len(v2)
                    _ = fi48
                }
                if (1 == 42) {
                    var v49 bool = (!true)
                    _ = v49
                }
            }
        }
        v50 := []int{128, 574, 1}
        _ = v50
    }
    v1 = 0.0
    if (-2147483648 != -1) {
        v3 -= (v3 | (^(v3)))
        if (0.1 >= 1e10) {
            v51 := GenS1{F0: (2 > 1), F1: -143}
            _ = v51
        } else {
            var v52 int16 = 1000
            if true {
                var v53 int32 = -2147483648
                _ = v53
            }
            _ = v52
        }
        v54 := GenS3{F0: -729}
        _ = v54
    } else {
        fmt.Println(v3)
        switch ((("foo bar" + "hello") + "hello") + "") {
        case "":
            for fi55 := 0; fi55 < 18; fi55++ {
                v56 := GenS4{F0: ""}
                v57 := GenS2{F0: true, F1: -1, F2: 34}
                _ = fi55
                _ = v56
                _ = v57
            }
        case "foo bar":
            fmt.Println(v1)
        }
        fmt.Println(v1)
    }
    v1 = float32(float64(v1))
    for fi58 := 0; fi58 < 14; fi58++ {
        if (0 == 9223372036854775807) {
            v59 := GenS4{F0: "\n\t"}
            v3 -= v3
            v1 = (0.0 / (1e10 * -2.25))
            _ = v59
        } else {
            var v60 float64 = (3.14159 + (1.5 * (1e10 * 1e10)))
            for fi61 := 0; fi61 < 5; fi61++ {
                fmt.Println(v60, v3)
                v62 := GenS5{F0: -2147483648}
                fmt.Println(v60, v3, v1)
                _ = fi61
                _ = v62
            }
            _ = v60
        }
        _ = fi58
    }
    for fi63 := 0; fi63 < 20; fi63++ {
        v64 := GenS3{F0: 962}
        var v65 uint8 = 2
        v66 := GenS6{F0: (v65 - v65), F1: "hello"}
        _ = fi63
        _ = v64
        _ = v65
        _ = v66
    }
    switch "" {
    case "a":
        v1 += (v1 / 1.0)
    case "\n\t":
        fmt.Println(v1, v3)
    case "":
        if (0 < 4294967295) {
            fmt.Println(v1, v1)
        }
    }
    for fi67 := 0; fi67 < 1; fi67++ {
        if (42 < 820) {
            if true {
                delete(v2, ("foo bar" + "foo bar"))
                v1 += v1
            } else {
                switch "hello" {
                case "foo bar":
                    for fi68 := 0; fi68 < 12; fi68++ {
                        v69 := []int{fi68}
                        if false {
                            switch 32767 {
                            case 516:
                                var v70 int = fi68
                                _ = v70
                            case -2147483648:
                                fmt.Println(v1, v3)
                            case 127:
                                v2["\n\t"] = 824
                            default:
                                if true {
                                    switch "hello" {
                                    case "hello":
                                        if true {
                                            _ = len(v2)
                                            switch 4294967295 {
                                            case 128:
                                                v71 := GenS4{F0: "\n\t"}
                                                _ = v71
                                            case -990:
                                                v69 = append(v69, fi68)
                                            case -25:
                                                switch (("\n\t" + "hello") + (("a" + "\n\t") + ("" + ("foo bar" + "hello")))) {
                                                case "hello":
                                                    v72 := GenS7{F0: true, F1: 0, F2: 0}
                                                    _ = v72
                                                default:
                                                    v73 := GenS8{F0: 2147483648}
                                                    _ = v73
                                                }
                                            default:
                                                if (0.0 > 0.1) {
                                                    var v74 bool = ((fi67 + fi67) != 191)
                                                    v3 -= (v3 + v3)
                                                    v1 += -2.25
                                                    v75 := v2["\n\t"]
                                                    var v76 float32 = (1.5 + v1)
                                                    if (42 < 981) {
                                                        v76 += (v1 + 0.1)
                                                        var v77 int8 = -101
                                                        v78 := GenS4{F0: (("hello" + "") + "hello")}
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    _ = v74
                                                    _ = v75
                                                    _ = v76
                                                }
                                            }
                                            var v79 bool = (1 != v3)
                                            _ = v79
                                        } else {
                                            var v80 uint8 = 197
                                            _ = v80
                                        }
                                    }
                                }
                            }
                            v3 = v3
                        } else {
                            var v81 float64 = float64(1.5)
                            v81 = -2.25
                            var v82 bool = ("a" >= "")
                            _ = v81
                            _ = v82
                        }
                        var v83 float32 = float32(1e10)
                        v83 -= v1
                        v84 := v2["foo bar"]
                        fmt.Println(fi68, fi68, v83)
                        _ = fi68
                        _ = v69
                        _ = v83
                        _ = v84
                    }
                case "hello":
                    v85 := []int{(fi67 ^ 32768), 9223372036854775807, fi67}
                    _ = v85
                }
                v86 := GenS9{F0: 0}
                _ = v86
            }
        }
        if true {
            v1 -= 0.0
            v3 -= 1
            var v87 float64 = 0.0
            _ = v87
        }
        v88 := v2["a"]
        _ = fi67
        _ = v88
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
}
