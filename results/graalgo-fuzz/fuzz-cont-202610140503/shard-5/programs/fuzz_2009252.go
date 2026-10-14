package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }
type GenS2 struct { F0 string; F1 uint }

func main() {
    var v1 int64 = 1
    var v2 uint64 = 643
    var v3 int64 = -9223372036854775808
    fmt.Println(v2, v1, v2)
    switch -799 {
    case 127:
        v4 := GenS1{F0: (v3 / ((1) | 1)), F1: -1, F2: -1}
        _ = v4
    case 32768:
        var v5 string = "\n\t"
        _ = v5
    }
    var v6 float32 = 0.1
    v7 := GenS1{F0: v3, F1: 42, F2: 100}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint64 = v2
        v9 = v2
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: "foo bar", F1: 427}
    fmt.Println(v6)
    v11 := []int{128}
    var v12 int32 = 2147483647
    _ = len(v11)
    switch (("a" + "\n\t") + ("foo bar" + "a")) {
    case "foo bar":
        _ = len(v11)
    default:
        if true {
            for fi13 := 0; fi13 < 16; fi13++ {
                for fi14 := 0; fi14 < 15; fi14++ {
                    fmt.Println(fi13, fi14)
                    _ = fi14
                }
                v15 := GenS1{F0: (^(v1)), F1: -1, F2: -1}
                v3 += (v3 - (v3 + (v3 << ((1) & 63))))
                _ = fi13
                _ = v15
            }
            v1 = (v1 >> ((2) & 63))
            var v16 uint64 = v2
            _ = v16
        } else {
            v11 = append(v11, -1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
