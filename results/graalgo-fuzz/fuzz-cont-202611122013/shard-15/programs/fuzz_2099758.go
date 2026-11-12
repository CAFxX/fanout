package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int32 }
func genFn0() uint64 {
    v7 := []int{}
    if ((9223372036854775807 <= -1) || false) {
        var v8 string = "foo bar"
        v8 = ""
        v8 = "a"
        _ = v8
    }
    var v9 int8 = 100
    _ = v7
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 int64 = 1
    var v2 uint64 = 42
    var v3 uint32 = 2
    v4 := []int{}
    v5 := []uint32{v3, (v3 >> ((uint(v2)) & 63)), (v3 - 233), 2147483647}
    v6 := make(map[uint32]string)
    fmt.Println(v1)
    for fi10 := 0; fi10 < 4; fi10++ {
        if true {
            fmt.Println(v2, v2)
        }
        var v11 string = "foo bar"
        _ = fi10
        _ = v11
    }
    switch "\n\t" {
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = (v3 >> ((65535) & 63))
        }
    case "a":
        v2 = v2
    default:
        delete(v6, 1)
    }
    v3 = (v3 | 65535)
    _ = len(v6)
    v1 = 42
    v12 := GenS1{F0: 1, F1: 127, F2: 2147483647}
    v4 = append(v4, -409)
    if len(v4) > 0 {
        v4[0] = 127
    }
    fmt.Println(v2, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
