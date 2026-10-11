package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
func genFn0(p0 int) int16 {
    var v7 uint = 4294967295
    v7 += (v7 >> (((v7 >> ((88) & 63))) & 63))
    _ = p0
    _ = v7
    return -1
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]int)
    var v3 uint32 = 768
    var v4 float32 = -2.25
    v5 := []uint32{(v3 + v3), v3, 2147483647, v3}
    var v6 uint32 = v3
    if true {
        var v8 uint64 = 42
        _ = v8
    }
    v3 = (v3 % (((v6 * v6)) | 1))
    v1 -= (v1 + (v1 - v1))
    var v9 float32 = (float32(0.0) / (((((((3.14159 - -2.25) + v4) + ((0.1 - 0.1) + 1.5)) + -2.25) / 0.1) / v4) / v4))
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 float64 = ((3.14159 - (((-2.25 + 1.5) * float64((1.5 / (1e10 / (v9 + v9))))) * float64(1e10))) * 1e10)
        v12 := GenS1{F0: 0, F1: 0, F2: "hello"}
        if false {
            var v13 uint32 = v6
            var v14 string = ("a" + "foo bar")
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = len(v2)
    var v15 uint64 = 4294967295
    v16 := GenS1{F0: 42, F1: 1, F2: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
