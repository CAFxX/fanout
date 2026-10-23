package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    var v1 int = -450
    var v2 uint8 = 128
    v3 := make(map[string]uint32)
    var v4 int32 = 1
    var v5 int32 = (v4 - -1)
    var v6 uint = 65535
    fmt.Println(v5, v4, v5)
    switch "" {
    case "":
        v6 = 680
    default:
        v7 := GenS1{F0: -1, F1: 127}
        _ = v7
    }
    var v8 uint64 = 4294967295
    if true {
        switch 9223372036854775807 {
        case 128:
            v2 = v2
        }
        for fi9 := 0; fi9 < 3; fi9++ {
            switch fi9 {
            case 0:
                var v10 int32 = (v5 % ((v5) | 1))
                _ = v10
            case -9223372036854775808:
                v11 := []int{-2147483648, fi9, fi9, fi9}
                _ = v11
            case -797:
                var v12 uint32 = 897
                _ = v12
            default:
                v13 := GenS2{F0: v2, F1: true}
                _ = v13
            }
            delete(v3, "foo bar")
            _ = fi9
        }
        v1 -= -9223372036854775808
    } else {
        v14 := []int{v1, v1, v1}
        var v15 bool = false
        _ = v14
        _ = v15
    }
    v16 := []int{128, (-(v1)), int(v8)}
    v5 -= v4
    var v17 float32 = -2.25
    v18 := GenS2{F0: 1, F1: true}
    v4 = (v5 >> (((v6 % ((v6) | 1))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
