package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint8 = 143
    v2 := make(map[uint32]uint32)
    var v3 int = 895
    var v4 uint32 = 118
    var v5 int = (v3 >> ((710) & 63))
    var v6 int8 = 1
    var v7 int16 = 0
    var v8 int8 = -128
    v9 := []int{334, v5, v5, (v3 - 128)}
    v6 = 1
    v5 += (v5 / ((-9223372036854775808) | 1))
    switch "" {
    case "foo bar":
        var v10 uint64 = 2
        _ = v10
    }
    v5 += (-(v5))
    if (float32((-2.25 / float64(float32(1.5)))) != float32(1.5)) {
        var v11 float64 = 0.1
        v12 := GenS1{F0: 4294967295, F1: ("hello" + "foo bar"), F2: v6}
        v1 += (v1 >> ((0) & 63))
        _ = v11
        _ = v12
    }
    if false {
        var v13 uint32 = (v4 * (v4 % ((uint32(v1)) | 1)))
        var v14 uint8 = (v1 - v1)
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 787}
    var v16 uint64 = 4294967295
    _ = len(v9)
    v3 = v5
    v17 := GenS2{F0: v4}
    var v18 int16 = -160
    v19 := GenS1{F0: 255, F1: (("a" + "foo bar") + ("\n\t" + "")), F2: 16}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
