package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 int32 }
type GenS3 struct { F0 uint; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 uint32; F2 int }
type GenS7 struct { F0 int64; F1 int64; F2 int32 }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 uint16 = 0
    v2 := []int{518, 32767}
    var v3 uint8 = 227
    var v4 int16 = int16(v1)
    if (4294967295 == 197) {
        if false {
            if (!((-2.25 <= 1.5) || false)) {
                v3 -= v3
                if true {
                    v2 = append(v2, -37)
                } else {
                    v1 += v1
                    var v5 uint64 = 2
                    _ = v5
                }
            }
        } else {
            fmt.Println(v4, v1)
            var v6 bool = ((!true) && ((-2.25 - (1.5 * 1e10)) <= -2.25))
            _ = v6
        }
        if (-53 < -1) {
            fmt.Println(v1, v1)
        } else {
            v2 = append(v2, 32768)
        }
        v4 = v4
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint = 65535
        v1 = (^(v1))
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (v4 >> ((4294967295) & 63)), F1: v1}
    v1 += (v1 << ((356) & 63))
    fmt.Println(v4, v3, v1)
    fmt.Println(v1, v4, v1)
    if (!(((!(77 <= 2)) || true) && false)) {
        var v10 int32 = -902
        var v11 uint = 2
        if ((v1 << (((v11 | (v11 & v11))) & 63)) <= (v1 * v1)) {
            switch "\n\t" {
            case "a":
                for fi12 := 0; fi12 < 9; fi12++ {
                    v11 -= v11
                    fmt.Println(v10)
                    _ = fi12
                }
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 int32 = int32(v3)
                _ = fi13
                _ = v14
            }
            if true {
                switch ("\n\t" + "") {
                case "hello":
                    if false {
                        for fi15 := 0; fi15 < 12; fi15++ {
                            v16 := GenS2{F0: 639, F1: "", F2: v10}
                            switch fi15 {
                            case 83:
                                var v17 float64 = -2.25
                                _ = v17
                            }
                            _ = fi15
                            _ = v16
                        }
                        var v18 bool = false
                        _ = v18
                    } else {
                        if true {
                            var v19 uint = v11
                            _ = v19
                        }
                        fmt.Println(v1)
                        if false {
                            v20 := GenS1{F0: v4, F1: v1}
                            if (1.5 <= -2.25) {
                                for fi21 := 0; fi21 < 9; fi21++ {
                                    if len(v2) > 0 {
                                        v2[0] = fi21
                                    }
                                    v22 := GenS3{F0: v11, F1: int16(v11), F2: 886}
                                    _ = fi21
                                    _ = v22
                                }
                                fmt.Println(v10, v4)
                                var v23 int16 = (-(v4))
                                _ = v23
                            } else {
                                for fi24 := 0; fi24 < 9; fi24++ {
                                    switch (fi24 - (-(fi24))) {
                                    case -9223372036854775808:
                                        switch ("\n\t" + "foo bar") {
                                        case "a":
                                            var v25 string = "hello"
                                            _ = v25
                                        }
                                    case -2147483648:
                                        switch "a" {
                                        case "hello":
                                            for fi26 := 0; fi26 < 6; fi26++ {
                                                var v27 int8 = 1
                                                v27 = v27
                                                v1 += (^(v1))
                                                if (!false) {
                                                    v28 := GenS1{F0: v4, F1: (^(v1))}
                                                    switch "hello" {
                                                    case "":
                                                        fmt.Println(fi26, fi24)
                                                    case "foo bar":
                                                        if (((1e10 * 1e10) - -2.25) <= 3.14159) {
                                                            var v29 int16 = v4
                                                            var v30 float64 = -2.25
                                                            fmt.Println(v29)
                                                            _ = v29
                                                            _ = v30
                                                        } else {
                                                            v31 := GenS4{F0: v11, F1: 2147483647}
                                                            _ = v31
                                                        }
                                                        var v32 uint = (v11 - (v11 ^ 65535))
                                                        var v33 uint32 = 2147483648
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if true {
                                                                var v34 int64 = -1
                                                                if true {
                                                                    var v35 float32 = 0.0
                                                                    if false {
                                                                        var v36 float64 = 1.5
                                                                        fmt.Println(v35, v3)
                                                                        _ = v36
                                                                    } else {
                                                                        _ = len(v2)
                                                                    }
                                                                    _ = v35
                                                                } else {
                                                                    v37 := GenS4{F0: 4294967295, F1: (v10 % ((v10) | 1))}
                                                                    v38 := GenS5{F0: false}
                                                                    fmt.Println(v3, v32)
                                                                    _ = v37
                                                                    _ = v38
                                                                }
                                                                _ = v34
                                                            } else {
                                                                if len(v2) > 0 {
                                                                    v2[0] = (-(fi26))
                                                                }
                                                                if (!(((fi24 | (fi24 & (fi24 * 127))) < (fi26 * (fi24 % (((fi26 << ((1) & 63))) | 1)))) || (!true))) {
                                                                    if true {
                                                                        var v39 int8 = 100
                                                                        _ = v39
                                                                    } else {
                                                                        var v40 bool = (!(!false))
                                                                        _ = v40
                                                                    }
                                                                    var v41 bool = false
                                                                    var v42 string = (("a" + ("a" + "a")) + ("a" + "foo bar"))
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                                v10 += (v10 * 0)
                                                            }
                                                        case "foo bar":
                                                            v43 := GenS4{F0: (v11 - 0), F1: (v10 - -1)}
                                                            _ = v43
                                                        }
                                                        _ = v32
                                                        _ = v33
                                                    case "hello":
                                                        var v44 uint32 = 0
                                                        _ = v44
                                                    }
                                                    v4 += v4
                                                    var v45 uint16 = v1
                                                    if (4294967295 > (v11 & (v11 ^ (v11 % (((v11 << (((^(v11))) & 63))) | 1))))) {
                                                        v2 = append(v2, fi26)
                                                    }
                                                    _ = v28
                                                    _ = v45
                                                } else {
                                                    switch (fi24 % ((int(v3)) | 1)) {
                                                    case 9223372036854775807:
                                                        _ = len(v2)
                                                    case 269:
                                                        switch 0 {
                                                        case -9223372036854775808:
                                                            fmt.Println(v4, v4, v11)
                                                        case 0:
                                                            if true {
                                                                var v46 float32 = -2.25
                                                                _ = v46
                                                            } else {
                                                                if ((v27 * (v27 >> ((42) & 63))) >= v27) {
                                                                    fmt.Println(fi26, v3)
                                                                    fmt.Println(v1, fi26)
                                                                }
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = fi26
                                                _ = v27
                                            }
                                        }
                                    case 32768:
                                        fmt.Println(v3, fi24)
                                    }
                                    var v47 uint16 = v1
                                    _ = fi24
                                    _ = v47
                                }
                            }
                            _ = v20
                        }
                    }
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = 127
                    }
                case "\n\t":
                    if false {
                        if (-128 >= -128) {
                            fmt.Println(v3)
                        } else {
                            var v48 int64 = -9223372036854775808
                            v1 -= (v1 ^ (^(v1)))
                            v10 = v10
                            _ = v48
                        }
                        switch 364 {
                        case 2147483647:
                            v2 = append(v2, int(v3))
                        case 42:
                            if (!false) {
                                var v49 int32 = v10
                                _ = v49
                            }
                        }
                        for fi50 := 0; fi50 < 9; fi50++ {
                            v2 = append(v2, 661)
                            _ = fi50
                        }
                    }
                default:
                    var v51 int64 = -1
                    _ = v51
                }
                if (615 < 65535) {
                    if len(v2) > 0 {
                        v2[0] = -519
                    }
                    v2 = append(v2, 4294967295)
                    if true {
                        switch 902 {
                        case 9223372036854775807:
                            v3 = (v3 << ((516) & 63))
                        case -2147483648:
                            var v52 int64 = int64(v10)
                            _ = v52
                        case 0:
                            for fi53 := 0; fi53 < 10; fi53++ {
                                v4 = (v4 % (((v4 & (v4 - -106))) | 1))
                                fmt.Println(v10)
                                if false {
                                    _ = len(v2)
                                    v11 = (^(v11))
                                    var v54 string = "\n\t"
                                    _ = v54
                                }
                                _ = fi53
                            }
                        default:
                            var v55 float64 = -2.25
                            _ = v55
                        }
                        fmt.Println(v11, v3, v10)
                    } else {
                        switch 0 {
                        case 9223372036854775807:
                            v11 += 607
                        case 128:
                            switch ("" + "\n\t") {
                            case "hello":
                                switch 9223372036854775807 {
                                case 127:
                                    v11 -= (v11 * 1)
                                case 9223372036854775807:
                                    v3 -= (v3 | v3)
                                case 2147483647:
                                    for fi56 := 0; fi56 < 2; fi56++ {
                                        fmt.Println(v4, v10)
                                        if (!(!false)) {
                                            fmt.Println(fi56, v3, v3)
                                            for fi57 := 0; fi57 < 3; fi57++ {
                                                v58 := GenS3{F0: (v11 | (^(v11))), F1: v4, F2: 42}
                                                _ = fi57
                                                _ = v58
                                            }
                                        }
                                        _ = fi56
                                    }
                                }
                            case "\n\t":
                                v1 = 1
                            }
                        case 127:
                            fmt.Println(v3, v1, v10)
                        }
                    }
                } else {
                    for fi59 := 0; fi59 < 12; fi59++ {
                        if (!true) {
                            for fi60 := 0; fi60 < 13; fi60++ {
                                if false {
                                    var v61 int16 = v4
                                    var v62 int32 = v10
                                    var v63 string = "a"
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                } else {
                                    var v64 float32 = 1e10
                                    v3 -= v3
                                    v10 = v10
                                    _ = v64
                                }
                                v4 -= v4
                                _ = fi60
                            }
                            v2 = append(v2, (fi59 >> (((v11 % ((v11) | 1))) & 63)))
                            if false {
                                if (v3 >= 2) {
                                    v2 = append(v2, 128)
                                    for fi65 := 0; fi65 < 2; fi65++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v11 -= (v11 >> ((4294967295) & 63))
                                        _ = fi65
                                    }
                                }
                                var v66 uint = (^(v11))
                                var v67 int8 = -1
                                _ = v66
                                _ = v67
                            } else {
                                v11 = 255
                                var v68 float64 = 1.5
                                _ = v68
                            }
                        }
                        fmt.Println(fi59, v3, v1)
                        switch (fi59 | (fi59 | (-(fi59)))) {
                        case -2147483648:
                            for fi69 := 0; fi69 < 5; fi69++ {
                                var v70 int64 = int64(fi69)
                                if true {
                                    v3 = v3
                                    if (true || ((!false) && ("a" >= ("foo bar" + "")))) {
                                        fmt.Println(fi69, v3)
                                        v10 = (^(v10))
                                        if true {
                                            v71 := GenS6{F0: fi69, F1: 0, F2: (fi59 + fi59)}
                                            if true {
                                                if false {
                                                    fmt.Println(v3, v1)
                                                    var v72 int64 = int64(v1)
                                                    _ = v72
                                                }
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v71
                                        }
                                    }
                                }
                                _ = fi69
                                _ = v70
                            }
                        case -560:
                            var v73 float64 = 0.1
                            _ = v73
                        }
                        _ = fi59
                    }
                    for fi74 := 0; fi74 < 20; fi74++ {
                        if false {
                            for fi75 := 0; fi75 < 2; fi75++ {
                                fmt.Println(fi74, fi75)
                                if true {
                                    v2 = append(v2, (fi74 * -2147483648))
                                    var v76 string = "hello"
                                    if false {
                                        var v77 int16 = v4
                                        _ = v77
                                    }
                                    _ = v76
                                } else {
                                    var v78 int = fi74
                                    var v79 int64 = 1
                                    v80 := GenS7{F0: v79, F1: (v79 | 9223372036854775807), F2: -748}
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                                var v81 int32 = (v10 ^ (^(v10)))
                                if false {
                                    var v82 uint = v11
                                    _ = v82
                                } else {
                                    switch fi74 {
                                    case 32768:
                                        v83 := GenS1{F0: v4, F1: v1}
                                        _ = v83
                                    default:
                                        v2 = append(v2, (fi74 / ((32768) | 1)))
                                    }
                                    fmt.Println(v11, v81)
                                }
                                switch (fi75 | 1) {
                                case 752:
                                    v84 := GenS6{F0: (fi75 | fi75), F1: 2147483648, F2: fi75}
                                    _ = v84
                                case 128:
                                    fmt.Println(v1, fi74)
                                case -9223372036854775808:
                                    if false {
                                        v11 -= v11
                                    } else {
                                        v81 += (-(v10))
                                        switch fi75 {
                                        case 9223372036854775807:
                                            v85 := GenS5{F0: (-82 == 127)}
                                            v10 = v10
                                            _ = v85
                                        case -1:
                                            if (!((v4 + v4) == (v4 * (v4 + (^(v4)))))) {
                                                if (!false) {
                                                    if (-109 == 100) {
                                                        v81 = (v10 | v81)
                                                        var v86 string = "hello"
                                                        var v87 float64 = -2.25
                                                        _ = v86
                                                        _ = v87
                                                    } else {
                                                        fmt.Println(fi74, fi74, v1)
                                                    }
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = (-(fi74))
                                                }
                                                v88 := GenS4{F0: (v11 >> (((v11 % ((989) | 1))) & 63)), F1: v10}
                                                _ = v88
                                            } else {
                                                v89 := GenS8{F0: 188}
                                                _ = v89
                                            }
                                        }
                                    }
                                }
                                fmt.Println(fi74)
                                var v90 int8 = 1
                                _ = fi75
                                _ = v81
                                _ = v90
                            }
                        } else {
                            var v91 string = "a"
                            v4 = (v4 & (v4 * v4))
                            _ = v91
                        }
                        _ = fi74
                    }
                }
                switch 711 {
                case 240:
                    for fi92 := 0; fi92 < 1; fi92++ {
                        var v93 uint32 = uint32(v1)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (!(v1 < 32767)) {
                            v93 = (v93 % (((^(v93))) | 1))
                            for fi94 := 0; fi94 < 18; fi94++ {
                                var v95 uint64 = 4294967295
                                switch "" {
                                case "\n\t":
                                    v93 = (v93 + 2147483648)
                                case "hello":
                                    v93 += (^(v93))
                                }
                                var v96 int16 = (v4 >> ((v11) & 63))
                                _ = len(v2)
                                if (0.1 >= 1.5) {
                                    if false {
                                        v10 += int32(v11)
                                        fmt.Println(v11, fi92, v4)
                                        v2 = append(v2, (fi92 | fi94))
                                        v97 := GenS8{F0: v3}
                                        _ = v97
                                    }
                                }
                                _ = fi94
                                _ = v95
                                _ = v96
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = fi92
                        _ = v93
                    }
                case 788:
                    v98 := GenS5{F0: false}
                    _ = v98
                case -1:
                    v1 = (^(v1))
                default:
                    if false {
                        v99 := GenS8{F0: v3}
                        switch 128 {
                        case -9223372036854775808:
                            v1 = v1
                        case 0:
                            if len(v2) > 0 {
                                v2[0] = -119
                            }
                        case -901:
                            for fi100 := 0; fi100 < 12; fi100++ {
                                if false {
                                    var v101 uint64 = 42
                                    fmt.Println(v101, v3)
                                    fmt.Println(v11, v101, fi100)
                                    _ = v101
                                } else {
                                    switch (fi100 | 42) {
                                    case -375:
                                        var v102 uint8 = 127
                                        _ = v102
                                    case -314:
                                        var v103 float64 = 0.0
                                        _ = v103
                                    default:
                                        var v104 int32 = (^(v10))
                                        _ = v104
                                    }
                                }
                                fmt.Println(v11, v4, v10)
                                v105 := GenS5{F0: true}
                                _ = fi100
                                _ = v105
                            }
                        default:
                            if (!false) {
                                var v106 string = ("foo bar" + "foo bar")
                                var v107 bool = true
                                v1 = 42
                                _ = v106
                                _ = v107
                            }
                        }
                        _ = v99
                    } else {
                        fmt.Println(v11)
                    }
                }
            }
        }
        _ = v10
        _ = v11
    }
    var v108 string = ""
    v4 += (v4 / ((int16(v3)) | 1))
    var v109 string = v108
    fmt.Println(v109, v108)
    var v110 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
