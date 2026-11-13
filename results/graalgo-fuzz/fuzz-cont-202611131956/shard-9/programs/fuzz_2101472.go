package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint16; F1 bool; F2 int16 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint; F1 uint }
type GenS11 struct { F0 uint; F1 string; F2 uint8 }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    switch 4294967295 {
    case 9223372036854775807:
        v2[408] = false
    case 32768:
        var v4 string = "\n\t"
        _ = v4
    case 2147483647:
        if (((!true) || true) && false) {
            v5 := v2[2147483648]
            _ = v5
        }
    default:
        v6 := GenS1{F0: 4294967295}
        _ = v6
    }
    switch -9223372036854775808 {
    case 1:
        var v7 uint = 255
        _ = v7
    case 32768:
        var v8 uint32 = 65535
        _ = v8
    }
    var v9 uint = 1
    if false {
        var v10 float64 = (0.1 / (3.14159 + 1.5))
        v10 += v10
        if len(v1) > 0 {
            v1[0] = -1
        }
        _ = v10
    } else {
        v11 := GenS2{F0: -1, F1: 255}
        switch "\n\t" {
        case "hello":
            for fi12 := 0; fi12 < 11; fi12++ {
                v9 = v9
                if false {
                    var v13 int32 = -944
                    _ = v13
                }
                fmt.Println(fi12, fi12)
                _ = fi12
            }
        case "":
            var v14 string = ""
            _ = v14
        case "foo bar":
            v9 = 18446744073709551615
        default:
            var v15 string = "\n\t"
            _ = v15
        }
        _ = v11
    }
    switch "hello" {
    case "a":
        fmt.Println(v9)
    case "\n\t":
        var v16 string = "foo bar"
        _ = v16
    case "":
        v9 += (^(v9))
    }
    delete(v2, 199)
    switch -977 {
    case -335:
        var v17 uint32 = 2147483648
        _ = v17
    default:
        if false {
            var v18 float64 = 0.1
            if false {
                for fi19 := 0; fi19 < 11; fi19++ {
                    for fi20 := 0; fi20 < 15; fi20++ {
                        if true {
                            v18 -= ((1.5 / float64(float32(1e10))) * v18)
                            _ = len(v1)
                            v18 = (v18 + 1e10)
                        } else {
                            if (("a" + "hello") == "hello") {
                                switch 1 {
                                case 0:
                                    if (!(!(((-1 > 820) || (true && true)) || false))) {
                                        fmt.Println(v9, v18)
                                        fmt.Println(v9, v9, v18)
                                        var v21 int16 = 976
                                        _ = v21
                                    } else {
                                        if false {
                                            if true {
                                                if len(v1) > 0 {
                                                    v1[0] = (fi19 | 0)
                                                }
                                            }
                                            if (42 > 9223372036854775807) {
                                                var v22 bool = (3.14159 <= (0.1 * 0.1))
                                                var v23 int16 = 32767
                                                _ = v22
                                                _ = v23
                                            }
                                            v18 = (1e10 + v18)
                                        }
                                    }
                                case 32767:
                                    var v24 float64 = v18
                                    _ = v24
                                default:
                                    switch "" {
                                    case "foo bar":
                                        if false {
                                            v1 = append(v1, fi20)
                                            v25 := GenS3{F0: 0, F1: (2147483647 > 478), F2: 0}
                                            _ = v25
                                        }
                                    }
                                }
                            }
                        }
                        var v26 int = -393
                        _ = fi20
                        _ = v26
                    }
                    v18 = (v18 * v18)
                    v18 += float64(1.5)
                    _ = fi19
                }
            }
            switch "\n\t" {
            case "foo bar":
                v3 = append(v3, 65535)
            case "a":
                for fi27 := 0; fi27 < 18; fi27++ {
                    for fi28 := 0; fi28 < 20; fi28++ {
                        if ((("foo bar" + "foo bar") + "\n\t") >= ("hello" + (("\n\t" + ("a" + "")) + ("hello" + "a")))) {
                            var v29 int8 = 100
                            v18 += (((v18 + ((3.14159 - v18) / v18)) * (3.14159 - (v18 + float64(0.1)))) * v18)
                            _ = v29
                        }
                        delete(v2, 2)
                        _ = fi28
                    }
                    switch fi27 {
                    case -9223372036854775808:
                        switch 32768 {
                        case 1:
                            v18 = v18
                        }
                    default:
                        v2[375] = true
                    }
                    _ = fi27
                }
            default:
                _ = len(v1)
            }
            _ = v18
        }
    }
    switch ("a" + "\n\t") {
    case "foo bar":
        v9 = v9
    case "hello":
        for fi30 := 0; fi30 < 13; fi30++ {
            if false {
                var v31 int32 = 0
                _ = v31
            } else {
                v32 := GenS1{F0: (^(v9))}
                for fi33 := 0; fi33 < 19; fi33++ {
                    switch 2147483647 {
                    case -48:
                        v9 += (v9 & 255)
                    case -9223372036854775808:
                        v9 += 42
                        switch 127 {
                        case -268:
                            fmt.Println(fi30)
                        case 286:
                            var v34 uint64 = uint64(v9)
                            _ = v34
                        case 42:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            _ = len(v2)
                        }
                        v2[0] = false
                    case 32767:
                        v9 = 18446744073709551615
                    default:
                        v9 -= (v9 ^ 141)
                    }
                    v9 = 246
                    v9 -= v9
                    _ = fi33
                }
                _ = v32
            }
            _ = fi30
        }
    case "\n\t":
        if (-1 < 759) {
            if (0 != 1) {
                if false {
                    switch (("a" + ("hello" + ("a" + ("hello" + "hello")))) + "hello") {
                    case "hello":
                        v35 := GenS4{F0: v9, F1: 1}
                        _ = v35
                    default:
                        v9 -= v9
                    }
                    v9 = v9
                    var v36 float64 = ((1.5 * 1.5) / 1.5)
                    _ = v36
                } else {
                    v9 = v9
                }
            } else {
                fmt.Println(v9)
            }
            switch 127 {
            case 828:
                if ((127 != 39) && (!true)) {
                    v37 := v2[2]
                    _ = v37
                } else {
                    switch ("a" + (("" + "") + ((("" + ("a" + "hello")) + "foo bar") + ""))) {
                    case "foo bar":
                        var v38 int8 = -1
                        _ = v38
                    case "hello":
                        for fi39 := 0; fi39 < 3; fi39++ {
                            v9 += (^(v9))
                            _ = fi39
                        }
                    case "":
                        v1 = append(v1, 9223372036854775807)
                    default:
                        if (3.14159 <= -2.25) {
                            fmt.Println(v9)
                            delete(v2, 42)
                        } else {
                            fmt.Println(v9)
                        }
                    }
                }
            case 0:
                switch ("hello" + ("a" + "foo bar")) {
                case "\n\t":
                    if (42 < 0) {
                        v40 := v2[493]
                        var v41 float64 = float64(1e10)
                        v42 := GenS5{F0: -319}
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                case "hello":
                    var v43 bool = (4294967295 >= 439)
                    _ = v43
                default:
                    v9 = v9
                }
            case -9223372036854775808:
                v9 = (v9 & 1)
            }
            var v44 int8 = 42
            _ = v44
        }
    default:
        v45 := GenS4{F0: 18446744073709551615, F1: 182}
        _ = v45
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        fmt.Println(fi46, fi46)
        _ = fi46
    }
    v9 += v9
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 float32 = 0.0
        _ = fi47
        _ = v48
    }
    if (("hello" + "a") > "a") {
        v2[0] = true
        fmt.Println(v9)
    } else {
        fmt.Println(v9)
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v9)
        default:
            delete(v2, 535)
        }
        for fi49 := 0; fi49 < 18; fi49++ {
            switch "a" {
            case "":
                _ = len(v2)
            case "a":
                if (!true) {
                    var v50 int16 = 32767
                    if (284 < 255) {
                        v51 := GenS3{F0: 228, F1: (0 < v50), F2: 1}
                        for fi52 := 0; fi52 < 2; fi52++ {
                            v50 -= v50
                            if true {
                                v53 := GenS2{F0: -2147483648, F1: 930}
                                _ = v53
                            } else {
                                switch "a" {
                                case "":
                                    _ = len(v3)
                                case "foo bar":
                                    if (!(!(42 < 0))) {
                                        if (false || true) {
                                            v50 += (v50 % ((v50) | 1))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        if true {
                                            v9 -= v9
                                            v54 := GenS1{F0: v9}
                                            var v55 bool = false
                                            v9 += (v9 % (((v9 * 255)) | 1))
                                            var v56 int8 = int8(v50)
                                            v56 = v56
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                        } else {
                                            v57 := GenS6{F0: "a", F1: v9}
                                            _ = v57
                                        }
                                    }
                                default:
                                    if (fi49 != 932) {
                                        fmt.Println(fi52)
                                        fmt.Println(v9)
                                        var v58 int = fi52
                                        _ = v58
                                    } else {
                                        v59 := GenS7{F0: 42, F1: (fi49 / (((fi52 << ((v9) & 63))) | 1))}
                                        if true {
                                            var v60 bool = (!true)
                                            _ = v60
                                        } else {
                                            var v61 uint32 = 2147483648
                                            _ = v61
                                        }
                                        var v62 uint = 4294967295
                                        _ = v59
                                        _ = v62
                                    }
                                }
                                v63 := GenS1{F0: 0}
                                v9 = (^(v9))
                                _ = v63
                            }
                            switch "foo bar" {
                            case "\n\t":
                                var v64 uint32 = 1
                                _ = v64
                            case "foo bar":
                                v65 := GenS8{F0: 0}
                                _ = v65
                            default:
                                var v66 int16 = v50
                                _ = v66
                            }
                            switch fi52 {
                            case 32767:
                                v67 := v2[4294967295]
                                _ = v67
                            case -284:
                                if false {
                                    v68 := GenS9{F0: v9}
                                    _ = v68
                                }
                            case -798:
                                switch fi49 {
                                case 128:
                                    if false {
                                        var v69 int8 = -1
                                        switch (fi52 + fi52) {
                                        case 591:
                                            var v70 uint64 = 2
                                            _ = v70
                                        case 87:
                                            var v71 uint16 = 0
                                            v71 -= 255
                                            v50 = v50
                                            _ = v71
                                        }
                                        _ = v69
                                    } else {
                                        v72 := GenS1{F0: (^(v9))}
                                        _ = v72
                                    }
                                case 1:
                                    switch "foo bar" {
                                    case "":
                                        var v73 int8 = -74
                                        _ = v73
                                    default:
                                        v9 += 255
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        v50 = v50
                                    case "hello":
                                        v74 := GenS4{F0: v9, F1: 2}
                                        _ = v74
                                    case "foo bar":
                                        v50 = (v50 % ((v50) | 1))
                                    }
                                    delete(v2, 2)
                                default:
                                    if ((v9 >> ((v9) & 63)) == (^(v9))) {
                                        if ((65535 > 0) && false) {
                                            v75 := GenS9{F0: v9}
                                            _ = v75
                                        }
                                    }
                                }
                            default:
                                _ = len(v2)
                            }
                            _ = fi52
                        }
                        v9 = v9
                        _ = v51
                    } else {
                        for fi76 := 0; fi76 < 16; fi76++ {
                            v2[4294967295] = true
                            v9 = (v9 + v9)
                            var v77 int8 = 42
                            if (!false) {
                                switch 127 {
                                case 833:
                                    v78 := GenS10{F0: 4294967295, F1: (v9 * (v9 * v9))}
                                    _ = v78
                                case 0:
                                    fmt.Println(fi49, v77)
                                case 9223372036854775807:
                                    _ = len(v1)
                                default:
                                    v79 := GenS11{F0: 2, F1: "hello", F2: 42}
                                    _ = v79
                                }
                            }
                            _ = fi76
                            _ = v77
                        }
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                        var v80 int32 = -1
                        _ = v80
                    }
                    _ = v50
                }
            case "hello":
                v9 = (^(v9))
            }
            v9 = 255
            _ = fi49
        }
    }
    fmt.Println(v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
