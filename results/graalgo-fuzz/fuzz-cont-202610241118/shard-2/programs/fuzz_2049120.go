package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 int; F1 int32; F2 string }

func main() {
    var v1 uint32 = 731
    v2 := make(map[int]string)
    v3 := []uint32{v1, (v1 & v1)}
    v4 := []string{("\n\t" + ("a" + "hello"))}
    fmt.Println(v1)
    fmt.Println(v1)
    v4 = append(v4, "")
    switch "hello" {
    case "hello":
        v1 -= v1
    case "\n\t":
        v1 = 0
    case "foo bar":
        v5 := GenS1{F0: 32768, F1: 255, F2: 326}
        _ = v5
    }
    v6 := GenS2{F0: -499, F1: -2147483648, F2: (("" + "hello") + "")}
    v1 += (v1 * (v1 | v1))
    switch 279 {
    case -2147483648:
        var v7 uint32 = v1
        _ = v7
    case 0:
        var v8 int16 = 32767
        _ = v8
    case 1:
        var v9 float64 = -2.25
        _ = v9
    default:
        if len(v3) > 0 {
            v3[0] = (v1 << ((255) & 63))
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v1
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int = -441
        var v12 int64 = 9223372036854775807
        _ = fi10
        _ = v11
        _ = v12
    }
    v1 = v1
    var v13 uint = 225
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
