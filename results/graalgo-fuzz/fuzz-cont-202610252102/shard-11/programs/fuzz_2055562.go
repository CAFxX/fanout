package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    v1 := []string{"", "foo bar"}
    v2 := make(map[string]uint32)
    var v3 int = 2147483647
    v4 := make(map[uint32]bool)
    var v5 float64 = 1e10
    v6 := GenS1{F0: true, F1: 0}
    v7 := GenS1{F0: (true && (!false)), F1: 968}
    v5 -= ((v5 - v5) + (v5 - (v5 * 3.14159)))
    v5 += v5
    if false {
        v3 += -189
        var v8 uint8 = 0
        fmt.Println(v3, v3, v5)
        _ = v8
    } else {
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    }
    v3 = -9223372036854775808
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := v2["\n\t"]
        var v11 int16 = -32768
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v5)
    fmt.Println(v3, v3)
    v2["hello"] = 898
    _ = len(v1)
    var v12 int16 = 206
    if len(v1) > 0 {
        v1[0] = ""
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
