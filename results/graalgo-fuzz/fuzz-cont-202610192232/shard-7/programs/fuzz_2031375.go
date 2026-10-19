package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int }
func genFn0(p0 uint16, p1 uint64) int8 {
    fmt.Println(p1, p0)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint16 = p0
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return 55
}

func main() {
    var v1 int32 = 2147483647
    var v2 bool = true
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    v5 := make(map[string]bool)
    var v6 uint16 = 32767
    for fi9 := 0; fi9 < 19; fi9++ {
        v1 -= (^(v1))
        fmt.Println(fi9, v1, v2)
        var v10 int8 = -128
        _ = fi9
        _ = v10
    }
    v4[("hello" + ((("a" + "") + "foo bar") + "\n\t"))] = "foo bar"
    v11 := []int{128, 1}
    var v12 int16 = -1
    var v13 string = ("a" + ("hello" + "hello"))
    var v14 float32 = 1.5
    v12 -= 32767
    v15 := GenS1{F0: (v12 ^ v12), F1: 289, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
