package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
func genFn0(p0 uint32, p1 uint64) bool {
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(p1, fi7, p0)
        var v8 int32 = 1
        _ = fi7
        _ = v8
    }
    p1 += (p1 % (((^(p1))) | 1))
    v9 := []int{42, int(p0), -2147483648}
    p0 = (p0 << ((42) & 63))
    _ = p0
    _ = p1
    _ = v9
    return (!false)
}
func genFn1(p0 int64, p1 uint16) uint {
    var v10 float32 = 1e10
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 543
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    v5 := []string{"\n\t", ("hello" + "\n\t")}
    v6 := make(map[string]bool)
    v12 := GenS1{F0: "hello", F1: 104}
    v2 = v2
    v2 = (v2 & (^(v2)))
    var v13 uint64 = 1
    v14 := GenS1{F0: "", F1: 255}
    if false {
        var v15 uint32 = 0
        var v16 uint32 = (v15 | v15)
        var v17 float32 = (1.5 * -2.25)
        _ = v15
        _ = v16
        _ = v17
    } else {
        v13 = (v13 << ((255) & 63))
        v13 = 0
    }
    var v18 string = ("" + "a")
    v18 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
