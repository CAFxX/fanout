package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 int; F1 string; F2 int32 }

func main() {
    v1 := []int{2147483647}
    v2 := make(map[uint32]int)
    v3 := []int{}
    var v4 uint16 = 1
    var v5 int32 = 2147483647
    var v6 int = 729
    v7 := GenS1{F0: "foo bar", F1: (((true || (1e10 < 3.14159)) || true) || false), F2: 42}
    v6 += v6
    v1 = append(v1, (v6 | v6))
    v8 := GenS1{F0: ("foo bar" + "\n\t"), F1: false, F2: 275}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        var v9 float64 = 1.5
        var v10 uint8 = 42
        _ = v10
        _ = v9
    } else {
        v11 := GenS2{F0: v6, F1: ("\n\t" + "foo bar"), F2: v5}
        var v12 uint32 = 65535
        _ = v11
        _ = v12
    }
    var v13 uint32 = 104
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
