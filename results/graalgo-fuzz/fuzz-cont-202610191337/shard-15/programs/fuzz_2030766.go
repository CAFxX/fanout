package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint64; F1 int8; F2 string }
type GenS4 struct { F0 uint; F1 int16; F2 int16 }
type GenS5 struct { F0 bool; F1 int; F2 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 uint8; F1 uint }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 int64; F1 bool }
type GenS12 struct { F0 int8; F1 int64 }
type GenS13 struct { F0 uint32; F1 uint }
type GenS14 struct { F0 uint64; F1 int8 }

func main() {
    v1 := []string{"hello"}
    var v2 float32 = (3.14159 / 1.0)
    v3 := []uint32{}
    var v4 uint8 = 81
    var v5 uint32 = 57
    switch 42 {
    case -781:
        v5 -= 1
    default:
        var v6 int16 = -432
        _ = v6
    }
    if false {
        var v7 uint64 = 2
        for fi8 := 0; fi8 < 10; fi8++ {
            if (!true) {
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 uint16 = 1
                    switch "foo bar" {
                    case "\n\t":
                        if (-2.25 > 0.1) {
                            v11 := GenS1{F0: 0, F1: 0}
                            switch (fi9 >> ((4294967295) & 63)) {
                            case -1:
                                v7 = (v7 | v7)
                            case -781:
                                switch 550 {
                                case -2147483648:
                                    v2 -= -2.25
                                }
                                var v12 bool = ((v4 | 127) != (v4 >> ((1) & 63)))
                                if ((v5 == v5) && true) {
                                    v5 = 2147483647
                                    v1 = append(v1, "foo bar")
                                } else {
                                    v13 := GenS2{F0: ("a" + "foo bar"), F1: 0}
                                    var v14 uint8 = v4
                                    fmt.Println(v10, fi9)
                                    _ = v13
                                    _ = v14
                                }
                                _ = v12
                            default:
                                v15 := GenS2{F0: (("" + "foo bar") + "a"), F1: 9223372036854775807}
                                _ = v15
                            }
                            _ = v11
                        } else {
                            if ((("" + ("" + "")) < "foo bar") || (455 > 255)) {
                                v3 = append(v3, uint32(v4))
                                var v16 float64 = -2.25
                                v17 := GenS3{F0: 4294967295, F1: 53, F2: ("a" + "\n\t")}
                                _ = v16
                                _ = v17
                            } else {
                                v2 += v2
                            }
                        }
                    case "hello":
                        fmt.Println(v4, v4)
                    }
                    _ = fi9
                    _ = v10
                }
                for fi18 := 0; fi18 < 1; fi18++ {
                    if (!false) {
                        if (v4 <= (v4 | (v4 << ((640) & 63)))) {
                            var v19 uint64 = v7
                            fmt.Println(fi18)
                            v5 = 2147483648
                            _ = v19
                        }
                        v20 := GenS4{F0: 65535, F1: 1000, F2: 0}
                        v21 := GenS2{F0: ("" + "hello"), F1: 721}
                        _ = v20
                        _ = v21
                    }
                    v4 = v4
                    var v22 string = "hello"
                    _ = fi18
                    _ = v22
                }
                if len(v3) > 0 {
                    v3[0] = (^(v5))
                }
            } else {
                v23 := GenS1{F0: 100, F1: 9223372036854775807}
                v4 += 0
                var v24 uint64 = v7
                _ = v23
                _ = v24
            }
            if ("hello" < (("foo bar" + "a") + "a")) {
                var v25 int8 = 127
                v5 = v5
                _ = v25
            }
            _ = fi8
        }
        _ = v7
    }
    var v26 float64 = ((3.14159 + 1.5) - 1.5)
    for fi27 := 0; fi27 < 19; fi27++ {
        v4 -= 4
        v3 = append(v3, (^(v5)))
        _ = fi27
    }
    var v28 uint64 = 2
    fmt.Println(v2, v26)
    var v29 uint32 = v5
    v2 = v2
    for fi30 := 0; fi30 < 4; fi30++ {
        if true {
            fmt.Println(v28)
            if (v2 <= 1e10) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v31 := GenS2{F0: "hello", F1: -268}
                v4 += v4
                _ = v31
            } else {
                var v32 int8 = 42
                if (true || (true || false)) {
                    switch int(v28) {
                    case 629:
                        var v33 int = (fi30 % ((9223372036854775807) | 1))
                        _ = v33
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                } else {
                    v26 -= v26
                }
                _ = v32
            }
        } else {
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 uint16 = 2
                _ = fi34
                _ = v35
            }
            switch (fi30 ^ 32768) {
            case -1:
                if true {
                    switch (fi30 | (fi30 >> ((1) & 63))) {
                    case 128:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 2147483647:
                        var v36 int64 = 42
                        _ = v36
                    default:
                        v37 := GenS4{F0: 18446744073709551615, F1: 32767, F2: -32768}
                        _ = v37
                    }
                }
            case 248:
                if ("\n\t" < "a") {
                    fmt.Println(v2, v26, v26)
                    if (42 > 1) {
                        fmt.Println(v5)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi38 := 0; fi38 < 16; fi38++ {
                            var v39 float64 = (float64((v2 + (0.0 / 0.1))) + (v26 + v26))
                            v40 := GenS3{F0: v28, F1: -1, F2: "hello"}
                            v29 = v5
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                    }
                } else {
                    v41 := GenS5{F0: true, F1: 9223372036854775807, F2: 575}
                    v1 = append(v1, ("a" + ("\n\t" + "\n\t")))
                    if (false || (false || ((^(v28)) > v28))) {
                        v3 = append(v3, (v29 - v29))
                    }
                    _ = v41
                }
            case -628:
                if (v26 >= v26) {
                    for fi42 := 0; fi42 < 10; fi42++ {
                        var v43 int64 = 42
                        v4 += 127
                        _ = len(v1)
                        _ = fi42
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 17; fi44++ {
                        switch "\n\t" {
                        case "a":
                            if (!(((1 == (v5 | uint32(v28))) || true) || true)) {
                                v29 = (v29 >> ((4294967295) & 63))
                                v2 = float32(float64((float32(float64(1e10)) + 3.14159)))
                                var v45 uint16 = 0
                                _ = v45
                            }
                        case "":
                            v5 = 107
                        }
                        _ = len(v3)
                        var v46 int16 = 1
                        fmt.Println(v26, fi30, v26)
                        _ = fi44
                        _ = v46
                    }
                    if ((0.0 / (0.0 - 1.5)) == v2) {
                        v4 = 0
                        _ = len(v3)
                    } else {
                        switch fi30 {
                        case 9223372036854775807:
                            v47 := GenS6{F0: (v4 % ((v4) | 1))}
                            _ = v47
                        default:
                            v5 += (v5 | v5)
                        }
                        v29 = v29
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                } else {
                    var v48 string = ((("a" + "a") + "foo bar") + ("foo bar" + "a"))
                    if (!((v4 << ((uint(v4)) & 63)) < (v4 - (v4 & v4)))) {
                        fmt.Println(v5)
                    } else {
                        switch (v48 + (v48 + (v48 + (v48 + "a")))) {
                        case "a":
                            if false {
                                v48 = v48
                                v49 := GenS4{F0: 195, F1: 1000, F2: -1}
                                _ = v49
                            }
                        case "":
                            if (!((-1 <= -40) || false)) {
                                if true {
                                    _ = len(v3)
                                } else {
                                    for fi50 := 0; fi50 < 8; fi50++ {
                                        v28 = 18446744073709551615
                                        _ = fi50
                                    }
                                    var v51 string = ((v48 + "a") + ((v48 + (((("hello" + (v48 + v48)) + v48) + ((("\n\t" + "a") + v48) + (v48 + (("foo bar" + ((v48 + v48) + v48)) + "foo bar")))) + ("\n\t" + "\n\t"))) + v48))
                                    _ = v51
                                }
                            }
                        case "\n\t":
                            fmt.Println(v29, v5)
                        }
                        if (("\n\t" + v48) == v48) {
                            var v52 uint64 = (^(v28))
                            for fi53 := 0; fi53 < 14; fi53++ {
                                v54 := GenS7{F0: int64(v52), F1: 9223372036854775807}
                                _ = fi53
                                _ = v54
                            }
                            switch ((v48 + v48) + "a") {
                            case "hello":
                                v55 := GenS8{F0: (fi30 | fi30), F1: ""}
                                _ = v55
                            case "\n\t":
                                for fi56 := 0; fi56 < 20; fi56++ {
                                    v57 := GenS1{F0: -64, F1: -1}
                                    if (32767 < 0) {
                                        fmt.Println(v29, v4)
                                        var v58 string = (v48 + (v48 + ((v48 + "\n\t") + "hello")))
                                        switch v48 {
                                        case "a":
                                            if true {
                                                if false {
                                                    v4 = (v4 % ((1) | 1))
                                                    switch v48 {
                                                    case "hello":
                                                        v59 := GenS9{F0: 131, F1: 4294967295}
                                                        _ = v59
                                                    case "a":
                                                        v29 += v29
                                                    default:
                                                        v60 := GenS9{F0: v4, F1: 255}
                                                        v3 = append(v3, v5)
                                                        _ = len(v1)
                                                        _ = v60
                                                    }
                                                    v2 = float32(v26)
                                                } else {
                                                    if (!true) {
                                                        switch (fi30 * (fi30 >> ((761) & 63))) {
                                                        case 9223372036854775807:
                                                            v58 = "a"
                                                        case 42:
                                                            if (true && false) {
                                                                var v61 uint16 = 1
                                                                v62 := GenS8{F0: -35, F1: ""}
                                                                _ = v61
                                                                _ = v62
                                                            }
                                                        default:
                                                            if false {
                                                                var v63 uint32 = v5
                                                                _ = v63
                                                            }
                                                        }
                                                        v5 = (v5 & 736)
                                                    } else {
                                                        var v64 string = "foo bar"
                                                        if true {
                                                            if false {
                                                                v29 = 1
                                                                var v65 uint32 = v29
                                                                _ = v65
                                                            }
                                                        }
                                                        _ = v64
                                                    }
                                                    switch 9223372036854775807 {
                                                    case 9223372036854775807:
                                                        v66 := GenS7{F0: 9223372036854775807, F1: -963}
                                                        _ = v66
                                                    case 127:
                                                        var v67 float64 = 1e10
                                                        _ = v67
                                                    case -1:
                                                        if len(v1) > 0 {
                                                            v1[0] = (v58 + "")
                                                        }
                                                    default:
                                                        var v68 int32 = 1
                                                        _ = v68
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            fmt.Println(v4)
                                        case "":
                                            fmt.Println(v48, v5)
                                        default:
                                            if false {
                                                switch (v58 + "foo bar") {
                                                case "":
                                                    v69 := GenS6{F0: 136}
                                                    _ = v69
                                                case "\n\t":
                                                    if (!(2 < 65535)) {
                                                        fmt.Println(v58, fi56, fi30)
                                                        var v70 uint32 = (v5 % (((v29 ^ 818)) | 1))
                                                        _ = v70
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = (v29 >> ((255) & 63))
                                                        }
                                                        v5 -= v29
                                                        v5 = v5
                                                    }
                                                case "foo bar":
                                                    if false {
                                                        var v71 float64 = (-2.25 / v26)
                                                        _ = v71
                                                    }
                                                default:
                                                    var v72 uint8 = (v4 % ((255) | 1))
                                                    _ = v72
                                                }
                                                fmt.Println(v26, v29)
                                                v5 -= 65535
                                            } else {
                                                v5 += (v29 << ((824) & 63))
                                                var v73 uint32 = v5
                                                fmt.Println(fi30, v29, v52)
                                                _ = v73
                                            }
                                        }
                                        fmt.Println(fi56)
                                        _ = v58
                                    }
                                    var v74 uint64 = v28
                                    _ = fi56
                                    _ = v57
                                    _ = v74
                                }
                            }
                            _ = v52
                        } else {
                            if false {
                                v2 += -2.25
                            } else {
                                fmt.Println(fi30)
                            }
                        }
                    }
                    var v75 float32 = (((3.14159 * (v2 * (v2 + v2))) / (v2 + (v2 / v2))) * v2)
                    _ = v48
                    _ = v75
                }
            }
        }
        fmt.Println(v5, v28)
        var v76 uint8 = v4
        _ = fi30
        _ = v76
    }
    _ = len(v3)
    v77 := GenS10{F0: uint8(v29), F1: 127}
    v2 = ((v2 - v2) / (v2 - ((v2 * v2) + v2)))
    switch "\n\t" {
    case "hello":
        if false {
            for fi78 := 0; fi78 < 2; fi78++ {
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                switch "foo bar" {
                case "hello":
                    v26 = v26
                case "foo bar":
                    if true {
                        fmt.Println(v4, fi78, v28)
                        if true {
                            v26 = (0.0 + 1.5)
                            fmt.Println(v4, v2, v26)
                        }
                    } else {
                        if (!true) {
                            var v79 uint16 = 65535
                            _ = v79
                        } else {
                            var v80 float32 = (v2 * (0.1 * ((float32(1.5) - (float32((v26 + v26)) + float32(((v26 / (v26 - 1.5)) * (-2.25 * v26))))) / 0.1)))
                            fmt.Println(v29, v28, v2)
                            _ = v80
                        }
                        for fi81 := 0; fi81 < 17; fi81++ {
                            if ((0.0 - v2) == (0.0 + v2)) {
                                var v82 uint32 = 1000
                                _ = v82
                            }
                            var v83 int16 = 0
                            var v84 float64 = 1e10
                            _ = fi81
                            _ = v83
                            _ = v84
                        }
                    }
                default:
                    v85 := GenS11{F0: 677, F1: (true && (false || (false || (-32768 > 32767))))}
                    _ = v85
                }
                for fi86 := 0; fi86 < 6; fi86++ {
                    v26 = (0.0 / 1.0)
                    var v87 int = (^(fi78))
                    switch ("" + "") {
                    case "hello":
                        fmt.Println(fi86)
                        var v88 uint = 42
                        fmt.Println(v88, v87, v88)
                        _ = v88
                    default:
                        if false {
                            v89 := GenS12{F0: 109, F1: 42}
                            fmt.Println(v29, fi78, v2)
                            _ = v89
                        }
                    }
                    _ = fi86
                    _ = v87
                }
                _ = fi78
            }
        }
    case "":
        var v90 int8 = 100
        _ = v90
    default:
        if (false || true) {
            if (((!false) || true) && (!true)) {
                switch "" {
                case "\n\t":
                    v91 := GenS4{F0: 18446744073709551615, F1: -201, F2: 697}
                    _ = v91
                case "a":
                    for fi92 := 0; fi92 < 13; fi92++ {
                        var v93 float32 = ((((v2 * (1e10 - ((v2 / v2) / (float32((1.5 + v26)) + (1e10 + v2))))) / v2) - v2) + (((0.0 / 3.14159) / float32((((((v26 * 3.14159) * -2.25) - 1e10) + (v26 + float64(1e10))) * -2.25))) - v2))
                        _ = fi92
                        _ = v93
                    }
                case "hello":
                    if len(v1) > 0 {
                        v1[0] = (("a" + "\n\t") + "hello")
                    }
                default:
                    v94 := GenS13{F0: 4294967295, F1: 18446744073709551615}
                    _ = v94
                }
                v95 := GenS14{F0: 2, F1: -43}
                _ = v95
            }
            for fi96 := 0; fi96 < 13; fi96++ {
                var v97 uint8 = 42
                var v98 bool = ((^(v4)) == (v97 & v97))
                _ = fi96
                _ = v97
                _ = v98
            }
            v2 = v2
        }
    }
    var v99 int = 4294967295
    fmt.Println(v99, v4, v29)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v99
}
