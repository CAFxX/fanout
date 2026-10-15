package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int32; F2 int }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 bool; F1 uint64; F2 bool }
type GenS7 struct { F0 int; F1 bool; F2 int16 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint; F1 string; F2 uint16 }
type GenS11 struct { F0 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 934
    v3 := []uint32{(v2 ^ (^(v2))), (v2 ^ (v2 & (v2 ^ (v2 ^ (v2 ^ v2)))))}
    var v4 int = 2147483647
    v5 := make(map[string]int)
    if (v2 == (v2 * 1)) {
        var v6 int = -102
        var v7 bool = ((float64(1.5) == 3.14159) || false)
        _ = v6
        _ = v7
    }
    v2 += 428
    for fi8 := 0; fi8 < 7; fi8++ {
        _ = len(v3)
        _ = fi8
    }
    if (1000 > 809) {
        var v9 bool = (2 < 128)
        switch ("hello" + (("hello" + "hello") + "foo bar")) {
        case "hello":
            for fi10 := 0; fi10 < 1; fi10++ {
                if ("\n\t" != "a") {
                    fmt.Println(v4)
                    v11 := GenS1{F0: 1, F1: (844 == -1)}
                    switch v4 {
                    case -9223372036854775808:
                        var v12 bool = (false && v9)
                        _ = v12
                    case 32768:
                        var v13 bool = (42 < 0)
                        _ = v13
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v9, v4)
                        case "hello":
                            switch "hello" {
                            case "":
                                for fi14 := 0; fi14 < 18; fi14++ {
                                    fmt.Println(v9, v9)
                                    _ = fi14
                                }
                            }
                        }
                    }
                    _ = v11
                } else {
                    fmt.Println(fi10)
                    fmt.Println(v2)
                }
                var v15 string = "hello"
                _ = fi10
                _ = v15
            }
        case "":
            v4 -= -1
        case "a":
            var v16 int = v4
            _ = v16
        default:
            var v17 string = ("" + ("a" + (("\n\t" + ("" + "foo bar")) + "")))
            _ = v17
        }
        v5["hello"] = v4
        _ = v9
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            if false {
                v4 = (fi19 | (fi18 << ((2) & 63)))
                v2 = (v2 / ((v2) | 1))
                delete(v5, "")
            } else {
                fmt.Println(fi19)
                v4 += (-(fi19))
            }
            if (!true) {
                var v20 int16 = int16(v2)
                if (!(!(!false))) {
                    v2 = (v2 | v2)
                    switch fi18 {
                    case 2147483647:
                        v2 = 2147483648
                    case -930:
                        var v21 int8 = -121
                        _ = v21
                    case 4294967295:
                        v22 := GenS1{F0: 128, F1: true}
                        if (0 <= 196) {
                            switch "" {
                            case "":
                                switch 945 {
                                case 128:
                                    fmt.Println(v2, fi19)
                                case 32767:
                                    v4 = fi19
                                default:
                                    switch (^(v4)) {
                                    case -664:
                                        var v23 bool = false
                                        fmt.Println(v20, fi19, v2)
                                        _ = v23
                                    default:
                                        switch (fi19 | v4) {
                                        case 32767:
                                            v24 := GenS1{F0: 149, F1: true}
                                            _ = v24
                                        case -2147483648:
                                            _ = len(v1)
                                        case 0:
                                            fmt.Println(fi18, v20, fi18)
                                        }
                                    }
                                }
                            }
                            v3 = append(v3, 2147483648)
                            fmt.Println(fi18, fi19)
                        } else {
                            v25 := GenS1{F0: 42, F1: ((^(v2)) > 2147483648)}
                            _ = v25
                        }
                        _ = v22
                    default:
                        if false {
                            var v26 float64 = ((-2.25 - 0.1) + (-2.25 * (float64((3.14159 / 0.1)) * 1.5)))
                            v27 := GenS1{F0: 16, F1: (2 <= 111)}
                            if (!false) {
                                v28 := GenS1{F0: 1, F1: true}
                                _ = v28
                            }
                            _ = v26
                            _ = v27
                        }
                    }
                } else {
                    v20 = (-(v20))
                }
                v2 = (^(v2))
                _ = v20
            }
            _ = fi19
        }
        v4 += (v4 / (((v4 << ((18446744073709551615) & 63))) | 1))
        fmt.Println(fi18)
        _ = fi18
    }
    v2 += (v2 >> ((775) & 63))
    for fi29 := 0; fi29 < 1; fi29++ {
        v4 += -1
        v30 := GenS2{F0: ("" + "a")}
        _ = fi29
        _ = v30
    }
    if len(v3) > 0 {
        v3[0] = 23
    }
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        v31 := GenS1{F0: 35, F1: true}
        _ = v31
    } else {
        v32 := GenS2{F0: "\n\t"}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v32
    }
    var v33 uint16 = 255
    switch "hello" {
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        _ = len(v3)
        if true {
            switch "hello" {
            case "foo bar":
                if (!(v33 <= 628)) {
                    var v35 int32 = 867
                    if (-815 == -32768) {
                        for fi36 := 0; fi36 < 19; fi36++ {
                            if false {
                                if true {
                                    var v37 uint32 = (v2 + v2)
                                    switch ("" + ((("a" + "\n\t") + "a") + "a")) {
                                    case "foo bar":
                                        v38 := v1[(-(fi34))]
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v39 uint64 = 4294967295
                                            _ = v39
                                        case "a":
                                            switch ("a" + "foo bar") {
                                            case "\n\t":
                                                var v40 float64 = 1e10
                                                _ = v40
                                            case "":
                                                switch -1 {
                                                case -412:
                                                    var v41 float32 = -2.25
                                                    _ = v41
                                                case 9223372036854775807:
                                                    v37 -= v2
                                                default:
                                                    var v42 int16 = 32767
                                                    _ = v42
                                                }
                                            default:
                                                if ((v33 << ((1) & 63)) < (^(v33))) {
                                                    v37 = v37
                                                    fmt.Println(v35, v37)
                                                    var v43 int8 = 1
                                                    _ = v43
                                                } else {
                                                    v3 = append(v3, 42)
                                                }
                                            }
                                        }
                                        v44 := v5["\n\t"]
                                        _ = v38
                                        _ = v44
                                    case "":
                                        v45 := GenS3{F0: true}
                                        _ = v45
                                    case "hello":
                                        if true {
                                            v37 = v2
                                            v37 += (v37 ^ (^(v37)))
                                            v46 := GenS4{F0: (^(v33)), F1: v35, F2: fi34}
                                            v4 -= (fi34 % (((^(fi36))) | 1))
                                            _ = v46
                                        } else {
                                            var v47 uint = 4294967295
                                            _ = v47
                                        }
                                    default:
                                        fmt.Println(v33)
                                    }
                                    var v48 uint8 = 32
                                    _ = v37
                                    _ = v48
                                }
                            }
                            _ = fi36
                        }
                        v35 -= (v35 ^ v35)
                        var v49 uint16 = (v33 << ((2) & 63))
                        _ = v49
                    }
                    _ = v35
                } else {
                    for fi50 := 0; fi50 < 15; fi50++ {
                        delete(v1, fi50)
                        if (2 <= 4294967295) {
                            if false {
                                fmt.Println(v4, fi50)
                                fmt.Println(fi34, fi50)
                            }
                            v4 = fi50
                            v51 := GenS5{F0: 333, F1: 127}
                            v33 += v33
                            _ = v51
                        } else {
                            v52 := GenS6{F0: false, F1: 1, F2: (fi34 < fi50)}
                            var v53 uint = 42
                            if false {
                                v2 += v2
                            }
                            _ = v52
                            _ = v53
                        }
                        _ = fi50
                    }
                }
            case "\n\t":
                if false {
                    switch (^(fi34)) {
                    case 2147483647:
                        switch 375 {
                        case 247:
                            for fi54 := 0; fi54 < 8; fi54++ {
                                var v55 bool = (1 >= 0)
                                _ = fi54
                                _ = v55
                            }
                        case 9223372036854775807:
                            var v56 uint8 = 0
                            _ = v56
                        case 56:
                            v4 = (v4 ^ (v4 | (fi34 & 127)))
                        default:
                            switch ("hello" + "") {
                            case "\n\t":
                                if false {
                                    _ = len(v5)
                                    v33 -= v33
                                } else {
                                    for fi57 := 0; fi57 < 6; fi57++ {
                                        if (true || (true && false)) {
                                            switch fi34 {
                                            case 830:
                                                v33 += v33
                                            case 4294967295:
                                                v33 += 13
                                                _ = len(v3)
                                            }
                                            var v58 int64 = 116
                                            _ = v58
                                        } else {
                                            var v59 int8 = 42
                                            _ = v59
                                        }
                                        v33 = v33
                                        _ = fi57
                                    }
                                    if false {
                                        if (-336 < 42) {
                                            v60 := GenS7{F0: 1, F1: ((v33 >> ((1) & 63)) > (^(v33))), F2: -339}
                                            if true {
                                                if (!((!(-577 >= -9223372036854775808)) && (true || (65535 >= 42)))) {
                                                    v5["a"] = 128
                                                } else {
                                                    v1[(^(fi34))] = false
                                                    switch "" {
                                                    case "":
                                                        v2 += (v2 + v2)
                                                    case "\n\t":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                }
                                                v2 += (v2 >> ((697) & 63))
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            if true {
                                                v33 -= (v33 & (v33 * 42))
                                            }
                                            _ = v60
                                        }
                                        for fi61 := 0; fi61 < 2; fi61++ {
                                            var v62 uint16 = v33
                                            _ = fi61
                                            _ = v62
                                        }
                                    }
                                    v63 := GenS4{F0: uint16(v2), F1: -560, F2: v4}
                                    _ = v63
                                }
                            case "foo bar":
                                v64 := GenS6{F0: (0 > 1), F1: 260, F2: true}
                                _ = v64
                            }
                        }
                    case 975:
                        fmt.Println(v2, v4, v33)
                    default:
                        var v65 int8 = 100
                        _ = v65
                    }
                    switch fi34 {
                    case -2147483648:
                        var v66 bool = true
                        _ = v66
                    case -1:
                        var v67 int64 = 9223372036854775807
                        _ = v67
                    case 899:
                        v68 := GenS8{F0: (^(v33)), F1: 1}
                        _ = v68
                    default:
                        delete(v1, fi34)
                    }
                    v2 += uint32(v33)
                }
            case "hello":
                v4 -= (fi34 * 127)
            }
        }
        _ = fi34
    }
    switch (("foo bar" + "a") + "\n\t") {
    case "\n\t":
        v2 -= (v2 ^ 4294967295)
    case "hello":
        v69 := GenS9{F0: 42}
        _ = v69
    case "":
        for fi70 := 0; fi70 < 14; fi70++ {
            v71 := v5["foo bar"]
            _ = fi70
            _ = v71
        }
    }
    v72 := GenS10{F0: 42, F1: "\n\t", F2: v33}
    v73 := GenS11{F0: uint16(v2)}
    var v74 uint16 = (v33 | (v33 % ((v33) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v72
    _ = v73
    _ = v74
}
