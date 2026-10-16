package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }

func main() {
    v1 := []uint32{2147483648, 313, 209, 2147483648}
    v2 := []uint32{2147483647, 1, 563, 2147483647}
    v3 := make(map[int]uint32)
    var v4 uint64 = 0
    fmt.Println(v4)
    v4 += 18446744073709551615
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v3[-1] = 2147483648
    fmt.Println(v4)
    switch 128 {
    case 42:
        switch (("" + ("hello" + "a")) + ("hello" + "hello")) {
        case "":
            var v5 int64 = -8
            _ = v5
        case "foo bar":
            var v6 float32 = -2.25
            _ = v6
        default:
            v4 = (v4 + v4)
        }
    default:
        var v7 uint8 = 1
        _ = v7
    }
    delete(v3, 2147483647)
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: v4, F1: 835}
        var v10 int32 = 727
        var v11 int16 = 1
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v4 = (v4 * (v4 * v4))
    v3[-770] = 458
    _ = len(v3)
    if true {
        v12 := GenS1{F0: v4, F1: 42}
        _ = v12
    }
    fmt.Println(v4)
    var v13 uint16 = 929
    v13 = 32767
    switch 2147483647 {
    case -967:
        for fi14 := 0; fi14 < 19; fi14++ {
            v13 = (v13 + v13)
            switch (^(fi14)) {
            case 4294967295:
                switch fi14 {
                case 42:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 1:
                    var v15 uint32 = 1
                    _ = v15
                case 29:
                    v2 = append(v2, 2147483647)
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            delete(v3, fi14)
            _ = fi14
        }
    default:
        if true {
            var v16 float32 = 3.14159
            _ = v16
        }
    }
    delete(v3, 42)
    var v17 float64 = float64((1.5 * (-2.25 * (1.5 / (-2.25 - (float32(0.0) / -2.25))))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
