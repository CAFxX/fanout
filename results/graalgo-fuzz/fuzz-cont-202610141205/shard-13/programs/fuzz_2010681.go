package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    v1 := []uint32{108, 42}
    v2 := make(map[string]uint32)
    var v3 int = 42
    v4 := []string{"a"}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 := GenS1{F0: -47, F1: (((("\n\t" + "") + "\n\t") + "a") + "foo bar")}
    v6 := GenS1{F0: 0, F1: "foo bar"}
    v7 := GenS1{F0: -14, F1: "\n\t"}
    v3 -= (v3 % (((v3 & (v3 << ((422) & 63)))) | 1))
    v8 := GenS1{F0: 100, F1: "\n\t"}
    v3 = (v3 | (v3 & (v3 ^ (v3 / ((257) | 1)))))
    var v9 uint16 = 465
    if ((1e10 / 0.1) >= (0.1 + (float64((3.14159 * (0.0 * float32(((0.1 * 0.1) - 3.14159))))) - 3.14159))) {
        v10 := v2[""]
        v10 -= (v10 << ((4294967295) & 63))
        _ = v10
    }
    fmt.Println(v3, v3)
    v3 = -326
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
