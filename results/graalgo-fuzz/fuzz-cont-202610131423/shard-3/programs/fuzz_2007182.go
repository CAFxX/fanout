package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 string = "\n\t"
    var v3 uint = 255
    var v4 int = 32768
    var v5 int32 = -666
    v1[(v4 & -1)] = "a"
    v6 := []int{32767, 685, -358, (^(v4))}
    var v7 uint64 = 18446744073709551615
    switch "foo bar" {
    case "a":
        var v8 uint = 42
        _ = v8
    }
    if (!false) {
        v4 += (v4 ^ (-(v4)))
    } else {
        if true {
            var v9 int16 = -514
            _ = v9
        }
    }
    var v10 int16 = 42
    v11 := GenS1{F0: 2147483647, F1: -564, F2: 443}
    var v12 int32 = 0
    var v13 string = "a"
    fmt.Println(v12, v4, v2)
    for fi14 := 0; fi14 < 8; fi14++ {
        v6 = append(v6, v4)
        _ = fi14
    }
    if false {
        fmt.Println(v4)
    }
    v15 := GenS1{F0: v12, F1: -1, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
