package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 int64, p1 bool, p2 int64) uint32 {
    v4 := GenS1{F0: ((65535 >= 255) && p1)}
    var v5 uint32 = 2147483648
    switch "" {
    case "foo bar":
        fmt.Println(p0, p1, v5)
    case "\n\t":
        p2 += p0
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (v5 >> ((255) & 63))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = float32(((0.0 - 1.5) * 3.14159))
    v3 := []string{}
    v6 := GenS1{F0: true}
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            if ((v2 * (1e10 * v2)) < v2) {
                v8 := GenS2{F0: 42}
                v2 += float32((0.0 - (1.5 * -2.25)))
                v2 = 0.1
                _ = v8
            }
            if ("hello" >= "") {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            fmt.Println(fi7, v2)
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int8 = 127
        v11 := GenS3{F0: 127, F1: 32767}
        _ = fi9
        _ = v10
        _ = v11
    }
    v2 = 1.5
    v12 := GenS3{F0: 1, F1: 0}
    var v13 int64 = -1
    fmt.Println(v13, v13)
    if true {
        v2 += ((v2 - v2) + (float32((1.5 + (float64(v2) + (1.5 / (1e10 - 1.5))))) * v2))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
}
