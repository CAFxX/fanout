package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16; F1 string; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int8; F1 int64; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint; F1 uint; F2 int8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint64; F1 int64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint8; F1 int16; F2 bool }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = (v1 << ((42) & 63))
    var v3 uint32 = 65535
    v4 := []int{9223372036854775807, -1, 0}
    v3 = (v3 + v3)
    if false {
        if false {
            fmt.Println(v1)
        }
    } else {
        v3 = (v3 ^ 65535)
        for fi5 := 0; fi5 < 20; fi5++ {
            for fi6 := 0; fi6 < 1; fi6++ {
                if (false && (!false)) {
                    var v7 uint64 = 663
                    var v8 uint8 = 42
                    _ = v7
                    _ = v8
                } else {
                    switch 128 {
                    case 32768:
                        switch "" {
                        case "hello":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "":
                            var v9 int8 = 67
                            _ = v9
                        case "\n\t":
                            var v10 int32 = 42
                            _ = v10
                        default:
                            if len(v4) > 0 {
                                v4[0] = fi5
                            }
                        }
                    case 32767:
                        switch (fi6 & (fi6 ^ fi5)) {
                        case 832:
                            var v11 int64 = -9223372036854775808
                            _ = v11
                        case 9223372036854775807:
                            if true {
                                v1 = (v1 & 1)
                            } else {
                                v3 -= 437
                                v12 := GenS1{F0: 526, F1: 9223372036854775807}
                                var v13 uint64 = (v2 - v2)
                                v3 += 65535
                                _ = v12
                                _ = v13
                            }
                        case -591:
                            if len(v4) > 0 {
                                v4[0] = 2147483647
                            }
                        }
                    case 4294967295:
                        var v14 uint64 = 848
                        _ = v14
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    switch (-(fi6)) {
                    case -2147483648:
                        fmt.Println(v2, v2, v2)
                    }
                    v3 += v3
                }
                switch (^(fi6)) {
                case 0:
                    switch "\n\t" {
                    case "hello":
                        v15 := GenS2{F0: 668, F1: "", F2: (^(v1))}
                        _ = v15
                    case "\n\t":
                        var v16 uint32 = v3
                        if (!true) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if (!(false || (255 < 42))) {
                                if (fi5 != fi6) {
                                    if false {
                                        if (221 != fi6) {
                                            v17 := GenS3{F0: ("a" + ("\n\t" + "hello"))}
                                            fmt.Println(v2, fi5)
                                            v18 := GenS1{F0: 1, F1: fi5}
                                            _ = v17
                                            _ = v18
                                        } else {
                                            v4 = append(v4, fi5)
                                            v4 = append(v4, fi6)
                                        }
                                    }
                                    v19 := GenS3{F0: (("foo bar" + "a") + (((("a" + "hello") + "foo bar") + ("foo bar" + "foo bar")) + "hello"))}
                                    _ = v19
                                }
                                var v20 float32 = (3.14159 / ((float32(1.5) / 0.1) - 0.0))
                                var v21 string = "foo bar"
                                _ = v20
                                _ = v21
                            } else {
                                v2 = v1
                                switch 42 {
                                case 30:
                                    v22 := GenS4{F0: (v2 << ((1) & 63)), F1: (^(v1)), F2: (v16 & 1)}
                                    _ = v22
                                case -650:
                                    var v23 int32 = -819
                                    _ = v23
                                default:
                                    v2 -= uint64(v16)
                                }
                                v1 += v1
                            }
                            var v24 int16 = 0
                            _ = v24
                        }
                        fmt.Println(fi5)
                        _ = v16
                    case "a":
                        fmt.Println(v2)
                    default:
                        if true {
                            v3 = v3
                            var v25 float32 = ((((((float32(3.14159) / 1e10) * 1e10) / 1.5) * (0.0 / 0.1)) - ((1e10 / -2.25) / float32(1e10))) - ((float32(1.5) - (((1e10 * -2.25) + ((1e10 * -2.25) / 1.5)) - (3.14159 * -2.25))) - 1e10))
                            _ = v25
                        } else {
                            if ((^(fi6)) > fi6) {
                                fmt.Println(fi5, fi5, fi6)
                                fmt.Println(v1, fi6)
                                if true {
                                    v1 -= 1
                                    var v26 int16 = 1000
                                    var v27 uint32 = v3
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 uint16 = 714
                                    _ = v28
                                }
                            }
                            fmt.Println(v3)
                        }
                    }
                case 42:
                    switch "" {
                    case "\n\t":
                        switch "foo bar" {
                        case "foo bar":
                            v29 := GenS3{F0: "a"}
                            _ = v29
                        case "":
                            if (v1 < (v2 & (v1 >> ((42) & 63)))) {
                                if (1.5 != ((-2.25 - (1e10 * 1e10)) + 1e10)) {
                                    v3 -= (v3 | (v3 << ((701) & 63)))
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if true {
                                        if false {
                                            var v30 int8 = 127
                                            fmt.Println(v3)
                                            fmt.Println(v30, v3, v3)
                                            _ = v30
                                        } else {
                                            switch "foo bar" {
                                            case "foo bar":
                                                switch (fi5 - (fi5 / ((fi5) | 1))) {
                                                case 128:
                                                    v1 = (v2 >> ((42) & 63))
                                                case 471:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                default:
                                                    v1 += 0
                                                }
                                            case "\n\t":
                                                var v31 int64 = 0
                                                _ = v31
                                            case "a":
                                                _ = len(v4)
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = (fi5 + 127)
                                            }
                                            var v32 string = "a"
                                            _ = v32
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    } else {
                                        var v33 int16 = 32767
                                        if len(v4) > 0 {
                                            v4[0] = fi6
                                        }
                                        _ = v33
                                    }
                                }
                                if (128 > 34) {
                                    if ((1.5 / 1.0) > (1.5 / 0.1)) {
                                        var v34 float32 = (float32((float64(0.0) + 1.5)) * (-2.25 / float32((float64(0.1) / 1.5))))
                                        _ = v34
                                    } else {
                                        if (!("\n\t" >= ("" + "foo bar"))) {
                                            if ("hello" >= ("\n\t" + "")) {
                                                var v35 uint64 = v2
                                                v36 := GenS1{F0: 32767, F1: fi5}
                                                _ = v35
                                                _ = v36
                                            }
                                            switch 128 {
                                            case 128:
                                                var v37 uint = 255
                                                _ = v37
                                            case 32767:
                                                v4 = append(v4, fi5)
                                            case 9223372036854775807:
                                                v1 = v2
                                            default:
                                                if ("\n\t" < "a") {
                                                    var v38 uint32 = v3
                                                    _ = v38
                                                } else {
                                                    fmt.Println(v2, v1, v1)
                                                    v39 := GenS5{F0: 127, F1: 624, F2: v3}
                                                    var v40 uint32 = v3
                                                    _ = v39
                                                    _ = v40
                                                }
                                            }
                                            var v41 uint16 = 165
                                            _ = v41
                                        }
                                        if true {
                                            v3 = (v3 & 2147483648)
                                        } else {
                                            if false {
                                                if false {
                                                    fmt.Println(v2, v3, v2)
                                                    var v42 int = 199
                                                    v42 += v42
                                                    _ = v42
                                                } else {
                                                    v3 = uint32(v1)
                                                    var v43 uint8 = 168
                                                    if true {
                                                        v3 = 2
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    } else {
                                                        if true {
                                                            v4 = append(v4, (fi6 * fi6))
                                                            var v44 bool = true
                                                            fmt.Println(fi5, fi6, v2)
                                                            _ = v44
                                                        }
                                                        var v45 uint64 = 717
                                                        _ = v45
                                                    }
                                                    _ = v43
                                                }
                                                v46 := GenS5{F0: 42, F1: -163, F2: (v3 / ((2147483648) | 1))}
                                                var v47 float64 = (-2.25 + 1e10)
                                                v48 := GenS6{F0: 65535, F1: 175, F2: -1}
                                                var v49 int32 = -1
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            } else {
                                                if false {
                                                    switch fi5 {
                                                    case 127:
                                                        if ((v3 - v3) == (v3 & v3)) {
                                                            switch int(v3) {
                                                            case 4294967295:
                                                                if false {
                                                                    var v50 float64 = 0.0
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        var v51 int16 = 170
                                                                        _ = v51
                                                                    case "foo bar":
                                                                        fmt.Println(v3)
                                                                    }
                                                                    var v52 uint = 255
                                                                    _ = v50
                                                                    _ = v52
                                                                } else {
                                                                    switch fi6 {
                                                                    case -26:
                                                                        var v53 int = 1
                                                                        _ = v53
                                                                    case -390:
                                                                        if (-9223372036854775808 <= -1) {
                                                                            if ((v3 % ((v3) | 1)) < (v3 + (v3 + (^(v3))))) {
                                                                                v3 = (v3 | (v3 | 65535))
                                                                                var v54 int16 = -1
                                                                                v2 = (v2 << ((42) & 63))
                                                                                var v55 uint8 = 1
                                                                                _ = v54
                                                                                _ = v55
                                                                            }
                                                                        }
                                                                    case 2147483647:
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    default:
                                                                        var v56 uint = 2
                                                                        _ = v56
                                                                    }
                                                                    var v57 int8 = -128
                                                                    _ = v57
                                                                }
                                                            }
                                                            v3 = (^(v3))
                                                            v1 += 4294967295
                                                            var v58 uint32 = (^(v3))
                                                            switch fi6 {
                                                            case 0:
                                                                fmt.Println(v58, fi6, fi5)
                                                            case 32767:
                                                                v2 -= v1
                                                                var v59 int64 = 0
                                                                _ = v59
                                                            default:
                                                                fmt.Println(v58, fi5, v2)
                                                            }
                                                            _ = v58
                                                        } else {
                                                            var v60 int64 = 42
                                                            _ = v60
                                                        }
                                                    default:
                                                        var v61 string = "hello"
                                                        _ = v61
                                                    }
                                                } else {
                                                    var v62 int = (fi5 >> ((255) & 63))
                                                    _ = v62
                                                }
                                            }
                                            v4 = append(v4, (fi5 ^ 4294967295))
                                            v63 := GenS7{F0: 86, F1: -1}
                                            _ = v63
                                        }
                                        var v64 int64 = 9223372036854775807
                                        var v65 int64 = (v64 / ((v64) | 1))
                                        _ = v64
                                        _ = v65
                                    }
                                    v66 := GenS8{F0: 42, F1: 425, F2: 100}
                                    if (566 >= 0) {
                                        v2 += 18446744073709551615
                                        if true {
                                            if len(v4) > 0 {
                                                v4[0] = 9223372036854775807
                                            }
                                            fmt.Println(v1, fi5)
                                        }
                                        if false {
                                            v67 := GenS9{F0: 45}
                                            _ = v67
                                        } else {
                                            switch (-(fi5)) {
                                            case 4294967295:
                                                v68 := GenS5{F0: 1, F1: 383, F2: 1}
                                                _ = v68
                                            }
                                            v69 := GenS6{F0: 289, F1: 65535, F2: 1}
                                            var v70 float32 = (float32(0.1) / (1e10 * (1.5 * float32(((0.1 / (3.14159 - (-2.25 + -2.25))) - 0.1)))))
                                            _ = v69
                                            _ = v70
                                        }
                                    } else {
                                        v4 = append(v4, fi6)
                                    }
                                    v3 = v3
                                    var v71 bool = true
                                    _ = v66
                                    _ = v71
                                } else {
                                    fmt.Println(v1)
                                }
                            } else {
                                v1 += v1
                                if (92 == 2) {
                                    v1 = v1
                                    var v72 uint64 = v1
                                    v3 -= v3
                                    _ = v72
                                } else {
                                    var v73 int = fi6
                                    switch ("a" + "\n\t") {
                                    case "hello":
                                        v4 = append(v4, fi5)
                                    }
                                    var v74 uint = 255
                                    _ = v73
                                    _ = v74
                                }
                            }
                        case "hello":
                            switch -661 {
                            case 42:
                                v75 := GenS10{F0: v2, F1: 488}
                                _ = v75
                            }
                        default:
                            v1 = 1
                        }
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            v1 -= (v1 | v2)
                        case "":
                            v76 := GenS3{F0: ("a" + "a")}
                            _ = v76
                        }
                    }
                case 127:
                    var v77 uint = 222
                    _ = v77
                default:
                    var v78 int32 = -974
                    _ = v78
                }
                _ = fi6
            }
            if (!(!true)) {
                v2 += v1
            }
            v79 := GenS8{F0: 0, F1: 937, F2: -1}
            _ = fi5
            _ = v79
        }
    }
    for fi80 := 0; fi80 < 17; fi80++ {
        switch (fi80 % ((fi80) | 1)) {
        case 682:
            var v81 string = "hello"
            _ = v81
        case -462:
            switch "\n\t" {
            case "foo bar":
                v82 := GenS3{F0: (("" + "") + "")}
                _ = v82
            case "hello":
                var v83 uint16 = 1
                _ = v83
            default:
                v84 := GenS5{F0: -51, F1: 315, F2: (v3 + v3)}
                _ = v84
            }
        }
        for fi85 := 0; fi85 < 7; fi85++ {
            v2 = (^(v1))
            _ = fi85
        }
        switch ((("foo bar" + ("a" + ("a" + (("hello" + (("a" + ("hello" + "a")) + "\n\t")) + "foo bar")))) + "") + "\n\t") {
        case "hello":
            _ = len(v4)
        default:
            _ = len(v4)
        }
        _ = fi80
    }
    for fi86 := 0; fi86 < 8; fi86++ {
        var v87 int64 = 9223372036854775807
        _ = fi86
        _ = v87
    }
    v3 -= (v3 + v3)
    switch "hello" {
    case "a":
        v88 := GenS11{F0: -949}
        _ = v88
    }
    if false {
        v3 = 219
    } else {
        for fi89 := 0; fi89 < 16; fi89++ {
            switch "" {
            case "foo bar":
                v4 = append(v4, (fi89 << ((0) & 63)))
            default:
                switch (fi89 ^ (fi89 % (((fi89 / ((fi89) | 1))) | 1))) {
                case 0:
                    if false {
                        v3 -= (^(v3))
                        var v90 uint32 = (v3 >> ((65535) & 63))
                        _ = v90
                    }
                default:
                    if false {
                        v91 := GenS12{F0: 1, F1: 1, F2: (42 == 381)}
                        switch "hello" {
                        case "":
                            v92 := GenS12{F0: 255, F1: 32767, F2: true}
                            _ = v92
                        case "hello":
                            v4 = append(v4, 128)
                        case "\n\t":
                            switch ("a" + ("" + (("foo bar" + "foo bar") + ""))) {
                            case "\n\t":
                                if len(v4) > 0 {
                                    v4[0] = (fi89 - fi89)
                                }
                            case "foo bar":
                                var v93 int8 = -128
                                _ = v93
                            case "":
                                for fi94 := 0; fi94 < 6; fi94++ {
                                    v95 := GenS1{F0: 438, F1: fi94}
                                    fmt.Println(fi94, v1)
                                    _ = fi94
                                    _ = v95
                                }
                            }
                        default:
                            v2 -= v1
                        }
                        _ = v91
                    }
                }
            }
            v96 := GenS1{F0: 42, F1: (-(fi89))}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi89
            _ = v96
        }
        var v97 uint8 = 0
        _ = v97
    }
    var v98 uint = 65535
    var v99 uint64 = 18446744073709551615
    var v100 uint16 = 1
    var v101 float32 = (float32(0.1) - 0.1)
    v3 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v98
    _ = v99
}
