package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{("hello" + ""), "foo bar", "", "foo bar"}
    v3 := []uint32{749, 2147483648, 2147483647}
    var v4 int16 = 1000
    _ = len(v2)
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(v4)
        v2 = append(v2, ("\n\t" + "\n\t"))
        _ = fi5
    }
    v6 := GenS1{F0: 0, F1: false, F2: -444}
    v4 = v4
    v7 := GenS2{F0: v4}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int8 = 100
        _ = fi8
        _ = v9
    }
    _ = len(v2)
    var v10 float32 = (1e10 * -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
