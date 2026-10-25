package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint64 }
func genFn0() uint64 {
    var v5 int32 = 1
    v5 = (v5 ^ v5)
    switch (("hello" + "hello") + "\n\t") {
    case "hello":
        var v6 uint = 255
        _ = v6
    default:
        v5 -= (v5 >> ((2) & 63))
    }
    v5 = (v5 / ((-1) | 1))
    _ = v5
    return 1
}

func main() {
    var v1 uint64 = 2
    var v2 int16 = 32767
    var v3 int32 = 42
    v4 := []string{}
    v3 = (v3 | (-(v3)))
    v7 := GenS1{F0: 2, F1: (-(v3)), F2: 644}
    v1 = (v1 << ((42) & 63))
    v8 := GenS1{F0: 201, F1: (v3 * -2147483648), F2: (v1 % ((v1) | 1))}
    if true {
        var v9 bool = true
        _ = v9
    } else {
        v2 += v2
    }
    fmt.Println(v3, v3, v1)
    v2 = (v2 ^ -1)
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
