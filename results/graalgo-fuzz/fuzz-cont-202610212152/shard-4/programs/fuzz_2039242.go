package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 bool; F1 int32; F2 int }
type GenS3 struct { F0 int16; F1 bool }
func genFn0(p0 uint8, p1 bool) uint16 {
    p0 += p0
    fmt.Println(p1, p0)
    v5 := []int{-2147483648, -397, -680, 128}
    _ = p0
    _ = p1
    _ = v5
    return 624
}
func genFn1(p0 uint16, p1 uint64) int {
    if true {
        fmt.Println(p1, p1)
        fmt.Println(p1, p0)
    }
    switch "a" {
    case "a":
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 int64 = -475
            for fi8 := 0; fi8 < 9; fi8++ {
                p0 += p0
                if false {
                    v9 := GenS1{F0: 885, F1: -50, F2: 18446744073709551615}
                    if (295 > 65535) {
                        switch (fi8 | fi6) {
                        case -2147483648:
                            var v10 string = "a"
                            _ = v10
                        case 2147483647:
                            p0 = genFn0(42, (-32768 < 32767))
                        case -1:
                            switch "hello" {
                            case "":
                                v11 := []int{127, fi6, (fi6 & int(p1)), -1}
                                _ = v11
                            }
                        default:
                            if (!false) {
                                v7 -= -695
                                p0 = 2
                            }
                        }
                        var v12 uint64 = p1
                        v13 := []int{fi8, (fi6 ^ (fi6 ^ fi6)), -2147483648}
                        p0 = (p0 >> ((187) & 63))
                        v13 = append(v13, fi8)
                        _ = v12
                        _ = v13
                    } else {
                        p0 = (p0 << ((255) & 63))
                        fmt.Println(v7)
                    }
                    p1 = (^(p1))
                    _ = v9
                } else {
                    if true {
                        v14 := []int{fi8, fi6}
                        v7 = v7
                        _ = v14
                    } else {
                        var v15 uint = 2
                        v16 := GenS1{F0: (v15 << ((v15) & 63)), F1: -128, F2: (v15 % ((v15) | 1))}
                        _ = v15
                        _ = v16
                    }
                    v7 = (v7 << ((42) & 63))
                    if (p0 >= p0) {
                        v17 := GenS1{F0: 1, F1: int8(v7), F2: 42}
                        var v18 string = "foo bar"
                        _ = v17
                        _ = v18
                    }
                }
                _ = fi8
            }
            v19 := GenS2{F0: true, F1: -618, F2: fi6}
            _ = fi6
            _ = v19
            _ = v7
        }
    case "\n\t":
        if true {
            fmt.Println(p0, p0)
        }
    case "":
        var v20 float64 = (-2.25 + (((0.1 + float64(0.0)) + (-2.25 * 1.5)) / ((0.0 / 0.1) + 1e10)))
        _ = v20
    default:
        switch 4294967295 {
        case 118:
            v21 := []int{724, 128}
            _ = v21
        default:
            v22 := []int{}
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 uint32 = 42
        p0 = p0
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 bool = false
        var v27 uint8 = uint8(p1)
        _ = fi25
        _ = v26
        _ = v27
    }
    _ = p0
    _ = p1
    return -253
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v4 int32 = 0
    v4 = 0
    v28 := GenS3{F0: -1, F1: (!(false && false))}
    var v29 uint8 = 211
    v4 += (v4 % (((^(v4))) | 1))
    v30 := GenS3{F0: 291, F1: false}
    v29 = v29
    var v31 bool = ((!((!(934 < 1)) || false)) && (-2.25 < (-2.25 + 3.14159)))
    v32 := []int{-994, -76, -1}
    var v33 uint16 = 2
    var v34 float64 = 1e10
    v35 := GenS1{F0: 0, F1: 127, F2: 52}
    v33 -= 1
    if false {
        var v36 int = 127
        _ = v36
    } else {
        for fi37 := 0; fi37 < 3; fi37++ {
            fmt.Println(v31, v4, fi37)
            fmt.Println(v31, v29, v4)
            switch 32767 {
            case 32768:
                var v38 int32 = v4
                _ = v38
            case 815:
                v31 = (v31 && true)
            case -15:
                var v39 uint64 = 42
                _ = v39
            default:
                v4 = 42
            }
            _ = fi37
        }
    }
    v33 = genFn0((v29 - v29), v31)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
