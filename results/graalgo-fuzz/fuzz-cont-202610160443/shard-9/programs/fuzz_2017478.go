package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint16; F1 bool }

func main() {
    var v1 int8 = 100
    var v2 uint64 = 376
    var v3 string = "hello"
    v4 := make(map[int]string)
    var v5 uint = 730
    var v6 int16 = -32768
    var v7 int64 = -1
    delete(v4, -29)
    if ((127 >= 128) || true) {
        v4[2147483647] = v3
        v8 := GenS1{F0: 1, F1: 1}
        _ = v8
    }
    _ = len(v4)
    v5 = 18446744073709551615
    v9 := []int{}
    switch v3 {
    case "a":
        v10 := GenS1{F0: 2147483648, F1: -604}
        _ = v10
    case "hello":
        if (((0.1 * float64(1.5)) + 3.14159) >= 0.1) {
            var v11 string = ((v3 + v3) + (v3 + v3))
            _ = v11
        }
    case "":
        v5 = (v5 >> (((v5 - 18446744073709551615)) & 63))
    default:
        v3 = v3
    }
    v7 = (-(v7))
    var v12 bool = (v7 >= (v7 % ((v7) | 1)))
    v7 = (v7 >> ((65535) & 63))
    if (42 > 211) {
        switch ((v3 + (v3 + (("" + ("" + (v3 + ("hello" + v3)))) + v3))) + "hello") {
        case "foo bar":
            v7 = 0
        }
    }
    v13 := v4[32767]
    v14 := GenS2{F0: 667, F1: v12}
    var v15 bool = ((false && ((0 > 774) || (!true))) && v12)
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
