package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 bool }
type GenS4 struct { F0 int32 }

func main() {
    var v1 int = 989
    var v2 int = 128
    v3 := []int{v2}
    v4 := []int{(^(v1))}
    v5 := make(map[string]int)
    v6 := make(map[uint32]bool)
    v1 = v1
    v7 := GenS1{F0: 42}
    var v8 bool = false
    v5["hello"] = 972
    v9 := v5["foo bar"]
    v10 := GenS2{F0: 1, F1: 2, F2: 474}
    v1 = (v2 | 42)
    v8 = v8
    v11 := GenS3{F0: 18446744073709551615, F1: 18446744073709551615, F2: (!v8)}
    switch (v2 >> ((42) & 63)) {
    case 436:
        v5["foo bar"] = (v2 << ((4294967295) & 63))
    default:
        fmt.Println(v1, v9, v9)
    }
    v12 := v5[("a" + "")]
    v13 := GenS4{F0: 0}
    v9 += (v2 & (v1 + (v1 & v1)))
    v14 := GenS1{F0: 127}
    fmt.Println(v8, v8, v2)
    v1 -= 32767
    var v15 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
