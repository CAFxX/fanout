package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int; F2 uint8 }
func genFn0(p0 uint8) int16 {
    fmt.Println(p0)
    v6 := GenS1{F0: 1}
    v7 := GenS2{F0: 127, F1: 0, F2: 12}
    var v8 uint64 = 1
    v9 := []int{32768, 32767, 949}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []int{-47, 127, 2147483647}
    var v2 int32 = 792
    v3 := make(map[int]string)
    var v4 float32 = ((0.1 + 0.0) - 0.1)
    v5 := make(map[string]int)
    fmt.Println(v2, v4)
    for fi10 := 0; fi10 < 8; fi10++ {
        switch ("\n\t" + ("hello" + "\n\t")) {
        case "a":
            fmt.Println(v2, v4)
        case "\n\t":
            fmt.Println(v4, v2)
        default:
            fmt.Println(fi10, v4, v4)
        }
        var v11 int16 = -1
        fmt.Println(v2)
        _ = fi10
        _ = v11
    }
    v5["\n\t"] = -2147483648
    fmt.Println(v4, v4)
    v12 := GenS2{F0: 9223372036854775807, F1: -9223372036854775808, F2: 255}
    fmt.Println(v2, v2)
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS1{F0: -1}
        v2 -= v2
        v2 -= (v2 << ((887) & 63))
        _ = fi13
        _ = v14
    }
    v4 = v4
    var v15 uint = 2
    for fi16 := 0; fi16 < 11; fi16++ {
        v15 = v15
        _ = fi16
    }
    var v17 float32 = float32(((0.0 + (float64(float32(0.0)) + ((0.1 / (3.14159 / 1.5)) + 3.14159))) + 1e10))
    v18 := GenS2{F0: 32767, F1: 32767, F2: 1}
    if len(v1) > 0 {
        v1[0] = -1
    }
    fmt.Println(v2, v17, v2)
    var v19 bool = false
    v20 := GenS2{F0: int(v15), F1: 647, F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
