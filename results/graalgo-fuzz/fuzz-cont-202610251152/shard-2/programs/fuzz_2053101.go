package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 string, p1 int8) int16 {
    v4 := GenS1{F0: true, F1: -9223372036854775808}
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 uint16 = 1
            _ = fi6
            _ = v7
        }
        var v8 string = p0
        _ = fi5
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int16 = -32768
        var v11 uint32 = 2147483648
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v4
    return -705
}
func genFn1(p0 int32) int64 {
    if false {
        p0 = -327
        var v12 int = 4294967295
        var v13 float32 = ((3.14159 * (3.14159 - 1e10)) * -2.25)
        _ = v12
        _ = v13
    } else {
        v14 := []int{128, -2147483648}
        p0 += p0
        v15 := GenS1{F0: true, F1: -9223372036854775808}
        _ = v14
        _ = v15
    }
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    var v1 float64 = (1e10 + 0.1)
    var v2 bool = true
    v3 := make(map[int]bool)
    v1 = 1e10
    fmt.Println(v1, v1)
    v16 := []int{-860, -75, 128, -64}
    v3[2147483647] = false
    var v17 int64 = 1
    if false {
        fmt.Println(v2, v17, v1)
        var v18 int8 = 127
        v19 := GenS2{F0: v18, F1: 68}
        _ = v18
        _ = v19
    }
    var v20 uint16 = 32767
    var v21 float32 = float32(float64((0.1 - (0.1 / -2.25))))
    fmt.Println(v2)
    v22 := GenS1{F0: v2, F1: v17}
    if (v21 > v21) {
        v21 = (v21 - v21)
        if ((0 < 2147483647) && false) {
            var v23 uint = 4294967295
            for fi24 := 0; fi24 < 11; fi24++ {
                v23 = v23
                v16 = append(v16, (fi24 ^ (fi24 | -1)))
                v25 := GenS3{F0: 254}
                _ = fi24
                _ = v25
            }
            _ = v23
        }
    } else {
        v26 := v3[307]
        var v27 string = "hello"
        var v28 uint8 = 176
        _ = v26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        _ = len(v16)
        var v30 int16 = 1000
        switch "foo bar" {
        case "hello":
            v31 := GenS3{F0: 128}
            _ = v31
        default:
            var v32 int8 = 1
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
