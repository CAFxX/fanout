package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
func genFn0(p0 int, p1 int8, p2 uint32) int {
    p1 = (p1 * -128)
    v5 := []int{}
    v6 := GenS1{F0: "", F1: uint64(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 65535
    var v3 uint8 = 216
    var v4 float64 = (0.0 + 1.5)
    if false {
        v4 += (float64(0.1) / 1.0)
    }
    if false {
        var v7 uint8 = v3
        _ = v7
    }
    fmt.Println(v2)
    switch "hello" {
    case "foo bar":
        var v8 string = "hello"
        _ = v8
    }
    v3 = v3
    fmt.Println(v2, v2)
    v9 := GenS1{F0: "", F1: 477}
    for fi10 := 0; fi10 < 7; fi10++ {
        v3 = (v3 / ((1) | 1))
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int64 = 9223372036854775807
            v4 -= (v4 * ((v4 - v4) + v4))
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
