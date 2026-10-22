package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
func genFn0(p0 uint32, p1 string, p2 int) uint32 {
    v4 := GenS1{F0: -1, F1: true}
    switch p2 {
    case 102:
        v5 := []int{p2, p2, 127}
        _ = v5
    case 789:
        v6 := []int{(p2 | (-(p2))), (p2 ^ p2), p2, p2}
        _ = v6
    }
    v7 := GenS1{F0: 13, F1: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return p0
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -81
    var v3 uint32 = 899
    var v8 int32 = (v2 % ((v2) | 1))
    v9 := GenS1{F0: 1, F1: (false || false)}
    if (!false) {
        var v10 int64 = 42
        v11 := v1[("\n\t" + ((("hello" + ("foo bar" + "")) + "\n\t") + ""))]
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint16 = 2
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    }
    v8 = v8
    v14 := []int{}
    v3 = (v3 & 2147483648)
    switch -862 {
    case -898:
        v8 = 42
    case 42:
        if true {
            var v15 int64 = -9223372036854775808
            _ = v15
        }
    default:
        v2 -= v8
    }
    v16 := GenS1{F0: -1, F1: (false || false)}
    switch 9223372036854775807 {
    case 0:
        delete(v1, "hello")
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
