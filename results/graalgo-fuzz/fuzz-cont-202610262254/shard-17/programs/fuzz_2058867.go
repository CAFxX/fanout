package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int8; F2 uint }
func genFn0(p0 int, p1 int, p2 string) bool {
    var v5 int16 = -32768
    if false {
        fmt.Println(v5)
    }
    p0 += p0
    v6 := []int{(p0 / (((p1 & p1)) | 1)), -9223372036854775808, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ((("foo bar" + p2) + ((p2 + p2) + ("" + "a"))) == p2)
}
func genFn1(p0 uint64, p1 int16, p2 string) int16 {
    p2 = p2
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint8 = 88
    v2 := make(map[int]string)
    var v3 float64 = (1.5 * -2.25)
    v4 := make(map[int]int)
    if (v3 > v3) {
        fmt.Println(v3, v1)
        switch "a" {
        case "foo bar":
            fmt.Println(v1, v1)
        case "a":
            v3 = 3.14159
        }
        fmt.Println(v3, v3)
    }
    fmt.Println(v1, v3)
    var v7 int64 = 42
    v8 := []int{-1, -493, 2147483647}
    v9 := GenS1{F0: -41}
    var v10 int8 = -64
    if true {
        v11 := GenS2{F0: 494, F1: (v10 ^ (v10 ^ (v10 % (((^(v10))) | 1)))), F2: 255}
        v3 -= float64(-2.25)
        v12 := GenS1{F0: 42}
        _ = v11
        _ = v12
    }
    v7 = v7
    var v13 float64 = ((((float64(((-2.25 / float32((0.1 / (v3 + 3.14159)))) * 0.1)) * v3) - v3) + (-2.25 / (((((v3 + ((float64((0.1 / -2.25)) / float64(((1.5 / float32(v3)) - (1e10 * 0.1)))) + v3)) - v3) * float64(float32(1.5))) / v3) + (((1.5 * v3) * ((((1e10 / (v3 + v3)) / (3.14159 * -2.25)) - v3) + (v3 - v3))) / v3)))) + float64(0.1))
    fmt.Println(v10)
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            v7 = 0
            v10 = v10
            v10 += (v10 + v10)
            _ = fi15
        }
        var v16 uint64 = 2
        fmt.Println(v13, v13)
        _ = fi14
        _ = v16
    }
    v17 := GenS2{F0: 32767, F1: 79, F2: 42}
    v3 = 3.14159
    _ = len(v8)
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v13, v3)
        var v19 float32 = 0.0
        v19 = ((v19 * 3.14159) / 3.14159)
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
