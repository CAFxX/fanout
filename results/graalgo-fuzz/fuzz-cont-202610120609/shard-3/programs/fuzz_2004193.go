package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }
func genFn0(p0 bool) string {
    p0 = ((1000 == 42) || ((451 > 1) && (false || ((109 != 42) || p0))))
    if p0 {
        fmt.Println(p0)
    }
    v5 := GenS1{F0: 42, F1: 2, F2: 182}
    var v6 int16 = 32767
    _ = p0
    _ = v5
    _ = v6
    return "foo bar"
}

func main() {
    var v1 float32 = float32(-2.25)
    v2 := make(map[int]bool)
    var v3 int = 1
    var v4 int16 = 183
    var v7 int8 = -44
    v3 -= v3
    v7 += int8(v3)
    v4 = v4
    if false {
        v2[v3] = (!((v3 - v3) > 128))
        var v8 int16 = v4
        v7 = v7
        _ = v8
    } else {
        v1 = float32(float64(((v1 + 3.14159) / 0.1)))
        var v9 uint32 = 172
        _ = v9
    }
    v7 = v7
    v10 := GenS1{F0: 21, F1: 2, F2: 255}
    var v11 int = v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
