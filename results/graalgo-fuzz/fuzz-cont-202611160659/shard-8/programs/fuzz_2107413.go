package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint64 }
func genFn0(p0 int, p1 uint) int {
    var v6 int16 = 42
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 1, F1: -148, F2: 42}
        v9 := []int{}
        var v10 uint16 = 255
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    p1 -= 18446744073709551615
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return (p0 + p0)
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]string)
    var v3 int16 = 1
    v4 := make(map[string]int)
    var v5 int64 = 1
    fmt.Println(v5, v5)
    v12 := GenS1{F0: (v5 << ((2) & 63)), F1: v5, F2: 449}
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v3, fi13)
        _ = fi13
    }
    fmt.Println(v3, v5)
    if true {
        switch "foo bar" {
        case "foo bar":
            var v14 uint32 = 88
            _ = v14
        case "\n\t":
            v15 := GenS1{F0: -1, F1: (v5 & (-(v5))), F2: 1}
            _ = v15
        }
        v5 += v5
        var v16 int16 = v3
        _ = v16
    }
    v17 := GenS1{F0: -605, F1: -858, F2: 4294967295}
    for fi18 := 0; fi18 < 15; fi18++ {
        v5 = 0
        _ = fi18
    }
    v3 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
