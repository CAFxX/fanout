package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 string = "hello"
    var v2 int16 = 0
    v3 := make(map[int]int)
    var v4 int = 0
    v2 -= (v2 >> ((455) & 63))
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            v1 = (v1 + "a")
            _ = fi6
        }
        _ = fi5
    }
    var v7 int16 = -32768
    _ = len(v3)
    fmt.Println(v2, v7)
    var v8 int64 = 0
    var v9 string = "a"
    if true {
        v10 := GenS1{F0: 0}
        var v11 uint32 = 42
        _ = v10
        _ = v11
    } else {
        _ = len(v3)
        v12 := GenS1{F0: -1}
        v2 = (v7 - (-(v7)))
        _ = v12
    }
    var v13 uint64 = 900
    switch v9 {
    case "hello":
        v1 = "a"
    default:
        v2 = (v2 % ((v7) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
