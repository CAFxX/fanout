package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }

func main() {
    var v1 int = -428
    v2 := make(map[int]string)
    var v3 int = v1
    var v4 float32 = 0.0
    v5 := []int{(^(v3))}
    v6 := make(map[int]bool)
    switch 297 {
    case -2147483648:
        fmt.Println(v1, v3, v3)
    case 42:
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 float32 = v4
            _ = fi7
            _ = v8
        }
    default:
        v3 = (v1 << ((255) & 63))
    }
    v1 += v1
    fmt.Println(v3)
    for fi9 := 0; fi9 < 6; fi9++ {
        v2[v3] = (("foo bar" + "hello") + ("foo bar" + ""))
        v4 = v4
        v3 = v1
        _ = fi9
    }
    if false {
        var v10 uint8 = 1
        v10 = 128
        _ = v10
    } else {
        v4 -= v4
        v11 := GenS1{F0: 121, F1: 42}
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 float32 = v4
        v3 = fi12
        _ = len(v5)
        _ = fi12
        _ = v13
    }
    if true {
        v14 := GenS1{F0: -1, F1: -1}
        _ = v14
    } else {
        v15 := GenS1{F0: 42, F1: -9223372036854775808}
        var v16 int64 = 9223372036854775807
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 0, F1: 9223372036854775807}
    var v18 uint64 = 18446744073709551615
    var v19 uint64 = 608
    var v20 int = v3
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
