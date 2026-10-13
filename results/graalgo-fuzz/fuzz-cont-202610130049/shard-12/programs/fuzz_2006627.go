package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 string, p1 uint) uint32 {
    p0 = "foo bar"
    v4 := []int{128, -2147483648}
    _ = p0
    _ = p1
    _ = v4
    return 643
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    var v5 int8 = 42
    for fi6 := 0; fi6 < 1; fi6++ {
        switch "" {
        case "":
            v7 := []int{(fi6 ^ (fi6 ^ fi6)), fi6}
            _ = v7
        case "foo bar":
            var v8 float64 = 0.0
            _ = v8
        }
        v5 = v5
        _ = fi6
    }
    fmt.Println(v5)
    v9 := v3["hello"]
    v10 := GenS1{F0: 245}
    v9 = (v9 + v9)
    v11 := v1[1]
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 float32 = float32((1e10 - -2.25))
        fmt.Println(v9)
        v9 = ("a" + v9)
        _ = fi12
        _ = v13
    }
    v14 := []int{-9223372036854775808}
    var v15 int64 = 557
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
