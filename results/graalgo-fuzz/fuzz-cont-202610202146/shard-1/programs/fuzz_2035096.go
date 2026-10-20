package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16, p1 int16) uint64 {
    switch "foo bar" {
    case "foo bar":
        p0 += (p1 + p0)
    case "\n\t":
        v7 := []int{32767, 955, 42, 2147483647}
        _ = v7
    case "":
        for fi8 := 0; fi8 < 2; fi8++ {
            p0 -= p0
            var v9 float32 = 1e10
            _ = fi8
            _ = v9
        }
    default:
        var v10 int16 = p1
        _ = v10
    }
    if true {
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return 774
}

func main() {
    var v1 int8 = 98
    v2 := make(map[int]bool)
    var v3 int64 = -1
    v4 := []int{-902, 32767}
    var v5 uint16 = 65535
    v6 := make(map[string]bool)
    for fi11 := 0; fi11 < 2; fi11++ {
        v3 += v3
        v12 := GenS1{F0: 2147483647}
        switch (fi11 << ((1) & 63)) {
        case -9223372036854775808:
            v4 = append(v4, fi11)
        }
        _ = fi11
        _ = v12
    }
    v3 = 42
    var v13 bool = false
    v14 := GenS1{F0: -688}
    var v15 string = "\n\t"
    v13 = v13
    fmt.Println(v15)
    v5 -= 826
    if len(v4) > 0 {
        v4[0] = 32767
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
