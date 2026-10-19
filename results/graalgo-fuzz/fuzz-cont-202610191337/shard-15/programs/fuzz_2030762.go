package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 string }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 bool = (65535 <= 2)
    var v4 int16 = 281
    var v5 string = ""
    for fi6 := 0; fi6 < 6; fi6++ {
        v4 = -826
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    v5 = (("\n\t" + ("a" + ("a" + ("hello" + (v5 + (("" + v5) + "foo bar")))))) + v5)
    v8 := v2["a"]
    v9 := GenS1{F0: v4, F1: 1, F2: "foo bar"}
    v10 := []int{-233}
    fmt.Println(v5)
    var v11 float32 = float32((-2.25 - (3.14159 + (0.0 - (((0.0 - (float64(3.14159) * (1.5 * 1.5))) * 3.14159) * 0.1)))))
    fmt.Println(v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
