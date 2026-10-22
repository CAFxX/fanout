package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64, p1 uint16) string {
    v6 := GenS1{F0: p1, F1: (p1 * p1)}
    if (!true) {
        p0 = 2
        switch "foo bar" {
        case "":
            var v7 bool = (float32(-2.25) == 1e10)
            _ = v7
        case "\n\t":
            var v8 int16 = 1000
            _ = v8
        case "foo bar":
            var v9 uint = 841
            _ = v9
        default:
            p1 -= p1
        }
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := []int{(-(fi10)), fi10, (fi10 ^ (fi10 ^ (fi10 | fi10))), -262}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        p1 -= p1
        for fi13 := 0; fi13 < 6; fi13++ {
            if false {
                v14 := GenS1{F0: p1, F1: p1}
                _ = v14
            } else {
                v15 := GenS1{F0: 110, F1: 573}
                fmt.Println(p1, fi12, p1)
                _ = v15
            }
            p1 = 255
            var v16 uint = 255
            _ = fi13
            _ = v16
        }
        var v17 bool = (9223372036854775807 != -466)
        _ = fi12
        _ = v17
    }
    p1 = (p1 >> ((255) & 63))
    _ = p0
    _ = p1
    _ = v6
    return "foo bar"
}

func main() {
    var v1 uint32 = 858
    var v2 int = 42
    var v3 uint8 = 0
    v4 := []string{"a"}
    var v5 uint32 = 2147483647
    v4 = append(v4, ("hello" + "foo bar"))
    var v18 float32 = ((1.5 - (-2.25 + (0.0 - (((0.1 - (0.1 / ((0.0 + 0.0) - 1.5))) * 1e10) + float32(0.0))))) - float32(0.1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v19 := GenS1{F0: 2, F1: 889}
    fmt.Println(v5, v2)
    v5 -= v5
    switch (v2 << ((0) & 63)) {
    case 0:
        var v20 int32 = 42
        _ = v20
    case 128:
        var v21 int16 = -874
        _ = v21
    }
    _ = len(v4)
    v4 = append(v4, "a")
    if false {
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 uint = 506
            _ = fi22
            _ = v23
        }
        v1 = (v5 << ((255) & 63))
        fmt.Println(v18, v3, v3)
    }
    var v24 bool = (65535 != 255)
    if (1e10 == -2.25) {
        var v25 uint = 1
        _ = v25
    } else {
        fmt.Println(v24, v18)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
