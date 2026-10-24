package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }

func main() {
    var v1 uint8 = 147
    var v2 uint16 = 42
    var v3 string = "foo bar"
    v4 := []uint32{2147483647}
    var v5 uint8 = (^(v1))
    v1 -= v1
    v2 -= 42
    v6 := GenS1{F0: 468, F1: 0, F2: 42}
    var v7 float64 = -2.25
    var v8 float32 = (0.0 - ((-2.25 * 0.1) + 0.1))
    fmt.Println(v7, v8, v8)
    var v9 int16 = -316
    for fi10 := 0; fi10 < 4; fi10++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
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
