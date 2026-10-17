package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int64, p2 int64) uint {
    var v5 int = 32767
    v6 := GenS1{F0: (("hello" + "foo bar") + "hello"), F1: int64(p0)}
    p0 = 576
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    var v3 uint32 = 42
    var v4 uint = 401
    var v7 uint64 = 4294967295
    var v8 uint32 = (v3 & 2)
    fmt.Println(v4)
    v9 := GenS2{F0: "foo bar", F1: (v4 ^ 65535)}
    if (((1 == 42) || true) && false) {
        _ = len(v1)
        v4 += (v4 | (^(v4)))
    }
    v10 := GenS3{F0: 813}
    v3 -= v3
    switch 32768 {
    case 128:
        v7 -= (^(v7))
    default:
        var v11 uint = uint(v7)
        _ = v11
    }
    v3 = (v8 ^ v8)
    v12 := []int{645}
    v1[4294967295] = 1
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint = (v4 & 599)
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 int16 = 42
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
