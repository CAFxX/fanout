package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int = -822
    v2 := make(map[uint32]string)
    var v3 int32 = 0
    var v4 string = "foo bar"
    fmt.Println(v1)
    switch (("\n\t" + v4) + v4) {
    case "\n\t":
        var v5 int64 = 0
        _ = v5
    }
    v1 += (v1 << ((65535) & 63))
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 float32 = -2.25
        _ = fi6
        _ = v7
    }
    v8 := []int{127, v1, (v1 | (^(v1)))}
    v1 = v1
    switch 32768 {
    case 32767:
        var v9 uint8 = 42
        _ = v9
    default:
        _ = len(v8)
    }
    delete(v2, 2147483648)
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 float32 = float32(0.1)
        _ = fi10
        _ = v11
    }
    if true {
        var v12 uint8 = 201
        _ = v12
    } else {
        var v13 uint64 = 493
        _ = v13
    }
    v14 := GenS1{F0: 0, F1: (v3 / (((^(v3))) | 1)), F2: 604}
    v15 := GenS2{F0: false}
    var v16 uint64 = 755
    v17 := GenS3{F0: 1}
    fmt.Println(v16, v4)
    fmt.Println(v3, v3)
    var v18 uint8 = 54
    v2[707] = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
