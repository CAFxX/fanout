package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint64; F1 int8; F2 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 int8; F1 uint8 }
func genFn0() uint16 {
    var v7 bool = (true || (!true))
    v8 := []int{}
    v7 = (834 == 42)
    v8 = append(v8, 2147483647)
    _ = v7
    _ = v8
    return 2
}
func genFn1() int64 {
    var v9 float64 = 1e10
    v10 := []int{-435, 188}
    v11 := GenS1{F0: "hello", F1: -775, F2: 255}
    _ = v10
    _ = v11
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 127
    var v2 int = 871
    v3 := make(map[uint32]uint32)
    v4 := []uint32{2, 4294967295}
    v5 := []int{(v2 | (^(v2))), 9223372036854775807, v2, 32768}
    var v6 uint64 = 18446744073709551615
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(fi12, v1)
            _ = fi13
        }
        v14 := GenS2{F0: 1}
        _ = fi12
        _ = v14
    }
    if (!false) {
        v15 := GenS2{F0: -1}
        var v16 uint32 = 920
        _ = v15
        _ = v16
    } else {
        var v17 uint64 = (v6 % ((v6) | 1))
        _ = len(v4)
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        v2 = v2
        fmt.Println(fi18, fi18, v2)
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 float32 = (float32(0.1) / 3.14159)
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 uint8 = v1
        _ = fi21
        _ = v22
    }
    v23 := GenS3{F0: (42 >= 2147483647), F1: 305}
    v24 := GenS3{F0: true, F1: -9223372036854775808}
    var v25 int8 = 1
    if (!(!true)) {
        if (!false) {
            switch v2 {
            case 32768:
                if true {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        if false {
                            var v27 int = -754
                            _ = v27
                        } else {
                            v6 += 2
                            var v28 int64 = 0
                            _ = v28
                        }
                        var v29 int = (v2 % ((-26) | 1))
                        for fi30 := 0; fi30 < 6; fi30++ {
                            v6 = (v6 | v6)
                            delete(v3, 1)
                            if (-1 > -2147483648) {
                                v31 := GenS1{F0: "", F1: 0, F2: 4294967295}
                                switch "a" {
                                case "a":
                                    switch fi30 {
                                    case 127:
                                        if false {
                                            v32 := GenS1{F0: "foo bar", F1: -1, F2: 345}
                                            var v33 uint16 = 12
                                            if (!false) {
                                                if len(v4) > 0 {
                                                    v4[0] = 4294967295
                                                }
                                                v1 -= (v1 / ((255) | 1))
                                                if (!true) {
                                                    var v34 uint32 = 2147483648
                                                    _ = v34
                                                } else {
                                                    _ = len(v5)
                                                    v29 = fi26
                                                }
                                            }
                                            _ = v32
                                            _ = v33
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = (-(v2))
                                            }
                                        }
                                        fmt.Println(v2, v2)
                                        fmt.Println(v1, v6)
                                        if true {
                                            v35 := GenS4{F0: v6, F1: (v25 << ((18446744073709551615) & 63)), F2: (v25 * v25)}
                                            _ = v35
                                        }
                                    default:
                                        if len(v5) > 0 {
                                            v5[0] = v29
                                        }
                                    }
                                case "hello":
                                    var v36 uint32 = 1
                                    _ = v36
                                default:
                                    fmt.Println(v2)
                                    var v37 uint16 = 2
                                    _ = v37
                                }
                                v38 := GenS2{F0: v25}
                                _ = v31
                                _ = v38
                            } else {
                                switch 128 {
                                case 1:
                                    var v39 float64 = float64(1e10)
                                    _ = v39
                                case -9223372036854775808:
                                    if false {
                                        switch "a" {
                                        case "hello":
                                            var v40 float32 = float32(((1.5 + float64((0.1 - 1e10))) + 0.1))
                                            _ = v40
                                        case "":
                                            var v41 bool = true
                                            _ = v41
                                        case "a":
                                            v42 := GenS4{F0: 4294967295, F1: (v25 | v25), F2: v25}
                                            _ = v42
                                        default:
                                            var v43 uint = 42
                                            _ = v43
                                        }
                                    }
                                    v29 += 4294967295
                                default:
                                    v44 := GenS5{F0: v29}
                                    v45 := GenS4{F0: v6, F1: int8(v1), F2: (v25 << ((994) & 63))}
                                    _ = v44
                                    _ = v45
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v46 int64 = -485
                                fmt.Println(v46)
                                var v47 int8 = v25
                                _ = v46
                                _ = v47
                            }
                            _ = fi30
                        }
                        _ = fi26
                        _ = v29
                    }
                    v48 := GenS2{F0: 100}
                    v49 := GenS6{F0: v25}
                    _ = v48
                    _ = v49
                }
            case 42:
                if (!false) {
                    delete(v3, 673)
                }
            case 4294967295:
                for fi50 := 0; fi50 < 6; fi50++ {
                    delete(v3, 85)
                    for fi51 := 0; fi51 < 5; fi51++ {
                        var v52 uint32 = 2147483648
                        _ = fi51
                        _ = v52
                    }
                    _ = fi50
                }
            }
        }
        if (0 < genFn1()) {
            var v53 int32 = -183
            _ = v53
        } else {
            switch "" {
            case "foo bar":
                for fi54 := 0; fi54 < 2; fi54++ {
                    var v55 int64 = 42
                    _ = fi54
                    _ = v55
                }
            case "":
                var v56 int32 = 338
                _ = v56
            default:
                v57 := GenS1{F0: "hello", F1: -1, F2: 4294967295}
                _ = v57
            }
        }
    } else {
        var v58 int = (v2 >> ((816) & 63))
        var v59 int = -538
        for fi60 := 0; fi60 < 3; fi60++ {
            switch v59 {
            case -9223372036854775808:
                fmt.Println(v6)
            case -505:
                var v61 float32 = float32((1.5 / (1.5 * 1e10)))
                _ = v61
            }
            v4 = append(v4, 2)
            v25 += v25
            _ = fi60
        }
        _ = v58
        _ = v59
    }
    switch "hello" {
    case "":
        v62 := GenS7{F0: 164, F1: 1}
        _ = v62
    default:
        switch (v2 + v2) {
        case 127:
            if true {
                var v63 float32 = (3.14159 - (float32(0.0) - 0.0))
                for fi64 := 0; fi64 < 16; fi64++ {
                    v65 := GenS6{F0: 42}
                    v3[65535] = 2147483648
                    fmt.Println(v25, v2, v1)
                    _ = fi64
                    _ = v65
                }
                _ = v63
            }
        case 1:
            v66 := GenS8{F0: (v25 - v25), F1: 1}
            _ = v66
        case 165:
            for fi67 := 0; fi67 < 14; fi67++ {
                v6 = v6
                if true {
                    v68 := GenS1{F0: "a", F1: -2147483648, F2: 468}
                    for fi69 := 0; fi69 < 4; fi69++ {
                        v6 += 42
                        _ = fi69
                    }
                    var v70 int8 = 100
                    _ = v68
                    _ = v70
                }
                v25 += 41
                _ = fi67
            }
        }
    }
    var v71 string = ""
    var v72 float32 = (float32(-2.25) / 3.14159)
    for fi73 := 0; fi73 < 7; fi73++ {
        fmt.Println(v25)
        v74 := GenS4{F0: v6, F1: v25, F2: v25}
        _ = fi73
        _ = v74
    }
    fmt.Println(v71, v1)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v72
}
