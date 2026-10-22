package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = ((1.5 / (-2.25 - 1e10)) + 1.5)
    var v3 uint16 = 1
    var v4 int64 = 0
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := []int{fi6}
            v7 = append(v7, (fi6 - fi6))
            v8 := GenS1{F0: 369}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 int32 = -2147483648
        _ = fi5
        _ = v9
    }
    switch (("\n\t" + "foo bar") + "hello") {
    case "foo bar":
        var v10 uint32 = 65535
        _ = v10
    case "a":
        v2 -= v2
    case "\n\t":
        fmt.Println(v4, v2)
    default:
        var v11 string = "foo bar"
        _ = v11
    }
    delete(v1, "\n\t")
    var v12 int8 = 1
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int16 = -1
        var v15 uint64 = 1
        v2 += ((0.0 + (((0.0 - float32(1.5)) + v2) * v2)) / (v2 * 1e10))
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 = 42
    v3 += v3
    v3 += v3
    v12 -= v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
