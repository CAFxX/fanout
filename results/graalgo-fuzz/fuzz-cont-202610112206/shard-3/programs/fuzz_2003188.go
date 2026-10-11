package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 int = 2147483647
    var v3 uint = 432
    var v4 int32 = 1
    v5 := make(map[int]string)
    var v6 int64 = 1
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: v6}
        _ = fi7
        _ = v8
    }
    if true {
        var v9 int64 = (v6 & (v6 + (v6 | v6)))
        v10 := GenS2{F0: 42, F1: 65535, F2: 392}
        fmt.Println(v2, v2)
        _ = v10
        _ = v9
    }
    var v11 uint32 = 1
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := []int{fi12, (-(v2)), (v2 + (v2 >> ((v3) & 63))), fi12}
        v3 += v3
        _ = fi12
        _ = v13
    }
    _ = len(v5)
    fmt.Println(v4)
    var v14 uint16 = 2
    switch "a" {
    case "a":
        v2 = v2
    case "foo bar":
        switch "foo bar" {
        case "a":
            v15 := GenS1{F0: (-(v6))}
            _ = v15
        default:
            v4 = (^(v4))
        }
    default:
        v11 += v11
    }
    v14 = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
