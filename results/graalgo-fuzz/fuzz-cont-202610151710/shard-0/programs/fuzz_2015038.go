package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int; F1 uint32; F2 uint64 }

func main() {
    v1 := []string{"hello", ("foo bar" + "foo bar"), ("foo bar" + "\n\t")}
    v2 := []string{(("foo bar" + "foo bar") + "foo bar")}
    v3 := []int{-9223372036854775808, 421}
    if ((0.1 < 3.14159) || true) {
        v4 := GenS1{F0: 4294967295, F1: 626}
        var v5 bool = (212 > 160)
        _ = v4
        _ = v5
    }
    var v6 bool = (0 >= -18)
    fmt.Println(v6)
    v6 = (((v6 || (255 == 179)) || (v6 || ((v6 && ((0.1 + 0.0) <= -2.25)) && v6))) && (17 < 0))
    fmt.Println(v6)
    v7 := GenS2{F0: 128, F1: 748, F2: 42}
    v6 = ((0.0 + (0.0 / (((0.0 + 0.1) / 1.0) - (1.5 + float64(0.1))))) <= 0.1)
    v6 = (!false)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
