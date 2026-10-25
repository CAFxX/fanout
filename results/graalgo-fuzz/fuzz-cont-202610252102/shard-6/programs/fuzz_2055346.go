package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 int64 = -1
    for fi4 := 0; fi4 < 4; fi4++ {
        v3 -= (v3 | 42)
        _ = fi4
    }
    var v5 string = ""
    var v6 int32 = 443
    v7 := GenS1{F0: v6, F1: 2147483647, F2: 42}
    for fi8 := 0; fi8 < 17; fi8++ {
        _ = len(v1)
        v9 := GenS2{F0: 100, F1: 491}
        _ = fi8
        _ = v9
    }
    var v10 int8 = 46
    var v11 bool = (v3 == -1)
    v10 = (v10 * v10)
    var v12 uint = 0
    v6 = v6
    if (!v11) {
        v13 := []int{9223372036854775807}
        _ = v13
    } else {
        v5 = "\n\t"
        v14 := GenS3{F0: -1, F1: 4}
        v10 = (v10 * (^(v10)))
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
