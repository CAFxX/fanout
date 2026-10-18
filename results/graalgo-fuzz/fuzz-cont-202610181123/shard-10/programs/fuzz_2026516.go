package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []uint32{1, 65535}
    v2 := make(map[int]string)
    var v3 uint = 0
    var v4 uint = (v3 | (v3 + v3))
    v5 := make(map[int]int)
    v6 := make(map[string]int)
    _ = len(v1)
    fmt.Println(v4, v4)
    switch -2147483648 {
    case -742:
        v3 += v3
    case -1:
        fmt.Println(v3)
    case 644:
        v7 := GenS1{F0: 244}
        _ = v7
    default:
        v3 = v4
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v4 = 255
        _ = fi8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        switch "foo bar" {
        case "hello":
            v4 = (v4 | 493)
        }
        _ = fi9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v3 -= (v3 / ((v4) | 1))
        _ = fi10
    }
    var v11 string = "a"
    var v12 int32 = -413
    delete(v6, v11)
    fmt.Println(v4, v12, v12)
    v13 := GenS1{F0: 1}
    var v14 int = 32768
    if false {
        fmt.Println(v12, v12, v12)
        for fi15 := 0; fi15 < 5; fi15++ {
            v3 -= v3
            _ = fi15
        }
        v11 = v11
    }
    v3 = (v4 << ((v4) & 63))
    fmt.Println(v3, v3, v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
