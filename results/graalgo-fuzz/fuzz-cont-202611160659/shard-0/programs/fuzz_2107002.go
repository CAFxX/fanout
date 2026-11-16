package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 string, p1 bool, p2 int) uint32 {
    p2 = 2147483647
    fmt.Println(p2, p1)
    fmt.Println(p2)
    switch (p2 ^ (p2 | -395)) {
    case -651:
        fmt.Println(p0, p0)
    case 393:
        p0 = p0
    case 479:
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 int32 = 2147483647
    v2 := []int{-2147483648}
    var v3 uint64 = 614
    v1 = (v1 | -79)
    if (!false) {
        var v4 uint16 = uint16(v3)
        if (!(109 >= 2)) {
            v5 := GenS1{F0: 42}
            v4 += (v4 / ((v4) | 1))
            _ = v5
        }
        v3 = (^(v3))
        _ = v4
    }
    v2 = append(v2, 42)
    var v6 int64 = 9223372036854775807
    v1 = (v1 << ((4294967295) & 63))
    var v7 int = 0
    v6 = (v6 >> ((807) & 63))
    v2 = append(v2, (v7 & 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
