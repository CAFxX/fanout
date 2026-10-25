package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 0
    var v3 int = 36
    fmt.Println(v3, v3)
    v3 += 1
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: "foo bar", F1: "foo bar", F2: 2}
        v2 = (v2 / ((v2) | 1))
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: "hello", F1: ("" + ""), F2: 127}
    v7 := []int{(-(v3)), v3, v3, 505}
    var v8 string = "\n\t"
    v2 = (v2 - v2)
    var v9 string = v8
    v2 -= (v2 / ((42) | 1))
    fmt.Println(v9)
    v2 = (v2 + (v2 ^ 0))
    var v10 float64 = 0.1
    var v11 int = v3
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 float64 = -2.25
        delete(v1, (v11 << ((252) & 63)))
        _ = fi12
        _ = v13
    }
    var v14 int16 = 42
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := GenS1{F0: "hello", F1: ((v8 + (v9 + "a")) + v8), F2: 42}
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v7 = append(v7, (^(v11)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
