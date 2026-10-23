package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
func genFn0(p0 uint32, p1 int32) int16 {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{(fi6 ^ (fi6 | -34)), fi6, (fi6 << ((65535) & 63)), fi6}
        fmt.Println(p1)
        p0 -= p0
        _ = fi6
        _ = v7
    }
    p1 += p1
    var v8 uint16 = 1
    _ = p0
    _ = p1
    _ = v8
    return 203
}

func main() {
    var v1 int32 = -1
    var v2 int = -1
    var v3 int32 = v1
    var v4 float64 = 3.14159
    v5 := make(map[uint32]string)
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: 65535, F1: 32767, F2: 1}
        _ = fi9
        _ = v10
    }
    var v11 float32 = 1e10
    v12 := GenS1{F0: 1, F1: -1, F2: 682}
    v13 := GenS1{F0: 42, F1: 999, F2: 0}
    var v14 float64 = ((((float64(v11) - (v4 + (v4 * (v4 - v4)))) * v4) * float64(float32((v4 * v4)))) / float64(1e10))
    var v15 float32 = (v11 + ((3.14159 / (((float32(float64(v11)) + v11) / float32(float64(1e10))) * v11)) - float32(float64(3.14159))))
    v16 := GenS1{F0: 196, F1: 42, F2: 1}
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int64 = -410
        fmt.Println(v4, v1)
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
