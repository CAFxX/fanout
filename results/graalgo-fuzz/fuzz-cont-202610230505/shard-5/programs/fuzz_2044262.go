package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint64; F2 int64 }
func genFn0(p0 int16, p1 int32) uint64 {
    v5 := []int{}
    switch ("\n\t" + "foo bar") {
    case "a":
        fmt.Println(p0)
    case "\n\t":
        p0 -= p0
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: (!(fi6 <= fi6))}
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: -2147483648}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 2
}

func main() {
    var v1 int16 = -895
    var v2 int64 = 610
    var v3 int16 = v1
    var v4 int32 = 1
    fmt.Println(v1, v3, v4)
    v9 := []int{4294967295}
    if ((v1 & v3) > (v3 / ((-1) | 1))) {
        _ = len(v9)
    }
    fmt.Println(v4)
    for fi10 := 0; fi10 < 8; fi10++ {
        v4 = (v4 << ((1) & 63))
        _ = fi10
    }
    var v11 int32 = 0
    v12 := GenS3{F0: 9223372036854775807, F1: 337, F2: v2}
    var v13 float32 = 0.1
    v11 = (v4 * 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
