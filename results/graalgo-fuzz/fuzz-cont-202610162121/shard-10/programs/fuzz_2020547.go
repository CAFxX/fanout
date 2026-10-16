package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint32 }

func main() {
    var v1 int32 = 2147483647
    var v2 uint = 262
    v3 := make(map[string]int)
    v4 := []uint32{2147483647, 4294967295, 1}
    var v5 int64 = 9223372036854775807
    switch 42 {
    case 662:
        v6 := GenS1{F0: 32, F1: "a", F2: 4294967295}
        _ = v6
    case -777:
        var v7 uint = (v2 + 968)
        _ = v7
    }
    v3[("foo bar" + "\n\t")] = 4294967295
    var v8 int32 = (v1 | v1)
    var v9 uint64 = 0
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(fi10, v2, v1)
        v8 -= (v1 - (^(v1)))
        _ = fi10
    }
    v8 += v1
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := v3[("foo bar" + (("hello" + "") + "hello"))]
        v13 := GenS1{F0: 216, F1: "\n\t", F2: 2}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = len(v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 uint64 = v9
    for fi15 := 0; fi15 < 9; fi15++ {
        v1 += v8
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
