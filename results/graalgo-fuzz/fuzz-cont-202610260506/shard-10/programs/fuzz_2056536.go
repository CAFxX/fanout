package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int8; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 int) int16 {
    switch p0 {
    case -1:
        v7 := []int{(p0 & (p0 & p0)), p0, -9223372036854775808}
        _ = v7
    case 127:
        var v8 float64 = float64(((0.1 - float32(0.0)) - (1.5 - -2.25)))
        _ = v8
    case -9223372036854775808:
        switch 4294967295 {
        case -9223372036854775808:
            p0 = p0
        case 127:
            fmt.Println(p0)
        }
    }
    p0 += (p0 ^ p0)
    if false {
        fmt.Println(p0)
    }
    v9 := []int{p0, (p0 & 631)}
    _ = p0
    _ = v9
    return -32768
}
func genFn1(p0 uint8) uint16 {
    var v10 bool = true
    for fi11 := 0; fi11 < 14; fi11++ {
        if v10 {
            v10 = true
            fmt.Println(v10, fi11, v10)
            if false {
                v12 := GenS1{F0: 2}
                v13 := GenS1{F0: uint64(p0)}
                p0 -= 127
                _ = v12
                _ = v13
            } else {
                v14 := GenS2{F0: -41, F1: 92, F2: 42}
                _ = v14
            }
        }
        switch fi11 {
        case 693:
            v10 = (v10 || v10)
        case 127:
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 float32 = ((((0.1 + float32((0.1 * float64((0.0 * (1e10 * 1e10)))))) + -2.25) / 1.5) * float32(0.0))
                _ = fi15
                _ = v16
            }
        case 32767:
            v17 := GenS3{F0: true}
            _ = v17
        }
        _ = fi11
    }
    v10 = (((float32(-2.25) * 0.1) / 1.0) >= 1e10)
    if true {
        v10 = false
    }
    _ = p0
    _ = v10
    return 131
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 766
    v3 := make(map[int]uint32)
    v4 := []int{}
    var v5 uint8 = 255
    v6 := []int{-2147483648, 9223372036854775807, -9223372036854775808, 93}
    v1[217] = -2147483648
    v5 = 128
    v5 -= v5
    for fi18 := 0; fi18 < 20; fi18++ {
        v5 = v5
        _ = fi18
    }
    v2 = 18446744073709551615
    var v19 int32 = -794
    fmt.Println(v19)
    fmt.Println(v5)
    delete(v1, 2147483647)
    if false {
        var v20 int16 = 1
        _ = v20
    } else {
        var v21 uint8 = v5
        _ = v21
    }
    v19 = v19
    for fi22 := 0; fi22 < 9; fi22++ {
        v5 -= 1
        _ = fi22
    }
    switch (((("foo bar" + "foo bar") + "hello") + "foo bar") + "hello") {
    case "foo bar":
        _ = len(v1)
    case "a":
        for fi23 := 0; fi23 < 7; fi23++ {
            for fi24 := 0; fi24 < 14; fi24++ {
                if (2 < 708) {
                    v3[fi24] = 518
                    v25 := GenS4{F0: -666, F1: false}
                    _ = v25
                }
                _ = fi24
            }
            fmt.Println(v19, fi23, v2)
            _ = fi23
        }
    default:
        if ((1e10 + ((3.14159 - 1.5) / (0.1 - ((float32(-2.25) - (-2.25 / -2.25)) + 0.0)))) != 0.1) {
            v26 := GenS3{F0: true}
            fmt.Println(v2, v2)
            _ = v26
        } else {
            var v27 uint8 = v5
            var v28 int16 = -470
            v5 += v5
            _ = v27
            _ = v28
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
