package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 uint64; F1 int8; F2 int }
type GenS7 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 bool; F2 string }

func main() {
    var v1 uint32 = 42
    var v2 int16 = 137
    var v3 int64 = -9223372036854775808
    v4 := []string{("\n\t" + ("foo bar" + ("" + ("" + "a"))))}
    var v5 uint16 = 42
    var v6 uint16 = (v5 ^ (^(v5)))
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint8 = uint8(v1)
            var v10 string = "hello"
            v11 := GenS1{F0: 0, F1: 65535}
            if (!true) {
                v6 = v6
            }
            if len(v4) > 0 {
                v4[0] = "hello"
            }
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        v6 -= v5
        v12 := GenS2{F0: 0, F1: 1}
        _ = fi7
        _ = v12
    }
    if true {
        if (int32(v3) < 1) {
            if false {
                var v13 float32 = 3.14159
                _ = v13
            } else {
                switch "a" {
                case "hello":
                    var v14 int16 = 1000
                    _ = v14
                case "":
                    var v15 int16 = (v2 & (v2 | 1000))
                    _ = v15
                case "a":
                    _ = len(v4)
                }
            }
            if (0.1 < (((float32(3.14159) + 3.14159) + 1e10) - -2.25)) {
                if (true || (138 != 255)) {
                    v5 -= uint16(v1)
                    v1 -= (v1 | (v1 % ((v1) | 1)))
                } else {
                    v2 = 0
                    fmt.Println(v5, v5)
                }
            } else {
                var v16 uint32 = v1
                v17 := GenS2{F0: (-(v2)), F1: 18446744073709551615}
                var v18 string = ("\n\t" + "foo bar")
                _ = v16
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 bool = true
                _ = fi19
                _ = v20
            }
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                if (false && true) {
                    var v22 float64 = 3.14159
                    v23 := GenS1{F0: -2147483648, F1: v1}
                    var v24 uint32 = (v1 >> ((18446744073709551615) & 63))
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    switch -2147483648 {
                    case 768:
                        v25 := GenS3{F0: fi21, F1: 42}
                        _ = v25
                    default:
                        if ("hello" <= (("" + "foo bar") + "a")) {
                            if true {
                                switch 9223372036854775807 {
                                case 32767:
                                    for fi26 := 0; fi26 < 9; fi26++ {
                                        if false {
                                            v1 -= v1
                                            v5 -= (^(v6))
                                            switch (("" + "\n\t") + "hello") {
                                            case "foo bar":
                                                var v27 float32 = 1.5
                                                _ = v27
                                            case "hello":
                                                v1 = v1
                                            case "\n\t":
                                                var v28 int8 = 0
                                                _ = v28
                                            default:
                                                if ((true || false) || false) {
                                                    if true {
                                                        v2 += (v2 ^ v2)
                                                        fmt.Println(v3, v6)
                                                    }
                                                    var v29 uint16 = 679
                                                    _ = v29
                                                } else {
                                                    var v30 uint8 = 255
                                                    if (true || ((v5 & 42) > (v5 & v5))) {
                                                        var v31 uint16 = v5
                                                        if (false || false) {
                                                            v32 := GenS2{F0: 32767, F1: 0}
                                                            var v33 uint64 = 18446744073709551615
                                                            fmt.Println(fi21)
                                                            _ = v32
                                                            _ = v33
                                                        } else {
                                                            switch (fi21 + (fi26 % (((fi21 << ((uint(v30)) & 63))) | 1))) {
                                                            case 736:
                                                                _ = len(v4)
                                                            case 127:
                                                                v34 := GenS2{F0: v2, F1: 775}
                                                                _ = v34
                                                            default:
                                                                fmt.Println(v6, v1, v3)
                                                            }
                                                            var v35 uint32 = 773
                                                            switch "a" {
                                                            case "":
                                                                if false {
                                                                    v36 := GenS3{F0: (fi21 & (-(fi26))), F1: uint8(v35)}
                                                                    v1 = v35
                                                                    v31 -= (v5 << ((4294967295) & 63))
                                                                    _ = v36
                                                                }
                                                            case "\n\t":
                                                                var v37 uint8 = (^(v30))
                                                                _ = v37
                                                            default:
                                                                v1 = 0
                                                            }
                                                            _ = v35
                                                        }
                                                        var v38 int8 = 100
                                                        _ = v31
                                                        _ = v38
                                                    } else {
                                                        v39 := GenS4{F0: 42, F1: 42, F2: 599}
                                                        if false {
                                                            if len(v4) > 0 {
                                                                v4[0] = "\n\t"
                                                            }
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            fmt.Println(fi26, fi26)
                                                        } else {
                                                            var v40 int64 = -991
                                                            _ = v40
                                                        }
                                                        _ = v39
                                                    }
                                                    _ = v30
                                                }
                                                v5 = (v5 << ((255) & 63))
                                            }
                                        } else {
                                            if (!(true || (4294967295 > 18446744073709551615))) {
                                                var v41 float32 = ((0.1 - 3.14159) - 0.1)
                                                _ = v41
                                            } else {
                                                v4 = append(v4, (("\n\t" + "\n\t") + "a"))
                                                var v42 bool = (!(true || true))
                                                _ = v42
                                            }
                                            var v43 float32 = 0.1
                                            _ = v43
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = "a"
                                        }
                                        var v44 int8 = 0
                                        _ = fi26
                                        _ = v44
                                    }
                                case 127:
                                    if true {
                                        var v45 string = "\n\t"
                                        v45 = (v45 + v45)
                                        var v46 float64 = (1.5 / (((0.1 - 0.0) + float64(-2.25)) - 1e10))
                                        _ = v45
                                        _ = v46
                                    }
                                case 4294967295:
                                    var v47 uint16 = v5
                                    _ = v47
                                }
                            } else {
                                fmt.Println(v2)
                            }
                            if ((!(2 > 128)) || true) {
                                switch ("hello" + "") {
                                case "a":
                                    if len(v4) > 0 {
                                        v4[0] = "a"
                                    }
                                case "":
                                    if len(v4) > 0 {
                                        v4[0] = "a"
                                    }
                                }
                                v48 := GenS1{F0: 0, F1: (v1 + (v1 >> ((2) & 63)))}
                                _ = v48
                            }
                        }
                    }
                }
                v5 = 1
                for fi49 := 0; fi49 < 10; fi49++ {
                    switch 2147483647 {
                    case -521:
                        v50 := GenS5{F0: (^(v1)), F1: 0}
                        _ = v50
                    case 128:
                        var v51 float64 = -2.25
                        _ = v51
                    case 42:
                        switch ("foo bar" + "hello") {
                        case "":
                            v52 := GenS2{F0: v2, F1: 1}
                            switch ("foo bar" + ("\n\t" + "")) {
                            case "\n\t":
                                fmt.Println(v6, fi21)
                            }
                            switch fi21 {
                            case 358:
                                var v53 uint8 = 0
                                _ = v53
                            default:
                                if true {
                                    fmt.Println(v5, v1)
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if (42 < 100) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v54 uint64 = 4294967295
                                        _ = v54
                                    } else {
                                        if true {
                                            var v55 int = fi21
                                            if (false || false) {
                                                fmt.Println(v5, v5, v6)
                                                v5 = v5
                                                fmt.Println(v3, fi49)
                                            } else {
                                                var v56 int32 = -2147483648
                                                _ = v56
                                            }
                                            var v57 int32 = 42
                                            _ = v55
                                            _ = v57
                                        }
                                    }
                                    var v58 bool = (!true)
                                    var v59 float64 = (float64((-2.25 * 1e10)) + 0.1)
                                    v6 -= 484
                                    _ = v58
                                    _ = v59
                                }
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                                v60 := GenS4{F0: 42, F1: uint(v1), F2: 42}
                                _ = v60
                            }
                            _ = v52
                        default:
                            if true {
                                if len(v4) > 0 {
                                    v4[0] = (("foo bar" + "") + "hello")
                                }
                                switch int(v1) {
                                case 2147483647:
                                    var v61 float64 = 0.0
                                    _ = v61
                                case 32768:
                                    v62 := GenS6{F0: 18446744073709551615, F1: 100, F2: (fi49 * fi49)}
                                    _ = v62
                                case 127:
                                    var v63 string = ("hello" + "foo bar")
                                    _ = v63
                                default:
                                    fmt.Println(fi49, v2, v5)
                                }
                            }
                        }
                    }
                    v64 := GenS6{F0: 0, F1: 42, F2: fi21}
                    _ = fi49
                    _ = v64
                }
                _ = fi21
            }
        }
    } else {
        for fi65 := 0; fi65 < 1; fi65++ {
            fmt.Println(v6, v2)
            v66 := GenS6{F0: 18446744073709551615, F1: -128, F2: 128}
            for fi67 := 0; fi67 < 12; fi67++ {
                fmt.Println(v3, v3)
                _ = fi67
            }
            _ = fi65
            _ = v66
        }
        var v68 uint = 0
        v2 += (v2 - v2)
        _ = v68
    }
    switch 4294967295 {
    case -9223372036854775808:
        if (v3 != v3) {
            if false {
                fmt.Println(v1)
                fmt.Println(v5, v2, v5)
                var v69 uint16 = (v5 & (v5 << ((uint(v1)) & 63)))
                _ = v69
            }
            v70 := GenS5{F0: (v1 << ((902) & 63)), F1: v2}
            var v71 uint = 47
            _ = v70
            _ = v71
        }
    default:
        switch 193 {
        case 724:
            v2 += 42
        default:
            if false {
                v2 += v2
                switch "a" {
                case "hello":
                    var v72 bool = true
                    _ = v72
                case "foo bar":
                    v4 = append(v4, "a")
                }
            } else {
                fmt.Println(v6, v6)
                fmt.Println(v1)
                v4 = append(v4, "foo bar")
            }
        }
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 0 {
    case 1:
        fmt.Println(v1)
    case 636:
        v73 := GenS6{F0: 781, F1: 42, F2: -148}
        _ = v73
    case -2147483648:
        fmt.Println(v6, v6, v5)
    default:
        fmt.Println(v5)
    }
    if ((1 != -184) && true) {
        switch "\n\t" {
        case "hello":
            fmt.Println(v6, v1)
        case "\n\t":
            v5 = v6
        case "foo bar":
            v4 = append(v4, "hello")
        }
    } else {
        v74 := GenS7{F0: (v1 / ((42) | 1)), F1: (v1 | v1), F2: (v2 | (v2 % ((v2) | 1)))}
        var v75 int = 4294967295
        if (v2 > v2) {
            var v76 int8 = -1
            _ = len(v4)
            switch ("" + ("a" + ("foo bar" + ((((("\n\t" + (((("foo bar" + ("foo bar" + "a")) + ((((("\n\t" + "foo bar") + (("hello" + ("foo bar" + "a")) + "\n\t")) + ("hello" + ("foo bar" + ("hello" + "foo bar")))) + (("" + "foo bar") + (((("\n\t" + ("foo bar" + "a")) + ("foo bar" + "hello")) + "\n\t") + "\n\t"))) + (("a" + "a") + "a"))) + (("hello" + "foo bar") + ("a" + "\n\t"))) + "hello")) + "hello") + "hello") + ("a" + ("" + ("a" + (("\n\t" + (("\n\t" + "hello") + ("" + "a"))) + ("foo bar" + "\n\t")))))) + "foo bar")))) {
            case "a":
                v76 = v76
            case "":
                var v77 float64 = (float64((((0.1 * 0.1) / float32(float64(float32(((3.14159 * 1.5) * -2.25))))) * (1.5 - 0.1))) - float64(1.5))
                _ = v77
            default:
                for fi78 := 0; fi78 < 12; fi78++ {
                    switch "a" {
                    case "":
                        var v79 uint8 = 42
                        _ = v79
                    case "foo bar":
                        fmt.Println(v6)
                    }
                    fmt.Println(v2)
                    v80 := GenS5{F0: (v1 / ((v1) | 1)), F1: (v2 + -582)}
                    _ = fi78
                    _ = v80
                }
            }
            _ = v76
        }
        _ = v74
        _ = v75
    }
    if (255 <= 65535) {
        if (1 <= 798) {
            for fi81 := 0; fi81 < 17; fi81++ {
                fmt.Println(v2)
                switch fi81 {
                case 32768:
                    v82 := GenS8{F0: v3}
                    _ = v82
                case -2147483648:
                    v83 := GenS4{F0: 599, F1: 18446744073709551615, F2: 823}
                    _ = v83
                }
                _ = fi81
            }
            var v84 uint8 = 2
            switch ("a" + "\n\t") {
            case "foo bar":
                v6 += v5
            case "a":
                for fi85 := 0; fi85 < 2; fi85++ {
                    v3 += v3
                    switch "foo bar" {
                    case "\n\t":
                        v4 = append(v4, "\n\t")
                    }
                    _ = fi85
                }
            default:
                var v86 int8 = 100
                _ = v86
            }
            _ = v84
        }
        if (false || true) {
            v87 := GenS3{F0: -902, F1: uint8(v1)}
            _ = v87
        } else {
            switch 463 {
            case 623:
                v2 = (-(v2))
            case -1:
                var v88 int16 = 1
                _ = v88
            case -304:
                for fi89 := 0; fi89 < 7; fi89++ {
                    for fi90 := 0; fi90 < 10; fi90++ {
                        v3 = v3
                        v5 = v5
                        if false {
                            var v91 uint32 = v1
                            switch "" {
                            case "a":
                                v2 = v2
                            case "foo bar":
                                if true {
                                    fmt.Println(v91, v2, fi89)
                                    if ("" >= "hello") {
                                        var v92 uint = 65535
                                        fmt.Println(fi90)
                                        var v93 int8 = 1
                                        _ = v92
                                        _ = v93
                                    } else {
                                        if true {
                                            v94 := GenS9{F0: 42, F1: (v6 <= 0), F2: (("a" + "") + "a")}
                                            _ = v94
                                        }
                                        var v95 bool = true
                                        v96 := GenS5{F0: 26, F1: (v2 >> ((2) & 63))}
                                        _ = v95
                                        _ = v96
                                    }
                                } else {
                                    var v97 string = ("" + "foo bar")
                                    _ = v97
                                }
                            default:
                                switch "" {
                                case "\n\t":
                                    _ = len(v4)
                                }
                            }
                            switch fi89 {
                            case 32767:
                                if false {
                                    if ((true || false) && ((v91 >> ((65535) & 63)) <= v91)) {
                                        var v98 int8 = 42
                                        _ = v98
                                    } else {
                                        var v99 uint64 = 4294967295
                                        v4 = append(v4, ("" + ""))
                                        if true {
                                            v100 := GenS6{F0: v99, F1: -128, F2: fi89}
                                            switch ("" + "") {
                                            case "foo bar":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "":
                                                if ((-(v2)) < int16(fi89)) {
                                                    switch fi89 {
                                                    case 1:
                                                        var v101 string = "hello"
                                                        fmt.Println(v101)
                                                        var v102 uint32 = (v91 | v1)
                                                        _ = len(v4)
                                                        fmt.Println(v5, fi89)
                                                        if (!((((v2 * 174) <= v2) || true) || (v101 == v101))) {
                                                            if (!false) {
                                                                if true {
                                                                    var v103 int64 = v3
                                                                    v4 = append(v4, v101)
                                                                    _ = v103
                                                                }
                                                                fmt.Println(v102)
                                                                v4 = append(v4, (v101 + (v101 + (v101 + ("\n\t" + "\n\t")))))
                                                            }
                                                            v101 = ("a" + ("foo bar" + v101))
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v101
                                                        _ = v102
                                                    case 128:
                                                        v104 := GenS8{F0: (v3 ^ (v3 * v3))}
                                                        _ = v104
                                                    }
                                                } else {
                                                    var v105 int32 = -74
                                                    _ = v105
                                                }
                                            }
                                            fmt.Println(v1, v99)
                                            _ = v100
                                        } else {
                                            var v106 uint8 = 128
                                            v1 = v91
                                            _ = v106
                                        }
                                        v2 = (v2 ^ (v2 + (v2 + v2)))
                                        _ = v99
                                    }
                                } else {
                                    fmt.Println(v1)
                                    if ((true || ((fi89 >= fi89) || false)) && (true || true)) {
                                        v4 = append(v4, "hello")
                                    }
                                }
                            case 1:
                                v2 = 1
                                var v107 uint16 = 42
                                if true {
                                    var v108 uint = 807
                                    var v109 int16 = 0
                                    _ = len(v4)
                                    _ = v108
                                    _ = v109
                                } else {
                                    if true {
                                        v5 = 32767
                                        if true {
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v107, v5, v5)
                                                var v110 bool = (0 >= 670)
                                                _ = v110
                                            default:
                                                v5 = (v107 ^ (^(v5)))
                                            }
                                            v3 -= (v3 % (((v3 & (v3 >> ((0) & 63)))) | 1))
                                            v111 := GenS9{F0: v6, F1: true, F2: ((("\n\t" + "") + "\n\t") + "a")}
                                            _ = v111
                                        }
                                    } else {
                                        switch (fi89 * (fi89 << ((2) & 63))) {
                                        case 280:
                                            v4 = append(v4, "\n\t")
                                        }
                                        v112 := GenS9{F0: v107, F1: false, F2: "hello"}
                                        _ = v112
                                    }
                                }
                                v91 = 1
                                _ = v107
                            }
                            if len(v4) > 0 {
                                v4[0] = ("foo bar" + "hello")
                            }
                            _ = v91
                        } else {
                            v2 = (v2 << ((4294967295) & 63))
                        }
                        _ = fi90
                    }
                    _ = fi89
                }
            }
            v3 -= 42
            fmt.Println(v5, v3)
        }
        v113 := GenS3{F0: -2147483648, F1: 42}
        _ = v113
    }
    switch 32768 {
    case 32768:
        fmt.Println(v6)
    case 32767:
        if len(v4) > 0 {
            v4[0] = ("hello" + "\n\t")
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
