package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint32; F1 uint8; F2 int }
type GenS8 struct { F0 uint8; F1 int16; F2 int }
type GenS9 struct { F0 int8; F1 uint }
func genFn0() int32 {
    var v5 string = "foo bar"
    var v6 uint8 = 214
    fmt.Println(v6)
    v7 := []int{32767, -274}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v5
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint = 65535
    v2 := []string{"", "hello"}
    var v3 uint32 = 2
    v4 := []uint32{(^(v3)), v3, (v3 | (v3 + v3))}
    switch 128 {
    case 9223372036854775807:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v8 string = "foo bar"
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(v1)
            v11 := GenS1{F0: v3, F1: 1, F2: -1}
            fmt.Println(fi10, fi10)
            _ = fi10
            _ = v11
        }
        var v12 uint8 = 237
        _ = fi9
        _ = v12
    }
    var v13 int8 = 100
    for fi14 := 0; fi14 < 18; fi14++ {
        switch 42 {
        case 32767:
            v3 = v3
        case -981:
            _ = len(v4)
        case -9223372036854775808:
            switch (fi14 & (fi14 * fi14)) {
            case 32768:
                if len(v2) > 0 {
                    v2[0] = (v8 + "\n\t")
                }
            case -9223372036854775808:
                var v15 uint64 = 4294967295
                _ = v15
            case -1:
                v13 += int8(fi14)
            }
        }
        switch (^(fi14)) {
        case -1:
            var v16 string = v8
            _ = v16
        case 884:
            var v17 uint = (^(v1))
            _ = v17
        case -766:
            var v18 float32 = float32(3.14159)
            _ = v18
        default:
            var v19 uint = (v1 ^ v1)
            _ = v19
        }
        _ = len(v4)
        _ = fi14
    }
    v20 := GenS1{F0: v3, F1: 65535, F2: 520}
    switch -9223372036854775808 {
    case 32767:
        for fi21 := 0; fi21 < 9; fi21++ {
            v1 += 18446744073709551615
            var v22 uint = v1
            v3 = v3
            _ = fi21
            _ = v22
        }
    case 697:
        if false {
            var v23 int64 = 42
            v4 = append(v4, 2147483648)
            _ = v23
        } else {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            fmt.Println(v3, v13, v13)
            for fi24 := 0; fi24 < 4; fi24++ {
                for fi25 := 0; fi25 < 17; fi25++ {
                    if (v1 > (v1 + 2)) {
                        var v26 uint32 = v3
                        if (false && false) {
                            if ((false && (false || false)) && (((1000 < 974) || (!(0.0 >= 1e10))) && true)) {
                                switch "hello" {
                                case "":
                                    switch ("hello" + "hello") {
                                    case "":
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                    default:
                                        var v27 int32 = genFn0()
                                        _ = v27
                                    }
                                }
                            }
                            if len(v2) > 0 {
                                v2[0] = v8
                            }
                            v8 = ""
                        } else {
                            var v28 uint16 = 255
                            v29 := GenS2{F0: (v26 >= 65535)}
                            _ = v28
                            _ = v29
                        }
                        _ = v26
                    }
                    v30 := GenS3{F0: 2}
                    var v31 uint8 = 0
                    _ = fi25
                    _ = v30
                    _ = v31
                }
                _ = fi24
            }
        }
    default:
        v8 = v8
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 17; fi33++ {
            fmt.Println(v13, v1)
            v34 := GenS4{F0: ((v8 + (v8 + v8)) + "foo bar")}
            if false {
                switch (fi32 & fi33) {
                case 127:
                    var v35 int32 = 1
                    _ = v35
                case -2147483648:
                    if ("foo bar" <= "") {
                        if true {
                            switch (v8 + (v8 + "")) {
                            case "a":
                                _ = len(v4)
                                var v36 uint16 = 221
                                _ = v36
                            case "foo bar":
                                v37 := GenS1{F0: (v3 ^ (^(v3))), F1: 473, F2: -897}
                                _ = v37
                            default:
                                v1 -= 42
                            }
                            if (42 != -1) {
                                var v38 int16 = 0
                                var v39 uint64 = 1
                                _ = v38
                                _ = v39
                            }
                        } else {
                            v13 = (v13 + 127)
                        }
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                    } else {
                        _ = len(v4)
                        if true {
                            var v40 bool = false
                            v1 -= v1
                            var v41 uint32 = v3
                            _ = v40
                            _ = v41
                        }
                    }
                default:
                    var v42 uint = v1
                    _ = v42
                }
                var v43 uint16 = 998
                _ = v43
            }
            _ = fi33
            _ = v34
        }
        if false {
            v44 := GenS2{F0: (!(!true))}
            var v45 int16 = -1
            _ = v44
            _ = v45
        } else {
            v46 := GenS3{F0: 220}
            _ = v46
        }
        var v47 int8 = v13
        _ = fi32
        _ = v47
    }
    var v48 float64 = ((3.14159 / 1e10) * 0.0)
    switch 32767 {
    case 827:
        for fi49 := 0; fi49 < 19; fi49++ {
            for fi50 := 0; fi50 < 4; fi50++ {
                if true {
                    v13 = (v13 - v13)
                }
                if ((v1 | (^(v1))) <= (v1 - 195)) {
                    fmt.Println(v13)
                    if false {
                        v3 = v3
                        v51 := GenS2{F0: (1000 > -234)}
                        var v52 uint32 = (v3 & 1)
                        _ = v51
                        _ = v52
                    } else {
                        if len(v4) > 0 {
                            v4[0] = (v3 - v3)
                        }
                        fmt.Println(fi49, fi49, fi50)
                    }
                }
                switch (fi50 >> ((42) & 63)) {
                case -1:
                    v2 = append(v2, v8)
                default:
                    v1 -= v1
                }
                v53 := GenS1{F0: 448, F1: 1, F2: 1000}
                _ = fi50
                _ = v53
            }
            var v54 int64 = 509
            _ = fi49
            _ = v54
        }
    case -9223372036854775808:
        if false {
            if false {
                v3 += v3
                for fi55 := 0; fi55 < 3; fi55++ {
                    var v56 int32 = -1
                    _ = fi55
                    _ = v56
                }
                var v57 bool = true
                _ = v57
            } else {
                switch "\n\t" {
                case "":
                    if true {
                        if false {
                            var v58 int64 = 0
                            v48 = 1e10
                            _ = v58
                        } else {
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                            v4 = append(v4, (v3 >> (((v1 % ((v1) | 1))) & 63)))
                        }
                        if true {
                            fmt.Println(v8, v8)
                            if (-2.25 < -2.25) {
                                var v59 int8 = v13
                                _ = v59
                            }
                        } else {
                            v13 += v13
                        }
                        switch 583 {
                        case 0:
                            v13 -= (v13 ^ 127)
                        case 42:
                            if true {
                                for fi60 := 0; fi60 < 19; fi60++ {
                                    if (32767 < 0) {
                                        if false {
                                            switch (fi60 >> (((v1 * 710)) & 63)) {
                                            case 42:
                                                fmt.Println(v48)
                                            default:
                                                for fi61 := 0; fi61 < 11; fi61++ {
                                                    var v62 uint16 = uint16(v1)
                                                    var v63 bool = true
                                                    v3 += v3
                                                    switch (v8 + v8) {
                                                    case "":
                                                        var v64 int32 = 42
                                                        _ = v64
                                                    default:
                                                        v63 = v63
                                                    }
                                                    _ = fi61
                                                    _ = v62
                                                    _ = v63
                                                }
                                            }
                                        } else {
                                            if false {
                                                if false {
                                                    if len(v2) > 0 {
                                                        v2[0] = (("" + v8) + "a")
                                                    }
                                                    v4 = append(v4, v3)
                                                    for fi65 := 0; fi65 < 19; fi65++ {
                                                        var v66 int = fi60
                                                        _ = fi65
                                                        _ = v66
                                                    }
                                                }
                                            } else {
                                                v48 -= float64(1.5)
                                                var v67 uint32 = (v3 << ((v1) & 63))
                                                var v68 float32 = 0.1
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                        if ("a" > v8) {
                                            _ = len(v2)
                                        } else {
                                            v1 += (v1 & v1)
                                            switch (("hello" + v8) + v8) {
                                            case "":
                                                v13 = 1
                                            }
                                            if true {
                                                for fi69 := 0; fi69 < 19; fi69++ {
                                                    fmt.Println(v13)
                                                    _ = fi69
                                                }
                                                if (true && true) {
                                                    if (!(858 <= 74)) {
                                                        fmt.Println(fi60, v48, v3)
                                                    }
                                                    switch v8 {
                                                    case "":
                                                        if (9223372036854775807 != -9223372036854775808) {
                                                            switch v8 {
                                                            case "a":
                                                                if ((v3 + (v3 * 65535)) > (v3 + v3)) {
                                                                    for fi70 := 0; fi70 < 18; fi70++ {
                                                                        fmt.Println(v8, v1)
                                                                        fmt.Println(v1)
                                                                        if true {
                                                                            if false {
                                                                                if (((-2.25 + 1e10) * 3.14159) <= (float32(v48) * (1.5 * 0.0))) {
                                                                                    v1 -= (^(v1))
                                                                                    v71 := GenS1{F0: (v3 & 65535), F1: uint16(v3), F2: 32767}
                                                                                    var v72 uint8 = 1
                                                                                    _ = v71
                                                                                    _ = v72
                                                                                } else {
                                                                                    v73 := GenS4{F0: (v8 + v8)}
                                                                                    v8 = v8
                                                                                    _ = v73
                                                                                }
                                                                                v74 := GenS5{F0: v3}
                                                                                v8 = v8
                                                                                switch 4294967295 {
                                                                                case -632:
                                                                                    v2 = append(v2, "hello")
                                                                                case -2147483648:
                                                                                    v48 = v48
                                                                                }
                                                                                _ = v74
                                                                            }
                                                                            _ = len(v2)
                                                                            switch 0 {
                                                                            case 4294967295:
                                                                                v75 := GenS6{F0: 42, F1: 42}
                                                                                _ = v75
                                                                            case -408:
                                                                                fmt.Println(v48, v48)
                                                                            default:
                                                                                _ = len(v2)
                                                                            }
                                                                        }
                                                                        _ = fi70
                                                                    }
                                                                }
                                                            default:
                                                                v3 = v3
                                                            }
                                                        } else {
                                                            v76 := GenS2{F0: (false || true)}
                                                            v48 += v48
                                                            var v77 int64 = 9223372036854775807
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                    case "a":
                                                        switch v8 {
                                                        case "":
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    default:
                                                        if true {
                                                            fmt.Println(v1, v3, v1)
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    }
                                                    v78 := GenS4{F0: "hello"}
                                                    _ = v78
                                                } else {
                                                    fmt.Println(v1, v13)
                                                }
                                            } else {
                                                v79 := GenS1{F0: 1, F1: 32767, F2: 42}
                                                _ = v79
                                            }
                                        }
                                        var v80 int16 = 1
                                        _ = v80
                                    } else {
                                        if false {
                                            var v81 int16 = 0
                                            v4 = append(v4, (v3 * (v3 << ((849) & 63))))
                                            _ = v81
                                        } else {
                                            v82 := GenS5{F0: v3}
                                            switch "foo bar" {
                                            case "hello":
                                                v13 -= v13
                                            case "foo bar":
                                                v83 := GenS1{F0: 12, F1: 42, F2: -1}
                                                _ = v83
                                            case "":
                                                var v84 uint = v1
                                                _ = v84
                                            }
                                            for fi85 := 0; fi85 < 14; fi85++ {
                                                _ = len(v4)
                                                switch fi85 {
                                                case 77:
                                                    v86 := GenS7{F0: v3, F1: 0, F2: (fi60 / (((fi85 % ((fi85) | 1))) | 1))}
                                                    _ = v86
                                                case -533:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v87 int64 = 1
                                                        switch v8 {
                                                        case "hello":
                                                            var v88 string = ("\n\t" + (v8 + ""))
                                                            _ = v88
                                                        case "a":
                                                            v87 = 9223372036854775807
                                                        case "":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        default:
                                                            v89 := GenS5{F0: (^(v3))}
                                                            _ = v89
                                                        }
                                                        var v90 uint32 = 42
                                                        _ = v87
                                                        _ = v90
                                                    case "a":
                                                        if (!(42 <= 0)) {
                                                            switch ((((v8 + "a") + ((v8 + v8) + v8)) + v8) + v8) {
                                                            case "\n\t":
                                                                v91 := GenS8{F0: 92, F1: 32767, F2: (fi85 * (fi60 / ((fi60) | 1)))}
                                                                _ = v91
                                                            case "foo bar":
                                                                if false {
                                                                    var v92 int = (-(fi60))
                                                                    var v93 string = "\n\t"
                                                                    _ = v92
                                                                    _ = v93
                                                                }
                                                            default:
                                                                v8 = (((v8 + v8) + v8) + (v8 + v8))
                                                            }
                                                        } else {
                                                            fmt.Println(v3)
                                                            v13 = 0
                                                        }
                                                    }
                                                case -889:
                                                    var v94 float32 = (-2.25 * float32(float64(3.14159)))
                                                    _ = v94
                                                default:
                                                    v8 = v8
                                                }
                                                var v95 int8 = v13
                                                _ = fi85
                                                _ = v95
                                            }
                                            _ = v82
                                        }
                                        v1 += (v1 | uint(v3))
                                        if true {
                                            var v96 int = fi60
                                            fmt.Println(v48, v1, v3)
                                            v97 := GenS9{F0: v13, F1: (^(v1))}
                                            _ = v96
                                            _ = v97
                                        } else {
                                            _ = len(v2)
                                            v13 += (v13 * (v13 / ((76) | 1)))
                                        }
                                    }
                                    _ = fi60
                                }
                                for fi98 := 0; fi98 < 7; fi98++ {
                                    var v99 float64 = (float64(-2.25) / v48)
                                    _ = fi98
                                    _ = v99
                                }
                                v3 = 2
                            }
                        }
                    }
                case "a":
                    v1 = (v1 + v1)
                case "foo bar":
                    fmt.Println(v8, v1)
                }
            }
        } else {
            var v100 int16 = -1
            _ = v100
        }
    default:
        switch v8 {
        case "a":
            for fi101 := 0; fi101 < 19; fi101++ {
                var v102 int64 = 42
                _ = fi101
                _ = v102
            }
        default:
            var v103 int16 = 32767
            _ = v103
        }
    }
    var v104 int64 = 0
    v104 += -50
    fmt.Println(v8, v1)
    fmt.Println(v48)
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v104
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v48
    _ = v8
}
