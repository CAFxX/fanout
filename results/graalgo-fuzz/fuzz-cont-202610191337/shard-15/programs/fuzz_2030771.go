package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []string{"\n\t", "", ("" + ("a" + (((("a" + "\n\t") + "") + ("a" + ("\n\t" + "foo bar"))) + ("hello" + ("" + ((("hello" + ("a" + "")) + "\n\t") + "a")))))), (("\n\t" + ("a" + "hello")) + "hello")}
    var v2 int = 228
    var v3 int32 = -1
    v4 := []uint32{}
    var v5 int8 = -1
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 2, F1: 2}
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: 201}
    v3 += v3
    v3 += (v3 ^ (v3 & (v3 * v3)))
    var v9 float64 = 1e10
    if false {
        var v10 bool = (!true)
        _ = v10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
