package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 string; F2 uint64 }
type GenS6 struct { F0 uint; F1 int16; F2 int64 }
type GenS7 struct { F0 uint; F1 uint8 }

func main() {
    v1 := []int{}
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]int)
    var v4 uint16 = 27
    v5 := []string{("" + (("hello" + "\n\t") + "")), (("foo bar" + "foo bar") + "foo bar"), ""}
    var v6 int32 = 0
    v4 += (v4 - v4)
    var v7 int = 0
    switch 42 {
    case 317:
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                v4 -= 255
                v10 := GenS1{F0: ("hello" + ""), F1: 645}
                _ = fi9
                _ = v10
            }
            v2 = (v2 * -1)
            if (929 >= 18446744073709551615) {
                v11 := GenS2{F0: 1, F1: (fi8 & fi8)}
                if (3.14159 <= ((0.0 + ((-2.25 - 0.0) * ((((1.5 - (1.5 - 1e10)) / (((-2.25 + 3.14159) + 0.0) / 1.5)) - 0.1) * float32(1e10)))) + ((float32(-2.25) * (0.1 / (((((1e10 - float32(0.1)) - -2.25) - 0.0) / float32(((((float64((float32((-2.25 + 1.5)) + 1.5)) + 1.5) * (1e10 / 1.5)) / (float64((((1e10 + 1e10) - 0.0) * 0.1)) - (0.0 - (((((1.5 - 1.5) - 0.0) / 3.14159) + 1.5) * -2.25)))) + 1.5))) + 0.0))) / 3.14159))) {
                    v2 += 471
                    for fi12 := 0; fi12 < 9; fi12++ {
                        var v13 bool = false
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 8; fi14++ {
                        if true {
                            v2 = (v2 - 42)
                            var v15 uint8 = 127
                            _ = v15
                        }
                        var v16 int8 = -128
                        _ = fi14
                        _ = v16
                    }
                }
                var v17 int8 = 88
                _ = v11
                _ = v17
            }
            _ = fi8
        }
    }
    v18 := GenS3{F0: int8(v7), F1: 14}
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 float64 = ((3.14159 * (-2.25 * -2.25)) * 1.5)
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := GenS1{F0: "hello", F1: 462}
            v2 -= v2
            _ = len(v3)
            _ = fi22
            _ = v23
        }
        v4 -= v4
        for fi24 := 0; fi24 < 17; fi24++ {
            v2 = v2
            _ = fi24
        }
        _ = fi21
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        var v26 bool = (!false)
        _ = fi25
        _ = v26
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        v27 := GenS3{F0: 127, F1: 42}
        _ = v27
    } else {
        switch "a" {
        case "foo bar":
            for fi28 := 0; fi28 < 1; fi28++ {
                v6 -= v6
                v6 = v6
                v7 -= v7
                _ = fi28
            }
        case "hello":
            v29 := GenS1{F0: "hello", F1: v4}
            _ = v29
        default:
            v7 = (v7 | -1)
        }
        v30 := GenS4{F0: (v7 * v7)}
        _ = v30
    }
    v5 = append(v5, "foo bar")
    for fi31 := 0; fi31 < 19; fi31++ {
        v2 = (v2 >> ((22) & 63))
        _ = fi31
    }
    var v32 int8 = -1
    if true {
        for fi33 := 0; fi33 < 10; fi33++ {
            switch v7 {
            case 32767:
                v34 := GenS1{F0: "a", F1: 255}
                _ = v34
            case 32768:
                v35 := v3[(^(fi33))]
                _ = v35
            default:
                v36 := GenS4{F0: 9223372036854775807}
                _ = v36
            }
            for fi37 := 0; fi37 < 11; fi37++ {
                v4 = (v4 & (v4 * v4))
                _ = fi37
            }
            if false {
                if len(v5) > 0 {
                    v5[0] = ("foo bar" + "hello")
                }
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 uint64 = 4294967295
                    if true {
                        v40 := GenS2{F0: (v32 << ((uint(v4)) & 63)), F1: v7}
                        if false {
                            if (!true) {
                                var v41 int8 = int8(v39)
                                _ = v41
                            }
                            v39 += (v39 & (v39 >> ((255) & 63)))
                        }
                        v1 = append(v1, fi33)
                        _ = v40
                    }
                    v42 := GenS3{F0: int8(v2), F1: 0}
                    _ = fi38
                    _ = v39
                    _ = v42
                }
                if ((42 <= 501) && true) {
                    if false {
                        v43 := GenS5{F0: 2, F1: "\n\t", F2: 1}
                        _ = v43
                    }
                } else {
                    fmt.Println(v6, v2)
                    for fi44 := 0; fi44 < 9; fi44++ {
                        v45 := GenS6{F0: 42, F1: 42, F2: 1}
                        switch fi33 {
                        case 0:
                            switch ("hello" + "hello") {
                            case "a":
                                switch "foo bar" {
                                case "hello":
                                    _ = len(v1)
                                case "":
                                    v2 = (v2 / ((v2) | 1))
                                case "\n\t":
                                    var v46 string = "\n\t"
                                    _ = v46
                                default:
                                    v47 := GenS7{F0: 1, F1: 1}
                                    _ = v47
                                }
                            default:
                                var v48 int64 = v2
                                _ = v48
                            }
                        case -13:
                            v7 = -2147483648
                        case 128:
                            v49 := GenS7{F0: 255, F1: 232}
                            _ = v49
                        }
                        _ = fi44
                        _ = v45
                    }
                    if true {
                        if (-2.25 == float32(1e10)) {
                            v4 -= v4
                            if (true && (v2 <= (v2 / (((v2 >> ((65535) & 63))) | 1)))) {
                                var v50 uint8 = 42
                                _ = v50
                            } else {
                                var v51 float64 = 0.1
                                v6 = -271
                                _ = v51
                            }
                            if true {
                                var v52 float64 = 3.14159
                                v2 = v2
                                if (!((-(v6)) == v6)) {
                                    v32 += (-(v32))
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        var v53 uint = 975
                                        _ = v53
                                    case "hello":
                                        v4 -= 0
                                    case "a":
                                        v5 = append(v5, "hello")
                                    }
                                    var v54 int8 = v32
                                    for fi55 := 0; fi55 < 19; fi55++ {
                                        fmt.Println(v32)
                                        switch ("a" + "") {
                                        case "hello":
                                            fmt.Println(fi55)
                                        case "":
                                            v32 = 127
                                        case "\n\t":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        switch "" {
                                        case "":
                                            v4 = (v4 >> ((2) & 63))
                                        case "a":
                                            if false {
                                                v56 := GenS4{F0: 4294967295}
                                                if (fi55 < fi33) {
                                                    if (0 <= 2147483648) {
                                                        if true {
                                                            fmt.Println(v4, fi33, v7)
                                                        } else {
                                                            var v57 float64 = v52
                                                            _ = v57
                                                        }
                                                        fmt.Println(v54)
                                                        v32 += (^(v32))
                                                    }
                                                    v52 = v52
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                } else {
                                                    _ = len(v1)
                                                    if true {
                                                        v58 := GenS5{F0: 42, F1: "a", F2: 421}
                                                        v4 = v4
                                                        v54 += (^(v54))
                                                        _ = v58
                                                    } else {
                                                        v52 -= float64(-2.25)
                                                        var v59 int8 = int8(v4)
                                                        _ = v59
                                                    }
                                                    fmt.Println(fi55)
                                                }
                                                _ = v56
                                            } else {
                                                fmt.Println(v4, v32, v32)
                                            }
                                        default:
                                            if ("a" <= "\n\t") {
                                                switch fi33 {
                                                case 128:
                                                    v32 = v32
                                                case 1:
                                                    _ = len(v5)
                                                }
                                            } else {
                                                v60 := v3[int(v4)]
                                                _ = v60
                                            }
                                        }
                                        _ = fi55
                                    }
                                    _ = v54
                                }
                                _ = v52
                            } else {
                                switch "a" {
                                case "":
                                    v7 = (v7 * (fi33 << ((18446744073709551615) & 63)))
                                case "foo bar":
                                    var v61 float64 = ((1e10 / 3.14159) / 1.0)
                                    _ = v61
                                case "hello":
                                    for fi62 := 0; fi62 < 7; fi62++ {
                                        v1 = append(v1, (v7 + (v7 << ((65535) & 63))))
                                        var v63 uint64 = 645
                                        _ = fi62
                                        _ = v63
                                    }
                                }
                            }
                        }
                        switch fi33 {
                        case 1:
                            var v64 int32 = (v6 & -1)
                            _ = v64
                        default:
                            v65 := GenS6{F0: 0, F1: 0, F2: int64(v32)}
                            _ = v65
                        }
                        fmt.Println(fi33, v2, v2)
                    }
                }
            }
            _ = fi33
        }
    }
    if false {
        v32 = (v32 >> ((300) & 63))
        v66 := GenS7{F0: 18446744073709551615, F1: 46}
        _ = v66
    } else {
        var v67 bool = true
        _ = v67
    }
    v32 = v32
    v1 = append(v1, (v7 << ((255) & 63)))
    if false {
        v68 := GenS3{F0: -1, F1: 2}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v68
    }
    v4 = (v4 % ((v4) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
