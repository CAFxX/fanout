package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint16 }

func main() {
    var v1 int64 = -517
    var v2 int = -551
    var v3 string = "\n\t"
    v4 := make(map[uint32]int)
    var v5 uint32 = 4294967295
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v3, v3, v5)
    case "hello":
        delete(v4, v5)
    default:
        if false {
            switch (v2 >> ((317) & 63)) {
            case 4294967295:
                v1 -= 42
            case 182:
                var v6 int8 = 46
                _ = v6
            default:
                for fi7 := 0; fi7 < 13; fi7++ {
                    v5 = (v5 * (^(v5)))
                    v5 = v5
                    v8 := GenS1{F0: -1, F1: 149}
                    _ = fi7
                    _ = v8
                }
            }
            v9 := GenS2{F0: 260, F1: 0, F2: 0}
            var v10 uint = 0
            _ = v10
            _ = v9
        } else {
            var v11 int16 = -190
            _ = len(v4)
            _ = v11
        }
    }
    fmt.Println(v2, v5, v5)
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v1, v1, v5)
        _ = fi12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint = 2
        _ = fi13
        _ = v14
    }
    v3 = ""
    for fi15 := 0; fi15 < 8; fi15++ {
        fmt.Println(v2)
        v16 := GenS2{F0: v5, F1: 1, F2: 42}
        _ = fi15
        _ = v16
    }
    var v17 uint8 = 1
    if (float64(3.14159) > float64(-2.25)) {
        v17 -= (^(v17))
    }
    var v18 int = (v2 >> ((65535) & 63))
    var v19 uint = 42
    fmt.Println(v1, v19, v17)
    if true {
        var v20 uint = (v19 >> ((v19) & 63))
        if ((9223372036854775807 > v1) && (3.14159 > 0.1)) {
            v4[v5] = v18
            fmt.Println(v17, v18)
            v21 := []int{(v2 | v2)}
            _ = v21
        }
        _ = v20
    }
    var v22 string = v3
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
