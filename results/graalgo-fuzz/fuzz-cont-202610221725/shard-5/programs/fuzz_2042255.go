package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 0
    v3 := make(map[uint32]bool)
    v2 = v2
    var v4 int8 = 100
    fmt.Println(v4)
    if (2 >= 465) {
        delete(v1, 42)
        var v5 uint32 = 65535
        var v6 uint16 = 42
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(v2)
        var v8 int8 = -54
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int64 = -207
            v11 := []int{fi7, fi7}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v12 := GenS1{F0: 1}
    v1[458] = 982
    v3[2147483647] = false
    if (-729 >= 42) {
        var v13 int32 = -615
        _ = v13
    }
    fmt.Println(v4, v4)
    var v14 uint32 = 680
    var v15 float64 = (0.0 - -2.25)
    for fi16 := 0; fi16 < 13; fi16++ {
        if (-2.25 <= -2.25) {
            if (fi16 < 764) {
                v15 = ((1e10 / v15) + (v15 + float64(float32(v15))))
                var v17 float64 = (float64(((-2.25 - ((1e10 * float32(v15)) / -2.25)) * 0.0)) / v15)
                _ = v17
            }
            var v18 int16 = 1000
            v19 := GenS2{F0: 42, F1: 342, F2: v2}
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    var v20 string = ((("\n\t" + ("a" + "hello")) + "\n\t") + "\n\t")
    fmt.Println(v2, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
