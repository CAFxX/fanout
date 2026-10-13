package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int16 = 1000
    var v2 float64 = ((float64((0.0 - (0.0 / 1.0))) / 3.14159) / 1e10)
    var v3 int16 = v1
    var v4 uint32 = 2147483647
    var v5 uint64 = 509
    var v6 int8 = 42
    var v7 int32 = -85
    if (true && false) {
        v1 = 1
        var v8 float32 = 1e10
        var v9 int16 = (v3 << ((93) & 63))
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := []int{int(v7), fi10, 127, (^(fi10))}
            _ = fi10
            _ = v11
        }
        switch "foo bar" {
        case "":
            var v12 int32 = v7
            _ = v12
        case "hello":
            v5 += 4294967295
        case "\n\t":
            v13 := []int{int(v7), 169, 2147483647}
            _ = v13
        default:
            fmt.Println(v3)
        }
        var v14 int64 = -655
        _ = v14
    }
    fmt.Println(v4, v6)
    v15 := []int{128, -1, 32767, -318}
    v1 = 0
    v1 -= -32768
    fmt.Println(v2, v1, v5)
    v15 = append(v15, 1)
    v16 := GenS1{F0: 1}
    v15 = append(v15, 127)
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
