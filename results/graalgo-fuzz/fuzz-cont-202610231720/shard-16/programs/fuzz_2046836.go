package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }

func main() {
    var v1 uint64 = 2
    var v2 float32 = 3.14159
    var v3 int32 = -821
    v4 := []string{(("" + "\n\t") + ("hello" + "")), "\n\t", "foo bar"}
    var v5 float64 = (1e10 / 1e10)
    for fi6 := 0; fi6 < 3; fi6++ {
        v3 = 0
        fmt.Println(fi6, v2, v3)
        _ = fi6
    }
    var v7 float32 = 1.5
    var v8 int8 = 41
    var v9 int64 = -93
    v7 -= 0.1
    _ = len(v4)
    v1 -= (^(v1))
    var v10 int16 = 146
    _ = len(v4)
    var v11 int32 = v3
    var v12 uint8 = 127
    v13 := GenS1{F0: -948, F1: -578, F2: (v10 & (^(v10)))}
    fmt.Println(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
