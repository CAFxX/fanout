package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0(p0 uint64, p1 uint, p2 uint16) int32 {
    p0 += 0
    p2 = (p2 >> ((475) & 63))
    _ = p0
    _ = p1
    _ = p2
    return -792
}

func main() {
    var v1 uint = 255
    var v2 int32 = 2147483647
    var v3 int8 = 127
    v4 := make(map[uint32]int)
    var v5 int64 = 551
    var v6 uint16 = 255
    var v7 int32 = 0
    fmt.Println(v1)
    v2 = (v7 - genFn0(42, 930, v6))
    var v8 uint = v1
    if (!("foo bar" < "\n\t")) {
        v5 = -473
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v4[2] = 1
        var v10 int32 = v2
        v11 := GenS1{F0: 255}
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 0, F1: 589, F2: v8}
    v6 += (v6 * (v6 * v6))
    v8 = v1
    var v13 string = "foo bar"
    v1 -= 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
