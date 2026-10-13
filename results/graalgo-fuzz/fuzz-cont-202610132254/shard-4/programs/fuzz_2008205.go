package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]uint32)
    v3 := []uint32{2147483648}
    var v4 float32 = ((0.1 * ((float32((1.5 / (3.14159 / 1.0))) - -2.25) / 1.0)) - (((v1 - float32(1e10)) - (v1 + 1e10)) / (1e10 - v1)))
    v4 += v4
    for fi5 := 0; fi5 < 7; fi5++ {
        v4 = v4
        switch (^(fi5)) {
        case -1:
            v3 = append(v3, 272)
        }
        _ = fi5
    }
    if (true || true) {
        if true {
            fmt.Println(v1, v4)
        }
        v1 = ((v1 - v1) * 1e10)
    }
    var v6 uint64 = 18446744073709551615
    _ = len(v3)
    switch 127 {
    case -9223372036854775808:
        fmt.Println(v1, v4, v6)
    case 628:
        var v7 int32 = 763
        _ = v7
    }
    v2[4294967295] = 0
    v6 -= (^(v6))
    _ = len(v3)
    var v8 uint64 = v6
    v9 := GenS1{F0: 2147483647, F1: 2147483647}
    if ((!true) || (623 == -883)) {
        v10 := v2[-9223372036854775808]
        v11 := GenS1{F0: v10, F1: (v10 - v10)}
        var v12 uint16 = 0
        _ = v10
        _ = v11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
