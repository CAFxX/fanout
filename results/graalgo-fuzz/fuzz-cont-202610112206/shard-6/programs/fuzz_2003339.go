package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8 }

func main() {
    v1 := make(map[string]int)
    var v2 int = 2147483647
    var v3 uint16 = 42
    var v4 int16 = -149
    v5 := []string{"hello", "a", ""}
    v6 := make(map[uint32]string)
    switch (-(v2)) {
    case 42:
        var v7 string = "hello"
        _ = v7
    case -9223372036854775808:
        var v8 float32 = (3.14159 * 0.1)
        _ = v8
    case 127:
        var v9 int32 = 0
        _ = v9
    default:
        if false {
            var v10 uint8 = 1
            _ = v10
        }
    }
    v2 = (^(v2))
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    v3 = 1
    v11 := v1[("foo bar" + "")]
    v4 += (-(v4))
    v12 := GenS1{F0: -119}
    v2 = v11
    v4 -= int16(v3)
    v3 -= (v3 << ((uint(v3)) & 63))
    switch "foo bar" {
    case "foo bar":
        var v13 int = int(v4)
        _ = v13
    }
    var v14 int16 = (^(v4))
    v4 = (v4 & 0)
    if false {
        v5 = append(v5, "hello")
    }
    fmt.Println(v14, v11)
    if false {
        var v15 uint64 = 129
        if true {
            v2 += -287
            switch (v2 ^ (v2 ^ (v2 & v11))) {
            case 42:
                var v16 int8 = 0
                _ = v16
            case -772:
                v17 := GenS1{F0: -128}
                _ = v17
            case 0:
                v18 := GenS2{F0: 0, F1: 85}
                _ = v18
            }
            v3 = (^(v3))
        } else {
            var v19 int64 = 1
            _ = v19
        }
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
