package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int16 = -1
    v2 := make(map[int]uint32)
    var v3 uint64 = 2
    v4 := make(map[uint32]bool)
    v5 := []int{}
    var v6 uint = 18446744073709551615
    var v7 bool = (!("foo bar" == "foo bar"))
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int64 = 9223372036854775807
        _ = fi8
        _ = v9
    }
    v10 := v4[1]
    var v11 float32 = 1.5
    var v12 int16 = 420
    switch "a" {
    case "hello":
        switch "a" {
        case "hello":
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := GenS1{F0: int32(v6)}
                _ = fi13
                _ = v14
            }
        default:
            fmt.Println(v12, v11)
        }
    case "foo bar":
        var v15 float32 = ((v11 / v11) * 0.1)
        _ = v15
    case "":
        var v16 uint = (v6 | v6)
        _ = v16
    }
    var v17 int16 = 618
    if v7 {
        v7 = v7
        v1 = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
