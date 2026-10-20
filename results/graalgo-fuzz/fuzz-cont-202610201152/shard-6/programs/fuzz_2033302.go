package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]int)
    var v3 float64 = v1
    var v4 int16 = -1
    var v5 uint16 = 255
    var v6 int = -9223372036854775808
    v7 := []int{(-(v6)), v6, -56}
    var v8 int8 = -128
    for fi9 := 0; fi9 < 8; fi9++ {
        v7 = append(v7, 408)
        var v10 string = ""
        _ = fi9
        _ = v10
    }
    var v11 float64 = 1.5
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v12 float64 = 0.0
    v8 += (^(v8))
    v7 = append(v7, 128)
    v1 -= 0.0
    var v13 uint16 = 32767
    v14 := GenS1{F0: 42}
    v4 += (v4 ^ v4)
    var v15 int = (v6 << ((881) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
