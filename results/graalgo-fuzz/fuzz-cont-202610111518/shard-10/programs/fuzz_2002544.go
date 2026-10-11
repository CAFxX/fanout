package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64 }

func main() {
    var v1 uint16 = 42
    var v2 uint = 42
    var v3 uint8 = 255
    if (v3 < v3) {
        var v4 string = "foo bar"
        _ = v4
    } else {
        fmt.Println(v2, v2, v1)
        if (18446744073709551615 != v2) {
            fmt.Println(v2, v3, v3)
        }
        if (v1 <= v1) {
            switch 2147483647 {
            case -66:
                var v5 uint = (v2 >> ((uint(v3)) & 63))
                _ = v5
            }
            fmt.Println(v3)
        }
    }
    if (-1 == 4294967295) {
        var v6 int32 = -170
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 int = (fi7 >> ((uint(v1)) & 63))
            _ = fi8
            _ = v9
        }
        v1 += v1
        _ = fi7
    }
    switch "" {
    case "foo bar":
        v2 = (v2 | v2)
    }
    var v10 int16 = -69
    var v11 int16 = (v10 - (v10 ^ (v10 * (v10 + (v10 + -1)))))
    v12 := []int{4294967295}
    v1 = v1
    _ = len(v12)
    v13 := GenS1{F0: 0}
    for fi14 := 0; fi14 < 4; fi14++ {
        v11 -= v11
        _ = fi14
    }
    v2 = uint(v1)
    v11 = v11
    var v15 string = ("foo bar" + "hello")
    var v16 uint = 759
    v10 -= (v10 >> ((278) & 63))
    if true {
        v2 += v2
        var v17 float32 = 1.5
        var v18 uint32 = 2
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 0, F1: int64(v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
