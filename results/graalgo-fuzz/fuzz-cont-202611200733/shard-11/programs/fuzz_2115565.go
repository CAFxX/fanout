package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{32768, 0, -195}
    var v2 int32 = 277
    v3 := []uint32{}
    var v4 string = ""
    var v5 float32 = (float32(3.14159) - float32(3.14159))
    v6 := GenS1{F0: 0, F1: 255, F2: 9223372036854775807}
    fmt.Println(v2)
    v7 := GenS2{F0: v2}
    _ = len(v3)
    var v8 int = 128
    v9 := GenS1{F0: 119, F1: 42, F2: 1}
    fmt.Println(v5, v8, v5)
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := GenS2{F0: int32(fi10)}
            var v13 bool = true
            _ = fi11
            _ = v12
            _ = v13
        }
        v4 = v4
        v5 = ((3.14159 * -2.25) * 0.1)
        _ = fi10
    }
    var v14 int = (v8 ^ v8)
    v15 := GenS1{F0: 1, F1: 437, F2: 9223372036854775807}
    var v16 int = v8
    _ = len(v3)
    v4 = (v4 + v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
