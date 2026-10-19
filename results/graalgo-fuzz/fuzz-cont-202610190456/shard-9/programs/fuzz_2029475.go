package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 int = -749
    var v2 string = ""
    var v3 int32 = 1
    v4 := []string{v2, ("foo bar" + v2)}
    v5 := GenS1{F0: 464, F1: 587}
    v1 = 93
    v1 = (v1 >> ((786) & 63))
    if (!false) {
        v2 = (("hello" + ("hello" + v2)) + "\n\t")
        var v6 float32 = -2.25
        _ = v6
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            v2 = "a"
            _ = fi7
        }
        if ((false || false) && true) {
            v3 = 848
            for fi8 := 0; fi8 < 1; fi8++ {
                v4 = append(v4, v2)
                v3 = v3
                for fi9 := 0; fi9 < 10; fi9++ {
                    v10 := GenS1{F0: 194, F1: 0}
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
        }
    }
    var v11 int16 = -450
    var v12 float64 = (0.1 / -2.25)
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        v12 += v12
        v11 = (v11 * (v11 + v11))
        _ = fi13
    }
    switch v1 {
    case 594:
        for fi14 := 0; fi14 < 9; fi14++ {
            for fi15 := 0; fi15 < 1; fi15++ {
                v4 = append(v4, "\n\t")
                var v16 int16 = v11
                _ = fi15
                _ = v16
            }
            if true {
                var v17 uint64 = 1
                var v18 uint16 = 1
                var v19 uint64 = (v17 << ((0) & 63))
                _ = v17
                _ = v18
                _ = v19
            }
            v20 := GenS1{F0: 2147483648, F1: 18446744073709551615}
            _ = fi14
            _ = v20
        }
    case 4294967295:
        var v21 int64 = 9223372036854775807
        _ = v21
    default:
        v11 += 32767
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
