package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 string = "hello"
    v2 := []uint32{1, 2, 900, 2}
    var v3 float32 = 3.14159
    var v4 int32 = -2147483648
    var v5 int16 = 887
    switch -9223372036854775808 {
    case 635:
        var v6 int64 = -1
        _ = v6
    case 4294967295:
        var v7 int8 = 1
        _ = v7
    }
    v4 += 0
    for fi8 := 0; fi8 < 4; fi8++ {
        v5 = 42
        for fi9 := 0; fi9 < 6; fi9++ {
            v5 = v5
            v10 := GenS1{F0: true, F1: 118, F2: 1}
            _ = fi9
            _ = v10
        }
        if ((273 > -307) && (false || false)) {
            v11 := GenS2{F0: 189}
            _ = v11
        }
        _ = fi8
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v1 = v1
    v12 := GenS1{F0: (true || false), F1: 255, F2: 2}
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS3{F0: (v4 & 0)}
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
