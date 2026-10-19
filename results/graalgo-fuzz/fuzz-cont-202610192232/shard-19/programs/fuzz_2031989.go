package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint32; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 uint; F2 uint64 }
type GenS6 struct { F0 int32; F1 int }
func genFn0(p0 uint8, p1 int) int32 {
    v4 := GenS1{F0: true, F1: (p1 / ((0) | 1)), F2: -84}
    v5 := []int{4294967295, (-(p1)), (p1 & p1)}
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: false, F1: (-(p1)), F2: 42}
        _ = fi6
        _ = v7
    }
    if (!true) {
        v8 := GenS1{F0: false, F1: (-(p1)), F2: 127}
        var v9 uint32 = 4294967295
        switch "a" {
        case "":
            var v10 uint64 = 1
            _ = v10
        case "hello":
            if true {
                p1 = 128
                for fi11 := 0; fi11 < 8; fi11++ {
                    v12 := GenS2{F0: 151, F1: -9223372036854775808, F2: (p0 << ((259) & 63))}
                    _ = fi11
                    _ = v12
                }
            }
        case "a":
            v5 = append(v5, p1)
        default:
            if false {
                p1 -= -1
                for fi13 := 0; fi13 < 20; fi13++ {
                    if true {
                        var v14 float64 = 1.5
                        v9 -= v9
                        v14 = 0.1
                        _ = v14
                    } else {
                        var v15 uint = 255
                        var v16 uint64 = 2
                        var v17 int16 = -114
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    fmt.Println(v9)
                    _ = len(v5)
                    _ = fi13
                }
                v9 -= 2147483647
            }
        }
        _ = v8
        _ = v9
    }
    if len(v5) > 0 {
        v5[0] = (p1 % ((32768) | 1))
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 86
}

func main() {
    v1 := []uint32{2, 2147483647}
    v2 := make(map[uint32]int)
    var v3 int32 = 2147483647
    v3 -= -1
    v3 = (v3 & (v3 | -236))
    _ = len(v1)
    for fi18 := 0; fi18 < 20; fi18++ {
        v3 -= (v3 | 0)
        v19 := GenS3{F0: 0, F1: 2, F2: -2147483648}
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 18446744073709551615
    var v21 float32 = (3.14159 * 1.5)
    v22 := GenS1{F0: true, F1: -92, F2: 1}
    var v23 float64 = (1e10 - (1.5 - (1e10 - (-2.25 + 1.5))))
    if (!(42 > 13)) {
        var v24 uint64 = 4294967295
        var v25 int32 = -954
        v26 := GenS4{F0: 128}
        _ = v24
        _ = v25
        _ = v26
    }
    if true {
        v27 := GenS5{F0: -1, F1: 42, F2: v20}
        for fi28 := 0; fi28 < 13; fi28++ {
            if len(v1) > 0 {
                v1[0] = 308
            }
            for fi29 := 0; fi29 < 15; fi29++ {
                v20 = 0
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi29
            }
            _ = fi28
        }
        _ = v27
    } else {
        if true {
            fmt.Println(v20)
            var v30 uint = 42
            fmt.Println(v21)
            _ = v30
        } else {
            v20 = 18446744073709551615
        }
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 uint = 764
            _ = fi31
            _ = v32
        }
        v33 := GenS1{F0: (false || (18446744073709551615 == 2)), F1: 127, F2: -61}
        _ = v33
    }
    if true {
        for fi34 := 0; fi34 < 15; fi34++ {
            var v35 int8 = -20
            var v36 bool = (v3 != -2147483648)
            _ = fi34
            _ = v35
            _ = v36
        }
        var v37 uint64 = 999
        switch 2147483647 {
        case 785:
            _ = len(v2)
        case 4294967295:
            fmt.Println(v3)
        case 490:
            for fi38 := 0; fi38 < 12; fi38++ {
                var v39 int16 = -32768
                v21 += v21
                _ = fi38
                _ = v39
            }
        }
        _ = v37
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        switch fi40 {
        case 284:
            switch "hello" {
            case "a":
                v41 := GenS5{F0: 42, F1: 404, F2: v20}
                _ = v41
            case "":
                v20 = (v20 >> ((53) & 63))
            }
        case -244:
            var v42 int32 = (v3 << ((18446744073709551615) & 63))
            _ = v42
        case 2147483647:
            v3 = v3
        default:
            fmt.Println(v20, v23, v23)
        }
        switch "hello" {
        case "hello":
            v43 := GenS1{F0: (false || (-698 >= 514)), F1: (fi40 << ((306) & 63)), F2: 4}
            _ = v43
        case "":
            var v44 float64 = float64((float32(1.5) + v21))
            _ = v44
        case "a":
            _ = len(v1)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi40
    }
    v45 := GenS2{F0: 121, F1: -196, F2: 194}
    fmt.Println(v21, v21, v3)
    v46 := GenS6{F0: (v3 + genFn0(97, 128)), F1: -9223372036854775808}
    v3 += (v3 ^ v3)
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v45
    _ = v46
}
