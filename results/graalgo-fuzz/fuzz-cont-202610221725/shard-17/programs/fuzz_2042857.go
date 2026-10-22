package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 int, p1 uint16) int8 {
    var v7 uint64 = 18446744073709551615
    v8 := []int{(^(p0)), (^(p0)), p0, (p0 + p0)}
    if true {
        switch 2147483647 {
        case 0:
            if len(v8) > 0 {
                v8[0] = int(p1)
            }
        case 9223372036854775807:
            fmt.Println(p0, p1)
        }
        v9 := GenS1{F0: 65535, F1: ("" + "foo bar")}
        var v10 uint = 233
        _ = v10
        _ = v9
    } else {
        if true {
            var v11 uint64 = (v7 - (^(v7)))
            var v12 uint32 = 840
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 100
}
func genFn1() uint64 {
    v13 := []int{}
    var v14 uint64 = 1
    var v15 bool = ((0 <= -32768) && true)
    _ = v13
    _ = v14
    _ = v15
    return 42
}

func main() {
    v1 := []int{-765, 9223372036854775807, 666, 128}
    var v2 uint = 42
    var v3 float32 = 0.0
    var v4 int64 = -9223372036854775808
    var v5 uint8 = 105
    var v6 uint16 = 220
    v16 := GenS2{F0: 42}
    var v17 uint8 = v5
    v4 = (v4 << ((v2) & 63))
    switch "a" {
    case "hello":
        v18 := GenS2{F0: 1}
        _ = v18
    case "a":
        var v19 uint8 = v17
        _ = v19
    }
    var v20 int64 = -9223372036854775808
    var v21 int64 = (v20 ^ (v20 + v20))
    for fi22 := 0; fi22 < 1; fi22++ {
        v17 += v5
        var v23 uint = 18446744073709551615
        _ = fi22
        _ = v23
    }
    v4 = v21
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
