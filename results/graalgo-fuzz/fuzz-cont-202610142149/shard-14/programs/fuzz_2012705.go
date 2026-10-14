package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }

func main() {
    var v1 float64 = (1e10 / 3.14159)
    v2 := []string{}
    var v3 uint32 = 170
    var v4 string = "hello"
    var v5 int32 = 42
    v6 := make(map[uint32]uint32)
    fmt.Println(v4, v4, v1)
    var v7 string = v4
    fmt.Println(v1, v4, v5)
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 148}
        var v10 uint = 898
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int = -201
    v12 := GenS2{F0: -32768, F1: v11, F2: -889}
    if ((!false) && (false || true)) {
        var v13 float32 = (3.14159 + 0.0)
        v3 = v3
        _ = v13
    } else {
        if true {
            v3 -= (v3 | v3)
            var v14 int16 = 0
            var v15 int = v11
            _ = v14
            _ = v15
        }
        if (!true) {
            var v16 int16 = 42
            _ = v16
        }
        switch (v11 - (v11 ^ v11)) {
        case 0:
            v4 = "foo bar"
        }
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
