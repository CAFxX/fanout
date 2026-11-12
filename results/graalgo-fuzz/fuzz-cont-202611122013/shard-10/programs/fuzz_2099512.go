package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 657
    var v3 float32 = 0.1
    var v4 uint64 = 562
    v5 := make(map[int]uint32)
    _ = len(v1)
    for fi6 := 0; fi6 < 12; fi6++ {
        v2 = (v2 >> ((18446744073709551615) & 63))
        if (!true) {
            var v7 float64 = 1.5
            fmt.Println(v2)
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(v2, v3, v2)
    if ((!(((594 < 35) && (1 <= 255)) && (false && (116 > 128)))) || true) {
        var v8 uint64 = (v4 << ((42) & 63))
        var v9 float32 = (v3 - (v3 - v3))
        _ = v8
        _ = v9
    }
    switch ("hello" + "\n\t") {
    case "a":
        v10 := GenS1{F0: 65535, F1: 0, F2: 1}
        _ = v10
    case "hello":
        if (-457 < -650) {
            switch 42 {
            case -529:
                if true {
                    var v11 int8 = -1
                    _ = v11
                }
            case 32767:
                v3 = 1e10
            default:
                v12 := v5[9223372036854775807]
                _ = v12
            }
        }
    case "":
        var v13 float64 = (0.0 / 1.0)
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 bool = true
        v16 := []int{1, (fi14 / ((1) | 1)), (fi14 * -1)}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int16 = 1
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
