package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0(p0 int16, p1 int32) int32 {
    var v5 uint16 = 42
    p1 += (p1 % ((-295) | 1))
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    v1 := []int{}
    var v2 uint64 = 2
    var v3 float64 = 0.0
    var v4 uint64 = v2
    v4 = 307
    v6 := GenS1{F0: 0, F1: (("\n\t" + ("hello" + "foo bar")) + "a")}
    v4 = v2
    v2 -= (^(v2))
    v7 := GenS1{F0: -47, F1: "foo bar"}
    _ = len(v1)
    fmt.Println(v4)
    var v8 uint32 = 516
    v9 := GenS1{F0: 100, F1: "\n\t"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
