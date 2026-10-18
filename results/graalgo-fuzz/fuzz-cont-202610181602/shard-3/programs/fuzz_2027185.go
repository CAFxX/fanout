package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS5 struct { F0 bool; F1 string; F2 bool }

func main() {
    var v1 uint32 = 2
    v2 := []string{"", "foo bar", "\n\t", "\n\t"}
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]string)
    var v5 uint = 831
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 0, F1: 2147483647}
        if (fi6 != (^(fi6))) {
            var v8 bool = (true && true)
            _ = len(v4)
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        v1 = (v1 << ((uint(v1)) & 63))
        _ = len(v2)
        var v10 int16 = -709
        _ = fi9
        _ = v10
    }
    var v11 uint = uint(v1)
    fmt.Println(v11, v5)
    v12 := GenS2{F0: ((0.1 * (0.1 / 1e10)) > 1e10)}
    v11 -= v11
    v13 := GenS3{F0: 0}
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS4{F0: 42, F1: v1, F2: 0}
        _ = fi14
        _ = v15
    }
    var v16 int64 = 802
    v17 := GenS5{F0: (!true), F1: "a", F2: (42 < -2147483648)}
    var v18 uint8 = 128
    var v19 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
