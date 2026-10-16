package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64; F1 bool; F2 int8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS13 struct { F0 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 uint64 = 42
    v4 := []int{4294967295, 4294967295}
    v5 := []int{9223372036854775807}
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: true, F1: 2, F2: "a"}
        v8 := GenS1{F0: true, F1: 2, F2: "hello"}
        v3 = v3
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint16 = 138
    switch "\n\t" {
    case "":
        v3 = v3
    case "foo bar":
        var v10 int = 42
        _ = v10
    }
    switch 32768 {
    case 128:
        if len(v4) > 0 {
            v4[0] = -2147483648
        }
    default:
        var v11 float64 = ((-2.25 * 1e10) * 3.14159)
        _ = v11
    }
    if true {
        if ((v9 >> ((922) & 63)) < v9) {
            v5 = append(v5, 0)
            var v12 string = ("" + "a")
            _ = v12
        }
        if (116 != -804) {
            switch 1 {
            case -2147483648:
                fmt.Println(v9, v3)
            default:
                v9 = 32767
            }
            for fi13 := 0; fi13 < 7; fi13++ {
                delete(v2, "hello")
                _ = fi13
            }
        }
        if (((!true) && true) || ("a" <= "hello")) {
            _ = len(v4)
        }
    }
    v9 -= v9
    var v14 uint = 18446744073709551615
    if (2147483647 == 65535) {
        v15 := GenS2{F0: 32768, F1: 127, F2: 817}
        if false {
            v3 -= 0
            v14 -= uint(v3)
        } else {
            if (-119 == 0) {
                var v16 string = "a"
                if (2147483647 == -1) {
                    var v17 uint64 = uint64(v9)
                    var v18 int32 = -1
                    _ = v17
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi19
                    }
                    v16 = v16
                    var v20 int64 = -9223372036854775808
                    _ = v20
                }
                _ = v16
            }
        }
        var v21 float32 = ((((1e10 - (((1.5 / (1e10 + 1e10)) / 3.14159) * (((((((3.14159 + 1e10) / -2.25) * float32(1e10)) / 1e10) * 1e10) - (0.0 * 3.14159)) / (((1e10 / 3.14159) / 1.0) - 0.1)))) - 3.14159) + 0.0) * 3.14159)
        _ = v15
        _ = v21
    }
    _ = len(v2)
    v14 = (v14 / (((v14 & uint(v9))) | 1))
    fmt.Println(v14, v14)
    for fi22 := 0; fi22 < 4; fi22++ {
        if ((true && (!true)) || ((3.14159 >= (0.0 + 1e10)) || true)) {
            if (!(1.5 != 0.0)) {
                var v23 int = (fi22 % (((fi22 / ((2147483647) | 1))) | 1))
                _ = v23
            }
            v9 = (v9 * (v9 & (v9 << (((v14 << (((v14 + v14)) & 63))) & 63))))
        } else {
            fmt.Println(v3, v14)
        }
        _ = fi22
    }
    if false {
        switch "" {
        case "foo bar":
            var v24 int8 = int8(v14)
            _ = v24
        case "":
            var v25 string = ("hello" + "\n\t")
            _ = v25
        case "a":
            if ((v14 ^ (v14 % (((^(v14))) | 1))) <= 0) {
                switch "a" {
                case "hello":
                    var v26 int32 = 42
                    _ = v26
                }
            } else {
                for fi27 := 0; fi27 < 18; fi27++ {
                    v28 := GenS2{F0: fi27, F1: 2, F2: 42}
                    _ = fi27
                    _ = v28
                }
                if true {
                    for fi29 := 0; fi29 < 20; fi29++ {
                        v30 := GenS3{F0: -757}
                        switch ("\n\t" + "") {
                        case "hello":
                            if len(v5) > 0 {
                                v5[0] = 32767
                            }
                        case "\n\t":
                            var v31 uint8 = 140
                            _ = v31
                        case "foo bar":
                            for fi32 := 0; fi32 < 2; fi32++ {
                                v33 := GenS4{F0: 0}
                                switch (fi32 * 32767) {
                                case 543:
                                    v14 += (v14 >> (((v14 * (^(v14)))) & 63))
                                case 127:
                                    v4 = append(v4, (fi29 - fi29))
                                case -1:
                                    fmt.Println(v9)
                                default:
                                    v3 -= (^(v3))
                                    if false {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v14 -= v14
                                        v3 -= (v3 % ((42) | 1))
                                    } else {
                                        v9 = 221
                                    }
                                    v34 := v1[0]
                                    _ = v34
                                }
                                v35 := GenS5{F0: 924, F1: 2}
                                _ = fi32
                                _ = v33
                                _ = v35
                            }
                        default:
                            switch "" {
                            case "a":
                                switch fi29 {
                                case 32768:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case -552:
                                    v9 = v9
                                case 4294967295:
                                    for fi36 := 0; fi36 < 9; fi36++ {
                                        fmt.Println(v3)
                                        _ = fi36
                                    }
                                default:
                                    if (1e10 < 1e10) {
                                        var v37 uint8 = 255
                                        for fi38 := 0; fi38 < 19; fi38++ {
                                            v39 := GenS3{F0: 9223372036854775807}
                                            var v40 bool = true
                                            _ = fi38
                                            _ = v39
                                            _ = v40
                                        }
                                        delete(v1, fi29)
                                        _ = v37
                                    } else {
                                        var v41 uint32 = 1
                                        v2["\n\t"] = ((fi29 >= (fi29 - int(v14))) && true)
                                        var v42 int = fi29
                                        _ = v41
                                        _ = v42
                                    }
                                }
                            default:
                                if false {
                                    v14 = 446
                                    for fi43 := 0; fi43 < 3; fi43++ {
                                        var v44 uint32 = 2
                                        _ = fi43
                                        _ = v44
                                    }
                                    v14 += v14
                                } else {
                                    if (!(!false)) {
                                        delete(v1, (fi29 & (fi29 - fi29)))
                                        if (4294967295 > 65535) {
                                            if (v9 <= v9) {
                                                var v45 int64 = 9223372036854775807
                                                if false {
                                                    v46 := GenS6{F0: v3}
                                                    if true {
                                                        fmt.Println(v14)
                                                        var v47 uint32 = 4294967295
                                                        _ = v47
                                                    }
                                                    var v48 int32 = -657
                                                    _ = v46
                                                    _ = v48
                                                } else {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        for fi49 := 0; fi49 < 13; fi49++ {
                                                            fmt.Println(v14)
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            switch (fi29 >> ((v14) & 63)) {
                                                            case 2147483647:
                                                                fmt.Println(v3, v14)
                                                            case 1:
                                                                v9 = 65535
                                                            case -2147483648:
                                                                if (!(400 > -1)) {
                                                                    v1[(fi29 + 266)] = "hello"
                                                                    if (v9 >= (v9 % (((v9 / (((v9 / (((v9 - (^(v9)))) | 1))) | 1))) | 1))) {
                                                                        var v50 float64 = (float64(0.0) / (0.1 + float64(0.1)))
                                                                        _ = v50
                                                                    }
                                                                    v3 = v3
                                                                }
                                                            }
                                                            _ = fi49
                                                        }
                                                    case "\n\t":
                                                        v51 := GenS7{F0: (v3 - (v3 + (v3 - 726))), F1: 65535}
                                                        _ = v51
                                                    case "hello":
                                                        var v52 float32 = float32(3.14159)
                                                        _ = v52
                                                    default:
                                                        v3 += 1
                                                    }
                                                }
                                                _ = v45
                                            } else {
                                                var v53 int = 0
                                                delete(v2, "")
                                                _ = v53
                                            }
                                            switch (^(fi29)) {
                                            case 2147483647:
                                                if (982 != 0) {
                                                    for fi54 := 0; fi54 < 9; fi54++ {
                                                        if (127 > 116) {
                                                            v55 := GenS4{F0: -2147483648}
                                                            v56 := GenS8{F0: 1}
                                                            switch (fi29 << (((v14 & v14)) & 63)) {
                                                            case 4294967295:
                                                                v3 = v3
                                                            case 984:
                                                                v57 := GenS9{F0: fi29}
                                                                _ = v57
                                                            case 127:
                                                                fmt.Println(fi54)
                                                            default:
                                                                v58 := GenS3{F0: 1}
                                                                _ = v58
                                                            }
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            fmt.Println(fi54, fi29, fi29)
                                                        }
                                                        switch fi54 {
                                                        case -517:
                                                            v14 += v14
                                                        default:
                                                            var v59 int = fi54
                                                            _ = v59
                                                        }
                                                        _ = fi54
                                                    }
                                                    v60 := GenS7{F0: v3, F1: v14}
                                                    _ = v60
                                                } else {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if ((0.1 + 1e10) >= 3.14159) {
                                                            var v61 int16 = 32767
                                                            _ = v61
                                                        }
                                                    case "a":
                                                        fmt.Println(v14, v14)
                                                    default:
                                                        for fi62 := 0; fi62 < 14; fi62++ {
                                                            v3 = v3
                                                            v63 := GenS4{F0: -1}
                                                            _ = fi62
                                                            _ = v63
                                                        }
                                                    }
                                                    v64 := GenS2{F0: fi29, F1: 128, F2: 15}
                                                    _ = v64
                                                }
                                            case 4294967295:
                                                fmt.Println(v3)
                                            }
                                            v14 = 652
                                        }
                                        if (false || true) {
                                            for fi65 := 0; fi65 < 19; fi65++ {
                                                v4 = append(v4, fi29)
                                                v3 = v3
                                                if (true && (!(0.1 <= 0.0))) {
                                                    if false {
                                                        fmt.Println(v14, fi65)
                                                    } else {
                                                        switch ("hello" + ("foo bar" + "")) {
                                                        case "hello":
                                                            var v66 float32 = (0.1 / 1e10)
                                                            _ = v66
                                                        case "":
                                                            v67 := GenS8{F0: -1}
                                                            _ = v67
                                                        default:
                                                            var v68 uint8 = 42
                                                            _ = v68
                                                        }
                                                        v69 := GenS6{F0: v3}
                                                        _ = v69
                                                    }
                                                    switch ("foo bar" + (("" + "a") + "\n\t")) {
                                                    case "hello":
                                                        v14 -= v14
                                                    case "a":
                                                        v70 := GenS10{F0: v3, F1: (true && false), F2: 7}
                                                        _ = v70
                                                    case "":
                                                        var v71 int64 = -948
                                                        _ = v71
                                                    default:
                                                        var v72 bool = true
                                                        _ = v72
                                                    }
                                                    v73 := GenS8{F0: 32767}
                                                    var v74 uint64 = uint64(v9)
                                                    fmt.Println(fi29, v3, v14)
                                                    if (false || ((780 > v14) || ((v14 << ((v14) & 63)) <= (^(v14))))) {
                                                        fmt.Println(fi65)
                                                        fmt.Println(v9, v14)
                                                    }
                                                    v75 := GenS11{F0: (!false)}
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                }
                                                _ = fi65
                                            }
                                            v76 := GenS12{F0: 1, F1: 42, F2: 0}
                                            _ = v76
                                        } else {
                                            var v77 int8 = -1
                                            v78 := GenS13{F0: -1}
                                            _ = v77
                                            _ = v78
                                        }
                                    }
                                    if true {
                                        v9 -= 408
                                        v14 -= v14
                                    } else {
                                        if (1 > (v14 + v14)) {
                                            var v79 int64 = 9223372036854775807
                                            v80 := GenS11{F0: (true || true)}
                                            fmt.Println(v9, v14, v14)
                                            _ = v79
                                            _ = v80
                                        }
                                    }
                                }
                            }
                        }
                        var v81 uint32 = 42
                        _ = fi29
                        _ = v30
                        _ = v81
                    }
                } else {
                    for fi82 := 0; fi82 < 11; fi82++ {
                        var v83 uint64 = (v3 | v3)
                        var v84 float32 = 1e10
                        _ = fi82
                        _ = v83
                        _ = v84
                    }
                }
            }
        }
        for fi85 := 0; fi85 < 17; fi85++ {
            if (fi85 < -9223372036854775808) {
                if false {
                    if false {
                        v9 = v9
                    }
                    var v86 float64 = (-2.25 / 1e10)
                    for fi87 := 0; fi87 < 10; fi87++ {
                        var v88 uint8 = 128
                        _ = len(v4)
                        v89 := GenS8{F0: -385}
                        _ = fi87
                        _ = v88
                        _ = v89
                    }
                    _ = v86
                } else {
                    v14 -= (v14 << ((1) & 63))
                }
                if (2147483648 == 65535) {
                    _ = len(v4)
                    v4 = append(v4, fi85)
                } else {
                    v3 = (v3 >> ((v14) & 63))
                }
            }
            for fi90 := 0; fi90 < 9; fi90++ {
                var v91 uint8 = 42
                _ = fi90
                _ = v91
            }
            v3 = uint64(v9)
            _ = fi85
        }
        for fi92 := 0; fi92 < 20; fi92++ {
            v2["foo bar"] = (v14 != (v14 ^ 259))
            _ = fi92
        }
    }
    if false {
        if len(v5) > 0 {
            v5[0] = -557
        }
        v9 -= v9
    } else {
        v14 = (v14 % ((1) | 1))
    }
    v93 := GenS2{F0: -824, F1: uint8(v3), F2: 2147483647}
    var v94 int16 = 0
    var v95 uint16 = (v9 * (v9 | (v9 & v9)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
    _ = v93
    _ = v94
    _ = v95
}
