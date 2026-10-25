package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0(p0 int16, p1 int, p2 int) uint64 {
    var v7 float32 = (0.0 - 0.1)
    var v8 int32 = 2147483647
    v9 := GenS1{F0: ("" + "foo bar"), F1: int32(p0)}
    var v10 uint16 = 650
    v11 := []int{(^(p1)), (-(p1)), -802}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 1
    var v3 int = 573
    v4 := make(map[int]string)
    var v5 uint16 = 0
    v6 := make(map[string]string)
    fmt.Println(v5, v3, v5)
    if (44 != -13) {
        v3 = int(v2)
        v2 = (v2 / (((v2 - (v2 + v2))) | 1))
    }
    v5 = v5
    for fi12 := 0; fi12 < 18; fi12++ {
        switch "" {
        case "\n\t":
            fmt.Println(v2, fi12, v5)
        case "foo bar":
            v13 := GenS2{F0: 42, F1: 127}
            _ = v13
        }
        v1[287] = "hello"
        _ = fi12
    }
    v3 = (^(v3))
    var v14 int64 = (^(v2))
    if true {
        v2 = -447
        fmt.Println(v2)
    }
    fmt.Println(v2, v2, v2)
    var v15 uint8 = 42
    var v16 uint16 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
