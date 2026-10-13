package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }

func main() {
    var v1 uint64 = 42
    var v2 uint32 = 65535
    var v3 int16 = -1
    v4 := []uint32{v2, 1, v2, (v2 << ((655) & 63))}
    var v5 uint16 = 255
    var v6 uint16 = 398
    var v7 uint64 = 1
    v1 -= (v7 + 614)
    var v8 float64 = (float64(0.1) - 0.0)
    fmt.Println(v2, v1, v2)
    v9 := GenS1{F0: -32768, F1: 9223372036854775807}
    v2 -= v2
    v6 = (^(v5))
    v6 -= v6
    if (true || false) {
        v4 = append(v4, 1)
        v10 := GenS2{F0: -333}
        _ = v10
    } else {
        v11 := GenS3{F0: 128}
        _ = v11
    }
    _ = len(v4)
    for fi12 := 0; fi12 < 10; fi12++ {
        v8 -= 0.0
        if false {
            var v13 string = "hello"
            _ = v13
        } else {
            v14 := GenS4{F0: ""}
            _ = v14
        }
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
