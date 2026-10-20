package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint; F1 int16; F2 int64 }
type GenS9 struct { F0 uint32 }

func main() {
    v1 := []uint32{535, 4294967295, 412, 2147483647}
    v2 := make(map[string]string)
    var v3 string = "hello"
    var v4 float32 = 1e10
    v5 := make(map[string]string)
    if (!true) {
        for fi6 := 0; fi6 < 9; fi6++ {
            fmt.Println(v4)
            v3 = "hello"
            v7 := GenS1{F0: 42, F1: ((v3 + "\n\t") + v3)}
            _ = fi6
            _ = v7
        }
    }
    v8 := GenS2{F0: 18446744073709551615, F1: 705}
    v4 = (3.14159 + v4)
    for fi9 := 0; fi9 < 2; fi9++ {
        if ((v4 * (v4 * v4)) < v4) {
            if (-10 < -128) {
                if ((!(65535 != 459)) && (-798 > 1)) {
                    v3 = "hello"
                }
                for fi10 := 0; fi10 < 19; fi10++ {
                    switch v3 {
                    case "a":
                        var v11 float64 = 1e10
                        if true {
                            if (((!false) && false) && (!((v3 + ("a" + v3)) > v3))) {
                                v12 := GenS1{F0: -128, F1: v3}
                                var v13 int = (fi10 / (((fi9 >> ((2) & 63))) | 1))
                                _ = v12
                                _ = v13
                            }
                            v14 := GenS3{F0: -9223372036854775808, F1: true}
                            _ = v14
                        }
                        if len(v1) > 0 {
                            v1[0] = 88
                        }
                        _ = v11
                    case "":
                        delete(v2, "hello")
                    }
                    _ = fi10
                }
            }
        }
        _ = fi9
    }
    if (-2147483648 != 42) {
        switch 4294967295 {
        case -163:
            for fi15 := 0; fi15 < 18; fi15++ {
                switch v3 {
                case "hello":
                    if false {
                        v1 = append(v1, 42)
                        for fi16 := 0; fi16 < 7; fi16++ {
                            var v17 uint8 = 42
                            v4 += v4
                            _ = fi16
                            _ = v17
                        }
                    } else {
                        v18 := GenS1{F0: -1, F1: v3}
                        for fi19 := 0; fi19 < 17; fi19++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch "\n\t" {
                            case "\n\t":
                                v20 := GenS1{F0: 42, F1: "a"}
                                _ = v20
                            case "a":
                                v21 := GenS3{F0: fi19, F1: (((((float32(((((0.0 / float64(1.5)) - ((0.0 * float64(v4)) - (0.0 * float64(1.5)))) / 1e10) / -2.25)) != v4) && false) || false) || true) && true)}
                                _ = v21
                            case "foo bar":
                                var v22 int = 42
                                _ = v22
                            default:
                                var v23 float32 = 0.1
                                _ = v23
                            }
                            switch v3 {
                            case "":
                                if (!((0.0 / (((-2.25 - 0.1) * 1e10) + float64(-2.25))) <= 1.5)) {
                                    delete(v2, v3)
                                }
                            case "hello":
                                v3 = (v3 + (v3 + (("a" + v3) + v3)))
                            case "foo bar":
                                switch (fi15 + (-(fi15))) {
                                case 424:
                                    v24 := GenS3{F0: (-(fi19)), F1: (true || (741 <= 1))}
                                    _ = v24
                                case -962:
                                    var v25 int8 = 42
                                    _ = v25
                                case -591:
                                    var v26 bool = (1 < 255)
                                    _ = v26
                                default:
                                    v3 = v3
                                    if false {
                                        var v27 int8 = 100
                                        _ = v27
                                    } else {
                                        if true {
                                            v4 = v4
                                            switch v3 {
                                            case "\n\t":
                                                v28 := GenS1{F0: 100, F1: "foo bar"}
                                                _ = v28
                                            default:
                                                v3 = v3
                                            }
                                            v29 := GenS1{F0: 1, F1: v3}
                                            _ = v29
                                        }
                                    }
                                }
                            default:
                                v3 = v3
                            }
                            _ = fi19
                        }
                        fmt.Println(v4)
                        _ = v18
                    }
                case "":
                    var v30 int64 = 299
                    _ = v30
                case "foo bar":
                    fmt.Println(fi15)
                }
                if (true && ((fi15 <= fi15) || ((false && false) && (true || (!true))))) {
                    if true {
                        var v31 int64 = 42
                        _ = v31
                    }
                    v3 = (v3 + ("a" + v3))
                }
                _ = fi15
            }
        case 32768:
            v3 = v3
        case -827:
            v32 := GenS4{F0: 42, F1: 42}
            _ = v32
        default:
            delete(v2, (v3 + "a"))
        }
        var v33 int8 = 42
        v3 = (v3 + ("foo bar" + v3))
        _ = v33
    }
    v4 -= float32(0.0)
    fmt.Println(v3, v3)
    switch 193 {
    case 128:
        if (false || false) {
            if true {
                if (65535 <= 75) {
                    fmt.Println(v4)
                    for fi34 := 0; fi34 < 2; fi34++ {
                        var v35 uint32 = 4294967295
                        _ = fi34
                        _ = v35
                    }
                    v1 = append(v1, 0)
                } else {
                    for fi36 := 0; fi36 < 16; fi36++ {
                        var v37 uint8 = 1
                        _ = fi36
                        _ = v37
                    }
                    fmt.Println(v3, v4)
                    var v38 int = -664
                    _ = v38
                }
                if ((1 > 28) || false) {
                    var v39 int64 = -9223372036854775808
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 12; fi40++ {
                        if (0.0 <= float32(((0.1 + 1.5) + float64((3.14159 * (v4 * v4)))))) {
                            var v41 float64 = 3.14159
                            v3 = v3
                            var v42 int = (-(fi40))
                            _ = v41
                            _ = v42
                        }
                        for fi43 := 0; fi43 < 7; fi43++ {
                            v3 = "a"
                            v44 := v2[v3]
                            _ = fi43
                            _ = v44
                        }
                        v45 := GenS2{F0: 65535, F1: 127}
                        _ = fi40
                        _ = v45
                    }
                }
            }
            v46 := GenS2{F0: 42, F1: 9223372036854775807}
            _ = v46
        }
    }
    fmt.Println(v4, v3)
    switch 1 {
    case -825:
        var v47 int = 2147483647
        _ = v47
    case 620:
        v2[(((((v3 + ((v3 + ((v3 + v3) + (("a" + v3) + v3))) + v3)) + (v3 + (v3 + "a"))) + v3) + "\n\t") + v3)] = (v3 + (("hello" + "") + (v3 + (v3 + (((((v3 + (v3 + "a")) + v3) + (v3 + (v3 + v3))) + v3) + "a")))))
    default:
        var v48 uint16 = 0
        _ = v48
    }
    v49 := GenS2{F0: 65535, F1: 2147483647}
    for fi50 := 0; fi50 < 20; fi50++ {
        if true {
            switch "foo bar" {
            case "a":
                switch -716 {
                case 42:
                    for fi51 := 0; fi51 < 12; fi51++ {
                        if (float32(0.0) >= v4) {
                            v3 = v3
                            v4 = 1e10
                            v5["a"] = "a"
                            v52 := GenS4{F0: 897, F1: 6}
                            _ = v52
                        } else {
                            var v53 uint64 = 2
                            _ = v53
                        }
                        _ = fi51
                    }
                case -938:
                    for fi54 := 0; fi54 < 1; fi54++ {
                        var v55 int16 = 1
                        switch v3 {
                        case "\n\t":
                            v4 -= 1e10
                        case "foo bar":
                            v4 = 0.1
                        case "":
                            switch v3 {
                            case "foo bar":
                                var v56 int8 = 42
                                _ = v56
                            }
                        }
                        switch fi54 {
                        case 127:
                            v4 -= float32(((-2.25 + 0.1) * 1e10))
                        case 32767:
                            var v57 uint16 = 2
                            _ = v57
                        case -1:
                            var v58 int = 32767
                            _ = v58
                        default:
                            switch 42 {
                            case -1:
                                if (-1 == 42) {
                                    v55 = (v55 ^ 588)
                                    fmt.Println(fi54, v55)
                                }
                            case -2147483648:
                                fmt.Println(v3)
                            }
                        }
                        var v59 uint = 42
                        v59 -= v59
                        _ = fi54
                        _ = v55
                        _ = v59
                    }
                }
            case "\n\t":
                var v60 uint16 = 65535
                _ = v60
            default:
                var v61 int16 = 1
                _ = v61
            }
            delete(v2, (v3 + v3))
        }
        if false {
            for fi62 := 0; fi62 < 3; fi62++ {
                switch (v3 + "") {
                case "foo bar":
                    var v63 uint64 = 4294967295
                    _ = v63
                }
                v4 -= v4
                _ = fi62
            }
            if (!true) {
                v3 = ((v3 + v3) + "a")
            } else {
                for fi64 := 0; fi64 < 17; fi64++ {
                    if false {
                        var v65 float64 = 0.0
                        switch v3 {
                        case "\n\t":
                            switch (v3 + "") {
                            case "a":
                                var v66 float64 = 1.5
                                _ = v66
                            case "foo bar":
                                delete(v5, ("hello" + "hello"))
                            default:
                                v67 := GenS4{F0: 1000, F1: 1}
                                _ = v67
                            }
                        }
                        switch (fi50 >> ((251) & 63)) {
                        case 9223372036854775807:
                            if (123 == 127) {
                                if len(v1) > 0 {
                                    v1[0] = 0
                                }
                            } else {
                                switch (((v3 + v3) + ("a" + (("a" + v3) + v3))) + v3) {
                                case "a":
                                    v3 = (v3 + ((v3 + ("" + v3)) + "foo bar"))
                                }
                                v68 := GenS2{F0: 255, F1: (fi50 - (fi64 / (((fi64 * 2147483647)) | 1)))}
                                if false {
                                    v65 = (v65 * (1e10 * v65))
                                } else {
                                    fmt.Println(v65)
                                }
                                _ = v68
                            }
                        case -2147483648:
                            v4 += 1.5
                        }
                        v69 := GenS5{F0: 0, F1: 255, F2: 42}
                        _ = v65
                        _ = v69
                    }
                    v3 = ("hello" + v3)
                    if (!true) {
                        var v70 bool = (1e10 == 1e10)
                        v70 = (42 == 0)
                        v3 = v3
                        _ = v70
                    } else {
                        var v71 string = v3
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        v71 = v3
                        _ = v71
                    }
                    _ = fi64
                }
                if false {
                    var v72 uint8 = 255
                    if false {
                        v4 -= v4
                        delete(v2, "")
                        _ = len(v1)
                    }
                    _ = v72
                } else {
                    v4 = ((v4 / (v4 + (0.0 + v4))) / float32((0.1 + 3.14159)))
                    var v73 string = v3
                    _ = v73
                }
            }
            var v74 uint64 = 42
            _ = v74
        }
        _ = fi50
    }
    _ = len(v1)
    for fi75 := 0; fi75 < 16; fi75++ {
        v3 = "hello"
        _ = len(v2)
        v3 = (v3 + ("a" + v3))
        _ = fi75
    }
    switch -9223372036854775808 {
    case -1:
        v3 = v3
    case 9223372036854775807:
        fmt.Println(v4, v4)
    case -841:
        if (!false) {
            v76 := GenS6{F0: 65535}
            _ = v76
        } else {
            v77 := GenS7{F0: 394}
            var v78 int16 = 32767
            v3 = ("foo bar" + ("foo bar" + "foo bar"))
            _ = v77
            _ = v78
        }
    }
    switch v3 {
    case "foo bar":
        switch "a" {
        case "hello":
            for fi79 := 0; fi79 < 15; fi79++ {
                var v80 bool = false
                if ((v80 || ((v80 && (42 < 32767)) && v80)) || (-1 >= -696)) {
                    switch fi79 {
                    case 214:
                        for fi81 := 0; fi81 < 15; fi81++ {
                            v3 = (("hello" + (v3 + v3)) + "")
                            if (0 <= -25) {
                                var v82 bool = v80
                                var v83 string = v3
                                v3 = "foo bar"
                                _ = v82
                                _ = v83
                            }
                            var v84 uint64 = 42
                            _ = fi81
                            _ = v84
                        }
                    }
                    for fi85 := 0; fi85 < 8; fi85++ {
                        v4 = v4
                        var v86 uint32 = 2147483647
                        _ = fi85
                        _ = v86
                    }
                } else {
                    switch (v3 + v3) {
                    case "hello":
                        v3 = "hello"
                    case "a":
                        var v87 int16 = -1
                        _ = v87
                    default:
                        var v88 uint16 = 624
                        _ = v88
                    }
                }
                if (!false) {
                    v89 := GenS8{F0: 706, F1: 1, F2: 1}
                    v90 := GenS5{F0: 179, F1: 65535, F2: -392}
                    _ = v89
                    _ = v90
                } else {
                    var v91 string = "\n\t"
                    v3 = (((v3 + ("" + v91)) + v3) + v3)
                    var v92 int32 = 0
                    _ = v91
                    _ = v92
                }
                _ = fi79
                _ = v80
            }
        case "\n\t":
            v93 := GenS9{F0: 256}
            _ = v93
        }
    case "\n\t":
        _ = len(v5)
    }
    var v94 int16 = 606
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v8
    _ = v94
}
