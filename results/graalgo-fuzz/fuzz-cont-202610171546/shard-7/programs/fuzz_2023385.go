package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int32; F1 uint32; F2 uint }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 int32 }

func main() {
    v1 := []string{"foo bar", ("a" + "")}
    var v2 int8 = -128
    v3 := make(map[int]string)
    var v4 float32 = (float32(1e10) / (1e10 / -2.25))
    if true {
        v2 = (^(v2))
        switch "" {
        case "":
            for fi5 := 0; fi5 < 20; fi5++ {
                switch fi5 {
                case -1:
                    fmt.Println(fi5, fi5)
                case 1:
                    v4 = -2.25
                case -827:
                    switch fi5 {
                    case 42:
                        var v6 float64 = 0.1
                        _ = v6
                    }
                }
                _ = fi5
            }
        case "\n\t":
            var v7 float64 = -2.25
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 19; fi8++ {
                v2 -= -128
                var v9 float32 = v4
                _ = fi8
                _ = v9
            }
        default:
            v10 := GenS1{F0: 42}
            _ = v10
        }
        switch 4294967295 {
        case 42:
            v2 -= v2
        }
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        if false {
            for fi12 := 0; fi12 < 20; fi12++ {
                if true {
                    v2 = v2
                }
                _ = fi12
            }
            var v13 int8 = 84
            for fi14 := 0; fi14 < 7; fi14++ {
                if true {
                    if (!(849 == 18446744073709551615)) {
                        v15 := GenS2{F0: 127}
                        if ((-2.25 * 0.0) == (0.1 / 0.1)) {
                            var v16 float32 = (-2.25 * v4)
                            if (112 >= 255) {
                                v4 -= 3.14159
                            } else {
                                v13 -= v13
                                var v17 int64 = -9223372036854775808
                                _ = v17
                            }
                            _ = v16
                        }
                        _ = v15
                    } else {
                        var v18 float64 = (3.14159 * 0.0)
                        v2 += -1
                        var v19 float64 = (1e10 + 1.5)
                        if true {
                            var v20 int16 = -1
                            var v21 int8 = 0
                            var v22 string = "a"
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        v18 -= (float64(v4) / 1.5)
                        _ = v18
                        _ = v19
                    }
                    v23 := GenS1{F0: 9223372036854775807}
                    _ = v23
                }
                if true {
                    var v24 float64 = 3.14159
                    switch "\n\t" {
                    case "":
                        switch fi14 {
                        case 2147483647:
                            var v25 float64 = (1e10 - ((v24 + v24) / v24))
                            _ = v25
                        case -1:
                            if (fi11 != (fi14 << ((4294967295) & 63))) {
                                v24 = 3.14159
                                switch (fi14 ^ 128) {
                                case -4:
                                    switch "\n\t" {
                                    case "":
                                        var v26 float64 = v24
                                        _ = v26
                                    case "a":
                                        v2 = v13
                                    case "hello":
                                        var v27 bool = ((true && false) && (255 <= 2))
                                        switch "\n\t" {
                                        case "\n\t":
                                            if len(v1) > 0 {
                                                v1[0] = ((("a" + "hello") + "hello") + ("" + ("foo bar" + "\n\t")))
                                            }
                                        case "a":
                                            switch "foo bar" {
                                            case "foo bar":
                                                fmt.Println(v24, v13)
                                            }
                                        default:
                                            v4 -= v4
                                        }
                                        if (!true) {
                                            if false {
                                                switch fi11 {
                                                case 268:
                                                    if ("\n\t" != "a") {
                                                        if (v4 == 1e10) {
                                                            v28 := GenS3{F0: 42, F1: "\n\t"}
                                                            _ = v28
                                                        }
                                                        fmt.Println(v4, v24, v24)
                                                    } else {
                                                        if (146 < -9223372036854775808) {
                                                            v2 -= v13
                                                        } else {
                                                            v27 = (42 < v2)
                                                            var v29 int = 32767
                                                            v4 -= 0.1
                                                            _ = v29
                                                        }
                                                        var v30 int64 = 42
                                                        _ = v30
                                                    }
                                                case -127:
                                                    v24 = (v24 * float64(((((float32(v24) * 0.1) + ((v4 + v4) * 0.1)) * ((-2.25 / v4) - v4)) * v4)))
                                                }
                                                v31 := GenS4{F0: 0, F1: 0}
                                                var v32 uint64 = 659
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(v13, fi14)
                                                    if true {
                                                        v33 := GenS5{F0: 2, F1: 42, F2: 253}
                                                        switch ("\n\t" + "a") {
                                                        case "hello":
                                                            fmt.Println(fi11, v27)
                                                        case "foo bar":
                                                            var v34 float64 = (v24 / 0.1)
                                                            _ = v34
                                                        default:
                                                            v24 = v24
                                                        }
                                                        _ = v33
                                                    }
                                                    v35 := GenS3{F0: 1, F1: "hello"}
                                                    _ = v35
                                                case "":
                                                    v27 = false
                                                default:
                                                    var v36 uint64 = v32
                                                    _ = v36
                                                }
                                                _ = v31
                                                _ = v32
                                            }
                                            var v37 int16 = -708
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v37
                                        }
                                        _ = v27
                                    default:
                                        if (("" < "\n\t") && (574 <= 9223372036854775807)) {
                                            fmt.Println(v13)
                                            fmt.Println(v2, v4, fi14)
                                            v2 += 1
                                        }
                                    }
                                case -661:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    _ = len(v1)
                                }
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = ("\n\t" + "")
                                }
                                v13 = -36
                            }
                        default:
                            if ((-335 == 128) && false) {
                                var v38 int = (fi11 / ((fi11) | 1))
                                _ = v38
                            }
                        }
                    default:
                        switch "" {
                        case "":
                            v39 := GenS4{F0: -354, F1: 42}
                            _ = v39
                        case "\n\t":
                            v24 = (v24 / v24)
                        }
                    }
                    delete(v3, (fi14 / (((-(fi11))) | 1)))
                    _ = len(v1)
                    var v40 uint32 = 979
                    _ = v24
                    _ = v40
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        if true {
                            v13 = (^(v2))
                        }
                    default:
                        fmt.Println(fi11, fi14)
                    }
                    fmt.Println(fi14, fi14)
                    var v41 bool = (!false)
                    fmt.Println(fi14, v4)
                    _ = v41
                }
                v42 := GenS4{F0: 1, F1: -1}
                _ = fi14
                _ = v42
            }
            _ = v13
        }
        v43 := GenS1{F0: 1}
        _ = fi11
        _ = v43
    }
    for fi44 := 0; fi44 < 7; fi44++ {
        v2 += v2
        v2 -= (v2 * v2)
        _ = fi44
    }
    v2 = (v2 - v2)
    v4 -= v4
    switch ("\n\t" + "hello") {
    case "":
        if false {
            var v45 int64 = 1
            v4 = 1e10
            _ = v45
        } else {
            if false {
                if ((!(v4 != v4)) || false) {
                    delete(v3, 32767)
                } else {
                    var v46 uint = 255
                    fmt.Println(v4)
                    if false {
                        v2 -= 100
                        if true {
                            v47 := GenS6{F0: 0}
                            _ = v47
                        } else {
                            v3[4294967295] = "hello"
                            v4 -= v4
                            var v48 uint = (v46 * (^(v46)))
                            _ = v48
                        }
                    } else {
                        for fi49 := 0; fi49 < 8; fi49++ {
                            for fi50 := 0; fi50 < 3; fi50++ {
                                var v51 int8 = (v2 * (v2 | v2))
                                _ = fi50
                                _ = v51
                            }
                            _ = fi49
                        }
                        fmt.Println(v2, v4)
                        delete(v3, -657)
                    }
                    _ = v46
                }
            }
            switch (("\n\t" + ((((("a" + "a") + "") + "hello") + ((("\n\t" + "hello") + "a") + ("a" + ("a" + "")))) + ("" + ("hello" + (("" + "") + ""))))) + "hello") {
            case "":
                v52 := GenS3{F0: 1, F1: "hello"}
                _ = v52
            default:
                var v53 float64 = 3.14159
                _ = v53
            }
            v2 += v2
        }
    }
    if true {
        v4 -= (-2.25 + v4)
    }
    var v54 uint16 = 733
    v2 -= (v2 * 127)
    switch 9223372036854775807 {
    case 2147483647:
        v54 = (v54 + (v54 << ((1) & 63)))
    case -42:
        if true {
            if (858 == 64) {
                if false {
                    v4 -= (((v4 - 1e10) * (3.14159 / 1.0)) - v4)
                    for fi55 := 0; fi55 < 17; fi55++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if ((fi55 == (fi55 / ((0) | 1))) && true) {
                            var v56 uint8 = 0
                            switch "\n\t" {
                            case "":
                                switch "" {
                                case "a":
                                    switch "" {
                                    case "foo bar":
                                        var v57 string = "foo bar"
                                        _ = v57
                                    default:
                                        v2 = v2
                                    }
                                case "hello":
                                    fmt.Println(fi55)
                                }
                            default:
                                for fi58 := 0; fi58 < 9; fi58++ {
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(v2)
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        if (fi55 != (fi58 | (^(fi55)))) {
                                            v56 -= (v56 >> ((255) & 63))
                                        } else {
                                            v1 = append(v1, "\n\t")
                                        }
                                        var v59 bool = (!false)
                                        _ = v59
                                    default:
                                        v2 = 72
                                        fmt.Println(fi58)
                                        switch "" {
                                        case "":
                                            v60 := GenS7{F0: -2147483648, F1: 2147483647, F2: 963}
                                            _ = v60
                                        case "a":
                                            if ((v4 * 0.1) == ((v4 * v4) - (3.14159 * ((1e10 - 0.1) / 1.0)))) {
                                                _ = len(v3)
                                                if (true && (2 < 1)) {
                                                    var v61 int64 = 1
                                                    v56 = (^(v56))
                                                    _ = v61
                                                }
                                                v54 = 137
                                            }
                                        }
                                    }
                                    _ = fi58
                                }
                            }
                            _ = v56
                        }
                        fmt.Println(fi55, v4)
                        _ = fi55
                    }
                }
                var v62 int16 = -369
                v63 := GenS3{F0: 1, F1: ""}
                _ = v62
                _ = v63
            }
            v64 := GenS4{F0: 898, F1: 0}
            _ = v64
        } else {
            if false {
                var v65 int8 = (v2 * v2)
                _ = v65
            } else {
                v66 := GenS1{F0: 9223372036854775807}
                _ = v66
            }
            v2 = -107
            if false {
                v54 -= (v54 / ((v54) | 1))
                for fi67 := 0; fi67 < 10; fi67++ {
                    v68 := GenS1{F0: -1}
                    switch fi67 {
                    case 32768:
                        v4 = v4
                    }
                    for fi69 := 0; fi69 < 5; fi69++ {
                        if true {
                            if true {
                                v70 := GenS1{F0: 113}
                                _ = v70
                            }
                        }
                        var v71 uint8 = 128
                        _ = fi69
                        _ = v71
                    }
                    _ = fi67
                    _ = v68
                }
            }
        }
    }
    switch -23 {
    case 128:
        switch -1 {
        case 32768:
            v54 -= 65535
        case 0:
            if true {
                switch "" {
                case "":
                    v72 := GenS8{F0: 1, F1: 4294967295, F2: 4294967295}
                    _ = v72
                default:
                    switch "a" {
                    case "":
                        v73 := v3[9223372036854775807]
                        _ = v73
                    case "\n\t":
                        var v74 uint8 = 152
                        _ = v74
                    }
                }
                var v75 int32 = -1
                _ = v75
            } else {
                v3[-9223372036854775808] = ("" + "a")
                var v76 float64 = float64(v4)
                _ = v76
            }
        default:
            for fi77 := 0; fi77 < 19; fi77++ {
                if (true || (-1 == -801)) {
                    for fi78 := 0; fi78 < 1; fi78++ {
                        v2 += (-(v2))
                        v79 := GenS5{F0: (v54 / ((v54) | 1)), F1: 0, F2: 1}
                        if false {
                            v54 -= v54
                            var v80 uint32 = 908
                            var v81 uint16 = 32767
                            _ = v80
                            _ = v81
                        }
                        if false {
                            v82 := GenS4{F0: -9223372036854775808, F1: 0}
                            fmt.Println(fi77, v2, fi77)
                            switch fi77 {
                            case 32767:
                                v83 := GenS9{F0: false, F1: (^(v54))}
                                _ = v83
                            default:
                                v4 = float32((1.5 + 1.5))
                            }
                            var v84 int16 = 1
                            _ = v82
                            _ = v84
                        }
                        var v85 int32 = -1
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        _ = fi78
                        _ = v79
                        _ = v85
                    }
                    for fi86 := 0; fi86 < 19; fi86++ {
                        fmt.Println(v54, fi77)
                        var v87 uint16 = 255
                        v4 += v4
                        if ((v4 - 0.1) < 1.5) {
                            var v88 uint16 = v54
                            var v89 string = "hello"
                            v90 := GenS10{F0: -859}
                            v91 := v3[-484]
                            _ = v88
                            _ = v89
                            _ = v90
                            _ = v91
                        } else {
                            if ((v87 >> ((uint(v87)) & 63)) >= (v54 / (((v87 - 32767)) | 1))) {
                                switch fi86 {
                                case 1:
                                    var v92 int8 = -28
                                    _ = v92
                                default:
                                    switch (("\n\t" + "hello") + "hello") {
                                    case "foo bar":
                                        if true {
                                            var v93 int16 = 0
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v93
                                        }
                                    case "\n\t":
                                        fmt.Println(v54, v4)
                                    case "a":
                                        var v94 int64 = 42
                                        _ = v94
                                    default:
                                        var v95 int64 = -1
                                        _ = v95
                                    }
                                }
                                var v96 uint64 = 2
                                var v97 int64 = 1
                                _ = v96
                                _ = v97
                            } else {
                                v4 += float32((3.14159 - float64(v4)))
                                fmt.Println(v4)
                            }
                        }
                        v4 = (v4 + v4)
                        _ = fi86
                        _ = v87
                    }
                    var v98 string = ("hello" + "foo bar")
                    _ = v98
                } else {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    var v99 int8 = 54
                    if true {
                        _ = len(v3)
                        for fi100 := 0; fi100 < 4; fi100++ {
                            v101 := GenS4{F0: 0, F1: 0}
                            var v102 int32 = 42
                            _ = fi100
                            _ = v101
                            _ = v102
                        }
                    }
                    _ = v99
                }
                fmt.Println(v2)
                _ = fi77
            }
        }
    }
    v54 = (^(v54))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
}
