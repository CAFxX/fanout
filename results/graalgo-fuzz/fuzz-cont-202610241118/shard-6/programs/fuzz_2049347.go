package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64, p1 bool) uint64 {
    v4 := []int{9223372036854775807}
    p0 += 1
    fmt.Println(p1, p0)
    p0 -= p0
    _ = p0
    _ = p1
    _ = v4
    return (p0 / (((p0 | (^(p0)))) | 1))
}

func main() {
    var v1 bool = (true && true)
    v2 := make(map[uint32]int)
    var v3 bool = (-1 <= 1)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int8 = 1
        v7 := GenS1{F0: fi5, F1: 57}
        var v8 string = "a"
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int32 = -119
    switch 32767 {
    case 881:
        var v10 uint = 4294967295
        _ = v10
    case -206:
        fmt.Println(v3, v3, v1)
    case 2147483647:
        _ = len(v2)
    default:
        switch "" {
        case "foo bar":
            var v11 uint8 = 0
            _ = v11
        case "\n\t":
            if true {
                var v12 float32 = -2.25
                _ = v12
            } else {
                var v13 int8 = -128
                _ = v13
            }
        case "hello":
            var v14 uint32 = 2147483648
            _ = v14
        }
    }
    fmt.Println(v3, v3, v1)
    var v15 int64 = -9223372036854775808
    fmt.Println(v15)
    v1 = v1
    fmt.Println(v1)
    if v1 {
        v16 := GenS2{F0: 0}
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
