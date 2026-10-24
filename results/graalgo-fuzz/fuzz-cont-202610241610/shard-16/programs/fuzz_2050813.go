package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 int64; F2 uint16 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int8; F1 uint8; F2 string }

func main() {
    var v1 int64 = 0
    var v2 uint8 = 1
    v3 := make(map[int]string)
    var v4 string = "hello"
    var v5 float64 = 3.14159
    var v6 int64 = (v1 | -1)
    switch v4 {
    case "":
        v1 += v1
    case "foo bar":
        fmt.Println(v2, v6)
    default:
        v5 += v5
    }
    v7 := GenS1{F0: 42}
    switch ((v4 + ("hello" + (v4 + v4))) + (v4 + v4)) {
    case "hello":
        var v8 uint64 = 1
        _ = v8
    default:
        if true {
            v9 := []int{2147483647, -2147483648, 32767}
            v6 += (v6 & 1)
            _ = v9
        }
    }
    v5 = ((0.0 - ((v5 * 1e10) * 0.0)) - 0.1)
    v10 := []int{42, -1, 165}
    switch 1 {
    case -444:
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 string = "foo bar"
            if (!(!((true && (!true)) || true))) {
                var v13 string = v12
                v14 := GenS2{F0: 4294967295, F1: -1, F2: int32(fi11)}
                switch "foo bar" {
                case "\n\t":
                    var v15 int16 = 32767
                    _ = v15
                case "":
                    for fi16 := 0; fi16 < 19; fi16++ {
                        switch v4 {
                        case "hello":
                            var v17 uint = 894
                            _ = v17
                        }
                        v1 = -40
                        var v18 int32 = 2147483647
                        _ = fi16
                        _ = v18
                    }
                case "foo bar":
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 uint64 = 0
                        var v21 string = ("hello" + "")
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                default:
                    for fi22 := 0; fi22 < 11; fi22++ {
                        v23 := GenS1{F0: 330}
                        _ = fi22
                        _ = v23
                    }
                }
                _ = v13
                _ = v14
            } else {
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 float64 = 1.5
                    v12 = v4
                    v26 := GenS1{F0: 584}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            if len(v10) > 0 {
                v10[0] = (fi11 / ((fi11) | 1))
            }
            _ = fi11
            _ = v12
        }
    case 1:
        if (true && (float32((((((v5 - v5) + 0.0) - 3.14159) * v5) * v5)) > 0.1)) {
            switch -572 {
            case 9223372036854775807:
                switch 354 {
                case -835:
                    var v27 int = -2147483648
                    _ = v27
                case 509:
                    if ((v5 / 0.1) > v5) {
                        for fi28 := 0; fi28 < 16; fi28++ {
                            var v29 uint64 = 0
                            var v30 int8 = 42
                            if (((v29 | v29) > v29) && ((217 != 1) || (v30 == v30))) {
                                v31 := GenS2{F0: v29, F1: 433, F2: -2147483648}
                                _ = v31
                            } else {
                                v32 := GenS1{F0: 32767}
                                for fi33 := 0; fi33 < 14; fi33++ {
                                    var v34 float32 = -2.25
                                    _ = fi33
                                    _ = v34
                                }
                                v35 := GenS3{F0: 806}
                                _ = v32
                                _ = v35
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    } else {
                        v1 = (v1 | v6)
                        _ = len(v10)
                        for fi36 := 0; fi36 < 6; fi36++ {
                            var v37 uint = uint(v2)
                            switch (("" + (v4 + (v4 + v4))) + "") {
                            case "foo bar":
                                v37 = 1
                            }
                            for fi38 := 0; fi38 < 4; fi38++ {
                                var v39 float32 = ((-2.25 / 1e10) - 1.5)
                                _ = fi38
                                _ = v39
                            }
                            _ = fi36
                            _ = v37
                        }
                    }
                }
            default:
                if (!false) {
                    if true {
                        if true {
                            fmt.Println(v5)
                            fmt.Println(v6, v5)
                            v4 = "\n\t"
                        } else {
                            for fi40 := 0; fi40 < 18; fi40++ {
                                var v41 uint8 = 255
                                switch (fi40 ^ 42) {
                                case 735:
                                    v42 := GenS4{F0: 1}
                                    _ = v42
                                case 9223372036854775807:
                                    v43 := GenS5{F0: -599, F1: (v6 << ((65535) & 63)), F2: 0}
                                    _ = v43
                                }
                                v44 := GenS6{F0: "", F1: 0}
                                _ = fi40
                                _ = v41
                                _ = v44
                            }
                            if false {
                                if (-2147483648 < 2147483647) {
                                    if false {
                                        for fi45 := 0; fi45 < 10; fi45++ {
                                            if false {
                                                var v46 int64 = v6
                                                _ = v46
                                            } else {
                                                for fi47 := 0; fi47 < 17; fi47++ {
                                                    var v48 uint8 = v2
                                                    if true {
                                                        if true {
                                                            fmt.Println(fi45)
                                                            v49 := GenS5{F0: -2147483648, F1: (v1 / ((v1) | 1)), F2: 65535}
                                                            _ = v49
                                                        }
                                                        if false {
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                            fmt.Println(v1, v6)
                                                            switch fi45 {
                                                            case -2147483648:
                                                                v50 := v3[(fi45 >> ((2) & 63))]
                                                                _ = v50
                                                            case 32767:
                                                                var v51 int64 = 9223372036854775807
                                                                _ = v51
                                                            }
                                                            fmt.Println(v48, v1, v5)
                                                            v2 += v2
                                                            v52 := GenS5{F0: -879, F1: v6, F2: 255}
                                                            _ = v52
                                                        } else {
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                            var v53 uint32 = 759
                                                            v54 := GenS3{F0: v6}
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        v6 = v6
                                                    }
                                                    fmt.Println(fi45, v5)
                                                    _ = fi47
                                                    _ = v48
                                                }
                                            }
                                            for fi55 := 0; fi55 < 12; fi55++ {
                                                if ((1.5 + 1e10) <= 1.5) {
                                                    var v56 uint32 = 2147483648
                                                    _ = v56
                                                }
                                                _ = fi55
                                            }
                                            _ = fi45
                                        }
                                    }
                                }
                            } else {
                                if (((true && true) || true) && false) {
                                    v4 = "foo bar"
                                    v4 = ("" + v4)
                                    switch "\n\t" {
                                    case "\n\t":
                                        _ = len(v3)
                                    case "hello":
                                        var v57 string = ("" + v4)
                                        _ = v57
                                    case "foo bar":
                                        v6 -= (v1 << ((346) & 63))
                                    }
                                } else {
                                    v2 = (v2 + 67)
                                    v58 := GenS6{F0: v4, F1: 2}
                                    _ = v58
                                }
                                v59 := GenS5{F0: 1, F1: v6, F2: 2}
                                _ = v59
                            }
                        }
                        if (286 == 65535) {
                            _ = len(v3)
                            var v60 uint8 = (^(v2))
                            var v61 int32 = 2147483647
                            _ = v60
                            _ = v61
                        } else {
                            if ("hello" != v4) {
                                fmt.Println(v4, v2, v4)
                                _ = len(v3)
                                v6 = v1
                            } else {
                                for fi62 := 0; fi62 < 11; fi62++ {
                                    fmt.Println(v6)
                                    if len(v10) > 0 {
                                        v10[0] = (fi62 - fi62)
                                    }
                                    v63 := GenS1{F0: 32767}
                                    _ = fi62
                                    _ = v63
                                }
                                var v64 uint32 = 42
                                v65 := GenS7{F0: 32767}
                                _ = v64
                                _ = v65
                            }
                        }
                    } else {
                        for fi66 := 0; fi66 < 18; fi66++ {
                            switch -925 {
                            case -2147483648:
                                if (((1e10 - 1e10) * (0.1 - float32(0.0))) > (3.14159 * 1e10)) {
                                    if (108 == 127) {
                                        switch (^(fi66)) {
                                        case -893:
                                            switch -2147483648 {
                                            case -702:
                                                if true {
                                                    var v67 int8 = int8(fi66)
                                                    _ = v67
                                                } else {
                                                    var v68 int16 = 1
                                                    _ = v68
                                                }
                                            case 1:
                                                for fi69 := 0; fi69 < 12; fi69++ {
                                                    switch (v4 + v4) {
                                                    case "\n\t":
                                                        v70 := GenS7{F0: 886}
                                                        _ = v70
                                                    default:
                                                        if false {
                                                            v4 = v4
                                                            v1 = 172
                                                            v10 = append(v10, 127)
                                                        } else {
                                                            var v71 int8 = 0
                                                            v10 = append(v10, (fi66 / ((-233) | 1)))
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = len(v10)
                                                    _ = fi69
                                                }
                                            case 127:
                                                var v72 float32 = 0.0
                                                _ = v72
                                            default:
                                                var v73 float32 = 3.14159
                                                _ = v73
                                            }
                                        case 533:
                                            _ = len(v10)
                                        }
                                        switch ((v4 + v4) + v4) {
                                        case "a":
                                            if true {
                                                fmt.Println(fi66, v1, v5)
                                                var v74 int32 = 2147483647
                                                _ = v74
                                            } else {
                                                delete(v3, fi66)
                                                for fi75 := 0; fi75 < 11; fi75++ {
                                                    var v76 int8 = 1
                                                    v77 := GenS8{F0: 65535}
                                                    var v78 int32 = 2147483647
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                    if false {
                                                        switch (("foo bar" + v4) + "a") {
                                                        case "\n\t":
                                                            if ((v2 - 128) != v2) {
                                                                v6 = v6
                                                            } else {
                                                                fmt.Println(fi75, v76, v78)
                                                                v10 = append(v10, (fi66 * fi75))
                                                                v79 := GenS9{F0: (v76 / ((100) | 1)), F1: (v2 + v2), F2: "foo bar"}
                                                                _ = v79
                                                            }
                                                        case "foo bar":
                                                            var v80 uint64 = 4294967295
                                                            _ = v80
                                                        case "a":
                                                            var v81 int8 = v76
                                                            _ = v81
                                                        }
                                                        v82 := GenS9{F0: v76, F1: v2, F2: "foo bar"}
                                                        _ = v82
                                                    } else {
                                                        var v83 float32 = 0.0
                                                        _ = v83
                                                    }
                                                    _ = fi75
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                }
                                            }
                                        case "":
                                            var v84 uint = uint(v2)
                                            _ = v84
                                        case "hello":
                                            for fi85 := 0; fi85 < 1; fi85++ {
                                                v5 = v5
                                                var v86 uint = 1
                                                _ = fi85
                                                _ = v86
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi66)
                                    }
                                    v5 += v5
                                    v3[42] = "a"
                                }
                            case 616:
                                v10 = append(v10, fi66)
                            default:
                                fmt.Println(v5)
                            }
                            v1 += v6
                            _ = fi66
                        }
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    }
                } else {
                    var v87 int32 = -1
                    var v88 uint = 248
                    _ = v87
                    _ = v88
                }
            }
            switch 0 {
            case 127:
                if false {
                    v1 = v1
                    _ = len(v3)
                } else {
                    fmt.Println(v4, v4, v4)
                }
            case 1:
                v2 = (^(v2))
            default:
                delete(v3, 2147483647)
            }
            var v89 uint8 = (v2 % (((v2 | (v2 << ((59) & 63)))) | 1))
            _ = v89
        } else {
            if false {
                var v90 int16 = 223
                fmt.Println(v2, v5, v4)
                _ = v90
            }
        }
    case 128:
        if len(v10) > 0 {
            v10[0] = -867
        }
    }
    switch -2147483648 {
    case -117:
        var v91 int = 128
        _ = v91
    case 982:
        v6 += -1
    }
    var v92 uint16 = 32767
    var v93 uint16 = v92
    if (2147483647 != uint32(v93)) {
        for fi94 := 0; fi94 < 3; fi94++ {
            if true {
                var v95 int = (^(fi94))
                for fi96 := 0; fi96 < 12; fi96++ {
                    switch "hello" {
                    case "":
                        fmt.Println(v93, v5)
                    case "a":
                        var v97 float64 = v5
                        _ = v97
                    default:
                        _ = len(v10)
                    }
                    v98 := GenS3{F0: v6}
                    var v99 int64 = v1
                    _ = fi96
                    _ = v98
                    _ = v99
                }
                _ = v95
            } else {
                v1 = v1
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v2 += v2
            }
            var v100 int32 = int32(v93)
            _ = fi94
            _ = v100
        }
        v4 = "foo bar"
        var v101 uint = uint(v2)
        _ = v101
    }
    v6 += -9223372036854775808
    v6 += v1
    _ = len(v3)
    _ = len(v3)
    v1 += 513
    if false {
        v2 -= 113
    } else {
        if true {
            var v102 int32 = -1
            for fi103 := 0; fi103 < 13; fi103++ {
                var v104 int8 = -128
                v6 -= (v6 / ((v1) | 1))
                _ = fi103
                _ = v104
            }
            if len(v10) > 0 {
                v10[0] = -1
            }
            _ = v102
        } else {
            var v105 int = 42
            var v106 float64 = v5
            for fi107 := 0; fi107 < 11; fi107++ {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi107
            }
            _ = v105
            _ = v106
        }
        v2 = 5
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v92
    _ = v93
}
