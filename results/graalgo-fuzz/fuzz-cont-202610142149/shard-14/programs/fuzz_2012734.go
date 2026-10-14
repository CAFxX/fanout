package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("foo bar" + ("" + ""))
    var v3 int64 = 42
    v4 := make(map[uint32]int)
    v2 = v2
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint32 = 2147483648
        _ = fi5
        _ = v6
    }
    v2 = "foo bar"
    v7 := GenS1{F0: 42}
    v2 = "\n\t"
    v8 := []int{}
    v9 := v4[618]
    var v10 float64 = (0.1 + 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
