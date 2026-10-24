package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint }
func genFn0() int8 {
    var v7 uint8 = 2
    v8 := GenS1{F0: "", F1: -32768, F2: -102}
    v9 := []int{4294967295, 2147483647, -17, 9223372036854775807}
    if len(v9) > 0 {
        v9[0] = 9223372036854775807
    }
    _ = v7
    _ = v8
    _ = v9
    return -67
}

func main() {
    v1 := []int{4294967295}
    var v2 uint8 = 51
    var v3 int32 = -1
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]bool)
    v10 := GenS1{F0: "", F1: 1000, F2: 0}
    v11 := GenS1{F0: "a", F1: 433, F2: 76}
    var v12 int64 = 1
    v13 := GenS2{F0: 2147483647, F1: (v2 >> ((18446744073709551615) & 63)), F2: 255}
    var v14 int32 = -350
    var v15 uint8 = (v2 - (v2 % (((v2 % ((128) | 1))) | 1)))
    v6["a"] = (v3 <= v14)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 int = 9223372036854775807
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
