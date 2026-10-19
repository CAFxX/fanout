package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ("a" + "")
    var v3 int = 1
    for fi4 := 0; fi4 < 20; fi4++ {
        v3 = fi4
        _ = fi4
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 float32 = (0.1 + 1e10)
        _ = fi5
        _ = v6
    }
    v7 := []int{-9223372036854775808}
    v3 += v3
    var v8 uint64 = 1
    switch v3 {
    case 9223372036854775807:
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint = 437
            var v11 uint32 = 2
            _ = fi9
            _ = v10
            _ = v11
        }
    case 1:
        var v12 bool = (0.1 > (float32(1e10) + 0.0))
        _ = v12
    }
    fmt.Println(v2, v2)
    fmt.Println(v8, v2)
    var v13 int8 = 100
    v13 += v13
    switch "foo bar" {
    case "":
        var v14 float32 = float32(float64((-2.25 - ((-2.25 * 0.1) / 1.5))))
        _ = v14
    case "a":
        fmt.Println(v13, v8)
    case "foo bar":
        var v15 bool = false
        _ = v15
    }
    v16 := GenS1{F0: "", F1: 661}
    var v17 float64 = -2.25
    if (v13 > (v13 / ((v13) | 1))) {
        for fi18 := 0; fi18 < 17; fi18++ {
            v2 = "\n\t"
            v19 := GenS1{F0: v2, F1: 597}
            _ = fi18
            _ = v19
        }
        var v20 bool = (255 < 966)
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
