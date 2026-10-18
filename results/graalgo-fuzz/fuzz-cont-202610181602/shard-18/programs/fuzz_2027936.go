package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint64; F1 int; F2 uint8 }
type GenS7 struct { F0 uint; F1 uint64 }

func main() {
    v1 := []string{}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 string = "\n\t"
    var v5 string = (("" + v4) + (v4 + (v4 + v4)))
    var v6 uint16 = 526
    v5 = v5
    switch (v5 + "") {
    case "foo bar":
        var v7 int16 = 32767
        _ = v7
    case "hello":
        var v8 int32 = 0
        _ = v8
    case "\n\t":
        var v9 float32 = (1e10 + (1.5 * (float32(1.5) - (3.14159 - 1.5))))
        _ = v9
    }
    var v10 uint8 = 127
    v5 = v4
    var v11 int = 127
    v11 -= v11
    if true {
        v11 -= v11
    }
    switch "foo bar" {
    case "hello":
        var v12 int64 = 536
        _ = v12
    case "\n\t":
        v11 += (v11 - (v11 % ((v11) | 1)))
    case "a":
        if (1e10 <= 0.1) {
            if true {
                _ = len(v1)
                var v13 int16 = -863
                v11 = 128
                _ = v13
            }
            var v14 int32 = 1
            _ = v14
        } else {
            for fi15 := 0; fi15 < 16; fi15++ {
                fmt.Println(v11)
                switch v4 {
                case "":
                    if (float32(3.14159) != -2.25) {
                        for fi16 := 0; fi16 < 14; fi16++ {
                            v17 := GenS1{F0: 9223372036854775807, F1: int(v6), F2: v11}
                            fmt.Println(v11, v5, v4)
                            var v18 uint16 = (v6 - v6)
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                case "a":
                    v5 = v4
                default:
                    v5 = "hello"
                }
                var v19 int32 = -2147483648
                _ = fi15
                _ = v19
            }
            var v20 string = ""
            v5 = v20
            _ = v20
        }
    default:
        for fi21 := 0; fi21 < 11; fi21++ {
            if (1 >= 2147483647) {
                v6 -= 2
            } else {
                v4 = v4
            }
            v4 = (v5 + v4)
            _ = fi21
        }
    }
    switch v4 {
    case "hello":
        v11 = v11
    case "foo bar":
        switch v5 {
        case "foo bar":
            var v22 float32 = (0.0 - (0.1 * float32(0.1)))
            _ = v22
        default:
            var v23 float32 = (float32((((0.0 * float64((float32(((((0.1 * 3.14159) + (-2.25 * (3.14159 - (0.0 - 3.14159)))) - -2.25) - ((3.14159 * float64(0.1)) / 1e10))) * 0.1))) - float64((float32(3.14159) + ((((float32((-2.25 / 0.1)) / (((3.14159 * 1e10) / 0.1) * 1e10)) + 3.14159) / (3.14159 / float32(-2.25))) / (3.14159 * (-2.25 + float32(0.0))))))) * float64(1e10))) - 1.5)
            _ = v23
        }
    case "a":
        fmt.Println(v11, v10)
    }
    if ((v6 * v6) < 939) {
        var v24 string = "hello"
        _ = v24
    } else {
        if true {
            _ = len(v1)
            if (!((-2.25 / ((0.0 + 1.5) + 0.0)) <= 3.14159)) {
                var v25 uint32 = 921
                v26 := GenS2{F0: 630}
                _ = v25
                _ = v26
            } else {
                v27 := GenS3{F0: -763}
                _ = v27
            }
            if false {
                switch (-(v11)) {
                case -1:
                    var v28 int32 = 2147483647
                    _ = v28
                case 4294967295:
                    v29 := GenS4{F0: 914, F1: 144, F2: -648}
                    _ = v29
                case -9223372036854775808:
                    v11 += 127
                }
                v30 := GenS5{F0: 42, F1: 0, F2: v10}
                _ = v30
            } else {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v10 = v10
                    _ = fi31
                }
                var v32 uint16 = (v6 >> ((1) & 63))
                switch v11 {
                case 1:
                    switch v11 {
                    case 0:
                        fmt.Println(v32, v10)
                    case 42:
                        for fi33 := 0; fi33 < 2; fi33++ {
                            switch 826 {
                            case 9223372036854775807:
                                for fi34 := 0; fi34 < 11; fi34++ {
                                    switch (-(v11)) {
                                    case 0:
                                        switch fi33 {
                                        case 42:
                                            if (!false) {
                                                _ = len(v1)
                                            }
                                        default:
                                            if false {
                                                var v35 bool = (v10 <= 127)
                                                _ = v35
                                            } else {
                                                var v36 int32 = -1
                                                _ = v36
                                            }
                                        }
                                    case 4294967295:
                                        v6 = 1
                                    case 221:
                                        v10 = 121
                                    default:
                                        if ((v10 <= (v10 >> ((42) & 63))) && true) {
                                            if (!(-2.25 == -2.25)) {
                                                v4 = "foo bar"
                                            }
                                            if false {
                                                if len(v1) > 0 {
                                                    v1[0] = ("hello" + "foo bar")
                                                }
                                            }
                                            var v37 int = fi34
                                            _ = v37
                                        } else {
                                            var v38 int64 = 191
                                            if (!false) {
                                                if (-147 >= 2147483647) {
                                                    switch "a" {
                                                    case "hello":
                                                        delete(v2, int(v38))
                                                    case "foo bar":
                                                        if (((3.14159 * (1e10 - 0.1)) / 1.0) < 0.0) {
                                                            var v39 int32 = 136
                                                            fmt.Println(v5, v39, v6)
                                                            _ = v39
                                                        }
                                                    case "":
                                                        if (-77 >= -128) {
                                                            var v40 string = ((v5 + "foo bar") + ("a" + v4))
                                                            var v41 int64 = (v38 << ((2) & 63))
                                                            var v42 int16 = -32768
                                                            _ = v40
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                    }
                                                }
                                                v2[9223372036854775807] = true
                                                v1 = append(v1, "a")
                                            } else {
                                                fmt.Println(v32, v4, v6)
                                                v43 := GenS6{F0: 729, F1: fi33, F2: (v10 + (v10 | (^(v10))))}
                                                v44 := GenS6{F0: 42, F1: fi34, F2: v10}
                                                v4 = "hello"
                                                switch 4294967295 {
                                                case 127:
                                                    switch fi34 {
                                                    case 261:
                                                        v4 = ""
                                                    case -2147483648:
                                                        switch (v5 + ("hello" + ("" + (v4 + v5)))) {
                                                        case "foo bar":
                                                            switch 127 {
                                                            case 838:
                                                                v45 := GenS5{F0: 733, F1: (v38 % ((-9223372036854775808) | 1)), F2: v10}
                                                                fmt.Println(fi34, v32, v38)
                                                                _ = v45
                                                            case 4294967295:
                                                                v46 := GenS6{F0: 1, F1: (fi33 >> ((255) & 63)), F2: v10}
                                                                _ = v46
                                                            }
                                                        default:
                                                            v47 := GenS1{F0: v38, F1: (v11 + v11), F2: -317}
                                                            _ = v47
                                                        }
                                                    default:
                                                        v11 += fi33
                                                    }
                                                case -831:
                                                    var v48 float64 = 1e10
                                                    _ = v48
                                                case 42:
                                                    v5 = "foo bar"
                                                }
                                                var v49 uint8 = (v10 / ((v10) | 1))
                                                var v50 uint64 = 42
                                                _ = v43
                                                _ = v44
                                                _ = v49
                                                _ = v50
                                            }
                                            v1 = append(v1, "a")
                                            _ = v38
                                        }
                                    }
                                    _ = fi34
                                }
                            case -9223372036854775808:
                                var v51 uint32 = 64
                                _ = v51
                            default:
                                v6 = (^(v32))
                            }
                            _ = fi33
                        }
                    default:
                        v52 := GenS7{F0: 880, F1: 0}
                        _ = v52
                    }
                case 32767:
                    fmt.Println(v11, v5, v11)
                case 9223372036854775807:
                    fmt.Println(v10, v4)
                default:
                    if (false || (v32 != (v6 % (((v6 * (v32 & (v6 ^ v32)))) | 1)))) {
                        v4 = v4
                        for fi53 := 0; fi53 < 7; fi53++ {
                            v32 = v6
                            var v54 string = (v4 + "hello")
                            v55 := v2[(fi53 << ((255) & 63))]
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                    }
                }
                _ = v32
            }
        }
    }
    var v56 uint = 0
    var v57 float64 = (1e10 / 1.0)
    v11 = (^(v11))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
