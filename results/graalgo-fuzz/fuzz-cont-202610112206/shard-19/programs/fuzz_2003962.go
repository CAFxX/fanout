package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int32; F1 uint }

func main() {
    v1 := []uint32{65535, 2147483648, 2147483647, 86}
    v2 := make(map[uint32]int)
    var v3 uint32 = 654
    var v4 float32 = -2.25
    switch ("hello" + "\n\t") {
    case "\n\t":
        var v5 float32 = (0.0 * v4)
        _ = v5
    }
    if false {
        v3 = (v3 >> ((42) & 63))
        switch (("a" + "foo bar") + "a") {
        case "":
            if len(v1) > 0 {
                v1[0] = (^(v3))
            }
        case "\n\t":
            var v6 int8 = -1
            _ = v6
        case "hello":
            var v7 int16 = 543
            _ = v7
        default:
            fmt.Println(v3, v3)
        }
    } else {
        if ((-24 > 9223372036854775807) || (59 >= -1)) {
            v4 -= 1.5
            v4 += ((v4 + (v4 - v4)) / float32(1.5))
        } else {
            v3 -= v3
            if false {
                var v8 int8 = 1
                _ = v8
            } else {
                if false {
                    var v9 uint32 = v3
                    delete(v2, 42)
                    _ = v9
                }
                var v10 uint8 = 127
                _ = v10
            }
        }
        fmt.Println(v3, v4)
    }
    v11 := GenS1{F0: 9223372036854775807, F1: -1}
    var v12 uint8 = 93
    var v13 uint = 42
    fmt.Println(v4, v13, v4)
    var v14 uint16 = 32767
    v15 := GenS2{F0: 2147483647, F1: (v13 + v13)}
    fmt.Println(v12, v13, v14)
    v13 = (v13 + (^(v13)))
    var v16 uint = v13
    var v17 float64 = float64(v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v18 uint8 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
