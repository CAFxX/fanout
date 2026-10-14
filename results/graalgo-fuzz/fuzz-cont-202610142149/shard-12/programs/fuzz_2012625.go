package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 float64 = (-2.25 / 1e10)
    var v2 float32 = 0.0
    v3 := make(map[int]string)
    v2 += v2
    v1 = 0.0
    var v4 float64 = 0.1
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{1}
        _ = fi5
        _ = v6
    }
    var v7 int = 801
    var v8 uint32 = 65535
    if (((1 == -2147483648) || true) || false) {
        for fi9 := 0; fi9 < 11; fi9++ {
            v4 = (v1 + float64((v2 * -2.25)))
            _ = fi9
        }
    }
    v10 := []int{}
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: 20}
        fmt.Println(v1, v2)
        var v13 uint8 = 179
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := v3[v7]
    var v15 uint64 = 18446744073709551615
    for fi16 := 0; fi16 < 4; fi16++ {
        v14 = v14
        _ = fi16
    }
    switch v7 {
    case 32768:
        var v17 float64 = float64(v2)
        _ = v17
    case 42:
        _ = len(v10)
    }
    v14 = v14
    v7 += v7
    fmt.Println(v14, v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
