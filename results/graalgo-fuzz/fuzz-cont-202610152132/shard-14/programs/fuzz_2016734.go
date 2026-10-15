package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 1
    var v3 float64 = float64(0.1)
    v4 := GenS1{F0: 212, F1: ("" + ((("\n\t" + "foo bar") + (("\n\t" + "") + ("" + ""))) + "a")), F2: int32(v2)}
    var v5 int32 = 14
    var v6 uint16 = 0
    var v7 float64 = v3
    switch 771 {
    case 2147483647:
        var v8 int64 = 0
        _ = v8
    case 128:
        v9 := []int{int(v6), 9223372036854775807}
        _ = v9
    }
    v6 = (^(v6))
    v10 := GenS2{F0: 29}
    switch "" {
    case "":
        v11 := GenS3{F0: -78}
        _ = v11
    case "a":
        var v12 bool = true
        _ = v12
    case "\n\t":
        _ = len(v1)
    default:
        v5 = (^(v5))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
