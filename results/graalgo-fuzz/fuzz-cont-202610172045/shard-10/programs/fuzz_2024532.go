package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint }

func main() {
    var v1 uint8 = 255
    v2 := []uint32{65535}
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    v5 := GenS1{F0: 0, F1: "a", F2: 953}
    delete(v3, "a")
    v6 := GenS2{F0: 255, F1: 1, F2: 1}
    v7 := GenS2{F0: 560, F1: 2147483647, F2: 2}
    if false {
        var v8 uint8 = (v1 - v1)
        var v9 uint32 = 2147483648
        _ = v8
        _ = v9
    }
    switch 876 {
    case 309:
        if true {
            var v10 uint = 42
            switch 42 {
            case 32767:
                if (42 != 0) {
                    var v11 float32 = 0.1
                    _ = v11
                }
            }
            _ = v10
        }
    case 42:
        v1 += 127
    default:
        fmt.Println(v1)
    }
    switch (("hello" + "a") + "a") {
    case "\n\t":
        for fi12 := 0; fi12 < 1; fi12++ {
            switch fi12 {
            case 0:
                if (false || true) {
                    v13 := GenS1{F0: 42, F1: "a", F2: 1}
                    fmt.Println(v1, fi12)
                    if false {
                        if (false || false) {
                            fmt.Println(v1, v1)
                        }
                        _ = len(v2)
                    }
                    _ = v13
                } else {
                    v14 := GenS1{F0: 486, F1: "a", F2: 0}
                    v1 = 42
                    v15 := v3[(("\n\t" + "hello") + "hello")]
                    _ = v14
                    _ = v15
                }
            case 322:
                if (0 <= 201) {
                    var v16 float64 = (-2.25 * (1.5 * 1.5))
                    v16 = float64(-2.25)
                    var v17 bool = (false || true)
                    _ = v16
                    _ = v17
                }
            case -9223372036854775808:
                if (4294967295 > 1) {
                    v1 = (^(v1))
                } else {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        v19 := GenS3{F0: 648}
                        _ = fi18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 14; fi20++ {
                        switch "" {
                        case "\n\t":
                            _ = len(v3)
                        case "a":
                            if false {
                                switch 372 {
                                case 4294967295:
                                    v21 := GenS2{F0: 32767, F1: 50, F2: 538}
                                    _ = v21
                                case 889:
                                    _ = len(v4)
                                default:
                                    var v22 uint16 = 757
                                    _ = v22
                                }
                                v23 := GenS4{F0: 42, F1: 130, F2: 0}
                                v24 := GenS5{F0: (^(v1))}
                                _ = v23
                                _ = v24
                            }
                        case "foo bar":
                            var v25 bool = false
                            _ = v25
                        default:
                            var v26 float32 = 0.0
                            _ = v26
                        }
                        var v27 int = (fi12 ^ (fi20 * (fi12 / ((fi12) | 1))))
                        v28 := GenS5{F0: v1}
                        _ = fi20
                        _ = v27
                        _ = v28
                    }
                    v29 := GenS6{F0: false, F1: 18446744073709551615, F2: 42}
                    _ = v29
                }
            }
            if (42 <= v1) {
                var v30 int16 = 785
                var v31 uint8 = 127
                _ = v30
                _ = v31
            }
            var v32 int8 = 0
            _ = fi12
            _ = v32
        }
    case "foo bar":
        switch "foo bar" {
        case "":
            if (-636 < 1000) {
                v33 := GenS4{F0: 9223372036854775807, F1: 2, F2: -124}
                v1 -= (^(v1))
                _ = v33
            } else {
                v34 := v4["a"]
                var v35 float32 = 1e10
                switch -9223372036854775808 {
                case 42:
                    v1 = v1
                }
                _ = v34
                _ = v35
            }
        case "a":
            if (!false) {
                var v36 uint8 = v1
                switch "hello" {
                case "":
                    var v37 int8 = 100
                    _ = v37
                case "hello":
                    if (!true) {
                        for fi38 := 0; fi38 < 6; fi38++ {
                            var v39 uint8 = (^(v1))
                            v1 = v36
                            if false {
                                fmt.Println(fi38, fi38, v1)
                                v39 = 127
                            } else {
                                v4["hello"] = 0
                                _ = len(v2)
                                var v40 uint64 = 1
                                _ = v40
                            }
                            _ = fi38
                            _ = v39
                        }
                    }
                case "foo bar":
                    switch 2147483647 {
                    case 127:
                        v2 = append(v2, 908)
                    default:
                        switch "" {
                        case "hello":
                            for fi41 := 0; fi41 < 3; fi41++ {
                                switch fi41 {
                                case 32768:
                                    if true {
                                        if (0 == 1) {
                                            for fi42 := 0; fi42 < 1; fi42++ {
                                                v1 = 127
                                                var v43 float64 = 0.1
                                                fmt.Println(fi41)
                                                _ = fi42
                                                _ = v43
                                            }
                                        } else {
                                            switch fi41 {
                                            case 128:
                                                switch "foo bar" {
                                                case "hello":
                                                    switch "" {
                                                    case "a":
                                                        var v44 float32 = (((0.1 - 3.14159) - 0.1) * 3.14159)
                                                        _ = v44
                                                    case "foo bar":
                                                        var v45 uint8 = (^(v36))
                                                        _ = v45
                                                    case "":
                                                        v46 := GenS5{F0: (v1 - 0)}
                                                        _ = v46
                                                    }
                                                case "a":
                                                    v47 := GenS6{F0: false, F1: 804, F2: 1}
                                                    _ = v47
                                                case "foo bar":
                                                    fmt.Println(fi41, fi41, fi41)
                                                default:
                                                    fmt.Println(v36, v36)
                                                }
                                            case 0:
                                                _ = len(v4)
                                            case 774:
                                                var v48 float64 = -2.25
                                                _ = v48
                                            default:
                                                if len(v2) > 0 {
                                                    v2[0] = 329
                                                }
                                            }
                                        }
                                        var v49 int16 = 1
                                        for fi50 := 0; fi50 < 18; fi50++ {
                                            v49 = v49
                                            _ = fi50
                                        }
                                        _ = v49
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 2147483647
                                        }
                                        var v51 float32 = 1.5
                                        _ = v51
                                    }
                                case 9223372036854775807:
                                    if len(v2) > 0 {
                                        v2[0] = 2147483647
                                    }
                                default:
                                    if true {
                                        v4[("\n\t" + "a")] = 1
                                        var v52 int8 = -102
                                        _ = v52
                                    }
                                }
                                for fi53 := 0; fi53 < 10; fi53++ {
                                    var v54 int32 = 0
                                    v55 := GenS1{F0: 2147483648, F1: "hello", F2: v54}
                                    _ = fi53
                                    _ = v54
                                    _ = v55
                                }
                                _ = len(v3)
                                _ = fi41
                            }
                        case "":
                            for fi56 := 0; fi56 < 6; fi56++ {
                                switch "hello" {
                                case "":
                                    v1 += v36
                                }
                                _ = fi56
                            }
                        case "foo bar":
                            var v57 float32 = (0.0 * -2.25)
                            _ = v57
                        default:
                            if (!true) {
                                v36 = v36
                                switch "hello" {
                                case "a":
                                    fmt.Println(v36)
                                default:
                                    v1 = v36
                                }
                            }
                        }
                    }
                }
                _ = v36
            }
        default:
            switch "\n\t" {
            case "hello":
                if true {
                    var v58 uint64 = 4294967295
                    v58 = 18446744073709551615
                    fmt.Println(v58)
                    _ = v58
                } else {
                    v1 = 217
                    var v59 int16 = -158
                    _ = v59
                }
            case "foo bar":
                v1 = 0
            case "\n\t":
                var v60 uint8 = (v1 >> ((714) & 63))
                _ = v60
            }
        }
    case "hello":
        if ((("\n\t" + ("" + "foo bar")) + "hello") == "") {
            if false {
                var v61 bool = ((!(-610 == -2147483648)) || false)
                _ = v61
            } else {
                var v62 int8 = int8(v1)
                for fi63 := 0; fi63 < 1; fi63++ {
                    var v64 float64 = ((0.0 / (((0.0 / 1.5) / 1.5) - (3.14159 * float64(float32(0.1))))) - ((3.14159 + ((-2.25 - (0.1 * float64(1e10))) / (((-2.25 + float64(float32(0.0))) / (-2.25 + -2.25)) - 0.1))) / float64(-2.25)))
                    v2 = append(v2, 1)
                    _ = fi63
                    _ = v64
                }
                v62 += -128
                _ = v62
            }
            for fi65 := 0; fi65 < 8; fi65++ {
                delete(v3, ("a" + ("hello" + "hello")))
                v1 -= v1
                _ = fi65
            }
            fmt.Println(v1)
        }
    default:
        for fi66 := 0; fi66 < 16; fi66++ {
            fmt.Println(v1, fi66)
            switch fi66 {
            case 32767:
                for fi67 := 0; fi67 < 12; fi67++ {
                    v1 += 128
                    _ = fi67
                }
            case -9223372036854775808:
                var v68 uint32 = 199
                _ = v68
            }
            v1 = v1
            _ = fi66
        }
    }
    var v69 bool = true
    var v70 float32 = 3.14159
    fmt.Println(v1, v70)
    for fi71 := 0; fi71 < 16; fi71++ {
        fmt.Println(v1)
        var v72 string = "foo bar"
        _ = fi71
        _ = v72
    }
    fmt.Println(v1, v69, v69)
    var v73 uint8 = v1
    v70 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v7
    _ = v70
    _ = v73
}
