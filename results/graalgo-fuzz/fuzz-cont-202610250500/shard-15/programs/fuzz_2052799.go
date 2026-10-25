package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int = 1
    var v2 float64 = 0.0
    v3 := make(map[uint32]bool)
    var v4 string = ("\n\t" + "hello")
    v5 := GenS1{F0: 441}
    fmt.Println(v2, v4, v2)
    v2 = float64(1e10)
    v6 := GenS2{F0: 0, F1: 32767}
    var v7 int = v1
    v2 = v2
    var v8 uint64 = 4294967295
    var v9 uint64 = v8
    if (2 == 2) {
        for fi10 := 0; fi10 < 2; fi10++ {
            if false {
                v1 -= v1
            } else {
                var v11 uint64 = (v8 + (v8 & v8))
                _ = v11
            }
            v12 := []int{}
            _ = fi10
            _ = v12
        }
        v13 := GenS2{F0: -1, F1: 36}
        _ = v13
    }
    if (!true) {
        fmt.Println(v7)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
