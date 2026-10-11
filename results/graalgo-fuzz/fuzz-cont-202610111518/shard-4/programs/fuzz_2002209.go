package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int8 = 52
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 48
    v5 := make(map[uint32]string)
    var v6 uint64 = 4294967295
    var v7 int64 = 341
    var v8 uint16 = 1
    v9 := GenS1{F0: 2147483647, F1: (v1 ^ 100)}
    v10 := GenS2{F0: int64(v8)}
    var v11 uint64 = (^(v6))
    switch "foo bar" {
    case "":
        v1 = v1
    case "hello":
        var v12 int8 = v1
        _ = v12
    }
    var v13 uint8 = 42
    var v14 string = ""
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS3{F0: 9223372036854775807, F1: 32767}
            var v18 float32 = ((float32(-2.25) / (float32(0.0) + (((0.0 - 0.1) / 1.0) + (float32(1e10) * -2.25)))) * 1e10)
            var v19 int = fi15
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
