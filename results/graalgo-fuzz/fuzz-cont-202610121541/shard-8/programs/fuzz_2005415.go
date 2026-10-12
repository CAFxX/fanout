package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int8; F1 int32; F2 uint16 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int8; F1 int; F2 uint64 }
type GenS6 struct { F0 int32; F1 int; F2 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS9 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    v3 := make(map[int]string)
    var v4 uint8 = 1
    v5 := make(map[uint32]int)
    var v6 string = "hello"
    v2 = (!v2)
    if v2 {
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(v2)
            _ = fi7
        }
        if v2 {
            var v8 uint16 = 162
            var v9 float32 = 1e10
            v6 = (("" + ("foo bar" + v6)) + "foo bar")
            _ = v8
            _ = v9
        } else {
            var v10 bool = (-1 != -712)
            _ = v10
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            v6 = v6
            _ = fi11
        }
    }
    var v12 float64 = -2.25
    if v2 {
        var v13 bool = (!v2)
        switch v6 {
        case "a":
            v14 := []int{128, 9223372036854775807, 9223372036854775807}
            _ = v14
        }
        _ = v13
    } else {
        var v15 int = 450
        fmt.Println(v12, v12, v2)
        switch (v6 + "hello") {
        case "":
            v2 = (!(!true))
        case "a":
            v16 := GenS1{F0: ""}
            _ = v16
        case "\n\t":
            v5[65535] = v15
        }
        _ = v15
    }
    if false {
        if v2 {
            v6 = (((((((v6 + ((v6 + "\n\t") + ("a" + "\n\t"))) + v6) + ("\n\t" + v6)) + "hello") + (((v6 + "hello") + (((v6 + (v6 + v6)) + "a") + v6)) + "foo bar")) + (v6 + ("" + (v6 + (("foo bar" + "\n\t") + ((v6 + v6) + v6)))))) + v6)
        } else {
            for fi17 := 0; fi17 < 15; fi17++ {
                switch v6 {
                case "\n\t":
                    if true {
                        for fi18 := 0; fi18 < 14; fi18++ {
                            v19 := GenS1{F0: v6}
                            _ = fi18
                            _ = v19
                        }
                        var v20 int8 = 0
                        _ = v20
                    }
                case "hello":
                    v21 := GenS1{F0: ("\n\t" + v6)}
                    _ = v21
                case "":
                    v12 += 3.14159
                default:
                    v22 := GenS1{F0: v6}
                    _ = v22
                }
                _ = fi17
            }
            var v23 int16 = 165
            for fi24 := 0; fi24 < 3; fi24++ {
                if (42 < 255) {
                    _ = len(v1)
                    for fi25 := 0; fi25 < 15; fi25++ {
                        v6 = v6
                        _ = fi25
                    }
                }
                _ = fi24
            }
            _ = v23
        }
        v12 -= 1.5
        var v26 int64 = 9223372036854775807
        _ = v26
    }
    v4 += v4
    switch 4294967295 {
    case -1:
        v2 = v2
    case -709:
        switch (v6 + v6) {
        case "":
            var v27 uint8 = 2
            _ = v27
        default:
            if (1e10 > 3.14159) {
                switch 589 {
                case 1:
                    switch 81 {
                    case 0:
                        if v2 {
                            v28 := GenS2{F0: -2147483648, F1: 75}
                            v29 := GenS3{F0: 0, F1: 1, F2: 255}
                            _ = v28
                            _ = v29
                        } else {
                            fmt.Println(v6, v4, v4)
                        }
                    case -714:
                        var v30 bool = false
                        _ = v30
                    default:
                        v2 = ((float64(0.0) / ((v12 / (((1e10 - v12) / float64(1e10)) / (float64(3.14159) * 0.1))) * float64((-2.25 * (0.0 + -2.25))))) != (1e10 / -2.25))
                    }
                case 128:
                    fmt.Println(v4, v4, v4)
                case 4294967295:
                    if (v4 == v4) {
                        v31 := GenS3{F0: 29, F1: -1, F2: 0}
                        _ = v31
                    } else {
                        var v32 uint16 = 32767
                        fmt.Println(v4, v4)
                        if v2 {
                            fmt.Println(v12, v32)
                            switch (v6 + "\n\t") {
                            case "hello":
                                var v33 uint32 = 1
                                _ = v33
                            case "a":
                                if (9223372036854775807 <= 2147483647) {
                                    var v34 int = 42
                                    switch (-(v34)) {
                                    case 42:
                                        switch "" {
                                        case "a":
                                            v35 := GenS4{F0: int32(v4), F1: ((v6 + ((("a" + ((("" + v6) + "a") + v6)) + v6) + "hello")) + (v6 + ((v6 + v6) + v6)))}
                                            _ = v35
                                        case "":
                                            v1[v34] = 403
                                        case "\n\t":
                                            v4 = (^(v4))
                                        default:
                                            for fi36 := 0; fi36 < 10; fi36++ {
                                                v12 -= 3.14159
                                                _ = fi36
                                            }
                                        }
                                    case 32767:
                                        if true {
                                            switch (v34 % (((v34 << ((82) & 63))) | 1)) {
                                            case 459:
                                                switch (-(v34)) {
                                                case -124:
                                                    switch (-(v34)) {
                                                    case -443:
                                                        v37 := []int{v34}
                                                        _ = v37
                                                    default:
                                                        switch (^(v34)) {
                                                        case -2147483648:
                                                            fmt.Println(v34, v6, v4)
                                                        case 2147483647:
                                                            for fi38 := 0; fi38 < 15; fi38++ {
                                                                v39 := []int{v34, v34, (fi38 & (v34 & (v34 + v34))), fi38}
                                                                for fi40 := 0; fi40 < 10; fi40++ {
                                                                    v41 := GenS4{F0: 0, F1: ("" + "a")}
                                                                    v4 += v4
                                                                    switch v6 {
                                                                    case "hello":
                                                                        v42 := GenS5{F0: 0, F1: fi38, F2: uint64(v4)}
                                                                        _ = v42
                                                                    default:
                                                                        var v43 uint8 = (v4 << ((360) & 63))
                                                                        _ = v43
                                                                    }
                                                                    _ = fi40
                                                                    _ = v41
                                                                }
                                                                if (!false) {
                                                                    v6 = "hello"
                                                                    v6 = v6
                                                                    if ((258 < 0) || (v6 > v6)) {
                                                                        if false {
                                                                            fmt.Println(v2)
                                                                        }
                                                                        v2 = v2
                                                                    }
                                                                }
                                                                _ = fi38
                                                                _ = v39
                                                            }
                                                        }
                                                    }
                                                case -348:
                                                    var v44 uint64 = 4294967295
                                                    _ = v44
                                                default:
                                                    fmt.Println(v2, v12, v4)
                                                }
                                            case 9223372036854775807:
                                                v45 := GenS1{F0: v6}
                                                _ = v45
                                            }
                                            v46 := GenS3{F0: 18, F1: -1, F2: 2}
                                            _ = v46
                                        }
                                    }
                                    if (!v2) {
                                        if (2 > 0) {
                                            v47 := []int{v34, int(v4), v34}
                                            v48 := GenS5{F0: int8(v4), F1: 1, F2: 2}
                                            v49 := GenS6{F0: 82, F1: 32768, F2: 1}
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        }
                                        v50 := []int{}
                                        _ = v50
                                    } else {
                                        var v51 int32 = 42
                                        if (4294967295 != 4294967295) {
                                            fmt.Println(v2, v2, v51)
                                            delete(v1, (^(v34)))
                                        } else {
                                            for fi52 := 0; fi52 < 20; fi52++ {
                                                delete(v3, fi52)
                                                _ = fi52
                                            }
                                        }
                                        _ = v51
                                    }
                                    _ = v34
                                }
                            default:
                                fmt.Println(v2, v4)
                            }
                            var v53 uint16 = (v32 >> ((361) & 63))
                            _ = v53
                        }
                        _ = v32
                    }
                default:
                    v54 := []int{}
                    _ = v54
                }
                var v55 uint8 = (v4 ^ v4)
                v56 := GenS6{F0: -1, F1: -674, F2: v4}
                _ = v55
                _ = v56
            }
        }
    }
    var v57 uint64 = 4294967295
    v12 += 1.5
    switch v6 {
    case "\n\t":
        fmt.Println(v57)
    case "a":
        if v2 {
            v58 := GenS7{F0: 182}
            for fi59 := 0; fi59 < 5; fi59++ {
                if v2 {
                    v2 = v2
                    if (42 == -1) {
                        v60 := []int{fi59, (fi59 / (((fi59 << ((65535) & 63))) | 1)), fi59}
                        v6 = "\n\t"
                        _ = v60
                    }
                }
                switch "a" {
                case "":
                    v61 := []int{(fi59 % ((fi59) | 1)), 0, (^(fi59))}
                    _ = v61
                case "hello":
                    var v62 int8 = 108
                    _ = v62
                default:
                    for fi63 := 0; fi63 < 12; fi63++ {
                        fmt.Println(v57)
                        v64 := GenS8{F0: 681, F1: 321, F2: 42}
                        v65 := GenS9{F0: 42}
                        v66 := GenS5{F0: 42, F1: fi63, F2: 486}
                        switch v6 {
                        case "a":
                            v67 := GenS7{F0: 720}
                            var v68 float64 = -2.25
                            _ = v67
                            _ = v68
                        case "\n\t":
                            if v2 {
                                v57 = (^(v57))
                                v69 := v1[fi59]
                                _ = v69
                            }
                        case "":
                            delete(v5, 2147483648)
                        default:
                            var v70 float64 = v12
                            _ = v70
                        }
                        fmt.Println(v12, v6, v12)
                        _ = fi63
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                }
                _ = fi59
            }
            for fi71 := 0; fi71 < 2; fi71++ {
                var v72 int64 = -323
                v73 := []int{(fi71 + fi71)}
                fmt.Println(v6, fi71)
                _ = fi71
                _ = v72
                _ = v73
            }
            _ = v58
        } else {
            v74 := GenS2{F0: 0, F1: 42}
            var v75 uint32 = 262
            switch v6 {
            case "foo bar":
                var v76 uint = 491
                _ = v76
            case "":
                v2 = v2
            default:
                var v77 int8 = 42
                _ = v77
            }
            _ = v74
            _ = v75
        }
    default:
        switch -2147483648 {
        case 32767:
            var v78 int8 = 1
            _ = v78
        }
    }
    var v79 uint32 = 926
    switch v6 {
    case "\n\t":
        v80 := GenS3{F0: 0, F1: 0, F2: 42}
        _ = v80
    case "a":
        var v81 uint8 = 42
        _ = v81
    default:
        var v82 uint32 = v79
        _ = v82
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v79
}
