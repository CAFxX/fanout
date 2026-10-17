package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int; F1 uint32; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 bool; F1 int16 }
func genFn0() bool {
    if false {
        v7 := []int{4294967295}
        _ = v7
    }
    var v8 int = 1
    switch "" {
    case "hello":
        v9 := []int{v8}
        _ = v9
    case "foo bar":
        if false {
            var v10 int8 = 0
            v11 := GenS1{F0: 18446744073709551615, F1: (true || true)}
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 uint8 = 2
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 uint16 = 42
                    if false {
                        if false {
                            var v16 uint16 = (v15 >> ((233) & 63))
                            _ = v16
                        }
                        var v17 int16 = 42
                        v18 := GenS2{F0: -681, F1: 42, F2: 65535}
                        _ = v17
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                }
                v19 := GenS3{F0: ("" != "a")}
                _ = fi12
                _ = v13
                _ = v19
            }
            v8 += (v8 << ((164) & 63))
            v8 = v8
        }
    case "\n\t":
        for fi20 := 0; fi20 < 7; fi20++ {
            if true {
                v8 += (v8 / ((1) | 1))
                var v21 uint8 = 151
                _ = v21
            } else {
                v8 += fi20
                fmt.Println(v8)
                var v22 int = (fi20 % ((v8) | 1))
                _ = v22
            }
            _ = fi20
        }
    default:
        if false {
            v8 -= (v8 * 127)
            v8 += v8
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 int16 = -32768
                var v25 int8 = -9
                if false {
                    if (!("\n\t" >= "")) {
                        v8 = (v8 | (v8 % (((v8 - 856)) | 1)))
                        v26 := GenS4{F0: true, F1: 2147483647}
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 9; fi27++ {
                        v24 += (v24 & v24)
                        switch "\n\t" {
                        case "":
                            v25 = (v25 >> ((133) & 63))
                        }
                        _ = fi27
                    }
                    var v28 float64 = 1e10
                    _ = v28
                } else {
                    var v29 uint = 255
                    if (v25 <= (-(v25))) {
                        for fi30 := 0; fi30 < 2; fi30++ {
                            v31 := GenS1{F0: 2, F1: (float64(3.14159) != (3.14159 / 1.0))}
                            _ = fi30
                            _ = v31
                        }
                        var v32 uint16 = 1
                        _ = v32
                    }
                    _ = v29
                }
                _ = fi23
                _ = v24
                _ = v25
            }
        } else {
            v8 += 9223372036854775807
            v33 := []int{48, v8, 32767}
            _ = v33
        }
    }
    v34 := GenS1{F0: 0, F1: (9223372036854775807 == -843)}
    for fi35 := 0; fi35 < 18; fi35++ {
        for fi36 := 0; fi36 < 18; fi36++ {
            var v37 uint = 1
            _ = fi36
            _ = v37
        }
        if false {
            v8 -= v8
            v8 += v8
            v8 += -334
        }
        v8 += v8
        _ = fi35
    }
    _ = v34
    _ = v8
    return ((v8 ^ v8) > v8)
}

func main() {
    var v1 int32 = -1
    var v2 int8 = -4
    var v3 uint32 = 2147483648
    v4 := []string{}
    var v5 bool = false
    var v6 int32 = v1
    var v38 int64 = 42
    for fi39 := 0; fi39 < 13; fi39++ {
        fmt.Println(v2)
        var v40 string = ""
        _ = fi39
        _ = v40
    }
    var v41 uint = 0
    v3 = v3
    var v42 bool = v5
    if (-9223372036854775808 > (v38 ^ (v38 - (v38 | v38)))) {
        fmt.Println(v42)
        if true {
            switch "hello" {
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "hello":
                if true {
                    switch ("" + ("" + "hello")) {
                    case "foo bar":
                        v42 = (!v5)
                    default:
                        v41 = (^(v41))
                    }
                    if v5 {
                        v43 := GenS5{F0: false, F1: 2}
                        var v44 string = ("" + "hello")
                        _ = v43
                        _ = v44
                    }
                    var v45 int32 = int32(v3)
                    _ = v45
                } else {
                    switch "hello" {
                    case "\n\t":
                        for fi46 := 0; fi46 < 19; fi46++ {
                            v38 = v38
                            switch fi46 {
                            case -17:
                                v47 := GenS6{F0: 42}
                                _ = v47
                            case 32768:
                                _ = len(v4)
                            case -2147483648:
                                fmt.Println(v1, v42)
                            }
                            var v48 uint16 = 32767
                            _ = fi46
                            _ = v48
                        }
                    }
                    v1 = v1
                    v38 = (v38 * (v38 % ((v38) | 1)))
                }
            default:
                if (-227 != 42) {
                    v6 -= (v6 / ((0) | 1))
                    var v49 uint = v41
                    _ = len(v4)
                    _ = v49
                } else {
                    for fi50 := 0; fi50 < 20; fi50++ {
                        if false {
                            v1 -= 42
                            if ((v1 * -1) < v1) {
                                for fi51 := 0; fi51 < 1; fi51++ {
                                    var v52 int32 = v6
                                    _ = fi51
                                    _ = v52
                                }
                                var v53 int64 = -1
                                _ = v53
                            }
                        }
                        _ = fi50
                    }
                }
            }
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            switch "" {
            case "\n\t":
                fmt.Println(v6)
            }
        }
    } else {
        var v54 float32 = -2.25
        switch "a" {
        case "a":
            if false {
                v5 = (209 != -721)
                if len(v4) > 0 {
                    v4[0] = ("a" + "\n\t")
                }
                v55 := GenS1{F0: 0, F1: v5}
                _ = v55
            } else {
                switch 42 {
                case 42:
                    v3 += (^(v3))
                case -1:
                    if (1 != 513) {
                        _ = len(v4)
                        if true {
                            for fi56 := 0; fi56 < 13; fi56++ {
                                for fi57 := 0; fi57 < 7; fi57++ {
                                    fmt.Println(v3, v1)
                                    _ = fi57
                                }
                                switch (fi56 - (fi56 & fi56)) {
                                case 42:
                                    var v58 float32 = (float32(0.1) - (1.5 / v54))
                                    _ = v58
                                case 9223372036854775807:
                                    if ((^(v2)) >= (^(v2))) {
                                        var v59 float32 = v54
                                        _ = v59
                                    } else {
                                        switch -9223372036854775808 {
                                        case 4294967295:
                                            if ((18446744073709551615 == 255) && (v42 || ((v3 / ((10) | 1)) == 564))) {
                                                if true {
                                                    v5 = v42
                                                    var v60 bool = v42
                                                    _ = v60
                                                } else {
                                                    v61 := GenS7{F0: -1, F1: (v6 - 42)}
                                                    _ = v61
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                switch "a" {
                                                case "a":
                                                    v41 = v41
                                                case "foo bar":
                                                    v62 := GenS8{F0: 266, F1: 641}
                                                    _ = v62
                                                case "":
                                                    if false {
                                                        switch "a" {
                                                        case "\n\t":
                                                            fmt.Println(fi56)
                                                        }
                                                    }
                                                default:
                                                    var v63 int = fi56
                                                    _ = v63
                                                }
                                            }
                                        case -311:
                                            v2 = v2
                                        }
                                        switch 32768 {
                                        case 0:
                                            var v64 string = "\n\t"
                                            _ = v64
                                        }
                                    }
                                }
                                for fi65 := 0; fi65 < 17; fi65++ {
                                    switch "\n\t" {
                                    case "a":
                                        if ("" <= "") {
                                            v66 := GenS9{F0: v5, F1: -1}
                                            _ = v66
                                        }
                                    case "hello":
                                        fmt.Println(v38)
                                    default:
                                        v6 += v1
                                    }
                                    switch ((("a" + "\n\t") + "foo bar") + "a") {
                                    case "\n\t":
                                        _ = len(v4)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    default:
                                        v6 = v6
                                    }
                                    v2 += 0
                                    _ = fi65
                                }
                                _ = fi56
                            }
                            var v67 float32 = (v54 * v54)
                            var v68 float64 = float64(v54)
                            _ = v67
                            _ = v68
                        }
                        for fi69 := 0; fi69 < 17; fi69++ {
                            fmt.Println(v54, v3, v42)
                            for fi70 := 0; fi70 < 14; fi70++ {
                                if true {
                                    var v71 uint64 = 105
                                    if len(v4) > 0 {
                                        v4[0] = "a"
                                    }
                                    _ = v71
                                } else {
                                    if (v5 && v42) {
                                        switch "" {
                                        case "foo bar":
                                            fmt.Println(v6, fi70, v41)
                                        case "\n\t":
                                            if (false && v42) {
                                                var v72 int64 = v38
                                                switch fi70 {
                                                case -216:
                                                    v3 += 2147483647
                                                case 2147483647:
                                                    if len(v4) > 0 {
                                                        v4[0] = ("a" + ("\n\t" + "\n\t"))
                                                    }
                                                default:
                                                    var v73 int64 = (v72 | (v72 % (((-(v38))) | 1)))
                                                    _ = v73
                                                }
                                                v4 = append(v4, "a")
                                                _ = v72
                                            } else {
                                                fmt.Println(fi70, v3, v54)
                                                var v74 int = fi69
                                                _ = v74
                                            }
                                        }
                                    }
                                    switch (-(fi69)) {
                                    case -2147483648:
                                        var v75 int16 = -1
                                        _ = v75
                                    case 9223372036854775807:
                                        if len(v4) > 0 {
                                            v4[0] = "a"
                                        }
                                    case 127:
                                        switch "" {
                                        case "foo bar":
                                            var v76 int16 = -1
                                            _ = v76
                                        }
                                    default:
                                        if (!v5) {
                                            var v77 uint32 = 728
                                            if v42 {
                                                v41 += 2
                                                fmt.Println(v1, v41)
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = ("" + "\n\t")
                                                }
                                            }
                                            _ = v77
                                        }
                                    }
                                }
                                var v78 uint32 = v3
                                _ = fi70
                                _ = v78
                            }
                            _ = fi69
                        }
                    }
                case 9223372036854775807:
                    for fi79 := 0; fi79 < 10; fi79++ {
                        if ((v38 >> (((v41 / ((v41) | 1))) & 63)) >= (^(v38))) {
                            var v80 float32 = (v54 * (0.1 + 1.5))
                            var v81 float64 = 3.14159
                            _ = v80
                            _ = v81
                        }
                        _ = fi79
                    }
                default:
                    fmt.Println(v54, v2)
                }
            }
        default:
            fmt.Println(v38, v3)
        }
        _ = v54
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = ((("a" + "") + "\n\t") + "\n\t")
    }
    fmt.Println(v42, v42, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
