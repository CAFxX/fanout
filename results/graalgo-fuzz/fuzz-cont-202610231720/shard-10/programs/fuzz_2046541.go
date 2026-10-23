package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 int; F2 uint32 }
type GenS8 struct { F0 uint; F1 int32; F2 bool }
func genFn0(p0 int64) uint64 {
    var v7 uint = 18446744073709551615
    var v8 string = "a"
    var v9 uint64 = 4294967295
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() string {
    var v10 int64 = 912
    v11 := []int{-1}
    var v12 uint32 = 0
    _ = v10
    _ = v11
    _ = v12
    return ""
}

func main() {
    var v1 int32 = 2147483647
    var v2 int8 = -128
    v3 := make(map[uint32]string)
    var v4 float64 = (3.14159 + 0.1)
    var v5 int32 = -683
    v6 := make(map[int]string)
    if true {
        for fi13 := 0; fi13 < 3; fi13++ {
            v2 -= v2
            v14 := GenS1{F0: (v1 % ((v1) | 1)), F1: 307}
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            if (65535 > 18446744073709551615) {
                var v16 float32 = 0.0
                delete(v3, 42)
                v4 = v4
                _ = v16
            }
            _ = fi15
        }
    } else {
        fmt.Println(v5)
        if ((^(v2)) != (^(v2))) {
            fmt.Println(v2)
            var v17 float64 = v4
            var v18 int8 = (v2 * (v2 >> ((3) & 63)))
            _ = v17
            _ = v18
        }
        v4 = v4
    }
    if ((((float32(float64(0.1)) - 1e10) / (0.1 - (0.1 + 1.5))) != (-2.25 + 1e10)) && true) {
        v2 -= 0
        for fi19 := 0; fi19 < 15; fi19++ {
            v2 = int8(fi19)
            v5 = (v5 << ((1) & 63))
            v1 = v5
            _ = fi19
        }
        if false {
            v1 = v1
            v20 := GenS1{F0: (v5 >> ((0) & 63)), F1: 784}
            _ = v20
        } else {
            v21 := GenS1{F0: (^(v1)), F1: 995}
            if (!true) {
                v2 = 127
                var v22 uint = 1
                _ = v22
            }
            _ = v21
        }
    } else {
        v1 -= v5
        v5 = v1
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        if false {
            delete(v3, 70)
            switch "" {
            case "":
                fmt.Println(v4, fi23)
            }
            var v24 uint = 18446744073709551615
            _ = v24
        } else {
            v5 += (v1 & 763)
        }
        _ = fi23
    }
    if true {
        v25 := []int{2147483647}
        _ = v25
    } else {
        var v26 string = ""
        _ = v26
    }
    v27 := []int{2147483647, 32768, 4294967295}
    switch -577 {
    case 443:
        v1 = v1
    case 9223372036854775807:
        switch -2147483648 {
        case 0:
            if false {
                v5 = v5
                v3[727] = ""
            } else {
                v2 += (v2 << ((18446744073709551615) & 63))
                v1 = 1
                if true {
                    v28 := GenS1{F0: 1, F1: 0}
                    v29 := v3[65535]
                    _ = v28
                    _ = v29
                }
            }
        case -2147483648:
            for fi30 := 0; fi30 < 9; fi30++ {
                for fi31 := 0; fi31 < 8; fi31++ {
                    if false {
                        v5 = v5
                    } else {
                        v2 += v2
                        v2 = v2
                        switch "a" {
                        case "foo bar":
                            if ((65535 <= 2) && (42 > 2147483647)) {
                                v32 := GenS1{F0: (v1 >> ((255) & 63)), F1: 0}
                                switch fi30 {
                                case 767:
                                    v5 -= (v5 | (v5 - 985))
                                case -389:
                                    var v33 float32 = 3.14159
                                    fmt.Println(v5, v5, v33)
                                    _ = v33
                                default:
                                    fmt.Println(fi30, v5)
                                }
                                switch "foo bar" {
                                case "a":
                                    var v34 uint8 = 230
                                    _ = v34
                                case "\n\t":
                                    if (false || false) {
                                        v35 := GenS1{F0: v5, F1: 0}
                                        var v36 uint64 = 42
                                        _ = v35
                                        _ = v36
                                    } else {
                                        switch ("" + "hello") {
                                        case "":
                                            v2 -= v2
                                        case "a":
                                            _ = len(v27)
                                        case "foo bar":
                                            v37 := GenS2{F0: 42, F1: fi31}
                                            _ = v37
                                        }
                                    }
                                case "":
                                    var v38 uint16 = 65535
                                    _ = v38
                                }
                                v39 := GenS3{F0: 76, F1: 65535}
                                var v40 uint32 = 65535
                                _ = v32
                                _ = v39
                                _ = v40
                            }
                            v2 -= v2
                        }
                    }
                    fmt.Println(fi31, v5, fi30)
                    _ = len(v27)
                    _ = fi31
                }
                var v41 float32 = (0.0 * float32(0.0))
                _ = fi30
                _ = v41
            }
        case -271:
            if ((v4 - ((v4 * (float64((0.0 + (float32(1e10) - 1e10))) / ((((v4 + v4) * 1.5) - (v4 / (v4 * (1e10 - float64(0.1))))) / v4))) - (((3.14159 * v4) * v4) - (1e10 * ((v4 / (v4 + (v4 / v4))) + (1e10 / float64(3.14159))))))) > v4) {
                _ = len(v27)
                v1 -= v5
            } else {
                for fi42 := 0; fi42 < 6; fi42++ {
                    fmt.Println(v4, fi42, v2)
                    v43 := GenS2{F0: 322, F1: 0}
                    for fi44 := 0; fi44 < 13; fi44++ {
                        if (true || (!true)) {
                            var v45 int8 = -1
                            v46 := GenS4{F0: fi42, F1: (!true)}
                            _ = v45
                            _ = v46
                        }
                        var v47 int16 = 1
                        if (1 == 824) {
                            switch "a" {
                            case "a":
                                v47 += (v47 * (v47 * v47))
                            case "\n\t":
                                var v48 float64 = (3.14159 + (v4 / v4))
                                v49 := GenS1{F0: -2147483648, F1: 509}
                                _ = v48
                                _ = v49
                            default:
                                if (0 != 18446744073709551615) {
                                    v50 := GenS1{F0: (v5 & int32(v47)), F1: genFn0(415)}
                                    _ = v50
                                }
                            }
                            if (127 <= 0) {
                                if false {
                                    fmt.Println(v5)
                                }
                                v27 = append(v27, fi44)
                                fmt.Println(v47)
                            } else {
                                var v51 string = "hello"
                                var v52 int16 = v47
                                _ = v51
                                _ = v52
                            }
                        } else {
                            if true {
                                v2 = (v2 - v2)
                            } else {
                                var v53 uint8 = 202
                                v3[4294967295] = "a"
                                v2 -= v2
                                _ = v53
                            }
                            switch "\n\t" {
                            case "\n\t":
                                _ = len(v27)
                            }
                            fmt.Println(v4)
                        }
                        _ = fi44
                        _ = v47
                    }
                    _ = fi42
                    _ = v43
                }
                v54 := GenS5{F0: "hello", F1: 4294967295}
                if (42 < genFn0(988)) {
                    var v55 int64 = -1
                    var v56 bool = false
                    _ = v55
                    _ = v56
                } else {
                    v2 -= (v2 & (v2 - v2))
                    var v57 int16 = 1
                    _ = v57
                }
                _ = v54
            }
        }
    case 42:
        fmt.Println(v1, v2)
    }
    var v58 float64 = -2.25
    v59 := GenS6{F0: 1}
    v6[0] = "\n\t"
    var v60 string = ("foo bar" + "foo bar")
    for fi61 := 0; fi61 < 14; fi61++ {
        switch v60 {
        case "a":
            for fi62 := 0; fi62 < 16; fi62++ {
                var v63 int8 = (-(v2))
                _ = fi62
                _ = v63
            }
        case "\n\t":
            v64 := GenS6{F0: 42}
            _ = v64
        }
        _ = fi61
    }
    for fi65 := 0; fi65 < 4; fi65++ {
        _ = len(v27)
        v66 := GenS7{F0: 358, F1: fi65, F2: 358}
        var v67 uint = 1
        _ = fi65
        _ = v66
        _ = v67
    }
    var v68 float32 = float32((0.1 / v58))
    v69 := GenS8{F0: 1, F1: 42, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v68
    _ = v69
}
