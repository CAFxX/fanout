package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int32; F2 int64 }

func main() {
    var v1 int32 = -99
    v2 := make(map[int]string)
    var v3 int8 = 0
    v4 := make(map[int]int)
    var v5 int16 = 1000
    var v6 string = "a"
    var v7 int8 = v3
    v3 = v3
    fmt.Println(v7)
    v8 := []int{}
    switch 1 {
    case -994:
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(v5, fi9, fi9)
            v3 = (v3 ^ (v7 & v3))
            _ = fi9
        }
    case -9223372036854775808:
        fmt.Println(v7)
    case 2147483647:
        v3 = 55
    }
    var v10 uint16 = 42
    v11 := GenS1{F0: v3}
    for fi12 := 0; fi12 < 4; fi12++ {
        if ((^(fi12)) >= (-(fi12))) {
            var v13 int32 = v1
            v5 = 32767
            _ = v13
        } else {
            v14 := GenS1{F0: -59}
            _ = v14
        }
        var v15 bool = (4294967295 >= 2147483648)
        _ = fi12
        _ = v15
    }
    v16 := GenS2{F0: 512, F1: v1, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
