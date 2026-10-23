package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 int32, p1 string) uint {
    switch (((p1 + p1) + "") + p1) {
    case "a":
        p1 = (p1 + p1)
    default:
        fmt.Println(p0)
    }
    v6 := []int{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = -2.25
    var v3 uint32 = 2147483647
    var v4 uint32 = 4294967295
    var v5 int32 = 0
    _ = len(v1)
    v7 := GenS1{F0: -182, F1: -2147483648}
    v8 := GenS1{F0: 0, F1: (v5 << ((18446744073709551615) & 63))}
    fmt.Println(v4, v4, v2)
    v1["a"] = ("\n\t" == "\n\t")
    v9 := GenS1{F0: 42, F1: (^(v5))}
    v3 = 231
    var v10 int16 = -32768
    v11 := GenS1{F0: -74, F1: v5}
    v4 = (v4 / (((v3 - v4)) | 1))
    v12 := GenS2{F0: "foo bar"}
    var v13 int = 2147483647
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
