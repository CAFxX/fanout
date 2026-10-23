package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 bool; F1 int; F2 bool }
type GenS6 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 int64; F2 string }
func genFn0(p0 int16, p1 int64, p2 int64) uint64 {
    var v6 uint = 42
    if false {
        fmt.Println(p1, v6)
        switch 9223372036854775807 {
        case 2147483647:
            v7 := GenS1{F0: 287, F1: 42, F2: (2 == 1)}
            _ = v7
        case -2147483648:
            if (!true) {
                var v8 int16 = p0
                var v9 float32 = 3.14159
                if true {
                    for fi10 := 0; fi10 < 17; fi10++ {
                        v11 := GenS2{F0: -1}
                        v12 := []int{(fi10 & fi10), (-(fi10))}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    v13 := []int{}
                    v14 := GenS3{F0: 65535}
                    _ = v13
                    _ = v14
                } else {
                    var v15 int64 = (p2 - 9223372036854775807)
                    if true {
                        p2 += 9223372036854775807
                        if true {
                            if false {
                                p2 += -1
                                var v16 uint64 = 755
                                for fi17 := 0; fi17 < 2; fi17++ {
                                    var v18 int32 = 2147483647
                                    v19 := GenS2{F0: (v18 * v18)}
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                                _ = v16
                            }
                            v9 -= ((v9 - (v9 - v9)) - v9)
                        } else {
                            var v20 uint16 = 65535
                            _ = v20
                        }
                        var v21 float32 = (1.5 + v9)
                        _ = v21
                    } else {
                        v9 = 3.14159
                    }
                    _ = v15
                }
                _ = v8
                _ = v9
            }
        default:
            for fi22 := 0; fi22 < 15; fi22++ {
                if true {
                    var v23 float32 = (float32(float64(0.0)) + 1.5)
                    v24 := GenS1{F0: 42, F1: 2147483648, F2: true}
                    _ = v23
                    _ = v24
                } else {
                    if true {
                        v25 := []int{fi22, (fi22 ^ (fi22 % ((fi22) | 1)))}
                        _ = v25
                    }
                    p0 = 1000
                    for fi26 := 0; fi26 < 17; fi26++ {
                        if (!(0.1 <= 3.14159)) {
                            var v27 string = ("hello" + "foo bar")
                            switch "\n\t" {
                            case "":
                                v28 := GenS3{F0: 32767}
                                _ = v28
                            case "a":
                                v29 := []int{(fi22 - (fi26 - (fi22 << (((v6 * (v6 | v6))) & 63))))}
                                _ = v29
                            case "foo bar":
                                var v30 float64 = ((1e10 / -2.25) + (3.14159 / 1.5))
                                _ = v30
                            default:
                                if (65535 <= (v6 - (v6 / ((255) | 1)))) {
                                    v31 := []int{}
                                    p1 = p2
                                    _ = len(v31)
                                    _ = v31
                                } else {
                                    v32 := []int{}
                                    _ = len(v32)
                                    if (!(false || false)) {
                                        fmt.Println(p0, fi22)
                                        var v33 float32 = ((1.5 / (1.5 * 3.14159)) / 1.0)
                                        switch (v27 + "foo bar") {
                                        case "\n\t":
                                            fmt.Println(v33, p1)
                                        default:
                                            switch (fi26 - (fi22 >> ((738) & 63))) {
                                            case -9223372036854775808:
                                                v6 = 703
                                                var v34 string = (v27 + "")
                                                _ = v34
                                            case -353:
                                                v6 = v6
                                            default:
                                                var v35 uint16 = 1
                                                _ = v35
                                            }
                                        }
                                        _ = v33
                                    } else {
                                        fmt.Println(p2, fi22)
                                    }
                                    _ = v32
                                }
                            }
                            _ = v27
                        }
                        p1 -= (p1 - (p2 >> ((v6) & 63)))
                        _ = fi26
                    }
                }
                _ = fi22
            }
        }
    } else {
        v36 := []int{1, 0}
        v36 = append(v36, 4294967295)
        var v37 int8 = -128
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        fmt.Println(v6, p0, p0)
        v39 := []int{fi38}
        var v40 uint64 = 777
        _ = fi38
        _ = v39
        _ = v40
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 float64 = float64(float32(0.1))
    var v2 uint = 4294967295
    var v3 uint64 = 18446744073709551615
    v4 := []uint32{}
    v5 := []int{127, 9223372036854775807, int(v3), 2147483647}
    v5 = append(v5, 1)
    switch "" {
    case "foo bar":
        fmt.Println(v3, v1, v3)
    case "a":
        for fi41 := 0; fi41 < 7; fi41++ {
            v42 := GenS4{F0: (-128 != 127), F1: 0}
            _ = fi41
            _ = v42
        }
    case "":
        var v43 float32 = (((1.5 + 1.5) / ((float32(float64((((1.5 + -2.25) * float32(((1e10 - v1) * 0.1))) / 1e10))) - (-2.25 / (float32((((v1 - (float64(0.1) + v1)) / (v1 * float64(3.14159))) - (1.5 + 1e10))) + 3.14159))) / ((((((1.5 - 0.0) * 1e10) - 3.14159) - ((float32(3.14159) * 1.5) / 1.0)) + 0.0) + ((0.1 - 0.0) + 0.0)))) * 0.1)
        _ = v43
    default:
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v44 := GenS4{F0: (true || false), F1: -90}
    if false {
        v2 = (v2 / ((4294967295) | 1))
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    v5 = append(v5, 128)
    if (1 > -1) {
        switch 817 {
        case -2147483648:
            v3 -= (v3 - (^(v3)))
        case 298:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            if false {
                if ((true && false) || false) {
                    _ = len(v4)
                    v1 = ((v1 * v1) - (((float64((0.1 - 0.0)) - (float64((-2.25 / 1.5)) - ((-2.25 / (v1 + ((float64(1e10) - v1) / 1.5))) + (((v1 / v1) - (v1 / (v1 * (0.1 * 1e10)))) * v1)))) * 1.5) + ((v1 + float64(0.0)) / v1)))
                    if (560 < 1) {
                        v2 += v2
                        fmt.Println(v3, v2)
                    }
                }
                for fi45 := 0; fi45 < 19; fi45++ {
                    v46 := GenS5{F0: (96 <= 14), F1: 32768, F2: false}
                    fmt.Println(v3)
                    _ = fi45
                    _ = v46
                }
                var v47 uint32 = 595
                _ = v47
            }
        }
    } else {
        if ((false || (true || false)) || (-1 > int64(v3))) {
            for fi48 := 0; fi48 < 10; fi48++ {
                for fi49 := 0; fi49 < 8; fi49++ {
                    if (fi48 != -1) {
                        fmt.Println(v1, fi48, fi49)
                        v1 = v1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    var v50 int32 = -2147483648
                    v3 += (^(v3))
                    v2 = v2
                    _ = fi49
                    _ = v50
                }
                _ = fi48
            }
            fmt.Println(v3, v2)
        }
        if ((!false) && true) {
            v51 := GenS6{F0: 65535, F1: 78, F2: (v3 + (v3 >> ((42) & 63)))}
            fmt.Println(v2)
            if false {
                v52 := GenS7{F0: -1}
                var v53 uint16 = 905
                _ = v52
                _ = v53
            } else {
                if true {
                    fmt.Println(v2, v2)
                }
                fmt.Println(v1, v1, v2)
            }
            _ = v51
        }
        if (!false) {
            v2 += 691
            if (231 > 1) {
                for fi54 := 0; fi54 < 14; fi54++ {
                    v3 += 42
                    _ = fi54
                }
                _ = len(v4)
                var v55 uint8 = 127
                _ = v55
            } else {
                if false {
                    var v56 uint32 = 4294967295
                    switch -120 {
                    case -1:
                        if (-1 > 630) {
                            for fi57 := 0; fi57 < 11; fi57++ {
                                v58 := GenS7{F0: 1}
                                _ = fi57
                                _ = v58
                            }
                        }
                    default:
                        v56 -= v56
                    }
                    _ = v56
                } else {
                    if (!(-1 <= 128)) {
                        for fi59 := 0; fi59 < 14; fi59++ {
                            var v60 int = (fi59 - fi59)
                            var v61 bool = true
                            switch v60 {
                            case -237:
                                v61 = v61
                            }
                            _ = fi59
                            _ = v60
                            _ = v61
                        }
                        switch 42 {
                        case 425:
                            if ((145 <= 128) && (-128 != 42)) {
                                if false {
                                    fmt.Println(v3, v2)
                                } else {
                                    fmt.Println(v3, v3)
                                    var v62 int = 42
                                    for fi63 := 0; fi63 < 15; fi63++ {
                                        var v64 string = "hello"
                                        for fi65 := 0; fi65 < 13; fi65++ {
                                            v66 := GenS7{F0: -9223372036854775808}
                                            _ = fi65
                                            _ = v66
                                        }
                                        _ = fi63
                                        _ = v64
                                    }
                                    _ = v62
                                }
                                v3 = v3
                            }
                        case 32767:
                            v67 := GenS8{F0: 42, F1: 42, F2: "\n\t"}
                            _ = v67
                        }
                    }
                    var v68 uint8 = 239
                    _ = v68
                }
                switch 32768 {
                case -1:
                    if (true && false) {
                        fmt.Println(v1)
                        v69 := GenS5{F0: true, F1: 510, F2: (!(32767 == 859))}
                        if false {
                            v70 := GenS1{F0: 2, F1: 828, F2: (144 > 728)}
                            var v71 float32 = float32(v1)
                            var v72 uint64 = (^(v3))
                            _ = v70
                            _ = v71
                            _ = v72
                        }
                        _ = v69
                    }
                default:
                    if false {
                        if false {
                            v4 = append(v4, 42)
                            v3 += v3
                        } else {
                            v73 := GenS3{F0: 255}
                            _ = v73
                        }
                        for fi74 := 0; fi74 < 8; fi74++ {
                            v2 = (^(v2))
                            switch "a" {
                            case "\n\t":
                                var v75 int8 = 127
                                _ = v75
                            case "foo bar":
                                var v76 int64 = 0
                                _ = v76
                            }
                            _ = fi74
                        }
                        v2 = (v2 + 4294967295)
                    }
                }
            }
            if true {
                if false {
                    for fi77 := 0; fi77 < 20; fi77++ {
                        for fi78 := 0; fi78 < 11; fi78++ {
                            fmt.Println(fi77)
                            v79 := GenS8{F0: 187, F1: -1, F2: ""}
                            var v80 uint = v2
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                        _ = fi77
                    }
                }
                for fi81 := 0; fi81 < 9; fi81++ {
                    switch fi81 {
                    case 1:
                        v3 = v3
                    case 591:
                        var v82 bool = (true && false)
                        _ = v82
                    }
                    _ = fi81
                }
            } else {
                v3 = v3
            }
        } else {
            var v83 float64 = (float64((1e10 / 3.14159)) - (v1 - 0.1))
            for fi84 := 0; fi84 < 13; fi84++ {
                var v85 int8 = 127
                _ = fi84
                _ = v85
            }
            _ = v83
        }
    }
    v2 += (^(v2))
    v3 -= (v3 / (((v3 >> ((v2) & 63))) | 1))
    _ = len(v4)
    v86 := GenS7{F0: -357}
    for fi87 := 0; fi87 < 4; fi87++ {
        if ("hello" != ("\n\t" + ("" + "\n\t"))) {
            var v88 float64 = v1
            var v89 uint32 = 0
            v3 += v3
            _ = v88
            _ = v89
        } else {
            var v90 uint = 175
            if len(v5) > 0 {
                v5[0] = (fi87 >> (((v2 | v2)) & 63))
            }
            _ = v90
        }
        _ = fi87
    }
    v91 := GenS2{F0: -1}
    for fi92 := 0; fi92 < 12; fi92++ {
        var v93 uint16 = 255
        _ = fi92
        _ = v93
    }
    if true {
        fmt.Println(v3, v3)
        v3 = (v3 + v3)
    }
    var v94 int = 701
    v1 -= 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v86
    _ = v91
    _ = v94
}
