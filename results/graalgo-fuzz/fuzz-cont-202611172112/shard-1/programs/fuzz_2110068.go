package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 float64 = (0.0 - 0.1)
    var v2 float32 = (float32(1e10) / 1.0)
    v3 := make(map[string]bool)
    var v4 float64 = 0.1
    var v5 int16 = 32767
    v6 := []int{-939}
    v7 := v3["foo bar"]
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint16 = 32767
        var v10 uint16 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(v2, v7)
            fmt.Println(fi11)
            var v12 int64 = -9223372036854775808
            _ = fi11
            _ = v12
        }
    }
    v2 = ((v2 * v2) - v2)
    v13 := GenS1{F0: 42}
    v14 := GenS1{F0: 61}
    var v15 uint16 = 65535
    if v7 {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
