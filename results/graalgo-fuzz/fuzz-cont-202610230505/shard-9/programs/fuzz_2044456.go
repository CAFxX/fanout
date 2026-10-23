package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint64; F1 int; F2 uint64 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int8; F1 int64; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 uint64 }
func genFn0() uint16 {
    var v5 int32 = 696
    switch "\n\t" {
    case "":
        var v6 int8 = -128
        _ = v6
    case "hello":
        fmt.Println(v5)
    }
    switch 32768 {
    case -987:
        v7 := GenS1{F0: -9223372036854775808}
        _ = v7
    case -1:
        fmt.Println(v5)
    case -958:
        v5 -= -2147483648
    default:
        for fi8 := 0; fi8 < 14; fi8++ {
            if ((19 < 65535) || true) {
                var v9 uint32 = 420
                for fi10 := 0; fi10 < 17; fi10++ {
                    v5 = v5
                    v11 := GenS1{F0: 0}
                    var v12 uint8 = 181
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := []int{}
                    _ = fi13
                    _ = v14
                }
                _ = v9
            }
            _ = fi8
        }
    }
    if (-1 < 127) {
        switch -2147483648 {
        case 32767:
            for fi15 := 0; fi15 < 20; fi15++ {
                fmt.Println(fi15, v5)
                _ = fi15
            }
        case 4294967295:
            v5 -= v5
        case 0:
            v5 = -1
        default:
            v5 = (v5 & (v5 ^ v5))
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := []int{(^(fi16)), 42, fi16}
            v18 := GenS1{F0: 9223372036854775807}
            var v19 float32 = (0.0 - ((((0.1 * 1e10) + (1.5 / (-2.25 - 0.1))) * 1e10) + ((1.5 * 0.1) / 3.14159)))
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
    } else {
        var v20 uint64 = 4294967295
        fmt.Println(v20, v20)
        _ = v20
    }
    v21 := []int{1, 396, 0, -843}
    _ = v21
    _ = v5
    return 2
}
func genFn1(p0 uint, p1 int64) int8 {
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := GenS1{F0: p1}
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        if true {
            if (42 != 42) {
                for fi25 := 0; fi25 < 3; fi25++ {
                    if (-2147483648 == 0) {
                        p0 = p0
                        var v26 bool = (true && false)
                        var v27 float64 = 0.0
                        _ = v26
                        _ = v27
                    } else {
                        switch ("a" + ("a" + "\n\t")) {
                        case "hello":
                            p0 -= (p0 / ((0) | 1))
                        }
                        v28 := []int{}
                        var v29 float64 = 3.14159
                        var v30 string = (("\n\t" + "\n\t") + "hello")
                        fmt.Println(p0)
                        switch v30 {
                        case "hello":
                            var v31 float32 = -2.25
                            _ = v31
                        }
                        var v32 uint16 = 2
                        _ = v28
                        _ = v29
                        _ = v30
                        _ = v32
                    }
                    _ = fi25
                }
                var v33 int = (fi24 | (fi24 * fi24))
                v33 += 4294967295
                _ = v33
            }
            fmt.Println(p0, p1, p0)
        }
        v34 := GenS2{F0: (p1 % (((p1 * p1)) | 1)), F1: false}
        switch ("\n\t" + "foo bar") {
        case "":
            v35 := GenS1{F0: p1}
            _ = v35
        case "foo bar":
            if ((^(p0)) <= (p0 * (p0 / ((p0) | 1)))) {
                if (744 >= 0) {
                    p0 = 841
                    var v36 uint = 65535
                    var v37 int16 = 0
                    _ = v36
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 18; fi38++ {
                        switch fi38 {
                        case 458:
                            var v39 int = -820
                            switch 2147483647 {
                            case 32768:
                                fmt.Println(p0, fi24, fi24)
                            }
                            v40 := []int{fi38, -9223372036854775808}
                            v39 = v39
                            _ = v39
                            _ = v40
                        case 239:
                            fmt.Println(p0, p0, p0)
                        case 32768:
                            p1 = 848
                        default:
                            v41 := GenS1{F0: (^(p1))}
                            _ = v41
                        }
                        _ = fi38
                    }
                    v42 := GenS3{F0: 164, F1: fi24, F2: 1}
                    _ = v42
                }
            }
        }
        _ = fi24
        _ = v34
    }
    p1 += p1
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := []uint32{}
    var v2 uint = 717
    v3 := []int{4294967295, int(v2), -2147483648, 4294967295}
    v4 := make(map[int]int)
    _ = len(v4)
    if true {
        switch 506 {
        case 2147483647:
            var v43 uint8 = 1
            _ = v43
        case 32768:
            var v44 int64 = 222
            _ = v44
        case 0:
            var v45 uint64 = 4294967295
            _ = v45
        default:
            v46 := GenS4{F0: 1, F1: 9223372036854775807}
            _ = v46
        }
        var v47 float32 = 1.5
        switch "\n\t" {
        case "hello":
            for fi48 := 0; fi48 < 7; fi48++ {
                _ = len(v4)
                if len(v3) > 0 {
                    v3[0] = fi48
                }
                var v49 uint16 = 32767
                _ = fi48
                _ = v49
            }
        }
        _ = v47
    }
    switch 2147483647 {
    case -1:
        fmt.Println(v2)
    }
    fmt.Println(v2)
    if (209 <= 1) {
        if true {
            switch "hello" {
            case "\n\t":
                v2 = (v2 | v2)
            default:
                var v50 int64 = 246
                _ = v50
            }
        } else {
            fmt.Println(v2)
        }
        if false {
            for fi51 := 0; fi51 < 14; fi51++ {
                _ = len(v4)
                _ = fi51
            }
            fmt.Println(v2)
            var v52 string = "\n\t"
            _ = v52
        } else {
            v2 -= v2
            for fi53 := 0; fi53 < 7; fi53++ {
                var v54 int16 = -308
                switch fi53 {
                case 32768:
                    switch (^(fi53)) {
                    case -1:
                        var v55 uint16 = genFn0()
                        _ = v55
                    case 32767:
                        v54 = (v54 & (v54 - -32768))
                    case 442:
                        for fi56 := 0; fi56 < 19; fi56++ {
                            v54 = 820
                            v54 -= v54
                            _ = fi56
                        }
                    default:
                        v57 := GenS5{F0: 1, F1: 1, F2: (fi53 & (-(fi53)))}
                        _ = v57
                    }
                case -905:
                    v54 = v54
                }
                _ = fi53
                _ = v54
            }
            for fi58 := 0; fi58 < 7; fi58++ {
                var v59 float32 = (-2.25 * -2.25)
                switch 32767 {
                case 9223372036854775807:
                    var v60 int32 = -2147483648
                    _ = v60
                }
                var v61 float64 = (((-2.25 + (((1e10 + 3.14159) * (((3.14159 + 0.0) + -2.25) / 3.14159)) + 1.5)) - -2.25) / (float64(1.5) - 0.1))
                _ = fi58
                _ = v59
                _ = v61
            }
        }
    } else {
        if true {
            for fi62 := 0; fi62 < 17; fi62++ {
                v63 := GenS1{F0: -1}
                _ = fi62
                _ = v63
            }
            var v64 bool = (1000 != -629)
            var v65 int32 = 2147483647
            _ = v64
            _ = v65
        }
    }
    for fi66 := 0; fi66 < 19; fi66++ {
        v67 := GenS5{F0: genFn1((v2 * v2), 1), F1: -1, F2: fi66}
        _ = fi66
        _ = v67
    }
    if (v2 != v2) {
        delete(v4, 4294967295)
        var v68 int32 = -1
        v69 := GenS2{F0: 231, F1: (!(!(-473 > 164)))}
        _ = v68
        _ = v69
    } else {
        v2 += 0
    }
    v2 = (v2 | 1)
    if (true && true) {
        v70 := GenS6{F0: -847}
        _ = v70
    }
    v2 -= v2
    v2 = v2
    if true {
        if true {
            delete(v4, 2147483647)
        }
        if (255 > (v2 >> (((v2 | 0)) & 63))) {
            delete(v4, 2147483647)
        } else {
            var v71 uint16 = 737
            _ = v71
        }
    } else {
        if false {
            var v72 float64 = 0.1
            v2 = 255
            v2 -= v2
            _ = v72
        }
        if false {
            v2 = v2
        } else {
            v73 := GenS3{F0: 739, F1: 32768, F2: 0}
            if (-9223372036854775808 == 215) {
                var v74 float64 = -2.25
                var v75 int = 1
                _ = v74
                _ = v75
            } else {
                var v76 string = ("" + ("" + ("foo bar" + "")))
                var v77 int8 = 0
                var v78 uint = (v2 >> ((2) & 63))
                _ = v76
                _ = v77
                _ = v78
            }
            if false {
                for fi79 := 0; fi79 < 8; fi79++ {
                    if (42 == 32767) {
                        v80 := GenS7{F0: 255}
                        for fi81 := 0; fi81 < 13; fi81++ {
                            if ((int16(fi81) <= 1000) && false) {
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v2)
                                default:
                                    fmt.Println(v2)
                                }
                            } else {
                                v82 := GenS8{F0: (0 == genFn1((v2 % (((v2 / ((v2) | 1))) | 1)), 42)), F1: 42}
                                _ = v82
                            }
                            var v83 int16 = 0
                            _ = fi81
                            _ = v83
                        }
                        var v84 bool = ((false && true) && (468 == 406))
                        _ = v80
                        _ = v84
                    }
                    var v85 float64 = 1.5
                    v85 = 1.5
                    _ = fi79
                    _ = v85
                }
                var v86 int16 = 0
                _ = v86
            } else {
                _ = len(v4)
            }
            _ = v73
        }
        fmt.Println(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
