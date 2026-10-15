package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
func genFn0() int64 {
    var v5 uint8 = 0
    v6 := GenS1{F0: -1, F1: 32767}
    v5 = 2
    v5 = 1
    v5 -= (v5 % (((^(v5))) | 1))
    _ = v5
    _ = v6
    return 818
}
func genFn1(p0 uint8) uint32 {
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 -= p0
        _ = fi7
    }
    if (true && ("foo bar" == "hello")) {
        v8 := []int{0}
        _ = v8
    }
    p0 = 2
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(fi9, p0)
        v10 := GenS1{F0: 2147483647, F1: 1000}
        var v11 float64 = 1.5
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    return 42
}

func main() {
    var v1 uint = 1
    var v2 int32 = -12
    var v3 string = ("" + "a")
    v4 := make(map[string]bool)
    v12 := []int{75, -1, 195, 128}
    v3 = "\n\t"
    fmt.Println(v3, v3)
    var v13 float32 = 3.14159
    switch 309 {
    case 4294967295:
        fmt.Println(v2, v2)
    }
    v14 := GenS1{F0: v2, F1: 1}
    for fi15 := 0; fi15 < 19; fi15++ {
        v13 += v13
        _ = fi15
    }
    v16 := GenS1{F0: v2, F1: 32767}
    var v17 int = 593
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
