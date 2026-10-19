package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0(p0 int64) bool {
    var v5 uint16 = 96
    v5 = v5
    p0 = p0
    _ = p0
    _ = v5
    return false
}

func main() {
    var v1 uint64 = 822
    v2 := make(map[string]bool)
    var v3 uint16 = 32767
    var v4 uint32 = 65535
    for fi6 := 0; fi6 < 5; fi6++ {
        _ = len(v2)
        _ = fi6
    }
    _ = len(v2)
    v1 -= (v1 + 4294967295)
    fmt.Println(v4)
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            fmt.Println(v1, v1, v1)
            _ = fi8
        }
        var v9 int = fi7
        v10 := []int{v9}
        _ = fi7
        _ = v10
        _ = v9
    }
    if false {
        var v11 int = -9223372036854775808
        var v12 int8 = 127
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: -128, F1: ""}
    var v14 int64 = -357
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
