package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int; F1 string; F2 int64 }
func genFn0(p0 uint32) int64 {
    v5 := []int{9223372036854775807}
    p0 = p0
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 int32) uint {
    v6 := GenS1{F0: 0, F1: 127}
    v7 := GenS2{F0: 236, F1: ("foo bar" + "\n\t"), F2: genFn0(2147483648)}
    v8 := []int{-1, 908, -982}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 995
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]int)
    var v3 float32 = 3.14159
    var v4 int64 = 263
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(v1)
        var v10 int64 = (-(v4))
        var v11 int64 = (v10 - v10)
        _ = fi9
        _ = v10
        _ = v11
    }
    switch 157 {
    case -39:
        switch (("\n\t" + ("foo bar" + ("\n\t" + "hello"))) + "a") {
        case "hello":
            switch "" {
            case "\n\t":
                v3 += (float32((-2.25 - float64(float32(float64((((0.0 * (((3.14159 / v3) + (0.1 + float32(1e10))) + v3)) + v3) - v3)))))) + float32(3.14159))
            }
        default:
            v12 := GenS1{F0: -128, F1: 1}
            _ = v12
        }
    }
    var v13 int16 = 32767
    _ = len(v2)
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint = 176
            v4 -= (v4 << (((v16 % ((18446744073709551615) | 1))) & 63))
            _ = fi15
            _ = v16
        }
        switch (fi14 | (-(fi14))) {
        case -9223372036854775808:
            v13 -= (v13 % (((v13 % (((v13 * 1000)) | 1))) | 1))
        default:
            v17 := []int{fi14, 659, fi14}
            _ = v17
        }
        switch -9223372036854775808 {
        case 541:
            var v18 int32 = 1
            _ = v18
        }
        _ = fi14
    }
    var v19 float64 = (0.1 / (3.14159 + 1e10))
    var v20 int = 633
    var v21 bool = true
    var v22 int64 = genFn0(2147483647)
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 int16 = (v13 + 1000)
        switch (v20 & fi23) {
        case 486:
            var v25 string = ("" + "hello")
            _ = v25
        default:
            fmt.Println(fi23, v21, v21)
        }
        _ = fi23
        _ = v24
    }
    var v26 float64 = 3.14159
    for fi27 := 0; fi27 < 2; fi27++ {
        var v28 int = (-(v20))
        v29 := GenS1{F0: 100, F1: 255}
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int32 = 0
    if (v19 != (v19 - (v26 / v26))) {
        _ = len(v2)
    }
    if v21 {
        v31 := []int{}
        var v32 bool = false
        v20 = int(v22)
        _ = v31
        _ = v32
    }
    delete(v2, v1)
    var v33 int = (-(v20))
    for fi34 := 0; fi34 < 7; fi34++ {
        _ = len(v2)
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v30
    _ = v33
    _ = v4
}
