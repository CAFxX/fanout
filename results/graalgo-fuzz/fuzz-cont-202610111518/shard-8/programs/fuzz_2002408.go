package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 0
    v3 := make(map[uint32]uint32)
    var v4 float32 = 3.14159
    var v5 int = 1
    var v6 uint32 = 2147483648
    v7 := v3[v6]
    v8 := GenS1{F0: (float64(v4) > ((3.14159 - -2.25) + (1e10 + (float64((v4 - 0.0)) + -2.25)))), F1: 2, F2: 2}
    _ = len(v1)
    var v9 int32 = 0
    fmt.Println(v9, v6)
    switch "" {
    case "hello":
        if ((v4 / v4) > v4) {
            var v10 float64 = 0.1
            _ = v10
        }
    case "\n\t":
        v1["a"] = (v6 << ((475) & 63))
    case "a":
        v4 = (float32(-2.25) * v4)
    }
    switch "a" {
    case "hello":
        var v11 uint64 = 42
        _ = v11
    }
    var v12 int8 = 0
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint64 = 1
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(v12, v5, v6)
        _ = fi15
    }
    v16 := GenS1{F0: (42 >= (^(v9))), F1: 18446744073709551615, F2: (v2 >> ((956) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
