package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int8; F1 uint8; F2 int16 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 string; F2 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int64 }

func main() {
    var v1 int = 4294967295
    v2 := []uint32{1, 65535, 42, 42}
    v3 := []uint32{65535, 17, 30, 65535}
    var v4 int = (v1 % (((^(v1))) | 1))
    var v5 string = ""
    if (true && false) {
        v1 += v1
    }
    fmt.Println(v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += v4
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v6 bool = false
    if v6 {
        var v7 int64 = 9223372036854775807
        v8 := GenS1{F0: 4294967295, F1: (264 < 42)}
        if (3.14159 != (((((((float32(0.0) + ((0.0 / -2.25) * -2.25)) - 1e10) - ((float32(3.14159) + (-2.25 + 0.0)) - (1e10 * 1e10))) * 3.14159) / -2.25) / 1.0) * (1.5 + 1e10))) {
            v6 = v6
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 uint64 = 4294967295
                v1 = (v4 | v1)
                if (-158 >= fi9) {
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    for fi11 := 0; fi11 < 18; fi11++ {
                        v3 = append(v3, 929)
                        v12 := GenS1{F0: (v10 ^ (v10 << ((0) & 63))), F1: true}
                        switch v5 {
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                        case "hello":
                            var v13 uint64 = (v10 + (v10 >> ((42) & 63)))
                            _ = v13
                        }
                        v10 = (v10 | 1)
                        v2 = append(v2, 1)
                        _ = fi11
                        _ = v12
                    }
                    switch (v1 | 2147483647) {
                    case -9223372036854775808:
                        if (-265 <= 42) {
                            _ = len(v2)
                            for fi14 := 0; fi14 < 13; fi14++ {
                                v1 = (^(fi14))
                                if false {
                                    v6 = (!((0.1 == 1.5) || (!(856 == (v7 - v7)))))
                                    _ = len(v2)
                                    v15 := GenS1{F0: v10, F1: v6}
                                    _ = v15
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                    v10 = 4294967295
                                }
                                var v16 int32 = -2147483648
                                switch (v5 + "") {
                                case "a":
                                    switch (v4 << ((65535) & 63)) {
                                    case -1:
                                        switch ((v5 + "\n\t") + v5) {
                                        case "hello":
                                            v4 -= fi14
                                        case "foo bar":
                                            var v17 float32 = 0.1
                                            _ = v17
                                        default:
                                            v10 = v10
                                        }
                                    case 97:
                                        var v18 int64 = v7
                                        _ = v18
                                    default:
                                        if false {
                                            v19 := GenS1{F0: 383, F1: false}
                                            fmt.Println(v6)
                                            _ = v19
                                        }
                                    }
                                }
                                v4 = -563
                                _ = fi14
                                _ = v16
                            }
                        } else {
                            switch "hello" {
                            case "":
                                v20 := GenS1{F0: 18446744073709551615, F1: (!(313 == 255))}
                                _ = v20
                            case "a":
                                if len(v3) > 0 {
                                    v3[0] = uint32(v10)
                                }
                            case "\n\t":
                                if ((-355 < 1000) && (!true)) {
                                    var v21 uint16 = 255
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v21
                                }
                            }
                            for fi22 := 0; fi22 < 3; fi22++ {
                                var v23 bool = false
                                if (127 == 2) {
                                    if len(v2) > 0 {
                                        v2[0] = 141
                                    }
                                    v5 = v5
                                } else {
                                    switch (fi22 + 2147483647) {
                                    case 2147483647:
                                        v7 = v7
                                    case 32767:
                                        v24 := GenS1{F0: (^(v10)), F1: true}
                                        _ = v24
                                    case -711:
                                        if true {
                                            if true {
                                                if len(v3) > 0 {
                                                    v3[0] = 4294967295
                                                }
                                                v7 += v7
                                            }
                                        }
                                    default:
                                        switch (^(fi9)) {
                                        case 127:
                                            if v23 {
                                                var v25 int = (fi22 | -652)
                                                var v26 string = v5
                                                fmt.Println(v5)
                                                var v27 uint64 = v10
                                                v25 = -837
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            }
                                        case 135:
                                            if true {
                                                if false {
                                                    v28 := GenS2{F0: -128, F1: 234, F2: -393}
                                                    v29 := GenS3{F0: v7, F1: 42, F2: 255}
                                                    var v30 int16 = -32768
                                                    _ = v28
                                                    _ = v29
                                                    _ = v30
                                                }
                                                v5 = v5
                                            }
                                        case -254:
                                            var v31 float32 = 1.5
                                            _ = v31
                                        default:
                                            switch "a" {
                                            case "a":
                                                var v32 uint32 = uint32(v10)
                                                _ = v32
                                            case "\n\t":
                                                v23 = (!true)
                                            default:
                                                switch v5 {
                                                case "":
                                                    var v33 uint8 = 2
                                                    _ = v33
                                                }
                                            }
                                        }
                                    }
                                }
                                if (876 != 934) {
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                    if (0 >= 339) {
                                        fmt.Println(v5)
                                        if len(v3) > 0 {
                                            v3[0] = 114
                                        }
                                    }
                                    v34 := GenS4{F0: 64, F1: v5, F2: (((v5 + "hello") + v5) + v5)}
                                    _ = v34
                                } else {
                                    switch (v4 + 1) {
                                    case 9223372036854775807:
                                        if v23 {
                                            v3 = append(v3, 42)
                                            if v6 {
                                                v5 = ""
                                            } else {
                                                if (v10 < (v10 | v10)) {
                                                    switch (-(fi9)) {
                                                    case 42:
                                                        if (!v23) {
                                                            v10 = 4294967295
                                                        }
                                                    }
                                                }
                                                v4 = (v1 ^ v1)
                                                fmt.Println(fi22, v5)
                                            }
                                            var v35 int = fi22
                                            v2 = append(v2, 65535)
                                            _ = v35
                                        } else {
                                            if false {
                                                v4 = (fi9 & (fi22 & v4))
                                                if v6 {
                                                    var v36 uint32 = 42
                                                    if (2 == 2) {
                                                        if len(v2) > 0 {
                                                            v2[0] = 569
                                                        }
                                                        v23 = true
                                                        if true {
                                                            if len(v2) > 0 {
                                                                v2[0] = (v36 * v36)
                                                            }
                                                            switch (-(v1)) {
                                                            case 625:
                                                                fmt.Println(v6)
                                                            default:
                                                                v37 := GenS5{F0: 42}
                                                                _ = v37
                                                            }
                                                        } else {
                                                            var v38 int16 = 42
                                                            _ = v38
                                                        }
                                                    }
                                                    _ = v36
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = 0
                                                    }
                                                    v39 := GenS5{F0: 42}
                                                    _ = v39
                                                }
                                                _ = len(v3)
                                            } else {
                                                v4 = -538
                                                v3 = append(v3, 42)
                                            }
                                        }
                                        v40 := GenS1{F0: 0, F1: true}
                                        var v41 int16 = 437
                                        _ = v40
                                        _ = v41
                                    case -2147483648:
                                        var v42 int32 = 0
                                        _ = v42
                                    case -1:
                                        var v43 uint = 722
                                        _ = v43
                                    }
                                    fmt.Println(fi9)
                                    v4 += fi22
                                }
                                _ = fi22
                                _ = v23
                            }
                        }
                    default:
                        switch ("foo bar" + (("\n\t" + ("a" + "hello")) + "foo bar")) {
                        case "hello":
                            v44 := GenS3{F0: (v7 - (^(v7))), F1: 0, F2: 32767}
                            _ = v44
                        }
                    }
                } else {
                    if true {
                        if len(v2) > 0 {
                            v2[0] = 2147483647
                        }
                        var v45 uint32 = 4294967295
                        _ = v45
                    }
                }
                _ = fi9
                _ = v10
            }
            v46 := GenS4{F0: 127, F1: (v5 + (v5 + ("" + "hello"))), F2: v5}
            _ = v46
        }
        _ = v7
        _ = v8
    } else {
        fmt.Println(v4)
    }
    var v47 int8 = 100
    v48 := GenS6{F0: 255, F1: (!v6)}
    if v6 {
        v49 := GenS7{F0: 0, F1: "\n\t"}
        v50 := GenS1{F0: 2, F1: v6}
        for fi51 := 0; fi51 < 1; fi51++ {
            var v52 int16 = -32768
            switch (v4 - (fi51 % ((v1) | 1))) {
            case 410:
                switch 9223372036854775807 {
                case -154:
                    for fi53 := 0; fi53 < 5; fi53++ {
                        v54 := GenS8{F0: 0}
                        _ = fi53
                        _ = v54
                    }
                case 499:
                    v6 = (487 == -1)
                case 9223372036854775807:
                    for fi55 := 0; fi55 < 2; fi55++ {
                        v56 := GenS9{F0: 4294967295}
                        v47 -= v47
                        v5 = (v5 + (v5 + v5))
                        _ = fi55
                        _ = v56
                    }
                }
            case 32767:
                var v57 uint8 = 33
                _ = v57
            case -2147483648:
                if (!v6) {
                    if v6 {
                        for fi58 := 0; fi58 < 6; fi58++ {
                            v59 := GenS10{F0: 42}
                            if (!v6) {
                                switch -793 {
                                case 4294967295:
                                    var v60 uint32 = 42
                                    _ = v60
                                case -838:
                                    if v6 {
                                        v4 = (^(fi58))
                                        v6 = (42 < 0)
                                        v6 = true
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            v6 = (!(!false))
                                        case "hello":
                                            var v61 int = (^(fi58))
                                            _ = v61
                                        default:
                                            if len(v2) > 0 {
                                                v2[0] = 42
                                            }
                                        }
                                        v62 := GenS11{F0: 42}
                                        _ = v62
                                    }
                                case 32768:
                                    v52 = v52
                                }
                            }
                            v1 = (^(fi51))
                            _ = fi58
                            _ = v59
                        }
                        var v63 string = v5
                        if ((^(fi51)) < (fi51 | (v1 * v4))) {
                            if v6 {
                                for fi64 := 0; fi64 < 16; fi64++ {
                                    v65 := GenS7{F0: 0, F1: v63}
                                    _ = fi64
                                    _ = v65
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v63
                    }
                    if len(v3) > 0 {
                        v3[0] = 891
                    }
                }
            }
            for fi66 := 0; fi66 < 18; fi66++ {
                v6 = (-792 != v52)
                var v67 uint8 = 218
                if v6 {
                    v67 = v67
                }
                var v68 int64 = -9223372036854775808
                _ = fi66
                _ = v67
                _ = v68
            }
            _ = fi51
            _ = v52
        }
        _ = v49
        _ = v50
    }
    if (!(701 == 509)) {
        var v69 uint8 = 200
        if (42 < 1) {
            v3 = append(v3, 65535)
            fmt.Println(v6, v4, v5)
        } else {
            for fi70 := 0; fi70 < 5; fi70++ {
                var v71 string = "\n\t"
                var v72 int16 = -190
                switch (-(v1)) {
                case 9223372036854775807:
                    switch v5 {
                    case "hello":
                        fmt.Println(fi70)
                    case "\n\t":
                        if (v69 < (v69 << ((0) & 63))) {
                            var v73 int8 = (v47 >> ((255) & 63))
                            var v74 uint64 = 280
                            _ = v73
                            _ = v74
                        }
                    case "":
                        var v75 float64 = 1.5
                        _ = v75
                    }
                }
                _ = fi70
                _ = v71
                _ = v72
            }
        }
        switch v1 {
        case 9223372036854775807:
            _ = len(v2)
        case -808:
            fmt.Println(v6, v5)
        }
        _ = v69
    } else {
        if false {
            for fi76 := 0; fi76 < 20; fi76++ {
                for fi77 := 0; fi77 < 15; fi77++ {
                    v5 = "a"
                    fmt.Println(v6, fi77)
                    var v78 int = v4
                    _ = fi77
                    _ = v78
                }
                _ = fi76
            }
            if v6 {
                if (!v6) {
                    v79 := GenS10{F0: 2}
                    _ = v79
                }
            }
        }
    }
    v1 = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
