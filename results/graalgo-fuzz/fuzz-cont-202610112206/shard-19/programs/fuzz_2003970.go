package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 int32 }

func main() {
    var v1 float64 = ((3.14159 + 3.14159) - 0.1)
    var v2 uint8 = 186
    v3 := make(map[uint32]string)
    v4 := make(map[string]string)
    v5 := GenS1{F0: 93, F1: -9223372036854775808, F2: 2}
    v1 = (1e10 - (3.14159 * 3.14159))
    v1 += ((v1 + v1) - v1)
    var v6 int8 = -128
    var v7 uint16 = 255
    fmt.Println(v1, v2)
    v8 := GenS2{F0: v2}
    v9 := GenS3{F0: 1, F1: -1, F2: -246}
    v6 += v6
    var v10 uint32 = 1
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 string = "foo bar"
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
