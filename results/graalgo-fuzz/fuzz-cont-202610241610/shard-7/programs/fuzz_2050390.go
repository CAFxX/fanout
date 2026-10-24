package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint8) int16 {
    var v4 uint16 = 42
    v5 := []int{}
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    fmt.Println(v4, v4)
    var v6 int32 = 594
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint16 = 701
    var v2 int8 = 0
    var v3 uint16 = (^(v1))
    fmt.Println(v2, v3, v2)
    v7 := GenS1{F0: 42}
    fmt.Println(v1)
    v1 = 65535
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int32 = -656
        _ = fi8
        _ = v9
    }
    v3 += 594
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int = (fi10 - (^(fi10)))
        switch "hello" {
        case "a":
            var v12 int8 = (v2 >> ((2) & 63))
            _ = v12
        case "\n\t":
            var v13 uint64 = 484
            _ = v13
        default:
            v11 += (v11 * v11)
        }
        _ = fi10
        _ = v11
    }
    v14 := []int{-554, 32768, 32767}
    var v15 float64 = -2.25
    v2 = (v2 & -100)
    v3 = (v1 * (^(v1)))
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 int64 = -540
        v18 := GenS1{F0: 1}
        if (4294967295 != 256) {
            v17 = v17
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    if len(v14) > 0 {
        v14[0] = 128
    }
    v15 = 0.0
    var v19 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v7
}
