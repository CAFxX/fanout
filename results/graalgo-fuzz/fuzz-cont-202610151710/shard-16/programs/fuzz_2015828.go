package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 string }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 bool }
type GenS4 struct { F0 int32; F1 int32; F2 uint8 }
type GenS5 struct { F0 bool; F1 int64; F2 uint32 }
type GenS6 struct { F0 int16; F1 int32; F2 int64 }
type GenS7 struct { F0 int16; F1 int; F2 int32 }
func genFn0() uint64 {
    var v5 uint32 = 400
    switch "foo bar" {
    case "":
        v6 := []int{4294967295, -1, -9223372036854775808, -9223372036854775808}
        _ = v6
    case "\n\t":
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := []int{(fi7 & -578), (fi7 & (fi7 >> ((0) & 63))), (fi7 + fi7), -9223372036854775808}
            var v9 int64 = 42
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    _ = v5
    return 4294967295
}
func genFn1(p0 string, p1 uint32) int16 {
    fmt.Println(p0, p0)
    p0 = p0
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v4 int16 = 0
    if true {
        fmt.Println(v1, v1)
    } else {
        v1 += v1
        var v10 uint8 = 127
        v11 := GenS1{F0: -2147483648, F1: 465, F2: "\n\t"}
        _ = v10
        _ = v11
    }
    v1 -= (((v1 / (v1 + float64((3.14159 / (3.14159 + 1e10))))) * v1) / (3.14159 - ((0.1 * (v1 * v1)) / 1e10)))
    for fi12 := 0; fi12 < 19; fi12++ {
        if true {
            v13 := GenS2{F0: 127, F1: -9223372036854775808}
            _ = v13
        } else {
            fmt.Println(fi12, fi12)
            v14 := GenS3{F0: 255, F1: 255, F2: true}
            _ = v14
        }
        _ = fi12
    }
    fmt.Println(v4, v1)
    if (1 != -128) {
        v1 -= 0.0
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v4 = genFn1("", 15)
        if false {
            if ((-2.25 - ((((3.14159 - 1.5) - 0.0) - 1e10) * ((float32((((v1 / 0.1) - 0.0) * (v1 * 0.1))) / 1.0) / -2.25))) < 3.14159) {
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 string = ("" + "a")
                    var v18 uint = 65535
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if false {
                    if true {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            v1 += -2.25
                            _ = fi19
                        }
                        var v20 float32 = (0.0 + float32(float64((0.0 + ((float32((v1 / 1.5)) - 3.14159) - (-2.25 * -2.25))))))
                        _ = v20
                    } else {
                        fmt.Println(v1, v4)
                        for fi21 := 0; fi21 < 19; fi21++ {
                            v4 = (v4 * (-(v4)))
                            _ = fi21
                        }
                        for fi22 := 0; fi22 < 8; fi22++ {
                            v23 := GenS3{F0: 42, F1: 238, F2: (0.1 < 0.0)}
                            v1 = ((v1 + (((1e10 + (1e10 * v1)) / (((0.1 / -2.25) - float64((1e10 - 0.0))) + (0.0 - (3.14159 - (float64((1e10 * 3.14159)) - ((v1 - 0.1) + v1)))))) + ((float64(1e10) - v1) * v1))) + 1e10)
                            var v24 string = "a"
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    }
                    if (!false) {
                        var v25 uint = 65535
                        for fi26 := 0; fi26 < 9; fi26++ {
                            v25 -= (v25 ^ v25)
                            switch fi26 {
                            case 1:
                                var v27 int32 = -2147483648
                                _ = v27
                            case -9223372036854775808:
                                if ((387 != 42) && false) {
                                    v28 := GenS1{F0: -182, F1: 32767, F2: ("a" + "\n\t")}
                                    var v29 int64 = 9223372036854775807
                                    switch fi26 {
                                    case 128:
                                        delete(v3, 128)
                                    default:
                                        if true {
                                            v30 := GenS3{F0: 128, F1: 0, F2: ("a" != "")}
                                            if (443 != 2) {
                                                fmt.Println(fi26, v29, v4)
                                                fmt.Println(v29, fi26, v4)
                                            }
                                            v1 -= 0.1
                                            _ = v30
                                        }
                                    }
                                    _ = v28
                                    _ = v29
                                } else {
                                    if true {
                                        fmt.Println(v4)
                                    } else {
                                        fmt.Println(v1)
                                        fmt.Println(v25, fi15)
                                    }
                                    v31 := []int{}
                                    _ = v31
                                }
                            case -956:
                                var v32 float32 = 0.1
                                _ = v32
                            default:
                                var v33 int8 = 33
                                _ = v33
                            }
                            if (-2.25 <= 1e10) {
                                var v34 bool = (-868 <= -2147483648)
                                _ = v34
                            } else {
                                v35 := GenS1{F0: 42, F1: 2, F2: "foo bar"}
                                v4 = v4
                                if (42 >= 42) {
                                    var v36 float32 = (0.0 + (((((float32((float64(float32(v1)) + 0.1)) / float32(1e10)) - -2.25) / -2.25) - 3.14159) / 1.5))
                                    var v37 bool = false
                                    var v38 int16 = -1
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                                var v39 float32 = 1.5
                                _ = v35
                                _ = v39
                            }
                            if false {
                                v40 := []int{}
                                v1 = v1
                                _ = v40
                            } else {
                                v1 = ((v1 / (v1 + (((((v1 / (((v1 + (1.5 + v1)) * v1) - v1)) * (3.14159 - v1)) * (float64((1.5 - ((-2.25 / 1e10) + (0.1 / 1.5)))) / (1e10 - v1))) * v1) + v1))) * float64(0.1))
                                v41 := []int{762}
                                v25 = (v25 << (((v25 + v25)) & 63))
                                _ = v41
                            }
                            v25 = v25
                            _ = fi26
                        }
                        _ = v25
                    } else {
                        var v42 bool = true
                        _ = v42
                    }
                }
            }
        } else {
            var v43 uint64 = 404
            var v44 int32 = 1
            for fi45 := 0; fi45 < 16; fi45++ {
                switch (fi45 | (fi45 | fi15)) {
                case 32768:
                    v43 -= genFn0()
                case 381:
                    v46 := GenS4{F0: (v44 ^ (v44 - -867)), F1: v44, F2: 128}
                    _ = v46
                default:
                    if (!true) {
                        var v47 uint64 = v43
                        _ = v47
                    } else {
                        var v48 bool = true
                        v2[42] = fi15
                        if v48 {
                            v49 := GenS5{F0: false, F1: 1, F2: 2147483648}
                            _ = v49
                        } else {
                            v50 := []int{fi15}
                            var v51 uint64 = (^(v43))
                            var v52 int16 = genFn1("\n\t", 432)
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                        v53 := []int{}
                        if false {
                            fmt.Println(v1, v43)
                            v54 := GenS1{F0: (v44 & (v44 << ((65535) & 63))), F1: 724, F2: "a"}
                            _ = v54
                        } else {
                            var v55 int64 = 42
                            _ = len(v53)
                            var v56 uint8 = 127
                            _ = v55
                            _ = v56
                        }
                        v48 = v48
                        v1 += 1e10
                        _ = v48
                        _ = v53
                    }
                }
                if (!true) {
                    fmt.Println(fi45, v43)
                    var v57 int16 = 0
                    v58 := GenS6{F0: -32768, F1: (-(v44)), F2: 324}
                    _ = v57
                    _ = v58
                } else {
                    fmt.Println(fi15)
                    v59 := []int{32768, 903, 127}
                    v4 = (v4 - (^(v4)))
                    var v60 uint16 = 0
                    if false {
                        v43 -= (v43 & (^(v43)))
                        var v61 uint64 = genFn0()
                        _ = v61
                    } else {
                        v62 := GenS7{F0: v4, F1: fi15, F2: v44}
                        v43 = v43
                        _ = v62
                    }
                    _ = v59
                    _ = v60
                }
                v63 := GenS5{F0: true, F1: 1, F2: 1}
                _ = fi45
                _ = v63
            }
            _ = v43
            _ = v44
        }
        _ = fi15
    }
    v64 := []int{}
    if true {
        v4 = (v4 - (v4 << ((0) & 63)))
        v65 := GenS6{F0: v4, F1: 897, F2: -1}
        var v66 float64 = (v1 / float64(-2.25))
        _ = v65
        _ = v66
    }
    v1 = float64(1.5)
    for fi67 := 0; fi67 < 17; fi67++ {
        var v68 int8 = -1
        _ = fi67
        _ = v68
    }
    v64 = append(v64, 9223372036854775807)
    var v69 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v64:", len(v64), v64)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v69
}
