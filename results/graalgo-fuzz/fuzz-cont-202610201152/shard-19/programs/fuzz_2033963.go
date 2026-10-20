package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    var v1 int = 32767
    v2 := make(map[string]string)
    v3 := []int{v1, 421, v1, (v1 ^ v1)}
    var v4 float64 = 0.1
    v5 := make(map[uint32]uint32)
    v6 := []int{}
    if (148 < 1000) {
        v7 := GenS1{F0: "\n\t", F1: 42}
        _ = v7
    }
    v1 = v1
    var v8 uint = 547
    v8 -= 0
    var v9 uint32 = 901
    v10 := GenS1{F0: "foo bar", F1: 1000}
    for fi11 := 0; fi11 < 7; fi11++ {
        v6 = append(v6, (^(fi11)))
        if len(v6) > 0 {
            v6[0] = fi11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 bool = false
            var v14 int8 = -128
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    if len(v3) > 0 {
        v3[0] = (v1 | 693)
    }
    var v15 int64 = -200
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
