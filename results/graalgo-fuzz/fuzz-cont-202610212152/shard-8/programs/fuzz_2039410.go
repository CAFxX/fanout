package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 1.5
    v3 := []string{}
    var v4 string = ""
    _ = len(v1)
    var v5 float64 = v2
    for fi6 := 0; fi6 < 3; fi6++ {
        if len(v3) > 0 {
            v3[0] = v4
        }
        _ = fi6
    }
    v7 := GenS1{F0: 2, F1: 127}
    switch -595 {
    case -396:
        v4 = (v4 + v4)
    default:
        var v8 uint16 = 1
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int8 = 42
        for fi11 := 0; fi11 < 10; fi11++ {
            fmt.Println(v4, fi11, v4)
            v5 += (0.1 * (float64(1.5) * v5))
            delete(v1, v4)
            _ = fi11
        }
        if true {
            var v12 int16 = -337
            var v13 float64 = v2
            v14 := GenS1{F0: 1, F1: (v10 / ((0) | 1))}
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    var v15 string = ((v4 + v4) + "a")
    fmt.Println(v5, v15)
    var v16 uint32 = 824
    var v17 int16 = 32767
    var v18 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
