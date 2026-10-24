package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int = 32767
    var v2 float32 = (0.1 + 1e10)
    v3 := make(map[uint32]int)
    if true {
        switch ("foo bar" + "") {
        case "\n\t":
            v4 := GenS1{F0: (!(3.14159 > v2))}
            _ = v4
        case "a":
            v5 := GenS1{F0: false}
            _ = v5
        }
        fmt.Println(v1, v1)
    } else {
        v6 := v3[0]
        _ = v6
    }
    if ((!(2147483647 == 42)) && false) {
        fmt.Println(v2, v1)
        v1 -= (v1 % ((9223372036854775807) | 1))
    } else {
        v7 := GenS1{F0: true}
        v8 := v3[16]
        _ = v7
        _ = v8
    }
    var v9 uint = 65535
    v10 := GenS2{F0: -32768}
    v3[0] = (v1 << (((v9 % ((2) | 1))) & 63))
    if true {
        var v11 uint16 = 977
        _ = v11
    } else {
        fmt.Println(v1)
        _ = len(v3)
        fmt.Println(v2)
    }
    fmt.Println(v1)
    var v12 float32 = v2
    if (!(-986 <= -134)) {
        v13 := GenS1{F0: false}
        v9 += (v9 / ((v9) | 1))
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
