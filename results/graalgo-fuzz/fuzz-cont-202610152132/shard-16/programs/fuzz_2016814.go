package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint; F2 int32 }
type GenS4 struct { F0 int; F1 int; F2 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 string; F1 int8 }
type GenS11 struct { F0 uint32; F1 int16; F2 uint }
type GenS12 struct { F0 uint64; F1 string; F2 uint16 }
type GenS13 struct { F0 uint16; F1 int16; F2 int32 }

func main() {
    var v1 int = 32767
    v2 := make(map[string]uint32)
    var v3 uint = 18446744073709551615
    v4 := []string{"a"}
    fmt.Println(v3, v3)
    for fi5 := 0; fi5 < 12; fi5++ {
        if ((!(0 >= (^(v3)))) && (true || false)) {
            if (true || (-2147483648 == -1)) {
                _ = len(v2)
                var v6 uint = v3
                var v7 int16 = -32768
                _ = v6
                _ = v7
            } else {
                if (-207 < -1) {
                    for fi8 := 0; fi8 < 7; fi8++ {
                        var v9 uint32 = 308
                        _ = fi8
                        _ = v9
                    }
                    var v10 int = (-(fi5))
                    _ = v10
                } else {
                    v3 = v3
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v3, fi5, v3)
                    case "hello":
                        if (!(!(1 < 1))) {
                            if ((float64(3.14159) + (1e10 / float64((1.5 * -2.25)))) <= -2.25) {
                                if (!false) {
                                    var v11 int32 = 42
                                    v12 := GenS1{F0: -9223372036854775808}
                                    fmt.Println(v3, v1, v3)
                                    _ = v11
                                    _ = v12
                                } else {
                                    var v13 int32 = -1
                                    v14 := GenS1{F0: -1}
                                    for fi15 := 0; fi15 < 12; fi15++ {
                                        switch "" {
                                        case "\n\t":
                                            fmt.Println(v13, v1)
                                        }
                                        _ = fi15
                                    }
                                    _ = v13
                                    _ = v14
                                }
                                if (796 != 42) {
                                    var v16 bool = true
                                    _ = v16
                                }
                                if false {
                                    var v17 float64 = -2.25
                                    v1 = (v1 << ((v3) & 63))
                                    _ = v17
                                } else {
                                    var v18 float32 = 0.1
                                    _ = v18
                                }
                            }
                            for fi19 := 0; fi19 < 12; fi19++ {
                                if false {
                                    v20 := GenS1{F0: 1}
                                    _ = v20
                                }
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                                fmt.Println(v3)
                                if false {
                                    if false {
                                        v3 += v3
                                        v1 = (fi5 | (-(fi5)))
                                    }
                                    if true {
                                        v21 := GenS1{F0: 468}
                                        _ = v21
                                    }
                                    v3 = v3
                                    var v22 uint32 = 547
                                    _ = v22
                                }
                                _ = fi19
                            }
                            if true {
                                var v23 float32 = ((1.5 + (float32(0.0) / 3.14159)) / 3.14159)
                                v23 += -2.25
                                _ = v23
                            }
                        } else {
                            if (!((!(!true)) || true)) {
                                var v24 uint16 = 65535
                                if true {
                                    v25 := GenS1{F0: -78}
                                    if len(v4) > 0 {
                                        v4[0] = "a"
                                    }
                                    _ = v25
                                }
                                _ = v24
                            } else {
                                for fi26 := 0; fi26 < 20; fi26++ {
                                    v4 = append(v4, "")
                                    fmt.Println(fi26, v3, fi26)
                                    v4 = append(v4, "foo bar")
                                    _ = fi26
                                }
                            }
                        }
                    }
                    if true {
                        v1 -= (fi5 * (v1 >> ((v3) & 63)))
                        switch "" {
                        case "foo bar":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "\n\t":
                            if (424 > -2147483648) {
                                if false {
                                    if (!true) {
                                        if (-9223372036854775808 < 105) {
                                            for fi27 := 0; fi27 < 11; fi27++ {
                                                if len(v4) > 0 {
                                                    v4[0] = "\n\t"
                                                }
                                                fmt.Println(v3)
                                                _ = fi27
                                            }
                                            if false {
                                                if true {
                                                    if true {
                                                        v28 := GenS1{F0: 9223372036854775807}
                                                        v1 += -1
                                                        v29 := GenS2{F0: 0}
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        v30 := GenS3{F0: 42, F1: v3, F2: 42}
                                                        v31 := GenS1{F0: 42}
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                } else {
                                                    var v32 uint = (^(v3))
                                                    fmt.Println(v1)
                                                    v2["\n\t"] = 2
                                                    _ = v32
                                                }
                                            } else {
                                                v1 -= v1
                                                v33 := GenS1{F0: -1}
                                                var v34 uint32 = 2
                                                _ = v33
                                                _ = v34
                                            }
                                        } else {
                                            v1 -= v1
                                            var v35 uint8 = 0
                                            _ = v35
                                        }
                                        if true {
                                            if true {
                                                v1 += v1
                                            } else {
                                                if true {
                                                    var v36 bool = (true || (2147483647 != 2))
                                                    var v37 bool = (int16(v3) > 0)
                                                    _ = v36
                                                    _ = v37
                                                }
                                                switch (v1 + (^(v1))) {
                                                case -935:
                                                    for fi38 := 0; fi38 < 18; fi38++ {
                                                        v1 = 104
                                                        v39 := GenS2{F0: 0}
                                                        v40 := GenS4{F0: (fi38 / ((fi38) | 1)), F1: fi5, F2: (v1 >> ((v3) & 63))}
                                                        v4 = append(v4, ("hello" + "\n\t"))
                                                        if true {
                                                            fmt.Println(v3, v1)
                                                            switch fi38 {
                                                            case -2147483648:
                                                                v41 := GenS3{F0: 42, F1: (^(v3)), F2: 2147483647}
                                                                _ = v41
                                                            default:
                                                                switch v1 {
                                                                case 127:
                                                                    var v42 float64 = 0.1
                                                                    _ = v42
                                                                case 32768:
                                                                    v43 := GenS5{F0: (1.5 < (-2.25 + 0.0))}
                                                                    _ = v43
                                                                case -1:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = fi38
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                default:
                                                    v1 += (^(fi5))
                                                }
                                                for fi44 := 0; fi44 < 7; fi44++ {
                                                    v3 = 0
                                                    _ = fi44
                                                }
                                            }
                                            for fi45 := 0; fi45 < 15; fi45++ {
                                                var v46 int8 = 78
                                                _ = fi45
                                                _ = v46
                                            }
                                        } else {
                                            v1 += (v1 >> ((v3) & 63))
                                        }
                                        var v47 int = v1
                                        _ = v47
                                    } else {
                                        for fi48 := 0; fi48 < 4; fi48++ {
                                            fmt.Println(fi5, fi5, v1)
                                            var v49 string = ("" + ("a" + "a"))
                                            var v50 uint16 = 929
                                            v1 -= (fi5 + v1)
                                            fmt.Println(v49)
                                            _ = fi48
                                            _ = v49
                                            _ = v50
                                        }
                                        var v51 string = "foo bar"
                                        _ = len(v4)
                                        _ = v51
                                    }
                                    if (55 < -123) {
                                        v52 := GenS2{F0: 1}
                                        v3 = (v3 - (v3 * (v3 - 65535)))
                                        _ = v52
                                    } else {
                                        for fi53 := 0; fi53 < 8; fi53++ {
                                            v54 := GenS4{F0: -889, F1: (fi53 * v1), F2: v1}
                                            if true {
                                                var v55 uint8 = 2
                                                _ = v55
                                            }
                                            var v56 uint = (^(v3))
                                            _ = fi53
                                            _ = v54
                                            _ = v56
                                        }
                                        v3 = 1
                                    }
                                    fmt.Println(fi5, fi5)
                                }
                                fmt.Println(v1, v3, v3)
                                v1 = 9223372036854775807
                            }
                        case "":
                            switch (v1 * v1) {
                            case 32768:
                                switch fi5 {
                                case -1:
                                    var v57 int = 128
                                    _ = v57
                                default:
                                    fmt.Println(v3)
                                }
                            case -2147483648:
                                var v58 int8 = 127
                                _ = v58
                            default:
                                v59 := GenS6{F0: 65535, F1: 18446744073709551615}
                                _ = v59
                            }
                        default:
                            v2[("foo bar" + "foo bar")] = 65535
                        }
                    } else {
                        for fi60 := 0; fi60 < 9; fi60++ {
                            fmt.Println(fi60)
                            var v61 int32 = -1
                            v62 := GenS7{F0: (233 > -812), F1: true}
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                        v1 -= fi5
                        if true {
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                for fi63 := 0; fi63 < 13; fi63++ {
                                    v3 = (^(v3))
                                    fmt.Println(fi63, fi63)
                                    if len(v4) > 0 {
                                        v4[0] = "a"
                                    }
                                    _ = fi63
                                }
                            } else {
                                switch (v1 / ((fi5) | 1)) {
                                case 543:
                                    v3 = v3
                                case -2147483648:
                                    for fi64 := 0; fi64 < 1; fi64++ {
                                        v65 := GenS8{F0: 408}
                                        var v66 float64 = 0.0
                                        v66 = v66
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                case 128:
                                    switch v1 {
                                    case -370:
                                        v1 = 127
                                    case 42:
                                        var v67 float32 = 1.5
                                        _ = v67
                                    case 9223372036854775807:
                                        delete(v2, "foo bar")
                                    }
                                default:
                                    var v68 uint8 = 1
                                    _ = v68
                                }
                                for fi69 := 0; fi69 < 1; fi69++ {
                                    v1 = v1
                                    _ = fi69
                                }
                                if false {
                                    v70 := GenS9{F0: 1, F1: 873}
                                    _ = v70
                                } else {
                                    if (("a" + "a") < "") {
                                        v71 := GenS5{F0: true}
                                        v1 = (v1 / ((976) | 1))
                                        _ = v71
                                    }
                                    v3 += (v3 % ((4294967295) | 1))
                                    delete(v2, "a")
                                }
                            }
                            _ = len(v2)
                            var v72 uint32 = 2147483648
                            _ = v72
                        } else {
                            fmt.Println(fi5, fi5)
                        }
                    }
                }
            }
            var v73 int16 = 42
            delete(v2, ("\n\t" + "\n\t"))
            _ = v73
        }
        var v74 float32 = 3.14159
        v1 += 42
        _ = fi5
        _ = v74
    }
    fmt.Println(v3)
    fmt.Println(v1, v1)
    if true {
        switch "foo bar" {
        case "":
            v3 = (^(v3))
        case "\n\t":
            switch (^(v1)) {
            case 2147483647:
                for fi75 := 0; fi75 < 9; fi75++ {
                    var v76 int32 = 300
                    _ = fi75
                    _ = v76
                }
            default:
                if false {
                    v1 += (v1 * v1)
                    switch v1 {
                    case 9223372036854775807:
                        if false {
                            var v77 float64 = (0.0 / ((0.0 + (0.1 + ((-2.25 / 1.0) / (0.0 + ((1.5 + 1e10) - (0.1 / 1e10)))))) - 0.0))
                            _ = v77
                        }
                    }
                } else {
                    v78 := GenS10{F0: ((("foo bar" + "\n\t") + "hello") + ""), F1: 0}
                    _ = v78
                }
            }
        case "hello":
            v3 -= v3
        }
    }
    v1 = (^(v1))
    switch ("a" + (("foo bar" + "foo bar") + "hello")) {
    case "hello":
        switch ("" + "") {
        case "hello":
            var v79 uint64 = 4294967295
            _ = v79
        case "a":
            var v80 uint8 = 2
            _ = v80
        case "":
            v3 = (v3 % (((v3 + (^(v3)))) | 1))
        }
    case "":
        _ = len(v4)
    default:
        switch (-(v1)) {
        case 672:
            v1 = 344
        case -1:
            if (128 < 127) {
                v3 = v3
                v81 := v2["\n\t"]
                _ = v81
            }
        case 274:
            v82 := GenS11{F0: 498, F1: 1000, F2: 18446744073709551615}
            _ = v82
        default:
            if false {
                v83 := GenS12{F0: 2, F1: "", F2: 255}
                var v84 int32 = -424
                _ = v83
                _ = v84
            }
        }
    }
    for fi85 := 0; fi85 < 17; fi85++ {
        if (0 != (^(v3))) {
            v3 = v3
            var v86 uint64 = 2
            _ = v86
        } else {
            var v87 string = "foo bar"
            var v88 uint64 = 230
            v89 := GenS7{F0: true, F1: ((-601 < 42) && true)}
            _ = v87
            _ = v88
            _ = v89
        }
        v1 = v1
        _ = fi85
    }
    for fi90 := 0; fi90 < 14; fi90++ {
        switch ((("foo bar" + ("hello" + "a")) + ("\n\t" + "foo bar")) + "foo bar") {
        case "a":
            fmt.Println(fi90, v1)
        case "\n\t":
            var v91 int16 = -917
            _ = v91
        }
        v2[""] = 4294967295
        v92 := GenS5{F0: (!(0 <= v1))}
        _ = fi90
        _ = v92
    }
    v2[("a" + "")] = 2147483648
    v1 = (v1 | -9223372036854775808)
    v3 += 18446744073709551615
    var v93 int64 = -388
    for fi94 := 0; fi94 < 19; fi94++ {
        var v95 int16 = 32767
        switch "\n\t" {
        case "foo bar":
            for fi96 := 0; fi96 < 12; fi96++ {
                var v97 int = (-(fi96))
                v98 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: 207}
                _ = fi96
                _ = v97
                _ = v98
            }
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        case "hello":
            if true {
                v1 -= v1
            }
        }
        v93 += v93
        _ = fi94
        _ = v95
    }
    v99 := GenS9{F0: 1, F1: 8}
    for fi100 := 0; fi100 < 13; fi100++ {
        switch v1 {
        case 0:
            var v101 uint32 = 0
            _ = v101
        case -1:
            switch (^(fi100)) {
            case -751:
                var v102 uint32 = 686
                _ = v102
            default:
                var v103 float64 = 1.5
                _ = v103
            }
        case -748:
            v1 += -321
        default:
            v104 := GenS1{F0: (v93 | (v93 % (((v93 >> ((0) & 63))) | 1)))}
            _ = v104
        }
        var v105 bool = false
        _ = fi100
        _ = v105
    }
    v106 := GenS13{F0: 255, F1: -482, F2: -333}
    for fi107 := 0; fi107 < 16; fi107++ {
        _ = len(v4)
        v93 = 0
        fmt.Println(fi107, fi107, v93)
        _ = fi107
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v93
    _ = v99
}
