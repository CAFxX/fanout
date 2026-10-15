package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 uint32, p1 bool, p2 string) int16 {
    v5 := []int{2147483647}
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint8 = 1
        var v9 string = "foo bar"
        var v10 uint8 = 255
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -700
}

func main() {
    v1 := []int{-547, -1}
    var v2 bool = true
    v3 := make(map[string]int)
    v4 := []int{-2147483648, 4294967295, 663, 4294967295}
    v11 := GenS1{F0: -1, F1: 0, F2: 32767}
    v2 = v2
    var v12 int16 = 1000
    v13 := GenS2{F0: 154}
    v2 = (2 <= 127)
    v14 := GenS3{F0: -1, F1: 9223372036854775807}
    fmt.Println(v12, v2)
    _ = len(v3)
    switch 81 {
    case -1:
        v15 := GenS3{F0: 1, F1: 1}
        _ = v15
    }
    v16 := GenS2{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
