package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 bool }
type GenS2 struct { F0 uint16 }
func genFn0() int {
    var v7 bool = true
    var v8 uint8 = 1
    var v9 bool = ((1 != -463) || (151 == 171))
    v10 := GenS1{F0: 42, F1: 624, F2: ("\n\t" >= "foo bar")}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 uint32 = 42
    var v2 float32 = ((0.0 + (1.5 / 1.0)) / (3.14159 * 0.1))
    var v3 uint32 = (^(v1))
    var v4 int64 = 522
    var v5 float32 = 1.5
    var v6 int = 2147483647
    v11 := []int{(v6 | (v6 ^ 9223372036854775807))}
    v5 -= v5
    if false {
        var v12 bool = (9223372036854775807 < v4)
        v13 := GenS1{F0: 1000, F1: v1, F2: v12}
        _ = v12
        _ = v13
    } else {
        fmt.Println(v2, v2)
    }
    var v14 uint64 = 1
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 uint32 = v3
        v17 := GenS2{F0: 255}
        fmt.Println(v3, v4, v5)
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float64 = (1.5 - 3.14159)
    v18 -= v18
    if len(v11) > 0 {
        v11[0] = v6
    }
    fmt.Println(v3)
    for fi19 := 0; fi19 < 13; fi19++ {
        v11 = append(v11, genFn0())
        _ = fi19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        switch v6 {
        case 42:
            var v21 float32 = v2
            _ = v21
        case -285:
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 float32 = v2
                v24 := GenS2{F0: 1}
                if len(v11) > 0 {
                    v11[0] = fi20
                }
                _ = fi22
                _ = v23
                _ = v24
            }
        case 9223372036854775807:
            var v25 uint64 = 0
            _ = v25
        default:
            if (!false) {
                if false {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        var v27 string = ""
                        _ = fi26
                        _ = v27
                    }
                    v6 += -628
                }
                if true {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                }
            } else {
                v5 -= (v2 + ((-2.25 / (0.1 / v2)) + v5))
            }
        }
        v5 -= float32(v18)
        _ = fi20
    }
    v6 -= v6
    var v28 float64 = v18
    var v29 uint16 = 1
    var v30 uint16 = (v29 & v29)
    var v31 float64 = v18
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
