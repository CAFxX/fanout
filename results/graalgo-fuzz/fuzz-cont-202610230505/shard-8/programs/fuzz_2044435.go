package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 int16, p1 int32, p2 uint64) uint32 {
    p2 += (p2 << ((47) & 63))
    switch "hello" {
    case "a":
        v5 := GenS1{F0: 897, F1: 1}
        _ = v5
    case "\n\t":
        p1 = (p1 & p1)
    case "hello":
        v6 := []int{}
        _ = v6
    }
    v7 := []int{716}
    v8 := GenS2{F0: ("hello" + "hello"), F1: 2147483647}
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 940, F1: 9223372036854775807}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 761
}
func genFn1() uint8 {
    var v11 int = 2147483647
    for fi12 := 0; fi12 < 4; fi12++ {
        if true {
            v13 := []int{1}
            _ = v13
        }
        _ = fi12
    }
    var v14 uint64 = 291
    _ = v11
    _ = v14
    return 156
}

func main() {
    v1 := []string{"a", "\n\t"}
    var v2 uint = 521
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    var v15 uint8 = 0
    var v16 int16 = -1
    v16 = (^(v16))
    fmt.Println(v2)
    var v17 string = ("hello" + "a")
    v3[42] = 42
    _ = len(v3)
    if len(v1) > 0 {
        v1[0] = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
