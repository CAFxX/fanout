package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 float32 = 1.5
    var v2 int16 = 42
    var v3 int = 4294967295
    v4 := []uint32{57}
    var v5 uint = 65535
    var v6 int8 = 42
    var v7 bool = (!false)
    for fi8 := 0; fi8 < 9; fi8++ {
        v7 = ((!false) || true)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v2, fi8)
        _ = fi8
    }
    v9 := GenS1{F0: -1}
    v1 += (-2.25 - v1)
    v1 += v1
    v1 = (float32((((-2.25 + 1.5) - 0.1) - 0.0)) / v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
