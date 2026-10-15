package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 int; F2 int8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 string; F2 int32 }
type GenS10 struct { F0 uint; F1 uint }
type GenS11 struct { F0 int64; F1 uint; F2 uint64 }
type GenS12 struct { F0 int8; F1 uint; F2 string }
type GenS13 struct { F0 uint8 }

func main() {
    var v1 string = (("foo bar" + "") + ("hello" + "hello"))
    v2 := []string{"hello", (v1 + "foo bar"), v1, "\n\t"}
    var v3 string = v1
    v4 := []int{0, -9223372036854775808}
    var v5 float64 = (0.1 - 0.0)
    v6 := make(map[string]int)
    if (!false) {
        v7 := GenS1{F0: 0, F1: 1}
        _ = v7
    } else {
        switch "" {
        case "\n\t":
            v1 = v1
        case "foo bar":
            v3 = v1
        case "a":
            _ = len(v4)
        }
        if ((!false) && (-1 == 274)) {
            if false {
                var v8 int64 = 0
                for fi9 := 0; fi9 < 8; fi9++ {
                    switch (fi9 | (fi9 % (((fi9 ^ (^(fi9)))) | 1))) {
                    case 32768:
                        v10 := GenS2{F0: 127}
                        _ = v10
                    default:
                        for fi11 := 0; fi11 < 8; fi11++ {
                            var v12 int32 = 585
                            v8 = (v8 % (((v8 ^ (v8 >> ((88) & 63)))) | 1))
                            var v13 uint64 = 42
                            v8 += v8
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                    fmt.Println(v3)
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v8 = v8
                        var v15 int64 = -1000
                        v15 = (-(v8))
                        _ = fi14
                        _ = v15
                    }
                    _ = fi9
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    _ = len(v2)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch (v1 + "") {
                    case "\n\t":
                        _ = len(v6)
                    case "hello":
                        if ((((0.0 - (1.5 / 3.14159)) + (v5 / float64(1.5))) >= 0.0) && (589 >= 65535)) {
                            switch "foo bar" {
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = v1
                                }
                            case "hello":
                                v17 := GenS3{F0: -967, F1: 593, F2: 643}
                                _ = v17
                            default:
                                fmt.Println(v1, v5, v5)
                            }
                        } else {
                            if true {
                                var v18 uint8 = 139
                                _ = v18
                            }
                        }
                    case "":
                        var v19 uint32 = 4294967295
                        _ = v19
                    }
                    _ = fi16
                }
                _ = v8
            } else {
                v20 := GenS4{F0: 42}
                var v21 float32 = (float32(float64(0.1)) - 3.14159)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v20
                _ = v21
            }
            switch 0 {
            case 127:
                if true {
                    switch 4294967295 {
                    case 1:
                        for fi22 := 0; fi22 < 8; fi22++ {
                            var v23 uint8 = 1
                            for fi24 := 0; fi24 < 10; fi24++ {
                                var v25 int8 = 42
                                if true {
                                    var v26 uint16 = 790
                                    if ((-1 >= 608) && (!(v23 > 255))) {
                                        var v27 uint64 = 191
                                        _ = v27
                                    }
                                    _ = v26
                                }
                                v6[v3] = 4294967295
                                _ = fi24
                                _ = v25
                            }
                            _ = fi22
                            _ = v23
                        }
                    }
                    _ = len(v2)
                } else {
                    if true {
                        for fi28 := 0; fi28 < 16; fi28++ {
                            if (float32(0.0) != 0.0) {
                                v29 := GenS4{F0: 155}
                                if false {
                                    switch v3 {
                                    case "":
                                        if (v5 <= 0.0) {
                                            if (!false) {
                                                v1 = v1
                                                v30 := GenS3{F0: 32767, F1: 0, F2: 2}
                                                _ = v30
                                            }
                                        } else {
                                            var v31 string = v1
                                            _ = v31
                                        }
                                    }
                                    if (!(false || false)) {
                                        v1 = v3
                                        for fi32 := 0; fi32 < 11; fi32++ {
                                            switch -333 {
                                            case 0:
                                                v6[(v3 + (v3 + v1))] = fi28
                                            case -1:
                                                if false {
                                                    if (true && (!((fi32 | (fi32 << ((229) & 63))) >= (fi32 | fi28)))) {
                                                        v33 := GenS3{F0: 32767, F1: 2, F2: 962}
                                                        v34 := GenS1{F0: 18446744073709551615, F1: -2147483648}
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        if true {
                                                            fmt.Println(fi32, fi32, v3)
                                                            if len(v4) > 0 {
                                                                v4[0] = (fi32 / ((fi32) | 1))
                                                            }
                                                            if (0.0 < v5) {
                                                                var v35 uint = 18446744073709551615
                                                                _ = v35
                                                            } else {
                                                                var v36 bool = false
                                                                switch ((v3 + v1) + v3) {
                                                                case "foo bar":
                                                                    v37 := GenS5{F0: 65535, F1: 0}
                                                                    _ = v37
                                                                case "a":
                                                                    if v36 {
                                                                        v5 -= (0.0 / (float64(0.1) * ((v5 - v5) * (-2.25 - v5))))
                                                                        _ = len(v6)
                                                                        switch fi32 {
                                                                        case 127:
                                                                            var v38 int64 = -94
                                                                            _ = v38
                                                                        }
                                                                        delete(v6, "hello")
                                                                    } else {
                                                                        delete(v6, "\n\t")
                                                                        var v39 int = (fi32 + (fi32 << ((0) & 63)))
                                                                        var v40 uint32 = 563
                                                                        _ = v39
                                                                        _ = v40
                                                                    }
                                                                case "\n\t":
                                                                    switch (fi32 & fi32) {
                                                                    case -880:
                                                                        if false {
                                                                            switch fi28 {
                                                                            case 4294967295:
                                                                                v5 = 1e10
                                                                            case 845:
                                                                                v41 := GenS6{F0: 0}
                                                                                _ = v41
                                                                            case 0:
                                                                                v3 = "\n\t"
                                                                            }
                                                                            v42 := GenS6{F0: 0}
                                                                            v43 := GenS7{F0: -1, F1: (fi28 - fi28), F2: 0}
                                                                            _ = v42
                                                                            _ = v43
                                                                        }
                                                                    case -764:
                                                                        v44 := GenS3{F0: 42, F1: 413, F2: 4294967295}
                                                                        _ = v44
                                                                    case 42:
                                                                        if (!(1.5 == (1e10 / -2.25))) {
                                                                            var v45 uint8 = 127
                                                                            fmt.Println(v1, fi28, v5)
                                                                            _ = v45
                                                                        } else {
                                                                            if v36 {
                                                                                v5 = (1.5 * v5)
                                                                                var v46 int8 = 42
                                                                                _ = v46
                                                                            } else {
                                                                                var v47 uint32 = 468
                                                                                var v48 float32 = 1.5
                                                                                switch (^(fi28)) {
                                                                                case -9223372036854775808:
                                                                                    v2 = append(v2, (v3 + ("a" + "a")))
                                                                                case -2147483648:
                                                                                    _ = len(v4)
                                                                                }
                                                                                _ = v47
                                                                                _ = v48
                                                                            }
                                                                        }
                                                                    default:
                                                                        v49 := GenS8{F0: 32767}
                                                                        _ = v49
                                                                    }
                                                                default:
                                                                    v50 := GenS5{F0: 255, F1: 100}
                                                                    _ = v50
                                                                }
                                                                var v51 bool = (false && (4294967295 > 1))
                                                                fmt.Println(fi32)
                                                                _ = v36
                                                                _ = v51
                                                            }
                                                        }
                                                        if (4294967295 == 0) {
                                                            v5 -= (v5 / float64((float32(v5) - 1e10)))
                                                        } else {
                                                            v1 = "a"
                                                            var v52 int16 = -1
                                                            _ = v52
                                                        }
                                                    }
                                                    fmt.Println(v1, v5, v1)
                                                }
                                                v53 := GenS9{F0: 2, F1: "\n\t", F2: -1}
                                                v4 = append(v4, (fi28 ^ fi28))
                                                _ = v53
                                            default:
                                                v54 := GenS10{F0: 4294967295, F1: 42}
                                                _ = v54
                                            }
                                            v55 := GenS5{F0: 64, F1: 26}
                                            _ = fi32
                                            _ = v55
                                        }
                                        fmt.Println(fi28, fi28)
                                    } else {
                                        for fi56 := 0; fi56 < 7; fi56++ {
                                            if false {
                                                fmt.Println(v1, fi28, v1)
                                                _ = len(v6)
                                            }
                                            fmt.Println(v3, fi56)
                                            delete(v6, (v3 + (v1 + v3)))
                                            v1 = (v1 + v3)
                                            _ = fi56
                                        }
                                        for fi57 := 0; fi57 < 6; fi57++ {
                                            var v58 int32 = -410
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = fi57
                                            _ = v58
                                        }
                                        if ((v5 - v5) != -2.25) {
                                            var v59 int16 = -1
                                            for fi60 := 0; fi60 < 15; fi60++ {
                                                v1 = v1
                                                _ = fi60
                                            }
                                            var v61 uint16 = 0
                                            _ = v59
                                            _ = v61
                                        }
                                    }
                                } else {
                                    if ((255 != 42) || (128 >= 35)) {
                                        var v62 int32 = 544
                                        var v63 uint16 = 42
                                        _ = v62
                                        _ = v63
                                    } else {
                                        if true {
                                            v64 := GenS4{F0: 0}
                                            var v65 uint32 = 65535
                                            _ = v64
                                            _ = v65
                                        } else {
                                            v66 := GenS11{F0: 42, F1: 259, F2: 633}
                                            _ = v66
                                        }
                                        if false {
                                            fmt.Println(fi28, v3)
                                        }
                                    }
                                    if (42 <= 1) {
                                        v3 = v3
                                    }
                                    if (373 >= int16(fi28)) {
                                        v67 := GenS6{F0: 246}
                                        _ = v67
                                    }
                                }
                                _ = v29
                            } else {
                                var v68 string = (v1 + v1)
                                _ = v68
                            }
                            v1 = ("" + (v1 + (v1 + v1)))
                            _ = fi28
                        }
                        var v69 float32 = -2.25
                        _ = v69
                    } else {
                        v70 := GenS3{F0: 42, F1: 409, F2: 1}
                        fmt.Println(v5)
                        fmt.Println(v5, v3)
                        _ = v70
                    }
                }
            case -655:
                for fi71 := 0; fi71 < 9; fi71++ {
                    var v72 int16 = 1000
                    _ = fi71
                    _ = v72
                }
            }
            var v73 uint = 396
            _ = v73
        } else {
            v74 := GenS7{F0: 42, F1: 32767, F2: 127}
            delete(v6, (v3 + v1))
            var v75 int16 = 0
            _ = v74
            _ = v75
        }
    }
    var v76 int32 = -2147483648
    var v77 float32 = 1e10
    switch -2147483648 {
    case -1:
        if (57 < 1) {
            var v78 uint = 0
            v79 := GenS9{F0: 234, F1: v3, F2: v76}
            _ = v78
            _ = v79
        } else {
            var v80 int = 42
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v80
        }
    case 1:
        _ = len(v6)
    }
    for fi81 := 0; fi81 < 13; fi81++ {
        fmt.Println(v77, v5)
        var v82 int16 = 32767
        var v83 int64 = 9223372036854775807
        _ = fi81
        _ = v82
        _ = v83
    }
    for fi84 := 0; fi84 < 10; fi84++ {
        v85 := GenS2{F0: 255}
        v5 += (((float64(((((((float32(v5) * -2.25) + v77) * v77) + v77) * 0.1) * v77)) + float64(v77)) * (v5 * 1.5)) * (1.5 - v5))
        v86 := GenS12{F0: 127, F1: 752, F2: v3}
        _ = fi84
        _ = v85
        _ = v86
    }
    switch v3 {
    case "hello":
        v76 -= v76
    }
    v77 = v77
    fmt.Println(v5, v3)
    for fi87 := 0; fi87 < 1; fi87++ {
        fmt.Println(v77, v77)
        var v88 uint64 = 742
        if true {
            if (1 >= -674) {
                if false {
                    v5 -= v5
                    switch (fi87 % ((fi87) | 1)) {
                    case 32768:
                        v89 := v6[v3]
                        _ = v89
                    case 4294967295:
                        v88 += v88
                    }
                } else {
                    v5 = v5
                    for fi90 := 0; fi90 < 3; fi90++ {
                        v91 := GenS2{F0: 2}
                        _ = fi90
                        _ = v91
                    }
                }
                v92 := GenS6{F0: 26}
                switch fi87 {
                case 2147483647:
                    v88 = v88
                case 9223372036854775807:
                    for fi93 := 0; fi93 < 4; fi93++ {
                        fmt.Println(v1, v77, v88)
                        var v94 int32 = -2147483648
                        if true {
                            v5 = v5
                        } else {
                            if (!true) {
                                v4 = append(v4, 32767)
                                var v95 int16 = 32767
                                v76 = v76
                                _ = v95
                            } else {
                                v96 := GenS1{F0: 428, F1: v76}
                                _ = v96
                            }
                            if false {
                                var v97 int = 128
                                var v98 int16 = 85
                                _ = v97
                                _ = v98
                            }
                            v88 += (v88 | (v88 | (^(v88))))
                        }
                        var v99 uint8 = 29
                        _ = fi93
                        _ = v94
                        _ = v99
                    }
                case 4294967295:
                    v88 += v88
                default:
                    fmt.Println(v5, v77)
                }
                _ = v92
            }
        } else {
            for fi100 := 0; fi100 < 8; fi100++ {
                v6["a"] = (fi100 - fi87)
                _ = fi100
            }
        }
        _ = fi87
        _ = v88
    }
    if false {
        var v101 uint8 = 128
        fmt.Println(v76, v3, v5)
        _ = v101
    }
    if len(v2) > 0 {
        v2[0] = (v1 + v1)
    }
    fmt.Println(v1, v1, v5)
    if false {
        var v102 float64 = v5
        for fi103 := 0; fi103 < 4; fi103++ {
            var v104 float64 = float64(v77)
            _ = fi103
            _ = v104
        }
        _ = v102
    } else {
        v105 := GenS13{F0: 0}
        _ = len(v6)
        _ = v105
    }
    fmt.Println(v76)
    v77 += (v77 + (v77 / (v77 * v77)))
    fmt.Println(v3, v1, v76)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
}
