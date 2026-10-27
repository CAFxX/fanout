package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 739
    var v3 uint = 164
    var v4 uint64 = 788
    var v5 float64 = 0.0
    v1[32767] = "a"
    var v6 float64 = 0.0
    switch "foo bar" {
    case "\n\t":
        v6 += (v5 / (((float64(0.0) / v6) * v6) - v5))
    case "a":
        v7 := GenS1{F0: v2}
        _ = v7
    default:
        v8 := GenS1{F0: (v2 + v2)}
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v3 += 879
        v4 = v4
        _ = fi9
    }
    var v10 int8 = -1
    delete(v1, -433)
    var v11 int64 = 1
    var v12 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
