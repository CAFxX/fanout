package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint8 }

func main() {
    var v1 int = 0
    v2 := []uint32{4294967295, 0, 42, 2}
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    v5 := []string{"", ""}
    if false {
        v5 = append(v5, "a")
        var v6 float64 = (float64(-2.25) / 1e10)
        _ = v6
    } else {
        var v7 uint16 = 32767
        if false {
            if (!false) {
                for fi8 := 0; fi8 < 13; fi8++ {
                    v9 := GenS1{F0: 42, F1: v1, F2: 109}
                    switch v1 {
                    case -2147483648:
                        var v10 int8 = -40
                        _ = v10
                    case 78:
                        var v11 int16 = -193
                        _ = v11
                    case 0:
                        var v12 int = 9223372036854775807
                        _ = v12
                    default:
                        v7 = v7
                    }
                    _ = fi8
                    _ = v9
                }
            }
        } else {
            v13 := GenS1{F0: -1, F1: (v1 + v1), F2: 1}
            _ = v13
        }
        _ = v7
    }
    v14 := GenS1{F0: -128, F1: v1, F2: 112}
    for fi15 := 0; fi15 < 6; fi15++ {
        delete(v3, "a")
        _ = fi15
    }
    var v16 uint64 = 940
    var v17 bool = (-1 > 42)
    v17 = true
    v18 := v4[("" + "\n\t")]
    var v19 uint8 = 2
    if v17 {
        for fi20 := 0; fi20 < 16; fi20++ {
            for fi21 := 0; fi21 < 18; fi21++ {
                delete(v3, (("\n\t" + "\n\t") + ("a" + ("" + "a"))))
                switch (^(fi20)) {
                case -93:
                    v1 = fi20
                case 2147483647:
                    v22 := GenS1{F0: -1, F1: (^(v18)), F2: v19}
                    _ = v22
                default:
                    v23 := GenS1{F0: 111, F1: 42, F2: v19}
                    _ = v23
                }
                _ = fi21
            }
            _ = fi20
        }
        v3["hello"] = 174
        fmt.Println(v16, v18, v18)
    } else {
        v24 := GenS1{F0: 55, F1: (v18 + v1), F2: 0}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
