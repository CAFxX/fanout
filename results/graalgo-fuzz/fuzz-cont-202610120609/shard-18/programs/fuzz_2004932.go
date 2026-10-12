package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint8) bool {
    p0 = p0
    var v6 uint32 = uint32(p0)
    _ = p0
    _ = v6
    return true
}

func main() {
    var v1 int64 = -412
    var v2 uint16 = 0
    var v3 string = ""
    v4 := make(map[uint32]string)
    var v5 int32 = 0
    var v7 string = v3
    if (-65 < int8(v2)) {
        v1 -= v1
        for fi8 := 0; fi8 < 19; fi8++ {
            switch fi8 {
            case 0:
                var v9 int = -9223372036854775808
                _ = v9
            case -2147483648:
                v1 -= (v1 >> ((4294967295) & 63))
            }
            _ = fi8
        }
    } else {
        var v10 uint8 = 105
        _ = v10
    }
    v11 := []int{-456, 128, 9223372036854775807}
    v12 := GenS1{F0: "hello", F1: 30}
    var v13 uint32 = 790
    fmt.Println(v2, v5)
    v5 = int32(v1)
    switch v7 {
    case "a":
        v14 := GenS2{F0: 42}
        _ = v14
    case "":
        fmt.Println(v2)
    default:
        v15 := GenS1{F0: "\n\t", F1: 1}
        _ = v15
    }
    v7 = ""
    _ = len(v11)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
