package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 float32 = -2.25
    var v2 int8 = 0
    var v3 int8 = 42
    var v4 uint = 18446744073709551615
    var v5 int64 = 907
    var v6 int8 = 100
    fmt.Println(v5, v2)
    for fi7 := 0; fi7 < 18; fi7++ {
        v2 = v2
        _ = fi7
    }
    fmt.Println(v1, v2)
    switch "hello" {
    case "a":
        var v8 int8 = v6
        _ = v8
    case "hello":
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: (v4 - 42)}
            v11 := []int{(fi9 >> (((v4 / (((v4 + v4)) | 1))) & 63)), fi9, 2147483647, (fi9 / (((^(fi9))) | 1))}
            fmt.Println(v3, v3)
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 uint8 = 255
    var v13 uint64 = 0
    for fi14 := 0; fi14 < 3; fi14++ {
        v2 = (v6 << ((v4) & 63))
        _ = fi14
    }
    v15 := []int{-831, -9223372036854775808, 32768, 9223372036854775807}
    if len(v15) > 0 {
        v15[0] = 279
    }
    var v16 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
