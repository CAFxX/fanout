package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS2 struct { F0 bool; F1 int; F2 uint8 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]bool)
    v3 := []int{127, 9223372036854775807}
    v4 := make(map[int]string)
    switch 1 {
    case 32767:
        for fi5 := 0; fi5 < 7; fi5++ {
            v1 -= 2147483647
            v6 := GenS1{F0: 2, F1: 2, F2: 11}
            _ = fi5
            _ = v6
        }
    case -2147483648:
        v1 += (^(v1))
    }
    var v7 int = -1
    v8 := GenS2{F0: (false && true), F1: v7, F2: 34}
    if false {
        fmt.Println(v1)
    }
    _ = len(v3)
    v1 = (v1 << ((230) & 63))
    var v9 int = 32768
    delete(v4, (-(v9)))
    for fi10 := 0; fi10 < 16; fi10++ {
        v9 -= fi10
        fmt.Println(v1)
        var v11 int64 = 42
        _ = fi10
        _ = v11
    }
    _ = len(v3)
    v9 -= (v9 << ((42) & 63))
    if true {
        var v12 uint32 = (v1 >> ((1) & 63))
        for fi13 := 0; fi13 < 15; fi13++ {
            switch ("hello" + ((("hello" + "") + "a") + "hello")) {
            case "\n\t":
                var v14 int64 = -857
                _ = v14
            case "foo bar":
                v7 = fi13
            default:
                _ = len(v4)
            }
            _ = fi13
        }
        v15 := v4[v9]
        _ = v12
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
