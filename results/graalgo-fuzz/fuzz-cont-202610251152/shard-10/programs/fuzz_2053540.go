package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }

func main() {
    var v1 int8 = -128
    var v2 int8 = -128
    v3 := make(map[string]string)
    var v4 uint64 = 4294967295
    var v5 float32 = 3.14159
    v6 := []int{127}
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int32 = 42
        var v9 float32 = 1e10
        _ = fi7
        _ = v8
        _ = v9
    }
    v2 += (v1 ^ 100)
    fmt.Println(v4, v2, v5)
    var v10 uint = 1
    v5 = 0.1
    var v11 uint64 = (v4 / (((v4 << ((v10) & 63))) | 1))
    _ = len(v3)
    _ = len(v6)
    var v12 uint32 = 483
    v13 := GenS1{F0: 208, F1: -393}
    v1 = v2
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: (v12 & v12), F1: 1}
        var v16 int64 = -9223372036854775808
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 bool = (-2.25 == ((0.1 + float64(v5)) - float64(3.14159)))
    var v18 float32 = 0.0
    var v19 string = ((("foo bar" + ("" + (("hello" + "foo bar") + "a"))) + (("hello" + (("\n\t" + "hello") + "hello")) + "foo bar")) + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
