package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = -415
    var v3 uint16 = 12
    var v4 uint = 4294967295
    var v5 int32 = 2147483647
    v6 := []int{2147483647}
    v7 := GenS1{F0: 666, F1: 169, F2: 2147483647}
    v8 := v1["hello"]
    fmt.Println(v4, v8, v5)
    _ = len(v6)
    v3 += v3
    v4 = v4
    v8 -= v8
    var v9 uint8 = 255
    if false {
        v6 = append(v6, 1)
    }
    var v10 int64 = 42
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint16 = 32767
        v13 := GenS1{F0: v8, F1: (^(v8)), F2: 0}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int8 = -44
    delete(v1, "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
