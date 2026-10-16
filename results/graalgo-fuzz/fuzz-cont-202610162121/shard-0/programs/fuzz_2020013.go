package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 int; F1 int64; F2 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16 }

func main() {
    v1 := []uint32{953, 2147483647}
    var v2 uint8 = 128
    var v3 uint64 = 4294967295
    for fi4 := 0; fi4 < 20; fi4++ {
        if true {
            v5 := GenS1{F0: ("foo bar" + (("foo bar" + "foo bar") + "a")), F1: 2147483647, F2: (v2 | v2)}
            _ = v5
        }
        v6 := GenS1{F0: "a", F1: 2, F2: (^(v2))}
        if false {
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 uint = 0
                v9 := GenS1{F0: ((("" + "hello") + ("\n\t" + "foo bar")) + (((("a" + (("a" + ("a" + "a")) + ("" + "hello"))) + "\n\t") + "") + "a")), F1: 0, F2: v2}
                var v10 int16 = -32768
                var v11 int8 = 127
                _ = fi7
                _ = v10
                _ = v11
                _ = v8
                _ = v9
            }
        } else {
            v12 := GenS1{F0: "a", F1: 1, F2: v2}
            _ = v12
        }
        _ = fi4
        _ = v6
    }
    if len(v1) > 0 {
        v1[0] = 526
    }
    v13 := GenS1{F0: "hello", F1: 815, F2: 127}
    if (!true) {
        for fi14 := 0; fi14 < 12; fi14++ {
            v3 -= v3
            v3 -= (v3 << ((696) & 63))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi14
        }
        if (true && (10 == -1)) {
            v3 -= (v3 ^ v3)
            if (861 < 1) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                switch -923 {
                case 2147483647:
                    v15 := GenS2{F0: "foo bar"}
                    _ = v15
                case 9223372036854775807:
                    var v16 bool = (-72 == 85)
                    _ = v16
                case 1:
                    v3 -= v3
                default:
                    var v17 int32 = 42
                    _ = v17
                }
                if true {
                    if (1 > 2147483648) {
                        fmt.Println(v2, v3)
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        if (("foo bar" + ("hello" + ("foo bar" + ("\n\t" + "hello")))) == "a") {
                            v2 = (v2 >> ((4294967295) & 63))
                            for fi18 := 0; fi18 < 9; fi18++ {
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(fi18, fi18, v3)
                                case "hello":
                                    fmt.Println(v3, v2)
                                case "foo bar":
                                    var v19 uint16 = 2
                                    _ = v19
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = 383
                                    }
                                }
                                _ = fi18
                            }
                            v1 = append(v1, 2147483647)
                        }
                    }
                    var v20 uint32 = 13
                    _ = v20
                }
            }
        }
    }
    _ = len(v1)
    fmt.Println(v2)
    fmt.Println(v2, v2)
    v2 += v2
    var v21 int16 = 0
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS1{F0: "foo bar", F1: 0, F2: 42}
            if false {
                fmt.Println(fi23, v21, v21)
                if true {
                    v25 := GenS2{F0: "foo bar"}
                    v2 -= v2
                    _ = v25
                }
                var v26 uint32 = 2147483647
                var v27 uint8 = v2
                _ = v26
                _ = v27
            }
            var v28 uint = 1
            if len(v1) > 0 {
                v1[0] = uint32(v3)
            }
            v29 := GenS3{F0: ("a" + "hello"), F1: false}
            _ = fi23
            _ = v24
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 9; fi30++ {
            v1 = append(v1, 209)
            if false {
                v1 = append(v1, 253)
                v21 -= v21
                switch fi22 {
                case -1:
                    _ = len(v1)
                case -9223372036854775808:
                    fmt.Println(fi22, v21, v3)
                case 249:
                    switch (("\n\t" + "foo bar") + "hello") {
                    case "a":
                        v31 := GenS4{F0: 0}
                        if false {
                            switch ("hello" + "hello") {
                            case "\n\t":
                                if true {
                                    var v32 int32 = -1
                                    switch fi22 {
                                    case 32768:
                                        switch ("\n\t" + (((("\n\t" + "") + "hello") + (("foo bar" + "\n\t") + ("foo bar" + ("foo bar" + "a")))) + (("a" + ("a" + "")) + ("foo bar" + "hello")))) {
                                        case "\n\t":
                                            var v33 uint32 = 2147483648
                                            _ = v33
                                        case "":
                                            if ((((v32 + v32) == v32) && ((!(true || false)) || false)) || false) {
                                                v21 -= (v21 >> ((255) & 63))
                                                var v34 uint = 18446744073709551615
                                                _ = v34
                                            } else {
                                                var v35 uint8 = 42
                                                if false {
                                                    v36 := GenS5{F0: fi30, F1: 307}
                                                    fmt.Println(v32, v32, v3)
                                                    _ = v36
                                                }
                                                _ = v35
                                            }
                                        case "a":
                                            if false {
                                                v1 = append(v1, 42)
                                                if (v21 != v21) {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v37 := GenS5{F0: (-(fi22)), F1: 124}
                                                        _ = v37
                                                    case "a":
                                                        v38 := GenS3{F0: "", F1: false}
                                                        if false {
                                                            _ = len(v1)
                                                        }
                                                        fmt.Println(fi30, fi30, fi30)
                                                        v2 = v2
                                                        v39 := GenS6{F0: (fi22 | fi30), F1: 1, F2: (v32 / (((v32 >> ((0) & 63))) | 1))}
                                                        if (true && (false || true)) {
                                                            var v40 string = "\n\t"
                                                            _ = v40
                                                        } else {
                                                            if (((-1 < -9223372036854775808) && (v2 > (v2 | 128))) && (false && true)) {
                                                                switch "a" {
                                                                case "a":
                                                                    _ = len(v1)
                                                                case "\n\t":
                                                                    fmt.Println(v3, v2, v3)
                                                                case "foo bar":
                                                                    var v41 int64 = -562
                                                                    _ = v41
                                                                default:
                                                                    switch ("" + "\n\t") {
                                                                    case "a":
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 660
                                                                        }
                                                                    }
                                                                    fmt.Println(v21, v2)
                                                                }
                                                                switch fi22 {
                                                                case 128:
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 232
                                                                    }
                                                                default:
                                                                    v3 = v3
                                                                    if (0.0 > 0.1) {
                                                                        var v42 uint = 42
                                                                        var v43 uint16 = 0
                                                                        fmt.Println(v43)
                                                                        _ = v42
                                                                        _ = v43
                                                                    }
                                                                }
                                                            } else {
                                                                if (!((true || ((!(v21 == v21)) && (!true))) || (int64(fi22) > -197))) {
                                                                    v44 := GenS6{F0: 32767, F1: -558, F2: 978}
                                                                    _ = v44
                                                                } else {
                                                                    var v45 float64 = 1.5
                                                                    fmt.Println(v2)
                                                                    _ = v45
                                                                }
                                                                var v46 uint16 = 2
                                                                _ = v46
                                                            }
                                                            if false {
                                                                _ = len(v1)
                                                            } else {
                                                                if true {
                                                                    v21 -= v21
                                                                    var v47 int32 = 114
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v48 int64 = 629
                                                                        _ = v48
                                                                    }
                                                                    if (-1 >= -38) {
                                                                        v21 = -1
                                                                    }
                                                                    if ((fi22 & -1) >= 42) {
                                                                        var v49 bool = true
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            v1 = append(v1, 2147483648)
                                                                        }
                                                                        switch ((("hello" + "hello") + ("hello" + "")) + "foo bar") {
                                                                        case "a":
                                                                            var v50 float64 = (((0.0 * 3.14159) * 0.1) + (((-2.25 - (float64(0.0) / 1.0)) + float64(3.14159)) * 1.5))
                                                                            _ = v50
                                                                        case "hello":
                                                                            switch fi30 {
                                                                            case 42:
                                                                                fmt.Println(v49, v21)
                                                                            case 2147483647:
                                                                                var v51 int = fi30
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        var v52 uint32 = 625
                                                                                        _ = v52
                                                                                    case "\n\t":
                                                                                        var v53 int8 = 0
                                                                                        _ = v53
                                                                                    case "foo bar":
                                                                                        switch v51 {
                                                                                        case 9223372036854775807:
                                                                                            var v54 int64 = 0
                                                                                            _ = v54
                                                                                        }
                                                                                    }
                                                                                case "a":
                                                                                    if v49 {
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 65535
                                                                                        }
                                                                                        v1 = append(v1, 0)
                                                                                        fmt.Println(v21, fi30)
                                                                                    }
                                                                                }
                                                                                _ = v51
                                                                            case -650:
                                                                                var v55 int = fi30
                                                                                _ = v55
                                                                            }
                                                                        case "foo bar":
                                                                            var v56 int64 = 1
                                                                            _ = v56
                                                                        }
                                                                        _ = v49
                                                                    } else {
                                                                        if true {
                                                                            v21 = int16(fi22)
                                                                            var v57 int64 = 1
                                                                            var v58 uint16 = 65535
                                                                            _ = v57
                                                                            _ = v58
                                                                        }
                                                                        v59 := GenS7{F0: false}
                                                                        v21 = v21
                                                                        var v60 int32 = (v47 & (-(v47)))
                                                                        var v61 float64 = 1.5
                                                                        _ = v59
                                                                        _ = v60
                                                                        _ = v61
                                                                    }
                                                                    _ = v47
                                                                }
                                                                v32 -= v32
                                                                var v62 string = "a"
                                                                _ = v62
                                                            }
                                                        }
                                                        _ = v38
                                                        _ = v39
                                                    case "hello":
                                                        var v63 float64 = float64(1e10)
                                                        _ = v63
                                                    default:
                                                        v64 := GenS8{F0: v32}
                                                        _ = v64
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v65 int64 = 0
                                                    _ = v65
                                                }
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v66 := GenS3{F0: ("a" + ""), F1: true}
                                                var v67 int8 = 100
                                                _ = v66
                                                _ = v67
                                            }
                                            switch "" {
                                            case "hello":
                                                v2 += (v2 + v2)
                                            case "a":
                                                v3 = v3
                                            }
                                            var v68 int32 = 42
                                            _ = v68
                                        }
                                    default:
                                        v3 = v3
                                    }
                                    _ = v32
                                }
                            default:
                                var v69 uint = 913
                                _ = v69
                            }
                        }
                        var v70 int = -632
                        v70 = (fi30 * fi30)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v1 = append(v1, 808)
                        _ = v31
                        _ = v70
                    case "hello":
                        v2 += v2
                    }
                default:
                    var v71 bool = ((-2.25 - ((3.14159 - float32((0.0 / 1.0))) / 1e10)) <= 0.0)
                    _ = v71
                }
            } else {
                if true {
                    var v72 uint = 4294967295
                    var v73 int8 = -1
                    _ = v72
                    _ = v73
                } else {
                    fmt.Println(v2, fi30)
                }
            }
            _ = fi30
        }
        _ = fi22
    }
    for fi74 := 0; fi74 < 16; fi74++ {
        var v75 int32 = 1
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi74
        _ = v75
    }
    for fi76 := 0; fi76 < 5; fi76++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi77 := 0; fi77 < 14; fi77++ {
            var v78 bool = (!false)
            v79 := GenS9{F0: int16(fi76)}
            _ = fi77
            _ = v78
            _ = v79
        }
        var v80 uint32 = 4294967295
        _ = fi76
        _ = v80
    }
    var v81 int16 = int16(v2)
    switch -1 {
    case -1:
        var v82 int16 = v81
        _ = v82
    case 443:
        for fi83 := 0; fi83 < 11; fi83++ {
            v2 += (v2 >> ((65535) & 63))
            _ = fi83
        }
    }
    var v84 int = 32767
    var v85 uint32 = 2
    v2 += 136
    if false {
        v1 = append(v1, v85)
        var v86 string = "hello"
        v2 += 73
        _ = v86
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v81:", v81)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v81
    _ = v84
    _ = v85
}
