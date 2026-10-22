package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = ((-2147483648 < 1) && false)
    var v3 int8 = 0
    var v4 float32 = ((3.14159 - (1.5 - (1.5 / float32((1.5 + 0.0))))) / ((-2.25 * 0.0) - (1e10 - float32(0.0))))
    v5 := make(map[int]int)
    var v6 uint = 65535
    var v7 string = "\n\t"
    fmt.Println(v7)
    if v2 {
        var v8 float32 = ((v4 - v4) / 0.1)
        _ = v8
    }
    v6 = 610
    v2 = false
    fmt.Println(v7, v2, v4)
    v9 := GenS1{F0: 127, F1: 1, F2: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
