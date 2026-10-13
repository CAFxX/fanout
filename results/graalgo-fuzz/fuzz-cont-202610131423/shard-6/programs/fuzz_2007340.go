package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 bool; F1 bool; F2 uint8 }
type GenS8 struct { F0 uint16; F1 string }
func genFn0() uint32 {
    if false {
        switch 0 {
        case 4294967295:
            v6 := GenS1{F0: 1}
            _ = v6
        default:
            var v7 uint32 = 42
            _ = v7
        }
    }
    v8 := GenS1{F0: 774}
    var v9 int = 26
    var v10 float64 = 1e10
    v10 += 0.0
    _ = v10
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 string) uint32 {
    v11 := GenS2{F0: p0, F1: 9223372036854775807}
    var v12 int32 = 0
    var v13 uint64 = 18446744073709551615
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 int32 = 42
    var v4 uint64 = 1
    var v5 uint64 = v4
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 bool = (755 > 9223372036854775807)
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 uint16 = 636
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    switch 2147483647 {
    case 1:
        switch "" {
        case "\n\t":
            var v18 int8 = -123
            _ = v18
        case "hello":
            for fi19 := 0; fi19 < 3; fi19++ {
                if (!(!true)) {
                    if true {
                        var v20 int = fi19
                        switch 1 {
                        case 9223372036854775807:
                            var v21 string = "a"
                            _ = v21
                        case -869:
                            v22 := []int{495, (v20 * 202), (v20 | (-(fi19)))}
                            _ = v22
                        case 1:
                            if true {
                                var v23 float64 = (float64((0.1 + ((((1e10 / 1.0) / (1e10 * (1e10 / (1.5 / (0.1 * 3.14159))))) - 1e10) / float32((0.1 - (3.14159 / 3.14159)))))) + 1e10)
                                if false {
                                    if (true || true) {
                                        if (((float64(-2.25) / 3.14159) + v23) < v23) {
                                            if (float64((3.14159 + 0.1)) == -2.25) {
                                                var v24 int64 = 9223372036854775807
                                                v24 += -180
                                                if (false || (!false)) {
                                                    switch (v20 + 9223372036854775807) {
                                                    case -691:
                                                        for fi25 := 0; fi25 < 19; fi25++ {
                                                            v20 = (fi19 << ((18446744073709551615) & 63))
                                                            fmt.Println(v20, fi19, v3)
                                                            var v26 bool = (!(-55 > -116))
                                                            v24 += v24
                                                            var v27 uint64 = v4
                                                            _ = fi25
                                                            _ = v26
                                                            _ = v27
                                                        }
                                                    }
                                                }
                                                _ = v24
                                            } else {
                                                if (int64(v20) > 0) {
                                                    if false {
                                                        v20 += v20
                                                        switch ((("hello" + "\n\t") + "") + "\n\t") {
                                                        case "a":
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v28 uint8 = 39
                                                                _ = v28
                                                            case "":
                                                                fmt.Println(v20, v23, v20)
                                                            default:
                                                                var v29 int16 = 1
                                                                _ = v29
                                                            }
                                                        case "\n\t":
                                                            v3 += v3
                                                        case "foo bar":
                                                            for fi30 := 0; fi30 < 8; fi30++ {
                                                                v5 += v4
                                                                if (32767 <= 0) {
                                                                    var v31 uint32 = 65535
                                                                    var v32 int16 = 32767
                                                                    v5 += (^(v4))
                                                                    _ = v31
                                                                    _ = v32
                                                                } else {
                                                                    var v33 uint = 0
                                                                    _ = v33
                                                                }
                                                                switch "" {
                                                                case "a":
                                                                    switch v20 {
                                                                    case 32767:
                                                                        fmt.Println(v4)
                                                                    case -2147483648:
                                                                        v2[2] = "\n\t"
                                                                    case 495:
                                                                        v34 := []int{}
                                                                        var v35 int16 = 107
                                                                        _ = v34
                                                                        _ = v35
                                                                    default:
                                                                        fmt.Println(fi19)
                                                                    }
                                                                }
                                                                _ = fi30
                                                            }
                                                        }
                                                    }
                                                }
                                                var v36 uint64 = 18446744073709551615
                                                _ = v36
                                            }
                                        } else {
                                            for fi37 := 0; fi37 < 6; fi37++ {
                                                v5 += (^(v4))
                                                _ = fi37
                                            }
                                            switch ("" + "hello") {
                                            case "":
                                                v5 += (v4 | (^(v4)))
                                            case "a":
                                                var v38 bool = (!((v4 != (^(v5))) && false))
                                                _ = v38
                                            case "hello":
                                                if true {
                                                    switch 32767 {
                                                    case 128:
                                                        v39 := GenS3{F0: 128, F1: 1, F2: false}
                                                        _ = v39
                                                    case 127:
                                                        var v40 string = "\n\t"
                                                        _ = v40
                                                    case -1:
                                                        var v41 bool = ((true || false) || (2 >= 2))
                                                        _ = v41
                                                    }
                                                } else {
                                                    if (818 < 65535) {
                                                        if false {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                var v42 int8 = 42
                                                                _ = v42
                                                            case "a":
                                                                var v43 float32 = ((float32(((v23 + ((v23 + v23) - v23)) - float64(1e10))) - 1.5) / 0.1)
                                                                _ = v43
                                                            }
                                                            v44 := GenS4{F0: 255}
                                                            v3 = v3
                                                            _ = v44
                                                        }
                                                        v23 = ((0.1 + ((1e10 / (((-2.25 / v23) / (1e10 * -2.25)) / ((v23 * (1e10 * (v23 - 1.5))) + v23))) + 1e10)) * (v23 + 1.5))
                                                    } else {
                                                        v45 := v2[genFn0()]
                                                        _ = v45
                                                    }
                                                    v46 := GenS1{F0: 65535}
                                                    _ = v46
                                                }
                                            default:
                                                v5 -= v5
                                            }
                                            var v47 uint64 = v5
                                            _ = v47
                                        }
                                        var v48 uint32 = genFn1("\n\t")
                                        _ = v48
                                    } else {
                                        switch v20 {
                                        case -994:
                                            var v49 bool = false
                                            _ = v49
                                        }
                                        delete(v1, (v20 ^ v20))
                                        v50 := []int{(fi19 ^ (-(fi19)))}
                                        _ = v50
                                    }
                                }
                                _ = v23
                            }
                        default:
                            if false {
                                if true {
                                    switch "foo bar" {
                                    case "hello":
                                        v20 -= fi19
                                    case "a":
                                        v5 = 4294967295
                                    case "":
                                        v3 += v3
                                    }
                                    v2[42] = "hello"
                                    v51 := v2[2147483648]
                                    _ = v51
                                }
                                for fi52 := 0; fi52 < 6; fi52++ {
                                    var v53 float64 = -2.25
                                    v54 := GenS5{F0: 255, F1: v4}
                                    var v55 uint16 = 2
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                                var v56 int32 = (v3 / (((v3 - (-(v3)))) | 1))
                                _ = v56
                            } else {
                                v57 := []int{}
                                for fi58 := 0; fi58 < 8; fi58++ {
                                    switch ("foo bar" + "hello") {
                                    case "\n\t":
                                        v20 = (^(fi58))
                                    case "foo bar":
                                        var v59 uint = 255
                                        _ = v59
                                    }
                                    if (!true) {
                                        fmt.Println(v4)
                                        var v60 uint32 = 883
                                        _ = v60
                                    } else {
                                        if true {
                                            fmt.Println(fi19)
                                            if (!(-1 >= 1000)) {
                                                var v61 int32 = -2147483648
                                                var v62 int8 = 1
                                                v4 += v4
                                                v57 = append(v57, (v20 * fi19))
                                                _ = v61
                                                _ = v62
                                            }
                                        }
                                    }
                                    v63 := GenS4{F0: 1}
                                    _ = fi58
                                    _ = v63
                                }
                                _ = v57
                            }
                        }
                        _ = v20
                    }
                    switch (-(fi19)) {
                    case 2147483647:
                        v4 -= v5
                    default:
                        var v64 uint8 = 255
                        _ = v64
                    }
                }
                _ = fi19
            }
        }
    default:
        var v65 int8 = 1
        _ = v65
    }
    fmt.Println(v4, v3, v4)
    for fi66 := 0; fi66 < 13; fi66++ {
        fmt.Println(v4, v4, fi66)
        v67 := GenS6{F0: 127, F1: 1}
        var v68 uint32 = 50
        _ = fi66
        _ = v67
        _ = v68
    }
    v1[42] = -521
    for fi69 := 0; fi69 < 4; fi69++ {
        v4 -= v4
        v70 := []int{(fi69 - (^(fi69))), fi69, fi69}
        _ = fi69
        _ = v70
    }
    var v71 int32 = (^(v3))
    var v72 int8 = -128
    switch 823 {
    case -9223372036854775808:
        v73 := []int{}
        _ = v73
    case 2147483647:
        fmt.Println(v72, v3)
    case 0:
        v74 := GenS7{F0: true, F1: true, F2: 0}
        _ = v74
    }
    for fi75 := 0; fi75 < 13; fi75++ {
        if ((3.14159 != 0.0) && true) {
            var v76 float32 = (1e10 + (3.14159 + ((((1e10 + 3.14159) - (float32(0.0) / -2.25)) - (-2.25 / 1e10)) + 1.5)))
            delete(v1, (fi75 * (fi75 * fi75)))
            _ = v76
        } else {
            v71 = v3
            v72 -= (^(v72))
            var v77 int32 = 116
            _ = v77
        }
        fmt.Println(v71)
        _ = fi75
    }
    switch 776 {
    case -871:
        v72 -= (v72 + (v72 ^ (v72 - v72)))
    case 93:
        v5 += (v4 | v4)
    case 128:
        v78 := []int{}
        _ = v78
    }
    var v79 float32 = (float32(0.1) / (3.14159 * 1.5))
    if true {
        var v80 uint16 = 32767
        var v81 string = ""
        v4 = v5
        _ = v80
        _ = v81
    } else {
        v82 := GenS5{F0: 311, F1: v4}
        _ = v82
    }
    switch 1 {
    case 49:
        v83 := []int{-908, -824, 553, 42}
        _ = v83
    case 585:
        switch 1 {
        case -244:
            fmt.Println(v5, v72)
        case 760:
            fmt.Println(v71)
        default:
            v84 := []int{-794}
            _ = v84
        }
    default:
        v85 := GenS8{F0: 32767, F1: ""}
        _ = v85
    }
    var v86 int32 = v71
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v79:", v79)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v72
    _ = v79
    _ = v86
}
