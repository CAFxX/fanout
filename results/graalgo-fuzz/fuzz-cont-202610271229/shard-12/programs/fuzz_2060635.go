package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32, p1 string, p2 int) uint32 {
    var v5 uint16 = 255
    p2 -= (-(p2))
    fmt.Println(p2, p0, v5)
    fmt.Println(p0, p2, v5)
    v5 = (^(v5))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 950
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    v6 := GenS1{F0: 128}
    v7 := []int{85, -1, -1, 128}
    switch "foo bar" {
    case "a":
        switch ("foo bar" + "hello") {
        case "":
            if (-2147483648 > -9223372036854775808) {
                v8 := GenS1{F0: 12}
                fmt.Println(v1)
                fmt.Println(v1)
                _ = v8
            }
        case "\n\t":
            v1 = 28
        }
    case "foo bar":
        v9 := v2["a"]
        _ = v9
    default:
        fmt.Println(v1)
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v10 := GenS2{F0: 32767}
    var v11 int64 = 9223372036854775807
    var v12 int = -9223372036854775808
    if len(v7) > 0 {
        v7[0] = v12
    }
    v1 = v1
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := v2[""]
        _ = fi13
        _ = v14
    }
    v1 = (v1 << ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
