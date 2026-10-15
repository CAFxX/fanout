package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 bool }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    v4 := []string{"a", ((("foo bar" + ("foo bar" + ("a" + ("\n\t" + "\n\t")))) + "hello") + "hello"), "foo bar", ("foo bar" + "\n\t")}
    var v5 uint32 = 4294967295
    var v6 uint32 = 42
    for fi7 := 0; fi7 < 20; fi7++ {
        v6 = (v6 * (v5 / ((2147483648) | 1)))
        _ = fi7
    }
    fmt.Println(v5)
    var v8 float64 = (-2.25 - 0.0)
    var v9 bool = (true && (123 <= 42))
    fmt.Println(v8, v9)
    delete(v3, (v6 * v6))
    var v10 int64 = 0
    if v9 {
        v11 := GenS1{F0: 18446744073709551615, F1: (42 >= 2), F2: v9}
        _ = v11
    } else {
        v12 := v2[(v6 ^ (^(v6)))]
        var v13 int = (v12 & (v12 * (v12 & v12)))
        v14 := GenS1{F0: 0, F1: true, F2: v9}
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        fmt.Println(v5, v10, v5)
        _ = fi15
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
