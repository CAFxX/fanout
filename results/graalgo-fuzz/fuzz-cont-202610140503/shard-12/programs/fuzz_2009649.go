package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int16; F1 string; F2 bool }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int32; F1 bool; F2 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 int64; F1 int16; F2 uint16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint32; F1 int }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 int8; F1 int8; F2 uint32 }
type GenS15 struct { F0 uint16; F1 bool; F2 int64 }
func genFn0(p0 int, p1 uint8) int8 {
    var v7 int16 = 42
    v7 = v7
    v8 := []int{(p0 & p0), (p0 >> ((4294967295) & 63)), int(v7)}
    v8 = append(v8, 32767)
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v7, p0)
        if false {
            v10 := GenS1{F0: "hello", F1: (("a" + (((("a" + "hello") + "\n\t") + "\n\t") + "\n\t")) + "\n\t"), F2: 0}
            var v11 bool = (84 != 109)
            _ = v10
            _ = v11
        } else {
            v12 := GenS1{F0: "a", F1: "\n\t", F2: 42}
            _ = v12
        }
        if (!false) {
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 int8 = 0
                p0 -= fi9
                v7 += (v7 ^ v7)
                _ = fi13
                _ = v14
            }
        } else {
            var v15 float32 = (0.0 - ((-2.25 - -2.25) + ((0.0 * 1.5) / float32((0.1 / 1e10)))))
            var v16 int32 = 731
            var v17 uint64 = 42
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -128
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []uint32{4294967295}
    v3 := make(map[uint32]int)
    var v4 int = 180
    var v5 string = "\n\t"
    var v6 string = (v5 + "")
    if (((v1 | v1) == v1) && ((!false) && ("a" < v6))) {
        v5 = (v5 + ((((v6 + "a") + (v6 + ("foo bar" + (v5 + v6)))) + (v5 + (v5 + "hello"))) + (v5 + v5)))
    } else {
        var v18 string = (v6 + v5)
        if true {
            _ = len(v3)
            v1 -= 466
        }
        _ = v18
    }
    var v19 bool = (!false)
    var v20 uint8 = 155
    fmt.Println(v19, v20)
    if false {
        if false {
            var v21 uint16 = 985
            if ((v19 && (!true)) && false) {
                if ((3.14159 + (((0.0 / (1.5 / float64(3.14159))) - 1.5) / 1.0)) == -2.25) {
                    var v22 string = v6
                    for fi23 := 0; fi23 < 4; fi23++ {
                        for fi24 := 0; fi24 < 13; fi24++ {
                            delete(v3, v1)
                            var v25 uint32 = v1
                            var v26 uint = 255
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 1; fi27++ {
                            var v28 string = "a"
                            var v29 int8 = int8(v20)
                            if (v21 != 0) {
                                v22 = "hello"
                                if v19 {
                                    v3[2] = fi27
                                    switch "a" {
                                    case "foo bar":
                                        v20 = v20
                                    default:
                                        var v30 uint = 18446744073709551615
                                        _ = v30
                                    }
                                }
                                v6 = v6
                            } else {
                                fmt.Println(v4, v21, v1)
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        _ = fi23
                    }
                    v4 = v4
                    _ = v22
                } else {
                    switch "" {
                    case "hello":
                        fmt.Println(v1)
                    case "":
                        if v19 {
                            v3[v1] = int(v21)
                            fmt.Println(v6, v1)
                        }
                    default:
                        fmt.Println(v5, v20)
                    }
                    var v31 uint = 200
                    _ = v31
                }
            }
            _ = v21
        } else {
            if v19 {
                fmt.Println(v19, v4)
            }
            switch (v4 * 763) {
            case 32768:
                v32 := GenS1{F0: "foo bar", F1: (((v6 + v6) + "a") + "foo bar"), F2: uint64(v20)}
                _ = v32
            }
        }
        var v33 int64 = 276
        for fi34 := 0; fi34 < 3; fi34++ {
            if len(v2) > 0 {
                v2[0] = 65535
            }
            for fi35 := 0; fi35 < 18; fi35++ {
                switch fi34 {
                case 645:
                    switch (v4 & fi34) {
                    case -1:
                        if v19 {
                            v4 = -706
                        } else {
                            if (genFn0(fi34, (v20 % (((v20 + (v20 & (^(v20))))) | 1))) < 127) {
                                if (!false) {
                                    fmt.Println(v6, v1, v6)
                                    v36 := GenS2{F0: 127, F1: 2, F2: v33}
                                    var v37 int8 = genFn0((-(v4)), (v20 + (v20 - v20)))
                                    _ = v36
                                    _ = v37
                                }
                            } else {
                                v4 = 127
                                var v38 float32 = 1.5
                                v3[v1] = (fi35 << ((18446744073709551615) & 63))
                                _ = v38
                            }
                            v39 := GenS3{F0: 255, F1: (^(v20))}
                            var v40 int32 = -1
                            _ = v39
                            _ = v40
                        }
                    default:
                        switch fi35 {
                        case -209:
                            v19 = (42 <= -1)
                        case 32767:
                            var v41 uint16 = 969
                            _ = v41
                        case 9223372036854775807:
                            v42 := GenS4{F0: 844, F1: -103}
                            _ = v42
                        }
                        v43 := v3[v1]
                        _ = v43
                    }
                case 0:
                    var v44 uint8 = v20
                    _ = v44
                default:
                    v20 += 67
                }
                v45 := GenS5{F0: 1, F1: v5, F2: (fi35 != -1)}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi35
                _ = v45
            }
            _ = fi34
        }
        _ = v33
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        v6 = ("\n\t" + v6)
        _ = fi46
    }
    var v47 int32 = -974
    for fi48 := 0; fi48 < 12; fi48++ {
        if true {
            if false {
                var v49 int32 = v47
                _ = v49
            }
        } else {
            var v50 int64 = -9223372036854775808
            _ = v50
        }
        for fi51 := 0; fi51 < 14; fi51++ {
            if v19 {
                v6 = v6
            } else {
                var v52 uint = 699
                var v53 uint32 = (v1 + (v1 * 2147483647))
                var v54 float32 = 1.5
                _ = v52
                _ = v53
                _ = v54
            }
            var v55 int64 = -584
            _ = fi51
            _ = v55
        }
        _ = fi48
    }
    if (631 != uint(v1)) {
        if true {
            fmt.Println(v6, v20, v47)
            v56 := GenS5{F0: -32768, F1: v5, F2: (v19 || true)}
            if ((v4 / (((v4 & 2147483647)) | 1)) > (^(v4))) {
                switch v6 {
                case "hello":
                    if (!((-32768 < 885) || v19)) {
                        if (0 != 255) {
                            var v57 uint = 14
                            v58 := GenS6{F0: v47, F1: 18446744073709551615}
                            switch v6 {
                            case "":
                                for fi59 := 0; fi59 < 18; fi59++ {
                                    for fi60 := 0; fi60 < 3; fi60++ {
                                        _ = len(v3)
                                        _ = fi60
                                    }
                                    v5 = ("hello" + v6)
                                    _ = fi59
                                }
                            case "hello":
                                if (v5 >= v5) {
                                    v61 := GenS7{F0: (v47 >> (((v57 % ((v57) | 1))) & 63)), F1: (!((!(1e10 > (0.1 - 1.5))) && false)), F2: v1}
                                    _ = v61
                                } else {
                                    var v62 uint = v57
                                    _ = v62
                                }
                            default:
                                v4 -= v4
                            }
                            _ = v57
                            _ = v58
                        }
                    }
                case "foo bar":
                    var v63 uint8 = v20
                    _ = v63
                default:
                    fmt.Println(v20, v20, v20)
                }
                var v64 uint = 429
                _ = len(v2)
                _ = v64
            } else {
                fmt.Println(v6)
                for fi65 := 0; fi65 < 8; fi65++ {
                    v4 += fi65
                    _ = fi65
                }
            }
            _ = v56
        } else {
            v19 = v19
        }
    }
    for fi66 := 0; fi66 < 2; fi66++ {
        var v67 int8 = 0
        var v68 bool = v19
        _ = fi66
        _ = v67
        _ = v68
    }
    for fi69 := 0; fi69 < 17; fi69++ {
        if v19 {
            switch fi69 {
            case 127:
                v2 = append(v2, v1)
            case 4294967295:
                fmt.Println(v5)
            case 32767:
                delete(v3, (^(v1)))
            }
            switch (v4 | (fi69 / ((fi69) | 1))) {
            case 128:
                v2 = append(v2, (^(v1)))
            case 197:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 32768:
                var v70 string = v6
                _ = v70
            }
            switch 50 {
            case 32767:
                if true {
                    switch v4 {
                    case 32768:
                        v71 := GenS8{F0: v4}
                        _ = v71
                    default:
                        var v72 float64 = float64(1e10)
                        _ = v72
                    }
                    if (1.5 <= (0.1 * -2.25)) {
                        for fi73 := 0; fi73 < 14; fi73++ {
                            v1 = v1
                            _ = fi73
                        }
                    }
                    if v19 {
                        v19 = true
                        var v74 float32 = 0.1
                        for fi75 := 0; fi75 < 4; fi75++ {
                            fmt.Println(v74, fi75, v5)
                            v2 = append(v2, uint32(v20))
                            switch v6 {
                            case "\n\t":
                                switch v6 {
                                case "foo bar":
                                    var v76 float64 = float64((v74 + (v74 - (v74 * (1e10 / v74)))))
                                    _ = v76
                                case "\n\t":
                                    var v77 int16 = 1
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    delete(v3, v1)
                                    v47 = v47
                                    _ = v77
                                case "":
                                    _ = len(v2)
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "hello":
                                var v78 uint64 = 42
                                _ = v78
                            case "a":
                                v79 := GenS9{F0: 1, F1: v20}
                                _ = v79
                            }
                            _ = fi75
                        }
                        _ = v74
                    } else {
                        if (506 != 628) {
                            v80 := GenS3{F0: 65535, F1: 42}
                            var v81 int = 42
                            _ = v80
                            _ = v81
                        } else {
                            if ((!true) || false) {
                                v82 := GenS2{F0: 255, F1: 976, F2: 42}
                                _ = v82
                            }
                            var v83 uint16 = 65535
                            _ = v83
                        }
                    }
                } else {
                    fmt.Println(v19)
                }
            case 1:
                switch 9223372036854775807 {
                case 4294967295:
                    v5 = v5
                case -664:
                    v84 := GenS10{F0: -875, F1: -942, F2: 0}
                    _ = v84
                default:
                    var v85 bool = true
                    _ = v85
                }
            case -848:
                v19 = (v19 && v19)
            default:
                for fi86 := 0; fi86 < 13; fi86++ {
                    fmt.Println(v4, v5, v6)
                    _ = len(v3)
                    if true {
                        _ = len(v3)
                        v87 := GenS10{F0: 553, F1: 32767, F2: 65535}
                        v88 := GenS11{F0: -1}
                        _ = v87
                        _ = v88
                    } else {
                        var v89 float64 = (1.5 / (float64(-2.25) * 3.14159))
                        switch v6 {
                        case "foo bar":
                            v90 := GenS12{F0: v1, F1: (^(fi86))}
                            _ = v90
                        case "a":
                            if false {
                                if v19 {
                                    var v91 int32 = 439
                                    v6 = v6
                                    _ = v91
                                }
                                v89 -= float64(0.0)
                                var v92 uint64 = 2
                                _ = v92
                            } else {
                                v89 += (float64((0.0 * 0.1)) * -2.25)
                                var v93 float64 = 0.1
                                _ = v93
                            }
                        case "hello":
                            v6 = (v6 + "a")
                            var v94 int32 = (v47 - (v47 / (((v47 >> ((674) & 63))) | 1)))
                            var v95 string = v6
                            v96 := GenS12{F0: (v1 + (v1 % (((v1 << ((18446744073709551615) & 63))) | 1))), F1: v4}
                            _ = v94
                            _ = v95
                            _ = v96
                        default:
                            _ = len(v3)
                        }
                        switch v4 {
                        case 16:
                            if ((-788 >= 741) || true) {
                                if true {
                                    var v97 int8 = 1
                                    if true {
                                        var v98 float64 = v89
                                        v97 -= v97
                                        v89 -= (v89 / v89)
                                        v6 = v6
                                        _ = v98
                                    }
                                    _ = v97
                                }
                            } else {
                                var v99 bool = (42 <= 32767)
                                var v100 uint8 = (v20 / ((v20) | 1))
                                _ = v100
                                _ = v99
                            }
                        case -69:
                            switch fi86 {
                            case -608:
                                _ = len(v3)
                            default:
                                v4 = (v4 * fi69)
                            }
                        default:
                            v1 = v1
                        }
                        var v101 uint16 = 453
                        if (1 >= 127) {
                            var v102 int = 1
                            _ = v102
                        } else {
                            fmt.Println(v20)
                            if v19 {
                                v1 -= v1
                            } else {
                                fmt.Println(v89)
                            }
                        }
                        switch fi86 {
                        case 170:
                            var v103 int8 = -1
                            v104 := GenS13{F0: (0.0 == 0.1), F1: 719}
                            _ = v103
                            _ = v104
                        case -396:
                            fmt.Println(v4, v89)
                        default:
                            switch v4 {
                            case -29:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 301:
                                if false {
                                    _ = len(v2)
                                    v105 := GenS14{F0: 20, F1: -92, F2: v1}
                                    _ = v105
                                } else {
                                    fmt.Println(fi69, v101)
                                    fmt.Println(v1, fi69)
                                }
                            case 32767:
                                if ((v1 >> ((374) & 63)) <= uint32(v101)) {
                                    fmt.Println(v89, v89)
                                }
                            }
                        }
                        v47 = (^(v47))
                        _ = v101
                        _ = v89
                    }
                    _ = fi86
                }
            }
        }
        if (!(878 == -854)) {
            v106 := GenS15{F0: 42, F1: false, F2: 0}
            _ = v106
        }
        _ = fi69
    }
    var v107 int32 = v47
    v5 = v5
    var v108 int32 = v47
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v107
    _ = v108
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
