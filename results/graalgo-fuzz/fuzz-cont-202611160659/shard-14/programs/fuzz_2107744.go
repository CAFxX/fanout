package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint16; F1 int; F2 int64 }
type GenS3 struct { F0 uint32; F1 int; F2 int64 }
type GenS4 struct { F0 uint16; F1 int8; F2 int }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 uint; F2 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16; F1 int64; F2 uint }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 int }

func main() {
    v1 := []int{128, 725, 32767}
    v2 := []uint32{625, 424}
    var v3 int32 = 634
    var v4 int16 = 1000
    switch ("" + ("foo bar" + "a")) {
    case "":
        v5 := GenS1{F0: 1, F1: 42, F2: -9223372036854775808}
        _ = v5
    case "foo bar":
        var v6 string = ""
        _ = v6
    }
    v7 := GenS1{F0: -2147483648, F1: 136, F2: 42}
    var v8 bool = (true || (v3 > -1))
    if v8 {
        var v9 uint8 = 127
        _ = v9
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint16 = 255
        _ = fi10
        _ = v11
    }
    if (-72 >= 1) {
        var v12 int = 1
        switch "a" {
        case "hello":
            var v13 float32 = 0.1
            _ = v13
        }
        var v14 float64 = (1e10 / 1.5)
        _ = v12
        _ = v14
    } else {
        fmt.Println(v3, v4)
        fmt.Println(v4, v8)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v15 := GenS1{F0: 9223372036854775807, F1: 78, F2: 42}
    v4 = v4
    var v16 uint16 = 42
    v8 = v8
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS2{F0: 0, F1: int(v3), F2: -1}
            if (fi18 >= (fi17 >> ((206) & 63))) {
                fmt.Println(v4)
                var v20 uint64 = 585
                v3 = int32(v16)
                _ = v20
            } else {
                if v8 {
                    v16 = (v16 + v16)
                    fmt.Println(fi17, v4, v3)
                }
                if (0 != 85) {
                    v21 := GenS2{F0: (v16 % (((v16 / (((v16 - 1)) | 1))) | 1)), F1: fi17, F2: -168}
                    _ = v21
                }
                if len(v2) > 0 {
                    v2[0] = 1
                }
            }
            _ = len(v1)
            var v22 uint8 = 76
            _ = fi18
            _ = v19
            _ = v22
        }
        var v23 uint = 65535
        _ = fi17
        _ = v23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        v3 = -247
        _ = fi24
    }
    switch ("hello" + "") {
    case "hello":
        v25 := GenS1{F0: 383, F1: 1, F2: -1}
        _ = v25
    case "a":
        for fi26 := 0; fi26 < 1; fi26++ {
            switch "foo bar" {
            case "foo bar":
                if (!(42 >= 1)) {
                    var v27 uint16 = v16
                    _ = v27
                }
            case "\n\t":
                fmt.Println(v4)
            default:
                for fi28 := 0; fi28 < 19; fi28++ {
                    v16 = 65535
                    v29 := GenS3{F0: 2147483648, F1: 1, F2: -598}
                    var v30 uint = 731
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            _ = fi26
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (!v8) {
        var v31 int8 = 0
        _ = v31
    }
    if (42 >= -1) {
        v32 := GenS4{F0: 1, F1: 100, F2: 32767}
        v4 -= (-(v4))
        if v8 {
            if v8 {
                _ = len(v1)
                var v33 int32 = (-(v3))
                _ = v33
            } else {
                fmt.Println(v3, v8)
                v4 = (v4 << ((42) & 63))
            }
            v8 = (189 > (^(v16)))
            var v34 float32 = float32(((0.0 / 3.14159) - (1e10 * ((0.0 * -2.25) - -2.25))))
            _ = v34
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi35 := 0; fi35 < 2; fi35++ {
                _ = len(v1)
                var v36 uint16 = 255
                switch (fi35 % (((fi35 * (fi35 / (((fi35 * (fi35 / ((197) | 1)))) | 1)))) | 1)) {
                case 4294967295:
                    v36 += v36
                default:
                    switch ("" + "a") {
                    case "foo bar":
                        if v8 {
                            var v37 float32 = 1.5
                            fmt.Println(v4, v16, v3)
                            _ = v37
                        }
                    case "":
                        for fi38 := 0; fi38 < 16; fi38++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v39 bool = ((^(v3)) >= (v3 << ((0) & 63)))
                            var v40 int8 = -1
                            v16 = 414
                            v41 := GenS5{F0: ("a" + (("\n\t" + "hello") + "hello")), F1: 133}
                            v42 := GenS2{F0: (v16 >> ((944) & 63)), F1: (fi38 + (-(fi35))), F2: 572}
                            v43 := GenS2{F0: v16, F1: fi38, F2: 217}
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                    case "\n\t":
                        var v44 uint16 = 204
                        _ = v44
                    default:
                        v3 = v3
                    }
                }
                _ = fi35
                _ = v36
            }
            _ = len(v1)
        }
        _ = v32
    } else {
        for fi45 := 0; fi45 < 6; fi45++ {
            for fi46 := 0; fi46 < 17; fi46++ {
                if len(v2) > 0 {
                    v2[0] = 2
                }
                if true {
                    v47 := GenS3{F0: 65535, F1: fi45, F2: 0}
                    v48 := GenS6{F0: 42}
                    if v8 {
                        if false {
                            v16 = v16
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch 32768 {
                        case 128:
                            if v8 {
                                if true {
                                    var v49 int64 = 172
                                    if len(v2) > 0 {
                                        v2[0] = 4294967295
                                    }
                                    v2 = append(v2, 4294967295)
                                    v50 := GenS7{F0: (^(v16)), F1: 648, F2: v8}
                                    _ = v49
                                    _ = v50
                                }
                                _ = len(v2)
                                var v51 float32 = float32(1e10)
                                _ = v51
                            }
                        case -84:
                            switch (fi46 << ((0) & 63)) {
                            case -971:
                                if (v16 <= v16) {
                                    var v52 int = (fi45 | -662)
                                    _ = v52
                                } else {
                                    var v53 int16 = v4
                                    _ = v53
                                }
                                v8 = (42 <= 127)
                                fmt.Println(v3, v8)
                            case 4294967295:
                                fmt.Println(v16)
                            default:
                                if (!(1e10 == (1e10 - ((1e10 * ((-2.25 + 3.14159) + 1.5)) * (float32(0.0) - 1.5))))) {
                                    fmt.Println(v3)
                                    v8 = true
                                    v54 := GenS8{F0: ""}
                                    _ = v54
                                }
                            }
                        default:
                            v3 = (v3 & 1)
                        }
                    }
                    _ = v47
                    _ = v48
                }
                if v8 {
                    switch (-(fi45)) {
                    case 4294967295:
                        if true {
                            v3 = v3
                            var v55 uint = 4294967295
                            if v8 {
                                if len(v1) > 0 {
                                    v1[0] = (fi46 ^ (fi45 >> (((v55 + v55)) & 63)))
                                }
                                if (-1 != v3) {
                                    v4 += v4
                                    var v56 uint16 = 32767
                                    _ = v56
                                }
                            }
                            _ = v55
                        }
                    case 9223372036854775807:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v57 int = fi46
                    v4 = (^(v4))
                    _ = v57
                } else {
                    var v58 int8 = -1
                    v58 = v58
                    _ = v58
                }
                if false {
                    v16 = (v16 ^ 42)
                    switch ("a" + "foo bar") {
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = fi45
                        }
                        if (v8 || (v4 == v4)) {
                            v4 = (v4 / ((v4) | 1))
                        }
                        switch (("" + "foo bar") + "") {
                        case "\n\t":
                            v4 += (v4 << ((957) & 63))
                        case "hello":
                            var v59 bool = false
                            _ = v59
                        default:
                            var v60 uint = 998
                            _ = v60
                        }
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v3 -= 2147483647
                    v61 := GenS3{F0: 259, F1: fi46, F2: 994}
                    _ = v61
                } else {
                    var v62 bool = v8
                    v1 = append(v1, (fi45 ^ fi45))
                    v8 = v62
                    var v63 string = "foo bar"
                    v64 := GenS9{F0: v4, F1: 881, F2: 512}
                    _ = v62
                    _ = v63
                    _ = v64
                }
                _ = fi46
            }
            if v8 {
                if (v3 < int32(v4)) {
                    for fi65 := 0; fi65 < 16; fi65++ {
                        switch fi65 {
                        case 2147483647:
                            v8 = (0 > -1)
                        case 0:
                            switch "\n\t" {
                            case "hello":
                                v8 = v8
                                switch fi65 {
                                case 127:
                                    if v8 {
                                        if (658 <= 1) {
                                            var v66 uint32 = 2147483648
                                            _ = v66
                                        } else {
                                            v3 += v3
                                            v4 += v4
                                        }
                                        if ((v3 * v3) >= (v3 << ((18446744073709551615) & 63))) {
                                            if true {
                                                v67 := GenS6{F0: 4294967295}
                                                _ = v67
                                            } else {
                                                var v68 int32 = 987
                                                var v69 uint8 = 51
                                                v70 := GenS1{F0: fi45, F1: 127, F2: 0}
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                            }
                                            var v71 uint = 224
                                            if len(v2) > 0 {
                                                v2[0] = 2147483648
                                            }
                                            _ = v71
                                        } else {
                                            v4 -= v4
                                        }
                                        v3 += (^(v3))
                                    }
                                    if v8 {
                                        if v8 {
                                            var v72 uint16 = (v16 + v16)
                                            if len(v1) > 0 {
                                                v1[0] = -436
                                            }
                                            var v73 int = fi65
                                            _ = v72
                                            _ = v73
                                        } else {
                                            if ((1 > 127) && v8) {
                                                if true {
                                                    v8 = (!(!v8))
                                                    v1 = append(v1, (fi65 ^ (fi45 / ((fi65) | 1))))
                                                }
                                                var v74 uint8 = 2
                                                switch fi65 {
                                                case 376:
                                                    if len(v1) > 0 {
                                                        v1[0] = (-(fi45))
                                                    }
                                                }
                                                _ = v74
                                            } else {
                                                var v75 uint8 = 127
                                                fmt.Println(v8, fi65, v16)
                                                var v76 int = fi65
                                                switch ("hello" + "a") {
                                                case "a":
                                                    var v77 int64 = 0
                                                    _ = v77
                                                case "":
                                                    v78 := GenS9{F0: (v4 / (((v4 << ((65535) & 63))) | 1)), F1: -9223372036854775808, F2: 255}
                                                    _ = v78
                                                default:
                                                    v75 = (v75 / ((128) | 1))
                                                }
                                                fmt.Println(v75, v3)
                                                v2 = append(v2, 56)
                                                _ = v75
                                                _ = v76
                                            }
                                        }
                                        v79 := GenS5{F0: "foo bar", F1: 42}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v79
                                    }
                                case 333:
                                    if (((!(v8 && false)) || ((v16 / (((v16 / ((v16) | 1))) | 1)) <= (^(v16)))) || (-1 >= 0)) {
                                        var v80 int8 = 1
                                        _ = v80
                                    } else {
                                        v4 -= v4
                                        v81 := GenS8{F0: "\n\t"}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v81
                                    }
                                }
                                v8 = v8
                            case "\n\t":
                                v4 += (v4 % ((32767) | 1))
                            case "":
                                var v82 uint8 = 13
                                _ = v82
                            default:
                                switch "\n\t" {
                                case "":
                                    v8 = true
                                }
                            }
                        default:
                            v4 = (v4 - v4)
                        }
                        _ = fi65
                    }
                    v1 = append(v1, fi45)
                    switch "" {
                    case "\n\t":
                        var v83 int = fi45
                        _ = v83
                    case "hello":
                        fmt.Println(fi45, v4, v8)
                    case "a":
                        fmt.Println(v16, v16)
                    default:
                        v84 := GenS6{F0: 164}
                        _ = v84
                    }
                } else {
                    v3 = 197
                    fmt.Println(fi45, v4)
                }
                if ((v3 ^ v3) == v3) {
                    v16 = (v16 << ((4294967295) & 63))
                    v85 := GenS4{F0: v16, F1: -33, F2: fi45}
                    if (42 != 65535) {
                        if ((-2.25 * 3.14159) >= 0.1) {
                            var v86 int32 = 2147483647
                            v2 = append(v2, 618)
                            _ = v86
                        }
                        v16 -= (v16 << ((42) & 63))
                    }
                    _ = v85
                }
                v8 = (((float64(((-2.25 - (0.0 * 0.0)) / 3.14159)) + (float64(0.1) / (-2.25 + (1e10 - ((3.14159 * ((3.14159 * (3.14159 + 3.14159)) + -2.25)) + (-2.25 - 1.5)))))) < 0.0) || false)
            } else {
                if v8 {
                    v2 = append(v2, 2)
                    var v87 uint = 42
                    _ = v87
                }
                if ((v3 * 42) >= (^(v3))) {
                    switch 228 {
                    case -554:
                        for fi88 := 0; fi88 < 18; fi88++ {
                            v4 = v4
                            _ = fi88
                        }
                    case -1:
                        var v89 int8 = 127
                        _ = v89
                    default:
                        if (!false) {
                            _ = len(v2)
                            v1 = append(v1, (fi45 / ((4294967295) | 1)))
                        } else {
                            if false {
                                switch -852 {
                                case 2147483647:
                                    v90 := GenS7{F0: v16, F1: 203, F2: (0.1 >= 1e10)}
                                    _ = v90
                                case 0:
                                    v1 = append(v1, (fi45 << ((1) & 63)))
                                case 36:
                                    fmt.Println(v8)
                                default:
                                    v2 = append(v2, 2147483647)
                                }
                                v4 -= v4
                                v4 -= v4
                            } else {
                                fmt.Println(v16)
                                switch (fi45 >> ((1) & 63)) {
                                case 9223372036854775807:
                                    v8 = v8
                                default:
                                    switch (((("foo bar" + "hello") + "a") + "a") + "\n\t") {
                                    case "hello":
                                        var v91 float32 = (0.0 - 1.5)
                                        _ = v91
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "a":
                                            v92 := GenS3{F0: 79, F1: (fi45 & fi45), F2: 42}
                                            _ = v92
                                        default:
                                            v93 := GenS1{F0: 128, F1: 235, F2: 1}
                                            _ = v93
                                        }
                                    }
                                }
                                for fi94 := 0; fi94 < 15; fi94++ {
                                    v95 := GenS7{F0: 291, F1: 343, F2: ((!true) && false)}
                                    var v96 uint = 543
                                    v97 := GenS10{F0: 42, F1: fi94}
                                    _ = fi94
                                    _ = v95
                                    _ = v96
                                    _ = v97
                                }
                            }
                        }
                    }
                } else {
                    v98 := GenS4{F0: (v16 >> ((65535) & 63)), F1: -1, F2: 288}
                    _ = v98
                }
            }
            v99 := GenS11{F0: (fi45 % (((-(fi45))) | 1))}
            _ = fi45
            _ = v99
        }
        var v100 int8 = -83
        fmt.Println(v16)
        _ = v100
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
