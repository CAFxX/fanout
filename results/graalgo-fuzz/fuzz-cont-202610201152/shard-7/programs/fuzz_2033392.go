package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint, p1 uint32) uint8 {
    fmt.Println(p1, p1)
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            p1 -= p1
            if true {
                v8 := GenS1{F0: 40, F1: 959}
                var v9 int8 = -124
                v10 := GenS2{F0: (v9 << ((p0) & 63))}
                if true {
                    switch (^(fi7)) {
                    case 128:
                        var v11 int = 1
                        _ = v11
                    }
                }
                _ = v10
                _ = v8
                _ = v9
            } else {
                v12 := []int{-126, -24}
                _ = v12
            }
            _ = fi7
        }
        fmt.Println(fi6)
        _ = fi6
    }
    v13 := GenS2{F0: 0}
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint = p0
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    return 127
}

func main() {
    var v1 int64 = -1
    var v2 bool = (18446744073709551615 == 4294967295)
    v3 := []int{}
    var v4 int = -1
    var v5 string = ("hello" + "a")
    for fi16 := 0; fi16 < 6; fi16++ {
        v2 = (!false)
        if len(v3) > 0 {
            v3[0] = (v4 * (v4 | fi16))
        }
        _ = fi16
    }
    fmt.Println(v4, v2, v2)
    v17 := GenS1{F0: 0, F1: 844}
    v18 := GenS1{F0: 0, F1: 255}
    var v19 string = v5
    var v20 uint16 = 1
    v19 = v19
    switch 9223372036854775807 {
    case 32767:
        var v21 int = v4
        _ = v21
    case 0:
        v19 = v5
    default:
        switch v5 {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = v4
            }
        default:
            if (!(72 >= 1)) {
                v5 = v19
                v1 -= (v1 - (v1 * -1))
                for fi22 := 0; fi22 < 3; fi22++ {
                    v3 = append(v3, (v4 & (v4 >> ((433) & 63))))
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 int32 = 186
                        _ = fi23
                        _ = v24
                    }
                    var v25 string = ""
                    _ = fi22
                    _ = v25
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
