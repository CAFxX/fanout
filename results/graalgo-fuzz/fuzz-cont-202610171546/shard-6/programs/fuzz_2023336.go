package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = -2.25
    v3 := []string{("a" + (((("foo bar" + ("" + "\n\t")) + "hello") + "foo bar") + "\n\t")), "foo bar"}
    var v4 bool = (float64(-2.25) >= v2)
    v4 = (255 < 1)
    var v5 int32 = -467
    switch 9223372036854775807 {
    case 16:
        if (57 < 374) {
            v5 = v5
        } else {
            for fi6 := 0; fi6 < 12; fi6++ {
                v7 := GenS1{F0: 756}
                _ = fi6
                _ = v7
            }
        }
    case -1:
        v5 = 0
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 0}
        for fi10 := 0; fi10 < 5; fi10++ {
            v2 = ((0.0 + 1e10) - v2)
            _ = fi10
        }
        var v11 int32 = (v5 << ((42) & 63))
        _ = fi8
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS1{F0: 18446744073709551615}
            _ = fi13
            _ = v14
        }
        fmt.Println(fi12, v4)
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi12
    }
    var v15 uint16 = 600
    if false {
        delete(v1, 0)
    } else {
        v5 = (v5 >> ((255) & 63))
        v16 := GenS1{F0: 2}
        _ = v16
    }
    v2 = ((((v2 - (3.14159 / -2.25)) - v2) / 3.14159) / 3.14159)
    if v4 {
        switch -171 {
        case 290:
            var v17 float32 = (1.5 * -2.25)
            _ = v17
        case 4294967295:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            var v18 uint16 = v15
            _ = v18
        }
        v1[-332] = 0
    } else {
        v4 = (255 > 255)
        var v19 uint8 = 128
        _ = v19
    }
    _ = len(v3)
    var v20 uint = 1
    v3 = append(v3, "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
