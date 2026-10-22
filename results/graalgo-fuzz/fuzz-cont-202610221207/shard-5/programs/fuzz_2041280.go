package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 int16 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []uint32{2147483648, 65535, 200, 0}
    var v2 float32 = 1e10
    var v3 int16 = -788
    var v4 int64 = 42
    v5 := []int{-596}
    var v6 bool = (983 != 16)
    switch 868 {
    case 1:
        var v7 uint64 = 1
        _ = v7
    case 728:
        var v8 int = 127
        _ = v8
    }
    v9 := GenS1{F0: int8(v4)}
    v4 -= -1
    var v10 int16 = (v3 >> ((592) & 63))
    v11 := GenS2{F0: 0, F1: 2, F2: v3}
    v2 = v2
    for fi12 := 0; fi12 < 16; fi12++ {
        v4 = (v4 + -1)
        _ = fi12
    }
    var v13 uint64 = 18446744073709551615
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v14 uint32 = 65535
    v14 += v14
    v3 = -424
    fmt.Println(v13, v2)
    v2 = v2
    switch ("\n\t" + "") {
    case "":
        v15 := GenS3{F0: int32(v3)}
        _ = v15
    default:
        var v16 uint16 = 32767
        _ = v16
    }
    fmt.Println(v10, v14, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
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
    _ = v9
}
