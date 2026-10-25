package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 int16; F1 int64 }
type GenS9 struct { F0 int8; F1 int64 }
type GenS10 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint64) string {
    var v6 int16 = -1
    v7 := GenS1{F0: -9223372036854775808, F1: 83, F2: 175}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ("" + "\n\t")
}
func genFn1() bool {
    v8 := []int{716, -1}
    var v9 uint32 = 42
    v10 := GenS2{F0: 18446744073709551615, F1: (^(v9)), F2: -20}
    switch ("\n\t" + ("foo bar" + "foo bar")) {
    case "a":
        v11 := GenS2{F0: 754, F1: 449, F2: 92}
        _ = v11
    case "foo bar":
        switch "hello" {
        case "":
            v9 += (v9 >> ((818) & 63))
        default:
            var v12 uint32 = v9
            _ = v12
        }
    case "hello":
        for fi13 := 0; fi13 < 13; fi13++ {
            if true {
                _ = len(v8)
            } else {
                var v14 int = fi13
                v9 = (v9 << ((42) & 63))
                if (int32(fi13) == -2147483648) {
                    if false {
                        var v15 string = "a"
                        var v16 string = "hello"
                        _ = v15
                        _ = v16
                    } else {
                        v14 = fi13
                        if false {
                            fmt.Println(v14, v14, v14)
                        }
                    }
                } else {
                    if len(v8) > 0 {
                        v8[0] = fi13
                    }
                    fmt.Println(v14, v9, v14)
                }
                _ = v14
            }
            _ = fi13
        }
    default:
        var v17 string = "foo bar"
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        v9 = v9
        v9 = v9
        _ = fi18
    }
    _ = v10
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 int32 = 42
    var v2 uint64 = 2
    var v3 uint = 1
    v4 := []int{}
    v5 := make(map[int]uint32)
    fmt.Println(v2, v1, v3)
    var v19 int8 = -41
    if (false && true) {
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 uint8 = 42
            delete(v5, fi20)
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 uint8 = (v21 | v21)
                var v24 int32 = v1
                v25 := GenS1{F0: (fi22 | (fi22 * 32767)), F1: 255, F2: v23}
                v21 = v23
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi20
            _ = v21
        }
        var v26 int16 = 1
        var v27 int8 = (v19 * (v19 % ((-1) | 1)))
        _ = v26
        _ = v27
    }
    var v28 string = ""
    if (true || false) {
        for fi29 := 0; fi29 < 6; fi29++ {
            var v30 float64 = 0.1
            if false {
                var v31 uint8 = 2
                v32 := GenS1{F0: 42, F1: 255, F2: v31}
                if len(v4) > 0 {
                    v4[0] = fi29
                }
                _ = v31
                _ = v32
            }
            _ = fi29
            _ = v30
        }
        for fi33 := 0; fi33 < 6; fi33++ {
            v3 -= 1
            for fi34 := 0; fi34 < 6; fi34++ {
                if true {
                    fmt.Println(fi33, v19, v19)
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v1 += (v1 + v1)
                if true {
                    var v35 uint = 579
                    _ = v35
                } else {
                    v5[32768] = 274
                }
                _ = fi34
            }
            _ = fi33
        }
        delete(v5, 783)
    } else {
        for fi36 := 0; fi36 < 11; fi36++ {
            var v37 uint16 = 2
            var v38 int32 = v1
            v28 = ""
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    delete(v5, 2147483647)
    v1 += (^(v1))
    for fi39 := 0; fi39 < 11; fi39++ {
        if (((v1 % ((int32(v3)) | 1)) > v1) && (false || ((fi39 <= -2147483648) || true))) {
            if false {
                var v40 int64 = -1
                v40 -= 42
                _ = v40
            } else {
                v41 := GenS1{F0: -328, F1: 32767, F2: 255}
                v42 := GenS1{F0: (fi39 | (-(fi39))), F1: 2, F2: 255}
                _ = v41
                _ = v42
            }
            v1 = (v1 << ((v3) & 63))
        } else {
            if false {
                for fi43 := 0; fi43 < 4; fi43++ {
                    var v44 int16 = 1000
                    if true {
                        fmt.Println(v19, fi39)
                        var v45 int = (fi39 & fi43)
                        var v46 int16 = v44
                        _ = v45
                        _ = v46
                    }
                    _ = fi43
                    _ = v44
                }
                if true {
                    for fi47 := 0; fi47 < 9; fi47++ {
                        if true {
                            v4 = append(v4, (fi47 * -470))
                            v48 := GenS3{F0: (-1 < -1), F1: 42}
                            v49 := GenS1{F0: fi39, F1: 32767, F2: 2}
                            _ = len(v5)
                            _ = v48
                            _ = v49
                        }
                        delete(v5, (^(fi39)))
                        if (94 == 255) {
                            var v50 int64 = 714
                            v51 := GenS4{F0: int32(v3)}
                            _ = v50
                            _ = v51
                        } else {
                            if len(v4) > 0 {
                                v4[0] = fi47
                            }
                            fmt.Println(v19, fi47, v1)
                        }
                        v52 := v5[9223372036854775807]
                        v1 = v1
                        _ = fi47
                        _ = v52
                    }
                    var v53 uint16 = 65535
                    _ = v53
                } else {
                    var v54 int64 = -402
                    for fi55 := 0; fi55 < 2; fi55++ {
                        v3 += v3
                        _ = fi55
                    }
                    switch fi39 {
                    case 128:
                        fmt.Println(v28, v3, fi39)
                    }
                    _ = v54
                }
                if false {
                    if true {
                        switch (v28 + (v28 + (v28 + v28))) {
                        case "hello":
                            v56 := GenS2{F0: (v2 / ((4294967295) | 1)), F1: 2147483648, F2: v19}
                            _ = v56
                        case "":
                            v1 += v1
                        default:
                            var v57 string = v28
                            _ = v57
                        }
                    }
                    v3 = (v3 << ((96) & 63))
                } else {
                    fmt.Println(v3, v28)
                }
            } else {
                _ = len(v5)
            }
        }
        fmt.Println(v28, v3, v3)
        _ = fi39
    }
    v3 += (v3 * 42)
    if true {
        v2 = v2
        var v58 int = int(v1)
        for fi59 := 0; fi59 < 10; fi59++ {
            for fi60 := 0; fi60 < 10; fi60++ {
                if (!("foo bar" > v28)) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi60
            }
            _ = fi59
        }
        _ = v58
    }
    for fi61 := 0; fi61 < 12; fi61++ {
        if true {
            v1 = (^(v1))
        }
        _ = fi61
    }
    var v62 int64 = 1
    for fi63 := 0; fi63 < 13; fi63++ {
        fmt.Println(v28)
        for fi64 := 0; fi64 < 13; fi64++ {
            v28 = "hello"
            if true {
                if (0 != 985) {
                    v1 = (v1 - (v1 * v1))
                }
                _ = len(v4)
                var v65 int32 = v1
                _ = v65
            } else {
                v66 := GenS5{F0: "\n\t", F1: 4294967295}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v66
            }
            var v67 uint16 = 659
            fmt.Println(v2, v2)
            _ = fi64
            _ = v67
        }
        _ = fi63
    }
    for fi68 := 0; fi68 < 13; fi68++ {
        if true {
            if true {
                v4 = append(v4, 128)
                for fi69 := 0; fi69 < 8; fi69++ {
                    if true {
                        v28 = ((v28 + (v28 + v28)) + ((v28 + v28) + (v28 + "foo bar")))
                    } else {
                        if (uint64(v3) >= (v2 & 42)) {
                            if ((^(v3)) < (v3 & v3)) {
                                var v70 int = -2147483648
                                _ = v70
                            } else {
                                v71 := GenS1{F0: fi68, F1: 255, F2: 26}
                                var v72 uint64 = v2
                                var v73 uint64 = v72
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                            fmt.Println(v1, v28)
                            v74 := GenS6{F0: v19}
                            _ = v74
                        } else {
                            if ((((v3 ^ 64) > v3) && (((42 >= 2) && (501 == 796)) || (-2.25 != (1e10 - 0.1)))) && false) {
                                v5[fi68] = 42
                                v19 = (v19 + (v19 / ((-21) | 1)))
                                v75 := GenS3{F0: ((v28 + "\n\t") <= ((v28 + v28) + (v28 + ((v28 + v28) + v28)))), F1: 127}
                                _ = v75
                            } else {
                                v28 = v28
                                v5[(fi68 * (fi69 ^ int(v2)))] = 1
                                fmt.Println(v62, fi69, fi68)
                            }
                            v76 := GenS7{F0: v1, F1: ("\n\t" + v28)}
                            fmt.Println(v28, fi69, v62)
                            _ = v76
                        }
                        switch v28 {
                        case "\n\t":
                            switch fi69 {
                            case -686:
                                switch 42 {
                                case -835:
                                    if (!(!((v2 + 2) >= v2))) {
                                        v77 := GenS4{F0: 213}
                                        if (false && false) {
                                            v2 += v2
                                            fmt.Println(v28, v3, v1)
                                            v78 := GenS1{F0: 4294967295, F1: 29, F2: 255}
                                            _ = v78
                                        } else {
                                            v79 := GenS4{F0: int32(v3)}
                                            _ = v79
                                        }
                                        v2 = (v2 + (v2 << (((v3 & v3)) & 63)))
                                        _ = v77
                                    } else {
                                        v80 := v5[9223372036854775807]
                                        if false {
                                            v28 = ""
                                            _ = len(v5)
                                            var v81 int64 = v62
                                            _ = v81
                                        } else {
                                            _ = len(v5)
                                            v82 := v5[9223372036854775807]
                                            var v83 uint32 = (v80 ^ v80)
                                            _ = v82
                                            _ = v83
                                        }
                                        _ = v80
                                    }
                                }
                            case -1:
                                v2 = v2
                            case 32767:
                                var v84 uint64 = (v2 - 0)
                                _ = v84
                            default:
                                v19 = (-(v19))
                            }
                        default:
                            fmt.Println(v28, v1)
                        }
                    }
                    if (42 > -1) {
                        v85 := GenS8{F0: 42, F1: v62}
                        if len(v4) > 0 {
                            v4[0] = (fi68 | fi69)
                        }
                        fmt.Println(fi69, v19, v28)
                        var v86 int32 = int32(v62)
                        _ = v85
                        _ = v86
                    } else {
                        var v87 string = v28
                        if (false && true) {
                            var v88 float32 = 3.14159
                            var v89 int64 = v62
                            v3 -= 0
                            _ = v88
                            _ = v89
                        }
                        fmt.Println(v2)
                        _ = v87
                    }
                    fmt.Println(v3, v19)
                    switch (((v28 + v28) + v28) + v28) {
                    case "foo bar":
                        v5[899] = 65535
                    case "hello":
                        var v90 uint16 = 2
                        _ = v90
                    case "a":
                        var v91 uint16 = 1
                        _ = v91
                    }
                    _ = fi69
                }
                v92 := GenS9{F0: (-(v19)), F1: v62}
                _ = v92
            }
            v1 = (v1 % (((v1 % ((-513) | 1))) | 1))
            fmt.Println(v1, v1, v3)
        }
        switch fi68 {
        case 0:
            for fi93 := 0; fi93 < 4; fi93++ {
                v94 := GenS9{F0: 112, F1: v62}
                var v95 int64 = v62
                _ = fi93
                _ = v94
                _ = v95
            }
        case 87:
            var v96 uint32 = 2147483648
            _ = v96
        case 751:
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
        }
        if true {
            for fi97 := 0; fi97 < 13; fi97++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if true {
                    fmt.Println(v1, v19)
                    var v98 float64 = -2.25
                    _ = v98
                } else {
                    if (!(v19 > (v19 | v19))) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if false {
                            v3 += v3
                        } else {
                            switch "a" {
                            case "hello":
                                delete(v5, int(v2))
                            default:
                                delete(v5, (fi68 ^ 127))
                            }
                        }
                        var v99 uint32 = 1
                        switch "" {
                        case "hello":
                            var v100 uint8 = 1
                            _ = v100
                        default:
                            v1 = v1
                        }
                        switch 0 {
                        case 9223372036854775807:
                            if true {
                                v2 = v2
                            }
                        case 310:
                            v2 = (^(v2))
                        case 32768:
                            v101 := GenS10{F0: v2}
                            _ = v101
                        }
                        _ = v99
                    } else {
                        v102 := v5[fi68]
                        var v103 float64 = 0.1
                        fmt.Println(fi68, v103, v103)
                        _ = v102
                        _ = v103
                    }
                }
                fmt.Println(v1)
                _ = fi97
            }
        }
        _ = fi68
    }
    if false {
        if false {
            if false {
                v104 := GenS7{F0: v1, F1: v28}
                fmt.Println(v3, v62)
                switch ((v28 + "hello") + v28) {
                case "":
                    var v105 int32 = 1
                    _ = v105
                }
                _ = v104
            }
            v19 -= -1
            for fi106 := 0; fi106 < 9; fi106++ {
                for fi107 := 0; fi107 < 20; fi107++ {
                    v3 = (v3 & (^(v3)))
                    _ = fi107
                }
                if true {
                    fmt.Println(v62)
                } else {
                    for fi108 := 0; fi108 < 18; fi108++ {
                        v109 := GenS5{F0: (v28 + (("foo bar" + v28) + "a")), F1: (v2 << (((v3 & (^(v3)))) & 63))}
                        var v110 int = (fi106 + fi108)
                        switch "foo bar" {
                        case "a":
                            if false {
                                fmt.Println(v2, v3, v3)
                                if len(v4) > 0 {
                                    v4[0] = (^(fi108))
                                }
                                if ((1e10 + 3.14159) < (1.5 - 1e10)) {
                                    v110 = (fi106 | int(v19))
                                    var v111 uint32 = 4294967295
                                    _ = v111
                                } else {
                                    v1 = (v1 - v1)
                                    var v112 float32 = (0.0 - float32(1.5))
                                    _ = v112
                                }
                            }
                        }
                        _ = fi108
                        _ = v109
                        _ = v110
                    }
                }
                fmt.Println(v3, v1)
                _ = fi106
            }
        } else {
            for fi113 := 0; fi113 < 18; fi113++ {
                if (fi113 == (^(fi113))) {
                    switch 202 {
                    case 0:
                        v62 = (v62 + 0)
                    case 9223372036854775807:
                        if (!true) {
                            fmt.Println(fi113, v2)
                        } else {
                            v62 -= (-(v62))
                            v2 = 18446744073709551615
                            var v114 int32 = v1
                            _ = v114
                        }
                    case 1:
                        fmt.Println(v19)
                    }
                    v1 += (v1 / ((v1) | 1))
                    v115 := GenS3{F0: (0.0 >= ((3.14159 - float64(1.5)) + 1.5)), F1: 0}
                    _ = v115
                } else {
                    fmt.Println(v1, fi113, v2)
                    v3 += (v3 >> (((v3 / ((v3) | 1))) & 63))
                }
                var v116 uint = v3
                _ = fi113
                _ = v116
            }
            fmt.Println(v19)
        }
    } else {
        v117 := GenS3{F0: (!(!true)), F1: 1}
        _ = v117
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v62
}
