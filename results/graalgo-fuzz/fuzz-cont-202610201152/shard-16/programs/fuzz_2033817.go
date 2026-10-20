package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    var v1 float64 = 0.1
    var v2 string = "hello"
    var v3 float32 = -2.25
    v4 := make(map[string]int)
    var v5 int64 = -9223372036854775808
    v6 := []string{}
    v7 := v4[v2]
    v8 := GenS1{F0: ("\n\t" + (v2 + ("" + v2))), F1: -395}
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 int64 = (v5 * v5)
        var v11 float64 = (1e10 + v1)
        _ = fi9
        _ = v10
        _ = v11
    }
    v5 = v5
    var v12 uint64 = 288
    for fi13 := 0; fi13 < 19; fi13++ {
        v3 += (float32(v1) * (float32(0.0) * 1e10))
        var v14 float32 = v3
        _ = fi13
        _ = v14
    }
    v1 -= v1
    v2 = "foo bar"
    v15 := GenS1{F0: v2, F1: 0}
    v16 := GenS1{F0: ("\n\t" + (("a" + (("a" + (("" + "\n\t") + (("hello" + v2) + (v2 + "a")))) + ("hello" + v2))) + v2)), F1: -32768}
    var v17 int8 = 0
    var v18 int16 = 1000
    if true {
        v19 := GenS1{F0: "a", F1: (-(v18))}
        fmt.Println(v5, v7)
        v4[v2] = v7
        _ = v19
    }
    v20 := GenS1{F0: v2, F1: (v18 + v18)}
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
