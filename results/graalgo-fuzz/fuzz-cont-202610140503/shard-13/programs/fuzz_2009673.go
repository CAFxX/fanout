package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 0
    var v3 int16 = 32767
    v4 := make(map[int]string)
    v5 := make(map[int]int)
    var v6 int8 = (v2 - (v2 | v2))
    delete(v4, -199)
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 497}
        v9 := GenS2{F0: "a", F1: 2, F2: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v3, v3)
    switch "hello" {
    case "\n\t":
        v6 -= v2
    case "":
        v10 := []int{}
        _ = v10
    case "foo bar":
        v11 := v1[("hello" + "\n\t")]
        _ = v11
    }
    v12 := GenS1{F0: 255}
    for fi13 := 0; fi13 < 2; fi13++ {
        if false {
            fmt.Println(v3, fi13)
        }
        _ = fi13
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
