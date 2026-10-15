package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 bool = true
    var v2 int8 = -128
    var v3 int32 = -383
    fmt.Println(v1, v2)
    v4 := GenS1{F0: ("hello" + "\n\t")}
    var v5 float64 = 1e10
    fmt.Println(v1)
    v2 += (v2 >> ((928) & 63))
    switch 0 {
    case 2147483647:
        var v6 int16 = -284
        _ = v6
    case 724:
        var v7 float64 = ((v5 * v5) - (0.1 + (float64((1e10 - 0.0)) * v5)))
        _ = v7
    case 785:
        v1 = v1
    default:
        var v8 uint = 643
        _ = v8
    }
    v9 := []int{1, -2147483648, 32768}
    switch "" {
    case "a":
        v2 += v2
    default:
        var v10 uint8 = 128
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
