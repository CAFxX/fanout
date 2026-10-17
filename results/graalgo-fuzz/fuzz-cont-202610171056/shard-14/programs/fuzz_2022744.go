package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint32; F1 uint; F2 string }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int; F1 int16; F2 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 int16 }

func main() {
    var v1 float32 = 0.1
    var v2 uint16 = 2
    var v3 uint32 = 0
    v4 := make(map[int]int)
    v5 := []int{-805}
    v6 := []int{-2147483648, 4294967295}
    if false {
        v3 = 1
        var v7 uint = 1
        _ = v7
    }
    if (1 <= 128) {
        v3 -= (v3 % (((v3 * 2147483647)) | 1))
        if len(v6) > 0 {
            v6[0] = 42
        }
    }
    v6 = append(v6, -9223372036854775808)
    v2 = 2
    var v8 int8 = 100
    for fi9 := 0; fi9 < 11; fi9++ {
        v3 = v3
        v10 := GenS1{F0: (v2 - v2), F1: true}
        _ = fi9
        _ = v10
    }
    if ((1 >= 1) || (int16(v8) >= 1)) {
        var v11 float32 = (1e10 / (v1 - 1.5))
        v2 = 42
        _ = v11
    }
    v1 += float32(0.1)
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 uint64 = 18446744073709551615
            if false {
                v15 := GenS1{F0: (v2 + (v2 & (v2 - 42))), F1: false}
                if false {
                    var v16 float32 = 1e10
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v17 := GenS1{F0: v2, F1: ((((!(v8 > v8)) && (true && false)) && false) && (!true))}
                    _ = v16
                    _ = v17
                } else {
                    if true {
                        fmt.Println(v2)
                        v4[(fi13 & fi12)] = (fi12 ^ fi13)
                    } else {
                        fmt.Println(v8, v1, v8)
                        if (1 >= (v3 % ((1) | 1))) {
                            var v18 float32 = 1.5
                            var v19 uint64 = 172
                            var v20 uint = 4294967295
                            var v21 int = (fi13 + (^(fi12)))
                            switch "foo bar" {
                            case "a":
                                v22 := v4[fi13]
                                v4[(fi12 >> (((v20 + v20)) & 63))] = fi12
                                _ = len(v4)
                                _ = v22
                            }
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            v23 := GenS1{F0: (v2 | (v2 << ((1) & 63))), F1: true}
                            if false {
                                v1 -= (v1 + v1)
                                fmt.Println(v2)
                                if false {
                                    if (2 != (v3 & 4294967295)) {
                                        v24 := GenS2{F0: 2, F1: false}
                                        v2 = v2
                                        delete(v4, int(v14))
                                        _ = v24
                                    }
                                    v2 = (v2 + (^(v2)))
                                } else {
                                    switch "" {
                                    case "\n\t":
                                        if (false && (!false)) {
                                            v25 := GenS3{F0: (v3 << ((42) & 63)), F1: 18446744073709551615, F2: "foo bar"}
                                            switch ("" + "a") {
                                            case "a":
                                                v6 = append(v6, fi13)
                                            case "hello":
                                                v1 = v1
                                            default:
                                                v1 -= float32(float64(0.0))
                                            }
                                            if ((v3 != (v3 / ((v3) | 1))) && (true || true)) {
                                                fmt.Println(v8)
                                                var v26 string = "a"
                                                var v27 bool = true
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v2 -= (^(v2))
                                            }
                                            _ = v25
                                        }
                                    }
                                    switch "\n\t" {
                                    case "":
                                        switch fi12 {
                                        case 383:
                                            v1 = v1
                                        case 135:
                                            v1 -= (((3.14159 * (((1.5 * (v1 * v1)) / ((1e10 + float32(0.1)) + v1)) / 1e10)) + 0.1) + (float32(0.0) * v1))
                                        case 311:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            var v28 float32 = (((v1 - -2.25) - (0.0 - v1)) * v1)
                                            fmt.Println(v8, v14)
                                            v29 := GenS4{F0: 2, F1: 128}
                                            switch -1 {
                                            case -1:
                                                if false {
                                                    switch "hello" {
                                                    case "hello":
                                                        var v30 int32 = -2147483648
                                                        _ = len(v6)
                                                        _ = v30
                                                    case "foo bar":
                                                        v2 -= v2
                                                    default:
                                                        switch fi13 {
                                                        case -177:
                                                            v1 += v1
                                                        case 9223372036854775807:
                                                            var v31 int16 = 1
                                                            _ = v31
                                                        default:
                                                            var v32 int64 = -9223372036854775808
                                                            _ = v32
                                                        }
                                                    }
                                                }
                                            }
                                            v28 -= ((3.14159 / ((v28 - v28) * (1.5 + -2.25))) / 1.5)
                                            _ = v28
                                            _ = v29
                                        }
                                        if false {
                                            if (771 == 42) {
                                                var v33 uint8 = 179
                                                _ = v33
                                            } else {
                                                v14 -= v14
                                            }
                                            var v34 uint64 = v14
                                            v14 = 0
                                            _ = v34
                                        } else {
                                            v35 := GenS2{F0: 169, F1: true}
                                            _ = v35
                                        }
                                        var v36 int64 = 42
                                        _ = v36
                                    case "\n\t":
                                        v37 := GenS5{F0: fi13, F1: 0, F2: 31}
                                        _ = v37
                                    default:
                                        if len(v5) > 0 {
                                            v5[0] = (fi13 << ((uint(v14)) & 63))
                                        }
                                    }
                                    switch (fi13 + fi12) {
                                    case 2147483647:
                                        v38 := GenS2{F0: 1, F1: true}
                                        _ = v38
                                    case 127:
                                        v39 := GenS1{F0: 2, F1: false}
                                        _ = v39
                                    default:
                                        v40 := GenS1{F0: v2, F1: true}
                                        _ = v40
                                    }
                                    v41 := v4[(fi13 & fi13)]
                                    _ = v41
                                }
                            }
                            if false {
                                var v42 int8 = (-(v8))
                                _ = v42
                            } else {
                                v43 := GenS2{F0: 30, F1: (true || true)}
                                var v44 int16 = -1
                                _ = v43
                                _ = v44
                            }
                            _ = v23
                        }
                        var v45 int32 = 0
                        _ = v45
                    }
                }
                var v46 int64 = -171
                _ = v15
                _ = v46
            } else {
                v3 -= 0
            }
            var v47 float64 = 1e10
            _ = fi13
            _ = v14
            _ = v47
        }
        _ = fi12
    }
    v48 := GenS6{F0: (v8 + (-(v8)))}
    for fi49 := 0; fi49 < 17; fi49++ {
        if true {
            v3 -= 0
        }
        for fi50 := 0; fi50 < 12; fi50++ {
            v2 = v2
            fmt.Println(v1, v8)
            v51 := GenS5{F0: fi50, F1: 42, F2: 166}
            switch int(v8) {
            case 42:
                v1 = 0.1
            case -1:
                fmt.Println(v2, v3)
            case -2147483648:
                if false {
                    var v52 uint8 = 128
                    _ = v52
                }
            }
            if false {
                _ = len(v6)
                if (2 > 110) {
                    v53 := GenS7{F0: "a"}
                    v3 = uint32(v2)
                    _ = v53
                } else {
                    v3 = v3
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v54 uint = 505
                    var v55 string = ""
                    var v56 float32 = v1
                    _ = v54
                    _ = v55
                    _ = v56
                }
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            delete(v4, fi49)
            v57 := GenS8{F0: (v2 % ((1) | 1)), F1: 1}
            _ = fi50
            _ = v51
            _ = v57
        }
        v3 += (v3 << ((0) & 63))
        _ = fi49
    }
    v58 := GenS5{F0: -614, F1: 1, F2: 255}
    switch ("hello" + "\n\t") {
    case "\n\t":
        v4[-559] = -43
    case "hello":
        switch -273 {
        case -9223372036854775808:
            var v59 float64 = (3.14159 + 3.14159)
            _ = v59
        }
    default:
        switch "" {
        case "":
            for fi60 := 0; fi60 < 9; fi60++ {
                v2 += v2
                v61 := GenS7{F0: "a"}
                v2 = (v2 >> ((uint(v3)) & 63))
                _ = fi60
                _ = v61
            }
        case "foo bar":
            for fi62 := 0; fi62 < 12; fi62++ {
                for fi63 := 0; fi63 < 18; fi63++ {
                    v2 += v2
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi63
                }
                v64 := GenS3{F0: (v3 + v3), F1: 476, F2: "foo bar"}
                _ = fi62
                _ = v64
            }
        case "\n\t":
            switch "foo bar" {
            case "a":
                v3 = 0
            case "":
                for fi65 := 0; fi65 < 7; fi65++ {
                    var v66 int32 = -1
                    var v67 uint16 = (v2 & 777)
                    var v68 string = "\n\t"
                    _ = fi65
                    _ = v66
                    _ = v67
                    _ = v68
                }
            }
        }
    }
    v69 := GenS6{F0: 0}
    v1 += 1e10
    v2 = (v2 & v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v58
    _ = v6
    _ = v69
    _ = v8
}
