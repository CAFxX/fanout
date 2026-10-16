package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 uint, p1 uint32, p2 int16) uint8 {
    fmt.Println(p2, p0, p1)
    v6 := GenS1{F0: -499, F1: false}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 216
}

func main() {
    var v1 int64 = -1
    var v2 int32 = 42
    v3 := make(map[int]int)
    v4 := []int{97, -1, 9223372036854775807}
    var v5 float64 = (float64(1e10) / 3.14159)
    v4 = append(v4, 1)
    v7 := GenS2{F0: 983, F1: (v2 - (v2 % ((v2) | 1)))}
    var v8 int16 = -32768
    fmt.Println(v8, v8)
    v1 = 464
    fmt.Println(v8)
    switch "a" {
    case "foo bar":
        v2 += (^(v2))
    case "hello":
        var v9 uint16 = 2
        _ = v9
    case "a":
        v2 -= (v2 * v2)
    }
    switch ((("" + ("foo bar" + ((("\n\t" + ("\n\t" + "foo bar")) + "a") + ""))) + "foo bar") + "") {
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = -65
        }
    default:
        switch "" {
        case "":
            var v10 float32 = 1.5
            _ = v10
        case "a":
            v2 -= (v2 | (^(v2)))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
