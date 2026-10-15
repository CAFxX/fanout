package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int64 }

func main() {
    v1 := []string{("" + "\n\t")}
    var v2 bool = false
    v3 := make(map[int]uint32)
    var v4 string = ""
    v5 := []string{"", ("\n\t" + (v4 + (v4 + "a"))), v4}
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint8 = 42
            v9 -= 0
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: 42}
    v11 := GenS1{F0: -128}
    v4 = "\n\t"
    fmt.Println(v4)
    var v12 uint = 646
    switch -459 {
    case 42:
        v4 = "foo bar"
    case -2147483648:
        v12 += v12
    default:
        if v2 {
            v4 = v4
            fmt.Println(v4)
        }
    }
    v2 = v2
    if (65535 < 255) {
        v12 = 4294967295
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS2{F0: 2147483648, F1: 944, F2: 281}
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 2
    fmt.Println(v12, v15, v4)
    v16 := GenS1{F0: 106}
    fmt.Println(v12, v2)
    fmt.Println(v15)
    v17 := GenS2{F0: 0, F1: 4294967295, F2: 0}
    v18 := GenS1{F0: 42}
    v6[1] = 825
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
