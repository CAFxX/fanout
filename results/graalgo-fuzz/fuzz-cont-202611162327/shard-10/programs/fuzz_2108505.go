package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0(p0 int64, p1 int32, p2 int8) uint32 {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float32 = 1.5
        _ = fi7
        _ = v8
    }
    var v9 int32 = (p1 % (((p1 << ((298) & 63))) | 1))
    v10 := []int{-39}
    var v11 bool = (!true)
    if (42 == (v9 * v9)) {
        var v12 int64 = (p0 % (((p0 & p0)) | 1))
        var v13 string = "a"
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 2147483647
}
func genFn1() int {
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: 42, F1: 453}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int = (-(fi16))
        v17 -= v17
        _ = fi16
        _ = v17
    }
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint32 = 404
    var v3 int8 = 63
    var v4 uint16 = 903
    var v5 string = "\n\t"
    var v6 uint16 = 356
    v18 := []int{32768, 2147483647}
    fmt.Println(v4, v4, v5)
    if true {
        v18 = append(v18, 4294967295)
        switch 4294967295 {
        case 32768:
            for fi19 := 0; fi19 < 1; fi19++ {
                v3 += (v3 % ((0) | 1))
                for fi20 := 0; fi20 < 8; fi20++ {
                    v18 = append(v18, genFn1())
                    v4 += 0
                    _ = fi20
                }
                _ = fi19
            }
        case 4294967295:
            switch 32767 {
            case -119:
                var v21 int8 = (v3 - v3)
                _ = v21
            }
        case -735:
            switch v5 {
            case "hello":
                for fi22 := 0; fi22 < 1; fi22++ {
                    v23 := GenS1{F0: 42, F1: 388}
                    var v24 float64 = -2.25
                    if len(v18) > 0 {
                        v18[0] = -54
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            default:
                v1 = (v1 >> ((1) & 63))
            }
        default:
            v5 = "a"
        }
        fmt.Println(v6)
    }
    var v25 bool = true
    for fi26 := 0; fi26 < 16; fi26++ {
        v5 = ("foo bar" + "")
        _ = fi26
    }
    if true {
        v27 := GenS1{F0: -1, F1: 172}
        _ = v27
    } else {
        if (v3 > -16) {
            v25 = v25
            for fi28 := 0; fi28 < 18; fi28++ {
                v4 -= 32767
                _ = fi28
            }
        } else {
            for fi29 := 0; fi29 < 11; fi29++ {
                v1 = v1
                _ = fi29
            }
            v1 = v1
            var v30 uint64 = 2
            _ = v30
        }
    }
    var v31 uint8 = 1
    var v32 float32 = (float32((3.14159 + 1.5)) - (0.1 + float32(0.0)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
