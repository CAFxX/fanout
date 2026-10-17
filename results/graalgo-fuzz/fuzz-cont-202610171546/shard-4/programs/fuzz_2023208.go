package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 int32; F1 uint32; F2 bool }
type GenS8 struct { F0 int; F1 bool }

func main() {
    var v1 int64 = -65
    v2 := []uint32{2147483647, 2, 0}
    var v3 int = 42
    var v4 bool = true
    var v5 float64 = float64(0.1)
    v6 := make(map[uint32]bool)
    v4 = true
    if v4 {
        if (127 >= 13) {
            if (v4 && (-51 <= 1)) {
                if ((1000 == 0) || false) {
                    var v7 int8 = 42
                    switch ("a" + "a") {
                    case "":
                        var v8 uint8 = 199
                        _ = v8
                    default:
                        v6[65535] = v4
                    }
                    delete(v6, 2)
                    _ = v7
                } else {
                    for fi9 := 0; fi9 < 20; fi9++ {
                        for fi10 := 0; fi10 < 19; fi10++ {
                            v11 := GenS1{F0: 1000}
                            _ = fi10
                            _ = v11
                        }
                        if (("\n\t" + ("hello" + ("foo bar" + ("" + ("" + "foo bar"))))) != ("a" + "")) {
                            v12 := v6[65535]
                            switch (fi9 << ((4294967295) & 63)) {
                            case 127:
                                v1 += v1
                            case 0:
                                fmt.Println(v12, fi9, v4)
                            }
                            switch (("foo bar" + "foo bar") + "\n\t") {
                            case "\n\t":
                                for fi13 := 0; fi13 < 8; fi13++ {
                                    var v14 int16 = 42
                                    _ = fi13
                                    _ = v14
                                }
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = 852
                                }
                            case "a":
                                if (42 <= 255) {
                                    v5 = v5
                                    v1 -= 42
                                }
                            default:
                                var v15 uint8 = 255
                                _ = v15
                            }
                            _ = v12
                        } else {
                            v4 = false
                            for fi16 := 0; fi16 < 20; fi16++ {
                                v17 := GenS1{F0: -32768}
                                v4 = (171 >= 2)
                                switch "" {
                                case "\n\t":
                                    v18 := GenS2{F0: 0, F1: 82}
                                    _ = v18
                                case "hello":
                                    _ = len(v6)
                                default:
                                    switch (v3 ^ v3) {
                                    case 128:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case 2147483647:
                                        if len(v2) > 0 {
                                            v2[0] = 42
                                        }
                                    default:
                                        v19 := GenS1{F0: -1}
                                        _ = v19
                                    }
                                }
                                _ = fi16
                                _ = v17
                            }
                            var v20 uint64 = 0
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 10; fi21++ {
                            if (!(!v4)) {
                                v22 := v6[0]
                                var v23 int32 = -1
                                switch fi9 {
                                case 32768:
                                    fmt.Println(fi9, v4, v1)
                                }
                                _ = v22
                                _ = v23
                            }
                            _ = fi21
                        }
                        _ = fi9
                    }
                }
                fmt.Println(v3)
            } else {
                for fi24 := 0; fi24 < 6; fi24++ {
                    v5 = v5
                    v25 := GenS3{F0: 732}
                    if (!v4) {
                        v3 += 822
                        var v26 int64 = (v1 << ((540) & 63))
                        for fi27 := 0; fi27 < 1; fi27++ {
                            var v28 int64 = (v26 >> ((1) & 63))
                            _ = fi27
                            _ = v28
                        }
                        _ = v26
                    } else {
                        v2 = append(v2, 722)
                        var v29 int8 = 100
                        _ = v29
                    }
                    _ = fi24
                    _ = v25
                }
                v5 += (((((1.5 - (-2.25 + v5)) - (v5 / 1.0)) - 0.0) + 0.1) + v5)
            }
            fmt.Println(v4, v1)
            _ = len(v2)
        }
        switch "" {
        case "a":
            switch -1 {
            case 2147483647:
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
            default:
                switch ("" + ("foo bar" + ((("foo bar" + "hello") + "a") + "hello"))) {
                case "hello":
                    if ((32767 > 255) && (65535 > 462)) {
                        var v30 int64 = (v1 >> ((18446744073709551615) & 63))
                        _ = v30
                    } else {
                        v3 -= 9223372036854775807
                        v31 := GenS4{F0: 91, F1: 160, F2: 1}
                        _ = v31
                    }
                }
            }
        case "\n\t":
            v1 = v1
        case "":
            switch (v3 << ((4294967295) & 63)) {
            case -172:
                fmt.Println(v3, v3)
            case -9223372036854775808:
                v3 = v3
            default:
                v4 = true
            }
        }
        v5 = 3.14159
    }
    fmt.Println(v1, v3)
    switch "\n\t" {
    case "foo bar":
        var v32 int32 = -2147483648
        _ = v32
    default:
        if len(v2) > 0 {
            v2[0] = 1
        }
    }
    switch -2147483648 {
    case 9223372036854775807:
        for fi33 := 0; fi33 < 19; fi33++ {
            v1 += (v1 % ((v1) | 1))
            _ = fi33
        }
    default:
        v4 = false
    }
    v4 = (255 >= 65535)
    var v34 int = v3
    if v4 {
        if (588 <= 157) {
            for fi35 := 0; fi35 < 4; fi35++ {
                for fi36 := 0; fi36 < 1; fi36++ {
                    switch v34 {
                    case 927:
                        var v37 uint = 333
                        _ = v37
                    case 2147483647:
                        switch 32767 {
                        case -616:
                            if (!v4) {
                                fmt.Println(fi36, v4)
                                switch "hello" {
                                case "\n\t":
                                    switch (v34 * (fi36 % (((fi35 ^ 1)) | 1))) {
                                    case 2147483647:
                                        switch fi36 {
                                        case -2147483648:
                                            fmt.Println(v1, v3)
                                        case 128:
                                            fmt.Println(v4)
                                        case -1:
                                            switch "\n\t" {
                                            case "a":
                                                fmt.Println(v3)
                                            case "hello":
                                                v1 = (v1 >> ((2) & 63))
                                            case "foo bar":
                                                var v38 string = "foo bar"
                                                _ = v38
                                            default:
                                                v39 := v6[2147483647]
                                                _ = v39
                                            }
                                        }
                                    }
                                case "hello":
                                    var v40 int = (v34 % ((-2147483648) | 1))
                                    _ = v40
                                default:
                                    v34 = 128
                                }
                            }
                        case -2147483648:
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                        case -9223372036854775808:
                            var v41 int = 32768
                            _ = v41
                        }
                    case 0:
                        delete(v6, 694)
                    default:
                        fmt.Println(fi35, v5, v4)
                    }
                    _ = fi36
                }
                _ = fi35
            }
            var v42 uint16 = 255
            _ = v42
        } else {
            var v43 string = "\n\t"
            fmt.Println(v1, v34, v4)
            v44 := GenS4{F0: -1, F1: 0, F2: 705}
            _ = v43
            _ = v44
        }
    }
    fmt.Println(v34, v5)
    if (!v4) {
        v45 := GenS5{F0: 42}
        _ = v45
    }
    if ((0.0 <= 3.14159) && (!v4)) {
        switch "foo bar" {
        case "a":
            var v46 int16 = 42
            _ = v46
        case "hello":
            v47 := GenS2{F0: 0, F1: -128}
            _ = v47
        default:
            _ = len(v2)
        }
        var v48 int8 = 127
        v4 = (-2147483648 > -1)
        _ = v48
    } else {
        fmt.Println(v5)
        switch "" {
        case "":
            if (2147483647 == -865) {
                v4 = true
                var v49 uint = 0
                _ = v49
            } else {
                switch (((("foo bar" + "\n\t") + ((("foo bar" + "") + "") + "a")) + (("foo bar" + "foo bar") + "a")) + "") {
                case "hello":
                    for fi50 := 0; fi50 < 1; fi50++ {
                        for fi51 := 0; fi51 < 5; fi51++ {
                            v34 -= (fi51 & (v34 % ((v3) | 1)))
                            v52 := GenS3{F0: 2}
                            v53 := v6[2]
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        _ = fi50
                    }
                case "":
                    switch v34 {
                    case 32767:
                        var v54 uint64 = 856
                        _ = v54
                    case -245:
                        var v55 uint8 = 42
                        _ = v55
                    case 518:
                        v56 := GenS5{F0: v1}
                        _ = v56
                    }
                default:
                    var v57 uint = 1
                    _ = v57
                }
                var v58 int16 = 42
                _ = v58
            }
        case "hello":
            if (v1 > (v1 + (v1 / ((-1) | 1)))) {
                v59 := GenS6{F0: 42, F1: 0}
                for fi60 := 0; fi60 < 3; fi60++ {
                    if v4 {
                        var v61 int32 = -1
                        _ = v61
                    }
                    for fi62 := 0; fi62 < 2; fi62++ {
                        if (1e10 <= -2.25) {
                            if true {
                                if (v4 && v4) {
                                    if v4 {
                                        fmt.Println(v4, fi60)
                                        v5 = (v5 - v5)
                                    } else {
                                        v4 = true
                                        var v63 float64 = v5
                                        var v64 int16 = -32768
                                        _ = v63
                                        _ = v64
                                    }
                                    v6[4294967295] = v4
                                    var v65 uint16 = 1
                                    _ = v65
                                } else {
                                    var v66 int64 = (v1 & v1)
                                    _ = v66
                                }
                                switch (("" + "\n\t") + "\n\t") {
                                case "hello":
                                    v67 := GenS3{F0: 2}
                                    _ = v67
                                case "foo bar":
                                    var v68 float64 = 3.14159
                                    _ = v68
                                }
                                var v69 uint = 425
                                _ = v69
                            }
                            v34 += fi60
                            var v70 int64 = (-(v1))
                            _ = v70
                        }
                        v71 := GenS7{F0: -2147483648, F1: 2147483648, F2: (((0.1 + (-2.25 - -2.25)) - -2.25) < 3.14159)}
                        _ = fi62
                        _ = v71
                    }
                    _ = fi60
                }
                _ = v59
            }
        default:
            for fi72 := 0; fi72 < 1; fi72++ {
                for fi73 := 0; fi73 < 13; fi73++ {
                    v74 := GenS2{F0: 42, F1: -104}
                    v3 = (v3 << ((667) & 63))
                    _ = fi73
                    _ = v74
                }
                _ = fi72
            }
        }
        for fi75 := 0; fi75 < 14; fi75++ {
            var v76 string = "foo bar"
            fmt.Println(v1, v5)
            v77 := GenS8{F0: v34, F1: (true && v4)}
            _ = fi75
            _ = v76
            _ = v77
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
