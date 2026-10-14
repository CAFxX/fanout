package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int }
type GenS2 struct { F0 int16; F1 int16; F2 uint16 }

func main() {
    var v1 float64 = 0.0
    var v2 string = ""
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    v5 := []int{-7, -216, 4294967295, 32767}
    v6 := GenS1{F0: 42, F1: 2147483647, F2: 9223372036854775807}
    v7 := GenS1{F0: 2147483647, F1: 440, F2: -1}
    v8 := GenS1{F0: 650, F1: 367, F2: 1}
    v9 := GenS2{F0: -32768, F1: 1, F2: 999}
    fmt.Println(v1)
    v2 = ("\n\t" + ("foo bar" + "foo bar"))
    fmt.Println(v1)
    var v10 int = 32768
    switch "hello" {
    case "foo bar":
        var v11 float64 = float64(((float32((3.14159 * v1)) + 1.5) * float32(v1)))
        _ = v11
    case "":
        var v12 int64 = -9223372036854775808
        _ = v12
    default:
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint64 = 18446744073709551615
            _ = fi13
            _ = v14
        }
    }
    v15 := GenS1{F0: -1, F1: 65535, F2: (v10 & 9223372036854775807)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
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
