package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int16 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]int)
    var v3 string = ("hello" + "a")
    fmt.Println(v3, v1)
    var v4 int16 = 0
    fmt.Println(v4, v1, v3)
    var v5 uint8 = 127
    v1 = (^(v1))
    delete(v2, 2147483648)
    v6 := []int{}
    v7 := GenS1{F0: 1, F1: -64, F2: v4}
    var v8 string = v3
    var v9 float64 = 0.0
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v4, v3, v5)
        delete(v2, 65535)
        v11 := GenS1{F0: 255, F1: -95, F2: (v4 * (v4 - 32767))}
        _ = fi10
        _ = v11
    }
    v8 = ""
    fmt.Println(v5, v1)
    v8 = "\n\t"
    v1 += 0
    v9 = 1e10
    var v12 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
