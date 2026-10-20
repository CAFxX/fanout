package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int8; F2 uint16 }
func genFn0(p0 int16, p1 uint8, p2 int) uint8 {
    switch "" {
    case "":
        var v7 uint64 = 0
        _ = v7
    case "foo bar":
        p1 -= p1
    default:
        if false {
            v8 := GenS1{F0: 4294967295}
            var v9 uint32 = 1
            v10 := GenS1{F0: 18446744073709551615}
            _ = v10
            _ = v8
            _ = v9
        }
    }
    v11 := GenS2{F0: (false || true), F1: -1, F2: 983}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (p1 % ((1) | 1))
}

func main() {
    var v1 int16 = 707
    v2 := make(map[int]string)
    var v3 uint8 = 27
    var v4 uint64 = 40
    v5 := make(map[uint32]int)
    var v6 int = 128
    v6 = v6
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v3, v3, v1)
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := []int{42, fi12, fi13}
            _ = fi13
            _ = v14
        }
        fmt.Println(v1, v3, fi12)
        _ = fi12
    }
    var v15 int32 = 42
    v16 := GenS2{F0: false, F1: 127, F2: 2}
    v17 := GenS2{F0: true, F1: -128, F2: 42}
    switch "" {
    case "foo bar":
        var v18 uint32 = 2
        _ = v18
    default:
        delete(v5, 2147483647)
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            v4 -= 235
            var v21 uint = 2
            switch "\n\t" {
            case "a":
                if false {
                    fmt.Println(fi19, fi20)
                }
            case "\n\t":
                var v22 int64 = -796
                _ = v22
            case "hello":
                var v23 uint = v21
                v24 := GenS1{F0: v4}
                _ = v23
                _ = v24
            }
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    var v25 uint = 42
    var v26 float32 = (1e10 - 0.1)
    v5[42] = (v6 | (v6 & (-(v6))))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
