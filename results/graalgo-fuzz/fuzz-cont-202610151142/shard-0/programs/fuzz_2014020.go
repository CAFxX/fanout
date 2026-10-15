package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int = 32768
    v2 := make(map[string]bool)
    var v3 bool = (!true)
    v4 := GenS1{F0: 127, F1: 246}
    fmt.Println(v1)
    v5 := v2["\n\t"]
    var v6 float32 = -2.25
    fmt.Println(v1, v5, v5)
    var v7 uint64 = 42
    var v8 string = "\n\t"
    var v9 uint32 = 42
    var v10 uint16 = 32767
    fmt.Println(v10, v3)
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS2{F0: v9}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
