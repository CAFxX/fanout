package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 0
    v3 := make(map[int]int)
    var v4 int64 = v2
    var v5 uint8 = 42
    var v6 uint64 = 18446744073709551615
    v7 := GenS1{F0: (-(v4))}
    var v8 string = "hello"
    for fi9 := 0; fi9 < 15; fi9++ {
        delete(v1, fi9)
        _ = fi9
    }
    v4 += (v4 - (v2 ^ (v2 % (((v4 % ((v4) | 1))) | 1))))
    switch v8 {
    case "foo bar":
        v10 := []int{}
        _ = v10
    case "hello":
        v4 -= v4
    default:
        var v11 uint16 = 32767
        _ = v11
    }
    var v12 uint16 = 32767
    v5 = v5
    switch v8 {
    case "\n\t":
        v13 := GenS1{F0: v4}
        _ = v13
    case "foo bar":
        v14 := GenS1{F0: 1}
        _ = v14
    case "a":
        if true {
            v2 -= 0
        }
    default:
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(v4)
            v6 = v6
            v4 -= -343
            _ = fi15
        }
    }
    var v16 uint32 = 2147483647
    var v17 uint32 = (v16 & v16)
    v5 = (v5 % ((v5) | 1))
    fmt.Println(v2, v6)
    var v18 int8 = -128
    v19 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
