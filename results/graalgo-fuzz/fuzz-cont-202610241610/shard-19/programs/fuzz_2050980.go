package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }

func main() {
    var v1 float32 = (1e10 + 1.5)
    var v2 int32 = -2147483648
    var v3 int = 0
    var v4 uint = 255
    var v5 int32 = -442
    var v6 int8 = -128
    var v7 float32 = v1
    var v8 int32 = (^(v2))
    var v9 uint64 = 2
    v10 := []int{v3, v3, (v3 % (((-(v3))) | 1)), v3}
    if len(v10) > 0 {
        v10[0] = (v3 + v3)
    }
    switch (v3 & (^(v3))) {
    case -709:
        var v11 uint = v4
        _ = v11
    case 89:
        var v12 bool = (float64(0.0) != 0.0)
        _ = v12
    }
    if false {
        v10 = append(v10, (v3 / ((4294967295) | 1)))
        var v13 float32 = v7
        var v14 string = ((("hello" + (("a" + "a") + ("a" + "\n\t"))) + (("foo bar" + (("foo bar" + "a") + (("\n\t" + "") + ("" + ("a" + "\n\t"))))) + "a")) + "\n\t")
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: -260, F1: "\n\t", F2: (v2 % ((-2147483648) | 1))}
    var v16 uint8 = 29
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int = (v3 >> ((v4) & 63))
        var v19 int8 = v6
        v19 = (v6 << (((v4 % (((v4 - (v4 ^ 42))) | 1))) & 63))
        _ = fi17
        _ = v18
        _ = v19
    }
    if len(v10) > 0 {
        v10[0] = (v3 / (((v3 + (v3 ^ v3))) | 1))
    }
    v5 = v8
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
