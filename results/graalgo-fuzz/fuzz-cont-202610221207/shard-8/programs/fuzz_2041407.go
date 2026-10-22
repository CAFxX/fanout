package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int; F1 uint64; F2 int }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int16; F1 uint; F2 int8 }
type GenS7 struct { F0 int8; F1 bool; F2 int }
type GenS8 struct { F0 int }
func genFn0() uint64 {
    var v7 int64 = -9223372036854775808
    v8 := []int{}
    var v9 int64 = v7
    var v10 uint = 39
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint8, p1 int32, p2 bool) uint64 {
    if (-498 >= -61) {
        v11 := []int{2147483647, -1}
        v12 := GenS1{F0: p1, F1: 576, F2: 875}
        for fi13 := 0; fi13 < 8; fi13++ {
            if false {
                p0 = p0
                var v14 int16 = -287
                for fi15 := 0; fi15 < 5; fi15++ {
                    v11 = append(v11, 1)
                    if p2 {
                        if (true || (-9223372036854775808 >= 9223372036854775807)) {
                            v14 -= v14
                            if (p2 && p2) {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                var v16 uint8 = (p0 % ((127) | 1))
                                if len(v11) > 0 {
                                    v11[0] = (fi15 | fi15)
                                }
                                if (65535 >= 65535) {
                                    v17 := GenS2{F0: 65535, F1: 2}
                                    _ = v17
                                } else {
                                    if len(v11) > 0 {
                                        v11[0] = 32767
                                    }
                                    p2 = ((1 <= 255) && true)
                                }
                                _ = v16
                            } else {
                                var v18 uint8 = p0
                                _ = v18
                            }
                        } else {
                            var v19 int32 = p1
                            v20 := GenS2{F0: 0, F1: 0}
                            var v21 int64 = 1
                            var v22 int64 = (v21 + (-(v21)))
                            _ = v19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    } else {
                        v14 += (-(v14))
                        var v23 uint16 = 614
                        var v24 int = int(p1)
                        _ = v23
                        _ = v24
                    }
                    switch (fi13 ^ fi13) {
                    case -2147483648:
                        v25 := GenS3{F0: 42, F1: (fi13 - fi15)}
                        _ = v25
                    case 4294967295:
                        var v26 int8 = 0
                        _ = v26
                    case 32767:
                        v14 -= 42
                    default:
                        switch "" {
                        case "\n\t":
                            p0 = (p0 << ((2) & 63))
                        }
                    }
                    _ = fi15
                }
                _ = v14
            } else {
                p0 = 0
            }
            _ = len(v11)
            _ = fi13
        }
        _ = v11
        _ = v12
    } else {
        var v27 uint = 42
        for fi28 := 0; fi28 < 3; fi28++ {
            switch "\n\t" {
            case "foo bar":
                v29 := GenS4{F0: (fi28 >> (((v27 % (((v27 >> (((^(v27))) & 63))) | 1))) & 63)), F1: 18446744073709551615, F2: fi28}
                _ = v29
            case "\n\t":
                if true {
                    switch fi28 {
                    case 32768:
                        var v30 uint32 = 0
                        _ = v30
                    case 26:
                        p0 += p0
                    case -9223372036854775808:
                        v27 = (^(v27))
                    default:
                        v31 := []int{fi28, (fi28 >> (((v27 - 2)) & 63)), (fi28 / (((fi28 / ((int(p0)) | 1))) | 1)), fi28}
                        _ = v31
                    }
                    p0 -= p0
                    v32 := []int{}
                    _ = v32
                }
            case "a":
                v33 := []int{(fi28 | fi28), (fi28 + int(p1))}
                _ = v33
            default:
                var v34 uint32 = 2
                _ = v34
            }
            switch 127 {
            case -2147483648:
                v35 := []int{fi28, 2147483647, (-(fi28))}
                _ = v35
            }
            _ = fi28
        }
        v27 -= v27
        _ = v27
    }
    if p2 {
        var v36 bool = p2
        _ = v36
    }
    var v37 uint8 = (p0 & p0)
    v38 := GenS3{F0: 42, F1: 415}
    var v39 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    _ = v38
    _ = v39
    return 18446744073709551615
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 string = "a"
    v3 := make(map[string]bool)
    var v4 uint64 = 573
    var v5 int8 = 100
    v6 := []int{-2147483648, -752, 163, 127}
    v5 = v5
    v4 += (v4 % ((26) | 1))
    _ = len(v3)
    v40 := GenS5{F0: 1, F1: 32767}
    if (!((!(0 == (v5 | -86))) && (4294967295 <= 42))) {
        fmt.Println(v1)
        v41 := GenS6{F0: -32768, F1: 774, F2: 100}
        v42 := GenS5{F0: 162, F1: 414}
        _ = v41
        _ = v42
    } else {
        v4 -= v4
    }
    if ((v4 % ((v4) | 1)) != v4) {
        switch (v2 + v2) {
        case "":
            v1 = (v1 - v1)
        case "hello":
            switch "foo bar" {
            case "":
                switch "\n\t" {
                case "hello":
                    for fi43 := 0; fi43 < 9; fi43++ {
                        for fi44 := 0; fi44 < 12; fi44++ {
                            v1 += v1
                            v2 = v2
                            v45 := v3[v2]
                            _ = fi44
                            _ = v45
                        }
                        var v46 uint8 = 142
                        fmt.Println(v2)
                        _ = fi43
                        _ = v46
                    }
                case "a":
                    v4 = 2
                }
            default:
                var v47 float64 = -2.25
                _ = v47
            }
        case "a":
            if true {
                delete(v3, v2)
                for fi48 := 0; fi48 < 9; fi48++ {
                    delete(v3, ("" + v2))
                    _ = fi48
                }
                switch v2 {
                case "foo bar":
                    if false {
                        if len(v6) > 0 {
                            v6[0] = 32767
                        }
                        _ = len(v6)
                        for fi49 := 0; fi49 < 13; fi49++ {
                            for fi50 := 0; fi50 < 4; fi50++ {
                                if (2147483647 >= -757) {
                                    var v51 uint8 = 128
                                    switch 4294967295 {
                                    case 0:
                                        v5 += 62
                                    default:
                                        if (!((v51 << ((255) & 63)) >= (v51 << ((4294967295) & 63)))) {
                                            v3[v2] = true
                                            var v52 int16 = -32768
                                            _ = v52
                                        }
                                    }
                                    var v53 string = v2
                                    switch (fi49 & -205) {
                                    case 127:
                                        v54 := GenS7{F0: -61, F1: true, F2: -304}
                                        _ = v54
                                    case -907:
                                        var v55 float32 = -2.25
                                        v4 -= 0
                                        v55 = 0.1
                                        _ = v55
                                    default:
                                        v1 = (v1 - v1)
                                    }
                                    _ = v51
                                    _ = v53
                                } else {
                                    var v56 int = (fi49 << ((1) & 63))
                                    v57 := GenS6{F0: -973, F1: 42, F2: -128}
                                    v4 = (^(v4))
                                    _ = v56
                                    _ = v57
                                }
                                fmt.Println(fi49, v5, fi49)
                                _ = fi50
                            }
                            fmt.Println(v4, v2)
                            _ = fi49
                        }
                    } else {
                        v3["hello"] = false
                    }
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "hello":
                    var v58 bool = true
                    _ = v58
                }
            } else {
                switch v2 {
                case "a":
                    v4 = (v4 * (v4 & 2))
                case "\n\t":
                    switch v2 {
                    case "":
                        if ((v5 % (((-(v5))) | 1)) > (v5 / (((v5 << ((173) & 63))) | 1))) {
                            for fi59 := 0; fi59 < 2; fi59++ {
                                if true {
                                    var v60 float64 = 0.0
                                    v2 = v2
                                    v4 -= v4
                                    _ = v60
                                }
                                _ = fi59
                            }
                            _ = len(v3)
                            var v61 int8 = (v5 & (v5 << ((0) & 63)))
                            _ = v61
                        } else {
                            var v62 float32 = (-2.25 + (0.1 * (0.0 * float32(-2.25))))
                            _ = v62
                        }
                    case "\n\t":
                        v4 = (v4 + v4)
                    }
                }
                var v63 uint8 = 2
                _ = v63
            }
        }
        var v64 uint64 = (v4 >> ((216) & 63))
        v65 := v3[v2]
        _ = v64
        _ = v65
    }
    v5 -= v5
    var v66 int = 42
    for fi67 := 0; fi67 < 5; fi67++ {
        var v68 float32 = -2.25
        for fi69 := 0; fi69 < 3; fi69++ {
            if false {
                var v70 int16 = -1
                var v71 int = fi67
                _ = v70
                _ = v71
            } else {
                if true {
                    v68 = (((v68 / ((1.5 * 1e10) + (3.14159 * float32(1e10)))) / v68) - float32(0.0))
                } else {
                    if ((^(v4)) <= v4) {
                        if true {
                            v5 += v5
                            var v72 bool = ("a" > "")
                            _ = v72
                        } else {
                            v66 = fi67
                        }
                        v73 := GenS2{F0: 2147483648, F1: 1}
                        v1 -= (v1 & v1)
                        v2 = "a"
                        v2 = ((v2 + ((((v2 + ("hello" + (v2 + v2))) + (v2 + (("" + v2) + "foo bar"))) + (("hello" + v2) + "\n\t")) + (v2 + v2))) + v2)
                        _ = v73
                    }
                    switch (fi67 & (v66 + v66)) {
                    case 42:
                        var v74 bool = false
                        _ = v74
                    case 1:
                        switch (-(v66)) {
                        case 9223372036854775807:
                            v68 = (v68 / v68)
                        case 1:
                            var v75 int8 = v5
                            _ = v75
                        case -1:
                            v1 -= 42
                        default:
                            v5 = 127
                        }
                    }
                }
            }
            v76 := v3[(v2 + v2)]
            _ = fi69
            _ = v76
        }
        _ = fi67
        _ = v68
    }
    v77 := GenS8{F0: int(v4)}
    switch (v2 + ((v2 + v2) + (((v2 + ("foo bar" + v2)) + v2) + v2))) {
    case "foo bar":
        var v78 uint16 = 255
        _ = v78
    }
    if (((v2 + v2) + "a") < ("\n\t" + v2)) {
        v66 = v66
    } else {
        _ = len(v6)
        fmt.Println(v5, v4, v5)
        v1 = (v1 / (((v1 << ((65535) & 63))) | 1))
    }
    fmt.Println(v5, v1, v4)
    if (false || true) {
        switch "\n\t" {
        case "":
            switch v2 {
            case "a":
                v5 += v5
            case "foo bar":
                v79 := GenS4{F0: -831, F1: (v4 - (^(v4))), F2: v66}
                _ = v79
            case "\n\t":
                switch "a" {
                case "":
                    v5 = -25
                }
            }
        case "a":
            if len(v6) > 0 {
                v6[0] = -9223372036854775808
            }
        case "foo bar":
            if false {
                fmt.Println(v66, v2, v4)
                v5 += v5
                var v80 int8 = int8(v66)
                _ = v80
            }
        }
        var v81 int32 = -1
        switch -2147483648 {
        case 1:
            var v82 int8 = v5
            _ = v82
        case 2147483647:
            var v83 uint64 = v4
            _ = v83
        }
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v66
    _ = v77
}
