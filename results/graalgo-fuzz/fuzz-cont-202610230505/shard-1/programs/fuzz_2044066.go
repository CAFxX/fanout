package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint; F1 bool; F2 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 255
    var v3 int16 = 42
    v4 := make(map[uint32]int)
    var v5 int8 = 0
    v6 := []string{"a", (("" + "foo bar") + "hello"), "foo bar"}
    switch 1 {
    case 123:
        if len(v6) > 0 {
            v6[0] = ("" + "")
        }
    case -2147483648:
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                if true {
                    var v9 uint16 = 1
                    v10 := GenS1{F0: 655, F1: 724}
                    fmt.Println(v2)
                    _ = v10
                    _ = v9
                }
                v11 := GenS1{F0: uint32(v2), F1: 1}
                _ = fi8
                _ = v11
            }
            v6 = append(v6, ("" + ("foo bar" + "foo bar")))
            v2 = (v2 + 203)
            _ = fi7
        }
    default:
        fmt.Println(v5, v2, v2)
    }
    if (!((v2 + v2) <= v2)) {
        v2 = (v2 / ((1) | 1))
        v5 = v5
        if true {
            var v12 bool = (!false)
            v3 += (v3 + v3)
            _ = v12
        } else {
            v5 -= v5
        }
    }
    fmt.Println(v3, v2, v5)
    v13 := GenS1{F0: 112, F1: 4294967295}
    for fi14 := 0; fi14 < 7; fi14++ {
        fmt.Println(v3)
        switch -421 {
        case 128:
            fmt.Println(v3)
        default:
            var v15 uint8 = uint8(v2)
            _ = v15
        }
        _ = fi14
    }
    v2 = 4294967295
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v16 string = "hello"
    if (2 >= 0) {
        v17 := GenS2{F0: (v2 % ((2) | 1)), F1: (1.5 <= (-2.25 * float32(1e10))), F2: 569}
        var v18 int32 = -2147483648
        _ = v17
        _ = v18
    }
    fmt.Println(v3, v3)
    for fi19 := 0; fi19 < 7; fi19++ {
        v2 = (v2 % ((2) | 1))
        var v20 int16 = 42
        _ = fi19
        _ = v20
    }
    fmt.Println(v5, v5, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
