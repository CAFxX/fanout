package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int8; F1 int32; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 uint16 = 1
    v3 -= (^(v3))
    if (1.5 >= 0.1) {
        v3 += 1
    }
    v3 = 65535
    v3 = v3
    v4 := []int{798}
    if true {
        if (-1 > 42) {
            var v5 bool = false
            _ = v5
        } else {
            var v6 int32 = -1
            _ = v6
        }
        v7 := GenS1{F0: true}
        _ = v7
    }
    if false {
        v3 -= (v3 << ((65535) & 63))
    } else {
        if true {
            v8 := GenS1{F0: (!true)}
            _ = v8
        }
        var v9 float32 = float32(1e10)
        _ = v9
    }
    var v10 int32 = 42
    fmt.Println(v3, v10)
    if (627 >= 65535) {
        v11 := GenS2{F0: "\n\t", F1: 705}
        _ = v11
    } else {
        v3 = 42
        var v12 bool = true
        if v12 {
            if len(v4) > 0 {
                v4[0] = 1
            }
        }
        _ = v12
    }
    var v13 float64 = 1e10
    if false {
        fmt.Println(v10, v3, v13)
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS3{F0: 1, F1: v10, F2: 83}
        fmt.Println(fi14)
        _ = fi14
        _ = v15
    }
    v3 += (v3 + 1)
    for fi16 := 0; fi16 < 6; fi16++ {
        fmt.Println(v10)
        if true {
            var v17 uint32 = 65535
            v17 += v17
            _ = v17
        } else {
            v18 := GenS2{F0: "", F1: 4294967295}
            fmt.Println(v3)
            _ = v18
        }
        _ = fi16
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        switch 9223372036854775807 {
        case -2147483648:
            var v20 int32 = (v10 - 42)
            _ = v20
        }
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
