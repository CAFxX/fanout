package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 string) uint32 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 5; fi4++ {
        p0 = "\n\t"
        _ = fi4
    }
    var v5 int32 = 429
    _ = p0
    _ = v5
    return 2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -1
    v3 := make(map[string]bool)
    fmt.Println(v2)
    v6 := []int{}
    v1[("\n\t" + ("a" + "\n\t"))] = (!false)
    fmt.Println(v2)
    v2 += -81
    fmt.Println(v2)
    fmt.Println(v2)
    v2 -= (v2 << ((42) & 63))
    v7 := GenS1{F0: 100}
    var v8 int32 = v2
    var v9 float64 = (float64(0.0) + 0.1)
    fmt.Println(v2, v2)
    var v10 uint8 = 0
    switch "hello" {
    case "a":
        var v11 int64 = 0
        _ = v11
    case "hello":
        fmt.Println(v2)
    default:
        v12 := v3[("hello" + (("\n\t" + "a") + "\n\t"))]
        _ = v12
    }
    v2 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
