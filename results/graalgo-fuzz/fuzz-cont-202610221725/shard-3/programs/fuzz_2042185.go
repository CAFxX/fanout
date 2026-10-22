package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int8 }
func genFn0(p0 int) int64 {
    v6 := []int{(-(p0)), (^(p0)), (p0 - p0), p0}
    var v7 uint = 0
    for fi8 := 0; fi8 < 13; fi8++ {
        v7 -= 0
        if true {
            if len(v6) > 0 {
                v6[0] = 9223372036854775807
            }
            fmt.Println(v7, v7)
        }
        fmt.Println(p0, v7, v7)
        _ = fi8
    }
    _ = p0
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{1, 2}
    v3 := []uint32{197, 1, 717, 711}
    var v4 uint = 4294967295
    v5 := make(map[uint32]int)
    fmt.Println(v4)
    if false {
        fmt.Println(v4)
    } else {
        var v9 int16 = -232
        v9 = v9
        _ = v9
    }
    var v10 int32 = -549
    if (!(1 > 1)) {
        v4 -= v4
    } else {
        fmt.Println(v4, v10)
        v10 -= -157
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v10 -= 1
        v4 = (^(v4))
        _ = fi11
    }
    var v12 float32 = -2.25
    if true {
        if true {
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := GenS1{F0: (v10 << (((v4 * v4)) & 63)), F1: 0, F2: (fi13 & fi13)}
                _ = fi13
                _ = v14
            }
            v12 -= (v12 / v12)
        }
        v12 = v12
        var v15 int8 = 1
        _ = v15
    }
    if false {
        var v16 int16 = 1000
        var v17 int16 = v16
        _ = v16
        _ = v17
    }
    if false {
        v2 = append(v2, 2147483647)
        v12 += v12
        v12 = (1.5 / 0.1)
    } else {
        v18 := GenS1{F0: (v10 / ((v10) | 1)), F1: 42, F2: 42}
        switch 32767 {
        case -1:
            if true {
                v19 := GenS1{F0: v10, F1: 255, F2: -272}
                v12 -= float32(1.5)
                _ = v19
            }
        case 32768:
            v20 := v5[2147483648]
            _ = v20
        case -2147483648:
            v12 += v12
        }
        v10 += (^(v10))
        _ = v18
    }
    switch "\n\t" {
    case "\n\t":
        var v21 int64 = -430
        _ = v21
    default:
        fmt.Println(v12, v4, v4)
    }
    if false {
        v4 -= v4
        var v22 string = "hello"
        for fi23 := 0; fi23 < 16; fi23++ {
            v4 -= 42
            fmt.Println(v4, v4)
            _ = fi23
        }
        _ = v22
    }
    var v24 int8 = 1
    var v25 float32 = (float32(1e10) - ((v12 * (v12 * (v12 * float32((0.1 / 1e10))))) / 0.1))
    v26 := GenS1{F0: v10, F1: 42, F2: 128}
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := GenS2{F0: 65535}
        var v29 float32 = v12
        _ = fi27
        _ = v28
        _ = v29
    }
    if true {
        if false {
            for fi30 := 0; fi30 < 15; fi30++ {
                switch "" {
                case "":
                    _ = len(v5)
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                var v31 bool = ((fi30 * (-(fi30))) < (fi30 % (((fi30 / (((fi30 | (-(fi30)))) | 1))) | 1)))
                v32 := GenS2{F0: 1}
                _ = fi30
                _ = v31
                _ = v32
            }
        } else {
            v10 = 2147483647
            if true {
                if (-134 <= genFn0(1)) {
                    if false {
                        v12 += (1e10 + 0.0)
                        if true {
                            var v33 uint8 = 127
                            var v34 int32 = v10
                            fmt.Println(v10, v34)
                            _ = v33
                            _ = v34
                        } else {
                            if (true && true) {
                                var v35 int64 = genFn0(32767)
                                switch 2147483647 {
                                case -603:
                                    var v36 uint32 = 445
                                    _ = v36
                                case 421:
                                    v37 := GenS1{F0: v10, F1: 645, F2: 4294967295}
                                    _ = v37
                                }
                                for fi38 := 0; fi38 < 7; fi38++ {
                                    v39 := GenS3{F0: v4, F1: -471, F2: 669}
                                    switch "hello" {
                                    case "a":
                                        var v40 int = fi38
                                        _ = v40
                                    case "":
                                        var v41 int8 = v24
                                        _ = v41
                                    }
                                    var v42 int16 = 70
                                    _ = fi38
                                    _ = v39
                                    _ = v42
                                }
                                _ = v35
                            }
                            v10 = (v10 | (v10 + (v10 / ((2147483647) | 1))))
                        }
                    } else {
                        v25 -= float32(-2.25)
                        v12 -= v25
                    }
                    _ = len(v1)
                }
                v43 := GenS4{F0: -943}
                _ = v43
            } else {
                var v44 uint32 = 159
                _ = v44
            }
            v12 += v25
        }
    }
    _ = len(v1)
    v45 := GenS5{F0: 215, F1: v24}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
