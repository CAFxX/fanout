package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    var v3 uint16 = 255
    v4 := make(map[int]uint32)
    v2[""] = 658
    v5 := GenS1{F0: -1, F1: 565, F2: 32767}
    v6 := GenS2{F0: 88, F1: (v3 ^ v3), F2: 2}
    fmt.Println(v3)
    var v7 int = 521
    var v8 uint64 = 184
    v9 := GenS2{F0: 2, F1: v3, F2: 0}
    v10 := GenS1{F0: 511, F1: -9223372036854775808, F2: (v3 ^ (v3 ^ v3))}
    switch (v7 - v7) {
    case 32768:
        v11 := []int{(v7 << ((746) & 63))}
        _ = v11
    }
    v3 -= v3
    v12 := GenS3{F0: 651, F1: "hello"}
    var v13 bool = ((("" + "\n\t") > "a") && (2147483647 >= 0))
    var v14 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
