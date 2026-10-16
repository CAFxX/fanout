package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 3.14159
    var v3 float32 = v2
    v2 = float32((-2.25 + 3.14159))
    for fi4 := 0; fi4 < 10; fi4++ {
        v2 += float32(-2.25)
        _ = fi4
    }
    v5 := GenS1{F0: (!false), F1: 42, F2: "a"}
    var v6 uint8 = 190
    v7 := v1["a"]
    var v8 uint64 = 393
    var v9 uint16 = 42
    v2 -= 3.14159
    v10 := v1[("" + ("\n\t" + "a"))]
    v11 := v1["a"]
    v12 := GenS1{F0: (0 != 771), F1: -1, F2: "foo bar"}
    var v13 float64 = 1e10
    v9 = (v9 >> ((255) & 63))
    v13 += v13
    v13 += 0.1
    _ = len(v1)
    if true {
        fmt.Println(v11, v3, v11)
        v14 := v1[""]
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
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
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
