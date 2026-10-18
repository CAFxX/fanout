package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }

func main() {
    var v1 int = 600
    v2 := []int{0, v1, (v1 | (v1 << ((0) & 63))), (v1 | (v1 + (v1 >> ((42) & 63))))}
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = (v1 >> ((65535) & 63))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        v1 -= fi6
        var v7 uint = 4294967295
        _ = fi6
        _ = v7
    }
    var v8 float64 = 1.5
    switch "hello" {
    case "foo bar":
        v9 := GenS1{F0: -1, F1: true, F2: -654}
        _ = v9
    case "a":
        var v10 int64 = 993
        _ = v10
    }
    v11 := GenS1{F0: 0, F1: true, F2: 312}
    fmt.Println(v8, v8)
    v1 -= v1
    if false {
        if (false || (!(!true))) {
            v12 := GenS1{F0: 0, F1: (false || (float32(v8) < (1e10 / 1e10))), F2: 302}
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 float64 = v8
                var v15 uint64 = 536
                var v16 uint8 = 128
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            if false {
                var v17 int8 = 1
                var v18 bool = true
                _ = v17
                _ = v18
            }
            _ = v12
        }
    }
    v8 += float64(float32(3.14159))
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(fi19, v8, fi19)
        v8 -= v8
        var v20 int8 = 27
        _ = fi19
        _ = v20
    }
    v1 -= (v1 % ((-1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
