package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 uint8; F1 uint; F2 string }
type GenS8 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS9 struct { F0 uint; F1 uint; F2 uint64 }
type GenS10 struct { F0 bool; F1 uint }
type GenS11 struct { F0 int64; F1 bool }
type GenS12 struct { F0 uint8; F1 uint }
type GenS13 struct { F0 uint16 }

func main() {
    var v1 int32 = -1
    var v2 int = -485
    var v3 int64 = -516
    v4 := make(map[int]int)
    v5 := []int{}
    fmt.Println(v3)
    if false {
        for fi6 := 0; fi6 < 13; fi6++ {
            for fi7 := 0; fi7 < 6; fi7++ {
                var v8 int = 128
                if true {
                    var v9 uint8 = 255
                    switch "\n\t" {
                    case "a":
                        v10 := GenS1{F0: 74, F1: "hello"}
                        _ = v10
                    case "foo bar":
                        var v11 int32 = 567
                        _ = v11
                    case "hello":
                        v12 := GenS2{F0: ("a" + ("a" + "foo bar")), F1: 1}
                        _ = v12
                    default:
                        if (3.14159 == float64(0.0)) {
                            var v13 uint16 = 979
                            _ = v13
                        }
                    }
                    _ = v9
                } else {
                    if false {
                        var v14 bool = (true && (!false))
                        var v15 bool = v14
                        _ = v14
                        _ = v15
                    }
                }
                _ = fi7
                _ = v8
            }
            var v16 uint = 615
            _ = fi6
            _ = v16
        }
        switch v2 {
        case -9223372036854775808:
            switch 4294967295 {
            case 9223372036854775807:
                var v17 string = ("foo bar" + ("\n\t" + "foo bar"))
                _ = v17
            case 4294967295:
                for fi18 := 0; fi18 < 5; fi18++ {
                    v19 := GenS3{F0: "\n\t"}
                    var v20 string = "\n\t"
                    if (!(42 > 593)) {
                        var v21 float64 = (((3.14159 * float64(1.5)) + 1e10) - 1e10)
                        v1 -= (v1 << ((404) & 63))
                        var v22 uint64 = 0
                        _ = v21
                        _ = v22
                    } else {
                        switch fi18 {
                        case 128:
                            if (-1 < (fi18 & 9223372036854775807)) {
                                v20 = v20
                            }
                        case 616:
                            v4[fi18] = v2
                        case -817:
                            var v23 int32 = int32(v2)
                            _ = v23
                        }
                        v24 := GenS4{F0: 36, F1: ((((v20 + v20) + v20) + (v20 + v20)) + v20)}
                        _ = v24
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case 0:
                var v25 uint8 = 224
                _ = v25
            }
        case 128:
            if true {
                if true {
                    if true {
                        for fi26 := 0; fi26 < 10; fi26++ {
                            for fi27 := 0; fi27 < 6; fi27++ {
                                if true {
                                    v1 -= -1
                                    fmt.Println(fi27, v1, fi26)
                                    v28 := GenS1{F0: 42, F1: "a"}
                                    _ = v28
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = (v2 >> ((255) & 63))
                                    }
                                }
                                if true {
                                    var v29 int8 = -53
                                    _ = v29
                                } else {
                                    if (65535 != 221) {
                                        if len(v5) > 0 {
                                            v5[0] = (-(fi26))
                                        }
                                        fmt.Println(v2)
                                        var v30 uint64 = 2
                                        var v31 int64 = v3
                                        _ = v30
                                        _ = v31
                                    }
                                }
                                _ = fi27
                            }
                            _ = fi26
                        }
                        switch v2 {
                        case -1:
                            v32 := GenS3{F0: ("" + "")}
                            _ = v32
                        case 32768:
                            var v33 int = v2
                            _ = v33
                        case 157:
                            if (2147483648 > 101) {
                                var v34 uint = 75
                                _ = v34
                            } else {
                                var v35 string = ("a" + (("foo bar" + "hello") + "\n\t"))
                                v36 := GenS5{F0: -1}
                                fmt.Println(v35)
                                _ = v35
                                _ = v36
                            }
                        default:
                            v37 := GenS2{F0: "\n\t", F1: (v1 >> ((4294967295) & 63))}
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 10; fi38++ {
                            fmt.Println(v2, fi38)
                            if len(v5) > 0 {
                                v5[0] = v2
                            }
                            _ = fi38
                        }
                    }
                    for fi39 := 0; fi39 < 15; fi39++ {
                        v3 = v3
                        for fi40 := 0; fi40 < 17; fi40++ {
                            if (true && true) {
                                v1 += v1
                                var v41 string = "foo bar"
                                if true {
                                    var v42 uint64 = 4294967295
                                    _ = v42
                                }
                                _ = v41
                            }
                            var v43 string = "foo bar"
                            _ = len(v5)
                            _ = fi40
                            _ = v43
                        }
                        var v44 bool = (0 != 0)
                        _ = fi39
                        _ = v44
                    }
                    v1 += v1
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        default:
            var v45 uint16 = 42
            _ = v45
        }
        var v46 string = (("a" + "a") + "\n\t")
        _ = v46
    } else {
        v47 := GenS3{F0: ""}
        if true {
            var v48 int16 = 988
            v49 := GenS6{F0: false, F1: (v2 >> ((0) & 63))}
            v50 := GenS7{F0: 24, F1: 65535, F2: "foo bar"}
            _ = v48
            _ = v49
            _ = v50
        }
        _ = v47
    }
    _ = len(v5)
    v1 += -1
    var v51 int16 = 0
    v51 = v51
    fmt.Println(v2, v2)
    var v52 uint16 = 0
    for fi53 := 0; fi53 < 4; fi53++ {
        switch "" {
        case "\n\t":
            var v54 bool = false
            _ = v54
        case "a":
            var v55 int64 = v3
            _ = v55
        case "foo bar":
            for fi56 := 0; fi56 < 9; fi56++ {
                var v57 int32 = 327
                switch ("a" + "foo bar") {
                case "":
                    v58 := GenS8{F0: 255, F1: 4294967295, F2: (v57 / ((v57) | 1))}
                    _ = v58
                case "a":
                    v1 = (v57 ^ (v1 - (v57 - (v1 + v57))))
                case "hello":
                    if false {
                        fmt.Println(fi56, v2)
                    } else {
                        v2 = fi56
                        var v59 uint16 = v52
                        _ = v59
                    }
                    fmt.Println(v3, v51)
                    var v60 uint64 = 309
                    _ = v60
                default:
                    var v61 float32 = 0.1
                    _ = v61
                }
                if (84 == (v51 >> ((673) & 63))) {
                    switch "foo bar" {
                    case "":
                        var v62 float64 = 0.0
                        _ = v62
                    case "foo bar":
                        v63 := GenS9{F0: 65535, F1: 390, F2: 0}
                        _ = v63
                    case "\n\t":
                        v64 := GenS10{F0: true, F1: 18446744073709551615}
                        _ = v64
                    default:
                        v1 -= v57
                    }
                    if false {
                        var v65 uint16 = (^(v52))
                        _ = v65
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            var v66 int = (fi53 % ((v2) | 1))
                            _ = v66
                        case "\n\t":
                            delete(v4, -447)
                        }
                    }
                    switch "\n\t" {
                    case "a":
                        var v67 float64 = 3.14159
                        _ = v67
                    default:
                        v2 = (v2 % ((fi56) | 1))
                    }
                } else {
                    v57 = (v57 - v57)
                    v68 := GenS1{F0: (v52 - 65535), F1: "\n\t"}
                    _ = len(v4)
                    _ = v68
                }
                _ = fi56
                _ = v57
            }
        default:
            if false {
                if false {
                    v2 -= (^(fi53))
                    if len(v5) > 0 {
                        v5[0] = int(v1)
                    }
                }
                if true {
                    switch ("a" + "") {
                    case "":
                        if len(v5) > 0 {
                            v5[0] = (v2 | fi53)
                        }
                    case "foo bar":
                        if true {
                            if true {
                                for fi69 := 0; fi69 < 10; fi69++ {
                                    var v70 uint = 65535
                                    v1 += 388
                                    _ = fi69
                                    _ = v70
                                }
                                v71 := v4[32768]
                                for fi72 := 0; fi72 < 9; fi72++ {
                                    var v73 string = "a"
                                    var v74 int = fi53
                                    _ = fi72
                                    _ = v73
                                    _ = v74
                                }
                                _ = v71
                            } else {
                                var v75 uint16 = 0
                                switch "\n\t" {
                                case "\n\t":
                                    v76 := GenS1{F0: (v52 % (((v75 ^ (v75 & 42))) | 1)), F1: "foo bar"}
                                    _ = v76
                                }
                                v5 = append(v5, v2)
                                _ = v75
                            }
                            if (2 <= v52) {
                                fmt.Println(v52)
                            } else {
                                v3 += v3
                            }
                            v51 = (v51 & -195)
                        } else {
                            var v77 string = "foo bar"
                            _ = v77
                        }
                    default:
                        for fi78 := 0; fi78 < 13; fi78++ {
                            v79 := GenS11{F0: v3, F1: (!((v51 ^ v51) > -32768))}
                            _ = fi78
                            _ = v79
                        }
                    }
                    var v80 uint = 18446744073709551615
                    _ = v80
                } else {
                    if (!false) {
                        for fi81 := 0; fi81 < 1; fi81++ {
                            switch "\n\t" {
                            case "hello":
                                var v82 int8 = 1
                                _ = v82
                            case "a":
                                switch int(v52) {
                                case 951:
                                    var v83 uint8 = 127
                                    _ = v83
                                case 2147483647:
                                    fmt.Println(v51, fi81)
                                    var v84 float32 = (0.0 * (-2.25 * (-2.25 + ((0.1 + (1e10 / 1.0)) * -2.25))))
                                    var v85 uint32 = 42
                                    var v86 uint64 = 4294967295
                                    _ = v84
                                    _ = v85
                                    _ = v86
                                case -1:
                                    v52 = v52
                                }
                            default:
                                v87 := GenS12{F0: 255, F1: 0}
                                _ = v87
                            }
                            _ = fi81
                        }
                        var v88 int16 = (v51 ^ v51)
                        var v89 int16 = v88
                        _ = v88
                        _ = v89
                    } else {
                        if false {
                            if (2 == 1) {
                                v52 = 0
                                for fi90 := 0; fi90 < 18; fi90++ {
                                    var v91 uint16 = 2
                                    _ = fi90
                                    _ = v91
                                }
                            } else {
                                v1 += (-(v1))
                            }
                            var v92 int32 = -2147483648
                            _ = v92
                        } else {
                            for fi93 := 0; fi93 < 11; fi93++ {
                                var v94 int16 = v51
                                var v95 bool = true
                                _ = fi93
                                _ = v94
                                _ = v95
                            }
                        }
                    }
                }
                v96 := GenS13{F0: (v52 >> ((564) & 63))}
                _ = v96
            } else {
                v97 := GenS1{F0: v52, F1: ""}
                if (!false) {
                    if ((^(v52)) < (^(v52))) {
                        fmt.Println(v3)
                        delete(v4, fi53)
                    } else {
                        if ((v2 / (((-(v2))) | 1)) > 1) {
                            if false {
                                if (!(127 >= 127)) {
                                    v1 += (v1 / ((v1) | 1))
                                    var v98 int = 1
                                    _ = v98
                                } else {
                                    var v99 int32 = v1
                                    _ = v99
                                }
                            }
                        }
                    }
                    var v100 uint64 = 275
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v100
                } else {
                    fmt.Println(v3)
                    var v101 string = "\n\t"
                    _ = v101
                }
                v51 = (v51 << ((4294967295) & 63))
                _ = v97
            }
        }
        var v102 float64 = float64(((0.0 * ((3.14159 + 3.14159) / ((float32(0.1) + 1.5) / 3.14159))) * (float32((((0.0 / (3.14159 + (float64(float32(1e10)) - float64(1.5)))) - -2.25) * (0.1 - 0.0))) - 3.14159)))
        _ = fi53
        _ = v102
    }
    v52 -= 787
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
