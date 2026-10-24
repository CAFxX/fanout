package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int32 }
type GenS2 struct { F0 uint; F1 int16; F2 uint64 }

func main() {
    var v1 string = ("hello" + "a")
    v2 := make(map[int]uint32)
    v3 := []string{}
    var v4 float64 = 3.14159
    var v5 bool = false
    if v5 {
        var v6 uint8 = 42
        v5 = v5
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: -1, F2: 1}
    v5 = (42 < -516)
    for fi8 := 0; fi8 < 17; fi8++ {
        switch (fi8 + fi8) {
        case 32768:
            v4 = v4
        case 330:
            v1 = v1
        case 2147483647:
            v9 := GenS1{F0: -128, F1: -1, F2: -1}
            _ = v9
        }
        v10 := GenS1{F0: 94, F1: -1, F2: -2147483648}
        v11 := GenS2{F0: 886, F1: 872, F2: 898}
        _ = fi8
        _ = v10
        _ = v11
    }
    v4 += (float64(float32(v4)) * v4)
    for fi12 := 0; fi12 < 13; fi12++ {
        v5 = v5
        _ = fi12
    }
    var v13 uint16 = 204
    var v14 uint = 4294967295
    fmt.Println(v5)
    var v15 float32 = float32((1e10 + float64(3.14159)))
    fmt.Println(v14)
    var v16 int64 = -1
    var v17 int32 = 2147483647
    if v5 {
        fmt.Println(v5)
    }
    switch v1 {
    case "":
        var v18 float32 = 1.5
        _ = v18
    case "hello":
        v2[1] = 2147483647
    }
    var v19 int8 = -53
    var v20 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
