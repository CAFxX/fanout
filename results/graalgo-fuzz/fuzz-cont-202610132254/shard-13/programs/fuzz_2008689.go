package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 bool; F2 int32 }
type GenS5 struct { F0 string; F1 int32; F2 uint }
type GenS6 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    var v1 uint32 = 65535
    v2 := []uint32{4294967295, v1, (v1 >> ((747) & 63)), v1}
    v3 := []int{-9223372036854775808}
    v4 := []uint32{}
    var v5 int16 = -367
    var v6 float32 = 0.0
    if true {
        var v7 int16 = v5
        _ = v7
    }
    v1 -= (v1 >> ((227) & 63))
    if true {
        v1 = 584
        switch 128 {
        case -880:
            switch 127 {
            case 9223372036854775807:
                _ = len(v2)
            case 32768:
                for fi8 := 0; fi8 < 4; fi8++ {
                    var v9 string = ("hello" + ("a" + (("" + "\n\t") + ("\n\t" + "foo bar"))))
                    fmt.Println(v9, v1, v9)
                    _ = fi8
                    _ = v9
                }
            }
        default:
            var v10 float64 = 1e10
            _ = v10
        }
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    fmt.Println(v1)
    switch 9223372036854775807 {
    case 42:
        var v11 int8 = -1
        _ = v11
    case 127:
        v6 -= float32((-2.25 + (1e10 * 3.14159)))
    case 2147483647:
        if (100 != 100) {
            for fi12 := 0; fi12 < 12; fi12++ {
                if false {
                    fmt.Println(fi12, v6)
                } else {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v14 := GenS1{F0: v5, F1: (v5 / (((v5 - 701)) | 1)), F2: 2147483647}
                        if (("\n\t" + ("a" + "hello")) >= "") {
                            if (true || ((^(fi12)) <= (fi13 | fi13))) {
                                v6 += v6
                                fmt.Println(fi13, v1)
                            }
                        }
                        v6 = ((3.14159 * float32(1.5)) + v6)
                        _ = fi13
                        _ = v14
                    }
                }
                var v15 uint16 = 543
                var v16 bool = true
                _ = fi12
                _ = v15
                _ = v16
            }
        }
    }
    v4 = append(v4, (^(v1)))
    v17 := GenS2{F0: "a", F1: 18446744073709551615}
    if false {
        if (163 != 0) {
            var v18 int16 = (v5 * (v5 % ((v5) | 1)))
            _ = v18
        } else {
            if (!(true || ((-2.25 - 1.5) <= 1e10))) {
                var v19 int8 = 1
                _ = v19
            }
            if false {
                v5 -= (v5 % (((v5 * 1)) | 1))
                switch 128 {
                case -9223372036854775808:
                    var v20 uint32 = 2
                    _ = v20
                case 2147483647:
                    if false {
                        if false {
                            fmt.Println(v1)
                            for fi21 := 0; fi21 < 8; fi21++ {
                                for fi22 := 0; fi22 < 3; fi22++ {
                                    switch (-(fi21)) {
                                    case 291:
                                        v23 := GenS2{F0: (((("" + "foo bar") + ("a" + ((("\n\t" + "hello") + ("a" + "a")) + "a"))) + "foo bar") + "\n\t"), F1: 42}
                                        _ = v23
                                    case -2147483648:
                                        if (("" >= ("" + "foo bar")) && (!false)) {
                                            switch (-(fi22)) {
                                            case -120:
                                                v1 = v1
                                            case 0:
                                                switch (-(fi21)) {
                                                case 0:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case -436:
                                                    switch (-(fi22)) {
                                                    case 9223372036854775807:
                                                        var v24 int8 = 1
                                                        _ = v24
                                                    case 1:
                                                        v5 = 177
                                                    default:
                                                        v25 := GenS1{F0: int16(v1), F1: v5, F2: fi21}
                                                        _ = v25
                                                    }
                                                }
                                            case -1:
                                                v1 += (v1 % ((v1) | 1))
                                                switch ("\n\t" + "") {
                                                case "hello":
                                                    var v26 uint16 = 1
                                                    _ = v26
                                                case "\n\t":
                                                    var v27 uint16 = 65535
                                                    _ = v27
                                                case "a":
                                                    var v28 uint8 = 143
                                                    _ = v28
                                                }
                                            default:
                                                if (!true) {
                                                    v29 := GenS1{F0: -1, F1: v5, F2: (fi21 << ((891) & 63))}
                                                    fmt.Println(v1, fi21, v1)
                                                    var v30 int8 = 1
                                                    _ = v29
                                                    _ = v30
                                                }
                                            }
                                            v1 -= 1
                                            var v31 uint = 860
                                            _ = v31
                                        } else {
                                            switch ("hello" + "") {
                                            case "foo bar":
                                                switch (("a" + "hello") + "a") {
                                                case "hello":
                                                    var v32 float64 = ((float64(float32(3.14159)) / 1e10) - 1.5)
                                                    _ = v32
                                                case "\n\t":
                                                    v5 = (^(v5))
                                                case "foo bar":
                                                    v2 = append(v2, (v1 ^ 0))
                                                default:
                                                    fmt.Println(fi22)
                                                }
                                            }
                                        }
                                    }
                                    _ = fi22
                                }
                                switch (fi21 - fi21) {
                                case 9223372036854775807:
                                    var v33 uint = 2
                                    _ = v33
                                case 128:
                                    if false {
                                        v6 += v6
                                        v5 = v5
                                        var v34 uint = 65535
                                        _ = v34
                                    }
                                case -370:
                                    if (-667 >= -9223372036854775808) {
                                        switch (("" + "a") + ("foo bar" + "hello")) {
                                        case "":
                                            var v35 string = "hello"
                                            _ = v35
                                        }
                                        for fi36 := 0; fi36 < 2; fi36++ {
                                            fmt.Println(fi36, v5)
                                            _ = fi36
                                        }
                                    } else {
                                        var v37 int16 = v5
                                        v37 = v5
                                        _ = v37
                                    }
                                default:
                                    v38 := GenS2{F0: "", F1: 592}
                                    _ = v38
                                }
                                _ = fi21
                            }
                        } else {
                            var v39 float32 = v6
                            v40 := GenS3{F0: 78}
                            _ = v39
                            _ = v40
                        }
                        fmt.Println(v1, v5, v1)
                    }
                default:
                    for fi41 := 0; fi41 < 11; fi41++ {
                        for fi42 := 0; fi42 < 8; fi42++ {
                            if (42 <= 68) {
                                var v43 string = "hello"
                                v44 := GenS4{F0: 42, F1: ("hello" >= v43), F2: 0}
                                _ = v43
                                _ = v44
                            } else {
                                v45 := GenS2{F0: "hello", F1: 0}
                                v1 = 2147483647
                                v46 := GenS3{F0: 0}
                                var v47 int16 = (v5 | -231)
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                            _ = len(v4)
                            var v48 bool = (237 <= 31)
                            _ = fi42
                            _ = v48
                        }
                        fmt.Println(fi41)
                        var v49 uint8 = 255
                        _ = fi41
                        _ = v49
                    }
                }
            } else {
                v5 -= v5
            }
        }
        if ((("hello" + "foo bar") >= "foo bar") || false) {
            for fi50 := 0; fi50 < 17; fi50++ {
                v51 := GenS5{F0: "hello", F1: 42, F2: 280}
                _ = fi50
                _ = v51
            }
            fmt.Println(v1, v5)
            v5 -= -1
        } else {
            for fi52 := 0; fi52 < 6; fi52++ {
                switch fi52 {
                case 230:
                    for fi53 := 0; fi53 < 2; fi53++ {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi53, v6, v6)
                        case "foo bar":
                            if (!false) {
                                fmt.Println(v5, v1, fi53)
                            }
                        default:
                            var v54 int = 0
                            _ = v54
                        }
                        switch fi53 {
                        case -9223372036854775808:
                            v2 = append(v2, (v1 - (v1 % ((v1) | 1))))
                        case 942:
                            if (592 != 2) {
                                if false {
                                    var v55 int = (fi53 / (((fi53 * 1)) | 1))
                                    fmt.Println(fi53, v6)
                                    switch 798 {
                                    case 4294967295:
                                        v2 = append(v2, 4294967295)
                                        v55 -= -1
                                        var v56 float64 = 0.1
                                        _ = v56
                                    default:
                                        v5 = v5
                                    }
                                    _ = v55
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v57 string = ("hello" + "\n\t")
                                    _ = v57
                                }
                                v6 -= (-2.25 / ((v6 * ((((-2.25 - v6) / v6) + (1e10 / (v6 * float32(1.5)))) * v6)) / 1.0))
                            }
                        default:
                            switch fi53 {
                            case 1:
                                var v58 float64 = (0.1 * 0.1)
                                _ = v58
                            case -1:
                                var v59 uint = 65535
                                _ = v59
                            }
                        }
                        fmt.Println(v6, v1)
                        v60 := GenS2{F0: ("a" + "\n\t"), F1: 536}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi53
                        _ = v60
                    }
                default:
                    for fi61 := 0; fi61 < 18; fi61++ {
                        v6 += v6
                        v6 = float32(1e10)
                        v5 = v5
                        v6 = (float32(-2.25) * (3.14159 * -2.25))
                        _ = fi61
                    }
                }
                _ = fi52
            }
            for fi62 := 0; fi62 < 18; fi62++ {
                if ((fi62 | -207) <= (fi62 << ((877) & 63))) {
                    var v63 uint = 255
                    fmt.Println(v63)
                    _ = v63
                }
                _ = fi62
            }
            v6 = 0.0
        }
    } else {
        if (!false) {
            switch "" {
            case "\n\t":
                v2 = append(v2, v1)
            case "":
                v64 := GenS6{F0: 0, F1: 42, F2: (100 < -112)}
                _ = v64
            case "a":
                for fi65 := 0; fi65 < 11; fi65++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v66 int64 = 9223372036854775807
                    _ = fi65
                    _ = v66
                }
            }
            var v67 int = -9223372036854775808
            _ = v67
        } else {
            var v68 int32 = -1
            switch 0 {
            case 4294967295:
                var v69 bool = false
                _ = v69
            }
            _ = v68
        }
        for fi70 := 0; fi70 < 10; fi70++ {
            if (!(-2147483648 <= 42)) {
                var v71 int16 = v5
                _ = v71
            }
            _ = len(v3)
            if (9223372036854775807 > 875) {
                var v72 uint64 = 548
                fmt.Println(v1, v72, v5)
                _ = v72
            }
            _ = fi70
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
