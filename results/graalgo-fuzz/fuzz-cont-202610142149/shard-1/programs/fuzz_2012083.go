package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint }

func main() {
    var v1 uint16 = 919
    var v2 uint16 = 689
    var v3 uint64 = 42
    fmt.Println(v3, v3, v2)
    v4 := []int{9223372036854775807, 0, 32768, 127}
    v2 += 445
    v5 := GenS1{F0: 260, F1: true, F2: 42}
    var v6 uint16 = v1
    v2 += v6
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    if len(v4) > 0 {
        v4[0] = -653
    }
    v3 = 18446744073709551615
    var v7 float32 = (((-2.25 * -2.25) - (1.5 + 0.1)) / 1e10)
    v8 := GenS2{F0: (uint32(v3) >= 4294967295), F1: 1}
    var v9 string = "foo bar"
    fmt.Println(v2, v6)
    var v10 int8 = 0
    v11 := GenS3{F0: (^(v2)), F1: 2}
    var v12 int = -516
    switch (v9 + ((v9 + v9) + v9)) {
    case "a":
        v13 := GenS1{F0: v12, F1: true, F2: (v10 * (v10 * v10))}
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
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
