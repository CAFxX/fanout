package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint32 }

func main() {
    v1 := []int{}
    var v2 uint = 0
    v3 := make(map[uint32]string)
    var v4 uint64 = 988
    switch "hello" {
    case "\n\t":
        v5 := GenS1{F0: 1, F1: 65535, F2: 0}
        _ = v5
    case "hello":
        var v6 uint16 = 807
        _ = v6
    case "foo bar":
        switch "hello" {
        case "\n\t":
            var v7 int32 = 83
            _ = v7
        case "foo bar":
            var v8 int = 128
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        _ = len(v1)
        _ = fi9
    }
    switch ((("" + "") + "a") + "\n\t") {
    case "a":
        v10 := GenS2{F0: 1, F1: 669}
        _ = v10
    case "":
        v4 = v4
    }
    var v11 int8 = 1
    var v12 bool = (2147483647 >= 565)
    var v13 string = "hello"
    var v14 uint32 = 4294967295
    var v15 float32 = (((1.5 * 0.1) - 0.0) * ((0.1 * 1.5) * (1.5 * 0.0)))
    if ((^(v2)) <= 42) {
        var v16 uint64 = v4
        v17 := GenS2{F0: 255, F1: (^(v14))}
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: v11, F1: v14, F2: 0}
    fmt.Println(v11, v14, v11)
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v11, v12)
        _ = fi19
    }
    v1 = append(v1, 42)
    v20 := GenS3{F0: v11, F1: (v14 << (((v2 + (^(v2)))) & 63))}
    var v21 uint = (v2 & v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
