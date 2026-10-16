package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }

func main() {
    var v1 float64 = ((1.5 - 3.14159) - 0.0)
    v2 := make(map[uint32]bool)
    v3 := []string{"foo bar", "foo bar", "a", ("\n\t" + "foo bar")}
    v4 := make(map[int]uint32)
    var v5 uint32 = 990
    var v6 float64 = float64((1e10 + (((1.5 - (((-2.25 * 0.1) - 1e10) * (3.14159 / float32(1.5)))) + (0.1 - ((0.0 / 1.0) / (1.5 / 1.5)))) - 0.0)))
    fmt.Println(v5)
    v7 := GenS1{F0: 0, F1: -9223372036854775808, F2: 0}
    if false {
        v5 += (v5 + (^(v5)))
        fmt.Println(v5, v6)
    }
    var v8 float64 = 0.1
    fmt.Println(v8)
    var v9 uint16 = 65535
    var v10 int8 = 100
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int16 = 599
        fmt.Println(fi11)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v4[(fi13 | fi13)] = 646
        if (4294967295 != 2) {
            v14 := GenS1{F0: 32767, F1: 127, F2: 1}
            _ = v14
        } else {
            if false {
                v9 += v9
                fmt.Println(v9)
                switch ("foo bar" + "a") {
                case "foo bar":
                    v15 := v4[fi13]
                    _ = v15
                }
            }
            delete(v2, (v5 | 828))
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v1 = v6
        if (86 < (v10 & v10)) {
            switch (fi16 & (-(fi16))) {
            case -269:
                fmt.Println(v10, v5)
            case -381:
                delete(v2, v5)
            default:
                delete(v4, fi16)
            }
        } else {
            if true {
                v5 += v5
            }
        }
        _ = fi16
    }
    var v17 uint32 = v5
    if false {
        delete(v4, 42)
    }
    v17 = (v5 / (((v17 * (v5 >> ((901) & 63)))) | 1))
    v3 = append(v3, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
