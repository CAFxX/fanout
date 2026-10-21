package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 uint32 }
type GenS7 struct { F0 int; F1 int64; F2 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool; F1 uint; F2 uint }
type GenS10 struct { F0 uint; F1 uint8 }
type GenS11 struct { F0 int64 }

func main() {
    var v1 uint16 = 794
    var v2 int32 = 42
    var v3 int16 = -1
    var v4 string = "hello"
    v5 := []int{-1, -348, -704, -9223372036854775808}
    v6 := []string{}
    if false {
        if len(v6) > 0 {
            v6[0] = ("foo bar" + "hello")
        }
        var v7 uint8 = 1
        _ = v7
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            fmt.Println(v4, v4)
            v4 = v4
            _ = fi8
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            for fi10 := 0; fi10 < 18; fi10++ {
                if (199 > 255) {
                    var v11 int = fi10
                    _ = len(v5)
                    _ = v11
                } else {
                    var v12 uint8 = 1
                    _ = v12
                }
                v4 = v4
                _ = fi10
            }
            var v13 uint16 = v1
            var v14 int = int(v3)
            _ = fi9
            _ = v13
            _ = v14
        }
        if (float64((3.14159 - (1e10 - (1.5 * 0.0)))) < 0.0) {
            if (100 >= 127) {
                for fi15 := 0; fi15 < 5; fi15++ {
                    v4 = v4
                    _ = fi15
                }
                var v16 int64 = -9223372036854775808
                _ = v16
            }
        }
    }
    switch -824 {
    case 32767:
        v4 = v4
    case 2147483647:
        fmt.Println(v3, v3)
    case -1:
        fmt.Println(v1, v3, v4)
    default:
        v17 := GenS1{F0: -1, F1: 236, F2: 139}
        _ = v17
    }
    if true {
        var v18 int8 = 42
        v19 := GenS1{F0: -907, F1: 128, F2: 18446744073709551615}
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: -1, F1: 1, F2: 0}
    if ((true && (!(v4 > (((v4 + ("" + v4)) + "\n\t") + (v4 + v4))))) || false) {
        switch 9223372036854775807 {
        case 0:
            if false {
                for fi21 := 0; fi21 < 20; fi21++ {
                    v3 -= -32768
                    for fi22 := 0; fi22 < 20; fi22++ {
                        if (0 <= 42) {
                            v2 = v2
                            var v23 bool = false
                            _ = v23
                        }
                        _ = fi22
                    }
                    _ = fi21
                }
            } else {
                for fi24 := 0; fi24 < 20; fi24++ {
                    if len(v5) > 0 {
                        v5[0] = int(v1)
                    }
                    for fi25 := 0; fi25 < 8; fi25++ {
                        v4 = (v4 + "hello")
                        var v26 bool = true
                        _ = fi25
                        _ = v26
                    }
                    _ = fi24
                }
                fmt.Println(v1)
            }
        }
        fmt.Println(v3, v1)
        switch (v4 + "a") {
        case "\n\t":
            v4 = v4
        default:
            v4 = v4
        }
    }
    if (false || false) {
        var v27 float64 = 0.1
        fmt.Println(v3, v1)
        _ = v27
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        if (3.14159 < float32(1.5)) {
            fmt.Println(v2, v1)
            var v29 bool = true
            if false {
                if v29 {
                    if (34 <= 9223372036854775807) {
                        if len(v6) > 0 {
                            v6[0] = (("\n\t" + v4) + "foo bar")
                        }
                    } else {
                        for fi30 := 0; fi30 < 4; fi30++ {
                            v31 := GenS1{F0: -1, F1: 42, F2: 3}
                            _ = fi30
                            _ = v31
                        }
                        if (127 != -8) {
                            var v32 uint8 = 238
                            for fi33 := 0; fi33 < 19; fi33++ {
                                if len(v6) > 0 {
                                    v6[0] = v4
                                }
                                v34 := GenS2{F0: v1, F1: fi33, F2: (v1 ^ 620)}
                                v29 = true
                                _ = fi33
                                _ = v34
                            }
                            v35 := GenS3{F0: v1, F1: v32}
                            _ = v32
                            _ = v35
                        }
                        var v36 uint16 = (v1 | 158)
                        _ = v36
                    }
                    v3 -= v3
                } else {
                    if len(v6) > 0 {
                        v6[0] = v4
                    }
                    _ = len(v5)
                    switch fi28 {
                    case 42:
                        for fi37 := 0; fi37 < 12; fi37++ {
                            v1 = v1
                            var v38 float32 = float32(((3.14159 / 0.1) - 0.0))
                            if true {
                                var v39 int64 = -443
                                if v29 {
                                    v38 -= (1.5 - float32(3.14159))
                                    var v40 string = "\n\t"
                                    _ = v40
                                }
                                _ = v39
                            }
                            fmt.Println(v3, v2, v1)
                            v4 = v4
                            _ = fi37
                            _ = v38
                        }
                    }
                }
                v3 = 1
                var v41 uint16 = (v1 * 0)
                _ = v41
            }
            _ = v29
        }
        fmt.Println(v2, v1)
        _ = fi28
    }
    switch "\n\t" {
    case "":
        if (1e10 < 3.14159) {
            if false {
                for fi42 := 0; fi42 < 12; fi42++ {
                    fmt.Println(v4, v3, v3)
                    var v43 uint32 = 42
                    var v44 uint = uint(v43)
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                if (v2 <= 2147483647) {
                    v5 = append(v5, -1)
                } else {
                    fmt.Println(v1, v1)
                    if (196 != 1) {
                        v1 = v1
                        for fi45 := 0; fi45 < 4; fi45++ {
                            v46 := GenS2{F0: (v1 ^ v1), F1: -9223372036854775808, F2: (v1 * (v1 + (v1 * (v1 ^ 65535))))}
                            var v47 int32 = 1
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                        fmt.Println(v3, v2, v3)
                    } else {
                        var v48 uint64 = 639
                        fmt.Println(v1)
                        v1 -= v1
                        _ = v48
                    }
                }
            } else {
                var v49 string = ""
                _ = v49
            }
            var v50 uint = 0
            fmt.Println(v50, v4, v3)
            _ = v50
        }
    case "\n\t":
        for fi51 := 0; fi51 < 7; fi51++ {
            var v52 int8 = 78
            _ = fi51
            _ = v52
        }
    }
    switch "\n\t" {
    case "\n\t":
        v5 = append(v5, -247)
    case "a":
        switch "foo bar" {
        case "":
            v4 = "hello"
        default:
            var v53 uint64 = 271
            _ = v53
        }
    }
    v54 := GenS4{F0: true}
    if len(v5) > 0 {
        v5[0] = 32767
    }
    if true {
        if ((v2 >> ((2) & 63)) != (v2 & -1)) {
            for fi55 := 0; fi55 < 8; fi55++ {
                v2 -= 0
                if (0 > 2) {
                    v1 = 65535
                    for fi56 := 0; fi56 < 16; fi56++ {
                        var v57 uint16 = (v1 % (((v1 | (v1 | v1))) | 1))
                        fmt.Println(v1, v4, v1)
                        _ = fi56
                        _ = v57
                    }
                }
                v2 = -65
                _ = fi55
            }
            v5 = append(v5, 32767)
            if true {
                for fi58 := 0; fi58 < 6; fi58++ {
                    v3 = (v3 & (v3 ^ (v3 & v3)))
                    v59 := GenS4{F0: (false || (true || true))}
                    _ = fi58
                    _ = v59
                }
            } else {
                v60 := GenS5{F0: 383}
                var v61 bool = (!true)
                v62 := GenS6{F0: v2, F1: 0}
                _ = v60
                _ = v61
                _ = v62
            }
        }
        if (2 != 42) {
            v2 = (v2 / ((v2) | 1))
            if false {
                v1 -= 0
                v1 -= v1
                v63 := GenS1{F0: 823, F1: 201, F2: 1}
                _ = v63
            } else {
                for fi64 := 0; fi64 < 16; fi64++ {
                    var v65 int64 = 42
                    v66 := GenS7{F0: (fi64 - fi64), F1: (v65 << ((979) & 63)), F2: 464}
                    switch v4 {
                    case "hello":
                        var v67 int16 = (v3 * v3)
                        _ = v67
                    case "":
                        for fi68 := 0; fi68 < 4; fi68++ {
                            fmt.Println(v3)
                            _ = fi68
                        }
                    case "\n\t":
                        for fi69 := 0; fi69 < 7; fi69++ {
                            if ((!(true || true)) && (-1 >= 2147483647)) {
                                var v70 int32 = v2
                                _ = v70
                            }
                            var v71 uint8 = 128
                            var v72 bool = (!("\n\t" >= v4))
                            _ = fi69
                            _ = v71
                            _ = v72
                        }
                    }
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                v1 += v1
                var v73 int32 = v2
                _ = v73
            }
        }
        var v74 uint8 = 2
        _ = v74
    }
    v1 = (v1 & 329)
    for fi75 := 0; fi75 < 5; fi75++ {
        v1 = v1
        var v76 float64 = (3.14159 / 1e10)
        _ = fi75
        _ = v76
    }
    for fi77 := 0; fi77 < 18; fi77++ {
        switch "" {
        case "hello":
            switch (fi77 * (fi77 + 4294967295)) {
            case 4294967295:
                v2 = v2
            }
        case "foo bar":
            v78 := GenS8{F0: -80}
            _ = v78
        case "a":
            var v79 bool = (0 == (fi77 | (fi77 / ((-574) | 1))))
            _ = v79
        }
        for fi80 := 0; fi80 < 7; fi80++ {
            var v81 int16 = (v3 ^ (v3 - v3))
            var v82 uint = 1
            if false {
                switch (fi80 - (fi80 * (fi80 % (((fi80 | (^(fi80)))) | 1)))) {
                case 42:
                    var v83 int32 = v2
                    v4 = ("a" + "")
                    _ = v83
                case 0:
                    fmt.Println(v82)
                case 128:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    switch (fi80 - fi80) {
                    case 358:
                        v5 = append(v5, (fi77 * fi77))
                    case 9223372036854775807:
                        var v84 float32 = -2.25
                        _ = v84
                    case 728:
                        if (true && (!true)) {
                            v85 := GenS9{F0: true, F1: v82, F2: v82}
                            _ = v85
                        } else {
                            var v86 int8 = -1
                            _ = v86
                        }
                    }
                    var v87 string = v4
                    _ = v87
                }
                switch 458 {
                case -2147483648:
                    if len(v5) > 0 {
                        v5[0] = (fi80 | fi77)
                    }
                case 1:
                    v88 := GenS10{F0: v82, F1: 205}
                    _ = v88
                default:
                    if len(v5) > 0 {
                        v5[0] = (fi80 & (fi77 | (fi77 - (fi77 & (fi80 ^ -807)))))
                    }
                    if false {
                        fmt.Println(v1, v2)
                        if len(v6) > 0 {
                            v6[0] = v4
                        }
                    }
                    switch v4 {
                    case "":
                        var v89 bool = false
                        _ = v89
                    case "hello":
                        var v90 float32 = 0.0
                        _ = v90
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if false {
                    if (255 != (v1 & v1)) {
                        if len(v6) > 0 {
                            v6[0] = "hello"
                        }
                        if len(v5) > 0 {
                            v5[0] = fi80
                        }
                    } else {
                        if false {
                            var v91 int16 = v81
                            fmt.Println(v4)
                            v82 -= v82
                            _ = v91
                        } else {
                            v6 = append(v6, v4)
                        }
                        v2 += 2147483647
                        v4 = "hello"
                    }
                    var v92 int32 = v2
                    v81 -= (-(v81))
                    if true {
                        v1 += 32767
                    } else {
                        v2 = 373
                    }
                    fmt.Println(v82, v2)
                    v2 -= (v92 >> (((v82 - (v82 + 2))) & 63))
                    _ = v92
                } else {
                    if len(v6) > 0 {
                        v6[0] = v4
                    }
                    fmt.Println(v82)
                }
                switch (v4 + (v4 + v4)) {
                case "foo bar":
                    if true {
                        fmt.Println(v81)
                        switch v4 {
                        case "hello":
                            v1 += v1
                        default:
                            switch v4 {
                            case "\n\t":
                                v81 = v3
                            }
                        }
                        v5 = append(v5, (fi77 << ((v82) & 63)))
                    } else {
                        if (true && false) {
                            var v93 float64 = (1.5 / ((1.5 + (0.1 / 1e10)) * 0.1))
                            if (v4 >= ((v4 + ((v4 + v4) + ((v4 + "foo bar") + v4))) + v4)) {
                                v94 := GenS7{F0: (fi80 + fi80), F1: -1, F2: 959}
                                v4 = v4
                                fmt.Println(fi77, fi80, fi77)
                                _ = v94
                            } else {
                                if (uint64(v1) <= 2) {
                                    switch (^(fi80)) {
                                    case 433:
                                        if (("a" > (v4 + (v4 + v4))) || (-32768 < -1)) {
                                            v93 = (v93 * v93)
                                        }
                                    default:
                                        fmt.Println(v81, v4)
                                    }
                                    v6 = append(v6, v4)
                                } else {
                                    switch "hello" {
                                    case "":
                                        v3 += v81
                                    case "a":
                                        fmt.Println(v1, fi80)
                                    }
                                    v95 := GenS11{F0: 571}
                                    var v96 int32 = (v2 ^ (v2 | 1))
                                    _ = v95
                                    _ = v96
                                }
                                fmt.Println(fi77, fi77)
                                switch (fi80 << (((v82 >> (((v82 * (v82 - v82))) & 63))) & 63)) {
                                case 0:
                                    if (!false) {
                                        fmt.Println(fi77, v3, v1)
                                        if false {
                                            var v97 int64 = 42
                                            var v98 int16 = (v3 | (v81 % ((42) | 1)))
                                            _ = v97
                                            _ = v98
                                        }
                                    }
                                }
                            }
                            var v99 float64 = (v93 / ((float64(1e10) - v93) * 1.5))
                            if len(v6) > 0 {
                                v6[0] = (v4 + v4)
                            }
                            switch (fi77 % (((-(fi77))) | 1)) {
                            case 0:
                                switch ((("\n\t" + v4) + v4) + "a") {
                                case "a":
                                    if len(v5) > 0 {
                                        v5[0] = fi80
                                    }
                                }
                                v2 = (v2 % ((v2) | 1))
                                fmt.Println(v99)
                            case -120:
                                v82 = v82
                            }
                            _ = v93
                            _ = v99
                        } else {
                            var v100 string = v4
                            _ = v100
                        }
                        v101 := GenS8{F0: -84}
                        if len(v6) > 0 {
                            v6[0] = "hello"
                        }
                        var v102 int16 = (v81 << (((^(v82))) & 63))
                        _ = v101
                        _ = v102
                    }
                case "":
                    switch "\n\t" {
                    case "":
                        var v103 string = v4
                        _ = v103
                    }
                case "hello":
                    fmt.Println(fi80)
                    v5 = append(v5, fi80)
                }
            }
            _ = fi80
            _ = v81
            _ = v82
        }
        _ = fi77
    }
    v3 = v3
    v3 = 746
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
