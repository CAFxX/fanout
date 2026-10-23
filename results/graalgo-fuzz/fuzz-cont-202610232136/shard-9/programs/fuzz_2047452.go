package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint; F2 int32 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int16; F1 uint8; F2 string }
func genFn0() uint16 {
    var v4 uint64 = 1
    if (!false) {
        v4 = 1
        if ((false || (973 < 42)) || (127 > 127)) {
            if false {
                v4 += 18446744073709551615
            }
        }
    } else {
        if (true && (true && true)) {
            v4 = (^(v4))
        }
        var v5 int = 4294967295
        _ = v5
    }
    v4 += (v4 & v4)
    v4 = 464
    _ = v4
    return 255
}
func genFn1(p0 uint32, p1 int8, p2 uint8) int64 {
    switch 19 {
    case -1:
        var v6 int8 = (p1 << ((uint(p0)) & 63))
        _ = v6
    case -19:
        v7 := []int{891, int(p1), 843, 0}
        _ = v7
    case 0:
        fmt.Println(p1, p1, p2)
    default:
        var v8 int32 = 42
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: 0}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    var v11 uint8 = 255
    var v12 int16 = -32768
    fmt.Println(v11, v12)
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS2{F0: v3}
        v15 := GenS3{F0: (v12 - 32767), F1: 42, F2: 751}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 float64 = 1.5
    if true {
        var v17 float64 = v16
        var v18 int32 = 42
        var v19 float32 = 0.1
        _ = v17
        _ = v18
        _ = v19
    } else {
        if false {
            if ((true || (v12 <= (^(v12)))) && (74 > 127)) {
                for fi20 := 0; fi20 < 16; fi20++ {
                    v16 = ((1e10 - 0.0) + v16)
                    for fi21 := 0; fi21 < 10; fi21++ {
                        v11 -= 201
                        _ = fi21
                    }
                    _ = fi20
                }
            } else {
                fmt.Println(v12, v3, v3)
                fmt.Println(v3)
                if false {
                    if ((-613 > 42) && false) {
                        var v22 uint32 = 1
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 14; fi23++ {
                            for fi24 := 0; fi24 < 11; fi24++ {
                                switch (fi23 >> ((808) & 63)) {
                                case 782:
                                    fmt.Println(v11, v3, v3)
                                default:
                                    if true {
                                        v16 = float64(3.14159)
                                    }
                                }
                                _ = fi24
                            }
                            _ = fi23
                        }
                    }
                    v25 := GenS4{F0: 0, F1: 1, F2: -1}
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 7; fi26++ {
                switch fi26 {
                case -775:
                    var v27 int = 2147483647
                    _ = v27
                default:
                    if true {
                        var v28 uint64 = 316
                        _ = v28
                    }
                }
                fmt.Println(v16, v12)
                switch v3 {
                case "\n\t":
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 float32 = ((-2.25 / (0.0 * (float32((v16 - (((((3.14159 - v16) * v16) - float64(3.14159)) * (float64(float32(-2.25)) + (((1.5 - (v16 * (v16 - v16))) - v16) / v16))) * 3.14159))) * 3.14159))) + ((1e10 * (3.14159 - (1.5 * float32((v16 * 0.1))))) + float32(float64(-2.25))))
                        if false {
                            var v31 float64 = (v16 / -2.25)
                            v30 += 1e10
                            _ = v31
                        } else {
                            var v32 string = ((v3 + (v3 + v3)) + v3)
                            v32 = v32
                            v33 := GenS5{F0: true, F1: (v11 / ((v11) | 1))}
                            _ = v32
                            _ = v33
                        }
                        v11 = v11
                        _ = fi29
                        _ = v30
                    }
                case "a":
                    var v34 uint64 = 794
                    _ = v34
                case "hello":
                    if false {
                        switch fi26 {
                        case -467:
                            switch v3 {
                            case "hello":
                                v35 := []int{fi26, fi26}
                                _ = v35
                            default:
                                fmt.Println(fi26)
                            }
                        }
                        var v36 int32 = -244
                        fmt.Println(v36)
                        _ = v36
                    } else {
                        if true {
                            switch fi26 {
                            case 0:
                                fmt.Println(v11, v16, v16)
                            }
                            var v37 uint16 = 2
                            for fi38 := 0; fi38 < 7; fi38++ {
                                if (false && ((-(fi38)) != 2147483647)) {
                                    v39 := GenS3{F0: v12, F1: -102, F2: 42}
                                    v11 -= (^(v11))
                                    if ((!false) || (-981 > 42)) {
                                        v16 -= 1.5
                                        v40 := []int{(fi38 - fi38)}
                                        var v41 int8 = -107
                                        _ = v40
                                        _ = v41
                                    }
                                    _ = v39
                                } else {
                                    v37 += (v37 % ((v37) | 1))
                                    v1[(fi38 + (fi26 ^ -1))] = v3
                                    v16 = 1e10
                                    v3 = ("a" + ((v3 + v3) + "\n\t"))
                                }
                                if false {
                                    var v42 int64 = 9223372036854775807
                                    _ = v42
                                } else {
                                    v43 := []int{fi38, (fi38 >> ((4294967295) & 63)), (^(fi26))}
                                    _ = v43
                                }
                                if true {
                                    var v44 float32 = 0.1
                                    fmt.Println(v11, v3, v12)
                                    _ = v44
                                } else {
                                    if (65535 == 907) {
                                        v45 := GenS6{F0: (^(v12)), F1: (v11 | 128), F2: "foo bar"}
                                        var v46 uint32 = 0
                                        _ = v45
                                        _ = v46
                                    }
                                }
                                if true {
                                    v47 := GenS1{F0: 2}
                                    v16 -= v16
                                    _ = v47
                                }
                                fmt.Println(v37, v37, v11)
                                _ = fi38
                            }
                            _ = v37
                        }
                        for fi48 := 0; fi48 < 16; fi48++ {
                            v3 = v3
                            if false {
                                v12 += (v12 | (v12 & v12))
                                v49 := []int{}
                                if (-1 == -1) {
                                    v11 = 128
                                }
                                if (1 > 133) {
                                    _ = len(v49)
                                }
                                _ = len(v49)
                                _ = v49
                            }
                            fmt.Println(fi26, v3)
                            v11 += v11
                            var v50 int16 = v12
                            _ = fi48
                            _ = v50
                        }
                        v51 := []int{fi26, fi26, (fi26 & fi26)}
                        _ = v51
                    }
                default:
                    v11 -= 79
                }
                _ = fi26
            }
            delete(v2, v3)
        } else {
            for fi52 := 0; fi52 < 17; fi52++ {
                switch ("hello" + "a") {
                case "":
                    v53 := GenS3{F0: (v12 >> ((750) & 63)), F1: -99, F2: 185}
                    _ = v53
                case "foo bar":
                    var v54 uint32 = 0
                    _ = v54
                }
                _ = fi52
            }
        }
    }
    if (v11 != v11) {
        v55 := GenS3{F0: 84, F1: -39, F2: 265}
        v3 = ((v3 + v3) + (v3 + ("" + v3)))
        v56 := GenS6{F0: (-(v12)), F1: v11, F2: v3}
        _ = v55
        _ = v56
    }
    if false {
        v57 := GenS3{F0: v12, F1: 0, F2: 441}
        _ = v57
    } else {
        var v58 int = -94
        fmt.Println(v12, v12)
        _ = v58
    }
    v59 := []int{-9223372036854775808, -2147483648, 2147483647, -556}
    v3 = v3
    v60 := v1[370]
    fmt.Println(v16, v16)
    if false {
        var v61 float32 = 0.0
        for fi62 := 0; fi62 < 8; fi62++ {
            for fi63 := 0; fi63 < 1; fi63++ {
                v61 += 3.14159
                _ = fi63
            }
            v12 = v12
            v59 = append(v59, fi62)
            _ = fi62
        }
        _ = v61
    } else {
        v64 := GenS2{F0: (("hello" + (v60 + (v60 + ""))) + v60)}
        _ = v64
    }
    switch v3 {
    case "\n\t":
        fmt.Println(v12)
    case "foo bar":
        v3 = v3
    }
    switch 0 {
    case 42:
        var v65 int = 32767
        _ = v65
    case 4294967295:
        switch -128 {
        case 4294967295:
            v16 = ((v16 - v16) * (3.14159 - (1e10 / v16)))
        default:
            var v66 uint32 = 2
            _ = v66
        }
    case -9223372036854775808:
        var v67 uint32 = 934
        _ = v67
    }
    if (v60 <= "a") {
        if (!true) {
            fmt.Println(v12, v3)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v59
    _ = v60
}
