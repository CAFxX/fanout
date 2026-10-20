package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 bool, p1 int) uint8 {
    switch "\n\t" {
    case "hello":
        var v4 uint8 = 0
        _ = v4
    case "\n\t":
        for fi5 := 0; fi5 < 12; fi5++ {
            fmt.Println(p1, fi5)
            var v6 string = ""
            _ = fi5
            _ = v6
        }
    }
    p0 = (0 == 210)
    v7 := []int{p1}
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint = 42
            _ = fi9
            _ = v10
        }
        fmt.Println(fi8)
        _ = fi8
    }
    v11 := GenS1{F0: 67}
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return 2
}
func genFn1(p0 uint64, p1 uint8) uint {
    var v12 int64 = 9223372036854775807
    v12 = 42
    _ = p0
    _ = p1
    _ = v12
    return 18446744073709551615
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    var v13 int8 = 42
    v14 := []int{-2147483648, -835, 783, 2147483647}
    for fi15 := 0; fi15 < 13; fi15++ {
        v1 -= (v1 & (v1 + v1))
        var v16 float32 = float32(0.0)
        if (genFn1(1, genFn0((1e10 < float64(1.5)), (fi15 * fi15))) > 4294967295) {
            v17 := GenS1{F0: 0}
            switch (fi15 >> ((4294967295) & 63)) {
            case 196:
                fmt.Println(v16)
            default:
                if (2147483648 != 228) {
                    var v18 float64 = ((1e10 / -2.25) + 1e10)
                    _ = v18
                } else {
                    var v19 int64 = v1
                    if (false || ((0.1 - 1.5) != (0.1 / 0.1))) {
                        fmt.Println(fi15, v1)
                        var v20 uint8 = 20
                        var v21 float64 = 1.5
                        _ = v20
                        _ = v21
                    } else {
                        var v22 float64 = -2.25
                        _ = v22
                    }
                    _ = v19
                }
            }
            v1 = v1
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v1)
    for fi23 := 0; fi23 < 13; fi23++ {
        if true {
            var v24 uint32 = 2
            _ = v24
        }
        _ = fi23
    }
    if (true && (193 != 11)) {
        v13 = (v13 * (^(v13)))
    }
    var v25 float64 = (float64((1.5 + 1e10)) - 0.1)
    var v26 uint64 = 0
    var v27 string = "hello"
    var v28 int16 = 0
    v14 = append(v14, -312)
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := v3[fi29]
        v25 = v25
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: 1}
    for fi32 := 0; fi32 < 12; fi32++ {
        if len(v14) > 0 {
            v14[0] = fi32
        }
        v2[((v27 + (v27 + "hello")) + (v27 + (((v27 + ((v27 + v27) + v27)) + (v27 + (v27 + (v27 + v27)))) + (v27 + "foo bar"))))] = v27
        _ = fi32
    }
    v14 = append(v14, -2147483648)
    v25 = v25
    v33 := GenS1{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v33
}
