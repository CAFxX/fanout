package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    var v1 uint16 = 32767
    var v2 uint64 = 311
    v3 := []string{(("\n\t" + "\n\t") + (((("hello" + "a") + "hello") + "foo bar") + "")), "\n\t"}
    var v4 uint16 = 42
    v5 := []int{}
    v6 := GenS1{F0: 42, F1: ((((1.5 * 1e10) / ((3.14159 - (-2.25 - float64(0.0))) * (((((((0.0 + (0.1 / 0.1)) / float64(3.14159)) + 1e10) + 1e10) + float64(0.0)) + float64((0.1 * -2.25))) / (0.1 / 1.0)))) > float64((1.5 + (0.0 + 0.0)))) || (!(v2 <= 4294967295)))}
    v1 = v1
    var v7 uint8 = 255
    var v8 uint64 = 521
    switch 668 {
    case -9223372036854775808:
        var v9 int = 32767
        _ = v9
    case 0:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 127:
        var v10 float64 = 1e10
        _ = v10
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v11 uint32 = 42
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint64 = (v2 % (((v8 + 2)) | 1))
        v8 += (v13 ^ v8)
        var v14 uint32 = 779
        _ = fi12
        _ = v13
        _ = v14
    }
    v8 = (v8 << ((581) & 63))
    var v15 int8 = -63
    v4 = (v4 - 0)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
