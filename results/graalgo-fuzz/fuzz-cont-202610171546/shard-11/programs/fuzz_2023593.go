package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint = 255
    var v2 int32 = 518
    var v3 int64 = -610
    v4 := []string{"foo bar"}
    switch int(v3) {
    case 127:
        switch 128 {
        case -1:
            var v5 uint32 = 789
            _ = v5
        case 4294967295:
            v3 += 482
        default:
            v2 -= v2
        }
    }
    var v6 uint16 = 255
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 1, F1: (^(v6))}
        v9 := GenS1{F0: 0, F1: (v6 % ((v6) | 1))}
        v2 -= (v2 - v2)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int = -798
    switch "\n\t" {
    case "hello":
        v4 = append(v4, "hello")
    case "a":
        if (!(!((v1 * (v1 + v1)) <= (v1 & v1)))) {
            var v11 uint64 = 390
            var v12 uint16 = (v6 + (^(v6)))
            _ = v11
            _ = v12
        } else {
            fmt.Println(v1, v2, v2)
        }
    }
    fmt.Println(v3, v1)
    v3 = (v3 << (((^(v1))) & 63))
    if false {
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                v1 -= (v1 % ((v1) | 1))
                fmt.Println(v3)
                _ = fi14
            }
            _ = fi13
        }
    }
    var v15 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
