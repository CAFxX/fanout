package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 string }
func genFn0() uint32 {
    if (65535 < 0) {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 string = "a"
            fmt.Println(v8)
            _ = fi7
            _ = v8
        }
        if false {
            var v9 int8 = 1
            v10 := []int{32768}
            for fi11 := 0; fi11 < 11; fi11++ {
                switch "a" {
                case "":
                    switch (^(fi11)) {
                    case -552:
                        var v12 string = "hello"
                        _ = v12
                    case -2147483648:
                        if true {
                            if false {
                                if len(v10) > 0 {
                                    v10[0] = fi11
                                }
                                var v13 uint8 = 1
                                _ = v13
                            }
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        } else {
                            switch fi11 {
                            case 910:
                                v14 := GenS1{F0: 42}
                                _ = v14
                            case 127:
                                fmt.Println(fi11, v9)
                            case 32767:
                                v10 = append(v10, fi11)
                            default:
                                v15 := GenS1{F0: 203}
                                _ = v15
                            }
                            v10 = append(v10, (fi11 << ((2) & 63)))
                            if (1.5 != 0.1) {
                                switch (("\n\t" + "") + "") {
                                case "a":
                                    fmt.Println(v9)
                                case "\n\t":
                                    for fi16 := 0; fi16 < 2; fi16++ {
                                        if (false || true) {
                                            switch ("a" + "a") {
                                            case "hello":
                                                var v17 uint64 = 18446744073709551615
                                                if (505 != 444) {
                                                    var v18 float64 = 0.0
                                                    var v19 uint8 = 255
                                                    _ = v18
                                                    _ = v19
                                                } else {
                                                    v17 += v17
                                                }
                                                var v20 int8 = v9
                                                if ((^(v20)) <= v9) {
                                                    v21 := GenS1{F0: 127}
                                                    _ = v21
                                                }
                                                _ = v17
                                                _ = v20
                                            default:
                                                if (4294967295 == 253) {
                                                    var v22 uint32 = 42
                                                    _ = v22
                                                }
                                            }
                                            v23 := GenS2{F0: 2, F1: (1 >= 4294967295)}
                                            _ = v23
                                        }
                                        _ = fi16
                                    }
                                case "foo bar":
                                    _ = len(v10)
                                }
                                v9 = (^(v9))
                                v9 += (-(v9))
                            } else {
                                v9 -= 100
                                for fi24 := 0; fi24 < 3; fi24++ {
                                    var v25 string = "foo bar"
                                    v26 := GenS3{F0: (v25 + v25), F1: 65535}
                                    v9 -= 127
                                    fmt.Println(fi24)
                                    v27 := GenS2{F0: 50, F1: false}
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                var v28 uint8 = 128
                                _ = v28
                            }
                        }
                    }
                case "\n\t":
                    var v29 float64 = (0.1 * (3.14159 + float64(3.14159)))
                    _ = v29
                }
                v30 := GenS2{F0: 2, F1: ("a" != ("" + "hello"))}
                _ = fi11
                _ = v30
            }
            _ = v10
            _ = v9
        } else {
            switch ("\n\t" + "\n\t") {
            case "":
                var v31 uint64 = 0
                _ = v31
            case "foo bar":
                var v32 bool = true
                _ = v32
            case "a":
                v33 := GenS2{F0: 127, F1: true}
                _ = v33
            }
            for fi34 := 0; fi34 < 18; fi34++ {
                v35 := []int{fi34}
                for fi36 := 0; fi36 < 3; fi36++ {
                    if (42 <= 631) {
                        if (((((1.5 / 1.5) + (-2.25 - 0.0)) + (1e10 + 3.14159)) > 3.14159) && (!true)) {
                            if (407 < 0) {
                                switch "hello" {
                                case "foo bar":
                                    var v37 int64 = 84
                                    _ = v37
                                case "a":
                                    switch (fi34 / (((fi34 * fi36)) | 1)) {
                                    case -9223372036854775808:
                                        if len(v35) >= 2 {
                                            v35 = v35[1:len(v35)]
                                        }
                                    default:
                                        var v38 int64 = 42
                                        _ = v38
                                    }
                                case "":
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(fi34, fi36)
                                    case "foo bar":
                                        if (0.1 < -2.25) {
                                            var v39 bool = (false || (!false))
                                            _ = v39
                                        }
                                    case "a":
                                        v40 := GenS4{F0: 42}
                                        _ = v40
                                    default:
                                        fmt.Println(fi36, fi34)
                                    }
                                default:
                                    if len(v35) >= 2 {
                                        v35 = v35[1:len(v35)]
                                    }
                                }
                                var v41 float64 = (1.5 + 0.1)
                                var v42 int64 = -677
                                _ = v41
                                _ = v42
                            }
                            var v43 int16 = 42
                            switch ((("a" + ("foo bar" + ("\n\t" + "\n\t"))) + (("hello" + "") + "a")) + "\n\t") {
                            case "":
                                v44 := GenS5{F0: 18446744073709551615}
                                _ = v44
                            case "a":
                                var v45 float64 = 0.0
                                _ = v45
                            default:
                                if (2 < 2) {
                                    switch fi36 {
                                    case 583:
                                        v43 -= (v43 - -769)
                                    case -364:
                                        switch "a" {
                                        case "hello":
                                            var v46 int32 = 429
                                            _ = v46
                                        default:
                                            v47 := GenS3{F0: "a", F1: 59}
                                            _ = v47
                                        }
                                    }
                                    v35 = append(v35, (fi34 | (fi34 >> ((0) & 63))))
                                    var v48 int16 = (^(v43))
                                    _ = v48
                                }
                            }
                            _ = v43
                        }
                        v49 := GenS3{F0: "a", F1: 1}
                        v35 = append(v35, 2147483647)
                        _ = v49
                    } else {
                        var v50 uint64 = 633
                        var v51 float64 = float64(0.1)
                        switch fi36 {
                        case 9223372036854775807:
                            fmt.Println(v51, fi34, fi36)
                        case 1:
                            fmt.Println(fi36, fi34, v51)
                        }
                        var v52 uint = 2
                        v51 -= (((0.0 * (((v51 - float64(0.1)) + (0.0 * v51)) - 1.5)) * 1.5) / (v51 * float64(1.5)))
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    _ = fi36
                }
                _ = fi34
                _ = v35
            }
            var v53 int32 = -24
            _ = v53
        }
    } else {
        if (0.1 < (0.1 / 1.5)) {
            if false {
                var v54 uint = 18446744073709551615
                var v55 uint32 = 2
                for fi56 := 0; fi56 < 2; fi56++ {
                    var v57 int = fi56
                    v58 := GenS5{F0: 18446744073709551615}
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                _ = v54
                _ = v55
            }
            var v59 int16 = 32767
            _ = v59
        }
    }
    switch 2147483647 {
    case -122:
        var v60 string = (("\n\t" + ("hello" + "")) + "a")
        _ = v60
    case -1:
        if true {
            var v61 float64 = -2.25
            _ = v61
        }
    default:
        if ((("\n\t" + "") + "foo bar") != "") {
            var v62 int = -830
            switch v62 {
            case 509:
                if false {
                    for fi63 := 0; fi63 < 10; fi63++ {
                        var v64 int = 32768
                        _ = fi63
                        _ = v64
                    }
                    switch v62 {
                    case 548:
                        for fi65 := 0; fi65 < 17; fi65++ {
                            if false {
                                fmt.Println(v62, fi65)
                                var v66 uint64 = 0
                                _ = v66
                            }
                            var v67 int16 = 1000
                            v68 := []int{}
                            _ = fi65
                            _ = v67
                            _ = v68
                        }
                    case 133:
                        switch ((("a" + "foo bar") + "a") + (("foo bar" + (("hello" + "") + "\n\t")) + (("a" + "a") + "foo bar"))) {
                        case "":
                            fmt.Println(v62)
                        }
                    case 128:
                        switch (v62 / ((1) | 1)) {
                        case 4294967295:
                            var v69 int64 = 0
                            _ = v69
                        case 952:
                            for fi70 := 0; fi70 < 3; fi70++ {
                                var v71 uint = 42
                                v62 = fi70
                                if true {
                                    var v72 uint16 = 2
                                    var v73 int32 = int32(v72)
                                    v62 -= v62
                                    _ = v72
                                    _ = v73
                                } else {
                                    var v74 int64 = -447
                                    _ = v74
                                }
                                _ = fi70
                                _ = v71
                            }
                        case -9223372036854775808:
                            if (false || true) {
                                var v75 float64 = ((0.0 / (0.1 / 1.0)) * (-2.25 / 1.0))
                                var v76 uint64 = 4294967295
                                v77 := GenS6{F0: (^(v62)), F1: 2}
                                _ = v75
                                _ = v76
                                _ = v77
                            } else {
                                if false {
                                    v78 := GenS5{F0: 1}
                                    _ = v78
                                } else {
                                    var v79 int = v62
                                    _ = v79
                                }
                                v62 = -202
                                if (255 < 41) {
                                    v62 = (v62 + v62)
                                    var v80 int = v62
                                    _ = v80
                                } else {
                                    switch "hello" {
                                    case "a":
                                        v81 := GenS1{F0: 212}
                                        _ = v81
                                    default:
                                        if (false || ((421 != 1) && false)) {
                                            var v82 int8 = 127
                                            _ = v82
                                        }
                                    }
                                    v83 := GenS5{F0: 596}
                                    _ = v83
                                }
                            }
                        }
                    }
                    v62 = v62
                }
            }
            _ = v62
        } else {
            v84 := []int{127}
            var v85 string = "\n\t"
            var v86 int32 = 42
            _ = v84
            _ = v85
            _ = v86
        }
    }
    return 4294967295
}
func genFn1(p0 uint, p1 uint, p2 bool) int32 {
    p0 += 42
    v87 := GenS1{F0: 0}
    switch 771 {
    case 850:
        v88 := []int{-2147483648, 1, -1, 366}
        _ = v88
    case 56:
        p0 = p0
    default:
        for fi89 := 0; fi89 < 11; fi89++ {
            v90 := GenS7{F0: "\n\t"}
            _ = fi89
            _ = v90
        }
    }
    if true {
        var v91 uint8 = 139
        _ = v91
    } else {
        p2 = (1 == -608)
        for fi92 := 0; fi92 < 12; fi92++ {
            for fi93 := 0; fi93 < 16; fi93++ {
                if false {
                    v94 := []int{fi93, fi92}
                    var v95 uint = (p1 << (((^(p1))) & 63))
                    _ = v94
                    _ = v95
                }
                _ = fi93
            }
            p1 = (p1 + p1)
            p0 += (^(p1))
            _ = fi92
        }
        v96 := []int{9223372036854775807, -9223372036854775808, -2147483648}
        _ = v96
    }
    v97 := GenS2{F0: 128, F1: (true || ("foo bar" != "a"))}
    _ = p0
    _ = p1
    _ = p2
    _ = v87
    _ = v97
    return 42
}

func main() {
    var v1 float32 = (-2.25 * -2.25)
    v2 := []uint32{2147483647}
    var v3 float32 = -2.25
    var v4 int64 = -769
    v5 := make(map[uint32]uint32)
    var v6 float32 = (1e10 / v3)
    fmt.Println(v3, v6, v1)
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v98 := GenS5{F0: 332}
    var v99 int32 = -603
    v2 = append(v2, 2147483648)
    v99 = v99
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch (("hello" + (("a" + (("" + "\n\t") + "\n\t")) + "hello")) + ("hello" + "hello")) {
    case "":
        var v100 int = 127
        _ = v100
    default:
        if (42 != 2147483648) {
            v1 -= v6
            v6 = (0.1 / v6)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
    _ = v99
}
