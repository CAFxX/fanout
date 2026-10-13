package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int8 }

func main() {
    var v1 float32 = (float32(((0.0 / (((0.0 / -2.25) / 1.5) - 1e10)) - 1.5)) / 1.0)
    var v2 uint8 = 42
    v3 := []string{"\n\t"}
    var v4 string = "hello"
    var v5 bool = false
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v6 uint64 = 371
    switch v4 {
    case "\n\t":
        fmt.Println(v6)
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int32 = 0
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    v10 := GenS1{F0: uint(v2)}
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS2{F0: (v4 + v4), F1: 224}
        var v13 int16 = 42
        var v14 int = 32767
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v15 := GenS3{F0: 1, F1: v2}
    v4 = "hello"
    fmt.Println(v6)
    fmt.Println(v1, v1, v9)
    v16 := GenS4{F0: v6, F1: -1}
    v17 := GenS3{F0: 42, F1: 127}
    var v18 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
