package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint64, p1 bool, p2 uint64) uint {
    v4 := GenS1{F0: 42, F1: 199}
    fmt.Println(p0, p1, p0)
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    p1 = p1
    var v6 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 255
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]string)
    var v3 float64 = (float64((1e10 + (v1 + 3.14159))) - float64((v1 - float32(1e10))))
    fmt.Println(v1, v1)
    v7 := []int{127}
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: 0, F1: 616}
            _ = fi9
            _ = v10
        }
        v11 := GenS2{F0: 2}
        v1 = v1
        _ = fi8
        _ = v11
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v12 float32 = float32(v3)
    v1 += float32((1.5 - float64(float32((v3 / 0.1)))))
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS2{F0: 49}
        v1 = (v12 + float32(((float64(3.14159) / v3) * (v3 / 1e10))))
        v1 += v12
        _ = fi13
        _ = v14
    }
    var v15 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v7
}
