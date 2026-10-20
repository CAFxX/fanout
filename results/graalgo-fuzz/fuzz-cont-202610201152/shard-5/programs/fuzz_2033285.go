package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 float64 = ((float64(3.14159) + -2.25) * (1e10 * 3.14159))
    v2 := []uint32{4294967295, 286, 606, 343}
    var v3 int16 = -32768
    if ("" > ("\n\t" + (("a" + "") + "a"))) {
        var v4 int32 = 2147483647
        _ = v4
    } else {
        fmt.Println(v3)
    }
    v1 = v1
    v3 = (-(v3))
    if (1 >= 42) {
        if true {
            for fi5 := 0; fi5 < 16; fi5++ {
                if len(v2) > 0 {
                    v2[0] = 2
                }
                _ = fi5
            }
            fmt.Println(v3, v1)
            v6 := GenS1{F0: 255}
            _ = v6
        } else {
            var v7 uint64 = 2
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v1 = v1
        _ = fi8
    }
    switch -462 {
    case 32768:
        var v9 int = 32768
        _ = v9
    case 128:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -999:
        var v10 uint16 = 0
        _ = v10
    default:
        v2 = append(v2, 713)
    }
    v1 -= v1
    if len(v2) > 0 {
        v2[0] = 42
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 uint64 = 1
        _ = fi11
        _ = v12
    }
    v3 = v3
    var v13 string = "\n\t"
    for fi14 := 0; fi14 < 7; fi14++ {
        switch fi14 {
        case 128:
            var v15 int64 = 42
            _ = v15
        }
        v3 = (v3 << ((42) & 63))
        var v16 string = (v13 + (v13 + v13))
        _ = fi14
        _ = v16
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
}
