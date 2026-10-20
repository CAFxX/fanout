package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 int16) bool {
    v5 := []int{}
    v6 := GenS1{F0: ("foo bar" + "hello"), F1: ((!(0 > 0)) && (true || (!((65535 >= 1) || (-9223372036854775808 > 0)))))}
    var v7 float64 = 1e10
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(p0)
        var v9 uint32 = 0
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 uint8 = 42
    var v2 uint8 = v1
    v3 := make(map[int]bool)
    var v4 uint64 = 383
    fmt.Println(v4, v4, v4)
    v10 := GenS2{F0: v4}
    fmt.Println(v1, v4, v1)
    v1 = v2
    for fi11 := 0; fi11 < 16; fi11++ {
        switch "\n\t" {
        case "\n\t":
            v4 += (v4 + (v4 | v4))
        default:
            v4 = v4
        }
        _ = fi11
    }
    v1 = (^(v1))
    fmt.Println(v4)
    if (128 <= int(v2)) {
        switch "foo bar" {
        case "foo bar":
            if (412 >= 2) {
                var v12 uint = 429
                _ = v12
            }
        }
    } else {
        var v13 int64 = -239
        var v14 string = "foo bar"
        v15 := []int{-790}
        _ = v13
        _ = v14
        _ = v15
    }
    if (!((v4 + (v4 - (^(v4)))) <= (v4 - v4))) {
        v1 = 104
        fmt.Println(v4, v1)
    } else {
        fmt.Println(v1)
    }
    v16 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
