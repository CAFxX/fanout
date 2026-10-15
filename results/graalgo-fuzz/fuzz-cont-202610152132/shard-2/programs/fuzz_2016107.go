package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint64 }

func main() {
    v1 := []uint32{}
    var v2 float32 = float32(3.14159)
    var v3 string = "hello"
    var v4 int16 = -32768
    v3 = v3
    v2 = 1.5
    var v5 int64 = -1
    if ((9223372036854775807 == 42) || (476 < 255)) {
        v3 = (((v3 + v3) + "hello") + v3)
        var v6 int16 = 0
        switch 32768 {
        case -848:
            fmt.Println(v5)
        default:
            v4 += v4
        }
        _ = v6
    }
    v7 := GenS1{F0: (false || ((v4 >> ((302) & 63)) == -1)), F1: -69, F2: 829}
    var v8 int16 = v4
    var v9 uint = 18446744073709551615
    v9 -= 424
    var v10 uint = (v9 & 255)
    v10 += v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
