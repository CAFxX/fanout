package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint16 }

func main() {
    var v1 int32 = 1
    var v2 int64 = 0
    var v3 int16 = 32767
    v4 := make(map[string]bool)
    v1 -= (v1 >> ((4294967295) & 63))
    var v5 float64 = 0.0
    v2 = v2
    switch "\n\t" {
    case "foo bar":
        var v6 float64 = 1.5
        _ = v6
    default:
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int64 = (v2 & v2)
            var v9 float32 = ((0.0 - (0.0 / 1.5)) + (0.0 - float32(v5)))
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v10 string = "hello"
    switch -1 {
    case -9223372036854775808:
        var v11 uint8 = 1
        _ = v11
    default:
        var v12 int = -9223372036854775808
        _ = v12
    }
    switch v10 {
    case "a":
        delete(v4, (v10 + v10))
    default:
        var v13 int32 = v1
        _ = v13
    }
    v2 = v2
    switch 0 {
    case 42:
        for fi14 := 0; fi14 < 16; fi14++ {
            if true {
                var v15 uint64 = 0
                _ = v15
            }
            fmt.Println(v3, v1, v1)
            _ = fi14
        }
    case 472:
        if false {
            if (v2 >= 0) {
                v16 := GenS1{F0: 2147483647, F1: 700, F2: 1}
                v17 := GenS1{F0: 0, F1: 0, F2: 2}
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 int = (-(fi18))
                    _ = fi18
                    _ = v19
                }
                _ = v16
                _ = v17
            }
        }
    case -391:
        var v20 uint = 18446744073709551615
        _ = v20
    default:
        v21 := GenS1{F0: 9223372036854775807, F1: 42, F2: 374}
        _ = v21
    }
    v2 = v2
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int16 = (v3 << ((1) & 63))
        var v24 uint64 = 18446744073709551615
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := v4[v10]
    if v25 {
        v4[("a" + v10)] = true
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
