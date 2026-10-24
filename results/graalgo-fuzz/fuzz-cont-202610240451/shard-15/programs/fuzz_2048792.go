package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }

func main() {
    v1 := []string{"a"}
    var v2 int8 = 0
    v3 := make(map[string]int)
    var v4 int = -755
    var v5 int8 = v2
    var v6 float32 = -2.25
    v5 = (-(v2))
    v5 = 42
    var v7 uint32 = 2147483647
    var v8 float32 = 1e10
    fmt.Println(v4, v8)
    var v9 int32 = 626
    v10 := v3["\n\t"]
    _ = len(v3)
    _ = len(v1)
    v3["foo bar"] = (v4 << ((4294967295) & 63))
    var v11 float64 = 3.14159
    var v12 int = v4
    v11 = 0.0
    if (56 <= 42) {
        if true {
            var v13 uint = 0
            v5 = v2
            _ = v13
        }
        var v14 uint32 = v7
        if ((v10 | (v4 - (v12 + v10))) >= 1) {
            v1 = append(v1, "\n\t")
            switch ("a" + "a") {
            case "\n\t":
                var v15 uint16 = 0
                _ = v15
            }
        }
        _ = v14
    } else {
        var v16 string = ("foo bar" + "\n\t")
        var v17 float32 = -2.25
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: "\n\t", F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
