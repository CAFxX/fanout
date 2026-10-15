package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 int64; F1 int64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint; F1 uint; F2 int }
func genFn0(p0 uint32, p1 int8, p2 int8) string {
    var v7 uint64 = 1
    var v8 uint = 18446744073709551615
    p1 = p1
    if true {
        fmt.Println(v8, p2, p0)
        v9 := []int{}
        p1 = (^(p1))
        _ = v9
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 uint16 = 32767
                v13 := GenS1{F0: uint8(p0), F1: 904, F2: 42}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    }
    v14 := GenS1{F0: uint8(v8), F1: 0, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v7
    _ = v8
    return ("hello" + "foo bar")
}

func main() {
    var v1 uint64 = 420
    var v2 int8 = -128
    v3 := make(map[string]string)
    var v4 string = "hello"
    v5 := make(map[uint32]int)
    var v6 uint64 = (v1 + v1)
    var v15 uint8 = 1
    var v16 int64 = 42
    fmt.Println(v16)
    var v17 int = -9223372036854775808
    var v18 string = "\n\t"
    if ((v15 / (((v15 ^ 128)) | 1)) == (v15 % (((^(v15))) | 1))) {
        v2 = (v2 / ((42) | 1))
    }
    var v19 uint16 = 0
    for fi20 := 0; fi20 < 1; fi20++ {
        if (-1 <= 517) {
            _ = len(v5)
            if (0 >= 42) {
                fmt.Println(v15, v18, fi20)
                v19 = (^(v19))
            }
            var v21 uint32 = 42
            _ = v21
        } else {
            if true {
                var v22 string = v18
                var v23 uint64 = (v6 | (v6 >> ((876) & 63)))
                _ = v22
                _ = v23
            } else {
                if (fi20 >= fi20) {
                    switch v17 {
                    case -9223372036854775808:
                        var v24 int64 = (v16 & 1)
                        _ = v24
                    case -2147483648:
                        if true {
                            var v25 uint = 1
                            v26 := GenS1{F0: uint8(v1), F1: 2147483647, F2: 42}
                            _ = v25
                            _ = v26
                        } else {
                            v27 := []int{}
                            delete(v5, 834)
                            for fi28 := 0; fi28 < 4; fi28++ {
                                switch 743 {
                                case 128:
                                    v15 -= (v15 >> ((255) & 63))
                                case -1:
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                default:
                                    fmt.Println(fi28, v1)
                                }
                                v29 := GenS1{F0: 1, F1: 0, F2: v15}
                                if (798 <= 42) {
                                    v15 -= 255
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                    var v30 bool = ((v16 + 0) >= (v16 | -371))
                                    v31 := GenS1{F0: 128, F1: 0, F2: v15}
                                    _ = len(v27)
                                    v32 := GenS2{F0: v1, F1: v15}
                                    v18 = v4
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                var v33 bool = true
                                _ = fi28
                                _ = v29
                                _ = v33
                            }
                            _ = v27
                        }
                    }
                }
                for fi34 := 0; fi34 < 5; fi34++ {
                    var v35 int64 = -1
                    if ((v16 | int64(v19)) < v35) {
                        fmt.Println(v6, v18, v1)
                        if true {
                            v36 := v3[v18]
                            v37 := GenS3{F0: 1}
                            v35 -= v16
                            v38 := GenS3{F0: (v19 & v19)}
                            var v39 int64 = (v16 << ((42) & 63))
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        v40 := GenS1{F0: v15, F1: 216, F2: (v15 & (v15 ^ v15))}
                        v41 := GenS1{F0: v15, F1: -1, F2: v15}
                        _ = v40
                        _ = v41
                    }
                    if false {
                        v42 := GenS3{F0: (v19 | v19)}
                        var v43 uint16 = v19
                        _ = v42
                        _ = v43
                    } else {
                        v44 := GenS4{F0: (v16 ^ (v35 << ((18446744073709551615) & 63)))}
                        switch (v18 + "a") {
                        case "":
                            v17 += -1
                            var v45 float64 = -2.25
                            switch "foo bar" {
                            case "hello":
                                v46 := []int{v17}
                                _ = v46
                            case "foo bar":
                                fmt.Println(fi20, v35, v17)
                            }
                            if false {
                                var v47 float64 = 0.0
                                var v48 int8 = (v2 + (v2 & v2))
                                _ = v47
                                _ = v48
                            } else {
                                v49 := GenS5{F0: v2, F1: 1}
                                delete(v5, 908)
                                _ = v49
                            }
                            _ = v45
                        case "\n\t":
                            v2 = (v2 ^ v2)
                        default:
                            v35 -= (v35 << ((0) & 63))
                        }
                        if (!true) {
                            var v50 uint8 = (v15 | (^(v15)))
                            var v51 float64 = (1e10 * 0.0)
                            _ = v50
                            _ = v51
                        }
                        _ = v44
                    }
                    if (358 != 42) {
                        v4 = v4
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v4)
                        case "\n\t":
                            fmt.Println(fi34)
                        case "a":
                            var v52 string = ((v4 + v18) + v18)
                            _ = v52
                        default:
                            v53 := GenS6{F0: "foo bar"}
                            _ = v53
                        }
                        v54 := GenS4{F0: v35}
                        if false {
                            v15 += v15
                            v55 := GenS1{F0: v15, F1: 0, F2: v15}
                            var v56 int64 = (v16 ^ v35)
                            _ = v55
                            _ = v56
                        } else {
                            v18 = "hello"
                            v17 += v17
                            v57 := []int{(fi34 % ((fi34) | 1)), -2147483648, (^(v17)), int(v1)}
                            _ = v57
                        }
                        v58 := []int{}
                        _ = v54
                        _ = v58
                    } else {
                        v2 = -1
                        var v59 int8 = (v2 ^ v2)
                        _ = v59
                    }
                    if ((fi34 + (fi34 ^ fi20)) > int(v6)) {
                        v16 -= 9223372036854775807
                    }
                    switch v4 {
                    case "\n\t":
                        switch "a" {
                        case "a":
                            if true {
                                v60 := []int{127}
                                fmt.Println(fi34)
                                switch fi20 {
                                case 32768:
                                    v61 := GenS7{F0: true, F1: 65535}
                                    _ = v61
                                case 128:
                                    v62 := v5[65535]
                                    _ = v62
                                default:
                                    v15 = 42
                                }
                                _ = v60
                            } else {
                                fmt.Println(v35, v4)
                            }
                        default:
                            var v63 uint = 255
                            _ = v63
                        }
                    default:
                        v64 := []int{32767, -811, 2147483647}
                        _ = v64
                    }
                    var v65 uint16 = (^(v19))
                    _ = fi34
                    _ = v35
                    _ = v65
                }
            }
        }
        v1 += v6
        _ = fi20
    }
    var v66 string = ""
    v67 := []int{v17}
    for fi68 := 0; fi68 < 7; fi68++ {
        switch (fi68 << ((2) & 63)) {
        case 9223372036854775807:
            switch "foo bar" {
            case "a":
                switch fi68 {
                case -9223372036854775808:
                    if true {
                        v6 += 2
                    }
                case -26:
                    if true {
                        var v69 uint32 = 0
                        v17 = v17
                        _ = v69
                    }
                case 9223372036854775807:
                    if false {
                        var v70 float32 = 3.14159
                        switch v4 {
                        case "\n\t":
                            var v71 bool = (431 >= 255)
                            _ = v71
                        default:
                            if false {
                                switch ("hello" + "\n\t") {
                                case "hello":
                                    fmt.Println(v66, fi68, v4)
                                case "a":
                                    switch v17 {
                                    case -455:
                                        v1 = v6
                                    default:
                                        if (!(v18 != v4)) {
                                            var v72 int = (v17 * 107)
                                            v70 = v70
                                            for fi73 := 0; fi73 < 12; fi73++ {
                                                v16 += (v16 << ((65535) & 63))
                                                _ = fi73
                                            }
                                            _ = v72
                                        } else {
                                            if true {
                                                v2 = int8(v16)
                                                v74 := GenS5{F0: (v2 / ((v2) | 1)), F1: 1}
                                                fmt.Println(fi68, v6)
                                                _ = v74
                                            }
                                            if (1 <= (v19 << ((2) & 63))) {
                                                v75 := GenS8{F0: false}
                                                v76 := GenS2{F0: 4294967295, F1: v15}
                                                _ = v75
                                                _ = v76
                                            }
                                            var v77 uint8 = 255
                                            _ = v77
                                        }
                                    }
                                }
                            } else {
                                v78 := GenS4{F0: v16}
                                if (2147483647 >= 65535) {
                                    v2 = (v2 + (v2 >> ((725) & 63)))
                                    for fi79 := 0; fi79 < 9; fi79++ {
                                        var v80 string = v66
                                        if ("a" != v66) {
                                            var v81 float64 = 1.5
                                            _ = v81
                                        } else {
                                            v82 := GenS9{F0: v16, F1: false}
                                            _ = v82
                                        }
                                        var v83 int32 = 42
                                        _ = fi79
                                        _ = v80
                                        _ = v83
                                    }
                                }
                                _ = v78
                            }
                        }
                        if len(v67) > 0 {
                            v67[0] = (fi68 / ((v17) | 1))
                        }
                        _ = v70
                    } else {
                        var v84 float32 = float32((-2.25 - 3.14159))
                        _ = v84
                    }
                }
            case "hello":
                if true {
                    v18 = v18
                }
            case "\n\t":
                v17 = 601
            }
        case -426:
            if (v15 >= v15) {
                if (18446744073709551615 > v1) {
                    var v85 uint64 = 18446744073709551615
                    fmt.Println(v17, v19)
                    var v86 float32 = 0.0
                    _ = v85
                    _ = v86
                } else {
                    v67 = append(v67, (-(v17)))
                    fmt.Println(v16, v16)
                }
                var v87 string = (v66 + ("a" + v4))
                v16 -= -9223372036854775808
                _ = v87
            } else {
                for fi88 := 0; fi88 < 17; fi88++ {
                    switch (v66 + v4) {
                    case "\n\t":
                        v89 := GenS10{F0: (v16 ^ v16), F1: (-(v16))}
                        _ = v89
                    case "foo bar":
                        if true {
                            var v90 uint16 = (v19 / (((v19 >> ((392) & 63))) | 1))
                            v90 += (v19 >> ((641) & 63))
                            _ = v90
                        }
                    case "hello":
                        var v91 bool = ((!(true && (!((v15 ^ (v15 << ((4294967295) & 63))) > (v15 * (v15 >> ((uint(v19)) & 63))))))) && (!(false || (false && ((false && (-9223372036854775808 >= v16)) && false)))))
                        _ = v91
                    default:
                        if len(v67) > 0 {
                            v67[0] = fi88
                        }
                    }
                    _ = fi88
                }
                if (!((true && false) || (false || ((^(v1)) == 86)))) {
                    v16 = v16
                }
            }
        case 42:
            fmt.Println(v66, v2, v1)
        default:
            var v92 uint = 893
            _ = v92
        }
        _ = fi68
    }
    for fi93 := 0; fi93 < 12; fi93++ {
        v66 = "foo bar"
        if true {
            switch (v18 + v4) {
            case "hello":
                if false {
                    v19 -= v19
                    for fi94 := 0; fi94 < 17; fi94++ {
                        fmt.Println(v6, v6)
                        _ = fi94
                    }
                    var v95 int16 = 32767
                    _ = v95
                }
            }
            v16 = (^(v16))
            if len(v67) > 0 {
                v67[0] = fi93
            }
        } else {
            fmt.Println(v6, v66, v17)
            var v96 bool = (true && false)
            _ = v96
        }
        for fi97 := 0; fi97 < 11; fi97++ {
            v19 -= uint16(v6)
            v17 += v17
            v6 += (v1 % ((0) | 1))
            _ = fi97
        }
        _ = fi93
    }
    fmt.Println(v17, v17, v66)
    switch v17 {
    case -580:
        switch v17 {
        case 9223372036854775807:
            switch v4 {
            case "":
                v16 = v16
            case "foo bar":
                v1 += 42
            case "\n\t":
                for fi98 := 0; fi98 < 12; fi98++ {
                    fmt.Println(v2, v17, v1)
                    if true {
                        v99 := GenS11{F0: (65535 > 690)}
                        var v100 float64 = 0.1
                        _ = v100
                        _ = v99
                    } else {
                        var v101 int32 = 1
                        if true {
                            switch 0 {
                            case 9223372036854775807:
                                var v102 uint16 = (^(v19))
                                _ = v102
                            case 0:
                                var v103 uint = 1
                                _ = v103
                            default:
                                if false {
                                    v104 := GenS12{F0: 1, F1: 436, F2: fi98}
                                    var v105 int8 = (v2 << ((847) & 63))
                                    v19 = v19
                                    _ = v104
                                    _ = v105
                                }
                            }
                        } else {
                            for fi106 := 0; fi106 < 17; fi106++ {
                                fmt.Println(v6, v1, v4)
                                _ = fi106
                            }
                        }
                        _ = v101
                    }
                    var v107 bool = (((-2.25 - 1e10) / 0.1) < -2.25)
                    _ = fi98
                    _ = v107
                }
            }
        }
    case 2147483647:
        v2 = (v2 << ((565) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v67:", len(v67), v67)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
}
