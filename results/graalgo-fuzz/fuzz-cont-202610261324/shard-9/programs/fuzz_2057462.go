package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 int64 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int; F1 int16; F2 uint8 }

func main() {
    v1 := []uint32{1, 0}
    var v2 uint16 = 2
    v3 := make(map[int]string)
    v2 = v2
    v2 -= 1
    if (155 < 0) {
        var v4 int32 = 0
        v2 = v2
        _ = v4
    } else {
        var v5 uint8 = 255
        v6 := GenS1{F0: "hello", F1: 331}
        _ = v5
        _ = v6
    }
    v7 := v3[128]
    v8 := GenS2{F0: v2, F1: 1, F2: -255}
    v9 := GenS2{F0: v2, F1: -1, F2: 1}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    var v10 uint64 = 42
    switch -9223372036854775808 {
    case 127:
        v11 := GenS3{F0: uint(v2), F1: 1}
        _ = v11
    case -9223372036854775808:
        v7 = (v7 + "hello")
    case 2147483647:
        v2 += v2
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v2 -= (v2 + 1)
        var v13 uint32 = 4294967295
        _ = fi12
        _ = v13
    }
    v14 := GenS4{F0: -2147483648, F1: -1, F2: 0}
    if (686 < -32768) {
        var v15 string = "foo bar"
        var v16 int = 1
        _ = v15
        _ = v16
    }
    delete(v3, 453)
    var v17 string = v7
    fmt.Println(v7, v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
