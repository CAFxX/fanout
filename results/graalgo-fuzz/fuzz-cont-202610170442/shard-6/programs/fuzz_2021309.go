package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int32 }
func genFn0(p0 int32, p1 int16, p2 int16) int8 {
    fmt.Println(p1)
    var v6 uint32 = 1
    v7 := GenS1{F0: 0}
    var v8 int8 = -1
    var v9 uint64 = 892
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return int8(v9)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 65535
    var v3 int8 = 127
    var v4 uint16 = 336
    v5 := make(map[string]int)
    v2 -= v2
    v3 -= v3
    for fi10 := 0; fi10 < 5; fi10++ {
        if true {
            v2 += v2
            var v11 uint8 = 228
            fmt.Println(v11)
            _ = v11
        } else {
            v12 := GenS2{F0: (v4 >> ((2) & 63)), F1: -132}
            if true {
                v2 = (v2 ^ v2)
            }
            _ = v12
        }
        switch (fi10 % ((fi10) | 1)) {
        case -1:
            v4 = (v4 / (((v4 % ((611) | 1))) | 1))
        case -2147483648:
            delete(v5, "hello")
        default:
            v4 = v4
        }
        _ = fi10
    }
    v13 := v5["foo bar"]
    fmt.Println(v2)
    fmt.Println(v2)
    var v14 uint = 726
    var v15 int16 = 1
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 float64 = 1.5
        fmt.Println(fi16)
        v15 -= v15
        _ = fi16
        _ = v17
    }
    v18 := []int{(v13 >> ((v14) & 63)), 726, v13, (v13 >> (((v14 - (v14 / ((v14) | 1)))) & 63))}
    if false {
        v2 -= v2
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 float32 = 0.0
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = fi19
            _ = v20
        }
    }
    v14 = v14
    v21 := v5["hello"]
    var v22 int64 = 5
    if len(v18) > 0 {
        v18[0] = -9223372036854775808
    }
    v23 := GenS2{F0: v4, F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
