package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int8 }
type GenS2 struct { F0 int8; F1 string }

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    var v3 uint32 = 742
    v4 := make(map[uint32]bool)
    v5 := []string{"hello"}
    v6 := []string{}
    switch 198 {
    case 127:
        for fi7 := 0; fi7 < 7; fi7++ {
            v2 = false
            _ = fi7
        }
    case -1:
        var v8 uint64 = 429
        _ = v8
    case 42:
        fmt.Println(v2, v2)
    default:
        v9 := GenS1{F0: 630, F1: 18446744073709551615, F2: 127}
        _ = v9
    }
    fmt.Println(v2, v3)
    var v10 uint64 = 389
    v2 = false
    var v11 float32 = 0.1
    fmt.Println(v10, v3, v10)
    v12 := GenS1{F0: -481, F1: 2, F2: int8(v10)}
    v3 = v3
    var v13 uint = 1
    var v14 bool = v2
    fmt.Println(v3, v14, v13)
    fmt.Println(v14, v2)
    v15 := GenS2{F0: 106, F1: ("\n\t" + "")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
