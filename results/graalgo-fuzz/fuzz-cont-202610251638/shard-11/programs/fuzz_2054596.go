package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int32 = 42
    var v2 float64 = -2.25
    var v3 bool = false
    var v4 bool = v3
    var v5 float64 = (v2 - v2)
    var v6 bool = ((v3 && ((!false) || true)) && (1 >= 2))
    v7 := []int{-9223372036854775808}
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            v5 -= v2
            var v10 int = -2147483648
            fmt.Println(v3)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    fmt.Println(v6)
    v11 := GenS1{F0: 0}
    var v12 uint16 = 255
    v13 := GenS2{F0: v1}
    switch "a" {
    case "\n\t":
        var v14 uint64 = 1
        _ = v14
    }
    if len(v7) > 0 {
        v7[0] = 32767
    }
    v3 = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
