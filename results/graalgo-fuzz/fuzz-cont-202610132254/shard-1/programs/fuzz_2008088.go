package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint8 = 148
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 string = "foo bar"
    var v5 int64 = 0
    v6 := []uint32{575, 240, 426, 999}
    switch ((v4 + v4) + v4) {
    case "foo bar":
        var v7 int64 = v5
        _ = v7
    case "":
        var v8 string = v4
        _ = v8
    }
    if false {
        v9 := GenS1{F0: (v4 < "a"), F1: 65535}
        _ = v9
    } else {
        v10 := v2[95]
        _ = v10
    }
    v1 += 42
    if ((1 > (v1 | (v1 & (v1 ^ 2)))) || false) {
        v5 = v5
    }
    v1 -= 164
    var v11 uint32 = 0
    fmt.Println(v1)
    if len(v6) > 0 {
        v6[0] = v11
    }
    if (4294967295 <= -2147483648) {
        v2[(^(v11))] = (^(v11))
        var v12 uint8 = v1
        fmt.Println(v12, v1, v11)
        _ = v12
    } else {
        v13 := GenS1{F0: (3.14159 >= 1.5), F1: 854}
        v11 = (v11 + 42)
        v5 = v5
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
