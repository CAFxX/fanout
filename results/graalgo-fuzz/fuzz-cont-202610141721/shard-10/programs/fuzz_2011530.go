package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint16 }
func genFn0(p0 int8, p1 int64) int16 {
    p0 = p0
    var v4 int = 853
    _ = p0
    _ = p1
    _ = v4
    return -430
}

func main() {
    v1 := []string{}
    v2 := []uint32{406}
    v3 := []uint32{}
    if false {
        var v5 uint64 = 18446744073709551615
        _ = v5
    }
    switch 762 {
    case -875:
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                _ = len(v3)
                if len(v3) > 0 {
                    v3[0] = 2
                }
                _ = fi7
            }
            _ = fi6
        }
    case 1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 798:
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int64 = 9223372036854775807
            _ = fi8
            _ = v9
        }
    default:
        var v10 uint32 = 2147483647
        _ = v10
    }
    var v11 float64 = -2.25
    var v12 string = (("\n\t" + ("a" + (("a" + ("" + "foo bar")) + "foo bar"))) + ("hello" + "foo bar"))
    if false {
        for fi13 := 0; fi13 < 19; fi13++ {
            v11 += float64((((3.14159 / -2.25) + float32((((v11 * v11) / v11) - ((((0.0 * v11) + (((1e10 - v11) + (0.0 * -2.25)) - v11)) * 1.5) + v11)))) + 0.1))
            for fi14 := 0; fi14 < 4; fi14++ {
                _ = len(v3)
                v12 = v12
                _ = fi14
            }
            _ = fi13
        }
        var v15 int8 = 88
        v15 = v15
        _ = v15
    } else {
        var v16 uint = 1
        switch (v12 + v12) {
        case "a":
            v17 := GenS1{F0: 0, F1: 32768, F2: 2}
            _ = v17
        case "":
            if len(v3) > 0 {
                v3[0] = 799
            }
        default:
            v18 := GenS1{F0: 32767, F1: 42, F2: 0}
            _ = v18
        }
        _ = v16
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v12 = v12
    fmt.Println(v11, v12)
    v3 = append(v3, 4294967295)
    v11 = ((v11 + v11) / (v11 + (0.0 - v11)))
    var v19 uint64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
}
