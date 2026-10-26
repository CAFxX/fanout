package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 2
    v3 := make(map[string]string)
    v4 := []int{42, int(v2)}
    var v5 int16 = 1
    fmt.Println(v5, v2)
    if ((v2 + 255) >= v2) {
        if true {
            v2 = (v2 & 0)
            var v6 int = 561
            _ = v6
        } else {
            var v7 int64 = 9223372036854775807
            v7 = (^(v7))
            _ = v7
        }
    }
    var v8 uint = 18446744073709551615
    var v9 uint = v8
    v8 += (^(v9))
    v10 := GenS1{F0: -984}
    v5 = (v5 & 1000)
    for fi11 := 0; fi11 < 9; fi11++ {
        _ = len(v4)
        v12 := GenS2{F0: fi11, F1: 242, F2: -485}
        fmt.Println(v5, v9, v2)
        _ = fi11
        _ = v12
    }
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
