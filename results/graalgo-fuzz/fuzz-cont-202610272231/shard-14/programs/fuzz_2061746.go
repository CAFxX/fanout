package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint }
type GenS3 struct { F0 string; F1 int32; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 int64; F2 int }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool; F1 uint64; F2 int16 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 uint8; F1 int32 }

func main() {
    var v1 bool = (true && false)
    var v2 int64 = -9223372036854775808
    var v3 int16 = 209
    v4 := make(map[int]int)
    v3 = 1000
    switch ("foo bar" + "foo bar") {
    case "":
        switch "" {
        case "\n\t":
            v5 := []int{127, 0, -668}
            _ = v5
        }
    case "foo bar":
        switch ("foo bar" + "a") {
        case "hello":
            v6 := GenS1{F0: 1}
            _ = v6
        }
    case "a":
        switch "" {
        case "":
            v4[178] = 9223372036854775807
        }
    default:
        _ = len(v4)
    }
    if v1 {
        if true {
            if (!false) {
                _ = len(v4)
            } else {
                v3 = 1000
                v7 := v4[42]
                _ = v7
            }
        }
        if v1 {
            var v8 bool = v1
            v8 = true
            _ = v8
        }
        switch 899 {
        case 128:
            for fi9 := 0; fi9 < 20; fi9++ {
                delete(v4, -671)
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 bool = (!v1)
                    v1 = (607 < 737)
                    if (305 < 1) {
                        v12 := []int{fi10, fi9, fi9, fi9}
                        switch fi10 {
                        case 791:
                            switch (fi9 * 128) {
                            case -94:
                                var v13 uint32 = 586
                                _ = v13
                            case -199:
                                v11 = v1
                            case 9223372036854775807:
                                fmt.Println(v2, v3, fi10)
                            }
                        case 805:
                            var v14 int32 = 0
                            _ = v14
                        default:
                            var v15 int8 = 42
                            _ = v15
                        }
                        if v11 {
                            v11 = true
                            if v1 {
                                v16 := GenS1{F0: 255}
                                fmt.Println(fi9, v3)
                                var v17 uint = 121
                                fmt.Println(v2)
                                if (v1 || v1) {
                                    switch "" {
                                    case "a":
                                        v18 := GenS2{F0: 235, F1: -1, F2: v17}
                                        _ = v18
                                    case "":
                                        v3 = 1
                                        var v19 float32 = (1.5 - 1.5)
                                        var v20 int32 = 2147483647
                                        _ = v19
                                        _ = v20
                                    }
                                } else {
                                    fmt.Println(fi10)
                                    if v11 {
                                        if true {
                                            _ = len(v4)
                                            if len(v12) > 0 {
                                                v12[0] = -9223372036854775808
                                            }
                                        } else {
                                            fmt.Println(v17, fi9)
                                        }
                                        v11 = false
                                    }
                                    fmt.Println(v11, v11)
                                    switch (fi9 + (fi9 >> ((1) & 63))) {
                                    case 9223372036854775807:
                                        var v21 uint = 0
                                        _ = v21
                                    default:
                                        v11 = (255 != 42)
                                    }
                                    v12 = append(v12, fi10)
                                    v2 = v2
                                }
                                _ = v16
                                _ = v17
                            }
                            var v22 uint16 = 2
                            _ = v22
                        } else {
                            v1 = (v3 < v3)
                            switch "" {
                            case "a":
                                v4[128] = fi9
                            }
                        }
                        _ = v12
                    } else {
                        v1 = true
                        v23 := GenS3{F0: (("" + "\n\t") + "a"), F1: 1, F2: -2147483648}
                        _ = v23
                    }
                    _ = fi10
                    _ = v11
                }
                if (27 <= 42) {
                    if true {
                        v1 = ("a" != "hello")
                    }
                    for fi24 := 0; fi24 < 6; fi24++ {
                        if (true || v1) {
                            if ((-2.25 - (-2.25 - 0.0)) <= 1e10) {
                                if (((((1.5 / 3.14159) - ((3.14159 * (1e10 / 1e10)) + (0.1 + 3.14159))) - (3.14159 / 1.5)) + -2.25) < 1.5) {
                                    v25 := GenS4{F0: 0}
                                    v1 = ((v1 && (!v1)) || (!v1))
                                    _ = v25
                                } else {
                                    fmt.Println(v1, fi24, fi9)
                                }
                                v2 += (v2 + (v2 ^ v2))
                                fmt.Println(fi24)
                            } else {
                                var v26 uint8 = 2
                                fmt.Println(v1)
                                _ = v26
                            }
                            var v27 bool = (42 >= 42)
                            delete(v4, (fi9 & -925))
                            _ = v27
                        }
                        _ = fi24
                    }
                    v2 = (v2 >> ((4294967295) & 63))
                }
                _ = fi9
            }
        case 4294967295:
            fmt.Println(v1)
        case 42:
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 uint16 = 117
                fmt.Println(v3, v29)
                v2 -= -571
                _ = fi28
                _ = v29
            }
        }
    }
    v3 += v3
    switch -1 {
    case 32767:
        var v30 float32 = 3.14159
        _ = v30
    case -9:
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 float64 = float64(-2.25)
            _ = fi31
            _ = v32
        }
    case 42:
        v33 := GenS4{F0: 127}
        _ = v33
    default:
        fmt.Println(v2, v2)
    }
    v3 = (^(v3))
    var v34 float32 = 0.0
    v1 = (!true)
    switch (("foo bar" + "foo bar") + "") {
    case "a":
        switch 4294967295 {
        case -1:
            for fi35 := 0; fi35 < 7; fi35++ {
                v36 := v4[-521]
                v37 := GenS5{F0: (v2 >> ((255) & 63)), F1: v2, F2: (v36 - (v36 / ((fi35) | 1)))}
                _ = fi35
                _ = v36
                _ = v37
            }
        case 599:
            switch 42 {
            case -2147483648:
                switch "\n\t" {
                case "a":
                    v38 := GenS5{F0: 42, F1: (v2 >> ((18446744073709551615) & 63)), F2: 0}
                    _ = v38
                }
            case 0:
                for fi39 := 0; fi39 < 13; fi39++ {
                    var v40 uint8 = 127
                    if (((v40 | (^(v40))) == v40) && v1) {
                        v3 += (v3 + -786)
                    }
                    for fi41 := 0; fi41 < 4; fi41++ {
                        switch (fi39 | (fi39 * (fi39 % (((fi41 - (fi39 | 4294967295))) | 1)))) {
                        case -595:
                            fmt.Println(v40)
                        case 128:
                            if (!v1) {
                                v42 := GenS2{F0: (v40 % (((v40 - (v40 << ((935) & 63)))) | 1)), F1: -2147483648, F2: 552}
                                var v43 uint32 = 669
                                _ = v42
                                _ = v43
                            }
                        default:
                            fmt.Println(fi41, v40, v2)
                        }
                        var v44 int64 = (v2 & (v2 | 0))
                        v45 := GenS6{F0: uint(v40), F1: v3}
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v1, v3)
                        case "":
                            fmt.Println(v3)
                        case "a":
                            v46 := GenS5{F0: (v44 / ((v44) | 1)), F1: v44, F2: (-(fi39))}
                            _ = v46
                        default:
                            var v47 float64 = 3.14159
                            _ = v47
                        }
                        var v48 int16 = v3
                        _ = fi41
                        _ = v44
                        _ = v45
                        _ = v48
                    }
                    _ = fi39
                    _ = v40
                }
            case 32768:
                if true {
                    v49 := GenS7{F0: 42}
                    _ = v49
                }
            default:
                switch 4294967295 {
                case -759:
                    fmt.Println(v3)
                case -2147483648:
                    switch 727 {
                    case -1:
                        v50 := []int{128, -263, 32767}
                        _ = v50
                    case 127:
                        for fi51 := 0; fi51 < 6; fi51++ {
                            if false {
                                fmt.Println(fi51, v34, fi51)
                            } else {
                                var v52 string = ""
                                _ = v52
                            }
                            var v53 uint8 = 2
                            var v54 bool = v1
                            _ = fi51
                            _ = v53
                            _ = v54
                        }
                    default:
                        if v1 {
                            v34 = v34
                            switch -1 {
                            case 42:
                                switch 32767 {
                                case -1:
                                    for fi55 := 0; fi55 < 1; fi55++ {
                                        v34 -= (v34 * 1e10)
                                        switch (fi55 + (fi55 % ((fi55) | 1))) {
                                        case 0:
                                            if (1e10 >= float64(1e10)) {
                                                for fi56 := 0; fi56 < 10; fi56++ {
                                                    v57 := []int{(fi56 / ((530) | 1))}
                                                    _ = fi56
                                                    _ = v57
                                                }
                                                var v58 int32 = -1
                                                _ = v58
                                            } else {
                                                v59 := GenS5{F0: (-(v2)), F1: (v2 / ((v2) | 1)), F2: int(v2)}
                                                var v60 uint = 255
                                                _ = v59
                                                _ = v60
                                            }
                                        case 127:
                                            v34 += ((1.5 + (v34 / v34)) - float32(1.5))
                                        }
                                        _ = fi55
                                    }
                                default:
                                    v61 := []int{2147483647, -2147483648}
                                    _ = v61
                                }
                            case 32767:
                                if (((69 == 255) && (v3 != -32768)) || (v34 <= v34)) {
                                    v3 -= (v3 << ((65535) & 63))
                                } else {
                                    for fi62 := 0; fi62 < 5; fi62++ {
                                        var v63 float64 = 1e10
                                        var v64 uint8 = 128
                                        if (1 > 100) {
                                            var v65 uint16 = 42
                                            for fi66 := 0; fi66 < 2; fi66++ {
                                                v63 = 3.14159
                                                v67 := GenS2{F0: 1, F1: -876, F2: 255}
                                                _ = fi66
                                                _ = v67
                                            }
                                            var v68 int8 = -1
                                            _ = v65
                                            _ = v68
                                        } else {
                                            if (!(18446744073709551615 >= 919)) {
                                                if (2 == 214) {
                                                    if (!(!v1)) {
                                                        var v69 uint32 = 0
                                                        for fi70 := 0; fi70 < 3; fi70++ {
                                                            v63 -= v63
                                                            _ = fi70
                                                        }
                                                        v71 := GenS1{F0: (v64 << ((857) & 63))}
                                                        _ = v69
                                                        _ = v71
                                                    } else {
                                                        v72 := GenS7{F0: 36}
                                                        v64 -= 42
                                                        _ = v72
                                                    }
                                                } else {
                                                    var v73 int16 = 1000
                                                    if ((228 >= 603) || false) {
                                                        for fi74 := 0; fi74 < 13; fi74++ {
                                                            fmt.Println(v1, v2)
                                                            _ = fi74
                                                        }
                                                        var v75 uint = 255
                                                        _ = v75
                                                    } else {
                                                        v34 = (1.5 + v34)
                                                        switch 2147483647 {
                                                        case 32768:
                                                            v4[(fi62 | -112)] = -428
                                                        case 2147483647:
                                                            var v76 int16 = (v3 * v3)
                                                            _ = v76
                                                        default:
                                                            if v1 {
                                                                var v77 int = 1
                                                                _ = v77
                                                            }
                                                        }
                                                        for fi78 := 0; fi78 < 18; fi78++ {
                                                            if false {
                                                                v3 = v73
                                                                v79 := GenS8{F0: v1, F1: 2, F2: (v3 << ((18446744073709551615) & 63))}
                                                                _ = v79
                                                            }
                                                            _ = fi78
                                                        }
                                                    }
                                                    _ = v73
                                                }
                                                var v80 bool = v1
                                                v81 := GenS9{F0: (v64 & v64), F1: 2147483647}
                                                _ = v80
                                                _ = v81
                                            } else {
                                                fmt.Println(fi62, fi62, v34)
                                                fmt.Println(v3, v64, v64)
                                            }
                                            v82 := GenS3{F0: "", F1: 1, F2: -2147483648}
                                            _ = v82
                                        }
                                        _ = fi62
                                        _ = v63
                                        _ = v64
                                    }
                                    fmt.Println(v34, v1, v3)
                                    switch 2147483647 {
                                    case -9223372036854775808:
                                        for fi83 := 0; fi83 < 8; fi83++ {
                                            var v84 int32 = 42
                                            _ = fi83
                                            _ = v84
                                        }
                                    case -44:
                                        for fi85 := 0; fi85 < 10; fi85++ {
                                            for fi86 := 0; fi86 < 7; fi86++ {
                                                var v87 int8 = 0
                                                v88 := GenS7{F0: (v87 >> ((279) & 63))}
                                                fmt.Println(v1)
                                                _ = fi86
                                                _ = v87
                                                _ = v88
                                            }
                                            v3 = (^(v3))
                                            _ = fi85
                                        }
                                    }
                                }
                            }
                            for fi89 := 0; fi89 < 12; fi89++ {
                                if v1 {
                                    for fi90 := 0; fi90 < 13; fi90++ {
                                        fmt.Println(v1)
                                        _ = fi90
                                    }
                                }
                                v91 := v4[(fi89 / ((fi89) | 1))]
                                var v92 uint64 = 0
                                _ = fi89
                                _ = v91
                                _ = v92
                            }
                        } else {
                            if (!(2147483647 != 2147483648)) {
                                v34 += (float32(0.1) - ((1e10 + (1e10 / ((v34 * (float32(((-2.25 - 3.14159) * (((-2.25 + 1e10) + 3.14159) * ((3.14159 - -2.25) * (float64(0.0) / 3.14159))))) + ((((v34 - (v34 + v34)) - (1e10 / 1.5)) / (((1e10 / ((v34 + (0.0 / v34)) * 1.5)) / 1e10) + 0.0)) * (float32((0.1 * 3.14159)) / v34)))) * v34))) - 0.0))
                                v3 = (v3 % ((v3) | 1))
                                if (!(int64(v3) < v2)) {
                                    v2 = v2
                                    for fi93 := 0; fi93 < 6; fi93++ {
                                        delete(v4, (fi93 % (((fi93 >> ((959) & 63))) | 1)))
                                        _ = fi93
                                    }
                                    v34 = v34
                                }
                            }
                        }
                    }
                case 32768:
                    _ = len(v4)
                }
            }
        case 42:
            v94 := []int{-2147483648, 128, 128, 2147483647}
            _ = v94
        default:
            v95 := v4[128]
            _ = v95
        }
    case "\n\t":
        if (v1 || (!(v1 && (!v1)))) {
            for fi96 := 0; fi96 < 16; fi96++ {
                if (!true) {
                    fmt.Println(fi96, v3)
                    v34 = (-2.25 + 1.5)
                    v97 := GenS10{F0: 255, F1: 1}
                    _ = v97
                }
                switch -2147483648 {
                case 9223372036854775807:
                    var v98 int64 = (v2 % ((860) | 1))
                    _ = v98
                }
                _ = fi96
            }
        }
    default:
        var v99 string = ("foo bar" + ("hello" + "foo bar"))
        _ = v99
    }
    var v100 string = ""
    v101 := GenS4{F0: 242}
    if (-111 > 1) {
        v2 = v2
    } else {
        var v102 uint16 = 822
        v103 := []int{127}
        _ = v102
        _ = v103
    }
    v3 += (v3 ^ -14)
    if false {
        if (32768 < int(v2)) {
            v104 := []int{429}
            var v105 int16 = -32768
            _ = v104
            _ = v105
        }
        var v106 uint8 = 0
        var v107 uint8 = v106
        _ = v106
        _ = v107
    }
    v34 = ((v34 * float32(3.14159)) * 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
