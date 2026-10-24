package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 string }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int8; F1 uint8; F2 int32 }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint }
type GenS10 struct { F0 uint; F1 int64; F2 int16 }
type GenS11 struct { F0 uint; F1 string }
type GenS12 struct { F0 uint8; F1 uint }
type GenS13 struct { F0 int16; F1 bool }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int64 }

func main() {
    var v1 uint16 = 42
    var v2 uint16 = (^(v1))
    var v3 uint32 = uint32(v2)
    v4 := make(map[uint32]string)
    v1 = 348
    v5 := []int{2147483647, -982, 2147483647, 49}
    if (42 <= -1) {
        var v6 int8 = 0
        _ = v6
    }
    if false {
        for fi7 := 0; fi7 < 14; fi7++ {
            for fi8 := 0; fi8 < 15; fi8++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi8
            }
            if (!true) {
                v9 := GenS1{F0: 124}
                if (42 == -177) {
                    v10 := GenS1{F0: 180}
                    _ = len(v4)
                    var v11 uint = 1
                    _ = v10
                    _ = v11
                }
                _ = v9
            } else {
                v12 := GenS1{F0: 127}
                v2 = (^(v1))
                if true {
                    _ = len(v5)
                    if (0.0 != (0.1 + 0.0)) {
                        v5 = append(v5, (fi7 + -1))
                        _ = len(v4)
                    }
                    if len(v5) > 0 {
                        v5[0] = (^(fi7))
                    }
                }
                _ = v12
            }
            v13 := GenS2{F0: false, F1: 0}
            _ = fi7
            _ = v13
        }
        _ = len(v4)
        v14 := GenS1{F0: 1}
        _ = v14
    } else {
        v1 -= (^(v2))
        v15 := GenS1{F0: uint8(v2)}
        _ = v15
    }
    var v16 int8 = -1
    _ = len(v4)
    var v17 bool = false
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int64 = -1
        switch (("a" + "") + "hello") {
        case "":
            var v20 uint32 = v3
            _ = v20
        case "hello":
            v2 = (v1 * (v2 % ((299) | 1)))
        case "\n\t":
            v21 := GenS2{F0: true, F1: 127}
            _ = v21
        default:
            v1 = (^(v1))
        }
        _ = fi18
        _ = v19
    }
    v3 = 4294967295
    switch 0 {
    case -9223372036854775808:
        v2 -= v2
    default:
        var v22 float32 = 0.1
        _ = v22
    }
    var v23 float32 = (((0.1 - ((1e10 * 3.14159) + 0.0)) * float32(1e10)) + 0.0)
    var v24 int8 = v16
    for fi25 := 0; fi25 < 4; fi25++ {
        switch "foo bar" {
        case "foo bar":
            if true {
                switch "\n\t" {
                case "":
                    switch "a" {
                    case "hello":
                        for fi26 := 0; fi26 < 17; fi26++ {
                            var v27 uint16 = 255
                            _ = fi26
                            _ = v27
                        }
                    }
                }
                if true {
                    if v17 {
                        switch ("" + "foo bar") {
                        case "\n\t":
                            switch "a" {
                            case "\n\t":
                                switch ("\n\t" + "\n\t") {
                                case "\n\t":
                                    if v17 {
                                        v28 := GenS1{F0: 2}
                                        if v17 {
                                            v23 += ((v23 + (float32(1.5) * -2.25)) / v23)
                                            v3 = 1
                                        } else {
                                            v23 -= (v23 / v23)
                                        }
                                        for fi29 := 0; fi29 < 4; fi29++ {
                                            v30 := GenS3{F0: (fi25 ^ fi25)}
                                            fmt.Println(fi25, v23, v17)
                                            var v31 float64 = 0.1
                                            _ = fi29
                                            _ = v30
                                            _ = v31
                                        }
                                        _ = v28
                                    }
                                case "a":
                                    var v32 uint = 4294967295
                                    _ = v32
                                default:
                                    var v33 int8 = v24
                                    _ = v33
                                }
                            case "":
                                v34 := GenS4{F0: 89}
                                _ = v34
                            case "foo bar":
                                var v35 bool = ((0 >= 1) && ((^(v24)) > v24))
                                _ = v35
                            }
                        case "":
                            if ((!v17) || v17) {
                                fmt.Println(v23)
                            }
                        default:
                            var v36 float32 = (0.0 / float32(1.5))
                            _ = v36
                        }
                        if true {
                            switch "a" {
                            case "foo bar":
                                v37 := GenS2{F0: (30 != 4294967295), F1: 151}
                                _ = v37
                            case "hello":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        }
                    } else {
                        switch "" {
                        case "\n\t":
                            var v38 uint16 = (v1 * v1)
                            _ = v38
                        case "a":
                            v2 -= (v1 - (v2 >> ((42) & 63)))
                        case "":
                            v2 = v1
                        }
                        if ((^(fi25)) >= fi25) {
                            var v39 int16 = -852
                            _ = v39
                        } else {
                            if ((v1 % (((^(v2))) | 1)) != 1) {
                                _ = len(v4)
                            } else {
                                if false {
                                    var v40 int = fi25
                                    if true {
                                        v41 := GenS5{F0: 255, F1: (v1 % ((uint16(v3)) | 1)), F2: ""}
                                        v42 := GenS1{F0: 132}
                                        _ = v41
                                        _ = v42
                                    } else {
                                        v24 = (v16 % (((-(v16))) | 1))
                                        var v43 float32 = v23
                                        _ = v43
                                    }
                                    if (152 > 255) {
                                        var v44 uint32 = (v3 % ((v3) | 1))
                                        _ = v44
                                    } else {
                                        if ((3.14159 - (1e10 + 1.5)) > 0.0) {
                                            v23 -= v23
                                            if (1 < 1) {
                                                for fi45 := 0; fi45 < 15; fi45++ {
                                                    v40 = fi25
                                                    var v46 float64 = 1.5
                                                    _ = fi45
                                                    _ = v46
                                                }
                                                v23 = v23
                                            } else {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v47 int = v40
                                                    _ = v47
                                                }
                                                _ = len(v5)
                                            }
                                            var v48 uint64 = 18446744073709551615
                                            _ = v48
                                        } else {
                                            var v49 bool = v17
                                            v50 := GenS6{F0: (fi25 >> ((366) & 63)), F1: 4294967295}
                                            if true {
                                                v49 = (724 == -759)
                                                var v51 float32 = v23
                                                _ = v51
                                            }
                                            _ = v49
                                            _ = v50
                                        }
                                        v52 := GenS7{F0: (v24 / ((v16) | 1)), F1: 255, F2: -1}
                                        v53 := GenS8{F0: 739, F1: 0}
                                        _ = v52
                                        _ = v53
                                    }
                                    _ = v40
                                } else {
                                    var v54 string = ("hello" + "")
                                    if (18446744073709551615 <= 662) {
                                        _ = len(v5)
                                        v2 = (^(v1))
                                        if (((false && false) && v17) && (-9223372036854775808 == -313)) {
                                            if true {
                                                var v55 int32 = 1
                                                switch "hello" {
                                                case "":
                                                    v4[(v3 | (v3 ^ v3))] = v54
                                                }
                                                v56 := GenS9{F0: (!v17), F1: 2, F2: 1}
                                                _ = v55
                                                _ = v56
                                            } else {
                                                var v57 int = 32767
                                                if (fi25 <= (v57 + (fi25 * (-(fi25))))) {
                                                    delete(v4, 4294967295)
                                                }
                                                _ = v57
                                            }
                                            if v17 {
                                                var v58 uint32 = 2
                                                _ = v58
                                            } else {
                                                v59 := GenS10{F0: 155, F1: 9223372036854775807, F2: 32767}
                                                v16 = v24
                                                var v60 uint16 = 32767
                                                _ = v59
                                                _ = v60
                                            }
                                            for fi61 := 0; fi61 < 6; fi61++ {
                                                v62 := v4[42]
                                                v4[65535] = v62
                                                v63 := GenS7{F0: 119, F1: 1, F2: -1}
                                                _ = fi61
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                    }
                                    _ = v54
                                }
                                switch (fi25 ^ -2147483648) {
                                case -352:
                                    var v64 int16 = 42
                                    _ = v64
                                case -784:
                                    v17 = v17
                                }
                                v65 := GenS1{F0: 0}
                                _ = v65
                            }
                            v66 := GenS11{F0: 745, F1: ("a" + "\n\t")}
                            _ = len(v4)
                            _ = v66
                        }
                        for fi67 := 0; fi67 < 15; fi67++ {
                            v68 := v4[0]
                            v24 += v16
                            if v17 {
                                if v17 {
                                    var v69 bool = (!(!v17))
                                    _ = v69
                                }
                                var v70 int = (fi25 - 127)
                                _ = v70
                            } else {
                                var v71 float32 = (v23 * 0.1)
                                switch (("hello" + v68) + v68) {
                                case "hello":
                                    v72 := v4[v3]
                                    _ = v72
                                default:
                                    var v73 string = (("\n\t" + ("hello" + ((("\n\t" + ("" + v68)) + "foo bar") + "hello"))) + v68)
                                    _ = v73
                                }
                                _ = v71
                            }
                            var v74 uint64 = 2
                            v17 = (true || (42 >= 32767))
                            _ = fi67
                            _ = v68
                            _ = v74
                        }
                    }
                    for fi75 := 0; fi75 < 15; fi75++ {
                        var v76 int32 = -2147483648
                        _ = fi75
                        _ = v76
                    }
                } else {
                    var v77 uint = 486
                    v3 = v3
                    v3 += (v3 ^ v3)
                    _ = v77
                }
                if true {
                    v78 := GenS4{F0: 1}
                    _ = v78
                } else {
                    if ((v3 | (v3 ^ (v3 & uint32(v2)))) == (v3 / ((v3) | 1))) {
                        v23 -= (((float32(0.1) - (0.0 + v23)) + v23) / ((((float32((3.14159 * 1.5)) + 3.14159) - v23) + float32((((-2.25 * 0.1) * (0.1 + (0.0 - ((0.0 / 1e10) / 0.1)))) + -2.25))) - 0.0))
                        switch (^(fi25)) {
                        case -1:
                            for fi79 := 0; fi79 < 15; fi79++ {
                                delete(v4, (v3 / (((v3 % ((v3) | 1))) | 1)))
                                _ = fi79
                            }
                        case 997:
                            for fi80 := 0; fi80 < 9; fi80++ {
                                if v17 {
                                    var v81 int16 = 123
                                    _ = len(v4)
                                    _ = v81
                                } else {
                                    fmt.Println(fi80, fi25, v23)
                                    if (-1 > 1000) {
                                        v82 := GenS12{F0: 111, F1: uint(v3)}
                                        switch ("" + "foo bar") {
                                        case "\n\t":
                                            v23 = (v23 - v23)
                                        case "a":
                                            if (!(3.14159 > ((float32(float64(-2.25)) * (v23 - float32(1e10))) / ((0.1 + -2.25) / float32(((((3.14159 / 1e10) - 3.14159) / 0.1) / float64(((-2.25 - (v23 * ((v23 * float32(0.0)) + 1e10))) / 1e10)))))))) {
                                                var v83 uint32 = (v3 | (v3 + v3))
                                                if v17 {
                                                    var v84 int8 = v16
                                                    v85 := GenS13{F0: int16(fi25), F1: v17}
                                                    fmt.Println(v2)
                                                    v83 = 4294967295
                                                    v24 += v24
                                                    v86 := GenS14{F0: 4294967295}
                                                    if v17 {
                                                        v16 = -1
                                                    } else {
                                                        var v87 uint8 = 255
                                                        _ = v87
                                                    }
                                                    _ = v84
                                                    _ = v85
                                                    _ = v86
                                                }
                                                _ = v83
                                            } else {
                                                var v88 int16 = -32768
                                                _ = v88
                                            }
                                            v1 -= (v2 << ((65535) & 63))
                                            var v89 int32 = -2147483648
                                            _ = v89
                                        case "":
                                            var v90 uint8 = 1
                                            _ = v90
                                        default:
                                            var v91 uint16 = v1
                                            _ = v91
                                        }
                                        switch (fi80 ^ -785) {
                                        case 128:
                                            var v92 uint16 = 42
                                            _ = v92
                                        default:
                                            fmt.Println(v1)
                                        }
                                        _ = v82
                                    } else {
                                        var v93 int16 = -32768
                                        if (false && true) {
                                            var v94 uint32 = 1
                                            v95 := GenS15{F0: -9223372036854775808}
                                            _ = v94
                                            _ = v95
                                        }
                                        v96 := GenS6{F0: (fi25 >> ((4294967295) & 63)), F1: 42}
                                        _ = v93
                                        _ = v96
                                    }
                                }
                                _ = fi80
                            }
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v23 = ((v23 * 0.0) + v23)
                    }
                    v3 -= (v3 * (v3 & v3))
                }
            } else {
                for fi97 := 0; fi97 < 15; fi97++ {
                    var v98 uint64 = 42
                    _ = fi97
                    _ = v98
                }
                v3 += v3
                for fi99 := 0; fi99 < 10; fi99++ {
                    v24 = v16
                    v100 := v4[v3]
                    _ = fi99
                    _ = v100
                }
            }
        default:
            var v101 uint32 = (v3 / (((v3 | v3)) | 1))
            _ = v101
        }
        _ = fi25
    }
    fmt.Println(v2, v17, v24)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
