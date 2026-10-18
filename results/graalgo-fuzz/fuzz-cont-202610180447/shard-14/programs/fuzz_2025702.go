package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint8; F2 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 float32 = -2.25
    var v4 float32 = v3
    var v5 uint64 = 18446744073709551615
    v3 = float32(((float64(1.5) - float64(0.0)) - float64(v3)))
    v6 := GenS1{F0: 205}
    fmt.Println(v3)
    v4 += (v4 / v3)
    var v7 string = ""
    fmt.Println(v3, v7)
    v8 := v2[-450]
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int64 = -903
        if (2 > 2) {
            if false {
                v11 := GenS1{F0: 420}
                v8 = (v3 >= v3)
                v12 := GenS2{F0: "", F1: 62, F2: 773}
                _ = v11
                _ = v12
            }
            v7 = v7
        } else {
            var v13 float32 = (v4 + 0.1)
            _ = v13
        }
        v7 = v7
        _ = fi9
        _ = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
