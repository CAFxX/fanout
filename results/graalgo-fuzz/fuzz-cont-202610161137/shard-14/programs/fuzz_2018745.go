package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 string }
type GenS2 struct { F0 int16; F1 int }
func genFn0(p0 int32, p1 int) int8 {
    v6 := GenS1{F0: p1, F1: 0, F2: ("" + "hello")}
    var v7 int16 = -32768
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int32) uint8 {
    v8 := GenS2{F0: 182, F1: 128}
    var v9 uint32 = 67
    switch "foo bar" {
    case "a":
        fmt.Println(p0)
    }
    p0 -= -1
    _ = p0
    _ = v8
    _ = v9
    return 152
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []int{42, -835, 0}
    var v3 bool = false
    v4 := make(map[uint32]bool)
    var v5 bool = v3
    v2 = append(v2, 42)
    if (1 != -9223372036854775808) {
        switch 2147483647 {
        case 2147483647:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            if v5 {
                v3 = false
            }
            _ = fi10
        }
    } else {
        v11 := GenS2{F0: 264, F1: 58}
        v12 := GenS1{F0: 1, F1: 104, F2: ("hello" + "hello")}
        _ = v11
        _ = v12
    }
    var v13 int16 = 42
    var v14 float64 = (1e10 / float64(0.1))
    v1 = (^(v1))
    if true {
        v15 := GenS2{F0: v13, F1: -876}
        var v16 int32 = -1
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: 0, F1: 42, F2: ""}
        v13 += (v13 & 686)
        _ = v17
    }
    if len(v2) > 0 {
        v2[0] = 84
    }
    var v18 float32 = 0.0
    var v19 int32 = -961
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
