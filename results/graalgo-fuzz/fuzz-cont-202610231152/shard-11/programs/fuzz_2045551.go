package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    var v1 int = 966
    v2 := make(map[int]string)
    var v3 int64 = 0
    v4 := []uint32{}
    var v5 int16 = 1
    fmt.Println(v1)
    v5 -= -32768
    var v6 int = v1
    var v7 bool = (!(1 != 0))
    v4 = append(v4, 987)
    var v8 int32 = 1
    v4 = append(v4, 42)
    if (v7 || ("a" >= "foo bar")) {
        var v9 uint = 2
        fmt.Println(v8)
        _ = v9
    } else {
        fmt.Println(v3, v5)
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 float32 = 1.5
            var v12 uint64 = 735
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v4 = append(v4, 65535)
        _ = fi13
    }
    v1 -= v6
    var v14 uint16 = 0
    for fi15 := 0; fi15 < 2; fi15++ {
        if v7 {
            v8 = v8
            _ = len(v4)
        } else {
            if len(v4) > 0 {
                v4[0] = 42
            }
            v16 := GenS1{F0: -768, F1: v1}
            _ = v16
        }
        var v17 float32 = 3.14159
        v17 = float32((((1e10 - (3.14159 / 1.0)) / 3.14159) - 0.0))
        _ = fi15
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
