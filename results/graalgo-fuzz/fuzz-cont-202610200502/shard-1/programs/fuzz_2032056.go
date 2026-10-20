package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 string = (("foo bar" + "foo bar") + "\n\t")
    v2 := make(map[string]uint32)
    v3 := []int{-9223372036854775808, 0, 128, 495}
    var v4 uint = 655
    var v5 float32 = (-2.25 + 1.5)
    if len(v3) > 0 {
        v3[0] = 536
    }
    var v6 int64 = -9223372036854775808
    var v7 uint64 = 42
    v3 = append(v3, 354)
    fmt.Println(v1, v7, v1)
    v6 = (v6 / ((42) | 1))
    var v8 float64 = (1e10 * -2.25)
    fmt.Println(v5, v4)
    switch 4294967295 {
    case -915:
        var v9 int32 = -520
        _ = v9
    case 966:
        var v10 int = 843
        _ = v10
    default:
        fmt.Println(v5, v5)
    }
    var v11 uint64 = 18446744073709551615
    for fi12 := 0; fi12 < 9; fi12++ {
        delete(v2, "a")
        v13 := GenS1{F0: v7}
        v11 = v7
        _ = fi12
        _ = v13
    }
    if (false && (!((!(!true)) || (false || true)))) {
        v14 := GenS1{F0: (^(v11))}
        _ = v14
    }
    var v15 uint = v4
    var v16 float64 = v8
    v1 = v1
    var v17 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
