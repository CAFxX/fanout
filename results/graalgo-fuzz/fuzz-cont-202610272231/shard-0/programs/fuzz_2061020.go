package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int; F1 uint8; F2 string }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := []int{739, 1, 128, 128}
    var v2 int = 9223372036854775807
    v3 := make(map[int]int)
    v4 := []uint32{}
    v5 := make(map[int]bool)
    _ = len(v3)
    if false {
        for fi6 := 0; fi6 < 6; fi6++ {
            v2 += fi6
            switch v2 {
            case -1:
                v2 -= 128
            case 997:
                switch v2 {
                case 529:
                    fmt.Println(v2, v2)
                case -632:
                    v7 := GenS1{F0: 2147483647, F1: 42}
                    _ = v7
                case 9223372036854775807:
                    for fi8 := 0; fi8 < 2; fi8++ {
                        v9 := GenS1{F0: 42, F1: 466}
                        v2 = -9223372036854775808
                        v2 -= fi8
                        _ = fi8
                        _ = v9
                    }
                }
            case 1:
                fmt.Println(v2, fi6)
            }
            _ = fi6
        }
    }
    switch "hello" {
    case "foo bar":
        v2 -= -1
    case "hello":
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(v2)
            _ = fi10
        }
    default:
        v11 := GenS2{F0: 696}
        _ = v11
    }
    if (-2.25 > 0.0) {
        var v12 float32 = (1.5 * 3.14159)
        _ = v12
    }
    switch (v2 >> ((4294967295) & 63)) {
    case 1:
        if (42 <= 0) {
            if false {
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 float64 = 1.5
                    _ = fi13
                    _ = v14
                }
            }
        } else {
            var v15 int32 = 42
            if (1.5 >= 0.1) {
                _ = len(v5)
            } else {
                delete(v3, -1)
                var v16 bool = ((v2 >> ((2) & 63)) <= v2)
                if (!false) {
                    fmt.Println(v16, v16)
                    fmt.Println(v15)
                    v2 -= (v2 & (v2 << ((255) & 63)))
                }
                _ = v16
            }
            _ = v15
        }
    }
    if (-151 > (v2 & (v2 % ((127) | 1)))) {
        switch "" {
        case "hello":
            v2 += (^(v2))
        case "\n\t":
            var v17 int16 = -32768
            _ = v17
        case "foo bar":
            if (3.14159 == (-2.25 - float64(0.1))) {
                if false {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        var v19 int = (v2 << ((18446744073709551615) & 63))
                        _ = fi18
                        _ = v19
                    }
                    if (1 <= 49) {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 string = (("\n\t" + "foo bar") + "\n\t")
                            _ = fi20
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 17; fi22++ {
                            if (false && (!true)) {
                                v2 = fi22
                                var v23 float32 = 3.14159
                                fmt.Println(v23, v23, fi22)
                                _ = v23
                            }
                            v2 = v2
                            _ = fi22
                        }
                        switch v2 {
                        case 1:
                            v2 += (v2 % ((v2) | 1))
                        case 0:
                            v2 = v2
                        case 2147483647:
                            for fi24 := 0; fi24 < 5; fi24++ {
                                if false {
                                    v25 := GenS2{F0: 1}
                                    v2 = 1
                                    for fi26 := 0; fi26 < 8; fi26++ {
                                        v2 = fi24
                                        if (0 < 255) {
                                            v2 -= 127
                                            if true {
                                                if (0 != 42) {
                                                    v2 -= (fi24 | 554)
                                                }
                                                var v27 uint64 = 46
                                                _ = v27
                                            } else {
                                                _ = len(v3)
                                                switch (-(fi26)) {
                                                case 4294967295:
                                                    var v28 uint16 = 891
                                                    _ = v28
                                                case -2147483648:
                                                    v29 := v3[-2147483648]
                                                    _ = v29
                                                default:
                                                    v30 := GenS1{F0: 998, F1: 65535}
                                                    _ = v30
                                                }
                                                var v31 int32 = 1
                                                var v32 int8 = 127
                                                _ = v31
                                                _ = v32
                                            }
                                        } else {
                                            v33 := GenS1{F0: 671, F1: 0}
                                            var v34 int16 = -738
                                            v35 := GenS3{F0: false, F1: 28}
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        var v36 uint32 = 639
                                        _ = fi26
                                        _ = v36
                                    }
                                    _ = v25
                                } else {
                                    var v37 uint8 = 128
                                    v38 := GenS4{F0: -2147483648, F1: (^(v37)), F2: ""}
                                    var v39 float32 = (-2.25 * float32((((float64((0.1 - float32(-2.25))) * float64(1.5)) - ((float64(float32(-2.25)) + ((1e10 / (0.1 * 3.14159)) / 1.0)) + 3.14159)) / (-2.25 + 0.1))))
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                _ = fi24
                            }
                        }
                    }
                    for fi40 := 0; fi40 < 11; fi40++ {
                        if true {
                            switch ("" + ("foo bar" + "\n\t")) {
                            case "foo bar":
                                if (false && (true && (42 == 32767))) {
                                    v41 := GenS4{F0: 0, F1: 88, F2: "foo bar"}
                                    _ = len(v1)
                                    _ = v41
                                }
                            case "":
                                _ = len(v4)
                            case "hello":
                                fmt.Println(v2, v2)
                            }
                        } else {
                            var v42 bool = ((541 <= 2) || (42 < 0))
                            v2 = fi40
                            v2 -= v2
                            _ = v42
                        }
                        switch v2 {
                        case -1:
                            v2 = (^(v2))
                        case 614:
                            for fi43 := 0; fi43 < 1; fi43++ {
                                v44 := GenS5{F0: -32768, F1: 453}
                                var v45 int8 = 1
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                        case 127:
                            switch ("hello" + "foo bar") {
                            case "\n\t":
                                var v46 float64 = -2.25
                                _ = v46
                            default:
                                _ = len(v3)
                            }
                        }
                        var v47 float32 = float32((0.1 - 3.14159))
                        _ = fi40
                        _ = v47
                    }
                }
            }
        default:
            var v48 int8 = 68
            _ = v48
        }
    } else {
        v49 := v5[(v2 ^ (v2 % ((v2) | 1)))]
        fmt.Println(v49)
        v50 := GenS5{F0: 0, F1: 378}
        _ = v49
        _ = v50
    }
    var v51 bool = (!(775 <= 4294967295))
    var v52 uint = 42
    delete(v3, -545)
    v4 = append(v4, 2)
    v2 -= (v2 % ((v2) | 1))
    switch (v2 >> ((v52) & 63)) {
    case 1:
        switch "hello" {
        case "":
            delete(v3, v2)
        default:
            v51 = v51
        }
    case 689:
        _ = len(v3)
    }
    v51 = (1 <= -9223372036854775808)
    v53 := GenS5{F0: 42, F1: 18446744073709551615}
    v51 = v51
    if v51 {
        switch "" {
        case "foo bar":
            v2 += (v2 % ((-342) | 1))
        case "\n\t":
            v51 = v51
        case "":
            v52 = (v52 >> ((2) & 63))
        }
    }
    switch (("foo bar" + "hello") + "foo bar") {
    case "foo bar":
        for fi54 := 0; fi54 < 19; fi54++ {
            if false {
                if false {
                    delete(v5, -1)
                    fmt.Println(v51, v2, v2)
                }
                fmt.Println(fi54, v2, v2)
                v55 := GenS5{F0: 998, F1: v52}
                _ = v55
            } else {
                v56 := GenS2{F0: 42}
                _ = v56
            }
            _ = fi54
        }
    default:
        if v51 {
            for fi57 := 0; fi57 < 16; fi57++ {
                fmt.Println(v51, v2, v51)
                _ = len(v3)
                _ = fi57
            }
            switch "hello" {
            case "":
                var v58 uint32 = 42
                _ = v58
            case "hello":
                for fi59 := 0; fi59 < 16; fi59++ {
                    v60 := GenS3{F0: ((v51 || ((!true) || v51)) || v51), F1: -53}
                    switch "foo bar" {
                    case "a":
                        v2 = (fi59 | (^(v2)))
                    case "foo bar":
                        for fi61 := 0; fi61 < 19; fi61++ {
                            switch 32768 {
                            case -232:
                                v62 := GenS6{F0: (0 > -1), F1: 2}
                                _ = v62
                            default:
                                var v63 float64 = 3.14159
                                _ = v63
                            }
                            switch fi61 {
                            case 32768:
                                switch v2 {
                                case 32768:
                                    var v64 uint8 = 0
                                    _ = v64
                                case 2147483647:
                                    var v65 uint16 = 65535
                                    _ = v65
                                case 511:
                                    var v66 int8 = 45
                                    _ = v66
                                default:
                                    if (!false) {
                                        var v67 string = "\n\t"
                                        v51 = (!true)
                                        var v68 uint32 = 2147483647
                                        if (!v51) {
                                            v68 += v68
                                        } else {
                                            var v69 float64 = 0.1
                                            switch v67 {
                                            case "foo bar":
                                                v52 = (^(v52))
                                                v51 = v51
                                            case "":
                                                v70 := GenS5{F0: 1, F1: v52}
                                                _ = v70
                                            case "hello":
                                                v67 = "a"
                                            }
                                            _ = v69
                                        }
                                        _ = v67
                                        _ = v68
                                    }
                                }
                            case 993:
                                var v71 bool = v51
                                _ = v71
                            }
                            fmt.Println(v52, fi59, fi59)
                            _ = fi61
                        }
                    case "\n\t":
                        v52 -= v52
                    default:
                        switch "hello" {
                        case "\n\t":
                            if (77 == fi59) {
                                fmt.Println(v2, v51, v51)
                            }
                        }
                    }
                    v72 := GenS7{F0: 1}
                    _ = fi59
                    _ = v60
                    _ = v72
                }
            }
        } else {
            v4 = append(v4, 0)
        }
    }
    var v73 int8 = 18
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v73
}
