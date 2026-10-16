package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint; F1 uint8; F2 string }
type GenS8 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS9 struct { F0 uint32; F1 uint16; F2 bool }
type GenS10 struct { F0 uint; F1 int; F2 int8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint8; F1 int64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint32 }
func genFn0() int {
    var v6 int64 = 42
    switch ("" + "hello") {
    case "foo bar":
        v7 := []int{}
        _ = v7
    default:
        var v8 uint = 255
        _ = v8
    }
    _ = v6
    return 868
}

func main() {
    v1 := []int{-2147483648, 1, 4294967295, 327}
    var v2 uint32 = 65535
    var v3 int8 = 0
    var v4 float32 = (1.5 - (1.5 + 3.14159))
    var v5 int16 = 1
    var v9 float64 = 0.0
    _ = len(v1)
    v10 := GenS1{F0: (("foo bar" + "a") + "hello"), F1: "hello"}
    for fi11 := 0; fi11 < 19; fi11++ {
        switch genFn0() {
        case -1:
            var v12 bool = true
            _ = v12
        }
        switch (("foo bar" + ("a" + "hello")) + "foo bar") {
        case "":
            var v13 uint = 636
            _ = v13
        case "foo bar":
            if (v4 != (v4 * v4)) {
                var v14 int32 = 2147483647
                v14 = v14
                _ = v14
            } else {
                v3 += (v3 % (((v3 & (v3 - 127))) | 1))
            }
        case "a":
            for fi15 := 0; fi15 < 16; fi15++ {
                if len(v1) > 0 {
                    v1[0] = 33
                }
                _ = fi15
            }
        }
        _ = fi11
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        v9 -= 0.1
        for fi17 := 0; fi17 < 18; fi17++ {
            if (!false) {
                _ = len(v1)
                v18 := GenS2{F0: v5, F1: 42}
                _ = v18
            }
            var v19 int8 = (v3 | int8(v2))
            v1 = append(v1, 1)
            var v20 int16 = (-(v5))
            switch "foo bar" {
            case "hello":
                var v21 int64 = -994
                _ = v21
            default:
                switch fi17 {
                case 1:
                    var v22 int32 = 0
                    _ = v22
                default:
                    v20 -= v20
                }
            }
            _ = fi17
            _ = v19
            _ = v20
        }
        _ = fi16
    }
    switch "hello" {
    case "\n\t":
        v3 += 127
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        for fi23 := 0; fi23 < 10; fi23++ {
            if ((v9 + v9) <= 3.14159) {
                v9 += ((0.0 / v9) - -2.25)
                v5 = (v5 << ((434) & 63))
            }
            if (!false) {
                if false {
                    v2 = (v2 << ((255) & 63))
                    fmt.Println(fi23)
                }
            } else {
                v1 = append(v1, 4294967295)
                if false {
                    switch 42 {
                    case 127:
                        fmt.Println(v5, v3)
                    default:
                        fmt.Println(v9)
                    }
                    v9 -= v9
                }
            }
            var v24 string = "foo bar"
            _ = fi23
            _ = v24
        }
    }
    _ = len(v1)
    v25 := GenS1{F0: "\n\t", F1: ""}
    if (819 > 65535) {
        if false {
            var v26 uint = 0
            if len(v1) > 0 {
                v1[0] = -2147483648
            }
            if true {
                v4 = 0.0
                if (!false) {
                    switch "\n\t" {
                    case "a":
                        if (v3 == (-(v3))) {
                            var v27 uint8 = 78
                            if true {
                                switch ("" + "foo bar") {
                                case "\n\t":
                                    if (9223372036854775807 == 4294967295) {
                                        fmt.Println(v27, v3, v9)
                                        if true {
                                            v27 += (v27 ^ (v27 | (v27 << (((v26 / ((v26) | 1))) & 63))))
                                            v1 = append(v1, genFn0())
                                            v5 += -54
                                        }
                                        var v28 uint64 = 927
                                        _ = v28
                                    } else {
                                        v3 = (v3 ^ v3)
                                        v29 := GenS3{F0: -830, F1: 1}
                                        v1 = append(v1, int(v3))
                                        _ = v29
                                    }
                                case "hello":
                                    for fi30 := 0; fi30 < 2; fi30++ {
                                        var v31 uint64 = 492
                                        _ = fi30
                                        _ = v31
                                    }
                                default:
                                    fmt.Println(v3, v4, v5)
                                }
                                for fi32 := 0; fi32 < 12; fi32++ {
                                    v5 = int16(v27)
                                    v33 := GenS4{F0: 2147483647, F1: (243 != (v2 * (v2 - (v2 % ((v2) | 1)))))}
                                    v27 += 2
                                    _ = fi32
                                    _ = v33
                                }
                                _ = len(v1)
                            }
                            _ = v27
                        }
                    default:
                        var v34 bool = (!((v26 | v26) > v26))
                        _ = v34
                    }
                    switch "hello" {
                    case "a":
                        if (!(genFn0() <= 4294967295)) {
                            for fi35 := 0; fi35 < 19; fi35++ {
                                var v36 int8 = v3
                                switch (fi35 - int(v5)) {
                                case -9223372036854775808:
                                    if (((true && (18446744073709551615 < 2)) && true) || (!((v2 | 0) < v2))) {
                                        var v37 int64 = 0
                                        _ = v37
                                    }
                                case -349:
                                    var v38 float32 = (1e10 * (0.1 * v4))
                                    _ = v38
                                case 42:
                                    fmt.Println(v36, v5)
                                default:
                                    var v39 float32 = (v4 * v4)
                                    _ = v39
                                }
                                _ = fi35
                                _ = v36
                            }
                        }
                    case "":
                        v26 -= (v26 | (v26 | 262))
                    }
                    switch 32767 {
                    case 9223372036854775807:
                        v9 -= ((float64(float32(0.1)) + v9) + ((-2.25 + 3.14159) + (0.1 / v9)))
                    case 0:
                        v40 := GenS5{F0: -935, F1: 626}
                        _ = v40
                    default:
                        if ((42 >= 0) && false) {
                            switch 9223372036854775807 {
                            case 1:
                                switch "" {
                                case "foo bar":
                                    if (913 <= 65535) {
                                        if false {
                                            v1 = append(v1, 2147483647)
                                            v3 += v3
                                        } else {
                                            var v41 int = -809
                                            var v42 uint = 18446744073709551615
                                            _ = v41
                                            _ = v42
                                        }
                                        var v43 uint16 = 2
                                        v3 += v3
                                        _ = v43
                                    } else {
                                        switch ("\n\t" + "hello") {
                                        case "a":
                                            v5 = (v5 ^ -1)
                                        default:
                                            var v44 string = ""
                                            _ = v44
                                        }
                                        for fi45 := 0; fi45 < 9; fi45++ {
                                            if (false || ("a" == "hello")) {
                                                v46 := GenS1{F0: ("\n\t" + "hello"), F1: "a"}
                                                if ((false && ((v26 - 618) != (v26 * uint(v2)))) && true) {
                                                    if (!(false || true)) {
                                                        for fi47 := 0; fi47 < 17; fi47++ {
                                                            fmt.Println(v26, v3, v3)
                                                            if true {
                                                                if false {
                                                                    if (!true) {
                                                                        v48 := GenS6{F0: v26, F1: (v5 * v5)}
                                                                        v1 = append(v1, fi47)
                                                                        _ = v48
                                                                    } else {
                                                                        if (!false) {
                                                                            v1 = append(v1, (fi47 ^ fi47))
                                                                        }
                                                                        var v49 int = 660
                                                                        v4 = v4
                                                                        _ = v49
                                                                    }
                                                                }
                                                                fmt.Println(fi45)
                                                            } else {
                                                                var v50 int16 = (v5 >> (((v26 << ((v26) & 63))) & 63))
                                                                v51 := GenS7{F0: v26, F1: 128, F2: "\n\t"}
                                                                fmt.Println(fi45)
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                            v1 = append(v1, 9223372036854775807)
                                                            _ = fi47
                                                        }
                                                    } else {
                                                        for fi52 := 0; fi52 < 5; fi52++ {
                                                            switch genFn0() {
                                                            case 0:
                                                                if len(v1) > 0 {
                                                                    v1[0] = genFn0()
                                                                }
                                                            case 4294967295:
                                                                var v53 bool = false
                                                                _ = v53
                                                            default:
                                                                var v54 uint = 191
                                                                _ = v54
                                                            }
                                                            _ = fi52
                                                        }
                                                    }
                                                    if true {
                                                        v26 -= v26
                                                        switch "\n\t" {
                                                        case "a":
                                                            v4 = (v4 + ((v4 * v4) - 1.5))
                                                        }
                                                        if ((v3 >> ((4294967295) & 63)) > v3) {
                                                            for fi55 := 0; fi55 < 15; fi55++ {
                                                                v3 = v3
                                                                var v56 int64 = 9223372036854775807
                                                                _ = fi55
                                                                _ = v56
                                                            }
                                                            fmt.Println(v9, v4)
                                                        } else {
                                                            for fi57 := 0; fi57 < 5; fi57++ {
                                                                v58 := GenS8{F0: -2147483648, F1: 4294967295, F2: v2}
                                                                _ = fi57
                                                                _ = v58
                                                            }
                                                            v9 = float64((((v4 + v4) - ((float32((3.14159 + ((0.1 + (0.0 + v9)) + 1.5))) * 1.5) - v4)) - v4))
                                                            var v59 uint = 42
                                                            _ = v59
                                                        }
                                                    } else {
                                                        if (true && (int8(v26) != -1)) {
                                                            v26 = (v26 & v26)
                                                            switch (^(fi45)) {
                                                            case 32767:
                                                                v9 = (3.14159 * 1.5)
                                                            case 0:
                                                                v9 = v9
                                                            }
                                                            v5 = v5
                                                        }
                                                        var v60 float64 = (1.5 + (v9 * 3.14159))
                                                        v3 = (v3 << ((v26) & 63))
                                                        _ = v60
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = fi45
                                                    }
                                                } else {
                                                    for fi61 := 0; fi61 < 7; fi61++ {
                                                        if ((v26 > 2) && false) {
                                                            if false {
                                                                if (int8(v26) != 127) {
                                                                    if (false && false) {
                                                                        var v62 int8 = (v3 ^ (^(v3)))
                                                                        _ = v62
                                                                    }
                                                                } else {
                                                                    if (false && false) {
                                                                        v63 := GenS5{F0: 0, F1: 2}
                                                                        _ = len(v1)
                                                                        v2 -= (^(v2))
                                                                        _ = v63
                                                                    }
                                                                    var v64 int32 = 0
                                                                    _ = v64
                                                                }
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            switch (fi45 & (fi45 << ((v26) & 63))) {
                                                            case -617:
                                                                v65 := GenS9{F0: 1, F1: 65535, F2: true}
                                                                _ = v65
                                                            case -1:
                                                                _ = len(v1)
                                                            case 32767:
                                                                var v66 uint64 = 116
                                                                _ = v66
                                                            }
                                                        }
                                                        fmt.Println(v3, v9)
                                                        if (fi61 >= fi61) {
                                                            switch "hello" {
                                                            case "hello":
                                                                switch ("" + "\n\t") {
                                                                case "\n\t":
                                                                    if len(v1) > 0 {
                                                                        v1[0] = genFn0()
                                                                    }
                                                                case "foo bar":
                                                                    if (int64(v3) == -1) {
                                                                        v1 = append(v1, genFn0())
                                                                    } else {
                                                                        v1 = append(v1, fi61)
                                                                    }
                                                                }
                                                            case "":
                                                                v4 -= -2.25
                                                            case "\n\t":
                                                                v67 := GenS10{F0: (v26 - 42), F1: fi61, F2: v3}
                                                                _ = v67
                                                            }
                                                            var v68 int16 = v5
                                                            _ = v68
                                                        } else {
                                                            switch "foo bar" {
                                                            case "":
                                                                fmt.Println(v26, v5)
                                                            case "a":
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    if ((!false) && (((0 >= -2147483648) || true) && false)) {
                                                                        switch ("a" + "a") {
                                                                        case "foo bar":
                                                                            switch fi45 {
                                                                            case -154:
                                                                                v69 := GenS11{F0: 561}
                                                                                _ = v69
                                                                            case 4294967295:
                                                                                fmt.Println(v26)
                                                                            default:
                                                                                _ = len(v1)
                                                                            }
                                                                        case "hello":
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                        var v70 int64 = -9223372036854775808
                                                                        v9 += 0.1
                                                                        _ = v70
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                v71 := GenS9{F0: 2147483648, F1: 32767, F2: true}
                                                                _ = v71
                                                            }
                                                        }
                                                        _ = fi61
                                                    }
                                                    for fi72 := 0; fi72 < 2; fi72++ {
                                                        var v73 uint64 = uint64(v26)
                                                        v73 = v73
                                                        v74 := GenS12{F0: 1, F1: -9223372036854775808}
                                                        _ = fi72
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                }
                                                _ = v46
                                            } else {
                                                v2 = (v2 % ((v2) | 1))
                                                v3 += v3
                                                if (true && true) {
                                                    v9 = 0.1
                                                    if (v2 <= (v2 / ((v2) | 1))) {
                                                        _ = len(v1)
                                                        switch "a" {
                                                        case "":
                                                            fmt.Println(v3, v4)
                                                        }
                                                    } else {
                                                        v75 := GenS9{F0: v2, F1: 882, F2: true}
                                                        _ = v75
                                                    }
                                                    var v76 string = "hello"
                                                    _ = v76
                                                }
                                            }
                                            _ = fi45
                                        }
                                    }
                                case "\n\t":
                                    if ((v26 & v26) > v26) {
                                        v4 -= v4
                                        v26 -= (v26 * v26)
                                        var v77 uint16 = 65535
                                        _ = v77
                                    }
                                case "hello":
                                    v78 := GenS9{F0: (v2 & v2), F1: 1, F2: (255 != 1)}
                                    _ = v78
                                default:
                                    switch 731 {
                                    case 2147483647:
                                        v79 := GenS6{F0: 1, F1: v5}
                                        _ = v79
                                    case 1:
                                        v80 := GenS13{F0: 52}
                                        _ = v80
                                    default:
                                        v2 = 730
                                    }
                                }
                            case 9223372036854775807:
                                if len(v1) > 0 {
                                    v1[0] = genFn0()
                                }
                            }
                            var v81 string = "a"
                            _ = v81
                        }
                    }
                }
            }
            _ = v26
        } else {
            switch "a" {
            case "hello":
                if len(v1) > 0 {
                    v1[0] = 32767
                }
            case "":
                if len(v1) > 0 {
                    v1[0] = -813
                }
            case "foo bar":
                if false {
                    v1 = append(v1, 128)
                    var v82 int16 = (v5 << ((2) & 63))
                    _ = v82
                }
            default:
                v1 = append(v1, -2147483648)
            }
        }
        if (genFn0() >= genFn0()) {
            v83 := GenS11{F0: 42}
            _ = v83
        } else {
            if true {
                var v84 float64 = (((v9 - float64(v4)) * (3.14159 / (v9 - v9))) / v9)
                _ = v84
            }
        }
        var v85 uint64 = 187
        _ = v85
    } else {
        var v86 uint64 = 2
        _ = v86
    }
    var v87 bool = (!((v5 - v5) < (v5 - 42)))
    for fi88 := 0; fi88 < 14; fi88++ {
        var v89 string = ("" + (("hello" + "hello") + ("a" + "foo bar")))
        switch "" {
        case "a":
            for fi90 := 0; fi90 < 18; fi90++ {
                fmt.Println(fi88, v4)
                switch fi90 {
                case 127:
                    v91 := GenS14{F0: (v2 / ((v2) | 1))}
                    _ = v91
                case 4294967295:
                    v9 -= v9
                case -1:
                    switch -519 {
                    case 0:
                        var v92 int = -9223372036854775808
                        _ = v92
                    case 162:
                        var v93 float64 = ((float64((0.0 + (v4 / 3.14159))) + (1e10 - v9)) - (v9 - (v9 + 0.1)))
                        fmt.Println(v89, v93, v3)
                        var v94 uint8 = 0
                        _ = v93
                        _ = v94
                    case 32768:
                        v2 += (v2 + (^(v2)))
                    }
                default:
                    var v95 int64 = 9223372036854775807
                    _ = v95
                }
                var v96 uint8 = 182
                _ = fi90
                _ = v96
            }
        case "\n\t":
            v1 = append(v1, genFn0())
        case "":
            var v97 bool = v87
            _ = v97
        }
        _ = len(v1)
        _ = fi88
        _ = v89
    }
    if len(v1) > 0 {
        v1[0] = genFn0()
    }
    var v98 uint32 = (v2 * v2)
    v4 = (v4 + 0.0)
    var v99 bool = v87
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v87:", v87)
    fmt.Println("v9:", v9)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v87
    _ = v9
    _ = v98
    _ = v99
}
