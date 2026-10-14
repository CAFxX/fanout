package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int64 = -786
    v2 := make(map[string]uint32)
    v3 := []int{}
    var v4 uint32 = 958
    var v5 int32 = 679
    var v6 float32 = ((1e10 - 1.5) - (-2.25 * 1e10))
    v7 := GenS1{F0: (("\n\t" + ("" + "hello")) + (("a" + "foo bar") + ""))}
    v1 = v1
    switch (("\n\t" + "hello") + "\n\t") {
    case "\n\t":
        _ = len(v3)
    default:
        _ = len(v3)
    }
    v8 := GenS1{F0: (("foo bar" + ("" + "")) + "\n\t")}
    if false {
        v4 = v4
    } else {
        v9 := GenS1{F0: "foo bar"}
        var v10 float64 = 0.1
        v5 -= v5
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v1 += v1
        _ = fi11
    }
    var v12 int16 = -313
    v2["a"] = (v4 * (v4 + v4))
    _ = len(v3)
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint64 = 1
        fmt.Println(v12, v12, fi13)
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: ((("" + ((("a" + "foo bar") + ("" + "")) + "")) + "a") + "foo bar")}
    var v16 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
