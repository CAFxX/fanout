package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 int32; F1 int16; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 bool }
func genFn0(p0 string) int16 {
    var v7 uint64 = 2
    p0 = "foo bar"
    fmt.Println(v7, v7)
    var v8 int32 = 320
    if (false && true) {
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: v7, F1: (fi9 * fi9), F2: 183}
            fmt.Println(v7, v7)
            _ = fi9
            _ = v10
        }
        var v11 bool = false
        v12 := GenS1{F0: (v7 ^ v7), F1: 127, F2: 2}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v7
    _ = v8
    return int16(v7)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    v3 := []uint32{0, 2147483648, 186}
    v4 := []int{2147483647, -751, -2147483648, 802}
    v5 := []uint32{0}
    var v6 uint = 1
    if true {
        fmt.Println(v6)
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            if len(v3) > 0 {
                v3[0] = 42
            }
            v6 = v6
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi13
        }
    }
    fmt.Println(v6)
    v6 = (^(v6))
    fmt.Println(v6)
    for fi14 := 0; fi14 < 7; fi14++ {
        fmt.Println(fi14, v6)
        var v15 int64 = -9223372036854775808
        var v16 uint = (^(v6))
        _ = fi14
        _ = v15
        _ = v16
    }
    if (-624 >= 128) {
        v17 := GenS2{F0: -1, F1: 1000, F2: uint16(v6)}
        _ = v17
    }
    v18 := GenS1{F0: 922, F1: 32768, F2: uint8(v6)}
    v19 := GenS1{F0: 1, F1: 127, F2: 255}
    var v20 uint32 = 65535
    if true {
        _ = len(v4)
        if (true && true) {
            var v21 float32 = float32(float64(0.0))
            v4 = append(v4, 136)
            _ = v21
        } else {
            v22 := GenS1{F0: 2, F1: int(v20), F2: 0}
            _ = v22
        }
        v23 := GenS2{F0: 2147483647, F1: int16(v6), F2: 985}
        _ = v23
    }
    var v24 int = 32767
    var v25 bool = false
    switch v24 {
    case -539:
        fmt.Println(v25, v25, v25)
    case 128:
        fmt.Println(v6, v6)
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        if v25 {
            if (90 >= 42) {
                fmt.Println(v24)
                v20 = (v20 + (^(v20)))
            } else {
                switch "" {
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = (v20 * v20)
                    }
                case "":
                    switch 379 {
                    case 127:
                        var v27 uint32 = v20
                        _ = v27
                    case 4294967295:
                        v24 += (v24 / ((353) | 1))
                    default:
                        fmt.Println(v25, v6, fi26)
                    }
                case "foo bar":
                    _ = len(v3)
                }
                if (true && false) {
                    v5 = append(v5, (v20 % (((v20 % ((v20) | 1))) | 1)))
                    var v28 int16 = -873
                    _ = v28
                }
                v24 = (fi26 | (v24 << (((v6 - v6)) & 63)))
            }
        } else {
            var v29 uint = uint(v20)
            var v30 uint16 = 65535
            if v25 {
                switch "foo bar" {
                case "a":
                    fmt.Println(v29, v6, v6)
                default:
                    var v31 int = (v24 & fi26)
                    _ = v31
                }
                v2[(v20 * 768)] = (v30 == v30)
            } else {
                if v25 {
                    for fi32 := 0; fi32 < 6; fi32++ {
                        fmt.Println(v29, v29)
                        v29 += v29
                        _ = fi32
                    }
                    v33 := GenS2{F0: -740, F1: 739, F2: (v30 / ((723) | 1))}
                    _ = v33
                }
                if v25 {
                    v34 := GenS3{F0: (v24 * (fi26 % (((fi26 >> ((v6) & 63))) | 1)))}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if (v20 == v20) {
                        v20 = (v20 | v20)
                        var v35 int32 = int32(v29)
                        _ = v35
                    }
                    _ = v34
                } else {
                    v36 := v1[(fi26 >> ((42) & 63))]
                    _ = v36
                }
                v37 := v2[(v20 - v20)]
                _ = v37
            }
            _ = v29
            _ = v30
        }
        _ = fi26
    }
    v38 := GenS1{F0: 4294967295, F1: 290, F2: uint8(v6)}
    var v39 uint = 930
    switch -163 {
    case 32767:
        v6 -= v6
    case 4294967295:
        if v25 {
            for fi40 := 0; fi40 < 10; fi40++ {
                if v25 {
                    v41 := GenS2{F0: 770, F1: genFn0("\n\t"), F2: uint16(v39)}
                    var v42 int = v24
                    for fi43 := 0; fi43 < 2; fi43++ {
                        v42 = v24
                        _ = fi43
                    }
                    _ = v41
                    _ = v42
                } else {
                    switch (v24 + (v24 ^ (v24 >> (((v39 | v39)) & 63)))) {
                    case 127:
                        var v44 int64 = -9223372036854775808
                        _ = v44
                    case 407:
                        var v45 uint16 = 207
                        _ = v45
                    }
                    v46 := GenS4{F0: 42, F1: (!false)}
                    if v25 {
                        var v47 int8 = 114
                        v25 = (207 > 42)
                        fmt.Println(v24, v25, fi40)
                        _ = v47
                    }
                    _ = v46
                }
                for fi48 := 0; fi48 < 5; fi48++ {
                    v25 = v25
                    _ = fi48
                }
                _ = fi40
            }
        }
    case -367:
        for fi49 := 0; fi49 < 6; fi49++ {
            for fi50 := 0; fi50 < 5; fi50++ {
                var v51 float64 = 1e10
                var v52 int32 = 2147483647
                _ = len(v4)
                _ = fi50
                _ = v51
                _ = v52
            }
            fmt.Println(fi49, v25, fi49)
            _ = fi49
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
