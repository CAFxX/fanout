package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 int16; F2 int64 }
type GenS5 struct { F0 bool; F1 int8; F2 int32 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int16) uint {
    if true {
        var v7 float64 = 1e10
        for fi8 := 0; fi8 < 11; fi8++ {
            switch "hello" {
            case "foo bar":
                v9 := []int{32767, fi8}
                _ = v9
            default:
                var v10 uint8 = 42
                _ = v10
            }
            switch "hello" {
            case "hello":
                switch "a" {
                case "hello":
                    if (4294967295 <= 2) {
                        var v11 uint32 = 2
                        fmt.Println(fi8, p0, v11)
                        var v12 int8 = -87
                        _ = v11
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 6; fi13++ {
                            v14 := GenS1{F0: -333, F1: 127, F2: 0}
                            switch ("foo bar" + "a") {
                            case "hello":
                                v15 := []int{32767, (^(fi8)), fi13}
                                _ = v15
                            case "\n\t":
                                switch -2147483648 {
                                case 2147483647:
                                    var v16 float64 = v7
                                    _ = v16
                                default:
                                    p0 = p0
                                }
                            case "":
                                var v17 uint32 = 1
                                _ = v17
                            default:
                                if true {
                                    v18 := []int{fi8, 32768, fi13, 0}
                                    _ = v18
                                }
                            }
                            v7 += (((((3.14159 + v7) * v7) - v7) + float64(1.5)) * (v7 / v7))
                            _ = fi13
                            _ = v14
                        }
                    }
                case "":
                    v7 += v7
                default:
                    var v19 int64 = 702
                    _ = v19
                }
            }
            v20 := GenS2{F0: -599, F1: 9223372036854775807}
            _ = fi8
            _ = v20
        }
        _ = v7
    }
    v21 := GenS1{F0: 0, F1: 0, F2: 4294967295}
    p0 = p0
    v22 := []int{-9223372036854775808, -689, -411, 42}
    _ = p0
    _ = v21
    _ = v22
    return 692
}

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 814
    var v3 int32 = -205
    var v4 float64 = -2.25
    v5 := make(map[uint32]bool)
    var v6 int16 = 1000
    v6 += 42
    v3 = 2147483647
    v2 -= v2
    var v23 int8 = -128
    if (!false) {
        var v24 bool = false
        for fi25 := 0; fi25 < 19; fi25++ {
            if (1 != 1) {
                switch 1 {
                case 713:
                    for fi26 := 0; fi26 < 12; fi26++ {
                        var v27 int = (fi25 - 812)
                        var v28 uint16 = 42
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                case 9223372036854775807:
                    var v29 int16 = (v6 % (((v6 | (^(v6)))) | 1))
                    _ = v29
                case -2147483648:
                    v30 := []int{(fi25 + fi25), 0, 32768, fi25}
                    _ = v30
                default:
                    fmt.Println(v6, v24, v6)
                }
                for fi31 := 0; fi31 < 1; fi31++ {
                    fmt.Println(v1, v23)
                    v32 := GenS3{F0: v1}
                    if v24 {
                        if v24 {
                            var v33 float64 = v4
                            fmt.Println(v3, v33, v23)
                            _ = v33
                        } else {
                            v34 := GenS4{F0: 2, F1: int16(fi31), F2: -240}
                            delete(v5, 426)
                            _ = v34
                        }
                    }
                    _ = fi31
                    _ = v32
                }
                for fi35 := 0; fi35 < 9; fi35++ {
                    switch "" {
                    case "hello":
                        v36 := []int{int(v23), 42}
                        _ = v36
                    case "foo bar":
                        v37 := GenS5{F0: (0 < 1), F1: v23, F2: (-(v3))}
                        _ = v37
                    }
                    v6 = -1
                    _ = fi35
                }
            } else {
                fmt.Println(v4)
            }
            _ = fi25
        }
        var v38 uint16 = 1
        _ = v24
        _ = v38
    } else {
        var v39 int8 = -110
        var v40 uint64 = v2
        _ = v39
        _ = v40
    }
    v41 := GenS2{F0: -1, F1: -238}
    v42 := GenS6{F0: v2}
    var v43 uint8 = (v1 % ((v1) | 1))
    v44 := GenS4{F0: v2, F1: -1, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
