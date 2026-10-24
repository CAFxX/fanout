package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]bool)
    var v3 uint = 0
    var v4 float64 = 1.5
    v3 = (v3 * 473)
    switch v1 {
    case "\n\t":
        v5 := GenS1{F0: uint64(v3)}
        _ = v5
    default:
        var v6 int64 = 9223372036854775807
        _ = v6
    }
    v4 -= (v4 * (v4 / v4))
    v7 := GenS1{F0: 1}
    fmt.Println(v1, v1, v3)
    var v8 string = "foo bar"
    fmt.Println(v8)
    v8 = "hello"
    var v9 uint64 = 304
    v3 -= (v3 >> ((v3) & 63))
    v10 := GenS1{F0: v9}
    if false {
        var v11 uint16 = 255
        v11 -= 32767
        v4 = (0.0 * 0.0)
        _ = v11
    } else {
        var v12 string = v8
        if (!(-128 <= -1)) {
            v1 = "hello"
            if ((1 <= 805) && (!true)) {
                var v13 bool = true
                _ = v13
            } else {
                var v14 uint32 = 2
                v15 := GenS1{F0: v9}
                _ = v14
                _ = v15
            }
            var v16 int8 = 0
            _ = v16
        }
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
