package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    var v1 uint16 = 0
    v2 := []uint32{438, 1}
    var v3 int = 128
    if (0.0 == (float32(0.0) - 0.0)) {
        v1 -= 65535
    }
    _ = len(v2)
    fmt.Println(v3, v1)
    v1 = v1
    var v4 float32 = 0.0
    fmt.Println(v1, v4, v3)
    _ = len(v2)
    v4 += (((1.5 / ((v4 + v4) * (((1.5 - ((v4 + v4) / (1e10 / v4))) - (v4 * 0.0)) / v4))) / v4) * 0.0)
    for fi5 := 0; fi5 < 7; fi5++ {
        if false {
            switch "\n\t" {
            case "a":
                var v6 string = "\n\t"
                _ = v6
            case "hello":
                var v7 uint8 = 128
                _ = v7
            default:
                v1 -= v1
            }
        }
        v8 := GenS1{F0: 2147483648, F1: (fi5 * (v3 >> ((42) & 63)))}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi5
        _ = v8
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
