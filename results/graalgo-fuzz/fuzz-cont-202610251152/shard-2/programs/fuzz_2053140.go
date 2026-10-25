package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }

func main() {
    v1 := []uint32{}
    var v2 bool = (true || (!("" > "a")))
    var v3 int32 = 42
    var v4 uint = 0
    v5 := []int{586, 2147483647, 128, 9223372036854775807}
    var v6 uint8 = 2
    var v7 int8 = 42
    v8 := GenS1{F0: -1, F1: -520, F2: 0}
    v3 -= v3
    v7 -= v7
    if false {
        var v9 uint64 = 104
        fmt.Println(v2, v6, v4)
        fmt.Println(v6)
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int = (fi10 | fi10)
        _ = fi10
        _ = v11
    }
    switch "foo bar" {
    case "":
        _ = len(v5)
    case "hello":
        fmt.Println(v6)
    case "a":
        switch 32767 {
        case 32768:
            var v12 int = 2147483647
            _ = v12
        case 729:
            v4 = (v4 >> ((1) & 63))
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        var v13 string = "a"
        _ = v13
    }
    fmt.Println(v4)
    var v14 int = 2147483647
    v14 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
