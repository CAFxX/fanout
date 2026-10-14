package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 string; F2 int64 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 uint64; F1 uint64 }
type GenS9 struct { F0 int8; F1 int }

func main() {
    v1 := []int{132, 0}
    var v2 uint64 = 0
    v3 := []uint32{0, 2}
    var v4 uint8 = 0
    var v5 int8 = 42
    switch 1 {
    case 0:
        v6 := GenS1{F0: -1}
        _ = v6
    case 32768:
        if ((v2 >= v2) && true) {
            v1 = append(v1, 32767)
        } else {
            v3 = append(v3, 1)
        }
    case -243:
        v7 := GenS2{F0: 793, F1: ("\n\t" + "\n\t"), F2: -491}
        _ = v7
    }
    switch -1 {
    case -1:
        v1 = append(v1, 1)
    case 127:
        var v8 string = "foo bar"
        _ = v8
    case 0:
        var v9 float32 = 1e10
        _ = v9
    default:
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS3{F0: 0, F1: -1}
            switch -9223372036854775808 {
            case 32767:
                _ = len(v3)
            case 127:
                fmt.Println(fi10, v4, v5)
            }
            v5 += 100
            _ = fi10
            _ = v11
        }
    }
    if (1 <= 0) {
        v2 += (v2 + v2)
        v1 = append(v1, 559)
    } else {
        if len(v3) > 0 {
            v3[0] = 386
        }
        if (4294967295 != 968) {
            for fi12 := 0; fi12 < 3; fi12++ {
                if ((!true) && (766 > 42)) {
                    if ((-69 != 9223372036854775807) || true) {
                        var v13 int8 = (-(v5))
                        fmt.Println(v2, v5, v4)
                        _ = v13
                    }
                    _ = len(v3)
                    var v14 int64 = -1
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        fmt.Println(fi15, v2)
                        v4 = 66
                        _ = fi15
                    }
                }
                var v16 uint64 = v2
                v17 := GenS4{F0: (v4 | 42), F1: 2}
                _ = fi12
                _ = v16
                _ = v17
            }
        } else {
            var v18 uint = 4294967295
            _ = v18
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 2; fi20++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v21 bool = (!(true || true))
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    }
    fmt.Println(v2, v4)
    v5 = -99
    var v22 int64 = 363
    if false {
        var v23 uint64 = 4294967295
        _ = len(v1)
        v5 = -120
        _ = v23
    } else {
        var v24 float64 = float64(float32(1e10))
        _ = v24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (v2 < v2) {
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 uint8 = 255
                if len(v1) > 0 {
                    v1[0] = fi25
                }
                v28 := GenS4{F0: v4, F1: (v4 ^ 0)}
                _ = fi26
                _ = v27
                _ = v28
            }
            var v29 int = fi25
            var v30 uint64 = uint64(v4)
            _ = v29
            _ = v30
        } else {
            var v31 float64 = ((3.14159 * 0.0) / 1.0)
            _ = v31
        }
        v4 -= v4
        _ = fi25
    }
    if true {
        for fi32 := 0; fi32 < 10; fi32++ {
            fmt.Println(v4, v4)
            v33 := GenS1{F0: 42}
            if (false && (("" >= "a") || false)) {
                switch (fi32 / (((^(fi32))) | 1)) {
                case 2147483647:
                    var v34 uint32 = 2147483648
                    _ = v34
                case -9223372036854775808:
                    v35 := GenS5{F0: v22, F1: 42}
                    _ = v35
                default:
                    var v36 int8 = (v5 << ((521) & 63))
                    _ = v36
                }
            } else {
                v5 = v5
                for fi37 := 0; fi37 < 13; fi37++ {
                    if (!(304 < fi37)) {
                        v2 -= 42
                        var v38 int64 = v22
                        if true {
                            if true {
                                v22 = v22
                            } else {
                                var v39 int = 42
                                v4 -= (v4 / ((127) | 1))
                                _ = v39
                            }
                            v40 := GenS6{F0: -2147483648, F1: 18446744073709551615}
                            v41 := GenS7{F0: 18446744073709551615, F1: true}
                            _ = v40
                            _ = v41
                        } else {
                            var v42 bool = ("foo bar" != "a")
                            var v43 int64 = (v38 + v22)
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                            _ = v42
                            _ = v43
                        }
                        fmt.Println(v5, v4)
                        _ = v38
                    } else {
                        var v44 uint = 579
                        _ = len(v3)
                        _ = v44
                    }
                    var v45 int64 = (v22 << ((855) & 63))
                    _ = fi37
                    _ = v45
                }
                var v46 int8 = 0
                _ = v46
            }
            _ = fi32
            _ = v33
        }
        fmt.Println(v22, v4)
    } else {
        for fi47 := 0; fi47 < 18; fi47++ {
            if (float64(3.14159) >= 1.5) {
                for fi48 := 0; fi48 < 16; fi48++ {
                    if false {
                        v22 = v22
                    } else {
                        v3 = append(v3, 42)
                        v49 := GenS2{F0: (v22 ^ (v22 * 9223372036854775807)), F1: "", F2: v22}
                        _ = v49
                    }
                    switch (fi47 - fi48) {
                    case -302:
                        v4 = v4
                    case -2147483648:
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            var v50 int32 = 2147483647
                            _ = v50
                        default:
                            v51 := GenS1{F0: 398}
                            _ = v51
                        }
                    case 32768:
                        if false {
                            v52 := GenS3{F0: 853, F1: (v5 & (-(v5)))}
                            v22 -= v22
                            _ = v52
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                    _ = fi48
                }
                if false {
                    fmt.Println(fi47)
                    var v53 uint8 = (v4 | v4)
                    _ = v53
                } else {
                    if false {
                        v54 := GenS7{F0: 956, F1: false}
                        for fi55 := 0; fi55 < 16; fi55++ {
                            switch (fi47 % ((-1) | 1)) {
                            case 9223372036854775807:
                                if true {
                                    v3 = append(v3, 849)
                                    v56 := GenS3{F0: 42, F1: -1}
                                    _ = v56
                                }
                            case 127:
                                if false {
                                    switch (^(fi47)) {
                                    case -2147483648:
                                        if ((0.1 - float32(0.1)) != ((0.0 / 1.0) - ((1.5 - float32((1.5 * 3.14159))) - (0.0 + (0.0 / -2.25))))) {
                                            v2 -= v2
                                            if false {
                                                if true {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v57 uint16 = 64
                                                        _ = v57
                                                    }
                                                }
                                                v58 := GenS5{F0: v22, F1: 2}
                                                fmt.Println(v2, v22)
                                                _ = v58
                                            }
                                            var v59 uint = 4294967295
                                            var v60 int16 = 0
                                            _ = v59
                                            _ = v60
                                        }
                                    }
                                    if (0.0 >= (-2.25 - (0.0 - (((-2.25 + -2.25) / 3.14159) + -2.25)))) {
                                        var v61 int16 = -1
                                        if false {
                                            var v62 uint32 = 2147483648
                                            v2 = v2
                                            _ = v62
                                        }
                                        _ = v61
                                    }
                                    var v63 uint8 = uint8(v2)
                                    _ = v63
                                }
                                var v64 uint32 = 0
                                _ = v64
                            case 2147483647:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                v1 = append(v1, fi55)
                            }
                            v65 := GenS7{F0: 456, F1: (-17 == v5)}
                            if (false || false) {
                                if true {
                                    var v66 uint = 18446744073709551615
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v66
                                }
                            }
                            _ = fi55
                            _ = v65
                        }
                        v5 = 0
                        _ = v54
                    }
                    var v67 int = fi47
                    v68 := GenS4{F0: 2, F1: 168}
                    _ = v67
                    _ = v68
                }
                v3 = append(v3, 4294967295)
            }
            for fi69 := 0; fi69 < 3; fi69++ {
                if false {
                    if true {
                        switch "\n\t" {
                        case "":
                            switch "a" {
                            case "foo bar":
                                switch fi69 {
                                case 32768:
                                    v5 += v5
                                case -9223372036854775808:
                                    if true {
                                        v4 -= 219
                                    } else {
                                        v5 += v5
                                        v2 -= uint64(v4)
                                        if (1.5 != 1.5) {
                                            fmt.Println(v2, v4)
                                            var v70 float64 = 1.5
                                            v70 -= v70
                                            _ = v70
                                        }
                                    }
                                default:
                                    var v71 int16 = -1
                                    _ = v71
                                }
                            }
                        default:
                            v72 := GenS4{F0: v4, F1: 128}
                            _ = v72
                        }
                        v2 = 42
                    } else {
                        v2 -= (v2 / ((1) | 1))
                        if ("foo bar" == "") {
                            fmt.Println(fi69)
                            v4 -= (^(v4))
                            var v73 float32 = 0.1
                            _ = v73
                        } else {
                            var v74 uint8 = v4
                            _ = v74
                        }
                        switch fi47 {
                        case 1:
                            fmt.Println(v4, v4)
                        case 42:
                            v75 := GenS8{F0: v2, F1: v2}
                            _ = v75
                        }
                    }
                    v76 := GenS9{F0: (v5 + v5), F1: fi47}
                    var v77 int64 = int64(v2)
                    _ = v76
                    _ = v77
                } else {
                    v22 = (v22 << ((65535) & 63))
                    if true {
                        var v78 float64 = -2.25
                        _ = v78
                    } else {
                        v4 = 0
                    }
                    switch ("a" + ("hello" + ("foo bar" + "\n\t"))) {
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = fi69
                        }
                    }
                }
                _ = fi69
            }
            if len(v1) > 0 {
                v1[0] = 0
            }
            _ = fi47
        }
        v79 := GenS2{F0: (v22 << ((29) & 63)), F1: (("foo bar" + ("" + ("\n\t" + (("foo bar" + "") + "\n\t")))) + (((("\n\t" + "\n\t") + ("\n\t" + "")) + "\n\t") + "hello")), F2: v22}
        v4 = (v4 / ((128) | 1))
        _ = v79
    }
    if (false && (true && false)) {
        if len(v3) > 0 {
            v3[0] = 0
        }
    }
    switch "a" {
    case "":
        for fi80 := 0; fi80 < 6; fi80++ {
            v4 = v4
            _ = fi80
        }
    case "hello":
        for fi81 := 0; fi81 < 9; fi81++ {
            v82 := GenS7{F0: uint(v2), F1: (0.1 > float32(-2.25))}
            v83 := GenS1{F0: 32767}
            _ = fi81
            _ = v82
            _ = v83
        }
    }
    v2 = 2
    for fi84 := 0; fi84 < 10; fi84++ {
        var v85 bool = ((false || (1e10 < 3.14159)) || false)
        _ = fi84
        _ = v85
    }
    if ((v2 % ((v2) | 1)) < v2) {
        v5 = -128
        var v86 int = -1
        _ = v86
    }
    fmt.Println(v22, v4)
    for fi87 := 0; fi87 < 1; fi87++ {
        v88 := GenS9{F0: (-(v5)), F1: fi87}
        _ = fi87
        _ = v88
    }
    for fi89 := 0; fi89 < 14; fi89++ {
        v2 -= v2
        switch fi89 {
        case 635:
            if (1 <= 127) {
                var v90 string = "a"
                var v91 string = ("a" + ("\n\t" + "foo bar"))
                fmt.Println(v91, v22)
                _ = v90
                _ = v91
            }
        case -9223372036854775808:
            var v92 int8 = v5
            _ = v92
        default:
            for fi93 := 0; fi93 < 12; fi93++ {
                if false {
                    v2 = (v2 - (^(v2)))
                    var v94 int16 = 1
                    if (1 != 42) {
                        fmt.Println(fi89, v4, v2)
                    }
                    _ = v94
                } else {
                    switch fi89 {
                    case 32767:
                        v3 = append(v3, 0)
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v95 := GenS3{F0: 65535, F1: (v5 * v5)}
                    _ = v95
                }
                _ = len(v1)
                v96 := GenS4{F0: v4, F1: (v4 | (^(v4)))}
                _ = fi93
                _ = v96
            }
        }
        _ = fi89
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
