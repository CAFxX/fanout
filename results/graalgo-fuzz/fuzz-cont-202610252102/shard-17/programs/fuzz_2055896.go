package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int32; F2 int64 }
func genFn0() int32 {
    var v5 int64 = -652
    fmt.Println(v5)
    var v6 uint16 = 65535
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: -1}
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return 538
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    var v3 uint64 = 2
    v4 := []string{}
    var v9 float32 = 3.14159
    var v10 uint8 = 255
    fmt.Println(v9, v9, v3)
    delete(v2, 834)
    if ((0 < 65535) && (2 > 491)) {
        v9 = v9
    }
    v9 = (0.1 + v9)
    if true {
        var v11 int = 2147483647
        _ = v11
    } else {
        var v12 uint8 = (v10 - 128)
        if (821 < 0) {
            fmt.Println(v10, v3)
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS2{F0: "hello", F1: -9223372036854775808}
            v15 := GenS2{F0: "a", F1: 0}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    }
    var v16 uint = 42
    v17 := GenS3{F0: v10}
    if (-83 == 0) {
        switch "a" {
        case "a":
            var v18 int32 = 2147483647
            _ = v18
        case "":
            var v19 uint32 = 4294967295
            _ = v19
        default:
            v9 = v9
        }
        v20 := GenS4{F0: "hello", F1: 2147483647, F2: 817}
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 float32 = (((3.14159 - v9) / v9) / v9)
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    v10 += (^(v10))
    var v23 uint32 = 65535
    var v24 float32 = v9
    v16 += (^(v16))
    var v25 int32 = -492
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
