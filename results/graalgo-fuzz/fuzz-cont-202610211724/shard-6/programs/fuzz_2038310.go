package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint }

func main() {
    v1 := []uint32{}
    var v2 uint8 = 255
    var v3 uint32 = 2
    var v4 int64 = -9223372036854775808
    v5 := make(map[uint32]string)
    var v6 uint32 = 2147483648
    if (-1 == 100) {
        v4 = int64(v6)
    }
    v2 = (v2 | v2)
    _ = len(v1)
    switch "a" {
    case "":
        v2 = 255
    case "a":
        for fi7 := 0; fi7 < 6; fi7++ {
            v5[v3] = "a"
            _ = len(v1)
            _ = fi7
        }
    case "\n\t":
        v8 := v5[65535]
        _ = v8
    }
    v1 = append(v1, (v6 / (((v6 - (v3 - v3))) | 1)))
    var v9 uint32 = v6
    var v10 int = 42
    delete(v5, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int = 127
        if (2 == 255) {
            switch (v12 * 126) {
            case 127:
                v6 = v6
            case -439:
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 uint = 1
                    v15 := GenS1{F0: v4, F1: 65535}
                    if true {
                        fmt.Println(v4)
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case 934:
                var v16 int64 = (v4 + v4)
                _ = v16
            }
            fmt.Println(v12, v10)
        }
        v12 = (fi11 ^ v10)
        _ = fi11
        _ = v12
    }
    var v17 float32 = float32(1.5)
    v2 -= v2
    var v18 bool = (-1 >= 1)
    v19 := GenS2{F0: 128, F1: 873}
    if (("\n\t" + "foo bar") > ("\n\t" + (("hello" + "\n\t") + ""))) {
        _ = len(v5)
    }
    v3 -= (v3 | (v3 * v9))
    v20 := GenS1{F0: 42, F1: 0}
    if (127 != 1) {
        switch (v10 / ((v10) | 1)) {
        case -9223372036854775808:
            v9 = 198
        default:
            switch "foo bar" {
            case "hello":
                var v21 uint8 = (v2 << ((42) & 63))
                _ = v21
            case "\n\t":
                v2 += (v2 - v2)
            case "foo bar":
                fmt.Println(v6, v10, v2)
            default:
                v9 = v9
            }
        }
    } else {
        var v22 bool = (0 != -1)
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
