package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 string }
type GenS2 struct { F0 uint16; F1 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 32768
    var v3 uint64 = 2
    var v4 float32 = 0.0
    var v5 float32 = 0.0
    var v6 int8 = -128
    if false {
        var v7 int32 = -2147483648
        var v8 float64 = float64(1.5)
        fmt.Println(v6)
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 1, F1: -32768, F2: "foo bar"}
    fmt.Println(v6, v4)
    var v10 uint = 65535
    v11 := v1[42]
    var v12 int16 = 1000
    if false {
        fmt.Println(v10, v6, v2)
    }
    var v13 int = int(v6)
    v14 := GenS2{F0: 412, F1: 0}
    v12 = v12
    v13 = (v2 - v13)
    v10 = v10
    v15 := []int{(v13 & v11), (v2 ^ 627), (v13 | v2)}
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS2{F0: 65535, F1: 814}
        _ = fi16
        _ = v17
    }
    var v18 float32 = 1e10
    v19 := GenS1{F0: 42, F1: 493, F2: (((("a" + "hello") + "a") + ("" + "\n\t")) + "\n\t")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
