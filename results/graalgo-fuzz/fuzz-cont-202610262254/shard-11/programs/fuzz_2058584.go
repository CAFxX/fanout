package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int64 }

func main() {
    var v1 uint16 = 42
    var v2 string = (("\n\t" + "foo bar") + ("hello" + "\n\t"))
    var v3 string = v2
    v4 := make(map[uint32]bool)
    v5 := make(map[string]uint32)
    var v6 uint32 = 65535
    switch "foo bar" {
    case "foo bar":
        var v7 uint16 = v1
        _ = v7
    case "":
        fmt.Println(v3, v3)
    default:
        var v8 int16 = -32768
        _ = v8
    }
    v9 := GenS1{F0: v6, F1: false, F2: 763}
    v10 := GenS1{F0: 4294967295, F1: false, F2: 9223372036854775807}
    if false {
        v11 := []int{}
        if false {
            if (!(((((((1e10 / 1.0) / (float64(((0.0 / 1.5) * 0.0)) - 1e10)) - float64(3.14159)) / 0.1) + (((3.14159 - 1.5) - ((1e10 - (float64((1e10 - 1.5)) - 0.0)) - (3.14159 - 3.14159))) - (0.0 / 1.5))) - 1e10) >= (0.0 * (0.0 - 1.5)))) {
                v12 := GenS1{F0: v6, F1: false, F2: -303}
                switch v3 {
                case "":
                    var v13 uint = 42
                    _ = v13
                }
                _ = v12
            }
        }
        _ = v11
    }
    if (((("foo bar" + "a") + (v3 + v3)) + v3) < (v3 + (v3 + v3))) {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 int64 = -1
            fmt.Println(v2)
            v4[v6] = (false || true)
            _ = fi14
            _ = v15
        }
    }
    var v16 string = (v2 + ("hello" + v3))
    if (-2.25 != 0.0) {
        var v17 string = v3
        var v18 uint8 = 128
        _ = v17
        _ = v18
    }
    var v19 int32 = 1
    var v20 uint16 = v1
    var v21 uint8 = 2
    v22 := GenS1{F0: (^(v6)), F1: true, F2: -735}
    v5[("hello" + (v2 + v2))] = uint32(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
