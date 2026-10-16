package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 bool; F2 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 int32; F1 uint16; F2 int32 }
type GenS9 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS10 struct { F0 string }
func genFn0(p0 uint8, p1 uint16) uint64 {
    if (0 > -1) {
        v5 := []int{32768, 32768}
        if false {
            v6 := GenS1{F0: -87, F1: 59, F2: 124}
            var v7 uint = 10
            var v8 int8 = 1
            _ = v6
            _ = v7
            _ = v8
        } else {
            fmt.Println(p1, p1)
        }
        for fi9 := 0; fi9 < 3; fi9++ {
            p0 = p0
            p0 = p0
            _ = fi9
        }
        _ = v5
    }
    if true {
        var v10 float64 = 0.1
        _ = v10
    } else {
        p1 = 1
    }
    var v11 string = "a"
    _ = p0
    _ = p1
    _ = v11
    return 42
}
func genFn1(p0 string) uint16 {
    if false {
        v12 := []int{771}
        _ = v12
    } else {
        if ((!(32767 > 86)) && (-2147483648 != -1)) {
            fmt.Println(p0)
            if true {
                if true {
                    v13 := []int{-9223372036854775808, 623, -572, 2147483647}
                    var v14 int8 = 100
                    _ = v13
                    _ = v14
                } else {
                    v15 := []int{}
                    _ = v15
                }
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 uint32 = 2147483648
                    v18 := GenS1{F0: 100, F1: (v17 / ((v17) | 1)), F2: 1}
                    v17 += 65535
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                p0 = "foo bar"
            } else {
                var v19 float64 = 0.1
                _ = v19
            }
        } else {
            v20 := []int{}
            _ = v20
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            fmt.Println(p0, fi21)
            var v22 float64 = (-2.25 + 3.14159)
            fmt.Println(fi21, p0, p0)
            _ = fi21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := []int{463, (-(fi23))}
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 uint16 = 42
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    var v27 uint8 = 0
    v27 = 42
    _ = p0
    _ = v27
    return 32
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -1
    v3 := make(map[string]string)
    var v4 string = "a"
    v2 -= (v2 % (((v2 >> ((4294967295) & 63))) | 1))
    var v28 string = v4
    var v29 uint64 = 4294967295
    switch v28 {
    case "\n\t":
        var v30 bool = false
        _ = v30
    case "":
        var v31 uint16 = 819
        _ = v31
    default:
        var v32 float32 = ((0.0 + 3.14159) * (-2.25 + ((-2.25 / (1e10 * (((0.0 + 1e10) / -2.25) - 0.0))) + -2.25)))
        _ = v32
    }
    fmt.Println(v28, v28)
    if (false && (-659 > 2147483647)) {
        var v33 float64 = (3.14159 / (3.14159 - (0.1 + 1e10)))
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 float64 = 0.0
            var v36 int16 = -564
            if false {
                fmt.Println(v2, fi34, v33)
                v2 = v2
                if (v35 < (0.1 - float64(float32(v33)))) {
                    v33 = 3.14159
                    v35 += (v33 * (v35 / (1e10 - (0.1 * v33))))
                } else {
                    for fi37 := 0; fi37 < 6; fi37++ {
                        fmt.Println(v4, v28, fi34)
                        v38 := []int{}
                        v36 -= 42
                        _ = fi37
                        _ = v38
                    }
                    v39 := GenS1{F0: 0, F1: 0, F2: 1}
                    _ = v39
                }
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v33
    } else {
        for fi40 := 0; fi40 < 10; fi40++ {
            v41 := GenS2{F0: 521, F1: 72}
            var v42 uint = 65535
            _ = fi40
            _ = v41
            _ = v42
        }
        v43 := GenS1{F0: -128, F1: 2147483647, F2: 65535}
        _ = v43
    }
    if true {
        v44 := GenS1{F0: 42, F1: 65535, F2: 1}
        _ = v44
    }
    var v45 int = -899
    var v46 uint32 = 594
    var v47 bool = false
    v28 = ""
    v48 := []int{(v45 - v45), v45}
    if ((v2 >> ((255) & 63)) != 696) {
        v29 += 1
        var v49 int32 = (v2 % ((v2) | 1))
        _ = len(v3)
        _ = v49
    }
    fmt.Println(v29)
    v29 = (v29 >> ((18446744073709551615) & 63))
    for fi50 := 0; fi50 < 18; fi50++ {
        if true {
            var v51 string = "a"
            var v52 int32 = v2
            var v53 int16 = 42
            _ = v51
            _ = v52
            _ = v53
        }
        if (648 < 255) {
            v54 := GenS3{F0: (^(v29))}
            for fi55 := 0; fi55 < 3; fi55++ {
                v56 := GenS3{F0: (^(v29))}
                var v57 float32 = 1.5
                _ = fi55
                _ = v56
                _ = v57
            }
            _ = v54
        }
        if ((v46 << ((uint(v46)) & 63)) > 65535) {
            v58 := GenS4{F0: (100 > 42)}
            for fi59 := 0; fi59 < 4; fi59++ {
                if (true || ("foo bar" >= v28)) {
                    switch -9223372036854775808 {
                    case 32767:
                        v60 := GenS5{F0: v2, F1: (!v47), F2: v2}
                        _ = v60
                    }
                    v4 = (v28 + v4)
                    if v47 {
                        v61 := GenS2{F0: 108, F1: 32767}
                        v4 = "hello"
                        _ = v61
                    } else {
                        fmt.Println(v4, v46)
                    }
                    if v47 {
                        v62 := GenS1{F0: -63, F1: 42, F2: 356}
                        _ = v62
                    } else {
                        var v63 float64 = (1.5 + -2.25)
                        v28 = (v4 + v28)
                        if (((v46 < v46) && v47) || v47) {
                            var v64 uint32 = v46
                            if v47 {
                                fmt.Println(v47)
                                if (((v28 + v28) + v28) <= v4) {
                                    v28 = "foo bar"
                                    v2 -= (v2 & (v2 / ((v2) | 1)))
                                    var v65 uint32 = (v46 * 486)
                                    _ = v65
                                }
                            } else {
                                switch v4 {
                                case "hello":
                                    switch "" {
                                    case "":
                                        v47 = (-1 <= -32768)
                                    }
                                default:
                                    v48 = append(v48, 339)
                                }
                                switch "hello" {
                                case "foo bar":
                                    var v66 string = v4
                                    _ = v66
                                case "a":
                                    switch fi50 {
                                    case 32768:
                                        switch v4 {
                                        case "hello":
                                            v67 := GenS6{F0: ""}
                                            v68 := GenS4{F0: ((false && (128 > 1)) && (false || false))}
                                            _ = v67
                                            _ = v68
                                        case "a":
                                            if (-1 <= 0) {
                                                _ = len(v48)
                                                fmt.Println(v64, v2, v47)
                                            }
                                        }
                                    }
                                }
                            }
                            v28 = v28
                            _ = v64
                        } else {
                            var v69 float32 = (((1.5 - ((float32(float64(1e10)) + float32(3.14159)) * float32(0.0))) + 0.0) * (0.0 / ((3.14159 * (-2.25 / 1e10)) - 1.5)))
                            v45 -= fi59
                            v70 := GenS5{F0: -1, F1: ((v47 && false) && ((3.14159 / (v63 - float64((v69 * 0.0)))) <= -2.25)), F2: (v2 & v2)}
                            _ = v69
                            _ = v70
                        }
                        _ = v63
                    }
                    v71 := GenS7{F0: (v46 >= v46), F1: 103}
                    _ = v71
                } else {
                    v47 = (0.0 < -2.25)
                    var v72 uint64 = (v29 * 1)
                    _ = v72
                }
                var v73 float64 = ((1.5 * 1e10) - 0.1)
                _ = fi59
                _ = v73
            }
            v4 = "foo bar"
            _ = v58
        } else {
            switch -1 {
            case 847:
                if (((1 != -32768) && (0.0 >= 1.5)) && v47) {
                    if true {
                        v48 = append(v48, (^(fi50)))
                        if ((fi50 / ((v45) | 1)) <= (fi50 % (((-(v45))) | 1))) {
                            v74 := GenS8{F0: -2147483648, F1: 550, F2: (-(v2))}
                            v29 -= uint64(v46)
                            _ = v74
                        } else {
                            var v75 int64 = int64(v29)
                            _ = v75
                        }
                    } else {
                        switch (v45 / ((v45) | 1)) {
                        case 324:
                            var v76 int = v45
                            _ = v76
                        case 42:
                            for fi77 := 0; fi77 < 11; fi77++ {
                                var v78 uint8 = 72
                                switch v4 {
                                case "\n\t":
                                    switch (^(fi50)) {
                                    case -896:
                                        v79 := GenS9{F0: -9223372036854775808, F1: v46, F2: v29}
                                        _ = v79
                                    case 4294967295:
                                        v80 := GenS10{F0: v28}
                                        _ = v80
                                    }
                                case "hello":
                                    var v81 string = (v28 + "hello")
                                    _ = v81
                                case "foo bar":
                                    var v82 float64 = (1e10 - 1.5)
                                    _ = v82
                                default:
                                    v46 -= (^(v46))
                                }
                                var v83 uint64 = (v29 << ((466) & 63))
                                v2 -= int32(v29)
                                var v84 float32 = (1e10 / (((1.5 * 1e10) - 3.14159) / (float32(-2.25) + float32(1.5))))
                                _ = fi77
                                _ = v78
                                _ = v83
                                _ = v84
                            }
                        case -55:
                            var v85 int8 = -112
                            _ = v85
                        }
                    }
                    v28 = v28
                }
            case 128:
                var v86 int64 = 1
                _ = v86
            default:
                v29 = v29
            }
            var v87 uint = 161
            _ = v87
        }
        _ = fi50
    }
    if len(v48) > 0 {
        v48[0] = (v45 & -1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
