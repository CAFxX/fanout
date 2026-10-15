package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 bool }

func main() {
    v1 := make(map[int]string)
    var v2 string = ("hello" + ("a" + "foo bar"))
    var v3 uint32 = 65535
    var v4 int = -549
    var v5 float32 = ((3.14159 * 0.0) * 1e10)
    fmt.Println(v2, v2)
    v6 := []int{(-(v4)), v4, (v4 + -693), v4}
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int = (v4 ^ fi7)
        _ = fi7
        _ = v8
    }
    v2 = v2
    v3 = v3
    v9 := v1[v4]
    fmt.Println(v9, v2)
    v9 = (v9 + ((v9 + v2) + "foo bar"))
    v9 = "foo bar"
    fmt.Println(v9)
    v10 := GenS1{F0: 2147483647, F1: 18446744073709551615, F2: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
