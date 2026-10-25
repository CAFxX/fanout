package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 int8; F1 int8; F2 uint8 }

func main() {
    var v1 int = 128
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 bool = false
    var v5 int32 = 42
    v6 := make(map[string]bool)
    var v7 bool = (!v4)
    if (!(v4 && v4)) {
        v4 = (2 <= 42)
        if (true || v7) {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 uint16 = 726
                var v10 uint32 = 2147483647
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(v1, v4, v4)
        }
        var v11 int = 2147483647
        _ = v11
    } else {
        if false {
            switch ((("a" + "\n\t") + (("" + "foo bar") + "foo bar")) + ("foo bar" + "a")) {
            case "hello":
                v4 = ((true || true) && (921 != 4294967295))
            default:
                v5 -= v5
            }
            var v12 int32 = v5
            var v13 float64 = (((1e10 / -2.25) - 3.14159) - 0.0)
            _ = v12
            _ = v13
        }
    }
    if ((!false) || (!(-32768 == 42))) {
        for fi14 := 0; fi14 < 9; fi14++ {
            v4 = true
            _ = fi14
        }
        var v15 uint16 = 65535
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        if false {
            for fi17 := 0; fi17 < 11; fi17++ {
                v18 := GenS1{F0: -238, F1: 0}
                if (!((0.0 + float64(-2.25)) <= 3.14159)) {
                    v19 := v3["hello"]
                    _ = v19
                } else {
                    fmt.Println(fi16)
                    if (true || (float64(3.14159) < -2.25)) {
                        switch 391 {
                        case 128:
                            v1 += (fi16 - (v1 & -2147483648))
                            if (65535 < 2147483647) {
                                v20 := GenS1{F0: 524, F1: 832}
                                v1 -= (fi16 | v1)
                                switch (v1 * (fi17 % (((fi17 | (-(v1)))) | 1))) {
                                case 4294967295:
                                    fmt.Println(v4)
                                    var v21 int16 = 1000
                                    switch fi17 {
                                    case -1:
                                        v21 -= v21
                                    default:
                                        v22 := GenS2{F0: 42, F1: 18446744073709551615, F2: true}
                                        if (41 >= 1) {
                                            v23 := GenS3{F0: 9223372036854775807}
                                            v4 = v4
                                            _ = v23
                                        } else {
                                            v24 := GenS3{F0: v1}
                                            v21 += (v21 * (v21 % (((^(v21))) | 1)))
                                            var v25 int64 = 9223372036854775807
                                            _ = v24
                                            _ = v25
                                        }
                                        var v26 bool = (v5 == (v5 / (((v5 & (v5 % ((1) | 1)))) | 1)))
                                        _ = v22
                                        _ = v26
                                    }
                                    _ = v21
                                case 0:
                                    v5 += (v5 % ((v5) | 1))
                                }
                                _ = v20
                            }
                        case -2147483648:
                            switch 127 {
                            case 155:
                                _ = len(v6)
                            case 32768:
                                var v27 int8 = 115
                                _ = v27
                            case 9223372036854775807:
                                v28 := GenS1{F0: -908, F1: -255}
                                if (-32768 > 0) {
                                    v29 := []int{fi17, 4294967295, (v1 ^ fi17)}
                                    v30 := GenS1{F0: -32768, F1: 1}
                                    _ = v29
                                    _ = v30
                                }
                                _ = v28
                            }
                        default:
                            if (152 == 0) {
                                v7 = true
                                v31 := GenS1{F0: 1, F1: 0}
                                v4 = v4
                                if true {
                                    var v32 int32 = (^(v5))
                                    v4 = false
                                    _ = v32
                                }
                                _ = v31
                            } else {
                                var v33 uint16 = 0
                                fmt.Println(fi17, fi17)
                                var v34 uint8 = 161
                                v5 += 0
                                delete(v3, ("\n\t" + "hello"))
                                v35 := GenS3{F0: fi17}
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        }
                        if v7 {
                            switch (fi16 << ((424) & 63)) {
                            case 2147483647:
                                if (784 != 32767) {
                                    var v36 int32 = 1
                                    _ = v36
                                } else {
                                    v1 = fi17
                                }
                            default:
                                if (911 < 0) {
                                    fmt.Println(v4)
                                }
                            }
                            fmt.Println(v7, v1, v7)
                            fmt.Println(v1, v7)
                            var v37 uint64 = 42
                            v38 := GenS4{F0: -429, F1: (v5 | v5)}
                            _ = v37
                            _ = v38
                        }
                    } else {
                        switch -1 {
                        case 9223372036854775807:
                            _ = len(v6)
                            v39 := GenS2{F0: 703, F1: 4294967295, F2: v7}
                            var v40 int = fi17
                            switch (v1 / ((fi17) | 1)) {
                            case 127:
                                var v41 int64 = -9223372036854775808
                                _ = v41
                            default:
                                var v42 uint64 = 1
                                _ = v42
                            }
                            v43 := GenS3{F0: 1}
                            _ = v39
                            _ = v40
                            _ = v43
                        case 4294967295:
                            v44 := GenS5{F0: 1}
                            _ = v44
                        case -692:
                            fmt.Println(v7, fi16, v1)
                        }
                        v45 := GenS2{F0: 65535, F1: 18446744073709551615, F2: false}
                        _ = v45
                    }
                }
                v7 = (42 == 0)
                var v46 int16 = -1
                v46 -= (v46 | (v46 % ((v46) | 1)))
                v4 = ((1e10 <= 0.0) && ((v46 | (v46 * 32767)) != v46))
                _ = fi17
                _ = v18
                _ = v46
            }
            fmt.Println(v7, v5, v4)
        }
        v47 := v2[(("" + (("" + (("\n\t" + "\n\t") + "\n\t")) + "")) + "a")]
        _ = fi16
        _ = v47
    }
    fmt.Println(v5, v4)
    if (1 > -997) {
        fmt.Println(v1, v1, v4)
    } else {
        v48 := GenS6{F0: 94, F1: 0}
        v49 := v6["hello"]
        var v50 int8 = 42
        _ = v48
        _ = v49
        _ = v50
    }
    switch "a" {
    case "foo bar":
        switch (-(v1)) {
        case -9223372036854775808:
            v4 = (!v4)
        case -353:
            var v51 float32 = 3.14159
            _ = v51
        case 32768:
            v5 += (v5 ^ v5)
        default:
            switch "" {
            case "hello":
                switch ("a" + "hello") {
                case "":
                    for fi52 := 0; fi52 < 17; fi52++ {
                        fmt.Println(fi52, v4, v4)
                        _ = fi52
                    }
                case "foo bar":
                    for fi53 := 0; fi53 < 7; fi53++ {
                        v54 := []int{127, fi53, 116, -535}
                        if v4 {
                            if v4 {
                                var v55 float64 = -2.25
                                _ = v55
                            }
                            v4 = (1000 > 0)
                        }
                        switch v1 {
                        case 2147483647:
                            switch fi53 {
                            case -907:
                                for fi56 := 0; fi56 < 7; fi56++ {
                                    if len(v54) > 0 {
                                        v54[0] = (fi53 * (fi56 ^ (fi53 * (fi53 + v1))))
                                    }
                                    var v57 float64 = 3.14159
                                    _ = fi56
                                    _ = v57
                                }
                            case 497:
                                for fi58 := 0; fi58 < 9; fi58++ {
                                    v1 = 32768
                                    _ = fi58
                                }
                            }
                        case 1:
                            for fi59 := 0; fi59 < 18; fi59++ {
                                switch (fi59 / ((42) | 1)) {
                                case -9223372036854775808:
                                    v60 := GenS1{F0: -800, F1: 0}
                                    _ = v60
                                case -165:
                                    v4 = (263 >= 1)
                                case -370:
                                    v4 = v7
                                    v61 := GenS2{F0: 849, F1: 2, F2: true}
                                    switch fi59 {
                                    case 2147483647:
                                        fmt.Println(v5, v7, fi59)
                                    default:
                                        fmt.Println(v4, v4)
                                    }
                                    _ = v61
                                default:
                                    fmt.Println(v4, v1, fi53)
                                }
                                _ = fi59
                            }
                        case 0:
                            v7 = ("foo bar" != "hello")
                        default:
                            for fi62 := 0; fi62 < 15; fi62++ {
                                v2[""] = 2147483647
                                v63 := GenS1{F0: -209, F1: 1}
                                v64 := v2["foo bar"]
                                _ = fi62
                                _ = v63
                                _ = v64
                            }
                        }
                        _ = fi53
                        _ = v54
                    }
                case "hello":
                    var v65 uint = 946
                    _ = v65
                }
            case "":
                v66 := []int{-80, v1, v1}
                _ = v66
            case "foo bar":
                if (2 < 4294967295) {
                    var v67 float64 = 1.5
                    _ = v67
                }
            }
        }
    case "\n\t":
        v68 := GenS3{F0: (v1 & (v1 << ((255) & 63)))}
        _ = v68
    case "":
        v69 := GenS6{F0: -83, F1: 1}
        _ = v69
    }
    v70 := GenS7{F0: v5}
    v4 = false
    if (!(!(((v5 >> ((646) & 63)) >= (v5 - v5)) || true))) {
        v71 := GenS7{F0: (v5 / (((v5 << ((4294967295) & 63))) | 1))}
        if true {
            for fi72 := 0; fi72 < 1; fi72++ {
                v7 = v4
                v73 := GenS8{F0: -1, F1: 32767}
                var v74 int = fi72
                _ = fi72
                _ = v73
                _ = v74
            }
            var v75 bool = v7
            for fi76 := 0; fi76 < 19; fi76++ {
                v77 := GenS1{F0: -32768, F1: -660}
                _ = fi76
                _ = v77
            }
            _ = v75
        } else {
            v7 = v4
        }
        _ = v71
    } else {
        for fi78 := 0; fi78 < 11; fi78++ {
            switch ("" + "a") {
            case "foo bar":
                if v4 {
                    fmt.Println(v1)
                }
            case "a":
                v79 := GenS9{F0: -1, F1: 0, F2: 255}
                _ = v79
            }
            _ = fi78
        }
        v80 := GenS2{F0: 65535, F1: 4294967295, F2: (((9223372036854775807 == -260) && (1 <= 0)) && v7)}
        _ = v80
    }
    if (128 != 199) {
        v81 := []int{(v1 | (v1 % ((v1) | 1))), v1}
        _ = v81
    }
    v7 = (!(370 != 255))
    v4 = ((-122 != 0) || v4)
    for fi82 := 0; fi82 < 13; fi82++ {
        switch "a" {
        case "foo bar":
            v1 += (^(v1))
        default:
            delete(v3, "foo bar")
        }
        v5 = v5
        for fi83 := 0; fi83 < 3; fi83++ {
            v5 = v5
            _ = fi83
        }
        _ = fi82
    }
    delete(v6, "\n\t")
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v70
}
