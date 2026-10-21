package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint }

func main() {
    v1 := []string{}
    var v2 int16 = 1
    v3 := []uint32{42, 65535}
    fmt.Println(v2)
    v2 -= v2
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int = (fi4 >> ((0) & 63))
        v2 -= (v2 % ((v2) | 1))
        _ = fi4
        _ = v5
    }
    if (true && false) {
        v2 += -1
        fmt.Println(v2)
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        v2 += v2
        _ = fi6
    }
    fmt.Println(v2)
    var v7 int = 127
    switch "\n\t" {
    case "hello":
        if (float32(0.0) >= float32(0.0)) {
            v7 = -2147483648
            switch "foo bar" {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                for fi8 := 0; fi8 < 13; fi8++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi8
                }
            }
        }
    case "\n\t":
        var v9 float32 = (0.1 / (1e10 * float32((0.1 * (((-2.25 - (float64((0.0 * -2.25)) + 0.1)) - 0.1) - -2.25)))))
        _ = v9
    case "a":
        v10 := GenS1{F0: 42, F1: 42, F2: 0}
        _ = v10
    default:
        v2 -= v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
