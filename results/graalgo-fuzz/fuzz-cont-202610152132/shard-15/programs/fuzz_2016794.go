package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 int8; F2 bool }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 int8 }

func main() {
    var v1 bool = true
    v2 := []int{32767}
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    v6 := make(map[int]uint32)
    delete(v4, "foo bar")
    fmt.Println(v1)
    v1 = v1
    v1 = (-2147483648 == -49)
    v5[""] = 123
    fmt.Println(v1)
    v1 = true
    v7 := GenS1{F0: ("\n\t" + "foo bar"), F1: 255, F2: -9223372036854775808}
    v2 = append(v2, 32768)
    v8 := GenS1{F0: "foo bar", F1: 0, F2: -218}
    fmt.Println(v1)
    v9 := GenS1{F0: (("" + ("foo bar" + ((("" + "\n\t") + "a") + "foo bar"))) + "hello"), F1: 42, F2: 128}
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "a" {
        case "foo bar":
            fmt.Println(fi10, fi10)
        case "hello":
            var v11 int8 = -1
            _ = v11
        case "\n\t":
            switch (fi10 * -326) {
            case -2147483648:
                for fi12 := 0; fi12 < 19; fi12++ {
                    if true {
                        v13 := GenS1{F0: "hello", F1: 0, F2: (fi12 % ((fi10) | 1))}
                        fmt.Println(v1, fi10)
                        v14 := GenS2{F0: 32767, F1: 461}
                        switch fi12 {
                        case 1:
                            var v15 uint = 0
                            _ = v15
                        case 32768:
                            fmt.Println(v1, fi12, v1)
                        case -2147483648:
                            if true {
                                var v16 float64 = 1e10
                                v1 = (399 <= 1)
                                _ = v16
                            } else {
                                v1 = true
                                switch "" {
                                case "\n\t":
                                    if (-1 <= 1) {
                                        if v1 {
                                            var v17 bool = true
                                            _ = v17
                                        }
                                    }
                                case "hello":
                                    switch fi10 {
                                    case 185:
                                        if (0 != -1) {
                                            if ((((3.14159 / float64(-2.25)) >= (-2.25 - -2.25)) || v1) && (v1 || (false && true))) {
                                                v18 := GenS2{F0: 668, F1: 718}
                                                fmt.Println(fi12, v1, fi12)
                                                _ = v18
                                            }
                                            v1 = v1
                                        }
                                    case 891:
                                        switch "hello" {
                                        case "":
                                            var v19 int16 = -1
                                            _ = v19
                                        default:
                                            v1 = v1
                                        }
                                    case 32768:
                                        switch "a" {
                                        case "hello":
                                            if (false || (!(fi10 > fi12))) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v20 uint8 = 0
                                                var v21 float32 = (1e10 * 3.14159)
                                                _ = v20
                                                _ = v21
                                            }
                                        case "":
                                            v22 := GenS3{F0: 170, F1: 866, F2: 111}
                                            _ = v22
                                        case "foo bar":
                                            var v23 bool = ((("foo bar" + "a") + "") <= "")
                                            _ = v23
                                        }
                                    }
                                case "a":
                                    var v24 uint32 = 2147483648
                                    _ = v24
                                default:
                                    fmt.Println(fi12, fi12)
                                }
                            }
                        default:
                            var v25 int = -1
                            _ = v25
                        }
                        v26 := GenS1{F0: ((("\n\t" + (((("" + "hello") + (("hello" + "foo bar") + (("foo bar" + "foo bar") + "hello"))) + "a") + ("" + "\n\t"))) + ("a" + (("\n\t" + "foo bar") + ""))) + ""), F1: 0, F2: fi12}
                        _ = v13
                        _ = v14
                        _ = v26
                    } else {
                        var v27 int = (fi12 ^ 2147483647)
                        _ = v27
                    }
                    var v28 int8 = 124
                    v28 -= (v28 - (v28 >> ((42) & 63)))
                    _ = fi12
                    _ = v28
                }
            default:
                if (!true) {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v30 := GenS1{F0: ("a" + ("\n\t" + ("a" + "\n\t"))), F1: 0, F2: 127}
                        if (2147483647 != 2147483647) {
                            switch fi10 {
                            case 107:
                                var v31 float64 = 1e10
                                _ = v31
                            default:
                                var v32 int64 = -9223372036854775808
                                _ = v32
                            }
                        } else {
                            v1 = (-1 <= 799)
                        }
                        delete(v4, (("" + "a") + "\n\t"))
                        _ = fi29
                        _ = v30
                    }
                    for fi33 := 0; fi33 < 1; fi33++ {
                        if (!true) {
                            var v34 float64 = 0.0
                            v1 = (!v1)
                            _ = v34
                        } else {
                            fmt.Println(v1)
                        }
                        _ = fi33
                    }
                    switch "" {
                    case "hello":
                        fmt.Println(fi10, fi10)
                    case "foo bar":
                        switch "\n\t" {
                        case "a":
                            var v35 int64 = 1
                            _ = v35
                        case "\n\t":
                            if (2 == 179) {
                                v36 := GenS4{F0: 106}
                                _ = v36
                            }
                        default:
                            v1 = (((525 <= 0) && false) || v1)
                        }
                    case "":
                        var v37 uint32 = 0
                        _ = v37
                    default:
                        for fi38 := 0; fi38 < 13; fi38++ {
                            var v39 int = (fi38 >> ((260) & 63))
                            var v40 float64 = 1.5
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                    }
                } else {
                    switch 876 {
                    case 476:
                        if v1 {
                            for fi41 := 0; fi41 < 16; fi41++ {
                                var v42 bool = (false && (((((3.14159 * -2.25) * 0.0) == 1.5) && v1) || false))
                                switch (^(fi10)) {
                                case 2147483647:
                                    v43 := GenS3{F0: 2147483648, F1: -9223372036854775808, F2: -128}
                                    _ = v43
                                }
                                _ = fi41
                                _ = v42
                            }
                        }
                    case 1:
                        var v44 int64 = -253
                        _ = v44
                    case -829:
                        if (!false) {
                            for fi45 := 0; fi45 < 5; fi45++ {
                                v46 := GenS4{F0: 67}
                                _ = fi45
                                _ = v46
                            }
                        } else {
                            v1 = v1
                        }
                    default:
                        for fi47 := 0; fi47 < 3; fi47++ {
                            v1 = (2147483647 != 740)
                            v1 = v1
                            if false {
                                switch "\n\t" {
                                case "a":
                                    var v48 uint16 = 381
                                    _ = v48
                                }
                                v1 = true
                            } else {
                                v49 := GenS3{F0: 859, F1: 42, F2: -53}
                                var v50 uint32 = 65535
                                var v51 int8 = 127
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            _ = fi47
                        }
                    }
                    v1 = v1
                    if v1 {
                        if (!v1) {
                            var v52 uint16 = 42
                            if ((!v1) || (((0.0 * (-2.25 / ((1e10 - 3.14159) + (3.14159 / 1.5)))) > 3.14159) || ((("" + "") + ("" + ("hello" + "\n\t"))) >= ("a" + "hello")))) {
                                v1 = (255 <= 255)
                                for fi53 := 0; fi53 < 1; fi53++ {
                                    if ((v52 ^ (v52 * v52)) > v52) {
                                        var v54 int16 = 147
                                        _ = v54
                                    }
                                    if false {
                                        if ((!((127 > 9) || (!(2147483647 != -23)))) && ((-2.25 + 0.1) <= 1e10)) {
                                            if (!((1e10 > ((-2.25 / (-2.25 - 1.5)) + -2.25)) || v1)) {
                                                v1 = ((!v1) || (42 < 288))
                                                switch (("a" + "foo bar") + "hello") {
                                                case "":
                                                    v5["\n\t"] = 2147483648
                                                case "hello":
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi53 - (fi10 % ((fi10) | 1)))
                                                    }
                                                case "foo bar":
                                                    v2 = append(v2, fi10)
                                                default:
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v55 bool = (1.5 <= (1e10 + 1.5))
                                                        _ = v55
                                                    default:
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v52 = (^(v52))
                                                        case "\n\t":
                                                            fmt.Println(fi10, fi10)
                                                        }
                                                    }
                                                }
                                            }
                                            var v56 int8 = 0
                                            var v57 string = "\n\t"
                                            _ = v56
                                            _ = v57
                                        } else {
                                            var v58 uint8 = 2
                                            _ = v58
                                        }
                                        if v1 {
                                            v1 = (32767 == -35)
                                            var v59 float32 = -2.25
                                            v52 += (v52 + (v52 - (v52 / ((630) | 1))))
                                            _ = v59
                                        }
                                    } else {
                                        fmt.Println(v1, fi53, fi53)
                                        switch int(v52) {
                                        case 1:
                                            v52 = v52
                                        default:
                                            var v60 int8 = -128
                                            _ = v60
                                        }
                                    }
                                    switch fi10 {
                                    case -102:
                                        if ((3.14159 + (float64(float32(-2.25)) * 1e10)) == (1.5 * float64(-2.25))) {
                                            fmt.Println(fi53, fi10, fi10)
                                            var v61 uint8 = 15
                                            fmt.Println(fi10, fi10, fi10)
                                            _ = v61
                                        } else {
                                            v62 := GenS3{F0: 1, F1: 42, F2: int8(v52)}
                                            _ = v62
                                        }
                                    case 2147483647:
                                        v52 = (^(v52))
                                    case 32768:
                                        v63 := GenS3{F0: 65535, F1: -9223372036854775808, F2: 1}
                                        _ = v63
                                    default:
                                        var v64 int64 = 1
                                        _ = v64
                                    }
                                    _ = fi53
                                }
                            } else {
                                for fi65 := 0; fi65 < 1; fi65++ {
                                    v1 = false
                                    _ = fi65
                                }
                                var v66 float64 = ((float64(float32(1.5)) + 0.1) - -2.25)
                                switch "\n\t" {
                                case "":
                                    v52 = v52
                                case "a":
                                    if (1000 <= 693) {
                                        var v67 uint64 = 138
                                        _ = v67
                                    }
                                case "\n\t":
                                    var v68 uint32 = 65535
                                    _ = v68
                                default:
                                    var v69 uint32 = 65535
                                    _ = v69
                                }
                                _ = v66
                            }
                            _ = v52
                        }
                        if v1 {
                            v70 := GenS3{F0: 810, F1: -952, F2: 42}
                            if true {
                                for fi71 := 0; fi71 < 4; fi71++ {
                                    if v1 {
                                        v1 = v1
                                        var v72 uint16 = 42
                                        var v73 int32 = -2147483648
                                        _ = v72
                                        _ = v73
                                    }
                                    _ = fi71
                                }
                                v1 = v1
                            } else {
                                fmt.Println(fi10, v1)
                            }
                            if (!((((0.0 + float32(-2.25)) + float32((-2.25 * 0.0))) / (1.5 + 0.0)) != float32(float64((1.5 * float32(float64(0.1))))))) {
                                var v74 uint64 = 355
                                v74 = (v74 ^ (^(v74)))
                                _ = v74
                            }
                            _ = v70
                        }
                    } else {
                        if (0.0 != (0.0 / 1.0)) {
                            for fi75 := 0; fi75 < 9; fi75++ {
                                if (v1 || (18446744073709551615 < 18446744073709551615)) {
                                    var v76 int8 = 100
                                    _ = v76
                                }
                                if (-128 != -48) {
                                    var v77 float64 = 1.5
                                    if (!(v1 && (!v1))) {
                                        fmt.Println(fi75, v77, fi75)
                                        if false {
                                            v77 -= (v77 - float64(0.1))
                                        }
                                        fmt.Println(v77, fi10)
                                    }
                                    _ = v77
                                } else {
                                    var v78 bool = (!(9223372036854775807 >= 0))
                                    _ = v78
                                }
                                var v79 int16 = -815
                                _ = fi75
                                _ = v79
                            }
                            if true {
                                fmt.Println(fi10, v1)
                                v80 := GenS5{F0: 127, F1: -91, F2: v1}
                                for fi81 := 0; fi81 < 7; fi81++ {
                                    if true {
                                        var v82 string = "foo bar"
                                        var v83 float32 = float32((((-2.25 + 1e10) + 1.5) / ((0.1 - (1e10 - 0.1)) - -2.25)))
                                        v84 := GenS2{F0: 42, F1: 1}
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            v1 = (32767 <= -32768)
                                        case "hello":
                                            var v85 uint = 178
                                            _ = v85
                                        }
                                        var v86 float32 = 1.5
                                        if v1 {
                                            switch fi10 {
                                            case -263:
                                                var v87 int32 = 0
                                                _ = v87
                                            case 1:
                                                v88 := GenS6{F0: (fi81 / ((-456) | 1)), F1: fi81}
                                                _ = v88
                                            case 127:
                                                _ = len(v5)
                                            }
                                        } else {
                                            var v89 int8 = -1
                                            _ = v89
                                        }
                                        _ = v86
                                    }
                                    if false {
                                        _ = len(v2)
                                    }
                                    v1 = (1 == 564)
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v90 uint = 907
                                        _ = v90
                                    case "\n\t":
                                        var v91 float32 = (float32(3.14159) * 1.5)
                                        _ = v91
                                    case "":
                                        v5[("\n\t" + "\n\t")] = 4294967295
                                    }
                                    var v92 uint8 = 128
                                    v93 := GenS7{F0: int16(fi81)}
                                    _ = fi81
                                    _ = v92
                                    _ = v93
                                }
                                _ = v80
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            v94 := GenS8{F0: 65535, F1: (18446744073709551615 <= 42)}
                            var v95 string = "\n\t"
                            _ = v94
                            _ = v95
                        }
                        v1 = (0 != 4294967295)
                        v3[4294967295] = 2
                    }
                }
            }
        default:
            var v96 uint32 = 1
            _ = v96
        }
        for fi97 := 0; fi97 < 3; fi97++ {
            fmt.Println(fi10)
            v98 := GenS8{F0: 65535, F1: (v1 && (true && v1))}
            v99 := v6[fi97]
            _ = fi97
            _ = v98
            _ = v99
        }
        var v100 uint32 = 2147483648
        _ = fi10
        _ = v100
    }
    v1 = (580 >= 415)
    if true {
        v1 = (-1 < 0)
        var v101 int32 = 2147483647
        fmt.Println(v101, v101)
        _ = v101
    }
    switch "" {
    case "hello":
        fmt.Println(v1)
    case "a":
        v1 = true
    }
    v102 := GenS9{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
