package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint32 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(fi4)
        fmt.Println(fi4)
        _ = fi4
    }
    switch 127 {
    case 765:
        var v5 uint16 = 65535
        _ = v5
    case 4294967295:
        var v6 uint = 65535
        _ = v6
    case -2147483648:
        var v7 int = 9223372036854775807
        _ = v7
    }
    return -1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    v1[""] = (0 <= 1)
    var v8 int = -2147483648
    v9 := GenS1{F0: 0, F1: 0}
    for fi10 := 0; fi10 < 15; fi10++ {
        v8 = 2147483647
        v8 += (^(fi10))
        _ = fi10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{9223372036854775807, (^(fi11)), v8}
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 856}
    if (2147483647 >= 2147483647) {
        if false {
            v14 := []int{}
            _ = v14
        } else {
            fmt.Println(v8)
        }
    } else {
        for fi15 := 0; fi15 < 5; fi15++ {
            v8 += (v8 * (fi15 ^ (fi15 >> ((109) & 63))))
            _ = fi15
        }
    }
    v16 := []int{}
    v8 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
