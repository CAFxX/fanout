package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool, p1 int8, p2 int32) uint32 {
    v5 := GenS1{F0: 0}
    var v6 string = "foo bar"
    var v7 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 float32 = (float32(((float64(3.14159) - float64(float32(0.1))) + (float64(float32(((-2.25 - 1.5) * 0.0))) * 1e10))) + (float32(1e10) * ((0.0 / ((((-2.25 - 1e10) * (((1e10 + 1.5) * (3.14159 - 0.1)) - float32(0.0))) - 0.0) - float32(0.0))) * float32(1.5))))
    var v2 int32 = -1
    var v3 int = 9223372036854775807
    v4 := make(map[uint32]bool)
    v8 := GenS2{F0: 76, F1: 133, F2: 0}
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint8 = 79
        _ = fi9
        _ = v10
    }
    v1 = float32(3.14159)
    var v11 int16 = -947
    var v12 int16 = (v11 & int16(v3))
    var v13 uint64 = 2
    v4[2147483648] = ((!false) && ((true || false) && false))
    v14 := []int{(v3 | (v3 + v3)), v3}
    v15 := GenS3{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
