package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int64; F2 uint }
func genFn0(p0 int8, p1 int32) uint64 {
    fmt.Println(p1, p0)
    v7 := GenS1{F0: p1}
    var v8 uint8 = 45
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 bool = true
    var v2 int64 = 1
    v3 := make(map[int]string)
    var v4 int16 = -730
    v5 := make(map[int]uint32)
    var v6 float64 = (float64(3.14159) / 0.1)
    if v1 {
        fmt.Println(v6, v2)
        v2 = 0
    } else {
        v9 := GenS1{F0: 0}
        switch "hello" {
        case "a":
            v4 = 32767
        case "\n\t":
            var v10 uint16 = 423
            _ = v10
        }
        v11 := GenS2{F0: 127}
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int16 = (v4 & v4)
        _ = fi12
        _ = v13
    }
    v1 = v1
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int = 127
        v16 := GenS1{F0: 832}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int = 32767
    delete(v5, (v17 & (^(v17))))
    v18 := GenS3{F0: int32(v4), F1: v2, F2: 131}
    v19 := []int{}
    var v20 int = v17
    var v21 bool = (!((0 != 2) || (true && v1)))
    if len(v19) > 0 {
        v19[0] = 218
    }
    if v1 {
        v3[(v20 >> ((798) & 63))] = "hello"
    } else {
        v22 := GenS2{F0: 67}
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 uint16 = 32767
            _ = fi23
            _ = v24
        }
        var v25 uint16 = 1
        _ = v22
        _ = v25
    }
    var v26 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
