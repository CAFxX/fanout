package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    var v3 int8 = -1
    var v4 float32 = 1e10
    v5 := []uint32{935, 724}
    v6 := GenS1{F0: 42}
    v7 := GenS2{F0: 42}
    v3 = (v3 + (^(v3)))
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 string = "foo bar"
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 float32 = 1.5
    v13 := GenS2{F0: 208}
    v12 -= v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
