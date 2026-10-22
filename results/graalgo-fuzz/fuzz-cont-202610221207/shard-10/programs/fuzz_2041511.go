package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 float32 = -2.25
    v2 := []string{"a", "", "\n\t", ""}
    var v3 int32 = 1
    var v4 uint32 = 1
    v5 := make(map[uint32]bool)
    if (false || true) {
        var v6 float64 = 3.14159
        fmt.Println(v3, v4, v6)
        if true {
            v1 += v1
            switch ("" + "a") {
            case "foo bar":
                v7 := GenS1{F0: 1, F1: 652, F2: 323}
                _ = v7
            case "":
                var v8 uint8 = 127
                _ = v8
            case "a":
                v9 := GenS1{F0: 1, F1: 2, F2: 0}
                _ = v9
            default:
                var v10 int16 = 739
                _ = v10
            }
        }
        _ = v6
    }
    v1 += (v1 * 0.0)
    fmt.Println(v4, v1)
    var v11 uint16 = 255
    v12 := GenS2{F0: 42, F1: 2}
    var v13 uint16 = (v11 - (v11 ^ v11))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v4, v11)
    if true {
        fmt.Println(v1, v3)
    } else {
        v14 := GenS2{F0: 2, F1: 18446744073709551615}
        _ = v14
    }
    v13 = uint16(v4)
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: 2, F1: v11, F2: 398}
        _ = fi15
        _ = v16
    }
    var v17 uint16 = v11
    for fi18 := 0; fi18 < 16; fi18++ {
        v3 -= (v3 | (-(v3)))
        _ = fi18
    }
    fmt.Println(v17, v11)
    v19 := GenS2{F0: 18446744073709551615, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
