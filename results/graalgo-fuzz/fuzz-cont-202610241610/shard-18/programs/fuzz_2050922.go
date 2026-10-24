package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
func genFn0(p0 uint8, p1 uint, p2 int) int32 {
    p0 += p0
    p2 = 4294967295
    v6 := []int{}
    if (true || false) {
        var v7 int32 = -675
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 479
}

func main() {
    var v1 int64 = 42
    var v2 uint16 = 1
    var v3 float64 = 0.0
    v4 := make(map[string]bool)
    var v5 uint8 = 2
    var v8 int = 388
    _ = len(v4)
    switch (v8 - int(v2)) {
    case -520:
        fmt.Println(v5)
    }
    v5 = 0
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int32 = 2147483647
        _ = fi9
        _ = v10
    }
    fmt.Println(v5, v5)
    v11 := GenS1{F0: 896}
    v5 -= (^(v5))
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := []int{(v8 / ((262) | 1)), (v8 + 0), (fi12 - 354)}
        var v14 float32 = 1e10
        v15 := GenS2{F0: "\n\t"}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := v4["a"]
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
