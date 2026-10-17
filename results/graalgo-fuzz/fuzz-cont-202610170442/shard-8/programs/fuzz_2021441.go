package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int }
func genFn0(p0 int16, p1 int16) int8 {
    v4 := []int{2147483647}
    var v5 float32 = 0.0
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int64 = -828
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 0, F2: 42}
    var v9 uint32 = 31
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 bool = false
    var v2 float64 = 0.1
    v3 := make(map[int]int)
    fmt.Println(v2)
    _ = len(v3)
    var v10 uint64 = 631
    for fi11 := 0; fi11 < 7; fi11++ {
        switch "a" {
        case "hello":
            var v12 uint32 = 0
            _ = v12
        }
        _ = fi11
    }
    v1 = true
    _ = len(v3)
    if v1 {
        v13 := GenS1{F0: v10, F1: 1, F2: 9223372036854775807}
        _ = v13
    }
    switch 42 {
    case 0:
        v10 = v10
    }
    v10 = 2
    switch 4294967295 {
    case -1:
        v10 -= 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
