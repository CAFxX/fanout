package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }

func main() {
    var v1 bool = ((-2.25 - 1.5) >= float32(((-2.25 + ((1.5 * -2.25) * 1.5)) + (1e10 * 1e10))))
    v2 := []int{352, 32767}
    var v3 float64 = (float64(((3.14159 / (3.14159 / 1.0)) - 0.0)) * 3.14159)
    var v4 int = 396
    v5 := make(map[int]string)
    v6 := []uint32{2, 989, 4294967295, 2}
    v7 := GenS1{F0: (v4 + v4), F1: 42}
    v5[v4] = "hello"
    var v8 int = v4
    v9 := GenS1{F0: -778, F1: -1}
    v5[v4] = "foo bar"
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS2{F0: -445}
        var v12 string = (("foo bar" + "a") + "hello")
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int64 = 1
    switch "\n\t" {
    case "\n\t":
        v4 = v8
    case "hello":
        switch "" {
        case "hello":
            v14 := GenS3{F0: -15}
            _ = v14
        }
    default:
        v1 = (!((1.5 - float64(-2.25)) >= v3))
    }
    var v15 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
