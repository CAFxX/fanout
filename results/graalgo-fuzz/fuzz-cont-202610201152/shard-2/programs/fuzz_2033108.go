package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]uint32)
    v3 := []string{}
    var v4 uint8 = 128
    for fi5 := 0; fi5 < 20; fi5++ {
        v4 = 127
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int8 = 42
            _ = fi6
            _ = v7
        }
        if (false || (true && false)) {
            _ = len(v3)
            fmt.Println(v4, v4, fi5)
        }
        _ = fi5
    }
    var v8 int = -317
    switch (v8 << (((v1 % ((699) | 1))) & 63)) {
    case 32768:
        v9 := GenS1{F0: (!(!(((v1 & (v1 << ((v1) & 63))) == (v1 / ((4294967295) | 1))) || ((true && false) || true)))), F1: 371}
        _ = v9
    default:
        for fi10 := 0; fi10 < 10; fi10++ {
            v1 += 65535
            _ = fi10
        }
    }
    fmt.Println(v4, v4)
    var v11 float32 = float32(((3.14159 / ((((-2.25 - (float64(float32(((-2.25 / 1.5) * float64(3.14159)))) / (1.5 - (1.5 * float64((3.14159 + (-2.25 / 1.0))))))) * 0.0) / float64(0.1)) - ((1e10 + 0.0) * 3.14159))) + 1e10))
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int32 = -2147483648
        var v14 float32 = v11
        var v15 bool = (-915 != 32767)
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v11 = -2.25
    v1 = 42
    v16 := GenS1{F0: (1.5 <= -2.25), F1: 2}
    _ = len(v2)
    if true {
        delete(v2, v8)
        var v17 uint32 = 987
        var v18 float64 = -2.25
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        v1 += v1
        v8 += (v8 - fi19)
        _ = fi19
    }
    v20 := GenS1{F0: (-2.25 >= -2.25), F1: 2}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v11 += (v11 * v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
