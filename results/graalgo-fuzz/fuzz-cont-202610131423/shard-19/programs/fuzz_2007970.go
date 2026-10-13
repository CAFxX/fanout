package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 string; F2 uint64 }
type GenS6 struct { F0 int32; F1 uint; F2 uint16 }
type GenS7 struct { F0 int8; F1 uint8 }
func genFn0(p0 int8, p1 string, p2 uint64) uint8 {
    switch p1 {
    case "a":
        v4 := GenS1{F0: -32768, F1: 65535}
        _ = v4
    case "":
        var v5 int32 = -205
        _ = v5
    case "\n\t":
        var v6 uint32 = 755
        _ = v6
    }
    fmt.Println(p1, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 uint32 = 661
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    for fi7 := 0; fi7 < 9; fi7++ {
        if false {
            if (!false) {
                var v8 uint64 = 0
                v8 += v8
                _ = v8
            }
            for fi9 := 0; fi9 < 12; fi9++ {
                v1 -= v1
                switch fi9 {
                case 32768:
                    v1 -= (v1 % ((65535) | 1))
                case -9223372036854775808:
                    var v10 int16 = -787
                    v1 += 4294967295
                    v11 := GenS1{F0: v10, F1: 4294967295}
                    _ = v10
                    _ = v11
                case 4294967295:
                    switch "foo bar" {
                    case "a":
                        if (32767 < 887) {
                            v12 := GenS2{F0: -32768}
                            _ = v12
                        }
                    case "":
                        v13 := []int{(fi7 % (((fi7 >> ((0) & 63))) | 1)), (fi7 + 32768), (fi7 & (fi7 | fi9)), 9223372036854775807}
                        _ = v13
                    case "foo bar":
                        v1 = 0
                        fmt.Println(fi9, fi9, fi9)
                    }
                default:
                    v1 -= v1
                    v14 := GenS1{F0: 1000, F1: 94}
                    v1 += 4294967295
                    _ = v14
                }
                _ = fi9
            }
            v1 = v1
        } else {
            delete(v2, (v1 / (((^(v1))) | 1)))
            var v15 string = (("\n\t" + "\n\t") + "hello")
            v16 := GenS3{F0: 1, F1: -383}
            _ = v15
            _ = v16
        }
        if (1e10 < 0.1) {
            delete(v2, 158)
            if false {
                if (!true) {
                    if false {
                        var v17 int16 = 55
                        v17 += (v17 - v17)
                        _ = v17
                    } else {
                        var v18 int64 = 42
                        _ = v18
                    }
                }
                v1 += (v1 % ((v1) | 1))
            } else {
                v1 = v1
                switch fi7 {
                case -1:
                    if (0 < 0) {
                        v19 := GenS4{F0: 9223372036854775807}
                        var v20 uint64 = 4294967295
                        var v21 bool = false
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        fmt.Println(v1, fi7)
                        var v22 uint = 2
                        v23 := v3[fi7]
                        _ = v22
                        _ = v23
                    }
                case 128:
                    v24 := GenS5{F0: 2, F1: (("foo bar" + "hello") + ("\n\t" + "hello")), F2: 384}
                    _ = v24
                case 2147483647:
                    switch "foo bar" {
                    case "":
                        if ("" <= "") {
                            v25 := GenS2{F0: 0}
                            var v26 int16 = 0
                            _ = v25
                            _ = v26
                        } else {
                            if true {
                                for fi27 := 0; fi27 < 20; fi27++ {
                                    v28 := []int{(^(fi7)), fi27, (fi7 >> ((2) & 63))}
                                    var v29 int64 = -747
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                                if (!(-9223372036854775808 >= 1)) {
                                    for fi30 := 0; fi30 < 17; fi30++ {
                                        fmt.Println(v1, fi30)
                                        var v31 int64 = -1
                                        v32 := []int{fi30, fi30, fi7}
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                    }
                                    var v33 float64 = float64((0.0 * float32(1e10)))
                                    v33 = v33
                                    _ = v33
                                } else {
                                    var v34 float64 = (1e10 * -2.25)
                                    if true {
                                        for fi35 := 0; fi35 < 12; fi35++ {
                                            v36 := GenS1{F0: 42, F1: 42}
                                            _ = fi35
                                            _ = v36
                                        }
                                        var v37 bool = (float64(3.14159) <= v34)
                                        _ = v37
                                    }
                                    _ = v34
                                }
                            }
                        }
                    case "foo bar":
                        for fi38 := 0; fi38 < 11; fi38++ {
                            fmt.Println(v1, v1)
                            v39 := GenS2{F0: 32767}
                            _ = fi38
                            _ = v39
                        }
                    default:
                        v40 := []int{(^(fi7))}
                        _ = v40
                    }
                default:
                    var v41 uint16 = 0
                    _ = v41
                }
            }
            switch (("hello" + "hello") + "a") {
            case "\n\t":
                var v42 bool = true
                _ = v42
            }
        } else {
            v43 := GenS3{F0: 1, F1: 469}
            _ = v43
        }
        v1 = v1
        _ = fi7
    }
    fmt.Println(v1)
    v3[-445] = (!false)
    if (true || (329 >= 0)) {
        for fi44 := 0; fi44 < 11; fi44++ {
            if true {
                v45 := []int{-2147483648, fi44, fi44}
                _ = v45
            }
            v1 = (v1 & v1)
            _ = fi44
        }
        var v46 int16 = -270
        _ = v46
    } else {
        v47 := GenS3{F0: -128, F1: 719}
        _ = v47
    }
    switch (("hello" + "\n\t") + "\n\t") {
    case "hello":
        var v48 float64 = 1.5
        _ = v48
    case "":
        var v49 int32 = -201
        _ = v49
    case "a":
        var v50 uint = 14
        _ = v50
    }
    var v51 string = ("foo bar" + "")
    var v52 uint32 = (v1 & 596)
    if ((42 != -93) || (-2147483648 > 1)) {
        fmt.Println(v52)
        var v53 int32 = -1
        for fi54 := 0; fi54 < 14; fi54++ {
            var v55 uint64 = 2
            _ = fi54
            _ = v55
        }
        _ = v53
    } else {
        for fi56 := 0; fi56 < 10; fi56++ {
            v57 := GenS6{F0: 42, F1: 716, F2: 927}
            v1 -= v52
            _ = fi56
            _ = v57
        }
    }
    var v58 uint8 = 1
    for fi59 := 0; fi59 < 3; fi59++ {
        if true {
            switch "a" {
            case "":
                var v60 float64 = (float64(0.0) + 0.1)
                _ = v60
            }
            var v61 uint8 = 255
            var v62 uint32 = v1
            _ = v61
            _ = v62
        } else {
            if (42 < 1) {
                for fi63 := 0; fi63 < 13; fi63++ {
                    var v64 uint32 = (v1 - v1)
                    v65 := []int{(fi63 * fi59), 127, 2147483647, fi63}
                    if (208 < 359) {
                        switch (-(fi59)) {
                        case 32767:
                            var v66 int = 0
                            _ = v66
                        case 0:
                            v52 += v52
                            var v67 bool = false
                            _ = v67
                        }
                    } else {
                        delete(v3, fi59)
                        fmt.Println(v58, v64, v51)
                        v68 := GenS7{F0: -1, F1: v58}
                        _ = v68
                    }
                    _ = fi63
                    _ = v64
                    _ = v65
                }
            }
            var v69 uint = 2
            _ = v69
        }
        var v70 float64 = 1e10
        v51 = v51
        _ = fi59
        _ = v70
    }
    var v71 int64 = 9223372036854775807
    for fi72 := 0; fi72 < 8; fi72++ {
        var v73 uint32 = (v1 >> ((4294967295) & 63))
        v74 := []int{fi72}
        _ = fi72
        _ = v73
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v58
    _ = v71
}
