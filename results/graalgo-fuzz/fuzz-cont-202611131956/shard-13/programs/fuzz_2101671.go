package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int16; F1 int; F2 int32 }
type GenS5 struct { F0 uint64; F1 uint8; F2 bool }
type GenS6 struct { F0 int8; F1 string; F2 int16 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int8, p1 uint64) int64 {
    v4 := []int{318}
    p0 = -5
    v5 := GenS1{F0: p0}
    p0 = -128
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            p0 -= (p0 / ((p0) | 1))
            _ = fi7
        }
        p1 -= p1
        for fi8 := 0; fi8 < 1; fi8++ {
            if (!false) {
                switch (("hello" + "") + "\n\t") {
                case "foo bar":
                    var v9 uint32 = 237
                    _ = v9
                case "a":
                    var v10 int32 = 0
                    _ = v10
                }
                p0 = 57
                var v11 uint16 = 1
                _ = v11
            } else {
                if (-1 <= 2147483647) {
                    p0 = p0
                    v4 = append(v4, fi6)
                    fmt.Println(fi6, fi6, fi8)
                }
                v12 := GenS2{F0: (!((p0 != (p0 * (p0 >> ((2) & 63)))) && (true && false))), F1: 938}
                fmt.Println(fi8)
                _ = v12
            }
            var v13 uint = 255
            _ = fi8
            _ = v13
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 262
}

func main() {
    v1 := []int{-1, -334}
    v2 := make(map[string]string)
    var v3 bool = true
    switch "a" {
    case "":
        if v3 {
            if (128 <= -9223372036854775808) {
                if v3 {
                    var v14 int16 = -1
                    _ = v14
                }
                v3 = (genFn0(42, 18446744073709551615) == -9223372036854775808)
            } else {
                if false {
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v3 = true
                        _ = fi15
                    }
                    var v16 int64 = -278
                    _ = v16
                }
            }
        } else {
            v3 = ("\n\t" > "hello")
        }
    case "\n\t":
        v2[("\n\t" + "hello")] = "foo bar"
    }
    _ = len(v2)
    switch 127 {
    case 32767:
        if false {
            var v17 float64 = (1.5 - float64(((-2.25 + ((float32(1e10) * (1.5 * (3.14159 * (float32(0.0) * 0.1)))) * (1.5 - 1e10))) - (0.0 * float32(3.14159)))))
            for fi18 := 0; fi18 < 19; fi18++ {
                v3 = ((v17 == -2.25) && (true || ((false || ("" <= "hello")) && true)))
                for fi19 := 0; fi19 < 12; fi19++ {
                    v20 := GenS2{F0: (791 > 0), F1: 42}
                    _ = fi19
                    _ = v20
                }
                v17 -= (float64(3.14159) - v17)
                _ = fi18
            }
            v21 := GenS2{F0: (!(true || v3)), F1: -556}
            _ = v17
            _ = v21
        } else {
            fmt.Println(v3)
            if true {
                fmt.Println(v3)
            }
            v22 := v2["\n\t"]
            _ = v22
        }
    case 42:
        var v23 uint32 = 2147483648
        _ = v23
    default:
        var v24 int = 128
        _ = v24
    }
    if v3 {
        v25 := v2["hello"]
        _ = v25
    } else {
        v26 := GenS3{F0: 65535, F1: "\n\t"}
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 uint16 = 6
            fmt.Println(v3)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi27
            _ = v28
        }
        switch 0 {
        case -2147483648:
            for fi29 := 0; fi29 < 6; fi29++ {
                v30 := GenS1{F0: 127}
                var v31 uint64 = 42
                _ = fi29
                _ = v30
                _ = v31
            }
        case 128:
            switch -1 {
            case 128:
                fmt.Println(v3)
            }
        }
        _ = v26
    }
    v3 = v3
    if (("hello" + "foo bar") != ("hello" + "\n\t")) {
        var v32 int64 = -9223372036854775808
        for fi33 := 0; fi33 < 1; fi33++ {
            switch "\n\t" {
            case "":
                v32 = (v32 ^ genFn0(1, 4294967295))
            case "hello":
                switch (fi33 | -1) {
                case 2147483647:
                    switch "\n\t" {
                    case "foo bar":
                        for fi34 := 0; fi34 < 2; fi34++ {
                            switch fi33 {
                            case -2147483648:
                                v35 := GenS3{F0: 0, F1: "a"}
                                _ = v35
                            }
                            v3 = (1000 <= 1)
                            _ = fi34
                        }
                    default:
                        var v36 int32 = 2147483647
                        _ = v36
                    }
                case 164:
                    v37 := GenS1{F0: 127}
                    _ = v37
                case -745:
                    fmt.Println(v3, v3, fi33)
                }
            }
            v3 = v3
            _ = fi33
        }
        for fi38 := 0; fi38 < 13; fi38++ {
            if (-715 != 1) {
                var v39 float64 = -2.25
                for fi40 := 0; fi40 < 7; fi40++ {
                    v32 = v32
                    if v3 {
                        fmt.Println(v32, v32)
                    } else {
                        _ = len(v2)
                        if len(v1) > 0 {
                            v1[0] = -318
                        }
                    }
                    switch fi38 {
                    case -621:
                        if (!true) {
                            switch (("\n\t" + "") + "\n\t") {
                            case "a":
                                v39 = v39
                                v41 := GenS2{F0: v3, F1: 1}
                                if v3 {
                                    var v42 uint32 = 0
                                    v39 += (v39 + 0.0)
                                    _ = v42
                                } else {
                                    switch (("hello" + "hello") + "\n\t") {
                                    case "a":
                                        if (!true) {
                                            var v43 uint64 = 0
                                            v2["a"] = "hello"
                                            _ = v43
                                        }
                                    case "\n\t":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        if (!v3) {
                                            v44 := GenS4{F0: -429, F1: fi40, F2: 42}
                                            _ = v44
                                        } else {
                                            v39 -= v39
                                            if v3 {
                                                if (2 >= 255) {
                                                    delete(v2, ((("hello" + "a") + "\n\t") + ("" + "hello")))
                                                    var v45 int64 = (v32 ^ v32)
                                                    _ = v45
                                                }
                                            }
                                        }
                                    }
                                    var v46 int32 = -2147483648
                                    _ = v46
                                }
                                _ = v41
                            case "":
                                fmt.Println(v32)
                            case "hello":
                                var v47 uint = 4294967295
                                _ = v47
                            default:
                                if v3 {
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v39 += 1.5
                                        v48 := GenS5{F0: 18446744073709551615, F1: 127, F2: (0 <= 2147483647)}
                                        _ = v48
                                    } else {
                                        v39 += 1e10
                                        var v49 int = (fi40 >> ((2) & 63))
                                        v39 += v39
                                        _ = v49
                                    }
                                }
                            }
                        } else {
                            _ = len(v1)
                        }
                    case 32768:
                        fmt.Println(fi38, fi38, v39)
                    case 32767:
                        v50 := GenS6{F0: -128, F1: "foo bar", F2: -32768}
                        _ = v50
                    }
                    _ = fi40
                }
                fmt.Println(v32, v32, v3)
                _ = v39
            }
            _ = fi38
        }
        _ = v32
    }
    v51 := GenS7{F0: 89}
    if ("\n\t" < "") {
        if true {
            if v3 {
                _ = len(v1)
                switch 42 {
                case 894:
                    fmt.Println(v3)
                case 32767:
                    var v52 int = 0
                    _ = v52
                case 1:
                    fmt.Println(v3)
                }
            }
        } else {
            for fi53 := 0; fi53 < 15; fi53++ {
                if (42 < 0) {
                    fmt.Println(fi53, fi53)
                    var v54 float32 = (0.1 / 1e10)
                    switch (-(fi53)) {
                    case 9223372036854775807:
                        for fi55 := 0; fi55 < 18; fi55++ {
                            fmt.Println(v54)
                            v54 = ((v54 + v54) * 0.0)
                            fmt.Println(v54, fi55)
                            _ = fi55
                        }
                    case -2147483648:
                        if len(v1) > 0 {
                            v1[0] = (-(fi53))
                        }
                    }
                    _ = v54
                } else {
                    v3 = false
                    v3 = ((549 <= 2147483648) || true)
                    v3 = v3
                }
                var v56 float32 = 1.5
                v3 = ((!(32767 <= 493)) && v3)
                _ = fi53
                _ = v56
            }
        }
        v3 = v3
        v3 = false
    }
    delete(v2, ("" + "hello"))
    for fi57 := 0; fi57 < 1; fi57++ {
        fmt.Println(fi57, v3)
        _ = fi57
    }
    if (0.1 >= 1.5) {
        fmt.Println(v3)
    } else {
        v2["\n\t"] = "\n\t"
        var v58 uint32 = 2147483648
        for fi59 := 0; fi59 < 7; fi59++ {
            for fi60 := 0; fi60 < 1; fi60++ {
                fmt.Println(fi59)
                _ = fi60
            }
            fmt.Println(v3)
            _ = fi59
        }
        _ = v58
    }
    switch "hello" {
    case "hello":
        switch 32768 {
        case -2147483648:
            v61 := v2[""]
            _ = v61
        case -1:
            fmt.Println(v3)
        case 9223372036854775807:
            switch "hello" {
            case "foo bar":
                var v62 bool = (((3.14159 * float32(0.0)) + -2.25) == 0.1)
                _ = v62
            case "":
                var v63 uint = 42
                _ = v63
            case "\n\t":
                if v3 {
                    if len(v1) > 0 {
                        v1[0] = 127
                    }
                    if (false && (293 <= 724)) {
                        v64 := v2[("hello" + "hello")]
                        v64 = "hello"
                        _ = v64
                    }
                }
            default:
                switch "foo bar" {
                case "hello":
                    var v65 bool = true
                    _ = v65
                case "foo bar":
                    v3 = v3
                default:
                    var v66 float64 = 3.14159
                    _ = v66
                }
            }
        default:
            switch "" {
            case "\n\t":
                for fi67 := 0; fi67 < 2; fi67++ {
                    fmt.Println(v3, fi67)
                    _ = fi67
                }
            case "a":
                var v68 uint64 = 4294967295
                _ = v68
            case "":
                v69 := GenS6{F0: 1, F1: "a", F2: -32768}
                _ = v69
            }
        }
    case "\n\t":
        v3 = false
    case "":
        if v3 {
            var v70 int64 = 0
            var v71 float32 = -2.25
            _ = v70
            _ = v71
        }
    default:
        v3 = (!(-2.25 <= float32(-2.25)))
    }
    v1 = append(v1, 127)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
}
