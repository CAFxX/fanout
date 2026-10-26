package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 int; F1 bool; F2 int32 }
type GenS10 struct { F0 bool; F1 int32; F2 string }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 int16; F1 uint64 }
func genFn0(p0 int32, p1 uint16, p2 bool) bool {
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 2, F1: 2147483648, F2: 0}
        switch "a" {
        case "hello":
            switch (((("foo bar" + "") + (("foo bar" + "hello") + "")) + (("" + "\n\t") + "a")) + "a") {
            case "foo bar":
                switch fi5 {
                case 127:
                    switch "hello" {
                    case "foo bar":
                        v7 := []int{}
                        _ = v7
                    case "\n\t":
                        v8 := GenS1{F0: uint32(p1), F1: 112, F2: 2}
                        _ = v8
                    case "":
                        if (856 != 2) {
                            for fi9 := 0; fi9 < 14; fi9++ {
                                var v10 uint64 = 18446744073709551615
                                v11 := []int{(-(fi9)), 127}
                                p1 = p1
                                p1 = (p1 - 42)
                                _ = fi9
                                _ = v10
                                _ = v11
                            }
                            if p2 {
                                p1 = p1
                                var v12 uint8 = 255
                                _ = v12
                            }
                        }
                    default:
                        p2 = (false || ((float32(((-2.25 / float64(-2.25)) / 1.5)) + (-2.25 / 1.0)) <= -2.25))
                    }
                default:
                    for fi13 := 0; fi13 < 12; fi13++ {
                        var v14 float64 = 3.14159
                        v15 := GenS2{F0: -1, F1: 123}
                        v16 := []int{fi5, (^(fi13))}
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                }
            case "hello":
                v17 := []int{(-(fi5)), (fi5 >> ((4294967295) & 63))}
                _ = v17
            }
        }
        _ = fi5
        _ = v6
    }
    var v18 int8 = 100
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return p2
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 2147483647
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    var v19 uint32 = 336
    v19 = (^(v19))
    v20 := []int{4294967295, (^(v2)), (v2 - int(v19)), (v2 << ((0) & 63))}
    v1[("\n\t" + ("" + "hello"))] = "a"
    if (32767 < 32767) {
        v21 := GenS3{F0: v2}
        if true {
            v20 = append(v20, v2)
            _ = len(v20)
            v22 := GenS4{F0: 255, F1: -742}
            _ = v22
        }
        switch v2 {
        case 128:
            v23 := GenS1{F0: 2, F1: 4294967295, F2: 0}
            _ = v23
        case 32767:
            var v24 float64 = 3.14159
            _ = v24
        case 0:
            var v25 float64 = -2.25
            _ = v25
        default:
            v19 -= v19
        }
        _ = v21
    } else {
        v2 = v2
        if len(v20) > 0 {
            v20[0] = v2
        }
        v26 := GenS5{F0: "hello", F1: -796}
        _ = v26
    }
    switch ("hello" + "a") {
    case "\n\t":
        if len(v20) > 0 {
            v20[0] = 42
        }
    default:
        v2 -= -9223372036854775808
    }
    v2 -= v2
    var v27 uint64 = 42
    v28 := GenS1{F0: (v19 >> ((2) & 63)), F1: (^(v19)), F2: 42}
    for fi29 := 0; fi29 < 5; fi29++ {
        v27 = (v27 * v27)
        if (9223372036854775807 > 9223372036854775807) {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            fmt.Println(v2)
        } else {
            switch ("" + "a") {
            case "a":
                v19 -= (v19 - 4294967295)
            default:
                fmt.Println(v27, fi29, v2)
            }
            v30 := GenS6{F0: v19}
            for fi31 := 0; fi31 < 8; fi31++ {
                switch "" {
                case "foo bar":
                    fmt.Println(v27, v19, fi31)
                case "a":
                    switch (v2 - fi29) {
                    case 4294967295:
                        if false {
                            if false {
                                switch "hello" {
                                case "hello":
                                    var v32 uint8 = 2
                                    _ = v32
                                case "a":
                                    fmt.Println(fi29, v27)
                                case "":
                                    if true {
                                        v27 = uint64(v19)
                                        v19 += 749
                                    }
                                }
                                var v33 int8 = 1
                                _ = v33
                            } else {
                                if true {
                                    v34 := GenS7{F0: 1}
                                    _ = v34
                                }
                            }
                            if true {
                                if (!false) {
                                    var v35 int32 = 2147483647
                                    if (true && false) {
                                        if len(v20) > 0 {
                                            v20[0] = 42
                                        }
                                        fmt.Println(v27, v2, fi31)
                                    } else {
                                        fmt.Println(v19, v27, v35)
                                        v35 -= (v35 >> ((255) & 63))
                                        switch (v2 ^ 42) {
                                        case 128:
                                            v36 := GenS3{F0: 70}
                                            _ = v36
                                        case 9223372036854775807:
                                            v19 -= (v19 << ((uint(v27)) & 63))
                                        case 442:
                                            var v37 uint8 = 0
                                            _ = v37
                                        }
                                    }
                                    _ = v35
                                } else {
                                    v38 := GenS5{F0: "foo bar", F1: -84}
                                    _ = v38
                                }
                                switch (fi29 << ((879) & 63)) {
                                case 586:
                                    v39 := GenS8{F0: ("" + "foo bar"), F1: 0}
                                    _ = v39
                                default:
                                    if true {
                                        v40 := GenS9{F0: -9223372036854775808, F1: true, F2: -251}
                                        var v41 uint16 = 42
                                        var v42 uint16 = (v41 - (v41 + (v41 / ((v41) | 1))))
                                        switch v2 {
                                        case 74:
                                            switch "\n\t" {
                                            case "a":
                                                v41 += v41
                                            case "foo bar":
                                                _ = len(v3)
                                            default:
                                                v2 = 32768
                                            }
                                        }
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                }
                                fmt.Println(fi29, fi29)
                                v27 += (v27 / ((356) | 1))
                            } else {
                                v19 -= (v19 / ((v19) | 1))
                                switch "a" {
                                case "":
                                    fmt.Println(v19, v2)
                                default:
                                    switch fi29 {
                                    case 572:
                                        var v43 int64 = 423
                                        _ = v43
                                    case 2147483647:
                                        fmt.Println(fi29, fi29)
                                        switch "" {
                                        case "hello":
                                            if (!false) {
                                                v19 += (v19 * v19)
                                                if true {
                                                    var v44 int32 = 0
                                                    v45 := v4[fi31]
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    var v46 bool = true
                                                    var v47 string = ""
                                                    fmt.Println(v47)
                                                    _ = v46
                                                    _ = v47
                                                }
                                            } else {
                                                fmt.Println(v27, fi31, v2)
                                                v48 := GenS7{F0: 537}
                                                if (((0.1 / 0.1) / 1.5) == (0.0 + (1e10 - 1.5))) {
                                                    if len(v20) > 0 {
                                                        v20[0] = fi29
                                                    }
                                                    var v49 uint8 = 127
                                                    var v50 int32 = 2147483647
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    v19 -= (v19 - v19)
                                                    v51 := GenS5{F0: "hello", F1: 42}
                                                    fmt.Println(v2, v19)
                                                    _ = v51
                                                }
                                                _ = v48
                                            }
                                        case "\n\t":
                                            v52 := GenS5{F0: "hello", F1: 42}
                                            _ = v52
                                        case "":
                                            var v53 uint8 = 128
                                            _ = v53
                                        }
                                        var v54 uint32 = (v19 - (v19 + (v19 / (((v19 ^ (v19 + v19))) | 1))))
                                        if (2147483647 == 977) {
                                            v2 -= (fi29 * int(v54))
                                        } else {
                                            v3[(^(v19))] = true
                                            v55 := GenS5{F0: "", F1: 2147483647}
                                            _ = v55
                                        }
                                        if (!false) {
                                            v2 -= (v2 % (((v2 % ((0) | 1))) | 1))
                                            v2 += fi29
                                        }
                                        v19 -= 65535
                                        var v56 int = fi29
                                        _ = v54
                                        _ = v56
                                    case -26:
                                        v27 += 18446744073709551615
                                    default:
                                        if (-1 == 42) {
                                            var v57 int8 = 127
                                            if false {
                                                var v58 int32 = 1
                                                fmt.Println(v57, fi29, v58)
                                                v57 = v57
                                                _ = v58
                                            }
                                            switch ("a" + "a") {
                                            case "hello":
                                                v19 = v19
                                            case "\n\t":
                                                var v59 uint8 = 240
                                                _ = v59
                                            }
                                            _ = v57
                                        }
                                        v3[v19] = false
                                        if (!false) {
                                            v60 := GenS2{F0: 377, F1: 65}
                                            var v61 int8 = 127
                                            if false {
                                                fmt.Println(v61, v61, v61)
                                                fmt.Println(v2, v61, v61)
                                                v62 := GenS10{F0: (false && ((v19 << ((255) & 63)) == v19)), F1: 2147483647, F2: ""}
                                                _ = v62
                                            } else {
                                                v2 += 85
                                                v63 := GenS11{F0: "foo bar", F1: (v27 < (v27 & (v27 / (((v27 | v27)) | 1))))}
                                                v61 -= v61
                                                var v64 float32 = (0.1 / (3.14159 - 0.1))
                                                _ = v63
                                                _ = v64
                                            }
                                            _ = v60
                                            _ = v61
                                        } else {
                                            v27 += (v27 / ((v27) | 1))
                                        }
                                        v65 := GenS2{F0: -368, F1: 127}
                                        if false {
                                            if (false && (!true)) {
                                                var v66 int16 = 1000
                                                _ = v66
                                            } else {
                                                var v67 bool = false
                                                _ = v67
                                            }
                                            v68 := GenS5{F0: "", F1: 19}
                                            _ = v68
                                        }
                                        _ = v65
                                    }
                                }
                            }
                        }
                    case -1:
                        switch fi31 {
                        case -9223372036854775808:
                            delete(v4, fi31)
                            var v69 int64 = 448
                            _ = v69
                        default:
                            if false {
                                v70 := v4[fi31]
                                v2 += v2
                                _ = v70
                            } else {
                                var v71 string = (("" + ("hello" + "")) + "a")
                                var v72 int16 = 1
                                v2 = 127
                                _ = v71
                                _ = v72
                            }
                        }
                    default:
                        v27 += v27
                    }
                }
                v19 = v19
                v73 := GenS12{F0: -1, F1: (^(v27))}
                _ = fi31
                _ = v73
            }
            _ = v30
        }
        _ = fi29
    }
    fmt.Println(v19, v2, v2)
    fmt.Println(v2, v27, v27)
    v19 = (v19 * 602)
    v74 := GenS7{F0: 2}
    if true {
        for fi75 := 0; fi75 < 14; fi75++ {
            var v76 int16 = -1
            v19 -= v19
            if (!false) {
                v77 := GenS3{F0: v2}
                _ = v77
            } else {
                var v78 float64 = ((0.0 * (0.1 - 0.0)) * -2.25)
                v79 := GenS1{F0: (v19 + 42), F1: v19, F2: (^(v27))}
                _ = v78
                _ = v79
            }
            _ = fi75
            _ = v76
        }
        switch (^(v2)) {
        case 127:
            if (151 == 845) {
                v27 -= 42
                var v80 uint16 = 1
                var v81 int16 = 32767
                _ = v80
                _ = v81
            }
        case -2147483648:
            switch "a" {
            case "\n\t":
                var v82 uint16 = 0
                _ = v82
            }
        case 0:
            v2 = (v2 - (v2 >> ((525) & 63)))
        default:
            switch v2 {
            case 1:
                switch "hello" {
                case "hello":
                    v2 -= (v2 - 127)
                default:
                    v83 := GenS3{F0: v2}
                    _ = v83
                }
            case 4294967295:
                v2 += -815
            }
        }
        var v84 int64 = -321
        _ = v84
    }
    v27 -= v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v74
}
