package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint8 }

func main() {
    v1 := []uint32{163, 0, 881}
    var v2 int32 = -726
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    var v5 string = ""
    var v6 int32 = (v2 | 920)
    v6 = 479
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint = 0
        fmt.Println(fi7, v6, v2)
        _ = fi7
        _ = v8
    }
    switch v5 {
    case "foo bar":
        fmt.Println(v2, v2)
    }
    switch 32767 {
    case -584:
        if false {
            v9 := GenS1{F0: (v5 + "\n\t"), F1: 136, F2: v2}
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := v4[(v5 + (v5 + "\n\t"))]
                v12 := GenS2{F0: 167, F1: (!v11), F2: 1}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        }
    case 32767:
        var v13 uint = 273
        _ = v13
    case 94:
        fmt.Println(v2, v2)
    }
    v5 = "\n\t"
    fmt.Println(v2, v5, v2)
    v6 += -1
    for fi14 := 0; fi14 < 10; fi14++ {
        if (!(((-9223372036854775808 != int64(v2)) && true) || false)) {
            if false {
                var v15 int = fi14
                v1 = append(v1, 0)
                _ = v15
            }
            if true {
                var v16 float64 = 1e10
                _ = v16
            }
        }
        _ = fi14
    }
    fmt.Println(v5, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
