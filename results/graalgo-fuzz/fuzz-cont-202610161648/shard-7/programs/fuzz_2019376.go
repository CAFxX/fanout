package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
func genFn0(p0 bool) uint8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint16 = 606
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return 128
}

func main() {
    var v1 int = -47
    var v2 float32 = 1e10
    v3 := make(map[uint32]bool)
    var v4 int = v1
    var v5 string = "hello"
    v6 := make(map[uint32]bool)
    v9 := []int{(-(v4)), (v1 << ((42) & 63)), v4}
    v4 = v1
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 uint16 = 32767
        for fi12 := 0; fi12 < 6; fi12++ {
            _ = len(v3)
            var v13 int32 = 0
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = len(v3)
    v14 := GenS1{F0: -742, F1: v4, F2: 8}
    var v15 int = v1
    var v16 float64 = ((1e10 - (-2.25 / ((3.14159 + (0.0 * 0.1)) - 0.0))) + (1e10 - -2.25))
    for fi17 := 0; fi17 < 17; fi17++ {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi17
    }
    _ = len(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
