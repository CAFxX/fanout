package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 0
    v2 := make(map[int]uint32)
    v3 := []uint32{2, 881}
    v4 := make(map[int]bool)
    var v5 bool = (65535 < 4294967295)
    var v6 uint32 = 0
    v6 = 1
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 float64 = 1e10
        _ = fi7
        _ = v8
    }
    v2[32767] = (v6 ^ v6)
    switch -898 {
    case 128:
        v6 = v6
    case 1:
        var v9 int = 590
        _ = v9
    case 32768:
        var v10 float64 = 1.5
        _ = v10
    default:
        var v11 uint = 4294967295
        _ = v11
    }
    switch -2147483648 {
    case 1:
        if (-207 != 1) {
            delete(v2, 0)
            v1 = v1
        } else {
            v12 := v2[32767]
            var v13 float64 = (((0.0 * 3.14159) + (((0.0 * float64((1e10 - 0.1))) - 3.14159) / -2.25)) * (1e10 - 1e10))
            switch 821 {
            case -256:
                v6 -= (v12 | (v6 * v6))
            }
            _ = v12
            _ = v13
        }
    }
    _ = len(v3)
    var v14 uint32 = v6
    v5 = false
    var v15 float32 = -2.25
    delete(v2, 488)
    switch "foo bar" {
    case "hello":
        if true {
            v14 = 947
        }
    case "":
        v16 := v4[int(v6)]
        _ = v16
    case "a":
        v6 -= v6
    }
    var v17 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
