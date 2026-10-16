package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int64, p1 string) uint16 {
    v6 := []int{1, 9223372036854775807}
    var v7 int32 = 618
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []int{42, 42}
    v2 := make(map[string]int)
    var v3 uint16 = 255
    var v4 uint8 = 255
    v5 := make(map[string]bool)
    v8 := GenS1{F0: 1}
    switch 9223372036854775807 {
    case 128:
        _ = len(v1)
    case 127:
        var v9 int32 = 416
        _ = v9
    case -1:
        fmt.Println(v4, v3)
    default:
        var v10 uint32 = 4294967295
        _ = v10
    }
    v11 := GenS1{F0: -1}
    v3 -= (v3 % (((v3 % ((v3) | 1))) | 1))
    var v12 int = 2147483647
    for fi13 := 0; fi13 < 3; fi13++ {
        fmt.Println(fi13)
        _ = fi13
    }
    var v14 uint64 = 4294967295
    if (678 <= 4294967295) {
        v15 := GenS1{F0: 42}
        fmt.Println(v4, v12)
        v14 = (^(v14))
        _ = v15
    }
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
