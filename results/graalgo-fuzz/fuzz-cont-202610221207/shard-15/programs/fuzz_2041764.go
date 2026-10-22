package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint8 }
func genFn0(p0 int32, p1 uint8) uint32 {
    v5 := []int{32768, -2147483648}
    p0 += p0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = v5
    return 752
}

func main() {
    var v1 int16 = 42
    var v2 int64 = 42
    var v3 int64 = -554
    v4 := []string{"\n\t", "hello", "\n\t"}
    var v6 uint64 = 18446744073709551615
    v3 -= v3
    var v7 uint16 = 1
    v8 := GenS1{F0: true, F1: uint(v6), F2: 1}
    fmt.Println(v6, v6)
    switch 2147483647 {
    case 32768:
        for fi9 := 0; fi9 < 13; fi9++ {
            fmt.Println(v1, v3)
            _ = fi9
        }
    case 32767:
        _ = len(v4)
    case 322:
        fmt.Println(v6, v7)
    }
    v6 -= (v6 << ((853) & 63))
    switch -9223372036854775808 {
    case 32768:
        v6 += (v6 - v6)
    case 32767:
        v3 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
